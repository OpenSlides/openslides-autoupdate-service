package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// ChatGroup handels restrictions for the collection chat_group.
type ChatGroup struct{}

// Modes give sthe FieldRestricter for a restriction_mode.
func (c ChatGroup) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return c.see
	}
	return nil
}

func (c ChatGroup) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, chatGroupID int) (bool, error) {
	meetingID, err := c.meetingID(ctx, fetch, chatGroupID)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	adminGroup := fetch.Field().Meeting_AdminGroupID(ctx, meetingID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting admin group id: %w", err)
	}

	if perms.InGroup(adminGroup) {
		return true, nil
	}

	readGroups := fetch.Field().ChatGroup_ReadGroupIDs(ctx, chatGroupID)
	for _, gid := range readGroups {
		if perms.InGroup(gid) {
			return true, nil
		}
	}

	return false, nil
}

func (c ChatGroup) meetingID(ctx context.Context, fetch *datastore.Fetcher, id int) (int, error) {
	mid := fetch.Field().ChatGroup_MeetingID(ctx, id)
	if err := fetch.Err(); err != nil {
		return 0, fmt.Errorf("fetching meeting_id for chat_group %d: %w", id, err)
	}
	return mid, nil
}
