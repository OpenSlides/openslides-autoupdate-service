package autoupdate_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestConnect(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	s := autoupdate.New(&test.MockRestricter{Data: map[string]string{"user/1/name": `"some value"`}}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}

	c := s.Connect(ctx, 1, kb)
	read, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	var data map[string]json.RawMessage
	decoder := json.NewDecoder(read)

	if err := decoder.Decode(&data); err != nil {
		t.Errorf("Can not decode connectoin stream: %v", err)
	}

	otherData, err := ioutil.ReadAll(decoder.Buffered())
	if err != nil {
		t.Errorf("Can not read buffer from decoder: %v", err)
	}
	if !(len(otherData) == 0 || (len(otherData) == 1 && otherData[0] == '\n')) {
		t.Errorf("Expected no more data, got: %v", otherData)
	}
}

func TestConnectionReadNoNewData(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	s := autoupdate.New(new(test.MockRestricter), keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	cancel()
	r, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}
	if r != nil {
		t.Errorf("Expect no new data, got: %v", r)
	}
}

func TestConnectionReadNewData(t *testing.T) {
	keychanges := test.NewMockKeysChanged()
	defer keychanges.Close()
	restricter := new(test.MockRestricter)
	s := autoupdate.New(restricter, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	kb := mockKeysBuilder{keys: keys("user/1/name")}
	c := s.Connect(ctx, 1, kb)
	if _, err := c.Next(); err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	restricter.Update(map[string]string{"user/1/name": `"new value"`})
	keychanges.Send(keys("user/1/name"))
	read, err := c.Next()
	if err != nil {
		t.Errorf("c.Next() returned an error: %v", err)
	}

	var data map[string]json.RawMessage
	if err := json.NewDecoder(read).Decode(&data); err != nil {
		t.Errorf("Can not decode connectoin stream: %v", err)
	}
	if got := len(data); got != 1 {
		t.Errorf("Expected data to have one key, got: %d", got)
	}
	if _, ok := data["user/1/name"]; !ok {
		t.Errorf("Returned value does not have key `user/1/name`")
	}
	if got := string(data["user/1/name"]); got != `"new value"` {
		t.Errorf("Expect value `new value` got: %s", got)
	}
}
