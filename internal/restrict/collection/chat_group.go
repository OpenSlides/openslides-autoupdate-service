package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ChatGroup handels restrictions for the collection chat_group.
//
// A user can see a chat group if any of:
//     The user has the permission chat.can_manage in the respective meeting (dedicated by the key meeting_id).
//     The user is assigned to groups in common with chat_group/read_group_ids.
//     The user is assigned to groups in common with chat_group/write_group_ids.
//
// Mode A: The user can see the chat_group.
type ChatGroup struct{}

// MeetingID returns the meetingID for the object.
func (c ChatGroup) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := c.meetingID(ctx, ds, id)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes give sthe FieldRestricter for a restriction_mode.
func (c ChatGroup) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return c.see
	}
	return nil
}

func (c ChatGroup) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, chatGroupID int) (bool, error) {
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

func (c ChatGroup) meetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, error) {
	mid, err := ds.ChatGroup_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("fetching meeting_id: %w", err)
	}
	return mid, nil
}
