package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Poll handels restrictions of the collection poll.
type Poll struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m Poll) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	case "B":
		return m.modeB
	case "C":
		return m.modeC
	case "D":
		return m.modeD
	}
	return nil
}

func (m Poll) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	contentObjectID, exist, err := ds.Poll_ContentObjectID(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}

	if !exist {
		meetingID, err := ds.Poll_MeetingID(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id of poll %d: %w", pollID, err)
		}

		see, err := Meeting{}.see(ctx, ds, mperms, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking see for meeting %d: %w", meetingID, err)
		}

		return see, nil
	}

	parts := strings.Split(contentObjectID, "/")
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("decoding id of content object id %q: %w", contentObjectID, err)
	}

	switch parts[0] {
	case "motion":
		see, err := Motion{}.see(ctx, ds, mperms, id)
		if err != nil {
			return false, fmt.Errorf("checking see motion %d: %w", id, err)
		}

		return see, nil

	case "assignment":
		see, err := Assignment{}.see(ctx, ds, mperms, id)
		if err != nil {
			return false, fmt.Errorf("checking see assignment %d: %w", id, err)
		}

		return see, nil

	default:
		return false, fmt.Errorf("unsupported collection for poll %d: %s", pollID, parts[0])
	}
}

func (m Poll) manage(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	contentObjectID, exist, err := ds.Poll_ContentObjectID(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}

	if !exist {
		meetingID, err := ds.Poll_MeetingID(pollID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting id of poll %d: %w", pollID, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.PollCanManage), nil
	}

	parts := strings.Split(contentObjectID, "/")
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

	default:
		return false, fmt.Errorf("unsupported collection for poll %d: %s", pollID, parts[0])
	}
}

func (m Poll) modeB(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	switch state {
	case "published":
		see, err := m.see(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see: %w", err)
		}
		return see, nil

	case "finished":
		manage, err := m.manage(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage: %w", err)
		}
		return manage, nil

	default:
		return false, nil

	}
}

func (m Poll) modeC(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	if state != "started" {
		return false, nil
	}

	return m.manage(ctx, ds, mperms, pollID)
}

func (m Poll) modeD(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	switch state {
	case "published":
		see, err := m.see(ctx, ds, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see: %w", err)
		}
		return see, nil

	case "finished":
		manage, err := m.manage(ctx, ds, mperms, pollID)
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
