package datastore

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Key represents a FQField.
type Key struct {
	Collection string
	ID         int
	Field      string
}

// KeyFromString parses a string to a Key.
func KeyFromString(in string) (Key, error) {
	if !keyValid(in) {
		return Key{}, invalidKeyError{in}
	}

	parts := strings.Split(in, "/")
	id, _ := strconv.Atoi(parts[1])
	return Key{parts[0], id, parts[2]}, nil
}

func (k Key) String() string {
	return fmt.Sprintf("%s/%d/%s", k.Collection, k.ID, k.Field)
}

// FQID returns the FQID part of the field
func (k Key) FQID() string {
	return fmt.Sprintf("%s/%d", k.Collection, k.ID)
}

// CollectionField returns the first and last part of the key.
func (k Key) CollectionField() string {
	return k.Collection + "/" + k.Field
}

// IDField retuns the the /id field for the key.
func (k Key) IDField() Key {
	k.Field = "id"
	return k
}

// MarshalJSON converts the key to a json string.
func (k Key) MarshalJSON() ([]byte, error) {
	return []byte(`"` + k.String() + `"`), nil
}

var reValidKeys = regexp.MustCompile(`^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*/[a-z][a-z0-9_]*\$?[a-z0-9_]*$`)

// keyValid checks if all of the given keys are valid. Invalid keys are
// returned.
//
// A return value of nil means, that all keys are valid.
func keyValid(key string) bool {
	return reValidKeys.MatchString(key)
}
