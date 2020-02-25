package autoupdate_test

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MockRestricter struct {
	data map[string]string
}

func (r MockRestricter) Restrict(ctx context.Context, uid int, keys []string) (map[string]string, error) {
	out := make(map[string]string, len(keys))
	for _, key := range keys {
		v, ok := r.data[key]
		if ok {
			out[key] = v
			continue
		}

		switch {
		case strings.HasPrefix(key, "error"):
			return nil, fmt.Errorf("Restricter got an error")
		case strings.HasSuffix(key, "_id"):
			out[key] = "1"
		case strings.HasSuffix(key, "_ids"):
			out[key] = "[1,2]"
		default:
			out[key] = "some value"
		}
	}
	return out, nil
}

type mockKeyChanged struct {
	c chan []string
	t *time.Ticker
}

func newMockKeyChanged() *mockKeyChanged {
	m := mockKeyChanged{}
	m.c = make(chan []string, 1)
	m.t = time.NewTicker(time.Second)
	return &m
}

func (m *mockKeyChanged) KeysChanged() ([]string, error) {
	select {
	case v := <-m.c:
		return v, nil
	case <-m.t.C:
		return nil, nil
	}
}

func (m *mockKeyChanged) send(keys []string) {
	m.c <- keys
}

func (m *mockKeyChanged) close() {
	m.t.Stop()
}

type mockKeysBuilder struct {
	keys []string
}

func (m mockKeysBuilder) Update([]string) error {
	return nil
}

func (m mockKeysBuilder) Keys() []string {
	return m.keys
}

// keyValue is a helper to create a map from string to string.
type keyValue map[string]string

func (kv keyValue) m() map[string]string {
	out := make(map[string]string)
	for key, value := range kv {
		out[key] = value
	}
	return out
}

// keys is a helper function to create a slice of strings.
func keys(keys ...string) []string { return keys }
