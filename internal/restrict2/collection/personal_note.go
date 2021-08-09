package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// PersonalNote handels restriction for the personal_node collection
type PersonalNote struct{}

// Modes returns the field restriction for each mode.
func (p PersonalNote) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p PersonalNote) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, personalNoteID int) (bool, error) {
	pUserID := fetch.Field().PersonalNote_UserID(ctx, personalNoteID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching user id of personal note: %w", err)
	}
	return mperms.UserID() == pUserID, nil
}

// SuperAdmin restricts the super admin.
func (p PersonalNote) SuperAdmin(mode string) FieldRestricter {
	return p.Modes(mode)
}
