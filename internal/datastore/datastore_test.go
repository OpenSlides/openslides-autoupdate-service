package datastore_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestDataStoreGet(t *testing.T) {
	ts := test.NewDatastoreServer()
	d := datastore.New(ts.TS.URL, new(test.UpdaterMock))

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
	ts := test.NewDatastoreServer()
	d := datastore.New(ts.TS.URL, new(test.UpdaterMock))

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
