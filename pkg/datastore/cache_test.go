package datastore

import (
	"context"
	"errors"
	"reflect"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

func TestCacheGetOrSet(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	c := newCache()
	got, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(keys []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("value")})
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}
	expect := []string{"value"}
	if len(got) != 1 || string(got[0]) != expect[0] {
		t.Errorf("GetOrSet() returned `%v`, expected `%v`", got, expect)
	}
}

func TestCacheGetOrSetMissingKeys(t *testing.T) {
	myKey1 := dskey.MustKey("key/1/field")
	myKey2 := dskey.MustKey("key/2/field")
	c := newCache()
	got, err := c.GetOrSet(context.Background(), []dskey.Key{myKey1, myKey2}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey1: []byte("value")})
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if len(got) != 2 {
		t.Errorf("got %d keys, expected 2", len(got))
	}

	if string(got[0]) != "value" {
		t.Errorf("%s has value %s, expected `value`", myKey1, got[0])
	}

	if got[1] != nil {
		t.Errorf("%s has value %s, expected nil", myKey2, got[1])
	}
}

func TestCacheGetOrSetNoSecondCall(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	c := newCache()
	c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("value")})
		return nil
	})

	var called bool

	got, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		called = true
		set(map[dskey.Key][]byte{myKey: []byte("value")})
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error %v", err)
	}

	if len(got) != 1 || string(got[0]) != "value" {
		t.Errorf("GetOrSet() returned %q, expected %q", got, "value")
	}
	if called {
		t.Errorf("GetOrSet() called the set method")
	}
}

func TestCacheGetOrSetBlockSecondCall(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	c := newCache()
	wait := make(chan struct{})
	go func() {
		c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
			<-wait
			set(map[dskey.Key][]byte{myKey: []byte("value")})
			return nil
		})
	}()

	// close done, when the second call is finished.
	done := make(chan struct{})
	go func() {
		c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
			set(map[dskey.Key][]byte{myKey: []byte("Shut not be returned")})
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

func TestCacheGetOrSetErrorInTheMiddle(t *testing.T) {
	myKey1 := dskey.MustKey("key/1/field")
	myKey2 := dskey.MustKey("key/2/field")
	c := newCache()
	_, err := c.GetOrSet(context.Background(), []dskey.Key{myKey1, myKey2}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey1: []byte("value")})
		return errors.New("some error")
	})
	if err == nil {
		t.Fatalf("got not error, expected some")
	}

	// Request key2 a second time, but this time outout an error
	got, err := c.GetOrSet(context.Background(), []dskey.Key{myKey2}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey2: []byte("expected Value")})
		return nil
	})

	expect := [][]byte{[]byte("expected Value")}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got value %v, expected %v", got, expect)
	}
}

func TestCacheSetIfExist(t *testing.T) {
	myKey1 := dskey.MustKey("key/1/field")
	myKey2 := dskey.MustKey("key/2/field")
	c := newCache()
	c.GetOrSet(context.Background(), []dskey.Key{myKey1}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey1: []byte("Shut not be returned")})
		return nil
	})

	// Set key1 and key2. key1 is in the cache. key2 should be ignored.
	c.SetIfExistMany(map[dskey.Key][]byte{
		myKey1: []byte("new_value"),
		myKey2: []byte("new_value"),
	})

	// Get key1 and key2 from the cache. The existing key1 should not be set.
	// key2 should be.
	got, _ := c.GetOrSet(context.Background(), []dskey.Key{myKey1, myKey2}, func(keys []dskey.Key, set func(map[dskey.Key][]byte)) error {
		for _, key := range keys {
			set(map[dskey.Key][]byte{key: []byte(key.String())})
		}
		return nil
	})

	expect := []string{"new_value", "key/2/field"}
	if len(got) != 2 || string(got[0]) != expect[0] || string(got[1]) != expect[1] {
		t.Errorf("Got %v, expected %v", got, expect)
	}
}

func TestCacheSetIfExistParallelToGetOrSet(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	c := newCache()

	waitForGetOrSet := make(chan struct{})
	go func() {
		c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
			// Signal, that GetOrSet was called.
			close(waitForGetOrSet)

			// Wait for some time.
			time.Sleep(10 * time.Millisecond)
			set(map[dskey.Key][]byte{myKey: []byte("shut not be used")})
			return nil
		})
	}()

	<-waitForGetOrSet

	// Set key1 to new value and stop the ongoing GetOrSet-Call
	c.SetIfExistMany(map[dskey.Key][]byte{myKey: []byte("new value")})

	got, _ := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("Expect values in cache")})
		return nil
	})

	expect := []string{"new value"}
	if len(got) != 1 || string(got[0]) != expect[0] {
		t.Errorf("Got `%s`, expected `%s`", got, expect)
	}
}

