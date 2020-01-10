// Package keysbuilder ...
package keysbuilder

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/keysrequest"
)

const keySep = "/"

// Builder ...
type Builder struct {
	user        int
	restricter  Restricter
	keysRequest keysrequest.KeysRequest
	cache       map[string][]int
	keys        map[string]bool
}

// New creates a new Builder instance
func New(user int, restricter Restricter, keysRequest keysrequest.KeysRequest) (*Builder, error) {
	b := &Builder{
		user:        user,
		restricter:  restricter,
		keysRequest: keysRequest,
		cache:       make(map[string][]int),
	}
	// Get ids from the request or all ids of the collection
	ids := b.keysRequest.IDs
	if ids == nil {
		var err error
		ids, err = b.restricter.IDsFromCollection(context.TODO(), b.user, b.keysRequest.MeetingID, b.keysRequest.Collection)
		if err != nil {
			return nil, fmt.Errorf("can not get all ids for collection \"%s\": %w", b.keysRequest.Collection, err)
		}
	}

	// Save the ids in the cache
	b.cache[b.keysRequest.Collection] = ids
	if err := b.genKeys(); err != nil {
		return nil, err
	}
	return b, nil
}

// Update triggers a keyupdate
func (b *Builder) Update(keys []string) error {
	for _, key := range keys {
		if !(strings.HasSuffix(key, "_id") || strings.HasSuffix(key, "_ids")) {
			continue
		}
		delete(b.cache, key)
	}
	return b.genKeys()
}

// Keys returns the keys
func (b *Builder) Keys() []string {
	out := make([]string, 0, len(b.keys))
	for key := range b.keys {
		out = append(out, key)
	}
	return out
}

func (b *Builder) genKeys() error {
	b.keys = make(map[string]bool)
	return b.run(b.keysRequest.Collection, b.keysRequest.FieldDescription)
}

func (b *Builder) run(name string, fd keysrequest.FieldDescription) error {
	ids, ok := b.cache[name]
	if !ok {
		var err error
		ids, err = b.restricter.IDsFromKey(context.TODO(), b.user, b.keysRequest.MeetingID, name)
		if err != nil {
			return err
		}
		b.cache[name] = ids
	}
	for _, id := range ids {
		for field, ifd := range fd.Fields {
			key := buildKey(fd.Collection, id, field)
			b.keys[key] = true
			if ifd.Null() {
				// field is not a reference
				continue
			}

			if err := b.run(key, ifd); err != nil {
				return err
			}
		}
	}
	return nil
}

func buildKey(collection string, id int, field string) string {
	return collection + keySep + strconv.Itoa(id) + keySep + field
}
