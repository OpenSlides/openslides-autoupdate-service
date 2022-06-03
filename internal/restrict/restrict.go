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
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Middleware can be used as a datastore.Getter that restrict the data for a
// user.
func Middleware(getter datastore.Getter, uid int) datastore.Getter {
	return restricter{
		getter: getter,
		uid:    uid,
	}
}

type restricter struct {
	getter datastore.Getter
	uid    int
}

// Get returns restricted data.
func (r restricter) Get(ctx context.Context, keys ...datastore.Key) (map[datastore.Key][]byte, error) {
	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("getting data: %w", err)
	}

	if err := restrict(ctx, r.getter, r.uid, data); err != nil {
		return nil, fmt.Errorf("restricting data: %w", err)
	}
	return data, nil
}

// restrict changes the keys and values in data for the user with the given user
// id.
func restrict(ctx context.Context, getter datastore.Getter, uid int, data map[datastore.Key][]byte) error {
	ds := dsfetch.New(getter)
	isSuperAdmin, err := perm.HasOrganizationManagementLevel(ctx, ds, uid, perm.OMLSuperadmin)
	if err != nil {
		var errDoesNotExist dsfetch.DoesNotExistError
		if errors.As(err, &errDoesNotExist) || datastore.Key(errDoesNotExist).Collection == "user" {
			// TODO LAST ERROR
			return fmt.Errorf("request user %d does not exist", uid)
		}
		return fmt.Errorf("checking for superadmin: %w", err)
	}

	if isSuperAdmin {
		return restrictSuperAdmin(ctx, getter, uid, data)
	}

	mperms := perm.NewMeetingPermission(ds, uid)

	type collectionMode struct {
		collection string
		mode       string
	}
	restrictModeIDs := make(map[collectionMode]*set.Set)

	// Group all ids with same collection-restrictionMode
	for key := range data {
		if data[key] == nil {
			continue
		}

		restrictionMode, err := buildFieldMode(key.Collection, key.Field)
		if err != nil {
			return fmt.Errorf("getting restriction Mode for %s: %w", key, err)
		}

		cm := collectionMode{key.Collection, restrictionMode}
		if restrictModeIDs[cm] == nil {
			restrictModeIDs[cm] = set.New()
		}
		restrictModeIDs[cm].Add(key.ID)
	}

	// Call restrict Mode function
	for cm, ids := range restrictModeIDs {
		modeFunc, err := restrictModefunc(cm.collection, cm.mode)
		if err != nil {
			return fmt.Errorf("getting restiction mode for %s/%s: %w", cm.collection, cm.mode, err)
		}

		allowedIDs, err := modeFunc(ctx, ds, mperms, ids.List()...)
		if err != nil {
			return fmt.Errorf("calling collection %s modefunc %s with ids %v: %w", cm.collection, cm.mode, ids.List(), err)
		}
		restrictModeIDs[cm].Remove(allowedIDs...)
	}

	// Remove restricted keys
	for key := range data {
		if data[key] == nil {
			continue
		}

		restrictionMode, err := buildFieldMode(key.Collection, key.Field)
		if err != nil {
			return fmt.Errorf("getting restriction Mode for %s: %w", key, err)
		}

		cm := collectionMode{key.Collection, restrictionMode}
		if restrictModeIDs[cm].Has(key.ID) {
			data[key] = nil
			continue
		}

		// Relation fields
		if toCollectionfield, ok := relationFields[templateKeyPrefix(key.CollectionField())]; ok {
			var id int
			if err := json.Unmarshal(data[key], &id); err != nil {
				return fmt.Errorf("decoding %q: %w", key, err)
			}

			parts := strings.Split(toCollectionfield, "/")
			fieldMode, err := buildFieldMode(parts[0], parts[1])
			if err != nil {
				return fmt.Errorf("building relation field field mode: %w", err)
			}

			modeFunc, err := restrictModefunc(parts[0], fieldMode)
			if err != nil {
				return fmt.Errorf("getting restict func: %w", err)
			}

			allowedID, err := modeFunc(ctx, ds, mperms, id)
			if err != nil {
				return fmt.Errorf("checking can see of relation field %s: %w", key, err)
			}
			if len(allowedID) != 1 {
				data[key] = nil
			}
		}

		keyPrefix := templateKeyPrefix(key.CollectionField())
		// Relation List fields
		if toCollectionfield, ok := relationListFields[keyPrefix]; ok {
			value, err := filterRelationList(ctx, ds, mperms, key, toCollectionfield, data[key])
			if err != nil {
				return fmt.Errorf("restrict relation-list ids of %q: %w", key, err)
			}
			data[key] = value
		}

		// Generic Relation fields
		if toCollectionFieldMap, ok := genericRelationFields[templateKeyPrefix(key.CollectionField())]; ok {
			var genericID string
			if err := json.Unmarshal(data[key], &genericID); err != nil {
				return fmt.Errorf("decoding %q: %w", key, err)
			}

			parts := strings.Split(genericID, "/")
			toField := toCollectionFieldMap[parts[0]]
			if toField == "" {
				// TODO LAST ERROR
				return fmt.Errorf("invalid generic relation for field %q: %s", key.CollectionField(), parts[0])
			}

			fieldMode, err := buildFieldMode(parts[0], toField)
			if err != nil {
				return fmt.Errorf("building generic relation field mode: %w", err)
			}

			modeFunc, err := restrictModefunc(parts[0], fieldMode)
			if err != nil {
				return fmt.Errorf("getting restict func: %w", err)
			}

			id, err := strconv.Atoi(parts[1])
			if err != nil {
				return fmt.Errorf("decoding genericID: %w", err)
			}

			allowedID, err := modeFunc(ctx, ds, mperms, id)
			if err != nil {
				return fmt.Errorf("checking can see: %w", err)
			}
			if len(allowedID) != 1 {
				data[key] = nil
			}
		}

		// Generic Relation List fields
		if toCollectionfieldMap, ok := genericRelationListFields[templateKeyPrefix(key.CollectionField())]; ok {
			value, err := filterGenericRelationList(ctx, ds, mperms, key, toCollectionfieldMap, data[key])
			if err != nil {
				return fmt.Errorf("restrict generic-relation-list ids of %q: %w", key, err)
			}
			data[key] = value
		}
	}

	return nil
}

