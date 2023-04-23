package datastore

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

func TestVoteCountSourceGet(t *testing.T) {
	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
		w.(http.Flusher).Flush()
		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	source := newVoteCountSource(env)
	eventer := func() (<-chan time.Time, func() bool) { return make(chan time.Time), func() bool { return true } }

	waitForResponse(ctx, source, func() {
		go source.Connect(ctx, eventer, func(error) {})
	})

	key1 := dskey.MustKey("poll/1/vote_count")

	t.Run("no data from vote-service", func(t *testing.T) {
		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if got[key1] != nil {
			t.Errorf("Get() without any data returned %s, expected nil", got[key1])
		}
	})

	t.Run("first data from vote-service", func(t *testing.T) {
		waitForResponse(ctx, source, func() {
			sender <- `{"1":42}`
		})

		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != "42" {
			t.Errorf("Get() after first data returned `%s`, expected `42`", got[key1])
		}
	})

	t.Run("second data from vote-service", func(t *testing.T) {
		waitForResponse(ctx, source, func() {
			sender <- `{"1":43}`
		})

		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != "43" {
			t.Errorf("Get() after first data returned `%s`, expected `43`", got[key1])
		}
	})

	t.Run("again data from vote-service", func(t *testing.T) {
		waitForResponse(ctx, source, func() {
			sender <- `{"1":44}`
		})

		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if string(got[key1]) != "44" {
			t.Errorf("Get() after first data returned `%s`, expected `44`", got[key1])
		}
	})

	t.Run("receive 0", func(t *testing.T) {
		waitForResponse(ctx, source, func() {
			sender <- `{"1":0}`
		})

		got, err := source.Get(ctx, key1)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		if got[key1] != nil {
			t.Errorf("Get() after receiving 0 returned %s, expected nil", got[key1])
		}
	})
}

func TestVoteCountSourceUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{}`)
		w.(http.Flusher).Flush()

		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	source := newVoteCountSource(env)
	eventer := func() (<-chan time.Time, func() bool) { return make(chan time.Time), func() bool { return true } }

	waitForResponse(ctx, source, func() {
		go source.Connect(ctx, eventer, func(error) {})
	})

	key1 := dskey.MustKey("poll/1/vote_count")

	t.Run("no data from vote-service", func(t *testing.T) {
		ctxTimeout, cancel := context.WithTimeout(ctx, time.Millisecond)
		defer cancel()

		_, err := source.Update(ctxTimeout)
		if err != context.DeadlineExceeded {
			t.Fatalf("Update: %v, expected context.DeadlineExceeded", err)
		}
	})

	t.Run("first data from vote-service", func(t *testing.T) {
		got, err := updateResult(ctx, source, func() {
			sender <- `{"1":42}`
		})
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[dskey.Key][]byte{key1: []byte("42")}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})

	t.Run("second data from vote-service", func(t *testing.T) {
		got, err := updateResult(ctx, source, func() {
			sender <- `{"1":43}`
		})
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[dskey.Key][]byte{key1: []byte("43")}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})

	t.Run("receive 0", func(t *testing.T) {
		got, err := updateResult(ctx, source, func() {
			sender <- `{"1":0}`
		})
		if err != nil {
			t.Fatalf("Update: %v", err)
		}

		expect := map[dskey.Key][]byte{key1: nil}
		if !reflect.DeepEqual(got, expect) {
			t.Errorf("Update() returned %v, expected %v", got, expect)
		}
	})
}

func TestReconnect(t *testing.T) {
	msg := `{"1":23}`
	sender := make(chan struct{})
	var counter int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, msg)
		w.(http.Flusher).Flush()
		counter++
		<-sender
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	event := make(chan time.Time)
	close(event)
	eventer := func() (<-chan time.Time, func() bool) {
		return event, func() bool { return false }
	}

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	source := newVoteCountSource(env)
	go source.Connect(ctx, eventer, func(error) {})

	sender <- struct{}{} // Close connection so there is a reconnect
	sender <- struct{}{} // Close connection again

	if counter < 2 {
		t.Errorf("Got %d connections, expected 2", counter)
	}
}

func TestReconnectWhenDeletedBetween(t *testing.T) {
	msg := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, <-msg)
		// This handler returns after the first data is send. This triggers a reconnect.
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Use a closed channel as eventer so the connection is reestablished immediatly.
	event := make(chan time.Time)
	close(event)
	eventer := func() (<-chan time.Time, func() bool) {
		return event, func() bool { return false }
	}

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	source := newVoteCountSource(env)
	go source.Connect(ctx, eventer, func(error) {})
	msg <- `{"1":23,"2":42}`
	msg <- `{"1":23}`

	source.Update(ctx) // wait for the service to process the update.

	key := dskey.MustKey("poll/2/vote_count")
	data, err := source.Get(ctx, key)
	if err != nil {
		t.Errorf("Get: %v", err)
	}

	if string(data[key]) != `42` {
		t.Errorf("Get for deleted key returned `%s`, expected 42", data[key])
	}
}

func TestGetWithoutConnect(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sender := make(chan string)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for msg := range sender {
			fmt.Fprintln(w, msg)
			w.(http.Flusher).Flush()
		}
	}))

	host, port, schema := parseURL(ts.URL)
	env := environment.ForTests(map[string]string{
		"VOTE_HOST":     host,
		"VOTE_PORT":     port,
		"VOTE_PROTOCOL": schema,
	})

	source := newVoteCountSource(env)

	key := dskey.MustKey("poll/1/vote_count")

	ctxTimeout, cancel := context.WithTimeout(ctx, time.Millisecond)
	defer cancel()

	_, err := source.Get(ctxTimeout, key)
	if err != context.DeadlineExceeded {
		t.Fatalf("Update: %v, expected context.DeadlineExceeded", err)
	}
}

// waitForResponse calls the given function and waits until the data is
// processed.
func waitForResponse(ctx context.Context, source *voteCountSource, fn func()) {
	received := make(chan struct{})
	go func() {
		source.Update(ctx) // wait for the service to process the update.
		close(received)
	}()

	fn()

	<-received
}

// updateResult returns the return values from source.Update after the given function is processed.
func updateResult(ctx context.Context, source *voteCountSource, fn func()) (map[dskey.Key][]byte, error) {
	type dataErr struct {
		data map[dskey.Key][]byte
		err  error
	}

	done := make(chan dataErr)
	go func() {
		v, err := source.Update(ctx)
		done <- dataErr{v, err}
	}()

	fn()

	got := <-done

	return got.data, got.err
}
