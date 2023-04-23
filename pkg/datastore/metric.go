package datastore

import (
	"sync/atomic"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
)

func (d *Datastore) metric(values metric.Container) {
	values.Add("datastore_cache_key_len", d.cache.len())
	values.Add("datastore_cache_size", d.cache.size())

	hitCount := atomic.LoadUint64(&d.metricGetHitCount)
	values.Add("datastore_get_calls", int(hitCount))

	if d.history != nil {
		ds, ok := d.history.(*sourceDatastore)
		if ok {
			values.Add("datastore_hits", int(ds.metricDSHitCount))
		}
	}
}
