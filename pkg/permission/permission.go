package permission

import (
	"context"
	"encoding/json"

	"github.com/OpenSlides/openslides-permission-service/internal/core"
)

// New returns a new permission service.
func New(dataprovider ExternalDataProvider) Permission {
	return core.NewPermissionService(dataprovider)
}

// ExternalDataProvider is the connection to the datastore. It returns the data
// required by the permission service.
type ExternalDataProvider interface {
	// If a field does not exist, it is not returned.
	Get(ctx context.Context, fqfields ...string) ([]json.RawMessage, error)
}

// Permission can tell, if a user has the permission for some data.
//
// See https://github.com/FinnStutzenstein/OpenSlides/blob/permissionService/docs/interfaces/permission-service.txt
type Permission interface {
	IsAllowed(name string, userID int, data map[string]string) (bool, map[string]interface{}, error)
	RestrictFQIDs(userID int, fqids []string) (map[string]bool, error)
	RestrictFQFields(userID int, fqfields []string) (map[string]bool, error)
	AdditionalUpdate(updated map[string]string) ([]int, error)
}
