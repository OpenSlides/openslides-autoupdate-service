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

	r, err := redis.New(environment.ForTests(tr.Env), nil)
	if err != nil {
		t.Fatalf("redis.New: %v", err)
	}
	r.Wait(ctx)

	done := make(chan error)
	var got map[dskey.Key][]byte
	go func() {
		data, err := r.Update(ctx)
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

	if _, err := conn.Do("XADD", "ModifiedFields", "*", "user/1/name", "Hubert", "user/2/name", "Isolde"); err != nil {
		t.Fatalf("Insert test data: %v", err)
	}

	if err := <-done; err != nil {
		t.Error(err)
	}

	expect := map[dskey.Key][]byte{
		dskey.MustKey("user/1/name"): []byte("Hubert"),
		dskey.MustKey("user/2/name"): []byte("Isolde"),
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

	r, err := redis.New(environment.ForTests(tr.Env), nil)
	if err != nil {
		t.Fatalf("redis.New: %v", err)
	}
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

func TestMetric(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tr := newTestRedis(t)
	defer tr.Close()

	t.Run("Save value", func(t *testing.T) {
		r, err := redis.New(environment.ForTests(tr.Env), nil)
		if err != nil {
			t.Fatalf("redis.New: %v", err)
		}
		r.Wait(ctx)

		if err := r.SaveMetric(ctx, "test1", 5); err != nil {
			t.Fatalf("save metric: %v", err)
		}

		v, err := r.Metric(ctx, "test1")
		if err != nil {
			t.Fatalf("metric: %v", err)
		}

		if v != 5 {
			t.Errorf("got %d, expected 5", v)
		}
	})

	t.Run("Save value on different instances", func(t *testing.T) {
		r1, err := redis.New(environment.ForTests(tr.Env), nil)
		if err != nil {
			t.Fatalf("redis.New: %v", err)
		}

		r2, err := redis.New(environment.ForTests(tr.Env), nil)
		if err != nil {
			t.Fatalf("redis.New: %v", err)
		}

		if err := r1.SaveMetric(ctx, "test2", 2); err != nil {
			t.Fatalf("save metric: %v", err)
		}

		if err := r2.SaveMetric(ctx, "test2", 3); err != nil {
			t.Fatalf("save metric: %v", err)
		}

		v1, err := r1.Metric(ctx, "test2")
		if err != nil {
			t.Fatalf("metric: %v", err)
		}

		v2, err := r2.Metric(ctx, "test2")
		if err != nil {
			t.Fatalf("metric: %v", err)
		}

		if v1 != 5 {
			t.Errorf("v1: got %d, expected 5", v1)
		}

		if v2 != 5 {
			t.Errorf("v2: got %d, expected 5", v2)
		}
	})

	t.Run("Ignore old instances", func(t *testing.T) {
		oldNow := func() time.Time {
			return time.Date(2023, time.January, 1, 5, 15, 0, 0, time.UTC)
		}

		oldInstance, err := redis.New(environment.ForTests(tr.Env), oldNow)
		if err != nil {
			t.Fatalf("redis.New: %v", err)
		}

		if err := oldInstance.SaveMetric(ctx, "test3", 2); err != nil {
			t.Fatalf("save metric: %v", err)
		}

		r, err := redis.New(environment.ForTests(tr.Env), nil)
		if err != nil {
			t.Fatalf("redis.New: %v", err)
		}

		if err := r.SaveMetric(ctx, "test3", 3); err != nil {
			t.Fatalf("save metric: %v", err)
		}

		v, err := r.Metric(ctx, "test3")
		if err != nil {
			t.Fatalf("metric: %v", err)
		}

		if v != 3 {
			t.Errorf("got %d, expected 3", v)
		}
	})
}
