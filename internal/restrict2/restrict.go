package restrict

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
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

	ctx, todoOldRestricter := oldRestrict.Middleware(ctx, recorder, userID)
	return ctx, &restrictedGetter{
		todoOldRestricter: todoOldRestricter,
		userID:            userID,
		restricter:        r,
		getter:            recorder,
	}, recorder
}

// precalculate calculates the attributes for modes.
func (r *Restricter) precalculate(ctx context.Context, collectionModes []dskey.Key) error {
	// TODO: Make concurency save
	if len(collectionModes) == 0 {
		return nil
	}

	log.Printf("precalculate: %v", collectionModes)

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
	keyToMode := make(map[dskey.Key]dskey.Key, len(keys))
	modeKeys := set.New[dskey.Key]()
	for _, key := range keys {
		if !r.restricter.implementedCollections.Has(key.Collection) {
			// TODO: Remove me
			continue
		}

		mode, ok := restrictionModes[key.CollectionField()]
		if !ok {
			// TODO
			// log.Printf("no restriction for %s", key.CollectionField())
			continue
		}
		modeKey := dskey.Key{Collection: key.Collection, ID: key.ID, Field: mode}
		modeKeys.Add(modeKey)
		keyToMode[key] = modeKey
	}

	needPrecalculate := r.restricter.attributes.NeedCalc(modeKeys.List())

	if err := r.restricter.precalculate(ctx, needPrecalculate); err != nil {
		return nil, fmt.Errorf("restricter precalculate: %w", err)
	}

	// Check the permissions from here

	user, err := buildUserAttributes(ctx, r.getter, r.userID)
	if err != nil {
		return nil, fmt.Errorf("calculate user permission: %w", err)
	}

	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("fetch full data: %w", err)
	}

	attrFuncs := r.restricter.attributes.Get(modeKeys.List()...)

	var oldKeys []dskey.Key // TODO: Remove me. This is only necessary for restrict1
	for key := range data {
		if !r.restricter.implementedCollections.Has(key.Collection) {
			oldKeys = append(oldKeys, key)
			continue
		}

		attrFunc := attrFuncs[keyToMode[key]]
		if attrFunc == nil {
			log.Printf("attrFunc for key %s, mode %s, is nil", key, keyToMode[key])
			data[key] = nil
			continue
		}

		if !attrFunc(user) {
			data[key] = nil
			continue
		}

		// TODO: relation fields
	}

	if len(oldKeys) > 0 {
		oldData, err := r.todoOldRestricter.Get(ctx, oldKeys...)
		if err != nil {
			return nil, fmt.Errorf("old restricter: %w", err)
		}

		for k, v := range oldData {
			data[k] = v
		}
	}

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

	meetingPermission := perm.NewMeetingPermission(fetcher, userID)
	meetingGetter := func(meetingID int) *perm.Permission {
		perm, err := meetingPermission.Meeting(ctx, meetingID)
		if err != nil {
			oserror.Handle(fmt.Errorf("getting meeting: %w", err))
			return nil
		}

		return perm
	}

	return attribute.UserAttributes{
		UserID:          userID,
		GetMeetingPerms: meetingGetter,
		OrgaLevel:       perm.OrganizationManagementLevel(globalLevelStr),
	}, nil
}
