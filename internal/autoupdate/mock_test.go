package autoupdate_test

import (
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func getConnection(closed <-chan struct{}) (*autoupdate.Connection, *test.MockDatastore) {
	datastore := test.NewMockDatastore(map[string]string{
		"user/1/name": `"Hello World"`,
	})
	s := autoupdate.New(datastore, new(test.MockRestricter), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	c := s.Connect(1, kb)

	return c, datastore
}
