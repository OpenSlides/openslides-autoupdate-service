package restrict

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
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

	start := time.Now()
	times, err := restrict(ctx, r.getter, r.uid, data)
	if err != nil {
		return nil, fmt.Errorf("restricting data: %w", err)
	}

	duration := time.Since(start)

	if times != nil && (duration > slowCalls || oserror.HasTagFromContext(ctx, "profile_restrict")) {
		body, ok := oserror.BodyFromContext(ctx)
		if !ok {
			body = "unknown body, probably simple request"
		}
		profile(body, duration, times)
	}

	return data, nil
}

// restrict changes the keys and values in data for the user with the given user
// id.
func restrict(ctx context.Context, getter datastore.Getter, uid int, data map[datastore.Key][]byte) (map[string]timeCount, error) {
	ds := dsfetch.New(getter)

	isSuperAdmin, err := perm.HasOrganizationManagementLevel(ctx, ds, uid, perm.OMLSuperadmin)
	if err != nil {
		var errDoesNotExist dsfetch.DoesNotExistError
		if errors.As(err, &errDoesNotExist) || datastore.Key(errDoesNotExist).Collection == "user" {
			// TODO LAST ERROR
			return nil, fmt.Errorf("request user %d does not exist", uid)
		}
		return nil, fmt.Errorf("checking for superadmin: %w", err)
	}

	if isSuperAdmin {
		if err := restrictSuperAdmin(ctx, getter, uid, data); err != nil {
			return nil, fmt.Errorf("restrict as superadmin: %w", err)
		}
		return nil, nil
	}

	restrictModeIDs := make(map[collectionMode]*set.Set[int])
	if err := groupKeysByCollection(data, restrictModeIDs); err != nil {
		return nil, fmt.Errorf("grouping keys by collection: %w", err)
	}

	// Get all required collectionModes from relation fields.
	relationKeys := make(map[datastore.Key]struct{})
	relationListKeys := make(map[datastore.Key][]int)
	genericRelationKeys := make(map[datastore.Key]struct{})
	genericRelationListKeys := make(map[datastore.Key][]string)
	for key := range data {
		if data[key] == nil {
			continue
		}

		keyPrefix := templateKeyPrefix(key.CollectionField())

		cm, id, ok, err := isRelation(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for relation: %w", key, err)
		}

		if ok {
			if restrictModeIDs[cm] == nil {
				restrictModeIDs[cm] = set.New[int]()
			}
			restrictModeIDs[cm].Add(id)
			relationKeys[key] = struct{}{}
			continue
		}

		cm, ids, ok, err := isRelationList(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for relation-list: %w", key, err)
		}

		if ok {
			if restrictModeIDs[cm] == nil {
				restrictModeIDs[cm] = set.New[int]()
			}
			restrictModeIDs[cm].Add(ids...)
			relationListKeys[key] = ids
			continue
		}

		cm, id, ok, err = isGenericRelation(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for generic-relation: %w", key, err)
		}

		if ok {
			if restrictModeIDs[cm] == nil {
				restrictModeIDs[cm] = set.New[int]()
			}
			restrictModeIDs[cm].Add(id)
			genericRelationKeys[key] = struct{}{}
			continue
		}

		mcm, value, ok, err := isGenericRelationList(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for generic-relation-list: %w", key, err)
		}

		if ok {
			for _, cmID := range mcm {
				if restrictModeIDs[cm] == nil {
					restrictModeIDs[cm] = set.New[int]()
				}
				restrictModeIDs[cmID.cm].Add(cmID.id)
			}
			genericRelationListKeys[key] = value

			continue
		}
	}

	if len(restrictModeIDs) == 0 {
		return nil, nil
	}

	times := make(map[string]timeCount, len(restrictModeIDs))

	// Call restrict Mode function
	mperms := perm.NewMeetingPermission(ds, uid)
	for cm, ids := range restrictModeIDs {
		idsCount := ids.Len()
		start := time.Now()
		modeFunc, err := restrictModefunc(cm.collection, cm.mode)
		if err != nil {
			return nil, fmt.Errorf("getting restiction mode for %s/%s: %w", cm.collection, cm.mode, err)
		}

		allowedIDs, err := modeFunc(ctx, ds, mperms, ids.List()...)
		if err != nil {
			var errDoesNotExist dsfetch.DoesNotExistError
			if !errors.As(err, &errDoesNotExist) {
				return nil, fmt.Errorf("calling collection %s modefunc %s with ids %v: %w", cm.collection, cm.mode, ids.List(), err)
			}
		}
		restrictModeIDs[cm].Remove(allowedIDs...)

		duration := time.Since(start)
		times[cm.collection+"/"+cm.mode] = timeCount{time: duration, count: idsCount}
	}

	// Remove restricted keys
	for key := range data {
		if data[key] == nil {
			continue
		}

		restrictionMode, err := buildFieldMode(key.Collection, key.Field)
		if err != nil {
			return nil, fmt.Errorf("getting restriction Mode for %s: %w", key, err)
		}

		cm := collectionMode{key.Collection, restrictionMode}
		if restrictModeIDs[cm].Has(key.ID) {
			data[key] = nil
			continue
		}
	}

	// Remove entries from relation fields
	for key := range relationKeys {
		if data[key] == nil {
			// The field was restricted.
			continue
		}

		keyPrefix := templateKeyPrefix(key.CollectionField())
		cm, id, _, err := isRelation(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for relation: %w", key, err)
		}

		if restrictModeIDs[cm].Has(id) {
			data[key] = nil
		}
	}

	// Remove entries from relation-list fields
	for key := range relationListKeys {
		if data[key] == nil {
			// The field was restricted.
			continue
		}

		keyPrefix := templateKeyPrefix(key.CollectionField())
		cm, ids, _, err := isRelationList(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for relation-list: %w", key, err)
		}

		allowed := make([]int, 0, len(ids))
		for _, id := range ids {
			if !restrictModeIDs[cm].Has(id) {
				allowed = append(allowed, id)
			}
		}

		if len(allowed) != len(ids) {
			newValue, err := json.Marshal(allowed)
			if err != nil {
				return nil, fmt.Errorf("marshal new value for key %s: %w", key, err)
			}
			data[key] = newValue
		}
	}

	// Remove entries from generic-relation fields
	for key := range genericRelationKeys {
		if data[key] == nil {
			// The field was restricted.
			continue
		}

		keyPrefix := templateKeyPrefix(key.CollectionField())
		cm, id, _, err := isGenericRelation(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for generic-relation: %w", key, err)
		}

		if restrictModeIDs[cm].Has(id) {
			data[key] = nil
		}
	}

	// Remove entries from generic-relation-list fields
	for key := range genericRelationListKeys {
		if data[key] == nil {
			// The field was restricted.
			continue
		}

		keyPrefix := templateKeyPrefix(key.CollectionField())
		mcm, genericIDs, _, err := isGenericRelationList(keyPrefix, data[key])
		if err != nil {
			return nil, fmt.Errorf("checking %s for generic-relation-list: %w", key, err)
		}

		allowed := make([]string, 0, len(genericIDs))
		for genericID, cmID := range mcm {
			if !restrictModeIDs[cmID.cm].Has(cmID.id) {
				allowed = append(allowed, genericID)
			}

		}

		if len(allowed) != len(genericIDs) {
			newValue, err := json.Marshal(allowed)
			if err != nil {
				return nil, fmt.Errorf("marshal new value for key %s: %w", key, err)
			}
			data[key] = newValue
		}
	}

	return times, nil
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

type collectionMode struct {
	collection string
	mode       string
}

// groupKeysByCollection groups all the keys in data by there collection.
func groupKeysByCollection(data map[datastore.Key][]byte, restrictModeIDs map[collectionMode]*set.Set[int]) error {
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
			restrictModeIDs[cm] = set.New[int]()
		}
		restrictModeIDs[cm].Add(key.ID)
	}

	return nil
}

