package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// HistoryEntry handles restrictions of the collection history_entry.
//
// The user can see the related history position.
//
// Mode A: The user can see the the history entry.
type HistoryEntry struct{}

// Name returns the collection name.
func (h HistoryEntry) Name() string {
	return "history_entry"
}

// MeetingID returns false since a HistoryEntry has no meeting.
func (h HistoryEntry) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the field restricters for the collection.
func (h HistoryEntry) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return h.see
	}
	return nil
}

func (h HistoryEntry) see(ctx context.Context, ds *dsfetch.Fetch, ids ...int) ([]int, error) {
	return eachRelationField(ctx, ds.HistoryEntry_PositionID, ids, func(positionID int, ids []int) ([]int, error) {
		allowed, err := Collection(ctx, HistoryPosition{}.Name()).Modes("A")(ctx, ds, positionID)
		if err != nil {
			return nil, fmt.Errorf("checking history position: %w", err)
		}
		if len(allowed) > 0 {
			return ids, nil
		}
		return nil, nil
	})
}
