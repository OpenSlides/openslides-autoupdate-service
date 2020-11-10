package http

import (
	"fmt"
)

type jsonError struct {
	msg string
	err error
}

func (e jsonError) Type() string {
	return "JSONError"
}

func (e jsonError) Error() string {
	return fmt.Sprintf("%s: %v", e.msg, e.err)
}

func (e jsonError) Unwrap() error {
	return e.err
}
