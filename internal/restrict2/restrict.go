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
	"sync"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsrecorder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/fastjson"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Restricter TODO
type Restricter struct {
	flow flow.Flow

	mu         sync.RWMutex
	attributes map[dskey.CollectionMode]attribute.Func
	hotKeys    set.Set[dskey.Key]
}

// New initializes a restricter
func New(flow flow.Flow) *Restricter {
	return &Restricter{
		flow:       flow,
		attributes: make(map[dskey.CollectionMode]attribute.Func),
		hotKeys:    set.New[dskey.Key](),
	}
}

// Get returns the full (unrestricted) data.
func (r *Restricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return r.flow.Get(ctx, keys...)
}

// Update updates the precalculation.
func (r *Restricter) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	r.flow.Update(ctx, func(data map[dskey.Key][]byte, err error) {
		var found bool
		for key := range data {
			if r.hotKeys.Has(key) {
				found = true
				break
			}
		}
		if !found {
			log.Printf("Update on not hot key")
			updateFn(data, err)
			return
		}

		r.mu.Lock()
		defer r.mu.Unlock()
		ctx = collection.ContextWithRestrictCache(ctx)

		start := time.Now()
		calculatedKeys := make([]dskey.CollectionMode, 0, len(r.attributes))
		for k := range r.attributes {
			calculatedKeys = append(calculatedKeys, k)
		}

		if preError := r.precalculate(ctx, calculatedKeys); err != nil {
			err = errors.Join(err, preError)
		}
		log.Printf("Update on hot key: %d keys in %s", len(calculatedKeys), time.Since(start))

		// Send a signal to the autoupdate so the connections recalculate
		data[dskey.UpdateKey] = nil

		updateFn(data, err)
	})
}

// precalculate calculates the attributes for modes.
//
// Has to be called with a locked r.mu
func (r *Restricter) precalculate(ctx context.Context, collectionModes []dskey.CollectionMode) error {
	recorder := dsrecorder.New(r.flow)
	fetcher := dsfetch.New(recorder)

	byCollection := make(map[string][]dskey.CollectionMode)
	for _, collectionMode := range collectionModes {
		byCollection[collectionMode.Collection()] = append(byCollection[collectionMode.Collection()], collectionMode)
	}

	for name, collectionModes := range byCollection {
		restricter := collection.FromName(ctx, name)

		byMode := make(map[string][]int)
		for _, collectionMode := range collectionModes {
			mode := collectionMode.Mode()
			byMode[mode] = append(byMode[mode], collectionMode.ID())
		}

		for mode, ids := range byMode {
			modefunc := restricter.Modes(mode)
			if modefunc == nil {
				// TODO: Maybe log something, that there is a key without a modfunc (hast to be done when all restricters are implemented)
				continue
			}

			attrFunc, err := modefunc(ctx, fetcher, ids)
			if err != nil {
				return fmt.Errorf("precalculate %s/%s: %w", name, mode, err)
			}

			for i, id := range ids {
				// TODO do not use FromPArts but use the original object
				attr, err := dskey.CollectionModeFromParts(name, id, mode)
				if err != nil {
					return fmt.Errorf("invalid key: %w", err)
				}
				r.attributes[attr] = attrFunc[i]
			}
		}
	}

	r.hotKeys.Merge(recorder.Keys())

	return nil
}

// missingKeys returns keys, that are not in the attributes.
//
// Has to be called with a read lock or write lock on r.mu.
func (r *Restricter) missingModKeys(modes []dskey.CollectionMode) []dskey.CollectionMode {
	missing := make([]dskey.CollectionMode, 0, len(modes))
	for _, key := range modes {
		if _, ok := r.attributes[key]; !ok {
			missing = append(missing, key)
		}
	}
	return missing
}

