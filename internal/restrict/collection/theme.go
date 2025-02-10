package collection

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// Theme handels the restrictions for the theme collection.
//
// Every user can see a theme.
type Theme struct{}

// Name returns the collection name.
func (t Theme) Name() string {
	return "theme"
}

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
