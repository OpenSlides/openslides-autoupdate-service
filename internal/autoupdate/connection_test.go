package autoupdate_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
)

func TestConnect(t *testing.T) {
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	s := autoupdate.New(MockRestricter{}, keychanges)
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
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	s := autoupdate.New(MockRestricter{}, keychanges)
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

func TestConntectionReadNewData(t *testing.T) {
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	restricter := &MockRestricter{}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Fatalf("Next returned false, expected true, err: %v", c.Err())
	}
	keychanges.send(keys("user/1/name"))
	restricter.data = map[string]string{"user/1/name": "new value"}

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

func TestConntectionFilterData(t *testing.T) {
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	restricter := &MockRestricter{data: keyValue{"user/1/name": "name1"}.m()}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Fatalf("Next returned false, expected true, err: %v", c.Err())
	}
	keychanges.send(keys("user/1/name")) // send again, value did not change in restricter

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
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	restricter := &MockRestricter{data: keyValue{"user/1/name": "name1", "user/2/name": "name2"}.m()}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name", "user/2/name")}
	c := s.Connect(ctx, 1, kb)
	if !c.Next() {
		t.Fatalf("Next returned false, expected true, err: %v", c.Err())
	}
	restricter.data["user/1/name"] = "newname" // Only change user/1 not user/2
	keychanges.send(keys("user/1/name", "user/2/name"))

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
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	restricter := &MockRestricter{data: make(map[string]string)}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keys := make([]string, keyCount)
	for i := 0; i < keyCount; i++ {
		keys = append(keys, fmt.Sprintf("user/%d/name", i))
	}
	kb := mockKeysBuilder{keys: keys}
	c := s.Connect(ctx, 1, kb)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Next()

		for i := 0; i < keyCount; i++ {
			restricter.data[fmt.Sprintf("user/%d/name", i)] = fmt.Sprintf("value %d", n)
		}
		keychanges.send(keys)
	}
}

func BenchmarkFilterNotChanging(b *testing.B) {
	const keyCount = 100
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	restricter := &MockRestricter{data: make(map[string]string)}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	keys := make([]string, keyCount)
	for i := 0; i < keyCount; i++ {
		keys = append(keys, fmt.Sprintf("user/%d/name", i))
	}
	kb := mockKeysBuilder{keys: keys}
	c := s.Connect(ctx, 1, kb)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		c.Next()
		keychanges.send(keys)
	}
}
