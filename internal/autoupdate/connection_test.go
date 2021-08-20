package autoupdate_test

import (
	"context"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConnect(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	next, _ := getConnection(closed)

	data, err := next(context.Background())
	if err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	if value, ok := data["user/1/name"]; !ok || string(value) != `"Hello World"` {
		t.Errorf("next() returned %v, expected map[user/1/name:\"Hello World\"", data)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	next, _ := getConnection(closed)
	ctx, disconnect := context.WithCancel(context.Background())

	if _, err := next(ctx); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	disconnect()
	data, err := next(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("next() returned error %v, expected context.Canceled", err)
	}
	if data != nil {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestConnectionReadNewData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	next, datastore := getConnection(closed)

	if _, err := next(context.Background()); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	datastore.Send(map[string]string{"user/1/name": `"new value"`})
	data, err := next(context.Background())

	if err != nil {
		t.Errorf("next() returned an error: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if value, ok := data["user/1/name"]; !ok || string(value) != `"new value"` {
		t.Errorf("next() returned %v, expected %v", data, map[string]string{"user/1/name": `"new value"`})
	}
}

func TestConnectionEmptyData(t *testing.T) {
	const (
		doesNotExistKey = "doesnot/1/exist"
		doesExistKey    = "user/1/name"
	)

	closed := make(chan struct{})
	defer close(closed)

	datastore := dsmock.NewMockDatastore(closed, map[string]string{
		doesExistKey: `"Hello World"`,
	})
	s := autoupdate.New(datastore, test.RestrictAllowed, closed)
	kb := test.KeysBuilder{K: test.Str(doesExistKey, doesNotExistKey)}

	t.Run("First responce", func(t *testing.T) {
		next := s.Connect(1, kb)

		data, err := next(context.Background())
		require.NoError(t, err)
		assert.Contains(t, data, doesExistKey, "next() should return the existing key")
		assert.NotContains(t, data, doesNotExistKey, "next() should not return a non existing key")
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
			next := s.Connect(1, kb)
			if _, err := next(context.Background()); err != nil {
				t.Errorf("next() returned an error: %v", err)
			}

			datastore.Send(tt.update)

			var data map[string][]byte
			var err error
			isBlocking := blocking(func() {
				data, err = next(context.Background())
			})

			require.NoError(t, err)
			if tt.expectBlocking {
				assert.True(t, isBlocking, "Expect next() to block")
			} else {
				assert.False(t, isBlocking, "Expect next() not to block.")
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
		next := s.Connect(1, kb)
		if _, err := next(context.Background()); err != nil {
			t.Errorf("next() returned an error: %v", err)
		}

		// First time not exist
		datastore.Send(map[string]string{doesExistKey: ""})

		blocking(func() {
			next(context.Background())
		})

		// Second time not exist
		datastore.Send(map[string]string{doesExistKey: ""})

		var err error
		isBlocking := blocking(func() {
			_, err = next(context.Background())
		})

		require.NoError(t, err)
		assert.True(t, isBlocking, "second request should be blocking")
	})
}

func TestConnectionFilterData(t *testing.T) {
	t.Skipf("TODO")
	closed := make(chan struct{})
	defer close(closed)

	datastore := dsmock.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	s := autoupdate.New(datastore, test.RestrictAllowed, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	next := s.Connect(1, kb)
	if _, err := next(context.Background()); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	// send again, value did not change in restricter
	datastore.Send(map[string]string{"user/1/name": `"Hello World"`})
	data, err := next(context.Background())

	if err != nil {
		t.Errorf("next() returned an error: %v", err)
	}
	if got := len(data); got != 0 {
		t.Errorf("Got %d keys, expected none", got)
	}
	if _, ok := data["user/1/name"]; ok {
		t.Errorf("next() returned %v, expected empty map", data)
	}
}

func TestConntectionFilterOnlyOneKey(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	datastore := dsmock.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	s := autoupdate.New(datastore, test.RestrictAllowed, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	next := s.Connect(1, kb)
	if _, err := next(context.Background()); err != nil {
		t.Errorf("next() returned an error: %v", err)
	}

	datastore.Send(map[string]string{"user/1/name": `"newname"`})
	data, err := next(context.Background())

	if err != nil {
		t.Errorf("next() returned an error: %v", err)
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

func TestNextNoReturnWhenDataIsRestricted(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	datastore := dsmock.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})

	s := autoupdate.New(datastore, test.RestrictNotAllowed, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}

	next := s.Connect(1, kb)

	t.Run("first call", func(t *testing.T) {
		var data map[string][]byte
		var err error
		isBlocked := blocking(func() {
			data, err = next(context.Background())

		})
		require.NoError(t, err, "next() returnd an error")
		assert.Empty(t, data, "next() should return data on first call.")
		assert.False(t, isBlocked, "next() should not block on first call.")
	})

	t.Run("next call", func(t *testing.T) {
		var data map[string][]byte
		var err error
		isBlocked := blocking(func() {
			data, err = next(context.Background())

		})
		require.NoError(t, err, "next() returned an error")
		assert.Empty(t, data, "next() returned data")
		assert.True(t, isBlocked, "next() did not block")
	})

}
