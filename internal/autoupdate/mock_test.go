package autoupdate_test

import (
	"context"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

type mockKeysBuilder struct {
	keys []string
}

func (m mockKeysBuilder) Update(context.Context) error {
	return nil
}

func (m mockKeysBuilder) Keys() []string {
	return m.keys
}

func getConnection(closed <-chan struct{}) (*autoupdate.Connection, *test.MockDatastore) {
	datastore := new(test.MockDatastore)
	s := autoupdate.New(datastore, new(test.MockRestricter), closed)
	kb := mockKeysBuilder{keys: test.Str("user/1/name")}
	c := s.Connect(1, kb)

	return c, datastore
}
