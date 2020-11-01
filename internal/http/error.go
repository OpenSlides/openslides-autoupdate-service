package http

type jsonError struct {
	msg string
	err error
}

func (e jsonError) Type() string {
	return "JSONError"
}

func (e jsonError) Error() string {
	return e.msg
}

func (e jsonError) Unwrap() error {
	return e.err
}
