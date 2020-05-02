package datastore_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestDataStoreGet(t *testing.T) {
	ts := test.NewDatastoreServer()
	d := datastore.New(ts.TS.URL, new(test.MockKeysChanged))

	value, err := d.Get(context.Background(), "collection/1/field")

	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}

	expected := test.Str(`"Hello World"`)
	if !test.CmpSlice(value, expected) {
		t.Errorf("Get() returned `%v`, expected `%v`", value, expected)
	}
}

func TestDataStoreGetMultiValue(t *testing.T) {
	ts := test.NewDatastoreServer()
	d := datastore.New(ts.TS.URL, new(test.MockKeysChanged))

	got, err := d.Get(context.Background(), "collection/1/field", "collection/2/field")

	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}

	expected := test.Str(`"Hello World"`, `"Hello World"`)
	if !test.CmpSlice(got, expected) {
		t.Errorf("Get() returned %v, expected %v", got, expected)
	}

	if ts.RequestCount != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1", ts.RequestCount)
	}
}
