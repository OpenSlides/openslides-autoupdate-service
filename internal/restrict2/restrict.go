package restrict

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	oldRestrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsrecorder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Restricter TODO
type Restricter struct {
	flow flow.Flow

	mu         sync.RWMutex
	attributes map[dskey.Key]attribute.Func
	hotKeys    set.Set[dskey.Key]

	// TODO: Remove me
	implementedCollections set.Set[string]
}

// New initializes a restricter
func New(flow flow.Flow) *Restricter {
	return &Restricter{
		flow:       flow,
		attributes: make(map[dskey.Key]attribute.Func),
		hotKeys:    set.New[dskey.Key](),

		implementedCollections: set.New(
			"agenda_item",
			"group",
			"motion",
			"motion_submitter",
			"list_of_speakers",
			"meeting",
			"motion_state",
			"motion_workflow",
			"organization",
			"projector",
			"user",
			"theme",
		),
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

		start := time.Now()
		calculatedKeys := make([]dskey.Key, 0, len(r.attributes))
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

// ForUser returns a getter that returns restricted data for an user id.
//
// Fetches keys from the flow and pre calculates the restriction for each key.
//
// TODO: Remove the ctx here and add it on every Get() call in the restricter
func (r *Restricter) ForUser(ctx context.Context, userID int) (context.Context, flow.Getter, *dsrecorder.Recorder) {
	recorder := dsrecorder.New(r.flow)

	ctx = perm.ContextWithGroupMap(ctx)

	ctx, todoOldRestricter := oldRestrict.Middleware(ctx, recorder, userID)
	return ctx, &userRestricter{
		todoOldRestricter: todoOldRestricter,
		userID:            userID,
		restricter:        r,
		getter:            recorder,
	}, recorder
}

// precalculate calculates the attributes for modes.
//
// Has to be called with a locked r.mu
func (r *Restricter) precalculate(ctx context.Context, collectionModes []dskey.Key) error {
	recorder := dsrecorder.New(r.flow)
	fetcher := dsfetch.New(recorder)

	byCollection := make(map[string][]dskey.Key)
	for _, collectionMode := range collectionModes {
		byCollection[collectionMode.Collection()] = append(byCollection[collectionMode.Collection()], collectionMode)
	}

	for name, collectionModes := range byCollection {
		restricter := collection.FromName(ctx, name)

		byMode := make(map[string][]int)
		for _, collectionMode := range collectionModes {
			mode := collectionMode.Field()
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
				attr, err := dskey.FromParts(name, id, mode)
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
func (r *Restricter) missingModKeys(keys []dskey.Key) []dskey.Key {
	missing := make([]dskey.Key, 0, len(keys))
	for _, key := range keys {
		if _, ok := r.attributes[key]; !ok {
			missing = append(missing, key)
		}
	}
	return missing
}

func (r *Restricter) precalcMissing(ctx context.Context, keys []dskey.Key) error {
	r.mu.RLock()
	missingKeys := r.missingModKeys(keys)
	r.mu.RUnlock()

	if len(missingKeys) == 0 {
		return nil
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	// This has to be done again with a write lock. Only the write lock makes
	// sure, that there was no call to Update() in between.
	missingKeys = r.missingModKeys(keys)
	if len(missingKeys) == 0 {
		return nil
	}

	if err := r.precalculate(ctx, missingKeys); err != nil {
		return fmt.Errorf("restricter precalculate: %w", err)
	}

	return nil
}

func (r *Restricter) calculatedAttributes(ctx context.Context, keys []dskey.Key) ([]attribute.Func, error) {
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

// userRestricter is a getter specific for an userID.
type userRestricter struct {
	todoOldRestricter flow.Getter
	userID            int
	restricter        *Restricter
	getter            flow.Getter
}

func (r *userRestricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	// start := time.Now()
	// defer func() {
	// 	log.Printf("full restrict %d keys took: %s", len(keys), time.Since(start))
	// }()

	// startIndexing := time.Now()
	modeKeys := make([]dskey.Key, len(keys))
	for i, key := range keys {
		// TODO: This would be faster with a autogenerated switch statement.
		// Or when there was a key.Mode() method with an autogenerated value
		mode := restrictionModes[key.CollectionField()]

		modeKey, err := dskey.FromParts(key.Collection(), key.ID(), mode)
		if err != nil {
			return nil, fmt.Errorf("invalid modekey: %w", err)
		}
		modeKeys[i] = modeKey
	}
	// log.Printf("building mod keys took: %s", time.Since(startIndexing))

	// Check the permissions from here

	// startUser := time.Now()
	user, err := attribute.NewUserAttributes(ctx, r.getter, r.userID)
	if err != nil {
		return nil, fmt.Errorf("calculate user permission: %w", err)
	}
	// log.Printf("building user took: %s", time.Since(startUser))

	// startData := time.Now()
	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("fetch full data: %w", err)
	}
	// log.Printf("fetching %d keys took: %s", len(keys), time.Since(startData))

	// startModeKeys := time.Now()
	attrFuncs, err := r.restricter.calculatedAttributes(ctx, modeKeys)
	if err != nil {
		return nil, fmt.Errorf("get precalculated functions: %w", err)
	}
	// log.Printf("getting mode funcs took: %s", time.Since(startModeKeys))

	// startRestrict := time.Now()
	var oldKeys []dskey.Key // TODO: Remove me. This is only necessary for old restrict

	for i, key := range keys {
		if !r.restricter.implementedCollections.Has(key.Collection()) {
			oldKeys = append(oldKeys, key)
			continue
		}

		attrFunc := attrFuncs[i]
		if attrFunc == nil {
			log.Printf("attrFunc for key %s, collection field %s, is nil", key, key.CollectionField())
			continue
		}

		if !attrFunc(user) {
			data[key] = nil
			continue
		}

		// TODO: relation fields
	}
	// log.Printf("precalculated restrict %d keys took: %s", len(keys), time.Since(startRestrict))

	// coundOld := make(map[string]int)
	// for _, key := range oldKeys {
	// 	coundOld[key.Collection]++
	// }
	// fmt.Println(coundOld)

	// startOld := time.Now()
	if len(oldKeys) > 0 {
		oldData, err := r.todoOldRestricter.Get(ctx, oldKeys...)
		if err != nil {
			return nil, fmt.Errorf("old restricter: %w", err)
		}

		for k, v := range oldData {
			data[k] = v
		}
	}
	// log.Printf("old restricter for %d keys took: %s", len(oldKeys), time.Since(startOld))

	return data, nil
}
