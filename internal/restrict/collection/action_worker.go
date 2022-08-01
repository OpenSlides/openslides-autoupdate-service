package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ActionWorker handles permission for action_worker.
type ActionWorker struct{}

// MeetingID returns no meeting.
func (a ActionWorker) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the restrictions modes for the action_worker collection.
func (a ActionWorker) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	}
	return nil
}
