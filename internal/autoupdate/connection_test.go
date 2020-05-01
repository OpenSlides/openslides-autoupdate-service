package autoupdate_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestConnect(t *testing.T) {
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	datastore.Data = map[string]string{"user/1/name": `"some value"`}
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}

	c := s.Connect(ctx, 1, kb)
	data, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	if value, ok := data["user/1/name"]; !ok || value != `"some value"` {
		t.Errorf("c.Next() returned %v, expected %v", data, datastore.Data)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	datastore.Data = map[string]string{"user/1/name": `"some value"`}
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	cancel()
	data, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if data != nil {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestConnectionReadNewData(t *testing.T) {
	datastore := test.NewMockDatastore()
	defer datastore.Close()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	datastore.Update(map[string]string{"user/1/name": `"new value"`})
	datastore.Send(keys("user/1/name"))
	data, err := c.Next()

	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if value, ok := data["user/1/name"]; !ok || value != `"new value"` {
		t.Errorf("c.Next() returned %v, expected %v", data, map[string]string{"user/1/name": `"new value"`})
	}
}
