package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// PersonalNote handels restriction for the personal_node collection.
//
// The user can see a personal node, if personal_note/user_id is the same as the id of the requested user.
//
// The superadmin can not see personal_notes from other users.
//
// Mode A: The user can see the personal note.
type PersonalNote struct{}

// Modes returns the field restriction for each mode.
func (p PersonalNote) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p PersonalNote) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, personalNoteID int) (bool, error) {
	pUserID, err := ds.PersonalNote_UserID(personalNoteID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching user id of personal note: %w", err)
	}
	return mperms.UserID() == pUserID, nil
}

// SuperAdmin restricts the super admin.
func (p PersonalNote) SuperAdmin(mode string) FieldRestricter {
	return p.Modes(mode)
}
