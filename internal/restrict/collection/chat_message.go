package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// ChatMessage handels restrictions for the collection chat_message.
type ChatMessage struct{}

// Modes give sthe FieldRestricter for a restriction_mode.
func (c ChatMessage) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return c.see
	}
	return nil
}

func (ChatMessage) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, chatGroupID int) (bool, error) {
	chatGroupID, err := ds.ChatMessage_ChatGroupID(chatGroupID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting chat_group_id: %w", err)
	}

	return ChatGroup{}.see(ctx, ds, mperms, chatGroupID)
}
