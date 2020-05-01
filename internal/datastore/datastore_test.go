package datastore_test

import (
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestDataStoreGet(t *testing.T) {
	ts := newTestServer()
	d := datastore.New(ts.ts.URL)

	value, err := d.Get("key")

	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}

	expected := `"value"`
	if len(value) != 1 || value[0] != expected {
		t.Errorf("Get() returned %v, expected [%s]", value, expected)
	}
}

func TestDataStoreGetMultiValue(t *testing.T) {
	ts := newTestServer()
	d := datastore.New(ts.ts.URL)

	got, err := d.Get("key1", "key2")

	if err != nil {
		t.Errorf("Get() returned an unexpected error: %v", err)
	}

	expected := test.Str(`"value"`, `"value"`)
	if !test.CmpSlice(got, expected) {
		t.Errorf("Get() returned %v, expected %v", got, expected)
	}

	if ts.requestCount != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1", ts.requestCount)
	}
}
