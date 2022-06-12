package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Poll handels restrictions of the collection poll.
// If the user can see a poll depends on the content object:
//     motion: The user can see the linked motion.
//     assignment: The user can see the linked assignment.
//     topic: The user can see the topic.
//
// If the user can manage the poll depends on the content object:
//     motion: The user needs motion.can_manage_polls.
//     assignment: The user needs assignment.can_manage.
//     topic: The user needs poll.can_manage.
//
// Mode A: The user can see the poll.
//
// Mode B: Depends on poll/state:
//     published: Accessible if the user can see the poll.
//     finished: Accessible if the user can manage the poll.
//     others: Not accessible for anyone.
//
// Mode C: The user can manage the poll and it is in the started state.
//
// Mode D: Same as Mode B, but for `finished`: Accessible if the user can manage the poll or the user has list_of_speakers.can_manage.
type Poll struct{}

// MeetingID returns the meetingID for the object.
func (p Poll) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Poll_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p Poll) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	case "B":
		return p.modeB
	case "C":
		return p.modeC
	case "D":
		return p.modeD
	}
	return nil
}

func (p Poll) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollIDs ...int) ([]int, error) {
	return eachContentObjectCollection(ctx, ds.Poll_ContentObjectID, pollIDs, func(objectCollection string, objectID int, ids []int) ([]int, error) {
		var collection interface {
			see(context.Context, *dsfetch.Fetch, *perm.MeetingPermission, ...int) ([]int, error)
		}

		switch objectCollection {
		case "motion":
			collection = Motion{}

		case "assignment":
			collection = Assignment{}

		case "topic":
			collection = Topic{}

		default:
			// TODO LAST ERROR
			return nil, fmt.Errorf("unsupported collection: %s", objectCollection)
		}

		see, err := collection.see(ctx, ds, mperms, objectID)
		if err != nil {
			return nil, fmt.Errorf("checking see of content object %d: %w", objectID, err)
		}

		if len(see) == 1 {
			return ids, nil
		}

		return nil, nil
	})
}

func (p Poll) manage(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollIDs ...int) ([]int, error) {
	return eachContentObjectCollection(ctx, ds.Poll_ContentObjectID, pollIDs, func(objectCollection string, objectID int, ids []int) ([]int, error) {
		switch objectCollection {
		case "motion":
			meetingID, err := ds.Motion_MeetingID(objectID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting meeting id of motion %d: %w", objectID, err)
			}

			perms, err := mperms.Meeting(ctx, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.MotionCanManagePolls) {
				return ids, nil
			}

			return nil, nil

		case "assignment":
			meetingID, err := ds.Assignment_MeetingID(objectID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting meeting id of assignment %d: %w", objectID, err)
			}

			perms, err := mperms.Meeting(ctx, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.AssignmentCanManage) {
				return ids, nil
			}
			return nil, nil

		case "topic":
			meetingID, err := ds.Topic_MeetingID(objectID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting meeting id of topic %d: %w", objectID, err)
			}

			perms, err := mperms.Meeting(ctx, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.PollCanManage) {
				return ids, nil
			}
			return nil, nil

		default:
			// TODO LAST ERROR
			return nil, fmt.Errorf("unsupported collection: %s", objectCollection)
		}
	})
}

func (p Poll) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollIDs ...int) ([]int, error) {
	return eachStringField(ctx, ds.Poll_State, pollIDs, func(state string, ids []int) ([]int, error) {
		switch state {
		case "published":
			see, err := p.see(ctx, ds, mperms, ids...)
			if err != nil {
				return nil, fmt.Errorf("checking see: %w", err)
			}
			return see, nil

		case "finished":
			manage, err := p.manage(ctx, ds, mperms, ids...)
			if err != nil {
				return nil, fmt.Errorf("checking manage: %w", err)
			}
			return manage, nil

		default:
			return nil, nil
		}
	})
}

func (p Poll) modeC(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollIDs ...int) ([]int, error) {
	return eachStringField(ctx, ds.Poll_State, pollIDs, func(state string, ids []int) ([]int, error) {
		if state != "started" {
			return nil, nil
		}

		return p.manage(ctx, ds, mperms, ids...)
	})
}

func (p Poll) modeD(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollIDs ...int) ([]int, error) {
	return eachStringField(ctx, ds.Poll_State, pollIDs, func(state string, pollIDs []int) ([]int, error) {
		switch state {
		case "published":
			see, err := p.see(ctx, ds, mperms, pollIDs...)
			if err != nil {
				return nil, fmt.Errorf("checking see: %w", err)
			}
			return see, nil

		case "finished":
			allowed, err := p.manage(ctx, ds, mperms, pollIDs...)
			if err != nil {
				return nil, fmt.Errorf("checking manage: %w", err)
			}

			if len(allowed) == len(pollIDs) {
				return allowed, nil
			}

			notAllowed := set.New(pollIDs...)
			notAllowed.Remove(allowed...)

			allowed2, err := meetingPerm(ctx, ds, p, notAllowed.List(), mperms, perm.ListOfSpeakersCanManage)
			if err != nil {
				return nil, fmt.Errorf("checking list of speaker permission: %w", err)
			}

			return append(allowed, allowed2...), nil

		default:
			return nil, nil
		}
	})
}
