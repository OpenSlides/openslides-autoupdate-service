package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// Poll handels restrictions of the collection poll. If the user can see a poll
// depends on the content_object:
//
//	motion: The user can see the linked motion.
//	assignment: The user can see the linked assignment.
//	topic: The user can see the topic.
//
// If the user can manage the poll also depends on the content_object:
//
//	motion: The user needs motion.can_manage_polls.
//	assignment: The user needs assignment.can_manage_polls.
//	topic: The user needs poll.can_manage.
//
// Mode A: Contains the fields to know, that the poll exists and how it is
// configured. It is allowed, if the user can see the poll.
//
// Mode B: Contains the fields to see the result of a poll. If the poll is
// published, the user has to be able to see the poll. If it is not published,
// he needs the permission to manage the poll.
type Poll struct{}

// Name returns the collection name.
func (p Poll) Name() string {
	return "poll"
}

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
	}
	return nil
}

func (p Poll) see(ctx context.Context, ds *dsfetch.Fetch, pollIDs ...int) ([]int, error) {
	return eachContentObjectCollection(
		ctx,
		ds.Poll_ContentObjectID,
		pollIDs, func(objectCollection string, objectID int, ids []int) ([]int, error) {
			var collection FieldRestricter

			switch objectCollection {
			case "motion":
				collection = Collection(ctx, Motion{}.Name()).Modes("C")

			case "assignment":
				collection = Collection(ctx, Assignment{}.Name()).Modes("A")

			case "topic":
				collection = Collection(ctx, Topic{}.Name()).Modes("A")

			default:
				return nil, fmt.Errorf("unsupported collection: %s", objectCollection)
			}

			see, err := collection(ctx, ds, objectID)
			if err != nil {
				return nil, fmt.Errorf("checking see of content object %d: %w", objectID, err)
			}

			if len(see) == 1 {
				return ids, nil
			}

			return nil, nil
		})
}

func (p Poll) manage(ctx context.Context, ds *dsfetch.Fetch, pollIDs ...int) ([]int, error) {
	return eachContentObjectCollection(ctx, ds.Poll_ContentObjectID, pollIDs, func(objectCollection string, objectID int, ids []int) ([]int, error) {
		switch objectCollection {
		case "motion":
			meetingID, err := ds.Motion_MeetingID(objectID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting meeting id of motion %d: %w", objectID, err)
			}

			perms, err := perm.FromContext(ctx, meetingID)
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

			perms, err := perm.FromContext(ctx, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.AssignmentCanManagePolls) {
				return ids, nil
			}
			return nil, nil

		case "topic":
			meetingID, err := ds.Topic_MeetingID(objectID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting meeting id of topic %d: %w", objectID, err)
			}

			perms, err := perm.FromContext(ctx, meetingID)
			if err != nil {
				return nil, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.PollCanManage) {
				return ids, nil
			}
			return nil, nil

		default:
			return nil, fmt.Errorf("unsupported collection: %s", objectCollection)
		}
	})
}

func (p Poll) modeB(ctx context.Context, ds *dsfetch.Fetch, pollIDs ...int) ([]int, error) {
	published := make([]bool, len(pollIDs))
	for i, pollID := range pollIDs {
		ds.Poll_Published(pollID).Lazy(&published[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll/publish: %w", err)
	}

	allowed := make([]int, 0, len(pollIDs))
	for i, pollID := range pollIDs {
		if published[i] {
			see, err := p.see(ctx, ds, pollID)
			if err != nil {
				return nil, fmt.Errorf("checking see: %w", err)
			}
			if len(see) > 0 {
				allowed = append(allowed, pollID)
			}
			continue
		}

		manage, err := p.manage(ctx, ds, pollID)
		if err != nil {
			return nil, fmt.Errorf("checking manage: %w", err)
		}
		if len(manage) > 0 {
			allowed = append(allowed, pollID)
		}
	}

	return allowed, nil
}
