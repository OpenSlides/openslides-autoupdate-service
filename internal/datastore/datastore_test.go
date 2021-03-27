package datastore_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestDataStoreGet(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, map[string]string{
		"collection/1/field": `"Hello World"`,
	})
	d := datastore.New(ts.TS.URL, closed, func(error) {}, ts)

	got, err := d.Get(context.Background(), "collection/1/field")

	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}

	expect := test.Str(`"Hello World"`)
	if len(got) != 1 || string(got[0]) != expect[0] {
		t.Errorf("Get() returned `%v`, expected `%v`", got, expect)
	}
}

func TestDataStoreGetMultiValue(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, map[string]string{
		"collection/1/field": `"Hello World"`,
		"collection/2/field": `"Hello World"`,
	})
	d := datastore.New(ts.TS.URL, closed, func(error) {}, ts)

	got, err := d.Get(context.Background(), "collection/1/field", "collection/2/field")

	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}

	expect := test.Str(`"Hello World"`, `"Hello World"`)
	if len(got) != 2 || string(got[0]) != expect[0] || string(got[1]) != expect[1] {
		t.Errorf("Get() returned %v, expected %v", got, expect)
	}

	if ts.RequestCount != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1", ts.RequestCount)
	}
}

func TestChangeListeners(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, nil)
	ds := datastore.New(ts.TS.URL, closed, func(error) {}, ts)

	var receivedData map[string]json.RawMessage
	received := make(chan struct{}, 1)

	ds.RegisterChangeListener(func(data map[string]json.RawMessage) error {
		receivedData = data
		close(received)
		return nil
	})

	ts.Send(map[string]string{"my/1/key": `"my value"`})

	<-received
	assert.Equal(t, map[string]json.RawMessage{"my/1/key": []byte(`"my value"`)}, receivedData)
}
