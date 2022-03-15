package datastore_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDataStoreGet(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{
		"collection/1/field": []byte(`"Hello World"`),
	}))
	ds := datastore.New(source, nil, source)

	got, err := ds.Get(context.Background(), "collection/1/field")
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := test.Str(`"Hello World"`)
	if len(got) != 1 || string(got["collection/1/field"]) != expect[0] {
		t.Errorf("Get() returned `%v`, expected `%v`", got, expect)
	}
}

func TestDataStoreGetMultiValue(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{
		"collection/1/field": []byte(`"v1"`),
		"collection/2/field": []byte(`"v2"`),
	}), dsmock.NewCounter)
	ds := datastore.New(source, nil, source)

	got, err := ds.Get(context.Background(), "collection/1/field", "collection/2/field")
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := test.Str(`"v1"`, `"v2"`)
	if len(got) != 2 || string(got["collection/1/field"]) != expect[0] || string(got["collection/2/field"]) != expect[1] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1: %v", counter.Value(), counter.Requests())
	}
}

func TestDataStoreGetKeyTwice(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{
		"collection/1/field": []byte(`"v1"`),
	}), dsmock.NewCounter)
	ds := datastore.New(source, nil, source)

	got, err := ds.Get(context.Background(), "collection/1/field", "collection/1/field")
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := test.Str(`"v1"`, `"v1"`)
	if len(got) != 1 || string(got["collection/1/field"]) != expect[0] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1: %v", counter.Value(), counter.Requests())
	}
}

func TestDataStoreGetInvalidKey(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}))
	ds := datastore.New(source, nil, source)

	_, err := ds.Get(context.Background(), "collection/1/Field")

	var errTyped interface {
		Type() string
	}
	if !errors.As(err, &errTyped) {
		t.Fatalf("Get() returned no error with Type method, got: %v", err)
	}

	if errTyped.Type() != "invalid" {
		t.Errorf("Error is of type %s, expected invalid", errTyped.Type())
	}
}

func TestCalculatedFields(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}))
	ds := datastore.New(source, nil, source)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		if changed == nil {
			return []byte("my value"), nil
		}

		return []byte(fmt.Sprintf("got %d changed keys", len(changed))), nil
	})

	t.Run("Fetch first time", func(t *testing.T) {
		got, err := ds.Get(context.Background(), "collection/1/myfield")
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "my value", string(got["collection/1/myfield"]))
	})

	t.Run("Fetch second time", func(t *testing.T) {
		got, err := ds.Get(context.Background(), "collection/1/myfield")
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "my value", string(got["collection/1/myfield"]))
	})
}

func TestCalculatedFieldsNewDataInReceiver(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{
		"collection/1/normal_field": []byte(`"original value"`),
	}))

	ds := datastore.New(source, nil, source)
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		fields, err := ds.Get(context.Background(), "collection/1/normal_field")
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields["collection/1/normal_field"])), nil
	})

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string][]byte) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	source.Send(dsmock.YAMLData("collection/1/normal_field: new value"))
	<-done

	got, err := ds.Get(context.Background(), "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got["collection/1/myfield"]))
}

