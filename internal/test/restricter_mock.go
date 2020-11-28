package test

import (
	"context"
	"encoding/json"
)

// MockRestricter implements the restricter interface.
type MockRestricter struct{}

// Restrict does currently nothing.
func (r *MockRestricter) Restrict(ctx context.Context, uid int, data map[string]json.RawMessage) error {
	return nil
}
