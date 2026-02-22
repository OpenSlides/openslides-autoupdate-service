package autoupdate

import (
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
	"github.com/OpenSlides/openslides-go/datastore"
	"github.com/OpenSlides/openslides-go/datastore/cache"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/environment"
)

// Flow is the connection to the database for the autoupdate service.
type Flow struct {
	flow.Flow

	cache    *cache.Cache
	postgres *datastore.FlowPostgres
}

// NewFlow initializes a flow for the autoupdate service.
func NewFlow(lookup environment.Environmenter) (*Flow, error) {
	postgres, err := datastore.NewFlowPostgres(lookup)
	if err != nil {
		return nil, fmt.Errorf("init postgres: %w", err)
	}

	cache := cache.New(postgres)

	flow := Flow{
		Flow:     cache,
		cache:    cache,
		postgres: postgres,
	}

	metric.Register(flow.metric)

	return &flow, nil
}

// ResetCache clears the cache.
func (f *Flow) ResetCache() {
	f.cache.Reset()
}

func (f *Flow) metric(values metric.Container) {
	values.Add("datastore_cache_key_len", f.cache.Len())
	values.Add("datastore_cache_size", f.cache.Size())
}
