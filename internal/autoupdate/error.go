package autoupdate

// ErrInput is returned, when the user input is wrong
// It is like an http 4xx error
type ErrInput struct {
	err error
}

func raiseErrInput(e error) ErrInput {
	return ErrInput{err: e}
}

func (e ErrInput) Error() string {
	return e.err.Error()
}

// Unwrap returns the wrapped error
func (e ErrInput) Unwrap() error {
	return e.err
}
