package restrict

import (
	"context"
	"fmt"

	oldRestrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

// Restricter TODO
type Restricter struct {
	flow flow.Flow
	// TODO: Probably needs a mutex
	attributes *attribute.Map
}

// New initializes a restricter
func New(flow flow.Flow) *Restricter {
	return &Restricter{
		flow:       flow,
		attributes: attribute.NewMap(),
	}
}

// Get returns the full (unrestricted) data.
func (r *Restricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	return r.flow.Get(ctx, keys...)
}

// Update updates the precalculation.
func (r *Restricter) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	r.flow.Update(ctx, updateFn)
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

func (r *Restricter) precalculate(ctx context.Context, keys []dskey.Key) error {
	// TODO: Make concurency save
	if len(keys) == 0 {
		return nil
	}

	ctx = collection.ContextWithRestrictCache(ctx)
	// TODO: Use a counter to get requested keys.
	fetcher := dsfetch.New(r.flow)

	// Group by collection
	byCollection := make(map[string][]dskey.Key)
	for _, key := range keys {
		byCollection[key.Collection] = append(byCollection[key.Collection], key)
	}

	for name, keys := range byCollection {
		restricter := collection.Collection(ctx, name)

		byMode := make(map[string][]int)
		for _, key := range keys {
			byMode[restrictionModes[key.CollectionField()]] = append(byMode[restrictionModes[key.CollectionField()]], key.ID)
		}

		for mode, ids := range byMode {
			// TODO: This sets the locks very ofeten. It would be better, if the lock on r.attribute would set on the start of precalculate
			if err := restricter.Modes(mode)(ctx, fetcher, r.attributes, ids); err != nil {
				return fmt.Errorf("precalculate %s/%s: %w", name, mode, err)
			}
		}
	}

	return nil
}

type restrictedGetter struct {
	todoOldRestricter flow.Getter
	userID            int
	restricter        *Restricter
}

func (r *restrictedGetter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	needPrecalculate := r.restricter.attributes.NeedCalc(keys)

	if err := r.restricter.precalculate(ctx, needPrecalculate); err != nil {
		return nil, fmt.Errorf("restricter precalculate: %w", err)
	}

	fetcher := dsfetch.New(r.restricter.flow)
	orgaLevel, groupIDs, err := userPermissions(ctx, fetcher, r.userID)
	if err != nil {
		return nil, fmt.Errorf("calculate user permission: %w", err)
	}

	for _, key := range keys {
		mode, ok := restrictionModes[key.CollectionField()]
		if !ok {
			return nil, fmt.Errorf("unknown restrictoin mode for %s", key.CollectionField())
		}

		requiredAttr, err := r.restricter.attributes.Get(ctx, fetcher, mperms, dskey.Key{Collection: key.Collection, ID: key.ID, Field: restrictionMode})
		if err != nil {
			return nil, fmt.Errorf("attributemap: %w", err)
		}

		if requiredAttr == nil {
			fmt.Printf("TODO: did not found attr for %s (%s)\n", key, restrictionMode)
			continue
		}

		if !AllowedByAttr(requiredAttr, uid, orgaLevel, groupIDs) {
			data[key] = nil
			continue
		}

		// TODO: relation fields

	}

	// TODO: only send not implemented collections
	return r.todoOldRestricter.Get(ctx, keys...)
}

var implementedCollections = []string{
	"agenda_item",
}

// UserPermissions returns the global permission and all group ids of an user.
func userPermissions(ctx context.Context, fetcher *dsfetch.Fetch, uid int) (perm.OrganizationManagementLevel, []int, error) {
	globalPermStr, err := fetcher.User_OrganizationManagementLevel(uid).Value(ctx)
	if err != nil {
		return "", nil, fmt.Errorf("getting orga level from user: %w", err)
	}

	groupIDs, err := userGroups(ctx, fetcher, uid)
	if err != nil {
		return "", nil, fmt.Errorf("getting group ids: %w", err)
	}

	orgaLevel := perm.OrganizationManagementLevel(globalPermStr)

	return orgaLevel, groupIDs, nil
}

// userGroups returns all groups from a user.
func userGroups(ctx context.Context, fetcher *dsfetch.Fetch, uid int) ([]int, error) {
	meetingIDs, err := fetcher.User_GroupIDsTmpl(uid).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting meetings of user: %w", err)
	}

	groupIDs := make([][]int, len(meetingIDs))
	for i := 0; i < len(meetingIDs); i++ {
		fetcher.User_GroupIDs(uid, meetingIDs[i]).Lazy(&groupIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting groupIDs: %w", err)
	}

	var result []int
	for _, ids := range groupIDs {
		result = append(result, ids...)
	}

	return result, nil
}
