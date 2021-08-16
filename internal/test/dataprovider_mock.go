package test

import (
	"context"
	"time"
)

// DataProvider is a mock that impelemts the keysbuilder.DataProvider interface.
type DataProvider struct {
	Err          error
	Data         map[string][]byte
	Sleep        time.Duration
	RequestCount int
}

// RestrictedData returns the restricted Data as fiven in the mock.
func (r *DataProvider) RestrictedData(ctx context.Context, uid int, keys ...string) (map[string][]byte, error) {
	time.Sleep(r.Sleep)
	if r.Err != nil {
		return nil, r.Err
	}

	r.RequestCount++

	data := make(map[string][]byte, len(keys))
	for _, key := range keys {
		v, ok := r.Data[key]
		if !ok {
			data[key] = nil
			continue
		}
		data[key] = v
	}

	return data, nil
}
