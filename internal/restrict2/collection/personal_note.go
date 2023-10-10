package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// PersonalNote handels restriction for the personal_node collection.
//
// The user can see a personal node, if personal_note/user_id is the same as the id of the requested user.
//
// Mode A: The user can see the personal note.
type PersonalNote struct{}

// Name returns the collection name.
func (p PersonalNote) Name() string {
	return "personal_note"
}

// MeetingID returns the meetingID for the object.
func (p PersonalNote) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.PersonalNote_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the field restriction for each mode.
func (p PersonalNote) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p PersonalNote) see(ctx context.Context, fetcher *dsfetch.Fetch, personalNoteIDs []int) ([]attribute.Func, error) {
	meetingUserID := make([]int, len(personalNoteIDs))
	for i, id := range personalNoteIDs {
		if id == 0 {
			continue
		}
		fetcher.PersonalNote_MeetingUserID(id).Lazy(&meetingUserID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching meeting user: %w", err)
	}

	userID := make([]int, len(personalNoteIDs))
	for i, id := range meetingUserID {
		if id == 0 {
			continue
		}
		fetcher.MeetingUser_UserID(id).Lazy(&userID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user: %w", err)
	}

	attr := make([]attribute.Func, len(personalNoteIDs))
	for i, id := range userID {
		if id == 0 {
			continue
		}
		attr[i] = attribute.FuncUserIDs([]int{id})
	}
	return attr, nil
}
