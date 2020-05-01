package datastore

import (
	"context"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestCacheGetOrSet(t *testing.T) {
	c := newCache()
	v, err := c.getOrSet(context.Background(), test.Str("key1"), func([]string) ([]string, error) {
		return test.Str("value"), nil
	})

	if err != nil {
		t.Errorf("getOrSet() returned the unexpected error %v", err)
	}
	expect := test.Str("value")
	if !test.CmpSlice(v, expect) {
		t.Errorf("getOrSet() returned %v, expected %v", v, expect)
	}
}

func TestCacheGetOrSetNoSecondCall(t *testing.T) {
	c := newCache()
	c.getOrSet(context.Background(), test.Str("key1"), func([]string) ([]string, error) {
		return test.Str("value"), nil
	})

	var called bool

	v, err := c.getOrSet(context.Background(), test.Str("key1"), func([]string) ([]string, error) {
		called = true
		return test.Str("Shut not be returned"), nil
	})

	if err != nil {
		t.Errorf("getOrSet() returned the unexpected error %v", err)
	}
	expect := test.Str("value")
	if !test.CmpSlice(v, expect) {
		t.Errorf("getOrSet() returned %v, expected %v", v, expect)
	}
	if called {
		t.Errorf("getOrSet() called the set method")
	}
}

func TestCacheGetOrSetBlockSecondCall(t *testing.T) {
	c := newCache()
	wait := make(chan struct{})
	go func() {
		c.getOrSet(context.Background(), test.Str("key1"), func([]string) ([]string, error) {
			<-wait
			return test.Str("value"), nil
		})
	}()

	// close done, when the second call is finished.
	done := make(chan struct{})
	go func() {
		c.getOrSet(context.Background(), test.Str("key1"), func([]string) ([]string, error) {
			return test.Str("Shut not be returned"), nil
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
	c.getOrSet(context.Background(), test.Str("key1"), func([]string) ([]string, error) {
		return test.Str("value"), nil
	})

	// Set key1 and key2. key1 is in the cache. key2 should be ignored.
	c.setIfExist(map[string]string{
		"key1": "new value",
		"key2": "new value",
	})

	// Get key1 and key2 from the cache. The existing key1 should not be set.
	// key2 should be.
	got, _ := c.getOrSet(context.Background(), test.Str("key1", "key2"), func(keys []string) ([]string, error) {
		return keys, nil
	})

	expect := test.Str("new value", "key2")
	if !test.CmpSlice(got, expect) {
		t.Errorf("Got %v, expected %v", got, expect)
	}
}

func TestCacheSetIfExistParallelToGetOrSet(t *testing.T) {
	c := newCache()

	waitForGetOrSet := make(chan struct{})
	go func() {
		c.getOrSet(context.Background(), test.Str("key1"), func(keys []string) ([]string, error) {
			// Signal, that getOrSet was called.
			close(waitForGetOrSet)

			// Wait for some time.
			time.Sleep(10 * time.Millisecond)
			return test.Str("shut not be used"), nil
		})
	}()

	<-waitForGetOrSet

	// Set key1 to new value and stop the ongoing getOrSet-Call
	c.setIfExist(map[string]string{"key1": "new value"})

	got, _ := c.getOrSet(context.Background(), test.Str("key1"), func([]string) ([]string, error) {
		return test.Str("Expect values in cache"), nil
	})

	expect := test.Str("new value")
	if !test.CmpSlice(got, expect) {
		t.Errorf("Got `%v`, expected `%v`", got, expect)
	}
}
