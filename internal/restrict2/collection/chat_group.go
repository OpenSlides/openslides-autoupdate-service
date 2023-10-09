package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
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

func (c ChatGroup) see(ctx context.Context, fetcher *dsfetch.Fetch, chatGroupIDs []int) ([]attribute.Func, error) {
	readGroupIDs := make([][]int, len(chatGroupIDs))
	writeGroupIDs := make([][]int, len(chatGroupIDs))
	meetingID := make([]int, len(chatGroupIDs))
	for i, id := range chatGroupIDs {
		if id == 0 {
			continue
		}
		fetcher.ChatGroup_MeetingID(id).Lazy(&meetingID[i])
		fetcher.ChatGroup_ReadGroupIDs(id).Lazy(&readGroupIDs[i])
		fetcher.ChatGroup_WriteGroupIDs(id).Lazy(&writeGroupIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching chat group data: %w", err)
	}

	attr := make([]attribute.Func, len(chatGroupIDs))
	for i, id := range chatGroupIDs {
		if id == 0 {
			continue
		}
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		attr[i] = attribute.FuncOr(
			attribute.FuncInGroup(groupMap[perm.ChatCanManage]),
			attribute.FuncInGroup(readGroupIDs[i]),
			attribute.FuncInGroup(writeGroupIDs[i]),
		)

	}

	return attr, nil
}
