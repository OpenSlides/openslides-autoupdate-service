package datastore

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestCacheGetOrSet(t *testing.T) {
	c := newCache()
	got, err := c.getOrSet(context.Background(), test.Str("key1"), func([]string) (map[string]json.RawMessage, error) {
		return map[string]json.RawMessage{"key1": json.RawMessage("value")}, nil
	})

	if err != nil {
		t.Errorf("getOrSet() returned the unexpected error %v", err)
	}
	expect := test.Str("value")
	if len(got) != 1 || string(got[0]) != expect[0] {
		t.Errorf("getOrSet() returned %v, expected %v", got, expect)
	}
}

func TestCacheGetOrSetNoSecondCall(t *testing.T) {
	c := newCache()
	c.getOrSet(context.Background(), test.Str("key1"), func([]string) (map[string]json.RawMessage, error) {
		return map[string]json.RawMessage{"key1": json.RawMessage("value")}, nil
	})

	var called bool

	got, err := c.getOrSet(context.Background(), test.Str("key1"), func([]string) (map[string]json.RawMessage, error) {
		called = true
		return map[string]json.RawMessage{"key1": json.RawMessage("Shut not be returned")}, nil
	})

	if err != nil {
		t.Errorf("getOrSet() returned the unexpected error %v", err)
	}
	expect := test.Str("value")
	if len(got) != 1 || string(got[0]) != expect[0] {
		t.Errorf("getOrSet() returned %v, expected %v", got, expect)
	}
	if called {
		t.Errorf("getOrSet() called the set method")
	}
}

func TestCacheGetOrSetBlockSecondCall(t *testing.T) {
	c := newCache()
	wait := make(chan struct{})
	go func() {
		c.getOrSet(context.Background(), test.Str("key1"), func([]string) (map[string]json.RawMessage, error) {
			<-wait
			return map[string]json.RawMessage{"key1": json.RawMessage("value")}, nil
		})
	}()

	// close done, when the second call is finished.
	done := make(chan struct{})
	go func() {
		c.getOrSet(context.Background(), test.Str("key1"), func([]string) (map[string]json.RawMessage, error) {
			return map[string]json.RawMessage{"key1": json.RawMessage("Shut not be returned")}, nil
		})
		close(done)
	}()

	select {
	case <-done:
		t.Errorf("done channel already closed")
	default:
	}

	close(wait)

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case <-done:
	case <-timer.C:
		t.Errorf("Second getOrSet-Call was not done one Millisecond after the frist getOrSet-Call was called.")
	}
}

func TestCacheSetIfExist(t *testing.T) {
	c := newCache()
	c.getOrSet(context.Background(), test.Str("key1"), func([]string) (map[string]json.RawMessage, error) {
		return map[string]json.RawMessage{"key1": json.RawMessage("value")}, nil
	})

	// Set key1 and key2. key1 is in the cache. key2 should be ignored.
	c.setIfExist(map[string]json.RawMessage{
		"key1": json.RawMessage("new_value"),
		"key2": json.RawMessage("new_value"),
	})

	// Get key1 and key2 from the cache. The existing key1 should not be set.
	// key2 should be.
	got, _ := c.getOrSet(context.Background(), test.Str("key1", "key2"), func(keys []string) (map[string]json.RawMessage, error) {
		data := make(map[string]json.RawMessage)
		for _, key := range keys {
			data[key] = json.RawMessage(key)
		}
		return data, nil
	})

	expect := test.Str("new_value", "key2")
	if len(got) != 2 || string(got[0]) != expect[0] || string(got[1]) != expect[1] {
		t.Errorf("Got %v, expected %v", got, expect)
	}
}

func TestCacheSetIfExistParallelToGetOrSet(t *testing.T) {
	c := newCache()

	waitForGetOrSet := make(chan struct{})
	go func() {
		c.getOrSet(context.Background(), test.Str("key1"), func(keys []string) (map[string]json.RawMessage, error) {
			// Signal, that getOrSet was called.
			close(waitForGetOrSet)

			// Wait for some time.
			time.Sleep(10 * time.Millisecond)
			return map[string]json.RawMessage{"key1": json.RawMessage("shut not be used")}, nil
		})
	}()

	<-waitForGetOrSet

	// Set key1 to new value and stop the ongoing getOrSet-Call
	c.setIfExist(map[string]json.RawMessage{"key1": json.RawMessage("new value")})

	got, _ := c.getOrSet(context.Background(), test.Str("key1"), func([]string) (map[string]json.RawMessage, error) {
		return map[string]json.RawMessage{"key1": json.RawMessage("Expect values in cache")}, nil
	})

	expect := test.Str("new value")
	if len(got) != 1 || string(got[0]) != expect[0] {
		t.Errorf("Got `%v`, expected `%v`", got, expect)
	}
}
