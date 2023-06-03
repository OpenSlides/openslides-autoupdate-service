package redis_test

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
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

func TestMetric(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tr := newTestRedis(t)
	defer tr.Close()

	r := redis.New(environment.ForTests(tr.Env))
	r.Wait(ctx)

	intCombine := func(value string, acc int) int {
		v, err := strconv.Atoi(value)
		if err != nil {
			fmt.Printf("can not convert value %s: %v", value, err)
			return acc
		}

		return acc + v
	}

	t.Run("Save value", func(t *testing.T) {
		m := redis.NewMetric[int](r, "test1", intCombine, time.Second, nil)

		if err := m.Save(ctx, 3); err != nil {
			t.Fatalf("save: %v", err)
		}

		v, err := m.Get(ctx)
		if err != nil {
			t.Fatalf("get: %v", err)
		}

		if v != 3 {
			t.Errorf("got %d, expected 3", v)
		}
	})

	t.Run("Save value on different instances", func(t *testing.T) {
		m1 := redis.NewMetric[int](r, "test2", intCombine, time.Second, nil)
		m2 := redis.NewMetric[int](r, "test2", intCombine, time.Second, nil)

		if err := m1.Save(ctx, 2); err != nil {
			t.Fatalf("m1 save: %v", err)
		}

		if err := m2.Save(ctx, 3); err != nil {
			t.Fatalf("m2 save: %v", err)
		}

		v1, err := m1.Get(ctx)
		if err != nil {
			t.Fatalf("v1 get: %v", err)
		}

		v2, err := m2.Get(ctx)
		if err != nil {
			t.Fatalf("v12 get: %v", err)
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

		oldInstance := redis.NewMetric[int](r, "test3", intCombine, time.Second, oldNow)

		if err := oldInstance.Save(ctx, 2); err != nil {
			t.Fatalf("save old: %v", err)
		}

		r := redis.NewMetric[int](r, "test3", intCombine, time.Second, nil)

		if err := r.Save(ctx, 3); err != nil {
			t.Fatalf("save: %v", err)
		}

		v, err := r.Get(ctx)
		if err != nil {
			t.Fatalf("get: %v", err)
		}

		if v != 3 {
			t.Errorf("got %d, expected 3", v)
		}
	})

	t.Run("Combine generic values", func(t *testing.T) {
		fn := func(value string, acc string) string {
			return value + acc
		}

		m1 := redis.NewMetric[string](r, "test4", fn, time.Second, nil)
		m2 := redis.NewMetric[string](r, "test4", fn, time.Second, nil)

		if err := m1.Save(ctx, "A"); err != nil {
			t.Fatalf("save m1: %v", err)
		}

		if err := m2.Save(ctx, "B"); err != nil {
			t.Fatalf("save m2: %v", err)
		}

		v, err := m1.Get(ctx)
		if err != nil {
			t.Fatalf("metric: %v", err)
		}

		if v != "AB" && v != "BA" {
			t.Errorf("got %s, expected 'AB' or 'BA'", v)
		}
	})
}
