package autoupdate_test

import (
	"context"
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
	kb := mockKeysBuilder{keys: []string{"user/1/name"}}

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
	kb := mockKeysBuilder{keys: []string{"user/1/name"}}
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
	kb := mockKeysBuilder{keys: []string{"user/1/name"}}
	c := s.Connect(ctx, 1, kb)
	if err := c.Err(); err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	keychanges.send([]string{"user/1/name"})
	restricter.Data = map[string]string{"user/1/name": "new value"}

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
