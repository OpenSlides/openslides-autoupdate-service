// Package restrict holds the Restricter obect that filters data for a specific
// user.
package restrict

import "encoding/json"

// Restricter implements the autoupdate.Restricter interface.
type Restricter struct{}

// Restrict filters and manipulates the given data for the user with the given
// uid.
//
// It is not allowed to manipulate a value in the dict. A value can only be
// replaced with a new value. If the user does not have the permission to see
// one key, it is not allowed to remove that key, the value has to be set to
// nil.
func (r *Restricter) Restrict(uid int, data map[string]json.RawMessage) {

}
