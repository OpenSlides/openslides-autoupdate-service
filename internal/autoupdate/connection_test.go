package autoupdate_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	datastore.Send(map[string]string{"user/1/name": `"new value"`})
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
	s := autoupdate.New(datastore, test.RestrictAllowed(), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str(doesExistKey, doesNotExistKey)}

	t.Run("First responce", func(t *testing.T) {
		c := s.Connect(1, kb)

		data, err := c.Next(context.Background())
		require.NoError(t, err)
		assert.Contains(t, data, doesExistKey, "c.Next() should return the existing key")
		assert.NotContains(t, data, doesNotExistKey, "c.Next() should not return a non existing key")
	})

	for _, tt := range []struct {
		name           string
		update         map[string]string
		expectBlocking bool
		expectExist    bool
		expectNotExist bool
	}{
		{
			name:           "not exist->not exist",
			update:         map[string]string{doesNotExistKey: ""},
			expectBlocking: true,
		},
		{
			name:           "not exist->exist",
			update:         map[string]string{doesNotExistKey: "value"},
			expectExist:    false, // existing key gets filtered.
			expectNotExist: true,
		},
		{
			name:           "exist->not exist",
			update:         map[string]string{doesExistKey: ""},
			expectExist:    true,
			expectNotExist: false,
		},
		{
			name:           "exist->exist",
			update:         map[string]string{doesExistKey: "new value"},
			expectExist:    true,
			expectNotExist: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := s.Connect(1, kb)
			if _, err := c.Next(context.Background()); err != nil {
				t.Errorf("c.Next() returned an error: %v", err)
			}

			datastore.Send(tt.update)

			var data map[string]json.RawMessage
			var err error
			isBlocking := blocking(func() {
				data, err = c.Next(context.Background())
			})

			require.NoError(t, err)
			if tt.expectBlocking {
				assert.True(t, isBlocking, "Expect c.Next() to block")
			} else {
				assert.False(t, isBlocking, "Expect c.Next() not to block.")
			}

			if tt.expectExist {
				assert.Contains(t, data, doesExistKey)
			} else {
				assert.NotContains(t, data, doesExistKey)
			}

			if tt.expectNotExist {
				assert.Contains(t, data, doesNotExistKey)
			} else {
				assert.NotContains(t, data, doesNotExistKey)
			}
		})
	}

	t.Run("exit->not exist-> not exist", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		// First time not exist
		datastore.Send(map[string]string{doesExistKey: ""})

		blocking(func() {
			c.Next(context.Background())
		})

		// Second time not exist
		datastore.Send(map[string]string{doesExistKey: ""})

		var err error
		isBlocking := blocking(func() {
			_, err = c.Next(context.Background())
		})

		require.NoError(t, err)
		assert.True(t, isBlocking, "second request should be blocking")
	})
}

func TestConnectionFilterData(t *testing.T) {
	t.Skipf("TODO")
	closed := make(chan struct{})
	defer close(closed)

	datastore := test.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	s := autoupdate.New(datastore, test.RestrictAllowed(), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	c := s.Connect(1, kb)
	if _, err := c.Next(context.Background()); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	// send again, value did not change in restricter
	datastore.Send(map[string]string{"user/1/name": `"Hello World"`})
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

	s := autoupdate.New(datastore, test.RestrictAllowed(), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	c := s.Connect(1, kb)
	if _, err := c.Next(context.Background()); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.Send(map[string]string{"user/1/name": `"newname"`})
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
	restricter := test.RestrictAllowed()
	s := autoupdate.New(datastore, restricter, userUpdater, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}

	t.Run("other user", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		restricter.Values = map[string]string{
			"user/1/name": `"New Value"`,
		}
		defer func() {
			// Reset values at the end.
			restricter.Values = nil
		}()

		// Send fulldata for other user. (additional update is triggert by an
		// datastore-update, so we have to change some key.)
		userUpdater.UserIDs = []int{2}
		datastore.Send(map[string]string{"some/5/data": "value"})

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

	t.Run("same user restricter changed", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		restricter.Values = map[string]string{
			"user/1/name": `"New Value"`,
		}
		defer func() {
			// Reset values at the end.
			restricter.Values = nil
		}()
		// Send fulldata for same user.
		userUpdater.UserIDs = []int{1}
		datastore.Send(map[string]string{"some/5/data": "value"})

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

		assert.Equal(t, map[string]json.RawMessage{"user/1/name": []byte(`"New Value"`)}, data)
	})

	t.Run("same user restricter data did not changed", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		// Send fulldata for same user.
		userUpdater.UserIDs = []int{1}
		datastore.Send(map[string]string{"some/5/data": "value"})

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var data map[string]json.RawMessage
		var err error
		isBlocking := blocking(func() {
			data, err = c.Next(ctx)
		})
		require.NoError(t, err, "Next returnd an undexpected error")
		assert.True(t, isBlocking, "Next should block if there is no new data")
		assert.Empty(t, data, "Data should be empty if data did not change")
	})

	t.Run("every user gets an full update on uid -1", func(t *testing.T) {
		c := s.Connect(1, kb)
		if _, err := c.Next(context.Background()); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		restricter.Values = map[string]string{
			"user/1/name": `"New Value"`,
		}
		defer func() {
			// Reset values at the end.
			restricter.Values = nil
		}()
		// Send fulldata for same user.
		userUpdater.UserIDs = []int{-1}
		datastore.Send(map[string]string{"some/5/data": "value"})

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

		assert.Equal(t, map[string]json.RawMessage{"user/1/name": []byte(`"New Value"`)}, data)
	})
}

func TestNextNoReturnWhenDataIsRestricted(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	datastore := test.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	userUpdater := new(test.UserUpdater)
	s := autoupdate.New(datastore, test.RestrictDenied(), userUpdater, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}

	c := s.Connect(1, kb)

	t.Run("first call", func(t *testing.T) {
		var data map[string]json.RawMessage
		var err error
		isBlocked := blocking(func() {
			data, err = c.Next(context.Background())

		})
		require.NoError(t, err, "c.Next() returnd an error")
		assert.Empty(t, data, "c.Next() should return data on first call.")
		assert.False(t, isBlocked, "c.Next() should not block on first call.")
	})

	t.Run("next call", func(t *testing.T) {
		var data map[string]json.RawMessage
		var err error
		isBlocked := blocking(func() {
			data, err = c.Next(context.Background())

		})
		require.NoError(t, err, "c.Next() returned an error")
		assert.Empty(t, data, "c.Next() returned data")
		assert.True(t, isBlocked, "c.Next() did not block")
	})

}
