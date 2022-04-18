package main

import (
	"runtime/metrics"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/metric"
)

func runtimeMetrics(con metric.Container) {
	const metricName = "/sched/goroutines:goroutines"

	sample := make([]metrics.Sample, 1)
	sample[0].Name = metricName
	metrics.Read(sample)

	if sample[0].Value.Kind() == metrics.KindBad {
		return
	}

	goroutines := sample[0].Value.Uint64()

	s := con.Sub("runtime")
	s.Add("goroutines", goroutines)
}