func (r *Restricter) precalcMissing(ctx context.Context, modes []dskey.CollectionMode) error {
	r.mu.RLock()
	missingKeys := r.missingModKeys(modes)
	r.mu.RUnlock()

	if len(missingKeys) == 0 {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	// This has to be done again with a write lock. Only the write lock makes
	// sure, that there was no call to Update() in between.
	missingKeys = r.missingModKeys(modes)
	if len(missingKeys) == 0 {
		return nil
	}

	if err := r.precalculate(ctx, missingKeys); err != nil {
		return fmt.Errorf("restricter precalculate: %w", err)
	}

	return nil
}

func (r *Restricter) calculatedAttributes(ctx context.Context, keys []dskey.CollectionMode) ([]attribute.Func, error) {
	if err := r.precalcMissing(ctx, keys); err != nil {
		return nil, fmt.Errorf("precalculate missing: %w", err)
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]attribute.Func, len(keys))
	for i, key := range keys {
		result[i] = r.attributes[key]
	}

	return result, nil
}

// ForUser returns a getter that returns restricted data for an user id.
//
// Fetches keys from the flow and pre calculates the restriction for each key.
//
// TODO: Remove the ctx here and add it on every Get() call in the restricter
func (r *Restricter) ForUser(ctx context.Context, userID int) (context.Context, flow.Getter, *dsrecorder.Recorder) {
	recorder := dsrecorder.New(r.flow)

	ctx = perm.ContextWithGroupMap(ctx)

	return ctx, &userRestricter{
		userID:     userID,
		restricter: r,
		getter:     recorder,
	}, recorder
}

// userRestricter is a getter specific for an userID.
type userRestricter struct {
	userID     int
	restricter *Restricter
	getter     flow.Getter
}

