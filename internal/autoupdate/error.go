package autoupdate

import (
	"fmt"
	"reflect"
)

// ValueError in returned by autoupdate.Value(), when the value of a key has not the
// expected format.
type ValueError struct {
	key        string
	gotType    string
	expectType reflect.Type
	err        error
}

func (e ValueError) Error() string {
	return fmt.Sprintf("invalid value in key %s. Got %s, expected %s", e.key, e.gotType, e.expectType)
}

// Type returns the name of the error.
func (e ValueError) Type() string {
	return "ValueError"
}

func (e ValueError) Unwrap() error {
	return e.err
}

// NotExistError is returned by autoupdate.Value, when the requested key does
// not exist or the user has not the permission to see it.
type NotExistError struct {
	Key string
}

func (e NotExistError) Error() string {
	return fmt.Sprintf("the key %s does not exist", e.Key)
}

// Type returns the name of the error.
func (e NotExistError) Type() string {
	return "NotExistError"
}

// KeyDoesNotExist returns true.
func (e NotExistError) KeyDoesNotExist() bool {
	return true
}
