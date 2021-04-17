package datastore_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataStoreGet(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, map[string]string{
		"collection/1/field": `"Hello World"`,
	})
	d := datastore.New(ts.TS.URL, closed, func(error) {}, ts)

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

	ts := test.NewDatastoreServer(closed, map[string]string{
		"collection/1/field": `"v1"`,
		"collection/2/field": `"v2"`,
	})
	d := datastore.New(ts.TS.URL, closed, func(error) {}, ts)

	got, err := d.Get(context.Background(), "collection/1/field", "collection/2/field")
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := test.Str(`"v1"`, `"v2"`)
	if len(got) != 2 || string(got[0]) != expect[0] || string(got[1]) != expect[1] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if ts.RequestCount != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1", ts.RequestCount)
	}
}

func TestCalculatedFields(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, nil)
	url := ts.TS.URL
	ds := datastore.New(url, closed, func(error) {}, ts)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string]json.RawMessage) ([]byte, error) {
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
	})

	t.Run("Fetch second time", func(t *testing.T) {
		ts.RequestCount = 0
		got, err := ds.Get(context.Background(), "collection/1/myfield")
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "my value", string(got[0]))
	})
}

func TestCalculatedFieldsNewDataInReceiver(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, map[string]string{
		"collection/1/normal_field": `"original value"`,
	})
	url := ts.TS.URL
	ds := datastore.New(url, closed, func(error) {}, ts)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string]json.RawMessage) ([]byte, error) {
		fields, err := ds.Get(context.Background(), "collection/1/normal_field")
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields[0])), nil
	})

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string]json.RawMessage) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	ts.Send(map[string]string{
		"collection/1/normal_field": `"new value"`,
	})
	<-done

	got, err := ds.Get(context.Background(), "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got[0]))
}

func TestCalculatedFieldsNewDataInReceiverAfterGet(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, map[string]string{
		"collection/1/normal_field": `"original value"`,
	})
	ds := datastore.New(ts.TS.URL, closed, func(error) {}, ts)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string]json.RawMessage) ([]byte, error) {
		fields, err := ds.Get(context.Background(), "collection/1/normal_field")
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields[0])), nil
	})

	// Call Get once to fill the cache
	ds.Get(context.Background(), "collection/1/myfield")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string]json.RawMessage) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	ts.Send(map[string]string{
		"collection/1/normal_field": `"new value"`,
	})
	<-done

	got, err := ds.Get(context.Background(), "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got[0]))
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTime(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, map[string]string{
		"collection/1/normal_field": `"original value"`,
	})
	ds := datastore.New(ts.TS.URL, closed, func(error) {}, ts)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string]json.RawMessage) ([]byte, error) {
		field, err := ds.Get(ctx, "collection/1/normal_field")
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return field[0], nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := ds.Get(ctx, "collection/1/normal_field", "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsNoDBQuery(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, nil)
	ds := datastore.New(ts.TS.URL, closed, func(error) {}, ts)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string]json.RawMessage) ([]byte, error) {
		return []byte("foobar"), nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := ds.Get(ctx, "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
	require.Equal(t, 0, ts.RequestCount)
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

func TestResetCache(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, nil)
	ds := datastore.New(ts.TS.URL, closed, func(error) {}, ts)

	// Fetch key to fill the cache.
	ds.Get(context.Background(), "some/1/key")
	ds.ResetCache()
	// Fetch key again.
	ds.Get(context.Background(), "some/1/key")

	// After a reset, the key should be fetched from the server again.
	assert.Equal(t, 2, ts.RequestCount)
}

func TestResetWhileUpdate(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ts := test.NewDatastoreServer(closed, nil)
	ds := datastore.New(ts.TS.URL, closed, func(error) {}, ts)

	// Fetch key to fill the cache.
	ds.Get(context.Background(), "some/1/key")

	doneReset := make(chan struct{})
	go func() {
		ds.ResetCache()
		close(doneReset)
	}()
	ts.Send(map[string]string{
		"some/1/key": "value",
	})

	<-doneReset
	// There is nothing to assert. This test is only for the race detector. Make
	// sure to run the tests with the -race flag.
}
