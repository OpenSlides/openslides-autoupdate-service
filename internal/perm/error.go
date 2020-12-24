package perm

import "fmt"

// NotAllowedf is an error that sends a message to the client that indicates,
// that the user has not the required permissions.
func NotAllowedf(format string, a ...interface{}) error {
	return notAllowedError{fmt.Sprintf(format, a...)}
}

type notAllowedError struct {
	msg string
}

func (e notAllowedError) Type() string {
	return "ClientError"
}

func (e notAllowedError) Error() string {
	return e.msg
}
