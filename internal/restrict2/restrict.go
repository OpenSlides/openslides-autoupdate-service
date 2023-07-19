package restrict

import (
	"context"
	"errors"
	"fmt"
	"log"

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
	attributes *attribute.Map
	hotKeys    set.Set[dskey.Key]

	// TODO: Remove me
	implementedCollections set.Set[string]
}

// New initializes a restricter
func New(flow flow.Flow) *Restricter {
	return &Restricter{
		flow:       flow,
		attributes: attribute.NewMap(),
		hotKeys:    set.New[dskey.Key](),

		implementedCollections: set.New("agenda_item"),
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
			updateFn(data, err)
			return
		}

		if preError := r.precalculate(ctx, r.attributes.Keys()); err != nil {
			err = errors.Join(err, preError)
		}

		updateFn(data, err)
	})
}

// ForUser returns a getter that returns restricted data for an user id.
//
// Fetches keys from the flow and pre calculates the restriction for each key.
//
// TODO: Remove the ctx here and add it on every Get() call in the restricter
func (r *Restricter) ForUser(ctx context.Context, userID int) (context.Context, flow.Getter) {
	ctx, todoOldRestricter := oldRestrict.Middleware(ctx, r.flow, userID)
	return ctx, &restrictedGetter{
		todoOldRestricter: todoOldRestricter,
		userID:            userID,
		restricter:        r,
	}
}

// precalculate calculates the attributes for modes.
func (r *Restricter) precalculate(ctx context.Context, collectionModes []dskey.Key) error {
	// TODO: Make concurency save
	if len(collectionModes) == 0 {
		return nil
	}

	recorder := dsrecorder.New(r.flow)
	fetcher := dsfetch.New(recorder)

	byCollection := make(map[string][]dskey.Key)
	for _, collectionMode := range collectionModes {
		byCollection[collectionMode.Collection] = append(byCollection[collectionMode.Collection], collectionMode)
	}

	ctx = perm.ContextWithGroupCache(ctx)

	for name, collectionModes := range byCollection {
		restricter := collection.Collection(ctx, name)

		byMode := make(map[string][]int)
		for _, collectionMode := range collectionModes {
			byMode[collectionMode.Field] = append(byMode[collectionMode.Field], collectionMode.ID)
		}

		for mode, ids := range byMode {
			modefunc := restricter.Modes(mode)
			if err := modefunc(ctx, fetcher, r.attributes, ids); err != nil {
				return fmt.Errorf("precalculate %s/%s: %w", name, mode, err)
			}
		}
	}

	r.hotKeys.Merge(recorder.Keys())

	return nil
}

type restrictedGetter struct {
	todoOldRestricter flow.Getter
	userID            int
	restricter        *Restricter
}

func (r *restrictedGetter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	keyToMode := make(map[dskey.Key]dskey.Key, len(keys))
	modeKeys := set.New[dskey.Key]()
	for _, key := range keys {
		mode, ok := restrictionModes[key.CollectionField()]
		if !ok {
			log.Printf("no restriction for %s", key.CollectionField())
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

	fetcher := dsfetch.New(r.restricter.flow)
	orgaLevel, groupIDs, err := userPermissions(ctx, fetcher, r.userID)
	if err != nil {
		return nil, fmt.Errorf("calculate user permission: %w", err)
	}

	data, err := r.restricter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("fetch full data: %w", err)
	}

	attributes := r.restricter.attributes.Get(modeKeys.List()...)

	var oldKeys []dskey.Key
	for key := range data {
		if !r.restricter.implementedCollections.Has(key.Collection) {
			oldKeys = append(oldKeys, key)
			continue
		}

		if !allowedByAttr(attributes[keyToMode[key]], r.userID, orgaLevel, groupIDs) {
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

// UserPermissions returns the global permission and all group ids of an user.
func userPermissions(ctx context.Context, fetcher *dsfetch.Fetch, userID int) (perm.OrganizationManagementLevel, []int, error) {
	var groupIDHelper [][]int
	if userID != 0 {
		meetingIDs, err := fetcher.User_GroupIDsTmpl(userID).Value(ctx)
		if err != nil {
			return "", nil, fmt.Errorf("getting meetings of user: %w", err)
		}

		groupIDHelper = make([][]int, len(meetingIDs))
		for i := 0; i < len(meetingIDs); i++ {
			fetcher.User_GroupIDs(userID, meetingIDs[i]).Lazy(&groupIDHelper[i])
		}
	}

	var globalPermStr string
	fetcher.User_OrganizationManagementLevel(userID).Lazy(&globalPermStr)

	if err := fetcher.Execute(ctx); err != nil {
		return "", nil, fmt.Errorf("getting groupIDs: %w", err)
	}

	var groupIDs []int
	for _, ids := range groupIDHelper {
		groupIDs = append(groupIDs, ids...)
	}

	orgaLevel := perm.OrganizationManagementLevel(globalPermStr)

	return orgaLevel, groupIDs, nil
}

// allowedByAttr tells if the user is allowed to see the attribute.
func allowedByAttr(requiredAttr *attribute.Attribute, uid int, orgaLevel perm.OrganizationManagementLevel, groupIDs []int) bool {
	if requiredAttr == nil {
		return false
	}

	if requiredAttr.GlobalPermission == 255 {
		return false
	}

	if allowedByGlobalPerm(orgaLevel, uid, requiredAttr.GlobalPermission) {
		return true
	}

	if allowedByGroup(requiredAttr.GroupIDs, groupIDs) {
		// if nextAttr := requiredAttr.GroupAnd; nextAttr != nil {
		// 	return allowedByAttr(nextAttr, uid, orgaLevel, groupIDs)
		// }
		return true
	}

	if allowedByUser(requiredAttr.UserIDs, uid) {
		return true
	}

	return false
}

// allowedByGlobalPerm tells, if the user has the correct global permission.
func allowedByGlobalPerm(orgaLevel perm.OrganizationManagementLevel, userID int, globalAttr attribute.GlobalPermission) bool {
	return globalAttr == attribute.GlobalAll ||
		globalAttr == attribute.GlobalLoggedIn && userID > 0 ||
		globalAttr == attribute.GlobalSuperadmin && orgaLevel == perm.OMLSuperadmin ||
		globalAttr == attribute.GlobalCanManageOrganization && (orgaLevel == perm.OMLSuperadmin || orgaLevel == perm.OMLCanManageOrganization) ||
		globalAttr == attribute.GlobalCanManageUsers && (orgaLevel == perm.OMLSuperadmin || orgaLevel == perm.OMLCanManageOrganization || orgaLevel == perm.OMLCanManageUsers)
}

func allowedByGroup(requiredGroups set.Set[int], hasGroups []int) bool {
	for _, gid := range hasGroups {
		if requiredGroups.Has(gid) {
			return true
		}
	}

	return false
}

func allowedByUser(requiredIDs set.Set[int], uid int) bool {
	return requiredIDs.Has(uid)
}
