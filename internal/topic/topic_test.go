package topic_test

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/topic"
)

func TestTopicSave(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"key1", "key2"}
	top.Save(expect)

	got, _, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if !cmpSlice(got, expect) {
		t.Errorf("Expected to get %v, got: %v", expect, got)
	}
}

func TestTopicSaveTwice(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"key1", "key2"}
	top.Save(expect[:1])
	top.Save(expect[1:])

	got, _, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if !cmpSlice(got, expect) {
		t.Errorf("Expected to get %v, got: %v", expect, got)
	}
}

func TestTopicReadSecond(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"key1", "key2"}
	id := top.Save(expect[:1])
	top.Save(expect[1:])

	got, _, err := top.Get(context.Background(), id)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if !cmpSlice(got, expect[1:]) {
		t.Errorf("Expected Get(%d) to return %v, got: %v", id, expect[1:], got)
	}
}

func TestTopicSaveTwiceSameValue(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"key1", "key1"}
	top.Save(expect[:1])
	top.Save(expect[1:])

	got, _, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if !cmpSlice(got, expect[1:]) {
		t.Errorf("Expected to get %v, got: %v", expect[1:], got)
	}
}

func TestTopicPrune(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"first", "key1"}
	top.Save(expect)
	top.Save(expect)
	ti := time.Now()
	top.Save(expect[1:])
	top.Save(expect[1:])

	top.Prune(ti)

	got, _, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if !cmpSlice(got, expect[1:]) {
		t.Errorf("Expected to get %v, got: %v", expect[1:], got)
	}
}

func TestTopicReadPrunedID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"first", "key1"}
	top.Save(expect)
	top.Save(expect)
	ti := time.Now()
	top.Save(expect[1:])
	top.Save(expect[1:])

	top.Prune(ti)

	_, _, err := top.Get(context.Background(), 1)
	topicErr, ok := err.(topic.ErrUnknownID)
	if !ok {
		t.Errorf("Expected topic.ErrUnknownID, got: %v", err)
	}
	if topicErr.First != 3 {
		t.Errorf("Expected the first id in the error to be 2, got: %d", topicErr.First)
	}
	if topicErr.ID != 1 {
		t.Errorf("Expected the id in the topic to be 1, got: %d", topicErr.ID)
	}
	expected := "id 1 is unknown in topic. Lowest id is 3"
	if got := topicErr.Error(); got != expected {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expected, got)
	}
}

func TestTopicReadBlocking(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	done := make(chan []string)
	go func() {
		time.Sleep(time.Millisecond)
		top.Save([]string{"value"})
	}()
	go func() {
		var err error
		got, _, err := top.Get(context.Background(), 0)
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

func TestTopicLastID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	top.Save([]string{"data"})
	top.Save([]string{"data"})
	top.Save([]string{"data"})

	if got := top.LastID(); got != uint64(3) {
		t.Errorf("Expected LastID to return 3, got: %d", got)
	}
}

func TestTopicEmptyLastID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}

	if got := top.LastID(); got != 0 {
		t.Errorf("Expected LastID to return 3, got: %d", got)
	}
}

func TestTopicDontPruneLastNode(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"first", "key1"}
	top.Save(expect)
	top.Save(expect)
	top.Prune(time.Now())
	_, id, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	if id != 2 {
		t.Fatalf("Expect id to be 2, got: %v", id)
	}
}

func TestTopicPruneExitSoon(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"first", "key1"}
	top.Save(expect)
	top.Prune(time.Now())
	_, id, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	if id != 1 {
		t.Fatalf("Expect id to be 1, got: %v", id)
	}
}

func TestTopicBlockUntilClose(t *testing.T) {
	t.Parallel()
	closed := make(chan struct{})
	top := topic.Topic{Closed: closed}

	done := make(chan struct{})
	go func() {
		if _, _, err := top.Get(context.Background(), 1); err != nil {
			t.Fatalf("Did not expect an error, got: %v", err)
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

func TestTopicBlockUntilCtxDone(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	ctx, closeCtx := context.WithCancel(context.Background())
	defer closeCtx()

	done := make(chan struct{})
	go func() {
		if _, _, err := top.Get(ctx, 1); err != nil {
			t.Fatalf("Did not expect an error, got: %v", err)
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

func TestTopicBlockOnHighestID(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	id := top.Save([]string{"data"})
	ctx, closeCtx := context.WithCancel(context.Background())
	defer closeCtx()

	done := make(chan struct{})
	go func() {
		if _, _, err := top.Get(ctx, id); err != nil {
			t.Fatalf("Did not expect an error, got: %v", err)
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

func TestTopicGet0AfterClosed(t *testing.T) {
	t.Parallel()
	closed := make(chan struct{})
	top := topic.Topic{Closed: closed}
	top.Save([]string{"foo"})
	close(closed)

	keys, _, err := top.Get(context.Background(), 0)
	if err != nil {
		t.Errorf("Expect no error, got: %v", err)
	}
	if len(keys) == 0 {
		t.Errorf("Expected %d keys in the topic, got none", 1)
	}
}

func TestTopicGetFutureAfterClosed(t *testing.T) {
	t.Parallel()
	closed := make(chan struct{})
	top := topic.Topic{Closed: closed}
	top.Save([]string{"foo"})
	close(closed)

	done := make(chan struct{})

	go func() {
		defer close(done)
		keys, _, err := top.Get(context.Background(), 2)
		if err != nil {
			t.Errorf("Expect no error, got: %v", err)
		}

		if len(keys) != 0 {
			t.Errorf("Expected empty data. got %d keys", len(keys))
		}
	}()

	timer := time.NewTimer(time.Millisecond)
	defer timer.Stop()

	select {
	case <-done:
	case <-timer.C:
		t.Errorf("Expect Get() not return immediatly")
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
