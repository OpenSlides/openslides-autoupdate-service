package datastore_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

func TestVoteCountSourceGet(t *testing.T) {
	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := datastore.NewVoteCountSource(ts.URL)
	go source.Connect(ctx, func(error) {})

	key1 := datastore.MustKey("poll/1/vote_count")

	t.Run("no data from vote-service", func(t *testing.T) {
		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if got[key1] != nil {
			t.Errorf("Get() without any data returned %v, expected nil", got[key1])
		}
	})

	t.Run("first data from vote-service", func(t *testing.T) {
		sender <- `{"1":42}`
		time.Sleep(time.Millisecond)
		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != "42" {
			t.Errorf("Get() after first data returned `%v`, expected `42`", got[key1])
		}
	})

	t.Run("second data from vote-service", func(t *testing.T) {
		sender <- `{"1":43}`
		time.Sleep(time.Millisecond)
		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != "43" {
			t.Errorf("Get() after first data returned `%v`, expected `42`", got[key1])
		}
	})

	t.Run("again data from vote-service", func(t *testing.T) {
		sender <- `{"1":44}`
		time.Sleep(time.Millisecond)
		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != "44" {
			t.Errorf("Get() after first data returned `%v`, expected `42`", got[key1])
		}
	})

	t.Run("receive 0", func(t *testing.T) {
		sender <- `{"1":0}`
		time.Sleep(time.Millisecond)
		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if got[key1] != nil {
			t.Errorf("Get() after receiving 0 returned %v, expected nil", got[key1])
		}
	})
}

func TestVoteCountSourceUpdate(t *testing.T) {
	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source := datastore.NewVoteCountSource(ts.URL)
	go source.Connect(ctx, func(error) {})

	key1 := datastore.MustKey("poll/1/vote_count")

	t.Run("no data from vote-service", func(t *testing.T) {
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Millisecond)
		defer cancel()

		_, err := source.Update(ctxTimeout)
		if err != context.DeadlineExceeded {
			t.Fatalf("Update: %v, expected context.DeadlineExceeded", err)
		}
	})

	t.Run("first data from vote-service", func(t *testing.T) {
		sender <- `{"1":42}`
		time.Sleep(time.Millisecond)
		got, err := source.Update(ctx)
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[datastore.Key][]byte{key1: []byte("42")}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})

	t.Run("second data from vote-service", func(t *testing.T) {
		sender <- `{"1":43}`
		time.Sleep(time.Millisecond)
		got, err := source.Update(ctx)
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[datastore.Key][]byte{key1: []byte("43")}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})

	t.Run("receive 0", func(t *testing.T) {
		sender <- `{"1":0}`
		time.Sleep(time.Millisecond)
		got, err := source.Update(ctx)
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[datastore.Key][]byte{key1: nil}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})
}
