package datastore_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestCalculdatedFields(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	dsmock := test.NewMockDatastore(nil)
	url := dsmock.StartServer(closed)
	updater := test.NewUpdaterMock()
	ds := datastore.New(url, closed, func(error) {}, updater)
	ds.RegisterCalculatedField("collection/myfield", func(key string, changed map[string]json.RawMessage) ([]byte, error) {
		if changed == nil {
			return []byte("my value"), nil
		}

		return []byte(fmt.Sprintf("got %d changed keys", len(changed))), nil
	})

	t.Run("Fetch first time", func(t *testing.T) {
		got, err := ds.Get(context.Background(), "collection/1/myfield")
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "my value", string(got[0]))
		assert.Equal(t, 1, dsmock.CountGetCalled)
	})

	t.Run("Fetch second time", func(t *testing.T) {
		dsmock.CountGetCalled = 0
		got, err := ds.Get(context.Background(), "collection/1/myfield")
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "my value", string(got[0]))
		assert.Equal(t, 0, dsmock.CountGetCalled)
	})

	t.Run("Update some key", func(t *testing.T) {
		dsmock.CountGetCalled = 0
		done := make(chan struct{})
		ds.RegisterChangeListener(func(map[string]json.RawMessage) error {
			// Signal, that the data is updated.
			close(done)
			return nil
		})

		updater.Send(map[string]string{
			"some/1/field": "some value",
		})
		<-done

		got, err := ds.Get(context.Background(), "collection/1/myfield")
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "got 1 changed keys", string(got[0]))
		assert.Equal(t, 0, dsmock.CountGetCalled)
	})
}
