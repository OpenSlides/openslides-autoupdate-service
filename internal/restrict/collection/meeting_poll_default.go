package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// MeetingPollDefault handels restrictions of the collection meeting poll default.
//
// # The user can see the meeting poll defaults if the user can see the meeting
//
// Mode A: The user can see the meeting.
type MeetingPollDefault struct{}

// Name returns the collection name.
func (m MeetingPollDefault) Name() string {
	return "meeting_poll_default"
}

// MeetingID returns the meetingID for the object.
func (m MeetingPollDefault) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return id, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MeetingPollDefault) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MeetingPollDefault) see(ctx context.Context, ds *dsfetch.Fetch, meetingPollDefaultIDs ...int) ([]int, error) {
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