func TestGetWhileUpdate(t *testing.T) {
	const count = 100
	var wg sync.WaitGroup

	c := newCache()

	myKey1 := dskey.MustKey("key/1/field")
	myKey2 := dskey.MustKey("key/2/field")
	c.GetOrSet(context.Background(), []dskey.Key{myKey1, myKey2}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey1: []byte("Init Value")})
		set(map[dskey.Key][]byte{myKey2: []byte("Init Value")})
		return nil
	})

	// Fetch Keys many times
	got := make([][][]byte, count)
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			got[i], _ = c.GetOrSet(context.Background(), []dskey.Key{myKey1, myKey2}, func([]dskey.Key, func(map[dskey.Key][]byte)) error { return nil })
		}(i)
	}

	// Update Keys many times to the same value
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()
			c.SetIfExistMany(map[dskey.Key][]byte{
				myKey1: []byte(strconv.Itoa(i)),
				myKey2: []byte(strconv.Itoa(i)),
			})
		}(i)
	}

	wg.Wait()

	for i, g := range got {
		if string(g[0]) != string(g[1]) {
			t.Fatalf("GetOrSet returned invalid data got[%d] has myKey1: %s but myKey2: %s", i, g[0], g[1])
		}
	}

}

// TODO: Flaky test
func TestCacheGetOrSetOldData(t *testing.T) {
	// GetOrSet is called only with key1. It returns key1 and key2 on version1
	// but takes a long time. In the meantime there is an update via setIfExist
	// for key1 and key2 on version2. At the end, there should not be the old
	// version1 in the cache (version2 or 'does not exist' is ok).
	myKey1 := dskey.MustKey("key/1/field")
	myKey2 := dskey.MustKey("key/2/field")
	c := newCache()

	waitForGetOrSetStart := make(chan struct{})
	waitForGetOrSetEnd := make(chan struct{})
	waitForSetIfExist := make(chan struct{})

	go func() {
		c.GetOrSet(context.Background(), []dskey.Key{myKey1}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
			close(waitForGetOrSetStart)
			set(map[dskey.Key][]byte{myKey1: []byte("v1"), myKey2: []byte("v1")})
			<-waitForSetIfExist
			return nil
		})
		close(waitForGetOrSetEnd)
	}()

	<-waitForGetOrSetStart
	c.SetIfExistMany(map[dskey.Key][]byte{
		myKey1: []byte("v2"),
		myKey2: []byte("v2"),
	})
	close(waitForSetIfExist)

	<-waitForGetOrSetEnd
	data, err := c.GetOrSet(context.Background(), []dskey.Key{myKey1, myKey2}, func(keys []dskey.Key, set func(map[dskey.Key][]byte)) error {
		data := make(map[dskey.Key][]byte, len(keys))
		for _, key := range keys {
			data[key] = []byte("key not in cache")
		}
		set(data)
		return nil
	})
	if err != nil {
		t.Errorf("GetOrSet returned unexpected error: %v", err)
	}

	if string(data[0]) != "v2" {
		t.Errorf("value for key1 is %s, expected `v2`", data[0])
	}

	if string(data[1]) == "v1" {
		t.Errorf("value for key2 is `v1`, expected `v2` or `key not in cache`")
	}
}

func TestCacheErrorOnFetching(t *testing.T) {
	// Make sure, that if a GetOrSet call fails the requested keys are not left
	// in pending state.
	myKey := dskey.MustKey("key/1/field")
	c := newCache()
	rErr := errors.New("GetOrSet Error")
	_, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		return rErr
	})

	if !errors.Is(err, rErr) {
		t.Errorf("GetOrSet returned err `%v`, expected `%v`", err, rErr)
	}

	done := make(chan [][]byte)
	go func() {
		data, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
			set(map[dskey.Key][]byte{myKey: []byte("value")})
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
		if string(data[0]) != "value" {
			t.Errorf("Second GetOrSet-Call returned value %q, expected value", data[0])
		}
	case <-timer.C:
		t.Errorf("Second GetOrSet-Call was not done after one Millisecond")
	}
}

func TestCacheConcurency(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	const count = 100
	c := newCache()
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			v, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(keys []dskey.Key, set func(map[dskey.Key][]byte)) error {
				time.Sleep(time.Millisecond)
				for _, k := range keys {
					set(map[dskey.Key][]byte{k: []byte("value")})
				}
				return nil
			})
			if err != nil {
				t.Errorf("goroutine %d returned error: %v", i, err)
			}

			if string(v[0]) != "value" {
				t.Errorf("goroutine %d returned %q", i, v)
			}

		}(i)
	}

	wg.Wait()
}

func TestGetNull(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	c := newCache()
	got, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("null")})
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if k1 := got[0]; k1 != nil {
		t.Errorf("GetOrSet() returned %q for key1, expected nil", k1)
	}
}

func TestUpdateNull(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	c := newCache()
	c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("value")})
		return nil
	})

	c.SetIfExist(myKey, []byte("null"))

	got, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("value that should not be fetched")})
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if k1 := got[0]; k1 != nil {
		t.Errorf("GetOrSet() returned %q for key1, expected nil", k1)
	}
}

func TestUpdateManyNull(t *testing.T) {
	myKey := dskey.MustKey("key/1/field")
	c := newCache()
	c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("value")})
		return nil
	})

	c.SetIfExistMany(map[dskey.Key][]byte{myKey: []byte("null")})

	got, err := c.GetOrSet(context.Background(), []dskey.Key{myKey}, func(key []dskey.Key, set func(map[dskey.Key][]byte)) error {
		set(map[dskey.Key][]byte{myKey: []byte("value that should not be fetched")})
		return nil
	})

	if err != nil {
		t.Errorf("GetOrSet() returned the unexpected error: %v", err)
	}

	if k1 := got[0]; k1 != nil {
		t.Errorf("GetOrSet() returned %q for key1, expected nil", k1)
	}
}
