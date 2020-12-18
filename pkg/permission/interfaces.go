package permission

import (
	"context"
	"encoding/json"
)

// DataProvider is the connection to the datastore. It returns the data
// required by the permission service.
type DataProvider interface {
	// If a field does not exist, it is not returned.
	Get(ctx context.Context, fqfields ...string) ([]json.RawMessage, error)
}
