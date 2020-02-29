package http

// err400 is returned, when the user input is wrong.
type err400 struct {
	err error
}

// raiseErr400 wraps an error in an err400.
func raiseErr400(e error) err400 {
	return err400{err: e}
}

func (e err400) Error() string {
	return e.err.Error()
}

// Unwrap returns the wrapped error.
func (e err400) Unwrap() error {
	return e.err
}
