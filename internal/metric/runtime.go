package metric

import "runtime/metrics"

// Runtime gathers metrics from the go runtime
func Runtime(con Container) {
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
