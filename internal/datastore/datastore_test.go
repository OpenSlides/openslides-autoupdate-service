package datastore_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/assert"
)

func TestDataStoreGet(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := test.NewMockDatastore(map[string]string{
		"collection/1/field": `"Hello World"`,
	})
	url := ds.StartServer(closed)
	d := datastore.New(url, closed, func(error) {}, test.NewUpdaterMock())

	got, err := d.Get(context.Background(), "collection/1/field")
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := test.Str(`"Hello World"`)
	if len(got) != 1 || string(got[0]) != expect[0] {
		t.Errorf("Get() returned `%v`, expected `%v`", got, expect)
	}
}

func TestDataStoreGetMultiValue(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := test.NewMockDatastore(map[string]string{
		"collection/1/field": `"Hello World"`,
		"collection/2/field": `"Hello World"`,
	})
	url := ds.StartServer(closed)
	d := datastore.New(url, closed, func(error) {}, test.NewUpdaterMock())

	got, err := d.Get(context.Background(), "collection/1/field", "collection/2/field")
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := test.Str(`"Hello World"`, `"Hello World"`)
	if len(got) != 2 || string(got[0]) != expect[0] || string(got[1]) != expect[1] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if ds.CountGetCalled != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1", ds.CountGetCalled)
	}
}

// func TestCalculdatedFields(t *testing.T) {
// 	closed := make(chan struct{})
// 	defer close(closed)
// 	ts := test.NewDatastoreServer()
// 	d := datastore.New(ts.TS.URL, closed, func(error) {}, test.NewUpdaterMock())
// 	_ = d
// }
