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
	case "D":
		return m.modeD
	}
	return nil
}

func (m Poll) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	contentObjectID := fetch.Field().Poll_ContentObjectID(ctx, pollID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}

	parts := strings.Split(contentObjectID, "/")

	var id int
	if len(parts) == 2 {
		var err error
		id, err = strconv.Atoi(parts[1])
		if err != nil {
			return false, fmt.Errorf("decoding id of content object id %q: %w", contentObjectID, err)
		}
	}

	switch parts[0] {
	case "motion":
		see, err := Motion{}.see(ctx, fetch, mperms, id)
		if err != nil {
			return false, fmt.Errorf("checking see motion %d: %w", id, err)
		}

		return see, nil

	case "assignment":
		see, err := Assignment{}.see(ctx, fetch, mperms, id)
		if err != nil {
			return false, fmt.Errorf("checking see assignment %d: %w", id, err)
		}

		return see, nil

	default:
		meetingID := fetch.Field().Poll_MeetingID(ctx, pollID)
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("gettin meeting id of poll %d: %w", pollID, err)
		}

		see, err := Meeting{}.see(ctx, fetch, mperms, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking see for meeting %d: %w", meetingID, err)
		}

		return see, nil
	}
}

func (m Poll) manage(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	contentObjectID := fetch.Field().Poll_ContentObjectID(ctx, pollID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}

	parts := strings.Split(contentObjectID, "/")

	var id int
	if len(parts) == 2 {
		var err error
		id, err = strconv.Atoi(parts[1])
		if err != nil {
			return false, fmt.Errorf("decoding id of content object id %q: %w", contentObjectID, err)
		}
	}

	switch parts[0] {
	case "motion":
		meetingID := fetch.Field().Motion_MeetingID(ctx, id)
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("gettin meeting id of motion %d: %w", id, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.MotionCanManagePolls), nil

	case "assignment":
		meetingID := fetch.Field().Assignment_MeetingID(ctx, id)
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("gettin meeting id of assignment %d: %w", id, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.AssignmentCanManage), nil

	default:
		meetingID := fetch.Field().Poll_MeetingID(ctx, pollID)
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("gettin meeting id of poll %d: %w", pollID, err)
		}

		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		return perms.Has(perm.PollCanManage), nil
	}
}

func (m Poll) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state := fetch.Field().Poll_State(ctx, pollID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	switch state {
	case "published":
		see, err := m.see(ctx, fetch, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see: %w", err)
		}
		return see, nil

	case "finished":
		manage, err := m.manage(ctx, fetch, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage: %w", err)
		}
		return manage, nil

	default:
		return false, nil

	}
}

func (m Poll) modeD(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, pollID int) (bool, error) {
	state := fetch.Field().Poll_State(ctx, pollID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	switch state {
	case "published":
		see, err := m.see(ctx, fetch, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking see: %w", err)
		}
		return see, nil

	case "finished":
		manage, err := m.manage(ctx, fetch, mperms, pollID)
		if err != nil {
			return false, fmt.Errorf("checking manage: %w", err)
		}
		if manage {
			return true, nil
		}

		meetingID := fetch.Field().Poll_MeetingID(ctx, pollID)
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("gettin meeting id of poll %d: %w", pollID, err)
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
