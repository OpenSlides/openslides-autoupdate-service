// Package restrict holds the Restricter obect that filters data for a specific
// user.
package restrict

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Restricter implements the autoupdate.Restricter interface.
type Restricter struct {
	perm   Permission
	checks map[string]Checker
}

// New creates an initialized Restricter.
func New(perm Permission, checker map[string]Checker) *Restricter {
	return &Restricter{
		perm:   perm,
		checks: checker,
	}
}

// Restrict filters and manipulates the given data for the user with the given
// uid.
//
// It is not allowed to manipulate a value in the dict. A value can only be
// replaced with a new value. If the user does not have the permission to see
// one key, it is not allowed to remove that key, the value has to be set to
// nil.
func (r *Restricter) Restrict(uid int, data map[string]json.RawMessage) error {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	allowed, err := r.perm.CheckFQFields(uid, keys)
	if err != nil {
		return fmt.Errorf("check permissions: %w", err)
	}

	for k, v := range data {
		if !allowed[k] {
			data[k] = nil
			continue
		}

		// Get key without model id
		t := strings.Split(k, "/")
		f := t[0] + "/" + t[2]

		checker, ok := r.checks[f]
		if !ok {
			continue
		}

		nv, err := checker.Check(uid, k, v)
		if err != nil {
			return fmt.Errorf("checker for key %s: %w", k, err)
		}
		data[k] = nv
	}
	return nil
}
