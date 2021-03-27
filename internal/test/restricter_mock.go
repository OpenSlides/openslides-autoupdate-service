package test

import (
	"context"
	"encoding/json"
)

// MockRestricter implements the restricter interface.
type MockRestricter struct {
	denie bool

	Values map[string]string
}

// RestrictAllowed creates a Restricter that allows everything.
func RestrictAllowed() *MockRestricter {
	return &MockRestricter{
		denie: false,
	}
}

// RestrictDenied create a Restricter, that disallowes everything.
func RestrictDenied() *MockRestricter {
	return &MockRestricter{
		denie: true,
	}
}

// Restrict does currently nothing.
func (r *MockRestricter) Restrict(ctx context.Context, uid int, data map[string]json.RawMessage) error {
	if r.denie {
		for k := range data {
			delete(data, k)
		}
		return nil
	}

	if r.Values != nil {
		for k := range data {
			v, ok := r.Values[k]
			if ok {
				data[k] = []byte(v)
			}
		}
	}
	return nil
}
