package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// ChatGroup handels restrictions for the collection chat_group.
//
// A user can see a chat group if any of:
//
//	The user has the permission chat.can_manage.
//	The user is assigned to groups in common with chat_group/read_group_ids or chat_group/write_group_ids.
//
// Mode A: The user can see the chat_group.
type ChatGroup struct{}

// Name returns the collection name.
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

func (c ChatGroup) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, chatGroupIDs ...int) error {
	return eachMeeting(ctx, ds, c, chatGroupIDs, func(meetingID int, ids []int) error {
		groupMap, err := mperms.Meeting(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("getting permissions: %w", err)
		}

		manageGroups := groupMap[perm.ChatCanManage].List()

		for _, chatGroupID := range ids {
			readGroups, err := ds.ChatGroup_ReadGroupIDs(chatGroupID).Value(ctx)
			if err != nil {
				return fmt.Errorf("getting chat read group ids: %w", err)
			}

			writeGroups, err := ds.ChatGroup_WriteGroupIDs(chatGroupID).Value(ctx)
			if err != nil {
				return fmt.Errorf("getting chat read group ids: %w", err)
			}

			allGroups := append(readGroups, writeGroups...)

			attrMap.Add(dskey.Key{Collection: c.Name(), ID: chatGroupID, Field: "A"}, &Attributes{
				GlobalPermission: byte(perm.OMLSuperadmin),
				GroupIDs:         set.New(append(manageGroups, allGroups...)...),
			})
		}

		return nil
	})
}
