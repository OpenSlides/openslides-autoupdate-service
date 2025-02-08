package collection

import (
	"context"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// ImportPreview handels restrictions of the collection import_preview.
//
// Noone can see the import_preview
type ImportPreview struct{}

// Name returns the collection name.
func (i ImportPreview) Name() string {
	return "import_preview"
}

// MeetingID returns the meetingID for the object.
func (i ImportPreview) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the field restricters for the collection.
func (i ImportPreview) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return never
	}
	return nil
}
