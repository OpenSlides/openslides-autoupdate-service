package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/set"
)

// Topic handels the restrictions for the topic collection.
//
// The user can see a topic, if the user can see the linked topic.
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

func (t Topic) see(ctx context.Context, ds *dsfetch.Fetch, topicIDs ...int) ([]int, error) {
	agendaIDs := make([]int, len(topicIDs))
	for i, tid := range topicIDs {
		ds.Topic_AgendaItemID(tid).Lazy(&agendaIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting agenda ids: %w", err)
	}

	allowedAgendaIDs, err := Collection(ctx, AgendaItem{}.Name()).Modes("A")(ctx, ds, agendaIDs...)
	if err != nil {
		return nil, fmt.Errorf("checking agenda permission: %w", err)
	}

	if len(allowedAgendaIDs) == len(topicIDs) {
		return topicIDs, nil
	}

	allowedAgendaSet := set.New(allowedAgendaIDs...)

	var allowed []int
	for i, tid := range topicIDs {
		if allowedAgendaSet.Has(agendaIDs[i]) {
			allowed = append(allowed, tid)
		}
	}

	return allowed, nil
}
