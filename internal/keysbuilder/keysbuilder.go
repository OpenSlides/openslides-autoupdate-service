// Package keysbuilder holds a datastructure to get and update requested keys.
package keysbuilder

import (
	"context"
	"strconv"
	"sync"
)

const keySep = "/"

// Builder builds the keys.
//
// Has to be created with keysbuilder.FromJSON() or keysbuilder.ManyFromJSON().
type Builder struct {
	ctx    context.Context
	ider   IDer
	bodies []body
	keys   []string
}

// newBuilder creates a new Builder instance from one or more bodies.
func newBuilder(ctx context.Context, ider IDer, bodys ...body) (*Builder, error) {
	b := &Builder{
		ctx:    ctx,
		ider:   ider,
		bodies: bodys,
	}
	if err := b.genKeys(); err != nil {
		return nil, err
	}
	return b, nil
}

// Update triggers a key update.
func (b *Builder) Update(keys []string) error {
	return b.genKeys()
}

// Keys returns the keys.
//
// This method reads the values from a cache. Therefore the method returns in
// constant time.
func (b *Builder) Keys() []string {
	return b.keys
}

// genKeys generates the list of keys, that can be requested with the Keys()
// method. It travels the KeysRequests object like a tree. Each branch is
// processed concurrently.
func (b *Builder) genKeys() error {
	var wg sync.WaitGroup
	keys := make(chan string, 1)
	errs := make(chan error)
	ctx, cancel := context.WithCancel(b.ctx)
	defer cancel()

	// Go though all bodies at the same time.
	for _, request := range b.bodies {
		wg.Add(1)
		go func(request body) {
			request.build(ctx, b, keys, errs)
			wg.Done()
		}(request)
	}

	// Close the keys channel as soon as all bodies are traveled.
	go func() {
		wg.Wait()
		close(keys)
	}()

	// Clears the keys slice without reallocating memory.
	b.keys = b.keys[:0]

	var err error
	for {
		select {
		case key, ok := <-keys:
			if !ok {
				// The keys channel was closed. This happens when everything is
				// done or when an error occurs.
				return err
			}
			b.keys = append(b.keys, key)
		case err = <-errs:
			cancel()
		}
	}
}

// buildGenericKey returns a valid key when the collection and id are already
// together.
//
// buildGenericKey("motion/5", "title") -> "motion/5/title".
func buildGenericKey(collectionID string, field string) string {
	return collectionID + keySep + field
}

func buildCollectionID(collection string, id int) string {
	return collection + keySep + strconv.Itoa(id)
}
