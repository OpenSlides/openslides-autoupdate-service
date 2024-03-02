package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Topic handels the restrictions for the topic collection.
//
// The user can see a topic, if the user can see the linked agenda_item.
//
// Mode A: The user can see the topic.
type Topic struct{}

// Name returns the collection name.
func (t Topic) Name() string {
	return "topic"
}

// MeetingID returns the meetingID for the object.
func (t Topic) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Topic_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the field restriction for each mode.
func (t Topic) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return t.see
	}
	return nil
}

func (t Topic) see(ctx context.Context, fetcher *dsfetch.Fetch, topicIDs []int) ([]attribute.Func, error) {
	return canSeeRelatedCollection(ctx, fetcher, fetcher.Topic_AgendaItemID, Collection(ctx, AgendaItem{}).Modes("A"), topicIDs)
}
