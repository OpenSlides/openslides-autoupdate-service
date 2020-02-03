package topic_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/topic"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}

	top.Add("v1", "v2")

	_, got, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	want := values("v1", "v2")
	if !cmpSlice(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestAddTwice(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}

	top.Add("v1")
	top.Add("v2")

	_, got, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	want := values("v1", "v2")
	if !cmpSlice(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestAddTwiceSameValue(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}

	top.Add("value")
	top.Add("value")

	_, got, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	want := values("value")
	if !cmpSlice(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestGetSecond(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	id := top.Add("v1")
	top.Add("v2")

	_, got, err := top.Get(context.Background(), id)

	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	want := values("v2")
	if !cmpSlice(got, want) {
		t.Errorf("Get(%d) == %v, want %v", id, got, want)
	}
}

func TestPrune(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	top.Add("first")
	top.Add("second")
	ti := time.Now()
	top.Add("third")
	top.Add("fourth")

	top.Prune(ti)

	_, got, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	want := values("third", "fourth")
	if !cmpSlice(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestGetPrunedID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	top.Add("first")
	top.Add("second")
	ti := time.Now()
	top.Add("third")
	top.Add("fourth")

	top.Prune(ti)

	_, _, err := top.Get(context.Background(), 1)
	topicErr, ok := err.(topic.ErrUnknownID)
	if !ok {
		t.Errorf("Expected err to be a topic.ErrUnknownID, got: %v", err)
	}
	if topicErr.First != 3 {
		t.Errorf("Expected the first id in the error to be 3, got: %d", topicErr.First)
	}
	if topicErr.ID != 1 {
		t.Errorf("Expected the id in the topic to be 1, got: %d", topicErr.ID)
	}
	want := "id 1 is unknown in topic. Lowest id is 3"
	if got := topicErr.Error(); got != want {
		t.Errorf("Got error message \"%s\", want \"%s\"", got, want)
	}
}

func TestDontPruneLastNode(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	top.Add("value")
	top.Add("value")

	top.Prune(time.Now())

	id, _, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if id != 2 {
		t.Errorf("Got id %d, want 2", id)
	}
}

func TestPruneOneValue(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	top.Add("value")

	top.Prune(time.Now())

	id, _, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if id != 1 {
		t.Errorf("Got id %d, want 1", id)
	}
}

func TestLastID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	top.Add("value")
	top.Add("value")
	top.Add("value")

	got := top.LastID()

	if got != uint64(3) {
		t.Errorf("LastID() == %d, want 3", got)
	}
}

func TestEmptyLastID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}

	got := top.LastID()

	if got != 0 {
		t.Errorf("LastID() == %d, want 0", got)
	}
}

func TestGetBlocking(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	done := make(chan []string)
	go func() {
		time.Sleep(time.Millisecond)
		top.Add("value")
	}()
	go func() {
		var err error
		_, got, err := top.Get(context.Background(), 0)
		if err != nil {
			t.Errorf("Did not expect an error, got: %v", err)
		}
		done <- got
	}()

	tick := time.NewTimer(100 * time.Millisecond)
	defer tick.Stop()
	select {
	case got := <-done:
		if len(got) != 1 || got[0] != "value" {
			t.Errorf("Expected to get [value] got: %v", got)
		}
	case <-tick.C:
		t.Errorf("Expected to get the data in time, wait for 100 Milliseconds")
	}
}

func TestBlockUntilClose(t *testing.T) {
	t.Parallel()
	closed := make(chan struct{})
	top := topic.Topic{Closed: closed}
	done := make(chan struct{})
	go func() {
		if _, _, err := top.Get(context.Background(), 1); err != nil {
			t.Errorf("Did not expect an error, got: %v", err)
		}
		close(done)
	}()

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case <-done:
		t.Errorf("Expect Get() not to return becore Close() is called")
	case <-timer.C:
		close(closed)
	}

	timer.Reset(100 * time.Millisecond)
	select {
	case <-done:
	case <-timer.C:
		t.Errorf("Expect Get() to return after Close() is called")
	}
}

func TestBlockUntilContexDone(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	ctx, closeCtx := context.WithCancel(context.Background())
	defer closeCtx()
	done := make(chan struct{})
	go func() {
		if _, _, err := top.Get(ctx, 1); err != nil {
			t.Errorf("Did not expect an error, got: %v", err)
		}
		close(done)
	}()

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case <-done:
		t.Errorf("Expect Get() not to return becore Close() is called")
	case <-timer.C:
		closeCtx()
	}

	timer.Reset(100 * time.Millisecond)
	select {
	case <-done:
	case <-timer.C:
		t.Errorf("Expect Get() to return after Close() is called")
	}
}

func TestBlockOnHighestID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	id := top.Add("value")
	ctx, closeCtx := context.WithCancel(context.Background())
	defer closeCtx()
	done := make(chan struct{})
	go func() {
		if _, _, err := top.Get(ctx, id); err != nil {
			t.Errorf("Did not expect an error, got: %v", err)
		}
		close(done)
	}()

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case <-done:
		t.Errorf("Expect Get() not to return becore Close() is called")
	case <-timer.C:
		closeCtx()
	}

	timer.Reset(100 * time.Millisecond)
	select {
	case <-done:
	case <-timer.C:
		t.Errorf("Expect Get() to return after Close() is called")
	}
}

func TestGetZeroAfterClosed(t *testing.T) {
	t.Parallel()
	closed := make(chan struct{})
	top := topic.Topic{Closed: closed}
	top.Add("value")
	close(closed)

	_, got, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Expect no error, got: %v", err)
	}
	if len(got) == 0 {
		t.Errorf("Got %d keys, want 0", len(got))
	}
}

func TestGetFutureAfterClosed(t *testing.T) {
	t.Parallel()
	closed := make(chan struct{})
	top := topic.Topic{Closed: closed}
	top.Add("value")
	close(closed)
	done := make(chan struct{})

	go func() {
		defer close(done)
		_, got, err := top.Get(context.Background(), 2)
		if err != nil {
			t.Errorf("Expect no error, got: %v", err)
		}

		if len(got) != 0 {
			t.Errorf("Got %d, want 0", len(got))
		}
	}()

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()
	select {
	case <-done:
	case <-timer.C:
		t.Errorf("Expect Get() not return immediately")
	}
}

func cmpSlice(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}

	sort.Strings(one)
	sort.Strings(two)
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}

func values(vs ...string) []string {
	return vs
}
