package redis

type closingError struct{}

func (e closingError) Closing()      {}
func (e closingError) Error() string { return "closing" }
