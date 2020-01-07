// Package topic holds the datastructure Topic to store
// a list of strings where changes of the strings
// can be requested
package topic

import (
	"sync"
	"time"
)

// Topic is a datastructure that holds list of strings.
// Each save to a topic creates a new id. It is possible
// to get all strings or the strings greater than a given
// id.
// A Topic is save for concourent use.
type Topic struct {
	// WaitDuration defines how long a get call waits for new data. Default is 1 second.
	WaitDuration time.Duration

	mu    sync.RWMutex
	first *node
	last  *node
	index map[uint64]*node

	waitingMu sync.Mutex
	waiting   []chan struct{}
}

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
	if t.first == nil {
		t.first = newNode
		t.index = make(map[uint64]*node)
	} else {
		id = t.last.id
		t.last.next = newNode
	}
	t.last = newNode
	newNode.id = id + 1
	newNode.t = time.Now()
	newNode.value = keys
	t.index[newNode.id] = newNode

	for _, c := range t.waiting {
		close(c)
	}
	t.waiting = make([]chan struct{}, 0)
	return newNode.id
}

// Get returns strings from a topic. If id is 0, all strings
// are returned, else, all strings that where inserted after
// the id are returned.
// If the id is lower then the lowest id, an error of type
// ErrUnknownTopicID is returned.
// If there is no new data, Get blocks for WaitDuration and returns
// keys if there are inserted in the meantime. If there are no
// inserted keys, it returns an empty list and the given id.
// If an id is given that is higher then the highest id Get could
// block much longer then WaitDuration, in theory it could block for
// ever. Therefore only the id returned by save() or get() should
// be used as input value.
func (t *Topic) Get(id uint64) ([]string, uint64, error) {
	t.mu.RLock()

	if t.first == nil || id > t.last.id {
		c := make(chan struct{})
		t.wait(c)
		t.mu.RUnlock()
		// Wait for c to get closed by a writer or for one second.
		d := t.WaitDuration
		if d == 0 {
			d = time.Second
		}
		timer := time.NewTimer(d)
		defer timer.Stop()
		select {
		case <-c:
			return t.Get(id)
		case <-timer.C:
			return []string{}, id, nil
		}
	}
	defer t.mu.RUnlock()

	maxID := uint64(0)
	if t.last != nil {
		maxID = t.last.id
	}

	if id == uint64(0) {
		out := runNode(t.first)
		return out, maxID, nil
	}

	n, ok := t.index[id]
	if !ok {
		return nil, 0, ErrUnknownID{ID: id, First: t.first.id}
	}

	out := runNode(n.next)
	return out, maxID, nil
}

// LastID returns the last if of topic
func (t *Topic) LastID() uint64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if t.last == nil {
		return 0
	}
	return t.last.id
}

// Prune removes entries from the topic that are older time the given time.
func (t *Topic) Prune(until time.Time) {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.first == nil || t.first.next == nil {
		return
	}

	t.index = make(map[uint64]*node)
	for n := t.first; n.next != nil; n = n.next {
		if n.t.After(until) {
			t.index[n.id] = n
			continue
		}
		t.first = n.next
	}
}

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

func (t *Topic) wait(c chan struct{}) {
	t.waitingMu.Lock()
	defer t.waitingMu.Unlock()
	t.waiting = append(t.waiting, c)
}
