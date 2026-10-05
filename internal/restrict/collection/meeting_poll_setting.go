package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// MeetingPollSetting handels restrictions of the collection meeting poll default.
//
// # The user can see the meeting poll defaults if the user can see the meeting
//
// Mode A: The user can see the meeting.
type MeetingPollSetting struct{}

// Name returns the collection name.
func (m MeetingPollSetting) Name() string {
	return "meeting_poll_setting"
}

// MeetingID returns the meetingID for the object.
func (m MeetingPollSetting) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := ds.MeetingPollSetting_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting_id: %w", err)
	}
	return mid, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MeetingPollSetting) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MeetingPollSetting) see(ctx context.Context, ds *dsfetch.Fetch, meetingPollDefaultIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, m, meetingPollDefaultIDs, func(meetingID int, ids []int) ([]int, error) {
		canSee, err := Collection(ctx, Meeting{}.Name()).Modes("B")(ctx, ds, meetingID)
		if err != nil {
			return nil, fmt.Errorf("can see meeting %d: %w", meetingID, err)
		}

		if len(canSee) == 1 {
			return ids, nil
		}
		return nil, nil
	})
}
