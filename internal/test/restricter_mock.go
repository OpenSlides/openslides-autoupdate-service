package test

import (
	"context"
	"encoding/json"
)

// MockRestricter implements the restricter interface.
type MockRestricter struct {
	allowAll    bool
	allowedKeys map[string]bool
}

// RestrictAllowed creates a Restricter that allows everything.
func RestrictAllowed() *MockRestricter {
	return &MockRestricter{
		allowAll: true,
	}
}

// RestrictDenied create a Restricter, that disallowes everything.
func RestrictDenied() *MockRestricter {
	return new(MockRestricter)
}

// RestrictOnlyAllow creates a Restricter that allows only some keys.
func RestrictOnlyAllow(keys []string) *MockRestricter {
	set := make(map[string]bool)
	for _, k := range keys {
		set[k] = true
	}
	return &MockRestricter{
		allowedKeys: set,
	}
}

// Restrict does currently nothing.
func (r *MockRestricter) Restrict(ctx context.Context, uid int, data map[string]json.RawMessage) error {
	if r.allowAll {
		return nil
	}

	for k := range data {
		if r.allowedKeys == nil || !r.allowedKeys[k] {
			delete(data, k)
		}
	}
	return nil
}
