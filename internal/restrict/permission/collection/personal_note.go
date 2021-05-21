package collection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// PersonalNote initializes a personal note.
func PersonalNote(dp dataprovider.DataProvider) perm.ConnecterFunc {
	p := &personalNote{dp}
	return func(s perm.HandlerStore) {
		s.RegisterAction("personal_note.create", perm.ActionFunc(p.create))
		s.RegisterAction("personal_note.update", perm.ActionFunc(p.modify))
		s.RegisterAction("personal_note.delete", perm.ActionFunc(p.modify))

		s.RegisterRestricter("personal_note", p)
	}
}

type personalNote struct {
	dp dataprovider.DataProvider
}

func (p personalNote) create(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	if userID == 0 {
		perm.LogNotAllowedf("Anonymous can not create personal notes.")
		return false, nil
	}
	return true, nil
}

func (p personalNote) modify(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	fqfield := fmt.Sprintf("personal_note/%s/user_id", payload["id"])
	var noteUserID int
	if err := p.dp.Get(ctx, fqfield, &noteUserID); err != nil {
		return false, fmt.Errorf("getting %s from datastore: %w", fqfield, err)
	}

	if noteUserID != userID {
		perm.LogNotAllowedf("Note belongs to a different user.")
		return false, nil
	}
	return true, nil
}

// RestrictFQFields checks for read permissions.
func (p personalNote) RestrictFQFields(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	var noteUserID int
	var lastID int
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
			lastID = fqfield.ID
			key := fmt.Sprintf("personal_note/%d/user_id", fqfield.ID)
			if err := p.dp.Get(ctx, key, &noteUserID); err != nil {
				return fmt.Errorf("getting %s from datastore: %w", key, err)
			}
		}

		if noteUserID != userID {
			continue
		}

		result[fqfield.String()] = true
	}
	return nil
}
