package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ChatMessage handels restrictions for the collection chat_message.
//
// A user can see a chat_message if they have the permission chat.can_manage
// or they are in the respective chat_group/read_group_ids (given by the key
// chat_group_id) or they have written the chat_message (dedicated by the key user_id).
//
// Mode A: A user can see a chat_message.
type ChatMessage struct{}

// Name returns the collection name.
func (c ChatMessage) Name() string {
	return "chat_message"
}

// MeetingID returns the meetingID for the object.
func (c ChatMessage) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	chatGroupID, err := ds.ChatMessage_ChatGroupID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting chat_group_id: %w", err)
	}

	return ChatGroup{}.MeetingID(ctx, ds, chatGroupID)
}

// Modes give sthe FieldRestricter for a restriction_mode.
func (c ChatMessage) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return c.see
	}
	return nil
}

func (ChatMessage) see(ctx context.Context, fetcher *dsfetch.Fetch, chatMessageIDs []int) ([]attribute.Func, error) {
	chatGroupID := make([]int, len(chatMessageIDs))
	meetingUserID := make([]int, len(chatMessageIDs))
	for i, id := range chatMessageIDs {
		if id == 0 {
			continue
		}
		fetcher.ChatMessage_MeetingUserID(id).Lazy(&meetingUserID[i])
		fetcher.ChatMessage_ChatGroupID(id).Lazy(&chatGroupID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching chat message data: %w", err)
	}

	userID := make([]int, len(chatMessageIDs))
	for i, id := range meetingUserID {
		if id == 0 {
			continue
		}
		fetcher.MeetingUser_UserID(id).Lazy(&userID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user from meeting user: %w", err)
	}

	canSeeCharGroup, err := Collection(ctx, ChatGroup{}).Modes("A")(ctx, fetcher, chatGroupID)
	if err != nil {
		return nil, fmt.Errorf("check chat group: %w", err)
	}

	attr := make([]attribute.Func, len(chatMessageIDs))
	for i, id := range chatMessageIDs {
		if id == 0 {
			continue
		}
		attr[i] = attribute.FuncOr(
			canSeeCharGroup[i],
			attribute.FuncUserIDs([]int{userID[i]}),
		)
	}
	return attr, nil
}
