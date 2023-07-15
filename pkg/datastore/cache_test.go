package datastore

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func TestCache_call_Get_returns_the_value_from_flow(t *testing.T) {
	ctx := context.Background()
	myKey := dskey.MustKey("key/1/field")
	flow := dsmock.NewFlow(dsmock.YAMLData(`---
	key/1/field: value
	`))

	c := newCache(flow)

	got, err := c.Get(ctx, myKey)
	if err != nil {
		t.Fatalf("cache.Get(): %v", err)
	}

	expect := map[dskey.Key][]byte{myKey: []byte(`"value"`)}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Get() == `%v`, expected `%v`", got, expect)
	}
}

func TestCache_Get_with_a_key_not_in_the_flow_returns_nil_as_value(t *testing.T) {
	ctx := context.Background()
	flow := dsmock.NewFlow(
		dsmock.YAMLData(``),
		dsmock.NewCounter,
	)
	counter := flow.Middlewares()[0].(*dsmock.Counter)
	myKey := dskey.MustKey("key/1/field")
	c := newCache(flow)

	if _, err := c.Get(ctx, myKey); err != nil {
		t.Errorf("cache.Get(): %v", err)
	}
	counter.Reset()

	got, err := c.Get(ctx, myKey)
	if err != nil {
		t.Errorf("cache.Get(): %v", err)
	}

	expect := map[dskey.Key][]byte{myKey: nil}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got %v, expected %v", got, expect)
	}

	if counter.Count() != 0 {
		t.Errorf("Got %d requests, expected 0", counter.Count())
	}
}

func TestCache_call_Get_two_times_only_calls_the_flow_one_time(t *testing.T) {
	ctx := context.Background()

	flow := dsmock.NewFlow(
		dsmock.Stub(dsmock.YAMLData(`---
		key/1/field: value
		`)),
		dsmock.NewCounter,
	)
	myKey := dskey.MustKey("key/1/field")
	c := newCache(flow)

	if _, err := c.Get(ctx, myKey); err != nil {
		t.Fatalf("cache.Get(): %v", err)
	}

	if _, err := c.Get(ctx, myKey); err != nil {
		t.Fatalf("cache.Get(): %v", err)
	}

	counter := flow.Middlewares()[0].(*dsmock.Counter)
	if counter.Count() != 1 {
		t.Errorf("Cache called flow %d times. Expected 1", counter.Count())
	}
}

func TestCache_calling_get_at_the_same_time_second_call_waits_until_first_is_finished(t *testing.T) {
	ctx := context.Background()
	wait := make(chan error)
	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		key/1/field: value
		`),
		dsmock.NewWait(wait),
	)
	myKey := dskey.MustKey("key/1/field")
	c := newCache(flow)

	err1 := make(chan error)
	go func() {
		if _, err := c.Get(ctx, myKey); err != nil {
			err1 <- fmt.Errorf("Get: %v", err)
			return
		}
		err1 <- nil
	}()

	err2 := make(chan error)
	go func() {
		if _, err := c.Get(ctx, myKey); err != nil {
			err2 <- fmt.Errorf("Get: %v", err)
			return
		}
		err2 <- nil
	}()

	select {
	case err := <-err1:
		t.Errorf("First get: %v", err)

	case err := <-err2:
		t.Errorf("Second get: %v", err)
	default:
	}

	close(wait)

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()

	select {
	case <-err2:
	case <-timer.C:
		t.Errorf("Second Get-Call was not done one Millisecond after the frist Get-Call returned.")
	}
}

func TestCache_Get_gets_an_error_from_flow_does_not_effect_a_second_call_to_Get(t *testing.T) {
	ctx := context.Background()
	waiter := make(chan error, 1)
	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		key/1/field: value
		`),
		dsmock.NewWait(waiter),
	)
	myKey1 := dskey.MustKey("key/1/field")
	c := newCache(flow)

	waiter <- fmt.Errorf("some error")
	if _, err := c.Get(ctx, myKey1); err == nil {
		t.Fatalf("Got no error, expected some")
	}

	waiter <- nil
	if _, err := c.Get(ctx, myKey1); err != nil {
		t.Fatalf("Second Get: %v", err)
	}
}

