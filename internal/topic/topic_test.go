package topic_test

import (
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

	got, _, err := top.Get(0)
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

	got, _, err := top.Get(0)
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

	got, _, err := top.Get(id)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	if !cmpSlice(got, expect[1:]) {
		t.Errorf("Expected to get %v, got: %v", expect[1:], got)
	}
}

func TestTopicSaveTwiceSameValue(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"key1", "key1"}
	top.Save(expect[:1])
	top.Save(expect[1:])

	got, _, err := top.Get(0)
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

	got, _, err := top.Get(0)
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

	_, _, err := top.Get(1)
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
		got, _, err := top.Get(0)
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

func TestTopicReadBlockingNoWrite(t *testing.T) {
	t.Parallel()
	top := topic.Topic{WaitDuration: 10 * time.Millisecond}
	done := make(chan struct{})
	go func() {
		var err error
		got, id, err := top.Get(1)
		if err != nil {
			t.Errorf("Did not expect an error, got: %v", err)
		}
		if len(got) != 0 {
			t.Errorf("did not expect data, got: %v", got)
		}
		if id != 1 {
			t.Errorf("Expect id to be 1, got: %d", id)
		}
		close(done)
	}()

	tick := time.NewTimer(time.Second)
	defer tick.Stop()
	select {
	case <-done:
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
	_, id, err := top.Get(0)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	if id != 2 {
		t.Fatalf("Did expect id to be 2, got: %v", id)
	}
}

func TestTopicPruneExitSoon(t *testing.T) {
	t.Parallel()
	top := topic.Topic{}
	expect := []string{"first", "key1"}
	top.Save(expect)
	top.Prune(time.Now())
	_, id, err := top.Get(0)
	if err != nil {
		t.Fatalf("Did not expect an error, got: %v", err)
	}
	if id != 1 {
		t.Fatalf("Expect id to be 1, got: %v", id)
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