func restrictSuperAdmin(ctx context.Context, getter datastore.Getter, uid int, data map[datastore.Key][]byte) error {
	ds := dsfetch.New(getter)
	mperms := perm.NewMeetingPermission(ds, uid)

	for key := range data {
		if data[key] == nil {
			continue
		}

		restricter := collection.Collection(key.Collection)
		if restricter == nil {
			// Superadmins can see unknown collections.
			continue
		}

		type superRestricter interface {
			SuperAdmin(mode string) collection.FieldRestricter
		}
		sr, ok := restricter.(superRestricter)
		if !ok {
			continue
		}

		restrictionMode, err := buildFieldMode(key.Collection, key.Field)
		if err != nil {
			return fmt.Errorf("getting restriction Mode for %s: %w", key, err)
		}

		modefunc := sr.SuperAdmin(restrictionMode)
		if modefunc == nil {
			// Do not restrict unknown fields that are not implemented.
			continue
		}

		allowed, err := modefunc(ctx, ds, mperms, key.ID)
		if err != nil {
			return fmt.Errorf("calling mode func: %w", err)
		}

		if len(allowed) == 0 {
			data[key] = nil
		}
	}
	return nil
}

// buildFieldMode returns the restriction mode for a collection and field.
//
// This is a string like "A" or "B" or any other name of a restriction mode.
func buildFieldMode(collection, field string) (string, error) {
	fieldMode, ok := restrictionModes[templateKeyPrefix(collection+"/"+field)]
	if !ok {
		// TODO LAST ERROR
		return "", fmt.Errorf("fqfield %q is unknown, maybe run go generate ./... to fetch all fields from the models.yml", collection+"/"+field)
	}
	return fieldMode, nil
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

func filterRelationList(
	ctx context.Context,
	ds *dsfetch.Fetch,
	mperms *perm.MeetingPermission,
	key datastore.Key,
	toCollectionField string,
	data []byte,
) ([]byte, error) {
	var ids []int
	if err := json.Unmarshal(data, &ids); err != nil {
		return nil, fmt.Errorf("decoding ids: %w", err)
	}

	parts := strings.Split(toCollectionField, "/")

	fieldMode, err := buildFieldMode(parts[0], parts[1])
	if err != nil {
		return nil, fmt.Errorf("building field mode: %w", err)
	}

	relationListModeFunc, err := restrictModefunc(parts[0], fieldMode)
	if err != nil {
		// Collection or field unknown. Handle it as no permission.
		log.Printf("Warning: getting restriction mode for values of relation list field %s: %v", key, err)
		return nil, nil
	}

	allowedIDs, err := relationListModeFunc(ctx, ds, mperms, ids...)
	if err != nil {
		// TODO: add currupted database warning if doesNotExist error from toCollectionField with one of the ids.
		return nil, fmt.Errorf("calling relation ist mode func: %w", err)
	}

	if allowedIDs == nil {
		// this is important for json
		allowedIDs = make([]int, 0)
	}

	encoded, err := json.Marshal(allowedIDs)
	if err != nil {
		return nil, fmt.Errorf("encoding ids: %w", err)
	}
	return encoded, nil
}

func filterGenericRelationList(
	ctx context.Context,
	ds *dsfetch.Fetch,
	mperms *perm.MeetingPermission,
	key datastore.Key,
	toCollectionFieldMap map[string]string,
	data []byte,
) ([]byte, error) {
	var genericIDs []string
	if err := json.Unmarshal(data, &genericIDs); err != nil {
		return nil, fmt.Errorf("decoding ids: %w", err)
	}

	allowedIDs := []string{} // Use empty list as default for json encoding.
	for _, genericID := range genericIDs {
		parts := strings.Split(genericID, "/")
		id, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid generic id: %w", err)
		}

		toField := toCollectionFieldMap[parts[0]]
		if toField == "" {
			// TODO LAST ERROR
			return nil, fmt.Errorf("invalid generic relation: %s", parts[0])
		}

		fieldMode, err := buildFieldMode(parts[0], toField)
		if err != nil {
			return nil, fmt.Errorf("building field mode: %w", err)
		}

		// TODO: call all restictMode of one collection at once
		relationListModeFunc, err := restrictModefunc(parts[0], fieldMode)
		if err != nil {
			// Collection or field unknown. Handle it as no permission.
			fmt.Printf("Warning: getting restiction mode values of generic key %s: %v", key, err)
			return nil, nil
		}

		allowedID, err := relationListModeFunc(ctx, ds, mperms, id)
		if err != nil {
			return nil, fmt.Errorf("checking %q for id %d: %w", toField, id, err)
		}
		if len(allowedID) == 1 {
			allowedIDs = append(allowedIDs, genericID)
		}
	}

	if allowedIDs == nil {
		// this is important for json
		allowedIDs = make([]string, 0)
	}

	encoded, err := json.Marshal(allowedIDs)
	if err != nil {
		return nil, fmt.Errorf("encoding ids: %w", err)
	}
	return encoded, nil
}

// restrictModefunc returns the field restricter function to use.
func restrictModefunc(collectionName, fieldMode string) (collection.FieldRestricter, error) {
	restricter := collection.Collection(collectionName)
	if restricter == nil {
		// TODO LAST ERROR
		return nil, fmt.Errorf("collection %q is not implemented, maybe run go generate ./... to fetch all fields from the models.yml", collectionName)
	}

	modefunc := restricter.Modes(fieldMode)
	if modefunc == nil {
		// TODO LAST ERROR
		return nil, fmt.Errorf("mode %q of models %q is not implemented", fieldMode, collectionName)
	}

	return modefunc, nil
}
