package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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

func (c ChatGroup) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, chatGroupID int) (bool, error) {
	meetingID, err := c.meetingID(ctx, ds, chatGroupID)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if perms.Has(perm.ChatCanManage) {
		return true, nil
	}

	readGroups, err := ds.ChatGroup_ReadGroupIDs(chatGroupID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting chat read group ids: %w", err)
	}

	writeGroups, err := ds.ChatGroup_WriteGroupIDs(chatGroupID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting chat read group ids: %w", err)
	}

	allGroups := append(readGroups, writeGroups...)

	for _, gid := range allGroups {
		if perms.InGroup(gid) {
			return true, nil
		}
	}

	return false, nil
}

func (c ChatGroup) meetingID(ctx context.Context, ds *datastore.Request, id int) (int, error) {
	mid, err := ds.ChatGroup_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching meeting_id for chat_group %d: %w", id, err)
	}
	return mid, nil
}