func (r *userRestricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	start := time.Now()
	defer func() {
		log.Printf("full restrict %d keys took: %s", len(keys), time.Since(start))
	}()

	startIndexing := time.Now()
	modeKeys := make([]dskey.CollectionMode, len(keys))
	for i := range keys {
		modeKeys[i] = keys[i].CollectionMode()
	}
	log.Printf("building mod keys took: %s", time.Since(startIndexing))

	startUser := time.Now()
	user, err := attribute.NewUserAttributes(ctx, r.getter, r.userID)
	if err != nil {
		return nil, fmt.Errorf("calculate user permission: %w", err)
	}
	log.Printf("building user took: %s", time.Since(startUser))

	startData := time.Now()
	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("fetch full data: %w", err)
	}
	log.Printf("fetching %d keys took: %s", len(keys), time.Since(startData))

	startModeKeys := time.Now()
	attrFuncs, err := r.restricter.calculatedAttributes(ctx, modeKeys)
	if err != nil {
		return nil, fmt.Errorf("get precalculated functions: %w", err)
	}
	log.Printf("getting mode funcs took: %s", time.Since(startModeKeys))

	startRestrict := time.Now()
	for i, key := range keys {
		checkAndRemove(data, key, attrFuncs[i], user)
	}
	log.Printf("precalculated restrict %d keys took: %s", len(keys), time.Since(startRestrict))

	startRelation := time.Now()
	for _, key := range keys {
		if key.RelationType() == dskey.RelationNone || data[key] == nil {
			continue
		}

		switch key.RelationType() {
		case dskey.RelationSingle:
			id, err := fastjson.DecodeInt(data[key])
			if err != nil {
				return nil, fmt.Errorf("relation: invalid value in %s, expected id: %w", key, err)
			}

			relatedKey, err := key.RelationTo(id)
			if err != nil {
				return nil, fmt.Errorf("get relation from key %s: %w", key, err)
			}

			attrFuncs, err := r.restricter.calculatedAttributes(ctx, []dskey.CollectionMode{relatedKey.CollectionMode()})
			if err != nil {
				return nil, fmt.Errorf("get precalculated functions: %w", err)
			}

			checkAndRemove(data, key, attrFuncs[0], user)

		case dskey.RelationGenericSingle:
			collection, id, err := parseGenericValue(data[key])
			if err != nil {
				return nil, fmt.Errorf("parse generic relation value %s: %w", key, err)
			}

			relatedKey, err := key.RelationGenericTo(collection, id)
			if err != nil {
				return nil, fmt.Errorf("get related mode from generic value in %s: %w", key, err)
			}

			attrFuncs, err := r.restricter.calculatedAttributes(ctx, []dskey.CollectionMode{relatedKey.CollectionMode()})
			if err != nil {
				return nil, fmt.Errorf("get precalculated functions: %w", err)
			}

			checkAndRemove(data, key, attrFuncs[0], user)

		case dskey.RelationList:
			ids, err := fastjson.DecodeIntList(data[key])
			if err != nil {
				return nil, fmt.Errorf("relation-list: invalid value in %s: `%s`, expected list of ids: %w", key, data[key], err)
			}

			relationModes := make([]dskey.CollectionMode, len(ids))
			for i, id := range ids {
				relatedKey, err := key.RelationTo(id)
				if err != nil {
					return nil, fmt.Errorf("get relation from key %s: %w", key, err)
				}
				relationModes[i] = relatedKey.CollectionMode()
			}

			attrFuncs, err := r.restricter.calculatedAttributes(ctx, relationModes)
			if err != nil {
				return nil, fmt.Errorf("get precalculated functions for relation list key %s: %w", key, err)
			}

			newList := make([]int, 0, len(ids))
			for i := range ids {
				if attrFuncs[i] == nil || !attrFuncs[i](user) {
					if attrFuncs[i] == nil {
						log.Printf("attrFunc for key %s, collection field %s, is nil", key, key.CollectionField())
					}
					continue
				}
				newList = append(newList, ids[i])
			}

			if len(newList) == len(ids) {
				continue
			}

			newValue, err := json.Marshal(newList)
			if err != nil {
				return nil, fmt.Errorf("marshal new value for key %s: %w", key, err)
			}
			data[key] = newValue

		case dskey.RelationGenericList:
			var listValue []json.RawMessage
			if err := json.Unmarshal(data[key], &listValue); err != nil {
				return nil, fmt.Errorf("generic-relation-list: invalid value in %s: `%s`, expected list: %w", key, data[key], err)
			}

			relationModes := make([]dskey.CollectionMode, len(listValue))
			for i, contentObectID := range listValue {
				collection, id, err := parseGenericValue(contentObectID)
				if err != nil {
					return nil, fmt.Errorf("parse generic relation value %s: %w", key, err)
				}

				relatedKey, err := key.RelationGenericTo(collection, id)
				if err != nil {
					log.Printf("WARNING: key %s has value %s, but collection %s is not a possible relation for that key", key, data[key], collection)
					// TODO: Do not return but remove value
					return nil, fmt.Errorf("invalid data")

				}
				relationModes[i] = relatedKey.CollectionMode()
			}

			attrFuncs, err := r.restricter.calculatedAttributes(ctx, relationModes)
			if err != nil {
				return nil, fmt.Errorf("get precalculated functions for relation list key %s: %w", key, err)
			}

			newList := make([]json.RawMessage, 0, len(listValue))
			for i := range listValue {
				if attrFuncs[i] == nil || !attrFuncs[i](user) {
					if attrFuncs[i] == nil {
						log.Printf("attrFunc for key %s, collection field %s, is nil", key, key.CollectionField())
					}
					continue
				}
				newList = append(newList, listValue[i])
			}

			if len(newList) == len(listValue) {
				continue
			}

			newValue, err := json.Marshal(newList)
			if err != nil {
				return nil, fmt.Errorf("marshal new value for key %s: %w", key, err)
			}
			data[key] = newValue
		}

	}
	log.Printf("modify relations %d keys took: %s", len(keys), time.Since(startRelation))

	return data, nil
}

func checkAndRemove(data map[dskey.Key][]byte, key dskey.Key, attrFunc attribute.Func, user attribute.UserAttributes) bool {
	if attrFunc == nil || !attrFunc(user) {
		if attrFunc == nil {
			log.Printf("attrFunc for key %s, collection field %s, is nil", key, key.CollectionField())
		}
		data[key] = nil
		return true
	}

	return false
}

func parseGenericValue(v []byte) (string, int, error) {
	var genericID string
	if err := json.Unmarshal(v, &genericID); err != nil {
		return "", 0, fmt.Errorf("decoding value: %w", err)
	}

	collection, rawID, found := strings.Cut(genericID, "/")
	if !found {
		return "", 0, fmt.Errorf("invalid generic value. No /")
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		return "", 0, fmt.Errorf("invalid generic value. Second part is no int: %w", err)
	}

	return collection, id, nil
}
