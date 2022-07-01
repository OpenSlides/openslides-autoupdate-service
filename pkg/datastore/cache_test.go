package datastore

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"testing"
	"time"
)

func MustKey(in string) Key {
	k, err := KeyFromString(in)
	if err != nil {
		panic(err)
	}
	return k
}

func TestCacheGetOrSet(t *testing.T) {
	myKey := MustKey("key/1/field")
	c := newCache()
	got, err := c.GetOrSet(context.Background(), []Key{myKey}, func(keys []Key, set func(Key, []byte)) error {
		set(myKey, []byte("value"))
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}
	expect := []string{"value"}
	if len(got) != 1 || string(got[myKey]) != expect[0] {
		t.Errorf("GetOrSet() returned `%v`, expected `%v`", got, expect)
	}
}

func TestCacheGetOrSetMissingKeys(t *testing.T) {
	myKey1 := MustKey("key/1/field")
	myKey2 := MustKey("key/2/field")
	c := newCache()
	got, err := c.GetOrSet(context.Background(), []Key{myKey1, myKey2}, func(key []Key, set func(Key, []byte)) error {
		set(myKey1, []byte("value"))
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if len(got) != 2 {
		t.Errorf("got %d keys, expected 2", len(got))
	}

	if string(got[myKey1]) != "value" {
		t.Errorf("%s has value %s, expected `value`", myKey1, got[myKey1])
	}

	if got[myKey2] != nil {
		t.Errorf("%s has value %s, expected nil", myKey2, got[myKey2])
	}
}

func TestCacheGetOrSetNoSecondCall(t *testing.T) {
	myKey := MustKey("key/1/field")
	c := newCache()
	c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		set(myKey, []byte("value"))
		return nil
	})

	var called bool

	got, err := c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		called = true
		set(myKey, []byte("value"))
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error %v", err)
	}

	if len(got) != 1 || string(got[myKey]) != "value" {
		t.Errorf("GetOrSet() returned %q, expected %q", got, "value")
	}
	if called {
		t.Errorf("GetOrSet() called the set method")
	}
}

func TestCacheGetOrSetBlockSecondCall(t *testing.T) {
	myKey := MustKey("key/1/field")
	c := newCache()
	wait := make(chan struct{})
	go func() {
		c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
			<-wait
			set(myKey, []byte("value"))
			return nil
		})
	}()

	// close done, when the second call is finished.
	done := make(chan struct{})
	go func() {
		c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
			set(myKey, []byte("Shut not be returned"))
			return nil
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
		t.Errorf("Second GetOrSet-Call was not done one Millisecond after the frist GetOrSet-Call was called.")
	}
}

func TestCacheSetIfExist(t *testing.T) {
	myKey1 := MustKey("key/1/field")
	myKey2 := MustKey("key/2/field")
	c := newCache()
	c.GetOrSet(context.Background(), []Key{myKey1}, func(key []Key, set func(Key, []byte)) error {
		set(myKey1, []byte("Shut not be returned"))
		return nil
	})

	// Set key1 and key2. key1 is in the cache. key2 should be ignored.
	c.SetIfExistMany(map[Key][]byte{
		myKey1: []byte("new_value"),
		myKey2: []byte("new_value"),
	})

	// Get key1 and key2 from the cache. The existing key1 should not be set.
	// key2 should be.
	got, _ := c.GetOrSet(context.Background(), []Key{myKey1, myKey2}, func(keys []Key, set func(Key, []byte)) error {
		for _, key := range keys {
			set(key, []byte(key.String()))
		}
		return nil
	})

	expect := []string{"new_value", "key/2/field"}
	if len(got) != 2 || string(got[myKey1]) != expect[0] || string(got[myKey2]) != expect[1] {
		t.Errorf("Got %v, expected %v", got, expect)
	}
}

func TestCacheSetIfExistParallelToGetOrSet(t *testing.T) {
	myKey := MustKey("key/1/field")
	c := newCache()

	waitForGetOrSet := make(chan struct{})
	go func() {
		c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
			// Signal, that GetOrSet was called.
			close(waitForGetOrSet)

			// Wait for some time.
			time.Sleep(10 * time.Millisecond)
			set(myKey, []byte("shut not be used"))
			return nil
		})
	}()

	<-waitForGetOrSet

	// Set key1 to new value and stop the ongoing GetOrSet-Call
	c.SetIfExistMany(map[Key][]byte{myKey: []byte("new value")})

	got, _ := c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		set(myKey, []byte("Expect values in cache"))
		return nil
	})

	expect := []string{"new value"}
	if len(got) != 1 || string(got[myKey]) != expect[0] {
		t.Errorf("Got `%s`, expected `%s`", got, expect)
	}
}

func TestGetWhileUpdate(t *testing.T) {
	const count = 100
	var wg sync.WaitGroup

	c := newCache()

	myKey1 := MustKey("key/1/field")
	myKey2 := MustKey("key/2/field")
	c.GetOrSet(context.Background(), []Key{myKey1, myKey2}, func(key []Key, set func(Key, []byte)) error {
		set(myKey1, []byte("Init Value"))
		set(myKey2, []byte("Init Value"))
		return nil
	})

	// Fetch Keys many times
	got := make([]map[Key][]byte, count)
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			got[i], _ = c.GetOrSet(context.Background(), []Key{myKey1, myKey2}, func([]Key, func(Key, []byte)) error { return nil })
		}(i)
	}

	// Update Keys many times to the same value
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			c.SetIfExistMany(map[Key][]byte{
				myKey1: []byte(strconv.Itoa(i)),
				myKey2: []byte(strconv.Itoa(i)),
			})
		}(i)
	}

	wg.Wait()

	for i, g := range got {
		if string(g[myKey1]) != string(g[myKey2]) {
			t.Fatalf("GetOrSet returned invalid data got[%d] has myKey1: %s but myKey2: %s", i, g[myKey1], g[myKey2])
		}
	}

}

