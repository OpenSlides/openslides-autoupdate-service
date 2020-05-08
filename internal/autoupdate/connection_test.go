package autoupdate_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestConnect(t *testing.T) {
	c, _, _, close := getConnection()
	defer close()

	data, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	if value, ok := data["user/1/name"]; !ok || string(value) != `"Hello World"` {
		t.Errorf("c.Next() returned %v, expected map[user/1/name:\"Hello World\"", data)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	c, _, disconnect, close := getConnection()
	defer close()

	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	disconnect()
	data, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if data != nil {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestConnectionReadNewData(t *testing.T) {
	c, datastore, _, close := getConnection()
	defer close()
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.Update(map[string]json.RawMessage{"user/1/name": []byte(`"new value"`)})
	datastore.Send(test.Str("user/1/name"))
	data, err := c.Next()

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

	datastore := test.NewMockDatastore()
	defer datastore.Close()

	datastore.Data = map[string]json.RawMessage{
		doesExistKey: []byte("exist"),
	}
	datastore.OnlyData = true

	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()

	kb := mockKeysBuilder{keys: test.Str(doesExistKey, doesNotExistKey)}

	t.Run("First responce", func(t *testing.T) {
		c := s.Connect(context.Background(), 1, kb)

		data, err := c.Next()

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
		update         map[string]json.RawMessage
		expectExist    bool
		expectNotExist bool
	}{
		{
			"not exist->not exist",
			map[string]json.RawMessage{doesNotExistKey: nil},
			false, // existing key gets filtered.
			false,
		},
		{
			"not exist->exist",
			map[string]json.RawMessage{doesNotExistKey: []byte("value")},
			false, // existing key gets filtered.
			true,
		},
		{
			"exist->not exist",
			map[string]json.RawMessage{doesExistKey: nil},
			true,
			false,
		},
		{
			"exist->exist",
			map[string]json.RawMessage{doesExistKey: []byte("new value")},
			true,
			false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			c := s.Connect(context.Background(), 1, kb)
			if _, err := c.Next(); err != nil {
				t.Errorf("c.Next() returned an error: %v", err)
			}

			datastore.Update(tt.update)
			datastore.Send([]string{doesExistKey, doesNotExistKey})
			data, err := c.Next()

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
		c := s.Connect(context.Background(), 1, kb)
		if _, err := c.Next(); err != nil {
			t.Errorf("c.Next() returned an error: %v", err)
		}

		// First time not exist
		datastore.Update(map[string]json.RawMessage{doesExistKey: nil})
		datastore.Send([]string{doesExistKey})
		c.Next()

		// Second time not exist
		datastore.Send([]string{doesExistKey})
		data, err := c.Next()

		if err != nil {
			t.Fatalf("c.Next() returned an error: %v", err)
		}
		if _, ok := data[doesExistKey]; ok {
			t.Errorf("key %s in second responce: true, expect: false", doesExistKey)
		}
	})
}

func TestConnectionFilterData(t *testing.T) {
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()
	kb := mockKeysBuilder{keys: test.Str("user/1/name")}
	c := s.Connect(context.Background(), 1, kb)
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.Send(test.Str("user/1/name")) // send again, value did not change in restricter
	data, err := c.Next()

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
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()
	kb := mockKeysBuilder{keys: test.Str("user/1/name")}
	c := s.Connect(context.Background(), 1, kb)
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.Update(map[string]json.RawMessage{"user/1/name": []byte(`"newname"`)}) // Only change user/1 not user/2
	datastore.Send(test.Str("user/1/name", "user/2/name"))
	data, err := c.Next()

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

func BenchmarkFilterChanging(b *testing.B) {
	const keyCount = 100
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()

	keys := make([]string, 0, keyCount)
	for i := 0; i < keyCount; i++ {
		keys = append(keys, fmt.Sprintf("user/%d/name", i))
	}
	kb := mockKeysBuilder{keys: keys}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := s.Connect(ctx, 1, kb)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Next()
		for i := 0; i < keyCount; i++ {
			datastore.Update(map[string]json.RawMessage{fmt.Sprintf("user/%d/name", i): []byte(fmt.Sprintf(`"value %d"`, n))})
		}
		datastore.Send(keys)
	}
}

func BenchmarkFilterNotChanging(b *testing.B) {
	const keyCount = 100
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()

	keys := make([]string, 0, keyCount)
	for i := 0; i < keyCount; i++ {
		keys = append(keys, fmt.Sprintf("user/%d/name", i))
	}
	kb := mockKeysBuilder{keys: keys}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := s.Connect(ctx, 1, kb)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Next()
		datastore.Send(keys)
	}
}
