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
	user        int
	restricter  Restricter
	keysRequest keysrequest.KeysRequest
	cache       *cache
	keys        []string
}

// New creates a new Builder instance
func New(user int, restricter Restricter, keysRequest keysrequest.KeysRequest) (*Builder, error) {
	b := &Builder{
		user:        user,
		restricter:  restricter,
		keysRequest: keysRequest,
		cache:       newCache(),
	}

	// Save the ids in the cache
	_, err := b.cache.get(b.keysRequest.Collection, func() ([]int, error) {
		// Get ids from the request or all ids of the collection
		ids := b.keysRequest.IDs
		if ids == nil {
			var err error
			ids, err = b.restricter.IDsFromCollection(context.TODO(), b.user, b.keysRequest.MeetingID, b.keysRequest.Collection)
			if err != nil {
				return nil, fmt.Errorf("can not get all ids for collection \"%s\": %w", b.keysRequest.Collection, err)
			}
		}
		return ids, nil
	})
	if err != nil {
		return nil, err
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
	keys, err := b.run(b.keysRequest.Collection, b.keysRequest.FieldDescription)
	if err != nil {
		return err
	}
	b.keys = keys
	return nil
}

func (b *Builder) run(name string, fd keysrequest.FieldDescription) ([]string, error) {
	ids, err := b.cache.get(name, func() ([]int, error) {
		return b.restricter.IDsFromKey(context.TODO(), b.user, b.keysRequest.MeetingID, name)
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
					keys, _ := b.run(name, ifd)
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
