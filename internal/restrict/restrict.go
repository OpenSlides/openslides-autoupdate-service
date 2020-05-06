// Package restrict holds the Restricter obect that filters data for a specific
// user.
package restrict

import "encoding/json"

// Restricter implements the autoupdate.Restricter interface.
type Restricter struct{}

// Restrict filters and manipulates the given data for the user with the given
// uid.
//
// It is not allowed to manipulate a value in the dict. Either delete a
// key-value pair or replace it with new data.
func (r *Restricter) Restrict(uid int, data map[string]json.RawMessage) {

}
