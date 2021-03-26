package test

import (
	"context"
	"encoding/json"
)

// MockRestricter implements the restricter interface.
type MockRestricter struct {
	NoPermission bool
}

// Restrict does currently nothing.
func (r *MockRestricter) Restrict(ctx context.Context, uid int, data map[string]json.RawMessage) error {
	if r.NoPermission {
		for k := range data {
			delete(data, k)
		}
	}
	return nil
}
