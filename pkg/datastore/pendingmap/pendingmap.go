package pendingmap

import (
	"context"
	"sync"
)

// PendingMap is like a map but values can be in a pending state. When a value
// is requested, the function blocks until the pending state is done.
//
// To get values, just use pendingmap.Get(ctx, key).
//
// Before calculating a value, set it as pending with
// pendingmap.MarkPending(key). When a key is marked as pending, it can not
// retrieved until it is set or unmarked.
//
// To set a value, there are different methods. SetIfExist() sets values if they
// are pending or already stored.
type PendingMap[K comparable, V any] struct {
	sync.RWMutex
	data    map[K]V
	pending map[K]chan struct{}
}

// New initializes a pendingDict.
func New[K comparable, V any]() *PendingMap[K, V] {
	return &PendingMap[K, V]{
		data:    map[K]V{},
		pending: map[K]chan struct{}{},
	}
}

// Get returns a list o keys from the pendingMap.
//
// The function blocks, until all values are not pending anymore.
//
// Returns nil for a value that does not exist.
//
// Makes sure that all values are returned at the same version. So if setIfExist
// is called while the function is running, then all values are returned at the
// latest version.
//
// Expects, that all keys are either pending or in the data. It is not allowed,
// that a key is not pending when this starts and gets pending while it runs.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded
func (pm *PendingMap[K, V]) Get(ctx context.Context, keys ...K) (map[K]V, error) {
	if err := pm.waitForPending(ctx, keys); err != nil {
		return nil, err
	}

	out := make(map[K]V, len(keys))
	reading(pm, func() {
		for _, k := range keys {
			out[k] = pm.data[k]
		}
	})

	return out, nil
}

// MarkPending marks one or more keys as pending.
//
// Skips keys that are already pending or are already in the datastructure.
//
// Returns all keys that where marked as pending (did not exist).
func (pm *PendingMap[K, V]) MarkPending(keys ...K) []K {
	var needMark []K
	reading(pm, func() {
		for _, key := range keys {
			if _, ok := pm.data[key]; ok {
				continue
			}
			if _, ok := pm.pending[key]; ok {
				continue
			}

			needMark = append(needMark, key)
		}
	})

	if len(needMark) == 0 {
		return nil
	}

	marked := make([]K, 0, len(needMark))
	pm.Lock()
	defer pm.Unlock()

	for _, key := range needMark {
		if _, ok := pm.pending[key]; ok {
			// It can happen, that another caller has already set the key.
			continue
		}

		pm.pending[key] = make(chan struct{})
		marked = append(marked, key)
	}
	return marked
}

// UnMarkPending sets any key that is still pending not to be pending.
//
// Skips keys that are already pending or are already in the database.
func (pm *PendingMap[K, V]) UnMarkPending(keys ...K) {
	pm.Lock()
	defer pm.Unlock()

	for _, key := range keys {
		if _, ok := pm.data[key]; ok {
			continue
		}
		pending := pm.pending[key]

		if pending == nil {
			continue
		}

		close(pending)
		delete(pm.pending, key)
	}
}

// waitForPending blocks until all the given keys are not pending anymore.
//
// Expects, that all keys are either pending or in the data. It is not allowed,
// that a key is not pending when this starts and gets pending whil it runs.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded
func (pm *PendingMap[K, V]) waitForPending(ctx context.Context, keys []K) error {
	for _, k := range keys {
		var pending chan struct{}
		reading(pm, func() {
			pending = pm.pending[k]
		})

		if pending == nil {
			continue
		}

		select {
		case <-pending:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}

// SetIfExist updates values, but only if the key already exists or is pending.
//
// If the key is pending, it is unmarked and all listeners are informed.
func (pm *PendingMap[K, V]) SetIfExist(data map[K]V) {
	pm.Lock()
	defer pm.Unlock()

	for key, value := range data {
		pending := pm.pending[key]
		_, exists := pm.data[key]

		if pending == nil && !exists {
			return
		}

		pm.data[key] = value

		if pending != nil {
			close(pending)
			delete(pm.pending, key)
		}
	}
}

// SetIfPending updates values but only if the key is pending.
//
// Informs all listeners.
func (pm *PendingMap[K, V]) SetIfPending(data map[K]V) {
	pm.Lock()
	defer pm.Unlock()

	for key, value := range data {
		if pending, isPending := pm.pending[key]; isPending {
			pm.data[key] = value
			close(pending)
			delete(pm.pending, key)
		}
	}
}

// SetEmptyIfPending set all keys that are still pending to nil.
func (pm *PendingMap[K, V]) SetEmptyIfPending(keys ...K) {
	pm.Lock()
	defer pm.Unlock()

	for _, key := range keys {
		if pending, isPending := pm.pending[key]; isPending {
			var v V
			pm.data[key] = v
			close(pending)
			delete(pm.pending, key)
		}
	}
}

func (pm *PendingMap[K, V]) Len() int {
	pm.RLock()
	defer pm.RUnlock()

	return len(pm.data)
}

// // size returns the size of all values in the cache in bytes.
// func (pm *pendingMap[K,V]) size() int {
// 	pm.RLock()
// 	defer pm.RUnlock()

// 	var size int
// 	for _, v := range pm.data {
// 		size += len(v)
// 	}
// 	return size
// }

type rlocker interface {
	RLock()
	RUnlock()
}

func reading(l rlocker, cmd func()) {
	l.RLock()
	defer l.RUnlock()
	cmd()
}
