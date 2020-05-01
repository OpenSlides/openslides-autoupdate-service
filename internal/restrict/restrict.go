package restrict

// Restricter implements the autoupdate.Restricter interface.
type Restricter struct{}

// Restrict filters and manipulates the given data for the user with the given
// uid.
func (r *Restricter) Restrict(uid int, data map[string]string) {

}
