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
//
//	The user has the permission chat.can_manage in the respective meeting (dedicated by the key meeting_id).
//	The user is assigned to groups in common with chat_group/read_group_ids.
//	The user is assigned to groups in common with chat_group/write_group_ids.
//
// Mode A: The user can see the chat_group.
type ChatGroup struct{}

// Name ChatGroup the collection name.
func (c ChatGroup) Name() string {
	return "chat_group"
}

// MeetingID returns the meetingID for the object.
func (c ChatGroup) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := ds.ChatGroup_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting_id: %w", err)
	}
	return mid, true, nil
}

// Modes give sthe FieldRestricter for a restriction_mode.
func (c ChatGroup) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return c.see
	}
	return nil
}

func (c ChatGroup) see(ctx context.Context, ds *dsfetch.Fetch, chatGroupIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, c, chatGroupIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.ChatCanManage) {
			return ids, nil
		}

		allowed, err := eachCondition(ids, func(chatGroupID int) (bool, error) {
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
		})
		if err != nil {
			return nil, fmt.Errorf("checking if user is in read or write group: %w", err)
		}

		return allowed, nil
	})
}
