package collection

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// Gender handles permission for action_worker.
//
// Everyone can see all genders.
type Gender struct{}

// Name returns the collection name.
func (a Gender) Name() string {
	return "gender"
}

// MeetingID returns no meeting.
func (a Gender) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the restrictions modes for the action_worker collection.
func (a Gender) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	}
	return nil
}
