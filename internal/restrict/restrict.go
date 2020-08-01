// Package restrict holds the Restricter obect that filters data for a specific
// user.
package restrict

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// Restricter implements the autoupdate.Restricter interface.
type Restricter struct {
	perm      Permission
	datastore Datastore

	checksMu sync.RWMutex
	checks   map[string]Checker
}

// New creates an initialized Restricter.
func New(perm Permission, db Datastore, checker map[string]Checker) *Restricter {
	// Initialize structured fields
	for k, check := range checker {
		sf, ok := check.(*structuredField)
		if !ok {
			continue
		}

		// TODO
		_ = sf
		_ = k
	}

	r := &Restricter{
		perm:      perm,
		checks:    checker,
		datastore: db,
	}

	db.RegisterChangeListener(r.structuredFieldUpdate)

	return r
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

		checker, ok := r.checks[fqfieldToModelField(k)]
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

func (r *Restricter) structuredFieldUpdate(data map[string]json.RawMessage) error {
	r.checksMu.Lock()
	defer r.checksMu.Unlock()

	for k, v := range data {
		if !strings.Contains(k, "$") {
			// Not a template field.
			continue
		}

		checkerKey := fqfieldToModelField(k)

		check, ok := r.checks[checkerKey]
		if !ok {
			return fmt.Errorf("unknown template field %s", checkerKey)
		}

		sf, ok := check.(*structuredField)
		if !ok {
			return fmt.Errorf("key %s is not a structured field", checkerKey)
		}

		// Delete old checks.
		for _, field := range sf.fields {
			delete(r.checks, field)
		}

		var replacements []string
		if err := json.Unmarshal(v, &replacements); err != nil {
			return fmt.Errorf("decoding template field %s: %w", k, err)
		}

		// Create new checks.
		for _, field := range structuredKeys(checkerKey, replacements) {
			r.checks[field] = sf.checker
		}
	}
	return nil
}

func structuredKeys(key string, replecments []string) []string {
	replaced := make([]string, len(replecments))
	for i, r := range replecments {
		replaced[i] = strings.Replace(key, "$", r, 1)
	}
	return replaced
}

func fqfieldToModelField(fqfied string) string {
	t := strings.Split(fqfied, "/")
	return t[0] + "/" + t[2]
}
