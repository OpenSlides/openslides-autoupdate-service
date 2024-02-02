package pendingmap

import (
	"context"
	"errors"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// ErrNotExist is returned from pendingmap.Get() when a key was not pending at
// the beginning but did not exist at the end. This can happen when a key gets
// unmarked.
var ErrNotExist = errors.New("key does not exist")

// PendingMap is like a map but values can be in a pending state. When a value
// is requested, the function blocks until the pending state is done.
//
// To get values, just use pendingmap.Get(ctx, key).
//
// Each key has one of three states: Not exists, pending, exists.
//
// A key that exists can be updated but not deleted. So if a key exists once, it
// will be in the existing state forever.
//
// A key that not exists can be set to pending or to existing.
//
// A key that is pending can get existing or (in error cases) go back to not
// existing.
//
// Before calculating a value, set it as pending with
// pendingmap.MarkPending(key). When a key is marked as pending, it can not
// retrieved until it is set or unmarked.
//
// pendingmap.UnMarkPending(key) can be used to undo a MarkPending call. This is
// usefull in error cases where a value should be freed for other callers.
//
// To set a value, there are different methods. SetIfExist() sets values if they
// are pending or already stored. SetIfPending() sets a value only if it is
// pending. SetEmptyIfPending() sets a value to its zero value if it is pending.
type PendingMap struct {
	mu      sync.RWMutex
	data    map[dskey.Key][]byte
	pending map[dskey.Key]chan struct{}
}

// New initializes a pendingDict.
func New() *PendingMap {
	return &PendingMap{
		data:    make(map[dskey.Key][]byte),
		pending: make(map[dskey.Key]chan struct{}),
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
// Only waits for keys that are pending when the function starts. If a key gets
// pending later (or switches between pending and not existing), the function
// returns an error.
//
// Possible Errors: context.Canceled, context.DeadlineExeeded or ErrNotExist
func (pm *PendingMap) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	if err := pm.waitForPending(ctx, keys); err != nil {
		return nil, err
	}

	out := make(map[dskey.Key][]byte, len(keys))
	err := pm.reading(func() error {
		for _, k := range keys {
			v, ok := pm.data[k]
			if !ok {
				return ErrNotExist
			}
			out[k] = v
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return out, nil
}

// waitForPending blocks until all the given keys are not pending anymore.
//
// Expects, that all keys are either pending or in the data. It is not allowed,
// that a key is not pending when this starts and gets pending whil it runs.
//
// Possible Errors: context.Canceled or context.DeadlineExeeded
func (pm *PendingMap) waitForPending(ctx context.Context, keys []dskey.Key) error {
	hasPendingKey := false
	pm.reading(func() error {
		for _, key := range keys {
			if _, ok := pm.pending[key]; ok {
				hasPendingKey = true
				return nil
			}
		}
		return nil
	})
	if !hasPendingKey {
		return nil
	}

	for _, k := range keys {
		var pending chan struct{}
		pm.reading(func() error {
			pending = pm.pending[k]
			return nil
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

// MarkPending marks one or more keys as pending.
//
// Skips keys that are already pending or are already in the map.
//
// Returns all keys that where marked as pending (did not exist).
func (pm *PendingMap) MarkPending(keys ...dskey.Key) []dskey.Key {
	var needMark []dskey.Key
	pm.reading(func() error {
		for _, key := range keys {
			if _, inStore := pm.data[key]; inStore {
				continue
			}
			if _, isPending := pm.pending[key]; isPending {
				continue
			}

			needMark = append(needMark, key)
		}
		return nil
	})

	if len(needMark) == 0 {
		return nil
	}

	marked := make([]dskey.Key, 0, len(needMark))
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for _, key := range needMark {
		if _, ok := pm.pending[key]; ok {
			// It can happen, that another caller has already set the key.
			continue
		}

		if _, inStore := pm.data[key]; inStore {
			// The other caller has already the data
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
func (pm *PendingMap) UnMarkPending(keys ...dskey.Key) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

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

// SetIfPendingOrExists updates values, but only if the key already exists or is pending.
//
// If the key is pending, it is unmarked and all listeners are informed.
func (pm *PendingMap) SetIfPendingOrExists(data map[dskey.Key][]byte) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for key, value := range data {
		pending := pm.pending[key]
		_, exists := pm.data[key]

		if pending == nil && !exists {
			continue
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
func (pm *PendingMap) SetIfPending(data map[dskey.Key][]byte) {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	for key, value := range data {
		if pending, isPending := pm.pending[key]; isPending {
			pm.data[key] = value
			close(pending)
			delete(pm.pending, key)
		}
	}
}

// Reset removes all data from PendingMap
func (pm *PendingMap) Reset() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.data = make(map[dskey.Key][]byte)
	pm.pending = make(map[dskey.Key]chan struct{})
}

// Len returns the amout of keys in the pending map.
func (pm *PendingMap) Len() int {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return len(pm.data)
}

func (pm *PendingMap) reading(cmd func() error) error {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return cmd()
}

// Size returns the size of all values in the cache in bytes.
func (pm *PendingMap) Size() int {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var size int
	for _, v := range pm.data {
		size += len(v)
	}
	return size
}
