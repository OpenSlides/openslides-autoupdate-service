package keysbuilder

import (
	"context"
	"encoding/json"
)

// DataProvider decodes a restricted value for an key.
type DataProvider interface {
	RestrictedData(ctx context.Context, uid int, keys ...string) (map[string]json.RawMessage, error)
}

type fieldDescription interface {
	keys(key string, value json.RawMessage, data map[string]fieldDescription) error
}
