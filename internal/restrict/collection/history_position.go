package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// HistoryPosition handles restrictions of the collection history_position.
//
// A user can see the position, if he can see one of its entries.
//
// Mode A: The user can see the the history position.
type HistoryPosition struct{}

// Name returns the collection name.
func (h HistoryPosition) Name() string {
	return "history_position"
}

// MeetingID returns false since a HistoryPosition has no meeting.
func (h HistoryPosition) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the field restricters for the collection.
func (h HistoryPosition) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return h.see
	}
	return nil
}

func (h HistoryPosition) see(ctx context.Context, ds *dsfetch.Fetch, historyPositionIDs ...int) ([]int, error) {
	ids := make([][]int, len(historyPositionIDs))
	for i, id := range historyPositionIDs {
		ds.HistoryPosition_EntryIDs(id).Lazy(&ids[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching history: %w", err)
	}

	allowed := make([]int, 0, len(historyPositionIDs))
	for i, id := range historyPositionIDs {
		allowedEntries, err := Collection(ctx, HistoryEntry{}.Name()).Modes("A")(ctx, ds, ids[i]...)
		if err != nil {
			return nil, fmt.Errorf("checking history entries: %w", err)
		}

		if len(allowedEntries) > 0 {
			allowed = append(allowed, id)
		}
	}
	return allowed, nil
}
