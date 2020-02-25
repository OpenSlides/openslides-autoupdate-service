// Package keysbuilder holds a datastructure to get all requested keys from  keysrequests.
package keysbuilder

import (
	"context"
	"strconv"
	"sync"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
)

const keySep = "/"

// Builder builds the keys from a list of keysrequest.
type Builder struct {
	ctx          context.Context
	ider         IDer
	keysRequests []keysrequest.Body
	cache        *cache
	keys         []string
}

// New creates a new Builder instance
func New(ctx context.Context, ider IDer, keysRequests ...keysrequest.Body) (*Builder, error) {
	b := &Builder{
		ctx:          ctx,
		ider:         ider,
		keysRequests: keysRequests,
		cache:        newCache(),
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

	for _, kr := range b.keysRequests {
		wg.Add(1)
		go func(kr keysrequest.Body) {
			b.run(ctx, kr.IDs, kr.Fields, kc, ec)
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

func (b *Builder) run(ctx context.Context, ids []int, fd keysrequest.Fields, kc chan<- string, ec chan<- error) {
	var wg sync.WaitGroup
	for _, id := range ids {
		for field, ifd := range fd.Names {
			key := buildKey(fd.Collection, id, field)
			kc <- key
			if ifd.Null() {
				// field is not a reference
				continue
			}

			wg.Add(1)
			go func(name string, ifd keysrequest.Fields) {
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
