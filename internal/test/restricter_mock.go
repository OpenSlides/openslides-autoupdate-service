package test

import "encoding/json"

// MockRestricter implements the restricter interface.
type MockRestricter struct{}

// Restrict does currently nothing.
func (r *MockRestricter) Restrict(uid int, data map[string]json.RawMessage) error {
	return nil
}
