package allowed

import (
	"fmt"
)

func NotAllowed(message string) error {
	return notAllowedError{message}
}

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
