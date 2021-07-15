package datastore

import "fmt"

type invalidKeyError struct {
	keys []string
}

func (i invalidKeyError) Error() string {
	return fmt.Sprintf("the given keys/fqfields are invalid: %v", i.keys)
}

func (i invalidKeyError) Type() string {
	return "invalid"
}
