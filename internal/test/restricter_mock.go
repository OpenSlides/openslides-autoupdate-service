package test

import (
	"context"
	"encoding/json"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// RestrictAllowed is a restricter that allows everything
func RestrictAllowed(ctx context.Context, fetch *datastore.Fetcher, uid int, data map[string]json.RawMessage) error {
	return nil
}

// RestrictNotAllowed is a restricter that removes everythin
func RestrictNotAllowed(ctx context.Context, fetch *datastore.Fetcher, uid int, data map[string]json.RawMessage) error {
	for k := range data {
		data[k] = nil
	}
	return nil
}