func isRelation(keyPrefix string, value []byte) (collectionMode, int, bool, error) {
	toCollectionfield, ok := relationFields[keyPrefix]
	if !ok {
		return collectionMode{}, 0, false, nil
	}

	var id int
	if err := json.Unmarshal(value, &id); err != nil {
		return collectionMode{}, 0, false, fmt.Errorf("decoding %q (`%s`): %w", keyPrefix, value, err)
	}

	collection, field, _ := strings.Cut(toCollectionfield, "/")
	fieldMode, err := buildFieldMode(collection, field)
	if err != nil {
		return collectionMode{}, 0, false, fmt.Errorf("building relation field field mode: %w", err)
	}

	return collectionMode{collection: collection, mode: fieldMode}, id, true, nil
}

func isRelationList(keyPrefix string, value []byte) (collectionMode, []int, bool, error) {
	toCollectionfield, ok := relationListFields[keyPrefix]
	if !ok {
		return collectionMode{}, nil, false, nil
	}

	var ids []int
	if err := json.Unmarshal(value, &ids); err != nil {
		return collectionMode{}, nil, false, fmt.Errorf("decoding value (size: %d): %w", len(value), err)
	}

	collection, field, _ := strings.Cut(toCollectionfield, "/")
	fieldMode, err := buildFieldMode(collection, field)
	if err != nil {
		return collectionMode{}, nil, false, fmt.Errorf("building relation field field mode: %w", err)
	}

	return collectionMode{collection: collection, mode: fieldMode}, ids, true, nil
}

