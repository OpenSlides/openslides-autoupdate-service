package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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
		return todoToSingle(p.see)
	case "B":
		return todoToSingle(p.modeB)
	case "C":
		return todoToSingle(p.modeC)
	case "D":
		return todoToSingle(p.modeD)
	}
	return nil
}

func (p Poll) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	contentObjectID, err := ds.Poll_ContentObjectID(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}

	parts := strings.Split(contentObjectID, "/")
	if len(parts) != 2 {
		// TODO LAST ERROR
		return false, fmt.Errorf("invalid value for poll/content_object_id: `%s`", contentObjectID)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("decoding id of content object id %q: %w", contentObjectID, err)
	}

	var collection interface {
		see(context.Context, *dsfetch.Fetch, *perm.MeetingPermission, ...int) ([]int, error)
	}

	switch parts[0] {
	case "motion":
		collection = Motion{}

	case "assignment":
		collection = Assignment{}

	case "topic":
		collection = Topic{}

	default:
		// TODO LAST ERROR
		return false, fmt.Errorf("unsupported collection for poll %d: %s", pollID, parts[0])
	}

	see, err := collection.see(ctx, ds, mperms, id)
	if err != nil {
		return false, fmt.Errorf("checking see of content objet %d: %w", id, err)
	}

	return len(see) > 0, nil
}

func (p Poll) manage(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	contentObjectID, err := ds.Poll_ContentObjectID(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}

	parts := strings.Split(contentObjectID, "/")
	if len(parts) != 2 {
		// TODO LAST ERROR
		return false, fmt.Errorf("invalid value for poll/content_object_id: `%s`", contentObjectID)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("decoding id of content object id %q: %w", contentObjectID, err)
	}

	switch parts[0] {
	case "motion":
		meetingID, err := ds.Motion_MeetingID(id).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id of motion %d: %w", id, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.MotionCanManagePolls), nil

	case "assignment":
		meetingID, err := ds.Assignment_MeetingID(id).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id of assignment %d: %w", id, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.AssignmentCanManage), nil

	case "topic":
		meetingID, err := ds.Topic_MeetingID(id).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id of topic %d: %w", id, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.PollCanManage), nil

	default:
		// TODO LAST ERROR
		return false, fmt.Errorf("unsupported collection for poll %d: %s", pollID, parts[0])
	}
}

func (p Poll) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	switch state {
	case "published":
		see, err := p.see(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see: %w", err)
		}
		return see, nil

	case "finished":
		manage, err := p.manage(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage: %w", err)
		}
		return manage, nil

	default:
		return false, nil

	}
}

func (p Poll) modeC(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	if state != "started" {
		return false, nil
	}

	return p.manage(ctx, ds, mperms, pollID)
}

func (p Poll) modeD(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	switch state {
	case "published":
		see, err := p.see(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see: %w", err)
		}
		return see, nil

	case "finished":
		manage, err := p.manage(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage: %w", err)
		}
		if manage {
			return true, nil
		}

		meetingID, err := ds.Poll_MeetingID(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id of poll %d: %w", pollID, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.ListOfSpeakersCanManage), nil

	default:
		return false, nil

	}
}
