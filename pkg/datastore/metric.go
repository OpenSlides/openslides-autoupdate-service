package datastore

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
)

func (d *Datastore) metric(values metric.Container) {
	c := values.Sub("datastore")
	c.Add("cache_key_len", d.cache.len())
	c.Add("cache_size", d.cache.size())
	c.Add("get_calls", d.metricGetHitCount)

	ds, ok := d.defaultSource.(*SourceDatastore)
	if ok {
		c.Add("ds_hits", ds.metricDSHitCount)
	}
}
