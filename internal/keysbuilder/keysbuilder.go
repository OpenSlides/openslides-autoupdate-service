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
	ids         []int
	restricter  Restricter
	keysRequest keysrequest.KeysRequest
	n           *node
}

// New creates a new Builder instance
func New(user int, restricter Restricter, keysRequest keysrequest.KeysRequest) (*Builder, error) {
	b := &Builder{
		user:        user,
		restricter:  restricter,
		keysRequest: keysRequest,
	}
	b.ids = b.keysRequest.IDs
	if b.ids == nil {
		var err error
		b.ids, err = b.restricter.IDsFromCollection(context.TODO(), b.user, b.keysRequest.MeetingID, b.keysRequest.Collection)
		if err != nil {
			return nil, fmt.Errorf("can not get all ids for collection \"%s\": %w", b.keysRequest.Collection, err)
		}
	}
	b.n = &node{fd: b.keysRequest.FieldDescription}
	if err := b.run(b.n); err != nil {
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
		if err := b.reset(b.n, key); err != nil {
			return err
		}
	}
	return nil
}

// Keys returns the keys
func (b *Builder) Keys() []string {
	return b.n.run()
}

func (b *Builder) reset(n *node, key string) error {
	for _, field := range n.fields {
		if field.fd.Null() {
			// No need to reset non relation fields
			continue
		}
		var err error
		// Either test sub fields or reset node
		if key != field.name {
			err = b.reset(field, key)
		} else {
			err = b.run(field)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
func (b *Builder) run(n *node) error {
	n.fields = []*node{}
	ids := b.ids
	if n.name != "" {
		var err error
		ids, err = b.restricter.IDsFromKey(context.TODO(), b.user, b.keysRequest.MeetingID, n.name)
		if err != nil {
			return err
		}
	}
	for _, id := range ids {
		for field, fd := range n.fd.Fields {
			key := buildKey(n.fd.Collection, id, field)
			node := &node{name: key, fd: fd}
			n.fields = append(n.fields, node)
			if fd.Null() {
				// field is not a reference
				continue
			}

			if err := b.run(node); err != nil {
				return err
			}
		}
	}
	return nil
}

func buildKey(collection string, id int, field string) string {
	return collection + keySep + strconv.Itoa(id) + keySep + field
}
