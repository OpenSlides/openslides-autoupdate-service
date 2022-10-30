package datastore

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
)

func (d *Datastore) metric(values metric.Container) {
	values.Add("datastore_get_calls", int(d.metricGetHitCount))

	if d.history != nil {
		ds, ok := d.history.(*sourceDatastore)
		if ok {
			values.Add("datastore_hits", int(ds.metricDSHitCount))
		}
	}
}
