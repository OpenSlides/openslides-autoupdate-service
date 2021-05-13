package models

import (
	"context"
	"encoding/json"
)

//go:generate go run gen_models/main.go -- collections.go

// Getter gets keys from the datastore. It is the same as datastore.Getter.
type Getter interface {
	Get(ctx context.Context, keys ...string) ([]json.RawMessage, error)
}
