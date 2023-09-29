package collection

import (
	"context"
	"errors"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ImportPreview handels restrictions of the collection import_preview.
//
// The user can see a import_preview, if TODO
//
// Mode A: The user can see the import_preview.
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
		return i.see
	}
	return nil
}

func (i ImportPreview) see(ctx context.Context, ds *dsfetch.Fetch, importPreviewIDs ...int) ([]int, error) {
	return nil, errors.New("TODO")
}
