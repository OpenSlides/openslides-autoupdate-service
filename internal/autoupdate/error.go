package autoupdate

import "fmt"

// ErrValue in returned by the ider, when the value of a key has not the expected format.
type ErrValue struct {
	key string
}

func (e ErrValue) Error() string {
	return fmt.Sprintf("Invalid value in key %s", e.key)
}

// Type returns the name of the error.
func (e ErrValue) Type() string {
	return "ValueError"
}
