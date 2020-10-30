package test

import (
	"context"
	"encoding/json"
	"time"
)

// DataProvider is a mock that impelemts the keysbuilder.DataProvider interface.
type DataProvider struct {
	Err          error
	Data         map[string]json.RawMessage
	Sleep        time.Duration
	RequestCount int
}

// RestrictedData returns the restricted Data as fiven in the mock.
func (r *DataProvider) RestrictedData(ctx context.Context, uid int, keys ...string) (map[string]json.RawMessage, error) {
	time.Sleep(r.Sleep)
	if r.Err != nil {
		return nil, r.Err
	}

	r.RequestCount++

	data := make(map[string]json.RawMessage, len(keys))
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
