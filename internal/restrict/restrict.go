package restrict

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Restrict changes the keys and values in data for the user with the given user
// id.
func Restrict(ctx context.Context, fetch *datastore.Fetcher, uid int, data map[string]json.RawMessage) error {
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
			// Collection or field unknown. Handle it as no permission.
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

		// Relation fields
		if toCollectionfield, ok := relationFields[templateKeyPrefix(fqfield.CollectionField())]; ok {
			var id int
			if err := json.Unmarshal(data[key], &id); err != nil {
				return fmt.Errorf("decoding %q: %w", key, err)
			}

			parts := strings.Split(toCollectionfield, "/")
			modeFunc, err := restrictMode(parts[0], parts[1], isSuperAdmin)
			if err != nil {
				return fmt.Errorf("getting restict func: %w", err)
			}

			cansee, err := modeFunc(ctx, fetch, mperms, id)
			if err != nil {
				return fmt.Errorf("checking can see: %w", err)
			}
			if !cansee {
				data[key] = nil
			}
		}

		foo := templateKeyPrefix(fqfield.CollectionField())
		// Relation List fields
		if toCollectionfield, ok := relationListFields[foo]; ok {
			value, err := filterRelationList(ctx, fetch, mperms, toCollectionfield, isSuperAdmin, data[key])
			if err != nil {
				return fmt.Errorf("restrict relation-list ids of %q: %w", key, err)
			}
			data[key] = value
		}

		// Generic Relation fields
		if toCollectionfield, ok := genericRelationFields[templateKeyPrefix(fqfield.CollectionField())]; ok {
			var genericID string
			if err := json.Unmarshal(data[key], &genericID); err != nil {
				return fmt.Errorf("decoding %q: %w", key, err)
			}

			parts := strings.Split(genericID, "/")
			modeFunc, err := restrictMode(toCollectionfield, parts[2], isSuperAdmin)
			if err != nil {
				return fmt.Errorf("getting restict func: %w", err)
			}

			id, err := strconv.Atoi(parts[0])
			if err != nil {
				return fmt.Errorf("decoding genericID: %w", err)
			}

			cansee, err := modeFunc(ctx, fetch, mperms, id)
			if err != nil {
				return fmt.Errorf("checking can see: %w", err)
			}
			if !cansee {
				data[key] = nil
			}
		}

		// Generic Relation List fields
		if toCollectionfield, ok := genericRelationListFields[templateKeyPrefix(fqfield.CollectionField())]; ok {
			value, err := filterGenericRelationList(ctx, fetch, mperms, toCollectionfield, isSuperAdmin, data[key])
			if err != nil {
				return fmt.Errorf("restrict generic-relation-list ids of %q: %w", key, err)
			}
			data[key] = value
		}
	}

	return nil
}

// templateKeyPrefix returns the index of the field list list. For template fields this is
// the key without the replacement.
func templateKeyPrefix(collectionField string) string {
	i := strings.IndexByte(collectionField, '$')
	if i < 0 || i == len(collectionField)-1 || collectionField[i+1] == '_' {
		// Normal field or $ at the end or $_
		return collectionField
	}

	return collectionField[:i+1]
}

func filterRelationList(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, toCollectionField string, isSuperAdmin bool, data []byte) ([]byte, error) {
	var ids []int
	if err := json.Unmarshal(data, &ids); err != nil {
		return nil, fmt.Errorf("decoding ids: %w", err)
	}

	parts := strings.Split(toCollectionField, "/")

	relationListModeFunc, err := restrictMode(parts[0], parts[1], isSuperAdmin)
	if err != nil {
		// Collection or field unknown. Handle it as no permission.
		log.Printf("Warning: %v", err)
		return nil, nil
	}

	var allowedIDs []int
	for _, id := range ids {
		allowed, err := relationListModeFunc(ctx, fetch, mperms, id)
		if err != nil {
			return nil, fmt.Errorf("checking %q for id %d: %w", toCollectionField, id, err)
		}
		if allowed {
			allowedIDs = append(allowedIDs, id)
		}
	}

	encoded, err := json.Marshal(allowedIDs)
	if err != nil {
		return nil, fmt.Errorf("encoding ids: %w", err)
	}
	return encoded, nil
}

func filterGenericRelationList(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, toCollectionField string, isSuperAdmin bool, data []byte) ([]byte, error) {
	var genericIDs []string
	if err := json.Unmarshal(data, &genericIDs); err != nil {
		return nil, fmt.Errorf("decoding ids: %w", err)
	}

	var allowedIDs []string
	for _, genericID := range genericIDs {
		parts := strings.Split(genericID, "/")
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid generic id: %w", err)
		}

		relationListModeFunc, err := restrictMode(parts[0], toCollectionField, isSuperAdmin)
		if err != nil {
			// Collection or field unknown. Handle it as no permission.
			fmt.Printf("Warning: %v", err)
			return nil, nil
		}

		allowed, err := relationListModeFunc(ctx, fetch, mperms, id)
		if err != nil {
			return nil, fmt.Errorf("checking %q for id %d: %w", toCollectionField, id, err)
		}
		if allowed {
			allowedIDs = append(allowedIDs, genericID)
		}
	}

	encoded, err := json.Marshal(allowedIDs)
	if err != nil {
		return nil, fmt.Errorf("encoding ids: %w", err)
	}
	return encoded, nil
}

// restrictMode returns the field restricter function to use.
func restrictMode(collectionName, fieldName string, isSuperAdmin bool) (collection.FieldRestricter, error) {
	restricter := collection.Collection(collectionName)
	if restricter == nil {
		if isSuperAdmin {
			// Superadmins can see unknown collections.
			return collection.Allways, nil
		}
		return nil, fmt.Errorf("collection %q is not implemented, maybe run go generate ./... to fetch all fields from the models.yml", collectionName)
	}

	fieldMode, ok := restrictionModes[templateKeyPrefix(collectionName+"/"+fieldName)]
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
