package redis_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/redis"
)

// useRealRedis desides, if a real redis instance is used or a fake redis
// mock.
const useRealRedis = false

func TestUpdateOnce(t *testing.T) {
	data, err := getRedis().Update()
	if err != nil {
		t.Errorf("Update() returned an unexpected error %v", err)
	}

	expect := map[string]json.RawMessage{
		"user/1/name": []byte("Hubert"),
		"user/2/name": []byte("Isolde"),
		"user/3/name": []byte("Igor"),
	}
	if !cmpMap(data, expect) {
		t.Errorf("Update() returned %v, expected %v", data, expect)
	}
}

func TestUpdateTwice(t *testing.T) {
	r := getRedis()
	if _, err := r.Update(); err != nil {
		t.Errorf("Update() returned an unexpected error %v", err)
	}

	keys, err := r.Update()
	if err != nil {
		t.Errorf("Update() returned an unexpected error %v", err)
	}

	expect := map[string]json.RawMessage{}
	if !cmpMap(keys, expect) {
		t.Errorf("Update() returned %v, expected %v", keys, expect)
	}
}

func TestRedisError(t *testing.T) {
	r := &redis.Service{Conn: mockConn{err: errors.New("my error")}}
	keys, err := r.Update()
	if err == nil {
		t.Errorf("Update() did not return an error, expected one.")
	}
	if keys != nil {
		t.Errorf("Update() returned %v, expected no keys.", keys)
	}
}
