package test

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

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

func (r mockRestricter) Get(ctx context.Context, keys ...datastore.Key) (map[datastore.Key][]byte, error) {
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
