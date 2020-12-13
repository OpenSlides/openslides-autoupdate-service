package core

import "fmt"

type clientError struct {
	msg string
}

func (e clientError) Type() string {
	return "ClientError"
}

func (e clientError) Error() string {
	return e.msg
}

type isAllowedError struct {
	err   error
	index int
	name  string
}

func (e isAllowedError) Error() string {
	return fmt.Sprintf("%s: %v", e.name, e.err)
}

func (e isAllowedError) Unwrap() error {
	return e.err
}

func (e isAllowedError) Index() int {
	return e.index
}
