package redis_test

import (
	"context"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
)

// useRealRedis desides, if a real redis instance is used or a fake redis
// mock.
const useRealRedis = false

func MustKey(in string) datastore.Key {
	k, err := datastore.KeyFromString(in)
	if err != nil {
		panic(err)
	}
	return k
}

func TestUpdateOnce(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	data, err := getRedis().Update(ctx)
	if err != nil {
		t.Errorf("Update() returned an unexpected error %v", err)
	}

	expect := map[datastore.Key][]byte{
		MustKey("user/1/name"): []byte("Hubert"),
		MustKey("user/2/name"): []byte("Isolde"),
		MustKey("user/3/name"): []byte("Igor"),
	}
	if !cmpMap(data, expect) {
		t.Errorf("Update() returned %v, expected %v", data, expect)
	}
}

func TestUpdateTwice(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := getRedis()
	if _, err := r.Update(ctx); err != nil {
		t.Errorf("Update() returned an unexpected error %v", err)
	}

	keys, err := r.Update(ctx)
	if err != nil {
		t.Errorf("Update() returned an unexpected error %v", err)
	}

	expect := map[datastore.Key][]byte{}
	if !cmpMap(keys, expect) {
		t.Errorf("Update() returned %v, expected %v", keys, expect)
	}
}

func TestRedisError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	r := &redis.Redis{Conn: mockConn{err: errors.New("my error")}}
	keys, err := r.Update(ctx)
	if err == nil {
		t.Errorf("Update() did not return an error, expected one.")
	}
	if keys != nil {
		t.Errorf("Update() returned %v, expected no keys.", keys)
	}
}
