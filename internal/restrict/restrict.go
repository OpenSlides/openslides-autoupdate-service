// Package restrict holds the Restricter obect that filters data for a specific
// user.
package restrict

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// Restricter implements the autoupdate.Restricter interface.
type Restricter struct {
	permer Permissioner
	checks map[string]Checker
}

// New creates an initialized Restricter.
func New(permer Permissioner, checker map[string]Checker) *Restricter {
	r := &Restricter{
		permer: permer,
		checks: checker,
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
func (r *Restricter) Restrict(ctx context.Context, uid int, data map[string]json.RawMessage) error {
	keys := make([]string, 0, len(data))
	for k, v := range data {
		if v == nil {
			// If the value is nil, there is no need to check it.
			continue
		}
		keys = append(keys, k)
	}
	allowed, err := r.permer.RestrictFQFields(ctx, uid, keys)
	if err != nil {
		return fmt.Errorf("check permissions: %w", err)
	}

	for k, v := range data {
		if v == nil {
			continue
		}

		if !allowed[k] {
			data[k] = nil
			continue
		}

		checker, ok := r.checks[checkerIndex(k)]
		if !ok {
			continue
		}

		nv, err := checker.Check(ctx, uid, k, v)
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

// checkerIndex returns the index of the checker list. This is the modelField
// (fqfield without middle part) or normal fields and template fields
// (foo/1/prefix_$_bar, foo/1/prefix_$). For fields that are derivated from a
// template field, this is only the prefix of the field.
func checkerIndex(fqfield string) string {
	t := strings.Split(fqfield, "/")
	modelField := t[0] + "/" + t[2]

	i := strings.IndexByte(modelField, '$')
	if i < 0 || i == len(modelField)-1 || modelField[i+1] == '_' {
		// Normal field or $ at the end or $_
		return modelField
	}

	return modelField[:i]
}
