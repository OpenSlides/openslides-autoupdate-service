// Package test is a test helper for the other packages.
package test

// Str builds a slice of strings.
func Str(keys ...string) []string {
	return keys
}

// CmpSlice checks, if slice a and b holds the same values.
func CmpSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

type closingError struct{}

func (e closingError) Closing()      {}
func (e closingError) Error() string { return "closing" }
