package autoupdate_test

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

func getConnection() (func(context.Context) (map[dskey.Key][]byte, error), *dsmock.MockDatastore, func(context.Context, func(error))) {
	datastore, dsBackground := dsmock.NewMockDatastore(dsmock.YAMLData(`---
	user/1/name: Hello World
	`))

	lookup := environment.ForTests{}
	s, _, _ := autoupdate.New(lookup, datastore, RestrictAllowed)
	kb, _ := keysbuilder.FromKeys(userNameKey.String())
	next, err := s.Connect(context.Background(), 1, kb)
	if err != nil {
		panic(err)
	}

	f, _ := next()

	return f, datastore, dsBackground
}

func blocking(f func()) bool {
	return blockingTime(10*time.Millisecond, f)
}

func blockingTime(wait time.Duration, f func()) bool {
	done := make(chan struct{})
	go func() {
		f()
		close(done)
	}()

	timer := time.NewTimer(wait)
	defer timer.Stop()
	select {
	case <-done:
		return false
	case <-timer.C:
		return true
	}
}

// RestrictAllowed is a restricter that allows everything
func RestrictAllowed(getter datastore.Getter, uid int) datastore.Getter {
	return mockRestricter{getter, true}
}

// RestrictNotAllowed is a restricter that removes everythin
func RestrictNotAllowed(getter datastore.Getter, uid int) datastore.Getter {
	return mockRestricter{getter, false}
}

type mockRestricter struct {
	getter datastore.Getter
	allow  bool
}

func (r mockRestricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("getting data: %w", err)
	}

	if r.allow {
		return data, nil
	}

	for k := range data {
		data[k] = nil
	}
	return data, nil
}
