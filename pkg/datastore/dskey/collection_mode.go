package dskey

import (
	"fmt"
	"strconv"
	"strings"
)

// CollectionMode represents a collection/Mode like motion/A.
type CollectionMode uint64

// CollectionModeFromString parses a string to a Key.
//
// This uses a regular expression to validate the key. This can be slow if
// called many times. It is faster to manually validate the key.
func CollectionModeFromString(format string, a ...any) (CollectionMode, error) {
	keyStr := fmt.Sprintf(format, a...)
	idx1 := strings.IndexByte(keyStr, '/')
	idx2 := strings.LastIndexByte(keyStr, '/')
	if idx1 == -1 || idx1 == idx2 {
		return 0, InvalidKeyError{keyStr}
	}

	id, _ := strconv.Atoi(keyStr[idx1+1 : idx2])

	cfID := collectionModeToID(keyStr[:idx1] + "/" + keyStr[idx2+1:])
	if cfID == -1 {
		return 0, InvalidKeyError{keyStr}
	}
	return CollectionMode(joinInt(cfID, id)), nil
}

// CollectionModeFromParts create a Mode from collection, id an field.
func CollectionModeFromParts(collection string, id int, field string) (CollectionMode, error) {
	cfID := collectionModeToID(collection + "/" + field)
	if cfID == -1 {
		return 0, InvalidKeyError{fmt.Sprintf("%s/%d/%s", collection, id, field)}
	}

	if id <= 0 {
		return 0, InvalidKeyError{fmt.Sprintf("%s/%d/%s", collection, id, field)}
	}

	return CollectionMode(joinInt(cfID, id)), nil
}

// MustCollectionMode is like FromString but panics, if the key is invalid.
//
// Should only be used in tests.
func MustCollectionMode(format string, a ...any) CollectionMode {
	k, err := CollectionModeFromString(format, a...)
	if err != nil {
		panic(err)
	}
	return k
}

func (k CollectionMode) String() string {
	return fmt.Sprintf("%s/%d/%s", k.Collection(), k.ID(), k.Mode())
}

// ID returns the id attribute from the Key.
func (k CollectionMode) ID() int {
	_, id := splitUInt64(uint64(k))
	return id
}

// Collection returns the collection attribute from the Key.
func (k CollectionMode) Collection() string {
	cfIdx, _ := splitUInt64(uint64(k))
	return collectionModeFields[cfIdx].collection
}

// Mode returns the Mode attribute from the CollectionMode.
func (k CollectionMode) Mode() string {
	cfIdx, _ := splitUInt64(uint64(k))
	return collectionModeFields[cfIdx].field
}

// InvalidCollectionModeError is returned from dskey.CollectionModeFromKey or
// dskey.CollectionModeFromParts, if the collection mode in not valid.
type InvalidCollectionModeError struct {
	key string
}

func (i InvalidCollectionModeError) Error() string {
	return fmt.Sprintf("the key/mode is invalid: %s", i.key)
}

// Type returns "invalid"
func (i InvalidCollectionModeError) Type() string {
	return "invalid"
}
