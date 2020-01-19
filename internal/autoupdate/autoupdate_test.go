package autoupdate

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
)

func TestPrepare(t *testing.T) {
	keychanges := newMockKeyChanged()
	s := New(mockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	krs, err := keysrequest.ManyFromJSON(strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"name":null}}]`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	_, _, data, err := s.Prepare(ctx, 1, krs)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	key := "user/1/name"
	if value, ok := data[key]; !ok || string(value) != "some value" {
		t.Errorf("Expected data to have key \"%s\" = \"%s\", got value \"%s\"", key, "some value", value)
	}
}

func TestEchoNoNewData(t *testing.T) {
	keychanges := newMockKeyChanged()
	s := New(mockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	krs, err := keysrequest.ManyFromJSON(strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"name":null}}]`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	tid, keys, _, err := s.Prepare(ctx, 1, krs)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	go func() {
		// Simulate, that the client closes the connection
		cancel()
	}()

	ntid, data, err := s.Echo(ctx, 1, tid, keys)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	if tid != ntid {
		t.Errorf("Expect no new tid, got: %d", ntid)
	}

	if len(data) != 0 {
		t.Errorf("Expect no new data, got: %v", data)
	}
}

func TestEchoNewData(t *testing.T) {
	keychanges := newMockKeyChanged()
	s := New(mockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	krs, err := keysrequest.ManyFromJSON(strings.NewReader(`[{"ids":[1],"collection":"user","fields":{"name":null}}]`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	tid, keys, _, err := s.Prepare(ctx, 1, krs)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	keychanges.send([]string{"user/1/name"})

	ntid, data, err := s.Echo(ctx, 1, tid, keys)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	if ntid < tid {
		t.Errorf("Expect a bigger tid as %d, got: %d", tid, ntid)
	}

	if len(data) != 1 || string(data["user/1/name"]) != "some value" {
		t.Errorf("Expect data[\"user/1/name\"] to be \"some value\", got: %v", data["user/1/name"])
	}
}

type mockRestricter struct {
	data map[string]string
}

func (r mockRestricter) Restrict(ctx context.Context, uid int, keys []string) (map[string][]byte, error) {
	out := make(map[string][]byte, len(keys))
	for _, key := range keys {
		v, ok := r.data[key]
		if ok {
			out[key] = []byte(v)
			continue
		}

		switch {
		case strings.HasPrefix(key, "error"):
			return nil, fmt.Errorf("Restricter got an error")
		case strings.HasSuffix(key, "_id"):
			out[key] = []byte("1")
		case strings.HasSuffix(key, "_ids"):
			out[key] = []byte("[1,2]")
		default:
			out[key] = []byte("some value")
		}
	}
	return out, nil
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
