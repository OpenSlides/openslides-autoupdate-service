package http

import "fmt"

type invalidRequestError struct {
	err error
}

func (e invalidRequestError) Error() string {
	return fmt.Sprintf("Invalid request: %v", e.err)
}

func (e invalidRequestError) Type() string {
	return "invalid_request"
}
