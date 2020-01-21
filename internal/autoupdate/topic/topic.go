// Package topic holds the datastructure Topic to store
// a list of strings where changes of the strings
// can be requested.
package topic

import (
	"context"
	"sync"
	"time"
)

// Topic is a datastructure that holds list of strings.
// Each save to a topic creates a new id. It is possible
// to get all strings or the strings greater than a given
// id.
//
// A Topic is save for concourent use.
// If a topic is initialized with a Closed-channel, it can be closed
// by closing this channel. It is not expected that the Closed channel is added
// or removed afterwards
type Topic struct {
	Closed chan struct{}

	mu      sync.RWMutex
	head    *node
	tail    *node
	waiting []chan struct{}
}

// node implements a linked list.
type node struct {
	id    uint64
	t     time.Time
	next  *node
	value []string
}

// Save saves a list of keys in a topic. Returns the current id.
func (t *Topic) Save(keys []string) uint64 {
	t.mu.Lock()
	defer t.mu.Unlock()

	newNode := &node{}
	id := uint64(0)
	if t.head == nil {
		t.head = newNode
	} else {
		id = t.tail.id
		t.tail.next = newNode
	}
	t.tail = newNode
	newNode.id = id + 1
	newNode.t = time.Now()
	newNode.value = keys

	for _, c := range t.waiting {
		close(c)
	}
	t.waiting = make([]chan struct{}, 0)
	return newNode.id
}

// Get returns strings from a topic. If id is 0, all strings
// are returned, else, all strings that where inserted after
// the id are returned.
//
// If the id is lower then the lowest id, an error of type
// ErrUnknownTopicID is returned.
//
// If there is no new data, Get blocks until threre is new data
// or the topic is closed or the given context is done.
func (t *Topic) Get(ctx context.Context, id uint64) ([]string, uint64, error) {
	t.mu.RLock()

	// No new data
	if t.tail == nil || id >= t.tail.id {
		c := make(chan struct{})
		t.waiting = append(t.waiting, c)
		t.mu.RUnlock()

		select {
		case <-c:
			return t.Get(ctx, id)
		case <-t.Closed:
		case <-ctx.Done():
		}
		return []string{}, id, nil
	}

	defer t.mu.RUnlock()
	maxID := t.LastID()

	if id == 0 {
		return runNode(t.head), maxID, nil
	}

	n := t.index(id)
	if n == nil {
		return nil, 0, ErrUnknownID{ID: id, First: t.head.id}
	}
	return runNode(n.next), maxID, nil
}

// LastID returns the last if of topic
func (t *Topic) LastID() uint64 {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if t.tail == nil {
		return 0
	}
	return t.tail.id
}

// Prune removes entries from the topic that are older time the given time.
func (t *Topic) Prune(until time.Time) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.head == nil || t.head.next == nil {
		return
	}

	for n := t.head; n.next != nil; n = n.next {
		if n.t.After(until) {
			return
		}
		t.head = n.next
	}
}

// runNode returns all keys from a node and the following nodes.
func runNode(n *node) []string {
	set := make(map[string]bool)
	for ; n != nil; n = n.next {
		for _, v := range n.value {
			set[v] = true
		}
	}

	out := make([]string, 0, len(set))
	for v := range set {
		out = append(out, v)
	}
	return out
}

// index returns the node with the given id.
func (t *Topic) index(id uint64) *node {
	for n := t.head; n != nil; n = n.next {
		if n.id == id {
			return n
		}
	}
	return nil
}
