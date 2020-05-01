package datastore

import (
	"context"
	"testing"
	"time"
)

func TestCacheGetOrSet(t *testing.T) {
	c := newCache()
	v, err := c.getOrSet("key1", func(context.Context) (string, error) {
		return "value", nil
	})

	if err != nil {
		t.Errorf("getOrSet() returned the unexpected error %v", err)
	}
	if v != "value" {
		t.Errorf("getOrSet() returned %v, expected \"value\"", v)
	}
}

func TestCacheGetOrSetNoSecondCall(t *testing.T) {
	c := newCache()
	c.getOrSet("key1", func(context.Context) (string, error) {
		return "value", nil
	})

	var called bool

	v, err := c.getOrSet("key1", func(context.Context) (string, error) {
		called = true
		return "Shut not be returned", nil
	})

	if err != nil {
		t.Errorf("getOrSet() returned the unexpected error %v", err)
	}
	if v != "value" {
		t.Errorf("getOrSet() returned %v, expected \"value\"", v)
	}
	if called {
		t.Errorf("getOrSet() called the set method")
	}
}

func TestCacheGetOrSetBlockSecondCall(t *testing.T) {
	c := newCache()
	wait := make(chan struct{})
	go func() {
		c.getOrSet("key1", func(context.Context) (string, error) {
			<-wait
			return "value", nil
		})
	}()

	// close done, when the second call is finished.
	done := make(chan struct{})
	go func() {
		c.getOrSet("key1", func(context.Context) (string, error) {
			return "Shut not be returned", nil
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
	c.getOrSet("key1", func(context.Context) (string, error) {
		return "value", nil
	})

	c.setIfExist("key1", "new value")
	c.setIfExist("key2", "new value")

	v1, _ := c.getOrSet("key1", func(context.Context) (string, error) {
		return "Shut not be returned", nil
	})
	v2, _ := c.getOrSet("key2", func(context.Context) (string, error) {
		return "Shut be returned", nil
	})

	if v1 != "new value" {
		t.Errorf("key1 is %s, expected %s", v1, "new value")
	}
	if v2 != "Shut be returned" {
		t.Errorf("key1 is %s, expected %s", v1, "Shut be returned")
	}
}

func TestCacheSetIfExistCloseGetOrSetCalls(t *testing.T) {
	c := newCache()

	var v string
	getOrSetDone := make(chan struct{})
	waitForGetOrSet := make(chan struct{})
	go func() {
		v, _ = c.getOrSet("key1", func(ctx context.Context) (string, error) {
			// Signal, that getOrSet was called.
			close(waitForGetOrSet)

			// wait until context is done
			<-ctx.Done()
			return "shut not be used", nil
		})
		close(getOrSetDone)
	}()

	<-waitForGetOrSet

	// Set key1 to new value and stop the ongoing getOrSet-Call
	c.setIfExist("key1", "new value")

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case <-getOrSetDone:
	case <-timer.C:
		t.Errorf("Expected getOrSet() to return after call to setIfExist. Took more then one millisecond.")
	}

	if v != "new value" {
		t.Errorf("key1 is `%s`, expected `%s`", v, "new value")
	}
}
