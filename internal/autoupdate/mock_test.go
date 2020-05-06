package autoupdate_test

import (
	"context"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

type mockKeysBuilder struct {
	keys []string
}

func (m mockKeysBuilder) Update([]string) error {
	return nil
}

func (m mockKeysBuilder) Keys() []string {
	return m.keys
}

func getConnection() (connection *autoupdate.Connection, datastore *test.MockDatastore, disconnect func(), close func()) {
	datastore = test.NewMockDatastore()
	s := autoupdate.New(datastore, new(test.MockRestricter))
	ctx, cancel := context.WithCancel(context.Background())
	kb := mockKeysBuilder{keys: test.Str("user/1/name")}
	c := s.Connect(ctx, 1, kb)

	return c, datastore, cancel, func() {
		cancel()
		s.Close()
		datastore.Close()
	}
}
