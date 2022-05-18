package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Theme handels the restrictions for the theme collection.
//
// Every user can see a theme.
type Theme struct{}

// MeetingID returns the meetingID for the object.
func (t Theme) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the field restriction for each mode.
func (t Theme) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	}
	return nil
}
