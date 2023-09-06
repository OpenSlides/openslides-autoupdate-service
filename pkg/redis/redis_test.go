package redis_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
)

func TestUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tr := newTestRedis(t)
	defer tr.Close()

	r := redis.New(environment.ForTests(tr.Env))
	r.Wait(ctx)

	done := make(chan error)
	var got map[dskey.Key][]byte
	go r.Update(ctx, func(data map[dskey.Key][]byte, err error) {
		if err != nil {
			done <- fmt.Errorf("Update() returned an unexpected error %w", err)
		}

		got = data
		done <- nil
	})

	time.Sleep(20 * time.Millisecond)
	conn, err := tr.conn(ctx)
	if err != nil {
		t.Fatalf("Creating test connection: %v", err)
	}

	if _, err := conn.Do("XADD", "ModifiedFields", "*", "user/1/username", "Hubert", "user/2/username", "Isolde"); err != nil {
		t.Fatalf("Insert test data: %v", err)
	}

	if err := <-done; err != nil {
		t.Error(err)
	}

	expect := map[dskey.Key][]byte{
		dskey.MustKey("user/1/username"): []byte("Hubert"),
		dskey.MustKey("user/2/username"): []byte("Isolde"),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Update() returned %v, expected %v", got, expect)
	}
}

func TestLogout(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tr := newTestRedis(t)
	defer tr.Close()

	r := redis.New(environment.ForTests(tr.Env))
	r.Wait(ctx)

	done := make(chan error)
	var got []string
	go func() {
		data, err := r.LogoutEvent(ctx)
		if err != nil {
			done <- fmt.Errorf("Update() returned an unexpected error %w", err)
		}

		got = data
		done <- nil
	}()

	time.Sleep(20 * time.Millisecond)
	conn, err := tr.conn(ctx)
	if err != nil {
		t.Fatalf("Creating test connection: %v", err)
	}

	if _, err := conn.Do("XADD", "logout", "*", "some Key", "Hubert", "sessionId", "12345", "sessionId", "6789"); err != nil {
		t.Fatalf("Insert test data: %v", err)
	}

	if err := <-done; err != nil {
		t.Error(err)
	}

	expect := []string{"12345", "6789"}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("LogoutEvent() returned %v, expected %v", got, expect)
	}
}
