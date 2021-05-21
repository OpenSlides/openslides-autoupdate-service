// Package permission handels the permissions of a user.
//
// It provides the Permission object with methods to tell, if a user is allowed
// to use an action or can see some fqfields.
package permission

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Permission provides methods to tell, if a user can use an action or can see
// fqfields. It has to be initializes with permission.New().
type Permission struct {
	hs *handlerStore

	dp dataprovider.DataProvider
}

// New returns a new permission service.
//
// It requires a permission.DataProvider to access the database.
func New(dp DataProvider) *Permission {
	p := &Permission{
		hs: newHandlerStore(),
		dp: dataprovider.DataProvider{External: dp},
	}

	for _, con := range openSlidesCollections(p.dp) {
		con.Connect(p.hs)
	}

	return p
}

// IsAllowed returns true, if the user can access the given action.
//
// One call to IsAllowed() handels a list of requests to this action. For each
// entry in the payloadList the method checks, if the user is allowed to use the
// action. The method returns true, if the user can the action for all of the
// given payloads.
func (ps *Permission) IsAllowed(ctx context.Context, action string, userID int, payloadList []map[string]json.RawMessage) (bool, error) {
	superadmin, err := ps.dp.IsSuperadmin(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("checking for superadmin: %w", err)
	}
	if superadmin {
		return true, nil
	}

	// TODO: after all handlers are implemented. Move this code above the superadmin check.
	handler, ok := ps.hs.actions[action]
	if !ok {
		return false, fmt.Errorf("unknown action: `%s`", action)
	}

	for i, payload := range payloadList {
		allowed, err := handler.IsAllowed(ctx, userID, payload)
		if err != nil {
			bs, jsonErr := json.Marshal(payload)
			if jsonErr != nil {
				bs = []byte("[payload can not be encoded]")
			}
			return false, fmt.Errorf("action: %s, payload-index %d: `%s`: %w", action, i, bs, err)
		}
		if !allowed {
			return false, nil
		}
	}

	return true, nil
}

// superadminFields handles fields that the superadmin is not allowed to see.
//
// Returns true, if the normal normal restricters should be skiped.
func superadminFields(result map[string]bool, collection string, fqfields []perm.FQField) (skip bool) {
	if collection == "personal_note" {
		return false
	}

	for _, k := range fqfields {
		if k.Collection == "user" && k.Field == "password" {
			continue
		}
		result[k.String()] = true
	}
	return true
}

// RestrictFQFields filters a list of fqfields and returns the fields, that the
// user can see.
//
// The return value is a set of fqfields. It can only contain fields, that where
// requested.
func (ps Permission) RestrictFQFields(ctx context.Context, userID int, fqfields []string) (map[string]bool, error) {
	allowedFields := make(map[string]bool, len(fqfields))

	superadmin, err := ps.dp.IsSuperadmin(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("checking for superadmin: %w", err)
	}

	grouped, err := groupFQFields(fqfields)
	if err != nil {
		return nil, fmt.Errorf("grouping fqfields: %w", err)
	}

	for name, fqfields := range grouped {
		if superadmin {
			if superadminFields(allowedFields, name, fqfields) {
				continue
			}
		}

		handler, ok := ps.hs.collections[name]
		if !ok {
			return nil, fmt.Errorf("unknown collection: `%s`", name)
		}

		if err := handler.RestrictFQFields(ctx, userID, fqfields, allowedFields); err != nil {
			return nil, fmt.Errorf("restrict for collection %s: %w", name, err)
		}
	}
	return allowedFields, nil
}

// groupFQFields sorts the fqfields and returns it grouped by collection.
func groupFQFields(fqfields []string) (map[string][]perm.FQField, error) {
	grouped := make(map[string][]perm.FQField)
	for _, f := range fqfields {
		fqfield, err := perm.ParseFQField(f)
		if err != nil {
			return nil, fmt.Errorf("decoding fqfield: %w", err)
		}
		grouped[fqfield.Collection] = append(grouped[fqfield.Collection], fqfield)
	}
	return grouped, nil
}

// AllRoutes returns the names of all known actions and collections.
func (ps *Permission) AllRoutes() (collections []string, actions []string) {
	rr := make([]string, 0, len(ps.hs.collections))
	for k := range ps.hs.collections {
		rr = append(rr, k)
	}

	wr := make([]string, 0, len(ps.hs.actions))
	for k := range ps.hs.actions {
		wr = append(wr, k)
	}
	return rr, wr
}

// AdditionalUpdate TODO
func (ps *Permission) AdditionalUpdate(ctx context.Context, updated map[string]json.RawMessage) ([]int, error) {
	return nil, nil
}

// DataProvider is the connection to the datastore. It returns the data
// required by the permission service.
type DataProvider interface {
	// Get returns the values for the given fqfields.
	//
	// The number of return values have to be the same then the number of
	// requested fields. If a field does not exist, nil is returned for this
	// value.
	Get(ctx context.Context, fqfields ...string) ([]json.RawMessage, error)
}

// handlerStore saves the known actions and collections
type handlerStore struct {
	actions     map[string]perm.Action
	collections map[string]perm.Collection
}

func newHandlerStore() *handlerStore {
	return &handlerStore{
		actions:     make(map[string]perm.Action),
		collections: make(map[string]perm.Collection),
	}
}

func (hs *handlerStore) RegisterRestricter(name string, collection perm.Collection) {
	if _, ok := hs.collections[name]; ok {
		panic(fmt.Sprintf("Collection with name `%s` allready exists", name))
	}
	hs.collections[name] = collection
}

func (hs *handlerStore) RegisterAction(name string, action perm.Action) {
	if _, ok := hs.actions[name]; ok {
		panic(fmt.Sprintf("Action with name `%s` allready exists", name))
	}
	hs.actions[name] = action
}
