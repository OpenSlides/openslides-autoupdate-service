package autoupdate_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestConnect(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	c, _ := getConnection(closed)

	data, err := c.Next(context.Background())
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	if value, ok := data["user/1/name"]; !ok || string(value) != `"Hello World"` {
		t.Errorf("c.Next() returned %v, expected map[user/1/name:\"Hello World\"", data)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	c, _ := getConnection(closed)
	ctx, disconnect := context.WithCancel(context.Background())

	if _, err := c.Next(ctx); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	disconnect()
	data, err := c.Next(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("c.Next() returned error %v, expected context.Canceled", err)
	}
	if data != nil {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestConnectionReadNewData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	c, datastore := getConnection(closed)

	if _, err := c.Next(context.Background()); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.SendValues(map[string]string{"user/1/name": `"new value"`})
	data, err := c.Next(context.Background())

	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if value, ok := data["user/1/name"]; !ok || string(value) != `"new value"` {
		t.Errorf("c.Next() returned %v, expected %v", data, map[string]string{"user/1/name": `"new value"`})
	}
}

func TestConnectionEmptyData(t *testing.T) {
	const (
		doesNotExistKey = "doesnot/1/exist"
		doesExistKey    = "user/1/name"
	)

	closed := make(chan struct{})
	defer close(closed)

	datastore := test.NewMockDatastore(closed, map[string]string{
		doesExistKey: `"Hello World"`,
	})

	s := autoupdate.New(datastore, new(test.MockRestricter), test.UserUpdater{}, closed)

	kb := test.KeysBuilder{K: test.Str(doesExistKey, doesNotExistKey)}

	t.Run("First responce", func(t *testing.T) {
		c := s.Connect(1, kb)

		data, err := c.Next(context.Background())

		if err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}
		if _, ok := data[doesExistKey]; !ok {
			t.Errorf("key %s not in first responce", doesExistKey)
		}
		if _, ok := data[doesNotExistKey]; ok {
			t.Errorf("key %s is in first responce", doesNotExistKey)
		}
	})

	for _, tt := range []struct {
		name           string
		update         map[string]string
		expectExist    bool
		expectNotExist bool
	}{
		{
			"not exist->not exist",
			map[string]string{doesNotExistKey: ""},
			false, // existing key gets filtered.
			false,
		},
		{
			"not exist->exist",
			map[string]string{doesNotExistKey: "value"},
			false, // existing key gets filtered.
			true,
		},
		{
			"exist->not exist",
			map[string]string{doesExistKey: ""},
			true,
			false,
		},
		{
			"exist->exist",
			map[string]string{doesExistKey: "new value"},
			true,
			false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := s.Connect(1, kb)
			if _, err := c.Next(context.Background()); err != nil {
				t.Errorf("c.Next() returned an error: %v", err)
			}

			datastore.SendValues(tt.update)
			data, err := c.Next(context.Background())

			if err != nil {
				t.Fatalf("c.Next() returned an error: %v", err)
			}
			if _, ok := data[doesExistKey]; ok != tt.expectExist {
				t.Errorf("key %s in second responce: %t, expect: %t", doesExistKey, ok, tt.expectExist)
			}
			if _, ok := data[doesNotExistKey]; ok != tt.expectNotExist {
				t.Errorf("key %s in second responce: %t, expect: %t", doesNotExistKey, ok, tt.expectExist)
			}

		})
	}

	t.Run("exit->not exist-> not exist", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		// First time not exist
		datastore.SendValues(map[string]string{doesExistKey: ""})
		c.Next(context.Background())

		// Second time not exist
		datastore.SendValues(map[string]string{doesExistKey: ""})
		data, err := c.Next(context.Background())

		if err != nil {
			t.Fatalf("c.Next() returned an error: %v", err)
		}
		if _, ok := data[doesExistKey]; ok {
			t.Errorf("key %s in second responce: true, expect: false", doesExistKey)
		}
	})
}

func TestConnectionFilterData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	datastore := test.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	s := autoupdate.New(datastore, new(test.MockRestricter), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	c := s.Connect(1, kb)
	if _, err := c.Next(context.Background()); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	// send again, value did not change in restricter
	datastore.SendValues(map[string]string{"user/1/name": `"Hello World"`})
	data, err := c.Next(context.Background())

	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if got := len(data); got != 0 {
		t.Errorf("Got %d keys, expected none", got)
	}
	if _, ok := data["user/1/name"]; ok {
		t.Errorf("c.Next() returned %v, expected empty map", data)
	}
}

func TestConntectionFilterOnlyOneKey(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	datastore := test.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	s := autoupdate.New(datastore, new(test.MockRestricter), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	c := s.Connect(1, kb)
	if _, err := c.Next(context.Background()); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.SendValues(map[string]string{"user/1/name": `"newname"`})
	data, err := c.Next(context.Background())

	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if _, ok := data["user/1/name"]; !ok {
		t.Errorf("Returned value does not have key `user/1/name`")
	}
	if got := string(data["user/1/name"]); got != `"newname"` {
		t.Errorf("Expect value `newname` got: %s", got)
	}
}

func TestFullUpdate(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	datastore := test.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	userUpdater := new(test.UserUpdater)
	s := autoupdate.New(datastore, new(test.MockRestricter), userUpdater, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}

	t.Run("other user", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		// Send fulldata for other user.
		userUpdater.UserIDs = []int{2}
		datastore.SendValues(map[string]string{"some/5/data": "value"})

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var data map[string]json.RawMessage
		var err error
		isBlocking := blocking(func() {
			data, err = c.Next(ctx)
		})

		if !isBlocking {
			t.Fatalf("fulldataupdate did not block")
		}

		if err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}

		if len(data) != 0 {
			t.Errorf("Got %s, expected no key update", data)
		}
	})

	t.Run("same user", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		// Send fulldata for same user.
		userUpdater.UserIDs = []int{1}
		datastore.SendValues(map[string]string{"some/5/data": "value"})

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var data map[string]json.RawMessage
		var err error
		isBlocking := blocking(func() {
			data, err = c.Next(ctx)
		})

		if isBlocking {
			t.Fatalf("fulldataupdate did block")
		}

		if err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}

		if len(data) != 1 || string(data["user/1/name"]) != `"Hello World"` {
			t.Errorf("Got %s, expected [user/1/name: Hello World]", data)
		}
	})
}

func blocking(f func()) bool {
	done := make(chan struct{})
	go func() {
		f()
		close(done)
	}()

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case <-done:
		return false
	case <-timer.C:
		return true
	}
}
