package restrict

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Restrict changes the keys and values in data for the user with the given user
// id.
func Restrict(ctx context.Context, fetch *datastore.Fetcher, uid int, data map[string][]byte) error {
	isSuperAdmin, err := perm.HasOrganizationManagementLevel(ctx, fetch, uid, perm.OMLSuperadmin)
	if err != nil {
		return fmt.Errorf("checking for superadmin: %w", err)
	}

	mperms := perm.NewMeetingPermission(fetch, uid)
	for key := range data {
		fqfield, err := parseFQField(key)
		if err != nil {
			return fmt.Errorf("parsing fqfield %s: %w", key, err)
		}

		modeFunc, err := restrictMode(fqfield.Collection, fqfield.Field, isSuperAdmin)
		if err != nil {
			log.Printf("Warning: %v", err)
			data[key] = nil
			continue
		}

		canSeeMode, err := modeFunc(ctx, fetch, mperms, fqfield.ID)
		if err != nil {
			var errDoesNotExist datastore.DoesNotExistError
			if !errors.As(err, &errDoesNotExist) {
				return fmt.Errorf("calling modefunc for key %s: %w", key, err)
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

		// TODO: Look for relation, relation-list, generic-* fields, also in templates, and remove items without permission.
	}

	return nil
}

// restrictMode returns the field restricter function to use.
func restrictMode(collectionName, fieldName string, isSuperAdmin bool) (collection.FieldRestricter, error) {
	restricter := collection.Collection(collectionName)
	if restricter == nil {
		if isSuperAdmin {
			// Superadmins can see unknown collections.
			return collection.Allways, nil
		}
		return nil, fmt.Errorf("collection %q is not implemented", collectionName)
	}

	fieldMode, ok := restrictionModes[collectionName+"/"+fieldName]
	if !ok {
		return nil, fmt.Errorf("fqfield %q is unknown, maybe run go generate ./... to fetch all fields from the models.yml", collectionName+"/"+fieldName)
	}

	if isSuperAdmin {
		type superRestricter interface {
			SuperAdmin(mode string) collection.FieldRestricter
		}
		sr, ok := restricter.(superRestricter)
		if !ok {
			// Do not restrict unknown fields from collections that do not
			// implement the superRestricter.
			return collection.Allways, nil
		}
		modefunc := sr.SuperAdmin(fieldMode)
		if modefunc == nil {
			return nil, fmt.Errorf("mode %q of models %q for superadmin is not implemented", fieldMode, collectionName)
		}

	}

	modefunc := restricter.Modes(fieldMode)
	if modefunc == nil {
		return nil, fmt.Errorf("mode %q of models %q is not implemented", fieldMode, collectionName)
	}

	return modefunc, nil
}
