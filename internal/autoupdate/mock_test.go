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

func getConnection() (connection *autoupdate.Connection, datastore *test.MockDatastore, close func()) {
	datastore = test.NewMockDatastore()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	kb := mockKeysBuilder{keys: test.Str("user/1/name")}
	c := s.Connect(1, kb, 0)

	return c, datastore, func() {
		s.Close()
		datastore.Close()
	}
}
