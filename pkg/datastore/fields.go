package datastore

//go:generate  sh -c "go run gen_fields/main.go > field_methods.go"

// Fields contains methods for all existing fields.
//
// Can only be used from Fetcher.Field().
type Fields struct {
	fetch *Fetcher
}
