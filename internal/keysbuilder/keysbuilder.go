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
	ctx   context.Context
	ider  IDer
	bodys []body
	cache *cache
	keys  []string
}

// newBuilder creates a new Builder instance.
func newBuilder(ctx context.Context, ider IDer, bodys ...body) (*Builder, error) {
	b := &Builder{
		ctx:   ctx,
		ider:  ider,
		bodys: bodys,
		cache: newCache(),
	}
	if err := b.genKeys(); err != nil {
		return nil, err
	}
	return b, nil
}

// Update triggers a keyupdate
func (b *Builder) Update(keys []string) error {
	b.cache.delete(keys)
	return b.genKeys()
}

// Keys returns the keys
func (b *Builder) Keys() []string {
	return b.keys
}

func (b *Builder) genKeys() error {
	var wg sync.WaitGroup
	keys := make(chan string, 1)
	errs := make(chan error, 1)
	ctx, cancel := context.WithCancel(b.ctx)
	defer cancel()

	for _, request := range b.bodys {
		wg.Add(1)
		go func(request body) {
			request.build(ctx, b, keys, errs)
			wg.Done()
		}(request)
	}

	go func() {
		wg.Wait()
		close(keys)
	}()

	b.keys = b.keys[:0]
	var err error
	for {
		select {
		case key, ok := <-keys:
			if !ok {
				return err
			}
			b.keys = append(b.keys, key)
		case err = <-errs:
			cancel()
		}
	}
}

func buildKey(collection string, id int, field string) string {
	return collection + keySep + strconv.Itoa(id) + keySep + field
}