func TestCache_Update_values_not_in_the_cache_do_not_update_the_cache(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		key/1/field: value
		key/2/field: value
		`),
		dsmock.NewCounter,
	)
	myKey1 := dskey.MustKey("key/1/field")
	myKey2 := dskey.MustKey("key/2/field")
	c := newCache(flow)

	// Calls update in background.
	go c.Update(ctx, nil)

	// Puts myKey1 in the cache
	if _, err := c.Get(ctx, myKey1); err != nil {
		t.Fatalf("Get: %v", err)
	}

	// Set key1 and key2. key1 is in the cache. key2 should be ignored.
	flow.Send(map[dskey.Key][]byte{
		myKey1: []byte("new_value"),
		myKey2: []byte("new_value"),
	})

	if _, err := c.Get(ctx, myKey1, myKey2); err != nil {
		t.Fatalf("Second Get: %v", err)
	}

	counter := flow.Middlewares()[0].(*dsmock.Counter)

	expect := [][]dskey.Key{
		{myKey1},
		{myKey2},
	}
	if got := counter.Requests(); !reflect.DeepEqual(got, expect) {
		t.Errorf("Got %v, expect %v", got, expect)
	}
}

func TestCache_Get_a_value_when_in_parallel_it_is_updated(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	waiter := make(chan error, 1)

	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		key/1/field: value
		`),
		dsmock.NewWait(waiter),
	)
	myKey := dskey.MustKey("key/1/field")
	c := newCache(flow)

	go c.Update(ctx, nil)

	type dataErr struct {
		data map[dskey.Key][]byte
		err  error
	}

	done := make(chan dataErr)
	go func() {
		got, err := c.Get(ctx, myKey)
		if err != nil {
			done <- dataErr{err: fmt.Errorf("Get1: %v", err)}
			return
		}
		done <- dataErr{data: got}
	}()

	flow.Send(map[dskey.Key][]byte{myKey: []byte("new value")})

	close(waiter)

	got := <-done

	if err := got.err; err != nil {
		t.Fatalf("Got: %v", err)
	}

	expect := map[dskey.Key][]byte{myKey: []byte("new value")}
	if !reflect.DeepEqual(got.data, expect) {
		t.Errorf("Got %v, expected %v", got.data, expect)
	}
}

func TestCache_flow_returns_null_should_return_nil(t *testing.T) {
	ctx := context.Background()

	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		key/1/field: null
		`),
	)
	myKey := dskey.MustKey("key/1/field")
	c := newCache(flow)

	got, err := c.Get(ctx, myKey)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := map[dskey.Key][]byte{myKey: nil}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got %v, expected %v", got, expect)
	}
}

func TestCache_flow_update_pushed_null_is_saved_as_nil(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		key/1/field: value
		`),
		dsmock.NewCounter,
	)
	counter := flow.Middlewares()[0].(*dsmock.Counter)
	myKey := dskey.MustKey("key/1/field")
	c := newCache(flow)

	waitForUpdated := make(chan struct{})
	go c.Update(ctx, func(map[dskey.Key][]byte, error) {
		close(waitForUpdated)
	})

	// Put myKey in the cache
	c.Get(ctx, myKey)
	counter.Reset()

	flow.Send(map[dskey.Key][]byte{myKey: []byte("null")})
	<-waitForUpdated

	got, err := c.Get(ctx, myKey)
	if err != nil {
		t.Fatalf("Get(): %v", err)
	}

	expect := map[dskey.Key][]byte{myKey: nil}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got %v, expected %v", got, expect)
	}

	if counter.Count() != 0 {
		t.Errorf("Got %d requests to the flow, expected non", counter.Count())
	}
}
