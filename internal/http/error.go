package http

import "fmt"

// noStatusCodeError helps the errorHandler do decide, if an status code can be
// set.
type noStatusCodeError struct {
	wrapped error
}

func (e noStatusCodeError) Error() string {
	return e.wrapped.Error()
}

type invalidRequestError struct {
	err error
}

func (e invalidRequestError) Error() string {
	return fmt.Sprintf("Invalid request: %v", e.err)
}

func (e invalidRequestError) Type() string {
	return "invalid_request"
}
