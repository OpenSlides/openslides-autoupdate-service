package test

import (
	"context"
	"encoding/json"
)

// UserUpdater implements implements autoupdater UserUpdater.
type UserUpdater struct {
	UserIDs []int
}

// AdditionalUpdate returns the userIDs.
func (u UserUpdater) AdditionalUpdate(ctx context.Context, updated map[string]json.RawMessage) ([]int, error) {
	return u.UserIDs, nil
}
