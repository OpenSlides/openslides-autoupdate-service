// Package permission handels the permissions of a user.
//
// It provides the Permission object with methods to tell, if a user is allowed
// to use an action or can see some fqfields.
package permission

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
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

	orgaLevel, err := ps.dp.OrgaLevel(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("checking for superadmin: %w", err)
	}

	grouped, err := groupFQFields(fqfields)
	if err != nil {
		return nil, fmt.Errorf("grouping fqfields: %w", err)
	}

	for name, fqfields := range grouped {
		if orgaLevel == "superadmin" {
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
	collections map[string]perm.Collection
}

func newHandlerStore() *handlerStore {
	return &handlerStore{
		collections: make(map[string]perm.Collection),
	}
}

func (hs *handlerStore) RegisterRestricter(name string, collection perm.Collection) {
	if _, ok := hs.collections[name]; ok {
		panic(fmt.Sprintf("Collection with name `%s` allready exists", name))
	}
	hs.collections[name] = collection
}
