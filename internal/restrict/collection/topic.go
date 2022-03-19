package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Topic handels the restrictions for the topic collection.
//
// The user can see a topic, if the user has agenda_item.can_see.
//
// Mode A: The user can see the topic.
type Topic struct{}

// MeetingID returns the meetingID for the object.
func (t Topic) MeetingID(ctx context.Context, ds *datastore.Request, id int) (int, bool, error) {
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

func (t Topic) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, topicID int) (bool, error) {
	meetingID, err := ds.Topic_MeetingID(topicID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting_id %d: %w", topicID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	return perms.Has(perm.AgendaItemCanSee), nil
}
