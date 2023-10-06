package datastore

import (
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// InvalidDataError is returned, when a key has an invalid value in the
// database.
type InvalidDataError struct {
	Key   dskey.Key
	Value []byte
}

func (err InvalidDataError) Error() string {
	return fmt.Sprintf("key %s has an invalid value %v", err.Key, err.Value)
}
