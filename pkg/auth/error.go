package auth

type authError struct {
	msg     string
	wrapped error
}

func (authError) Type() string {
	return "auth"
}

func (a authError) Error() string {
	return a.msg
}

func (a authError) Unwrap() error {
	return a.wrapped
}

func (a authError) StatusCode() int {
	return 403
}
