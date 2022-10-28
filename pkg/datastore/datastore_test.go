package datastore_test

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	myKey1       = datastore.MustKey("collection/1/field")
	myKey2       = datastore.MustKey("collection/2/field")
	myField1     = "collection/calculated"
	myCalculated = datastore.MustKey("collection/2/calculated")
)

func TestDataStoreGet(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{
		myKey1: []byte(`"Hello World"`),
	}))
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	got, err := ds.Get(context.Background(), myKey1)
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := []string{`"Hello World"`}
	if len(got) != 1 || string(got[myKey1]) != expect[0] {
		t.Errorf("Get() returned `%v`, expected `%v`", got, expect)
	}
}

func TestDataStoreGetMultiValue(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{
		myKey1: []byte(`"v1"`),
		myKey2: []byte(`"v2"`),
	}), dsmock.NewCounter)
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	got, err := ds.Get(context.Background(), myKey1, myKey2)
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := []string{`"v1"`, `"v2"`}
	if len(got) != 2 || string(got[myKey1]) != expect[0] || string(got[myKey2]) != expect[1] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1: %v", counter.Value(), counter.Requests())
	}
}

func TestDataStoreGetKeyTwice(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{
		myKey1: []byte(`"v1"`),
	}), dsmock.NewCounter)
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	got, err := ds.Get(context.Background(), myKey1, myKey1)
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := []string{`"v1"`, `"v1"`}
	if len(got) != 1 || string(got[myKey1]) != expect[0] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1: %v", counter.Value(), counter.Requests())
	}
}

func TestCalculatedFields(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}))
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		if changed == nil {
			return []byte("my value"), nil
		}

		return []byte(fmt.Sprintf("got %d changed keys", len(changed))), nil
	})

	t.Run("Fetch first time", func(t *testing.T) {
		got, err := ds.Get(context.Background(), myCalculated)
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "my value", string(got[myCalculated]))
	})

	t.Run("Fetch second time", func(t *testing.T) {
		got, err := ds.Get(context.Background(), myCalculated)
		require.NoError(t, err, "Get returned unexpected error")
		assert.Len(t, got, 1)
		assert.Equal(t, "my value", string(got[myCalculated]))
	})
}

func TestCalculatedFieldsNewDataInReceiver(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))

	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		fields, err := ds.Get(context.Background(), myKey1)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields[myKey1])), nil
	})

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[datastore.Key][]byte) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	source.Send(dsmock.YAMLData("collection/1/field: new value"))
	<-done

	got, err := ds.Get(context.Background(), myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got[myCalculated]))
}

func TestCalculatedFieldsNewDataInReceiverAfterGet(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))

	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		fields, err := ds.Get(context.Background(), myKey1)
		if err != nil {
			return nil, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields[myKey1])), nil
	})

	// Call Get once to fill the cache
	ds.Get(context.Background(), myKey1)

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[datastore.Key][]byte) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	source.Send(dsmock.YAMLData("collection/1/field: new value"))
	<-done

	got, err := ds.Get(context.Background(), myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got[myCalculated]))
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTime(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))

	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		field, err := ds.Get(ctx, myKey1)
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return field[myKey1], nil
	})

	ctx, cancel = context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myKey1)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeTwice(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		fields, err := ds.Get(ctx, myKey1, myKey1)
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return append(fields[myKey1], fields[myKey1]...), nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeAtDoesNotExist(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}))
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		field, err := ds.Get(ctx, myKey1)
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return field[myKey1], nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myKey1)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeAtDoesNotExistTwice(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}))
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		fields, err := ds.Get(ctx, myKey1, myKey1)
		if err != nil {
			return nil, fmt.Errorf("getting normal field: %w", err)
		}
		return append(fields[myKey1], fields[myKey1]...), nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsNoDBQuery(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}), dsmock.NewCounter)
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		return []byte("foobar"), nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myCalculated)
	require.NoError(t, err, "Get returned unexpected error")

	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 0 {
		t.Errorf("Got %d requests to the datastore, expected 0: %v", counter.Value(), counter.Requests())
	}
}

func TestChangeListeners(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}))
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	var receivedData map[datastore.Key][]byte
	received := make(chan struct{}, 1)

	ds.RegisterChangeListener(func(data map[datastore.Key][]byte) error {
		receivedData = data
		close(received)
		return nil
	})

	source.Send(dsmock.YAMLData("collection/1/field: my value"))

	<-received
	assert.Equal(t, []byte(`"my value"`), receivedData[myKey1])
}

func TestChangeListenersWithCalculatedFields(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}))
	ds, _, err := datastore.New(shutdownCtx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	var callCounter int
	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error) {
		callCounter++
		return []byte("foobar" + strconv.Itoa(callCounter)), nil
	})

	// Load calculated field in cache.
	ds.Get(context.Background(), myCalculated)

	var receivedData map[datastore.Key][]byte
	received := make(chan struct{}, 1)

	ds.RegisterChangeListener(func(data map[datastore.Key][]byte) error {
		receivedData = data
		close(received)
		return nil
	})

	source.Send(map[datastore.Key][]byte{myKey1: []byte(`"my value"`)})

	<-received
	assert.Equal(t, map[datastore.Key][]byte{
		myKey1:       []byte(`"my value"`),
		myCalculated: []byte("foobar2"),
	}, receivedData)
}

func TestResetCache(t *testing.T) {
	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}), dsmock.NewCounter)
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	// Fetch key to fill the cache.
	ds.Get(context.Background(), myKey1)
	ds.ResetCache()
	// Fetch key again.
	ds.Get(context.Background(), myKey1)

	// After a reset, the key should be fetched from the server again.
	if counter := source.Middlewares()[0].(*dsmock.Counter); counter.Value() != 2 {
		t.Errorf("Got %d requests to the datastore, expected 2: %v", counter.Value(), counter.Requests())
	}
}

func TestResetWhileUpdate(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := dsmock.NewStubWithUpdate(dsmock.Stub(map[datastore.Key][]byte{}))
	ctx := context.Background()
	ds, _, err := datastore.New(ctx, environment.ForTests{}, nil, datastore.WithDefaultSource(source))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go ds.ListenOnUpdates(shutdownCtx, func(err error) { log.Println(err) })

	// Fetch key to fill the cache.
	ds.Get(context.Background(), myKey1)

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
