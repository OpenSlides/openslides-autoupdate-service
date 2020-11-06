package core

type clientError struct {
	msg string
}

func (e clientError) Type() string {
	return "ClientError"
}

func (e clientError) Error() string {
	return e.msg
}
