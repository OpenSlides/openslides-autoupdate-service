package autoupdate

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const keySep = "/"

// ErrUnknownKey ist returned from RestrictedIDs, when the requested key is not
// returned from the restricter.
var ErrUnknownKey = errors.New("key does not exist")

// RestrictedIDs implements the IDer interface by using a restricer.
type RestrictedIDs struct {
	user int
	r    Restricter
}

// ID returns the id in the key.
func (i RestrictedIDs) ID(ctx context.Context, key string) (int, error) {
	data, err := i.decodedRestricter(ctx, key)
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, ValueError{key: key}
	}

	return id, nil
}

// IDList returns the a list of ids in the key.
func (i RestrictedIDs) IDList(ctx context.Context, key string) ([]int, error) {
	data, err := i.decodedRestricter(ctx, key)
	if err != nil {
		return nil, err
	}

	var value []int
	if err := json.Unmarshal(data, &value); err != nil {
		return nil, ValueError{key: key}
	}
	return value, nil
}

// GenericID returns a collection-id tuple.
func (i RestrictedIDs) GenericID(ctx context.Context, key string) (string, error) {
	data, err := i.decodedRestricter(ctx, key)
	if err != nil {
		return "", err
	}

	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return "", ValueError{key: key}
	}
	return value, nil
}

// GenericIDs returns a list of collection-id tuples.
func (i RestrictedIDs) GenericIDs(ctx context.Context, key string) ([]string, error) {
	data, err := i.decodedRestricter(ctx, key)
	if err != nil {
		return nil, err
	}

	var values []string
	if err := json.Unmarshal(data, &values); err != nil {
		return nil, ValueError{key: key}
	}
	return values, nil
}

// Template returns the strings from a template field.
func (i RestrictedIDs) Template(ctx context.Context, key string) ([]string, error) {
	return i.GenericIDs(ctx, key)
}

func (i RestrictedIDs) decodedRestricter(ctx context.Context, key string) (json.RawMessage, error) {
	r, err := i.r.Restrict(ctx, i.user, []string{key})
	if err != nil {
		return nil, fmt.Errorf("restrict key %s: %w", key, err)
	}

	return fromModel(key, r)
}

func fromModel(key string, r io.Reader) (json.RawMessage, error) {
	keyParts := strings.SplitN(key, keySep, 3)
	if len(keyParts) != 3 {
		return nil, fmt.Errorf("invalid key %s", key)
	}

	collection := keyParts[0]
	id := keyParts[1]
	field := keyParts[2]

	var data map[string]map[string]map[string]json.RawMessage
	if err := json.NewDecoder(r).Decode(&data); err != nil {
		return nil, fmt.Errorf("decode restricted data: %w", err)
	}

	if _, ok := data[collection]; !ok {
		return nil, ErrUnknownKey
	}
	if _, ok := data[collection][id]; !ok {
		return nil, ErrUnknownKey
	}
	if _, ok := data[collection][id][field]; !ok {
		return nil, ErrUnknownKey
	}

	return data[collection][id][field], nil
}
