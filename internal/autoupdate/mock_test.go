package autoupdate_test

import (
	"context"
	"encoding/json"

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

func (u mockUserUpdater) AdditionalUpdate(ctx context.Context, updated map[string]json.RawMessage) ([]int, error) {
	return u.userIDs, nil
}
