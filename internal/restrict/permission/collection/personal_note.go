package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// PersonalNote initializes a personal note.
func PersonalNote(dp dataprovider.DataProvider) perm.ConnecterFunc {
	p := &personalNote{dp}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("personal_note", p)
	}
}

type personalNote struct {
	dp dataprovider.DataProvider
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
