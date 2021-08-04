package restrict

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Restrict changes the keys and values in data for the user with the given user
// id.
func Restrict(ctx context.Context, fetch *datastore.Fetcher, uid int, data map[string][]byte) error {
	mperms := perm.NewMeetingPermission(fetch, uid)
	for key := range data {
		fqfield, err := parseFQField(key)
		if err != nil {
			return fmt.Errorf("parsing fqfield %s: %w", key, err)
		}

		restricter, ok := collections[fqfield.Collection]
		if !ok {
			data[key] = nil
			continue
		}

		canSee, err := restricter.See(ctx, fetch, mperms, fqfield.ID)
		if err != nil {
			var errDoesNotExist datastore.DoesNotExistError
			if !errors.As(err, &errDoesNotExist) {
				return fmt.Errorf("see permission for %s: %w", fqfield, err)
			}

			// If an element does not exist, then just handel it as no
			// permission.
			data[key] = nil
			continue
		}

		if !canSee {
			data[key] = nil
			continue
		}

		fieldMode, ok := restrictionModes[fqfield.CollectionField()]
		if !ok {
			// Field unknown
			data[key] = nil
			continue
		}

		canSeeMode, err := restricter.Modes(fieldMode)(ctx, fetch, mperms, fqfield.ID)
		if err != nil {
			var errDoesNotExist datastore.DoesNotExistError
			if !errors.As(err, &errDoesNotExist) {
				return fmt.Errorf("Mode %s permission for %s: %w", fieldMode, fqfield, err)
			}

			// If an element does not exist, then just handel it as no
			// permission.
			data[key] = nil
			continue
		}

		if !canSeeMode {
			data[key] = nil
			continue
		}

		// TODO: Lock for relation, relation-list, generic-* fields, also in templates, and remove items without permission.
	}

	return nil
}

type collectionRestricter interface {
	See(ctx context.Context, fetch *datastore.Fetcher, mperms perm.MeetingPermission, agendaID int) (bool, error)
	Modes(mode string) collection.FieldRestricter
}

var collections = map[string]collectionRestricter{
	"agenda_item": collection.AgendaItem{},
	"assignment":  collection.Assignment{},
}
