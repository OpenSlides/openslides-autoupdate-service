package datastore_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	myKey1       = dskey.MustKey("collection/1/field")
	myKey2       = dskey.MustKey("collection/2/field")
	myField1     = "collection/calculated"
	myCalculated = dskey.MustKey("collection/2/calculated")
)

func TestDataStoreGet(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{
		myKey1: []byte(`"Hello World"`),
	}))

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
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
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{
		myKey1: []byte(`"v1"`),
		myKey2: []byte(`"v2"`),
	}), dsmock.NewCounter)

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	got, err := ds.Get(context.Background(), myKey1, myKey2)
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := []string{`"v1"`, `"v2"`}
	if len(got) != 2 || string(got[myKey1]) != expect[0] || string(got[myKey2]) != expect[1] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if counter := flow.Middlewares()[0].(*dsmock.Counter); counter.Count() != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1: %v", counter.Count(), counter.Requests())
	}
}

func TestDataStoreGetKeyTwice(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{
		myKey1: []byte(`"v1"`),
	}), dsmock.NewCounter)

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	got, err := ds.Get(context.Background(), myKey1, myKey1)
	assert.NoError(t, err, "Get() returned an unexpected error")

	expect := []string{`"v1"`, `"v1"`}
	if len(got) != 1 || string(got[myKey1]) != expect[0] {
		t.Errorf("Get() returned %s, expected %s", got, expect)
	}

	if counter := flow.Middlewares()[0].(*dsmock.Counter); counter.Count() != 1 {
		t.Errorf("Got %d requests to the datastore, expected 1: %v", counter.Count(), counter.Requests())
	}
}

func TestCalculatedFields(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}))

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		if changed == nil {
			return []byte("my value"), true, nil
		}

		return []byte(fmt.Sprintf("got %d changed keys", len(changed))), true, nil
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

	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))

	ds, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go bg(shutdownCtx, oserror.Handle)

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		fields, err := ds.Get(context.Background(), myKey1)
		if err != nil {
			return nil, false, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields[myKey1])), true, nil
	})

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[dskey.Key][]byte) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	flow.Send(dsmock.YAMLData("collection/1/field: new value"))
	<-done

	got, err := ds.Get(context.Background(), myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got[myCalculated]))
}

func TestCalculatedFieldsNewDataInReceiverAfterGet(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))

	ds, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go bg(shutdownCtx, oserror.Handle)

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		fields, err := ds.Get(context.Background(), myKey1)
		if err != nil {
			return nil, false, err
		}
		return []byte(fmt.Sprintf(`"normal_field is %s"`, fields[myKey1])), true, nil
	})

	// Call Get once to fill the cache
	ds.Get(context.Background(), myKey1)

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[dskey.Key][]byte) error {
		// Signal, that the data is updated.
		close(done)
		return nil
	})

	flow.Send(dsmock.YAMLData("collection/1/field: new value"))
	<-done

	got, err := ds.Get(context.Background(), myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, "\"normal_field is \"new value\"\"", string(got[myCalculated]))
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTime(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))

	ds, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go bg(shutdownCtx, oserror.Handle)

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		field, err := ds.Get(ctx, myKey1)
		if err != nil {
			return nil, false, fmt.Errorf("getting normal field: %w", err)
		}
		return field[myKey1], true, nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myKey1)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeTwice(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{
		myKey1: []byte(`"original value"`),
	}))

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		fields, err := ds.Get(ctx, myKey1, myKey1)
		if err != nil {
			return nil, false, fmt.Errorf("getting normal field: %w", err)
		}
		return append(fields[myKey1], fields[myKey1]...), true, nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeAtDoesNotExist(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}))

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		field, err := ds.Get(ctx, myKey1)
		if err != nil {
			return nil, false, fmt.Errorf("getting normal field: %w", err)
		}
		return field[myKey1], true, nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myKey1)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsRequireNormalFieldFetchedAtTheSameTimeAtDoesNotExistTwice(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}))

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		fields, err := ds.Get(ctx, myKey1, myKey1)
		if err != nil {
			return nil, false, fmt.Errorf("getting normal field: %w", err)
		}
		return append(fields[myKey1], fields[myKey1]...), true, nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myKey1, myCalculated)
	require.NoError(t, err, "Get returned unexpected error")
}

func TestCalculatedFieldsNoDBQuery(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}), dsmock.NewCounter)

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		return []byte("foobar"), true, nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err = ds.Get(ctx, myCalculated)
	require.NoError(t, err, "Get returned unexpected error")

	if counter := flow.Middlewares()[0].(*dsmock.Counter); counter.Count() != 0 {
		t.Errorf("Got %d requests to the datastore, expected 0: %v", counter.Count(), counter.Requests())
	}
}

func TestChangeListeners(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}))

	ds, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go bg(shutdownCtx, oserror.Handle)

	var receivedData map[dskey.Key][]byte
	received := make(chan struct{}, 1)

	ds.RegisterChangeListener(func(data map[dskey.Key][]byte) error {
		receivedData = data
		close(received)
		return nil
	})

	flow.Send(dsmock.YAMLData("collection/1/field: my value"))

	<-received
	assert.Equal(t, []byte(`"my value"`), receivedData[myKey1])
}

func TestChangeListenersWithCalculatedFields(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}))
	ds, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go bg(shutdownCtx, oserror.Handle)

	var callCounter int
	ds.RegisterCalculatedField(myField1, func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error) {
		callCounter++
		return []byte("foobar" + strconv.Itoa(callCounter)), true, nil
	})

	// Load calculated field in cache.
	ds.Get(context.Background(), myCalculated)

	var receivedData map[dskey.Key][]byte
	received := make(chan struct{}, 1)

	ds.RegisterChangeListener(func(data map[dskey.Key][]byte) error {
		receivedData = data
		close(received)
		return nil
	})

	flow.Send(map[dskey.Key][]byte{myKey1: []byte(`"my value"`)})

	<-received
	assert.Equal(t, map[dskey.Key][]byte{
		myKey1:       []byte(`"my value"`),
		myCalculated: []byte("foobar2"),
	}, receivedData)
}

func TestResetCache(t *testing.T) {
	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}), dsmock.NewCounter)

	ds, _, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}

	// Fetch key to fill the cache.
	ds.Get(context.Background(), myKey1)
	ds.ResetCache()
	// Fetch key again.
	ds.Get(context.Background(), myKey1)

	// After a reset, the key should be fetched from the server again.
	if counter := flow.Middlewares()[0].(*dsmock.Counter); counter.Count() != 2 {
		t.Errorf("Got %d requests to the datastore, expected 2: %v", counter.Count(), counter.Requests())
	}
}

func TestResetWhileUpdate(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlowFromStub(dsmock.Stub(map[dskey.Key][]byte{}))

	ds, bg, err := datastore.New(environment.ForTests{}, nil, datastore.WithDefaultFlow(flow))
	if err != nil {
		t.Fatalf("init ds: %v", err)
	}
	go bg(shutdownCtx, oserror.Handle)

	// Fetch key to fill the cache.
	ds.Get(context.Background(), myKey1)

	doneReset := make(chan struct{})
	go func() {
		ds.ResetCache()
		close(doneReset)
	}()
	flow.Send(dsmock.YAMLData("some/1/key: value"))

	<-doneReset
	// There is nothing to assert. This test is only for the race detector. Make
	// sure to run the tests with the -race flag.
}
