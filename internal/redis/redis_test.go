package redis_test

import (
	"errors"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/redis"
)

// useRealRedis desides, if a real redis instance is used or a fake redis
// mock.
const useRealRedis = false

func TestKeysChangedOnce(t *testing.T) {
	keys, err := getRedis().KeysChanged()
	if err != nil {
		t.Errorf("KeysChanged() returned an unexpected error %v", err)
	}

	expect := []string{"key1", "key2", "key3"}
	if !cmpSlice(keys, expect) {
		t.Errorf("KeysChanged() returned %v, expected %v", keys, expect)
	}
}

func TestKeysChangedTwice(t *testing.T) {
	r := getRedis()
	if _, err := r.KeysChanged(); err != nil {
		t.Errorf("KeysChanged() returned an unexpected error %v", err)
	}

	keys, err := r.KeysChanged()
	if err != nil {
		t.Errorf("KeysChanged() returned an unexpected error %v", err)
	}

	expect := []string{}
	if !cmpSlice(keys, expect) {
		t.Errorf("KeysChanged() returned %v, expected %v", keys, expect)
	}
}

func TestRedisError(t *testing.T) {
	r := &redis.Service{Conn: mockConn{err: errors.New("my error")}}
	keys, err := r.KeysChanged()
	if err == nil {
		t.Errorf("KeysChange() did not return an error, expected one.")
	}
	if keys != nil {
		t.Errorf("KeysChanged() returned %v, expected no keys.", keys)
	}
}
