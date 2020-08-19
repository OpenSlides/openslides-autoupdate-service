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
	perm             Permission
	checks           map[string]Checker
	structuredFields []*structuredField
}

// New creates an initialized Restricter.
func New(perm Permission, checker map[string]Checker) *Restricter {
	r := &Restricter{
		perm:   perm,
		checks: checker,
	}

	for _, c := range checker {
		if s, ok := c.(*structuredField); ok {
			r.structuredFields = append(r.structuredFields, s)
		}
	}
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

		modelField := fqfieldToModelField(k)
		checker, ok := r.checks[modelField]
		if !ok {
			for _, sf := range r.structuredFields {
				if sf.Match(modelField) {
					checker = sf.checker
					break
				}
			}
			if checker == nil {
				// Not a check and not a structured field.
				continue
			}
		}

		nv, err := checker.Check(uid, k, v)
		if err != nil {
			return fmt.Errorf("checker for key %s: %w", k, err)
		}
		data[k] = nv
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

func fqfieldToModelField(fqfield string) string {
	t := strings.Split(fqfield, "/")
	return t[0] + "/" + t[2]
}
