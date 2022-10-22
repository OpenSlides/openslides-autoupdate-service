package datastore

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
)

func (d *Datastore) metric(values metric.Container) {
	values.Add("datastore_cache_key_len", d.cache.len())
	values.Add("datastore_cache_size", d.cache.size())
	values.Add("datastore_get_calls", int(d.metricGetHitCount))

	ds, ok := d.defaultSource.(*SourceDatastoreReader)
	if ok {
		values.Add("datastore_hits", int(ds.metricDSHitCount))
	}
}
