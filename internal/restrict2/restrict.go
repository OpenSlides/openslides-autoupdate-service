package restrict

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	oldRestrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsrecorder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Restricter TODO
type Restricter struct {
	flow flow.Flow
	// TODO: Probably needs a mutex
	attributes *attrMap
	hotKeys    set.Set[dskey.Key]

	// TODO: Remove me
	implementedCollections set.Set[string]
}

// New initializes a restricter
func New(flow flow.Flow) *Restricter {
	return &Restricter{
		flow:       flow,
		attributes: newAttrMap(),
		hotKeys:    set.New[dskey.Key](),

		implementedCollections: set.New(
			"agenda_item",
			"motion",
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

		start := time.Now()
		if preError := r.precalculate(ctx, r.attributes.Keys()); err != nil {
			err = errors.Join(err, preError)
		}
		log.Printf("Update on hot key: %d keys in %s", r.attributes.Len(), time.Since(start))

		// Send a signal to the autoupdate so the connections recalculate
		data[dskey.Key{Collection: "meta", ID: 1, Field: "update"}] = nil

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
	return ctx, &restrictedGetter{
		todoOldRestricter: todoOldRestricter,
		userID:            userID,
		restricter:        r,
		getter:            recorder,
	}, recorder
}

func slice[T any](list []T, c int) []T {
	if len(list) <= c {
		return list
	}
	return list[:c]
}

// precalculate calculates the attributes for modes.
func (r *Restricter) precalculate(ctx context.Context, collectionModes []dskey.Key) error {
	// TODO: Make concurency save
	if len(collectionModes) == 0 {
		return nil
	}

	log.Printf("precalculate %d modes: %v", len(collectionModes), slice(collectionModes, 10))

	recorder := dsrecorder.New(r.flow)
	fetcher := dsfetch.New(recorder)

	byCollection := make(map[string][]dskey.Key)
	for _, collectionMode := range collectionModes {
		byCollection[collectionMode.Collection] = append(byCollection[collectionMode.Collection], collectionMode)
	}

	// Put all tuples together so they can be added at once (with one lock)
	var allTuples []collection.Tuple
	for name, collectionModes := range byCollection {
		restricter := collection.Collection(ctx, name)

		byMode := make(map[string][]int)
		for _, collectionMode := range collectionModes {
			byMode[collectionMode.Field] = append(byMode[collectionMode.Field], collectionMode.ID)
		}

		for mode, ids := range byMode {
			modefunc := restricter.Modes(mode)
			tuples, err := modefunc(ctx, fetcher, ids)
			if err != nil {
				return fmt.Errorf("precalculate %s/%s: %w", name, mode, err)
			}
			allTuples = append(allTuples, tuples...)
		}
	}

	r.attributes.Add(allTuples...)
	r.hotKeys.Merge(recorder.Keys())

	return nil
}

type restrictedGetter struct {
	todoOldRestricter flow.Getter
	userID            int
	restricter        *Restricter
	getter            flow.Getter
}

func (r *restrictedGetter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	start := time.Now()
	defer func() {
		log.Printf("full restrict %d keys took: %s", len(keys), time.Since(start))
	}()

	relevantKey := make([]dskey.Key, 0, len(keys))
	for _, key := range keys {
		if !(key.Collection == "agenda_item" || key.Collection == "motion") {
			// TODO: Remove me
			continue
		}

		_, ok := restrictionModes[key.CollectionField()]
		if !ok {
			// TODO
			// log.Printf("no restriction for %s", key.CollectionField())
			continue
		}

		relevantKey = append(relevantKey, key)
	}

	startIndexing := time.Now()
	modeKeys := make([]dskey.Key, len(relevantKey))
	for i, key := range relevantKey {
		mode := restrictionModes[key.CollectionField()]

		modeKey := dskey.Key{Collection: key.Collection, ID: key.ID, Field: mode}
		modeKeys[i] = modeKey
	}
	log.Printf("indexing took: %s", time.Since(startIndexing))

	startPrecalc := time.Now()
	needPrecalculate := r.restricter.attributes.NeedCalc(modeKeys)
	if err := r.restricter.precalculate(ctx, needPrecalculate); err != nil {
		return nil, fmt.Errorf("restricter precalculate: %w", err)
	}
	log.Printf("precalcing took: %s", time.Since(startPrecalc))

	// Check the permissions from here

	startUser := time.Now()
	user, err := buildUserAttributes(ctx, r.getter, r.userID)
	if err != nil {
		return nil, fmt.Errorf("calculate user permission: %w", err)
	}
	log.Printf("building user took: %s", time.Since(startUser))

	startModeKeys := time.Now()
	attrFuncs := r.restricter.attributes.Get(modeKeys...)
	log.Printf("getting mode funcs took: %s", time.Since(startModeKeys))

	startRestrict := time.Now()
	var oldKeys []dskey.Key // TODO: Remove me. This is only necessary for restrict1

	allowedKeys := make([]dskey.Key, 0, len(keys))

	for _, key := range keys {
		if !r.restricter.implementedCollections.Has(key.Collection) {
			oldKeys = append(oldKeys, key)
			continue
		}

		mode := restrictionModes[key.CollectionField()]
		modeKey := dskey.Key{Collection: key.Collection, ID: key.ID, Field: mode}

		attrFunc := attrFuncs[modeKey]
		if attrFunc == nil {
			log.Printf("attrFunc for key %s, mode %s, is nil", key, modeKey)
			continue
		}

		if !attrFunc(user) {
			continue
		}

		allowedKeys = append(allowedKeys, key)

		// TODO: relation fields
	}
	log.Printf("precalculated restrict %d keys took: %s", len(keys), time.Since(startRestrict))

	startData := time.Now()
	data, err := r.getter.Get(ctx, allowedKeys...)
	if err != nil {
		return nil, fmt.Errorf("fetch full data: %w", err)
	}
	log.Printf("fetching %d keys took: %s", len(allowedKeys), time.Since(startData))

	startOld := time.Now()
	if len(oldKeys) > 0 {
		oldData, err := r.todoOldRestricter.Get(ctx, oldKeys...)
		if err != nil {
			return nil, fmt.Errorf("old restricter: %w", err)
		}

		for k, v := range oldData {
			data[k] = v
		}
	}
	log.Printf("old restricter for %d keys took: %s", len(oldKeys), time.Since(startOld))

	return data, nil
}

func buildUserAttributes(ctx context.Context, getter flow.Getter, userID int) (attribute.UserAttributes, error) {
	var zero attribute.UserAttributes
	fetcher := dsfetch.New(getter)

	if userID == 0 {
		return zero, nil
	}

	var meetingIDs []int
	var globalLevelStr string
	fetcher.User_OrganizationManagementLevel(userID).Lazy(&globalLevelStr)
	fetcher.User_GroupIDsTmpl(userID).Lazy(&meetingIDs)

	if err := fetcher.Execute(ctx); err != nil {
		return zero, fmt.Errorf("getting meeting ids and global level for user %d: %w", userID, err)
	}

	groupIDList := make([][]int, len(meetingIDs))
	for i, mid := range meetingIDs {
		fetcher.User_GroupIDs(userID, mid).Lazy(&groupIDList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return zero, fmt.Errorf("getting group IDs %d: %w", userID, err)
	}

	groupIDs := set.New[int]()
	for _, idList := range groupIDList {
		groupIDs.Add(idList...)
	}

	return attribute.UserAttributes{
		UserID:    userID,
		GroupIDs:  groupIDs,
		OrgaLevel: perm.OrganizationManagementLevel(globalLevelStr),
	}, nil
}
