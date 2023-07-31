package dskey

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Key ...
type Key uint64

// FromString parses a string to a Key.
//
// This uses a regular expression to validate the key. This can be slow if
// called many times. It is faster to manually validate the key.
func FromString(in string) (Key, error) {
	// if !keyValid(in) {
	// 	return 0, invalidKeyError{in}
	// }

	idx1 := strings.IndexByte(in, '/')
	idx2 := strings.LastIndexByte(in, '/')
	if idx1 == -1 || idx1 == idx2 {
		return 0, invalidKeyError{in}
	}

	id, _ := strconv.Atoi(in[idx1+1 : idx2])

	cfID := collectionFieldToID(in[:idx1] + "/" + in[idx2+1:])
	if cfID == -1 {
		return 0, invalidKeyError{in}
	}
	return Key(joinInt(cfID, id)), nil
}

// FromParts create a key from collection, id an field.
func FromParts(collection string, id int, field string) (Key, error) {
	// TODO: Use a separate function with different namespace for mode-keys
	cfID := collectionFieldToID(collection + "/" + field)
	if cfID == -1 {
		return 0, invalidKeyError{fmt.Sprintf("%s/%d/%s", collection, id, field)}
	}

	return Key(joinInt(cfID, id)), nil
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

// ID returns the id attribute from the Key.
func (k Key) ID() int {
	_, id := splitUInt64(uint64(k))
	return id
}

// Collection returns the collection attribute from the Key.
func (k Key) Collection() string {
	cfIdx, _ := splitUInt64(uint64(k))
	return collectionFields[cfIdx].collection
}

// Field returns the Field attribute from the key.
func (k Key) Field() string {
	cfIdx, _ := splitUInt64(uint64(k))
	return collectionFields[cfIdx].field
}

// FQID returns the FQID part of the field
func (k Key) FQID() string {
	cfIdx, id := splitUInt64(uint64(k))
	return fmt.Sprintf("%s/%d", collectionFields[cfIdx].collection, id)
}

// CollectionField returns the first and last part of the key.
func (k Key) CollectionField() string {
	cfIdx, _ := splitUInt64(uint64(k))
	return collectionFields[cfIdx].collection + "/" + collectionFields[cfIdx].field
}

// IDField retuns the the /id field for the key.
func (k Key) IDField() Key {
	idCfID := collectionFieldToID(k.Collection() + "/id")

	return Key(joinInt(idCfID, k.ID()))
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
