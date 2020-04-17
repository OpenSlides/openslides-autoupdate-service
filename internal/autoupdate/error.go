package autoupdate

import "fmt"

// ValueError in returned by the ider, when the value of a key has not the expected format.
type ValueError struct {
	key string
}

func (e ValueError) Error() string {
	return fmt.Sprintf("invalid value in key %s", e.key)
}

// Type returns the name of the error.
func (e ValueError) Type() string {
	return "ValueError"
}