func isGenericRelation(keyPrefix string, value []byte) (collectionMode, int, bool, error) {
	toCollectionfield, ok := genericRelationFields[keyPrefix]
	if !ok {
		return collectionMode{}, 0, false, nil
	}

	var genericID string
	if err := json.Unmarshal(value, &genericID); err != nil {
		return collectionMode{}, 0, false, fmt.Errorf("decoding %q: %w", keyPrefix, err)
	}

	cm, id, err := genericKeyToCollectionMode(genericID, toCollectionfield)
	if err != nil {
		return collectionMode{}, 0, false, fmt.Errorf("parsing generic key: %w", err)
	}

	return cm, id, true, nil
}

type collectionModeID struct {
	cm collectionMode
	id int
}

func isGenericRelationList(keyPrefix string, value []byte) (map[string]collectionModeID, []string, bool, error) {
	toCollectionfield, ok := genericRelationListFields[keyPrefix]
	if !ok {
		return nil, nil, false, nil
	}

	var genericIDs []string
	if err := json.Unmarshal(value, &genericIDs); err != nil {
		return nil, nil, false, fmt.Errorf("decoding %q: %w", keyPrefix, err)
	}

	mcm := make(map[string]collectionModeID, len(genericIDs))
	for _, genericID := range genericIDs {
		cm, id, err := genericKeyToCollectionMode(genericID, toCollectionfield)
		if err != nil {
			return nil, nil, false, fmt.Errorf("parsing generic key: %w", err)
		}
		mcm[genericID] = collectionModeID{cm, id}
	}

	return mcm, genericIDs, true, nil
}

// genericKeyToCollectionMode calls f for each collection mode.
func genericKeyToCollectionMode(genericID string, toCollectionFieldMap map[string]string) (collectionMode, int, error) {
	collection, rawID, found := strings.Cut(genericID, "/")
	if !found {
		// TODO LAST ERROR
		return collectionMode{}, 0, fmt.Errorf("invalid generic relation: %s", genericID)
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		// TODO LAST ERROR
		return collectionMode{}, 0, fmt.Errorf("invalid generic relation, no id: %s", genericID)
	}

	toField := toCollectionFieldMap[collection]
	if toField == "" {
		// TODO LAST ERROR
		return collectionMode{}, 0, fmt.Errorf("unknown generic relation: %s", collection)
	}

	fieldMode, err := buildFieldMode(collection, toField)
	if err != nil {
		return collectionMode{}, 0, fmt.Errorf("building generic relation field mode: %w", err)
	}

	return collectionMode{collection: collection, mode: fieldMode}, id, nil
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
