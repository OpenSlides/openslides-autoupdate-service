package http

// noStatusCodeError helps the errorHandler do decide, if an status code can be
// set.
type noStatusCodeError struct {
	wrapped error
}

func (e noStatusCodeError) Error() string {
	return e.wrapped.Error()
}
