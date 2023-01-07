package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Topic handels the restrictions for the topic collection.
//
// The user can see a topic, if the user has agenda_item.can_see.
//
// Mode A: The user can see the topic.
type Topic struct {
	name string
}

// Name returns the collection name.
func (t Topic) Name() string {
	return t.name
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

func (t Topic) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, topicIDs ...int) error {
	return meetingPerm(ctx, ds, t, "A", topicIDs, mperms, perm.AgendaItemCanSee, attrMap)
}
