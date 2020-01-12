// Package keysbuilder ...
package keysbuilder

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/openslides/openslides-autoupdate-service/internal/keysrequest"
)

const keySep = "/"

// Builder ...
type Builder struct {
	user         int
	restricter   Restricter
	keysRequests []keysrequest.KeysRequest
	cache        *cache
	keys         []string
}

// New creates a new Builder instance
func New(user int, restricter Restricter, keysRequests ...keysrequest.KeysRequest) (*Builder, error) {
	b := &Builder{
		user:         user,
		restricter:   restricter,
		keysRequests: keysRequests,
		cache:        newCache(),
	}

	// TODO: Run in parallel
	for idx, kr := range keysRequests {
		// Save the ids in the cache
		_, err := b.cache.get(string(idx), func() ([]int, error) {
			// Get ids from the request or all ids of the collection
			ids := kr.IDs
			if ids == nil {
				var err error
				ids, err = b.restricter.IDsFromCollection(context.TODO(), b.user, kr.MeetingID, kr.Collection)
				if err != nil {
					return nil, fmt.Errorf("can not get all ids for collection \"%s\": %w", kr.Collection, err)
				}
			}
			return ids, nil
		})
		if err != nil {
			return nil, err
		}
	}
	if err := b.genKeys(); err != nil {
		return nil, err
	}
	return b, nil
}

// Update triggers a keyupdate
func (b *Builder) Update(keys []string) error {
	b.cache.mu.Lock()
	for _, key := range keys {
		if !(strings.HasSuffix(key, "_id") || strings.HasSuffix(key, "_ids")) {
			continue
		}
		delete(b.cache.data, key)
	}
	b.cache.mu.Unlock()
	return b.genKeys()
}

// Keys returns the keys
func (b *Builder) Keys() []string {
	return b.keys
}

func (b *Builder) genKeys() error {
	b.keys = make([]string, 0)
	for idx, kr := range b.keysRequests {
		keys, err := b.run(string(idx), kr.FieldDescription, kr.MeetingID)
		if err != nil {
			return err
		}
		b.keys = append(b.keys, keys...)
	}
	return nil
}

func (b *Builder) run(name string, fd keysrequest.FieldDescription, meeting int) ([]string, error) {
	ids, err := b.cache.get(name, func() ([]int, error) {
		return b.restricter.IDsFromKey(context.TODO(), b.user, meeting, name)
	})
	if err != nil {
		return nil, err
	}

	done := make(chan struct{})
	kc := make(chan string)
	go func() {
		var wg sync.WaitGroup
		for _, id := range ids {
			for field, ifd := range fd.Fields {
				key := buildKey(fd.Collection, id, field)
				kc <- key
				if ifd.Null() {
					// field is not a reference
					continue
				}

				// TODO handle error
				wg.Add(1)
				go func(name string, ifd keysrequest.FieldDescription) {
					keys, _ := b.run(name, ifd, meeting)
					for _, key := range keys {
						kc <- key
					}
					wg.Done()
				}(key, ifd)
			}
		}
		wg.Wait()
		close(done)
	}()
	keys := make([]string, 0)
	for {
		select {
		case key := <-kc:
			keys = append(keys, key)
		case <-done:
			return keys, nil
		}
	}
}

func buildKey(collection string, id int, field string) string {
	return collection + keySep + strconv.Itoa(id) + keySep + field
}
