package autoupdate_test

import (
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func getConnection(closed <-chan struct{}) (*autoupdate.Connection, *test.MockDatastore) {
	datastore := new(test.MockDatastore)
	s := autoupdate.New(datastore, new(test.MockRestricter), mockUserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	c := s.Connect(1, kb)

	return c, datastore
}

type mockUserUpdater struct {
	userIDs []int
}

func (u mockUserUpdater) AdditionalUpdate(updated map[string]string) ([]int, error) {
	return u.userIDs, nil
}
