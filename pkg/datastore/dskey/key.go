package dskey

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Key struct {
	collectionField string
	fieldIdx        int
	id              int
}

// FromString parses a string to a Key.
//
// This uses a regular expression to validate the key. This can be slow if
// called many times. It is faster to manually validate the key.
func FromString(in string) (Key, error) {
	if !keyValid(in) {
		return Key{}, invalidKeyError{in}
	}

	idx1 := strings.IndexByte(in, '/')
	idx2 := strings.LastIndexByte(in, '/')
	id, _ := strconv.Atoi(in[idx1+1 : idx2])

	key := Key{
		collectionField: in[:idx1] + in[idx2+1:],
		fieldIdx:        idx1,
		id:              id,
	}

	return key, nil
}

func FromParts(collection string, id int, field string) Key {
	return Key{
		collectionField: collection + field,
		fieldIdx:        len(collection),
		id:              id,
	}
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
	return fmt.Sprintf("%s/%d/%s", k.Collection(), k.ID(), k.Field())
}

func (k Key) ID() int {
	return k.id
}

func (k Key) Collection() string {
	return k.collectionField[:k.fieldIdx]
}

func (k Key) Field() string {
	return k.collectionField[k.fieldIdx:]
}

// FQID returns the FQID part of the field
func (k Key) FQID() string {
	return fmt.Sprintf("%s/%d", k.Collection(), k.id)
}

// CollectionField returns the first and last part of the key.
func (k Key) CollectionField() string {
	return k.Collection() + "/" + k.Field()
}

// IDField retuns the the /id field for the key.
func (k Key) IDField() Key {
	return Key{
		collectionField: k.collectionField[:k.fieldIdx] + "id",
		fieldIdx:        k.fieldIdx,
		id:              k.id,
	}
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
