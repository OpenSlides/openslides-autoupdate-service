package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// PersonalNote handels permissions for personal notes.
type PersonalNote struct {
	dp dataprovider.DataProvider
}

// NewPersonalNote initializes a personal note.
func NewPersonalNote(dp dataprovider.DataProvider) *PersonalNote {
	return &PersonalNote{
		dp: dp,
	}
}

// Connect creates the routes.
func (p *PersonalNote) Connect(s collection.HandlerStore) {
	s.RegisterWriteHandler("personal_note.create", collection.WriteCheckerFunc(p.create))
	s.RegisterWriteHandler("personal_note.update", collection.WriteCheckerFunc(p.modify))
	s.RegisterWriteHandler("personal_note.delete", collection.WriteCheckerFunc(p.modify))

	s.RegisterReadHandler("personal_note", p)
}

func (p PersonalNote) modify(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	fqfield := fmt.Sprintf("personal_note/%s/user_id", payload["id"])
	var noteUserID int
	if err := p.dp.Get(ctx, fqfield, &noteUserID); err != nil {
		return nil, fmt.Errorf("getting %s from datastore: %w", fqfield, err)
	}

	if noteUserID != userID {
		return nil, collection.NotAllowedf("Not your note")
	}
	return nil, nil
}

func (p PersonalNote) create(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	if userID == 0 {
		collection.NotAllowedf("Anonymous can not create personal notes.")
	}
	return nil, nil
}

// RestrictFQFields checks for read permissions.
func (p PersonalNote) RestrictFQFields(ctx context.Context, userID int, fqfields []collection.FQField, result map[string]bool) error {
	var noteUserID int
	var lastID int
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
			f := fmt.Sprintf("personal_note/%d/user_id", fqfield.ID)
			if err := p.dp.Get(ctx, f, &noteUserID); err != nil {
				return fmt.Errorf("getting %s from datastore: %w", f, err)
			}
			lastID = fqfield.ID
		}

		if noteUserID != userID {
			continue
		}

		result[fqfield.String()] = true
	}
	return nil
}