func TestCacheGetOrSetOldData(t *testing.T) {
	// GetOrSet is called with key1. It returns key1 and key2 on version1 but
	// takes a long time. In the meantime there is an update via setIfExist for
	// key1 and key2 on version2. At the end, there should not be the old
	// version1 in the cache (version2 or 'does not exist' is ok).
	myKey1 := MustKey("key/1/field")
	myKey2 := MustKey("key/2/field")
	c := newCache()

	waitForGetOrSetStart := make(chan struct{})
	waitForGetOrSetEnd := make(chan struct{})
	waitForSetIfExist := make(chan struct{})

	go func() {
		c.GetOrSet(context.Background(), []Key{myKey1}, func(key []Key, set func(Key, []byte)) error {
			close(waitForGetOrSetStart)
			set(myKey1, []byte("v1"))
			set(myKey2, []byte("v1"))
			<-waitForSetIfExist
			return nil
		})
		close(waitForGetOrSetEnd)
	}()

	<-waitForGetOrSetStart
	c.SetIfExistMany(map[Key][]byte{
		myKey1: []byte("v2"),
		myKey2: []byte("v2"),
	})
	close(waitForSetIfExist)

	<-waitForGetOrSetEnd
	data, err := c.GetOrSet(context.Background(), []Key{myKey1, myKey2}, func(keys []Key, set func(Key, []byte)) error {
		for _, key := range keys {
			set(key, []byte("key not in cache"))
		}
		return nil
	})
	if err != nil {
		t.Errorf("GetOrSet returned unexpected error: %v", err)
	}

	if string(data[myKey1]) != "v2" {
		t.Errorf("value for key1 is %s, expected `v2`", data[myKey1])
	}

	if string(data[myKey2]) == "v1" {
		t.Errorf("value for key2 is `v1`, expected `v2` or `key not in cache`")
	}
}

func TestCacheErrorOnFetching(t *testing.T) {
	// Make sure, that if a GetOrSet call fails the requested keys are not left
	// in pending state.
	myKey := MustKey("key/1/field")
	c := newCache()
	rErr := errors.New("GetOrSet Error")
	_, err := c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		return rErr
	})

	if !errors.Is(err, rErr) {
		t.Errorf("GetOrSet returned err `%v`, expected `%v`", err, rErr)
	}

	done := make(chan map[Key][]byte)
	go func() {
		data, err := c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
			set(myKey, []byte("value"))
			return nil
		})
		if err != nil {
			t.Errorf("Second GetOrSet returned unexpected err: %v", err)
		}
		done <- data
	}()

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case data := <-done:
		if string(data[myKey]) != "value" {
			t.Errorf("Second GetOrSet-Call returned value %q, expected value", data[myKey])
		}
	case <-timer.C:
		t.Errorf("Second GetOrSet-Call was not done after one Millisecond")
	}
}

func TestCacheConcurency(t *testing.T) {
	myKey := MustKey("key/1/field")
	const count = 100
	c := newCache()
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			v, err := c.GetOrSet(context.Background(), []Key{myKey}, func(keys []Key, set func(k Key, v []byte)) error {
				time.Sleep(time.Millisecond)
				for _, k := range keys {
					set(k, []byte("value"))
				}
				return nil
			})
			if err != nil {
				t.Errorf("goroutine %d returned error: %v", i, err)
			}

			if string(v[myKey]) != "value" {
				t.Errorf("goroutine %d returned %q", i, v)
			}

		}(i)
	}

	wg.Wait()
}

func TestGetNull(t *testing.T) {
	myKey := MustKey("key/1/field")
	c := newCache()
	got, err := c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		set(myKey, []byte("null"))
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if k1, ok := got[myKey]; k1 != nil || !ok {
		t.Errorf("GetOrSet() returned (%q, %t) for key1, expected (nil, true)", k1, ok)
	}
}

func TestUpdateNull(t *testing.T) {
	myKey := MustKey("key/1/field")
	c := newCache()
	c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		set(myKey, []byte("value"))
		return nil
	})

	c.SetIfExist(myKey, []byte("null"))

	got, err := c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		set(myKey, []byte("value that should not be fetched"))
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if k1, ok := got[myKey]; k1 != nil || !ok {
		t.Errorf("GetOrSet() returned (%q, %t) for key1, expected (nil, true)", k1, ok)
	}
}

func TestUpdateManyNull(t *testing.T) {
	myKey := MustKey("key/1/field")
	c := newCache()
	c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		set(myKey, []byte("value"))
		return nil
	})

	c.SetIfExistMany(map[Key][]byte{myKey: []byte("null")})

	got, err := c.GetOrSet(context.Background(), []Key{myKey}, func(key []Key, set func(Key, []byte)) error {
		set(myKey, []byte("value that should not be fetched"))
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if k1, ok := got[myKey]; k1 != nil || !ok {
		t.Errorf("GetOrSet() returned (%q, %t) for key1, expected (nil, true)", k1, ok)
	}
}
