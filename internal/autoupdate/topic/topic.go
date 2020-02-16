// Package topic is a in process pubsub system where new values have to
// be pulled instead of beeing pushed.
//
// The idea of pulling updates is inspired by Kafka or Redis-Streams. A subscriber
// does not have to register or unsubscribe to a topic and can take as long as it needs to
// process the messages. Therefore, the system is less error-prone.
package topic

import (
	"context"
	"sync"
	"time"
)

// Topic is a datastructure that holds a set of strings.
// Each time a list of streams are added to the topic, a new id is
// created. It is possible to get all strings at once or the strings that added after
// a specivic id.
//
// A Topic is save for concourent use.
// If a topic is initialized with a Closed-channel, it can be closed
// by closing this channel. It is not expected that the Closed channel is added
// or removed afterwards.
type Topic struct {
	Closed chan struct{}

	mu    sync.RWMutex
	head  *node
	tail  *node
	index map[uint64]*node

	waitingMu sync.Mutex
	waiting   []chan struct{}
}

// New creates a new topic.
func New(options ...Option) *Topic {
	top := &Topic{}

	for _, o := range options {
		o(top)
	}
	return top
}

// node implements a linked list.
type node struct {
	id    uint64
	t     time.Time
	next  *node
	value []string
}

// Add adds a list of strings to a topic. It creates a new id and returns it.
func (t *Topic) Add(value ...string) uint64 {
	t.mu.Lock()
	defer t.mu.Unlock()

	newNode := &node{}
	var id uint64
	if t.head == nil {
		t.head = newNode
	} else {
		id = t.tail.id
		t.tail.next = newNode
	}
	t.tail = newNode
	newNode.id = id + 1
	newNode.t = time.Now()
	newNode.value = value

	if t.index == nil {
		t.index = make(map[uint64]*node)
	}
	t.index[newNode.id] = newNode

	for _, c := range t.waiting {
		close(c)
	}
	// This clears the waiting slice by keeping the underlieing array. The idea is, that it
	// is very likely, that the same subscribers will subsribe again, the same size of array
	// will be needed again.
	t.waiting = t.waiting[:0]
	return newNode.id
}

// Get returns a slice of uniq strings from the topic. If id is 0, all strings
// are returned, else, all strings that where inserted after the id are returned.
//
// If the id is lower then the lowest id in the topic, an error of type
// ErrUnknownTopicID is returned.
//
// If there is no new data, Get blocks until threre is new data or the topic is closed or the
// given context is done.
func (t *Topic) Get(ctx context.Context, id uint64) (uint64, []string, error) {
	t.mu.RLock()

	// No new data
	if t.tail == nil || id >= t.tail.id {
		c := make(chan struct{})
		t.wait(c)
		t.mu.RUnlock()

		select {
		case <-c:
			return t.Get(ctx, id)
		case <-t.Closed:
		case <-ctx.Done():
		}
		return id, []string{}, nil
	}

	defer t.mu.RUnlock()

	if id == 0 {
		return t.tail.id, runNode(t.head), nil
	}

	n := t.index[id]
	if n == nil {
		return 0, nil, ErrUnknownID{ID: id, First: t.head.id}
	}
	return t.tail.id, runNode(n.next), nil
}

// LastID returns the last if of topic. Returns 0 for an empty topic.
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
		delete(t.index, n.id)
		t.head = n.next
	}
}

func (t *Topic) wait(c chan struct{}) {
	t.waitingMu.Lock()
	t.waiting = append(t.waiting, c)
	t.waitingMu.Unlock()
}

// runNode returns all strings from a node and the following nodes.
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
