package redis_test

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/redis"
)

func TestMetric(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tr := newTestRedis(t)
	defer tr.Close()

	r := redis.New(environment.ForTests(tr.Env))
	r.Wait(ctx)

	t.Run("Save value", func(t *testing.T) {
		m := redis.NewMetric[int](r, "test1", intCombiner{}, time.Second, nil)

		if err := m.Save(ctx, strconv.Itoa(3)); err != nil {
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
		m1 := redis.NewMetric[int](r, "test2", intCombiner{}, time.Second, nil)
		m2 := redis.NewMetric[int](r, "test2", intCombiner{}, time.Second, nil)

		if err := m1.Save(ctx, strconv.Itoa(2)); err != nil {
			t.Fatalf("m1 save: %v", err)
		}

		if err := m2.Save(ctx, strconv.Itoa(3)); err != nil {
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

		oldInstance := redis.NewMetric[int](r, "test3", intCombiner{}, time.Second, oldNow)

		if err := oldInstance.Save(ctx, strconv.Itoa(2)); err != nil {
			t.Fatalf("save old: %v", err)
		}

		r := redis.NewMetric[int](r, "test3", intCombiner{}, time.Second, nil)

		if err := r.Save(ctx, strconv.Itoa(3)); err != nil {
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

	t.Run("Combine json map", func(t *testing.T) {
		m1 := redis.NewMetric[map[int]int](r, "test4", mapIntCombiner{}, time.Second, nil)
		m2 := redis.NewMetric[map[int]int](r, "test4", mapIntCombiner{}, time.Second, nil)

		value, _ := json.Marshal(map[int]int{1: 2, 3: 4})

		if err := m1.Save(ctx, string(value)); err != nil {
			t.Fatalf("save m1: %v", err)
		}

		value, _ = json.Marshal(map[int]int{1: 1, 2: 2})

		if err := m2.Save(ctx, string(value)); err != nil {
			t.Fatalf("save m2: %v", err)
		}

		got, err := m1.Get(ctx)
		if err != nil {
			t.Fatalf("get: %v", err)
		}

		expect := map[int]int{1: 3, 2: 2, 3: 4}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("got %v, expected %v", got, expect)
		}
	})
}

type intCombiner struct{}

func (intCombiner) Combine(value string, acc int) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("convert from string %s: %w", value, err)
	}

	return acc + v, nil
}

type mapIntCombiner struct{}

func (mapIntCombiner) Combine(rawValue string, acc map[int]int) (map[int]int, error) {
	var value map[int]int
	if err := json.Unmarshal([]byte(rawValue), &value); err != nil {
		return nil, err
	}

	if acc == nil {
		acc = make(map[int]int)
	}

	for k, v := range value {
		acc[k] += v
	}

	return acc, nil
}
