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
// Has to be created with keysbuilder.FromJSON or keysbuilder.ManyFromJSON.
type Builder struct {
	ctx   context.Context
	ider  IDer
	bodys []body
	cache *cache
	keys  []string
}

// newBuilder creates a new Builder instance
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
	kc := make(chan string, 1)
	ec := make(chan error, 1)
	ctx, cancel := context.WithCancel(b.ctx)
	defer cancel()

	for _, kr := range b.bodys {
		wg.Add(1)
		go func(kr body) {
			b.run(ctx, kr.IDs, kr.fields, kc, ec)
			wg.Done()
		}(kr)
	}

	go func() {
		wg.Wait()
		close(kc)
	}()

	b.keys = make([]string, 0)
	var err error
	for {
		select {
		case key, ok := <-kc:
			if !ok {
				return err
			}
			b.keys = append(b.keys, key)
		case err = <-ec:
			cancel()
		}
	}
}

func (b *Builder) run(ctx context.Context, ids []int, fd fields, kc chan<- string, ec chan<- error) {
	var wg sync.WaitGroup
	for _, id := range ids {
		for field, ifd := range fd.Names {
			key := buildKey(fd.Collection, id, field)
			kc <- key
			if ifd.null() {
				// field is not a reference
				continue
			}

			wg.Add(1)
			go func(name string, ifd fields) {
				defer wg.Done()
				ids := b.cache.getOrSet(name, func() []int {
					ids, err := b.ider.IDs(ctx, name)
					if err != nil {
						ec <- err
						return nil
					}
					return ids
				})

				b.run(ctx, ids, ifd, kc, ec)
			}(key, ifd)
		}
	}
	wg.Wait()
}

func buildKey(collection string, id int, field string) string {
	return collection + keySep + strconv.Itoa(id) + keySep + field
}
