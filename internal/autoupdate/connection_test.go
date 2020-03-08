package autoupdate_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestConnect(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	s := autoupdate.New(test.MockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}

	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Errorf("Next returned false, expected true")
	}
	if err := c.Err(); err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	key := "user/1/name"
	if value, ok := c.Data()[key]; !ok || value != "some value" {
		t.Errorf("Expected data to have key \"%s\" = \"%s\", got value \"%s\"", key, "some value", value)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	s := autoupdate.New(test.MockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Errorf("Next returned false, expected true")
	}
	if err := c.Err(); err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	go func() {
		// Simulate, that the client closes the connection
		cancel()
	}()

	if c.Next() {
		t.Errorf("Did not expect data")
	}
	if c.Err() != nil {
		t.Fatalf("Did not expect an error, got: %v", c.Err())
	}
	if len(c.Data()) != 0 {
		t.Errorf("Expect no new data, got: %v", c.Data())
	}
}

func TestConnectionReadNewData(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	restricter := &test.MockRestricter{}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Fatalf("Next returned false, expected true, err: %v", c.Err())
	}
	keychanges.Send(keys("user/1/name"))
	restricter.Data = keyValue{"user/1/name": "new value"}.m()

	if !c.Next() {
		t.Errorf("Next returned false, expected true")
	}
	if c.Err() != nil {
		t.Fatalf("Did not expect an error, got: %v", c.Err())
	}

	if data := c.Data(); len(data) != 1 || data["user/1/name"] != "new value" {
		t.Errorf("Expect data[\"user/1/name\"] to be \"new value\", got: \"%v\"", data["user/1/name"])
	}
}

func TestConnectionFilterData(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	restricter := &test.MockRestricter{Data: keyValue{"user/1/name": "name1"}.m()}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Fatalf("Next returned false, expected true, err: %v", c.Err())
	}
	keychanges.Send(keys("user/1/name")) // send again, value did not change in restricter

	if !c.Next() {
		t.Errorf("Next returned false, expected true")
	}
	if c.Err() != nil {
		t.Fatalf("Did not expect an error, got: %v", c.Err())
	}

	if data := c.Data(); len(data) != 0 {
		t.Errorf("Expect emty data; got \"%v\"", data)
	}
}

func TestConntectionOnlyDifferentData(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	restricter := &test.MockRestricter{Data: keyValue{"user/1/name": "name1", "user/2/name": "name2"}.m()}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name", "user/2/name")}
	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Fatalf("Next returned false, expected true, err: %v", c.Err())
	}
	restricter.Data["user/1/name"] = "newname" // Only change user/1 not user/2
	keychanges.Send(keys("user/1/name", "user/2/name"))

	if !c.Next() {
		t.Errorf("Next returned false, expected true")
	}
	if c.Err() != nil {
		t.Fatalf("Did not expect an error, got: %v", c.Err())
	}

	if data := c.Data(); len(data) != 1 || data["user/1/name"] != "newname" {
		t.Errorf("Expect data[\"user/1/name\"] to be \"newname\", got: \"%v\"", data)
	}
}

func BenchmarkFilterChanging(b *testing.B) {
	const keyCount = 100
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	restricter := &test.MockRestricter{Data: make(map[string]string)}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keys := make([]string, 0, keyCount)
	for i := 0; i < keyCount; i++ {
		keys = append(keys, fmt.Sprintf("user/%d/name", i))
	}
	kb := mockKeysBuilder{keys: keys}
	c := s.Connect(ctx, 1, kb)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Next()

		for i := 0; i < keyCount; i++ {
			restricter.Data[fmt.Sprintf("user/%d/name", i)] = fmt.Sprintf("value %d", n)
		}
		keychanges.Send(keys)
	}
}

func BenchmarkFilterNotChanging(b *testing.B) {
	const keyCount = 100
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	restricter := &test.MockRestricter{Data: make(map[string]string)}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keys := make([]string, 0, keyCount)
	for i := 0; i < keyCount; i++ {
		keys = append(keys, fmt.Sprintf("user/%d/name", i))
	}
	kb := mockKeysBuilder{keys: keys}
	c := s.Connect(ctx, 1, kb)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Next()
		keychanges.Send(keys)
	}
}
