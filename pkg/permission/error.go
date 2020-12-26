package permission

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

type indexError struct {
	err   error
	index int
	name  string
}

func (e indexError) Error() string {
	return fmt.Sprintf("%s: %v", e.name, e.err)
}

func (e indexError) Unwrap() error {
	return e.err
}

func (e indexError) Index() int {
	return e.index
}
