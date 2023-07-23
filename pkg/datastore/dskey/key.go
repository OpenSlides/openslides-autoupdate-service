package dskey

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Key string

// Key represents a FQField.
type KeyOld struct {
	Collection string
	ID         int
	Field      string
}

// FromString parses a string to a Key.
//
// This uses a regular expression to validate the key. This can be slow if
// called many times. It is faster to manually validate the key.
func FromString(in string) (Key, error) {
	if !keyValid(in) {
		return "", invalidKeyError{in}
	}

	return Key(in), nil
}

func FromParts(collection string, id int, field string) Key {
	return Key(fmt.Sprintf("%s/%d/%s", collection, id, field))
}

// MustKey is like FromString but panics, if the key is invalid.
//
// Should only be used in tests.
func MustKey(in string) Key {
	k, err := FromString(in)
	if err != nil {
		panic(err)
	}
	return k
}

func (k Key) String() string {
	return string(k)
}

func (k Key) ID() int {
	idx1 := strings.IndexByte(string(k), '/')
	if idx1 < 0 {
		panic(fmt.Sprintf("invalid key %s", k))
	}

	idx2 := strings.LastIndexByte(string(k), '/')
	if idx2 < 0 {
		panic(fmt.Sprintf("invalid key %s", k))
	}

	n, err := strconv.Atoi(string(k)[idx1+1 : idx2])
	if err != nil {
		panic(fmt.Sprintf("invalid key: %s: %v", k, err))
	}

	return n
}

func (k Key) Collection() string {
	idx1 := strings.IndexByte(string(k), '/')
	if idx1 < 0 {
		panic(fmt.Sprintf("invalid key %s", k))
	}

	return string(k)[:idx1]
}

func (k Key) Field() string {
	idx2 := strings.LastIndexByte(string(k), '/')
	if idx2 < 0 {
		panic(fmt.Sprintf("invalid key %s", k))
	}

	return string(k)[idx2+1:]
}

// FQID returns the FQID part of the field
func (k Key) FQID() string {
	idx2 := strings.LastIndexByte(string(k), '/')
	if idx2 < 0 {
		panic(fmt.Sprintf("invalid key %s", k))
	}
	return string(k)[:idx2]
}

// CollectionField returns the first and last part of the key.
func (k Key) CollectionField() string {
	return k.Collection() + "/" + k.Field()
}

// IDField retuns the the /id field for the key.
func (k Key) IDField() Key {
	idx2 := strings.LastIndexByte(string(k), '/')
	if idx2 < 0 {
		panic(fmt.Sprintf("invalid key %s", k))
	}

	return Key(string(k)[:idx2] + "/id")
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

type invalidKeyError struct {
	key string
}

func (i invalidKeyError) Error() string {
	return fmt.Sprintf("the key/fqfield is invalid: %s", i.key)
}

func (i invalidKeyError) Type() string {
	return "invalid"
}