func TestCalculatedFieldsNewDataInReceiverAfterGet(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{
		"collection/1/normal_field": []byte(`"original value"`),
	}))

	ds := datastore.New(source, nil, source)
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		fields, err := ds.Get(context.Background(), "collection/1/normal_field")
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields["collection/1/normal_field"])), nil
	})

	// Call Get once to fill the cache
	ds.Get(context.Background(), "collection/1/myfield")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string][]byte) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	source.Send(dsmock.YAMLData("collection/1/normal_field: new value"))
	<-done

	got, err := ds.Get(context.Background(), "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got["collection/1/myfield"]))
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTime(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{
		"collection/1/normal_field": []byte(`"original value"`),
	}))

	ds := datastore.New(source, nil, source)
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		field, err := ds.Get(ctx, "collection/1/normal_field")
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return field["collection/1/normal_field"], nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := ds.Get(ctx, "collection/1/normal_field", "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeTwice(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{
		"collection/1/normal_field": []byte(`"original value"`),
	}))
	ds := datastore.New(source, nil, source)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		fields, err := ds.Get(ctx, "collection/1/normal_field", "collection/1/normal_field")
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return append(fields["collection/1/normal_field"], fields["collection/1/normal_field"]...), nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := ds.Get(ctx, "collection/1/normal_field", "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeAtDoesNotExist(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}))
	ds := datastore.New(source, nil, source)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		field, err := ds.Get(ctx, "collection/1/normal_field")
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return field["collection/1/normal_field"], nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := ds.Get(ctx, "collection/1/normal_field", "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeAtDoesNotExistTwice(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}))
	ds := datastore.New(source, nil, source)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		fields, err := ds.Get(ctx, "collection/1/normal_field", "collection/1/normal_field")
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return append(fields["collection/1/normal_field"], fields["collection/1/normal_field"]...), nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := ds.Get(ctx, "collection/1/normal_field", "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsNoDBQuery(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}), dsmock.NewCounter)
	ds := datastore.New(source, nil, source)
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		return []byte("foobar"), nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := ds.Get(ctx, "collection/1/myfield")
	require.NoError(t, err, "Get returned unexpected error")

	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 0 {
		t.Errorf("Got %d requests to the datastore, expected 0: %v", counter.Value(), counter.Requests())
	}
}

func TestChangeListeners(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}))
	ds := datastore.New(source, nil, source)
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	var receivedData map[string][]byte
	received := make(chan struct{}, 1)

	ds.RegisterChangeListener(func(data map[string][]byte) error {
		receivedData = data
		close(received)
		return nil
	})

	source.Send(dsmock.YAMLData("my/1/key: my value"))

	<-received
	assert.Equal(t, []byte(`"my value"`), receivedData["my/1/key"])
}

func TestChangeListenersWithCalculatedFields(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}))
	ds := datastore.New(source, nil, source)
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	var callCounter int
	ds.RegisterCalculatedField("collection/myfield", func(ctx context.Context, key string, changed map[string][]byte) ([]byte, error) {
		callCounter++
		return []byte("foobar" + strconv.Itoa(callCounter)), nil
	})

	// Load calculated field in cache.
	ds.Get(context.Background(), "collection/1/myfield")

	var receivedData map[string][]byte
	received := make(chan struct{}, 1)

	ds.RegisterChangeListener(func(data map[string][]byte) error {
		receivedData = data
		close(received)
		return nil
	})

	source.Send(map[string][]byte{"my/1/key": []byte(`"my value"`)})

	<-received
	assert.Equal(t, map[string][]byte{
		"my/1/key":             []byte(`"my value"`),
		"collection/1/myfield": []byte("foobar2"),
	}, receivedData)
}

func TestResetCache(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}), dsmock.NewCounter)
	ds := datastore.New(source, nil, source)

	// Fetch key to fill the cache.
	ds.Get(context.Background(), "some/1/key")
	ds.ResetCache()
	// Fetch key again.
	ds.Get(context.Background(), "some/1/key")

	// After a reset, the key should be fetched from the server again.
	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 2 {
		t.Errorf("Got %d requests to the datastore, expected 2: %v", counter.Value(), counter.Requests())
	}
}

func TestResetWhileUpdate(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[string][]byte{}))
	ds := datastore.New(source, nil, source)
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	// Fetch key to fill the cache.
	ds.Get(context.Background(), "some/1/key")

	doneReset := make(chan struct{})
	go func() {
		ds.ResetCache()
		close(doneReset)
	}()
	source.Send(dsmock.YAMLData("some/1/key: value"))

	<-doneReset
	// There is nothing to assert. This test is only for the race detector. Make
	// sure to run the tests with the -race flag.
}
