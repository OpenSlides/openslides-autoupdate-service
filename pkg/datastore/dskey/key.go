package dskey

import (
	"fmt"
	"strconv"
	"strings"
)

// Key represents a FQField.
type Key uint64

// FromString parses a string to a Key.
//
// This uses a regular expression to validate the key. This can be slow if
// called many times. It is faster to manually validate the key.
func FromString(format string, a ...any) (Key, error) {
	keyStr := fmt.Sprintf(format, a...)
	idx1 := strings.IndexByte(keyStr, '/')
	idx2 := strings.LastIndexByte(keyStr, '/')
	if idx1 == -1 || idx1 == idx2 {
		return 0, InvalidKeyError{keyStr}
	}

	id, _ := strconv.Atoi(keyStr[idx1+1 : idx2])

	cfID := collectionFieldToID(keyStr[:idx1] + "/" + keyStr[idx2+1:])
	if cfID == -1 {
		return 0, InvalidKeyError{keyStr}
	}
	return Key(joinInt(cfID, id)), nil
}

// FromParts create a key from collection, id an field.
func FromParts(collection string, id int, field string) (Key, error) {
	// TODO: Use a separate function with different namespace for mode-keys
	cfID := collectionFieldToID(collection + "/" + field)
	if cfID == -1 {
		return 0, InvalidKeyError{fmt.Sprintf("%s/%d/%s", collection, id, field)}
	}

	if id <= 0 {
		return 0, InvalidKeyError{fmt.Sprintf("%s/%d/%s", collection, id, field)}
	}

	return Key(joinInt(cfID, id)), nil
}

// MustKey is like FromString but panics, if the key is invalid.
//
// Should only be used in tests.
func MustKey(format string, a ...any) Key {
	k, err := FromString(format, a...)
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

// InvalidKeyError is returned from dskey.FromKey or dskey.FromParts, if the key
// in not valid.
type InvalidKeyError struct {
	key string
}

func (i InvalidKeyError) Error() string {
	return fmt.Sprintf("the key/fqfield is invalid: %s", i.key)
}

// Type returns "invalid"
func (i InvalidKeyError) Type() string {
	return "invalid"
}
