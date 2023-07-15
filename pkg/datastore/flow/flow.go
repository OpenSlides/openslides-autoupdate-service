package flow

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// Getter implements the Get function to fetch keys.
type Getter interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
}

// Updater is a blocking function. It expects a callback. The callback is
// called, when there is new data.
type Updater interface {
	Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error))
}

// Flow combines a Getter with an Updater.
//
// It represents data that can be fetched and gets updated.
type Flow interface {
	Getter
	Updater
}
