package autoupdate

import (
	"context"
	"fmt"
	"github.com/openslides/openslides-autoupdate-service/internal/keysrequest"
	"strings"
	"testing"
	"time"
)

func TestPrepare(t *testing.T) {
	keychanges := newMockKeyChanged()
	s := New(mockRestricter{}, keychanges)
	defer s.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	kr, err := keysrequest.FromJSON(strings.NewReader(`{"ids":[1],"collection":"user","fields":{"name":null}}`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	_, _, data, err := s.prepare(ctx, 1, kr)
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
	kr, err := keysrequest.FromJSON(strings.NewReader(`{"ids":[1],"collection":"user","fields":{"name":null}}`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	tid, keys, _, err := s.prepare(ctx, 1, kr)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	ntid, data, err := s.echo(ctx, 1, tid, keys)
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
	kr, err := keysrequest.FromJSON(strings.NewReader(`{"ids":[1],"collection":"user","fields":{"name":null}}`))
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	tid, keys, _, err := s.prepare(ctx, 1, kr)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}

	keychanges.send(KeyChanges{Updated: []string{"user/1/name"}})

	ntid, data, err := s.echo(ctx, 1, tid, keys)
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

type mockRestricter struct{}

func (r mockRestricter) Restrict(ctx context.Context, uid int, keys []string) (map[string][]byte, error) {
	out := make(map[string][]byte, len(keys))
	for _, key := range keys {
		switch {
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

func (r mockRestricter) IDsFromKey(ctx context.Context, uid int, mid int, key string) ([]int, error) {
	if strings.HasPrefix(key, "not_exist") {
		return nil, nil
	}
	if strings.HasSuffix(key, "_id") {
		return []int{1}, nil
	}
	if !strings.HasSuffix(key, "_ids") {
		return nil, fmt.Errorf("Key %s can not be a reference; expected suffex _id or _ids", key)
	}
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
}

func (r mockRestricter) IDsFromCollection(ctx context.Context, uid int, mid int, collection string) ([]int, error) {
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
}

type mockKeyChanged struct {
	c chan KeyChanges
	t *time.Ticker
}

func newMockKeyChanged() mockKeyChanged {
	m := mockKeyChanged{}
	m.c = make(chan KeyChanges, 1)
	m.t = time.NewTicker(time.Second)
	return m
}

func (m mockKeyChanged) KeysChanged() (KeyChanges, error) {
	select {
	case v := <-m.c:
		return v, nil
	case <-m.t.C:
		return KeyChanges{}, nil
	}
}

func (m mockKeyChanged) send(kc KeyChanges) {
	m.c <- kc
}

func (m mockKeyChanged) close() {
	m.t.Stop()
}
