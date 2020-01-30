package autoupdate_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
)

func TestConnect(t *testing.T) {
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	s := autoupdate.New(autoupdate.MockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	krs, err := keysrequest.ManyFromJSON(strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"name":null}}]`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	_, data, err := s.Connect(ctx, 1, krs)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	key := "user/1/name"
	if value, ok := data[key]; !ok || value != "some value" {
		t.Errorf("Expected data to have key \"%s\" = \"%s\", got value \"%s\"", key, "some value", value)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	s := autoupdate.New(autoupdate.MockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	krs, err := keysrequest.ManyFromJSON(strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"name":null}}]`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	c, _, err := s.Connect(ctx, 1, krs)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	go func() {
		// Simulate, that the client closes the connection
		cancel()
	}()

	data, err := c.Read()
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestConntectionReadNewData(t *testing.T) {
	keychanges := newMockKeyChanged()
	defer keychanges.close()
	restricter := &autoupdate.MockRestricter{}
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	krs, err := keysrequest.ManyFromJSON(strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"name":null}}]`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	c, _, err := s.Connect(ctx, 1, krs)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	keychanges.send([]string{"user/1/name"})
	restricter.Data = map[string]string{"user/1/name": "new value"}

	data, err := c.Read()
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	if len(data) != 1 || data["user/1/name"] != "new value" {
		t.Errorf("Expect data[\"user/1/name\"] to be \"new value\", got: \"%v\"", data["user/1/name"])
	}
}

type mockKeyChanged struct {
	c chan []string
	t *time.Ticker
}

func newMockKeyChanged() mockKeyChanged {
	m := mockKeyChanged{}
	m.c = make(chan []string, 1)
	m.t = time.NewTicker(time.Second)
	return m
}

func (m mockKeyChanged) KeysChanged() ([]string, error) {
	select {
	case v := <-m.c:
		return v, nil
	case <-m.t.C:
		return nil, nil
	}
}

func (m mockKeyChanged) send(keys []string) {
	m.c <- keys
}

func (m mockKeyChanged) close() {
	m.t.Stop()
}
