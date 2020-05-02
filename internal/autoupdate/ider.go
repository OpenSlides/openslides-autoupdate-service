package autoupdate

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// RestrictedIDs implements the keybuilder.IDer interface by using a restricer.
type RestrictedIDs struct {
	user       int
	autoupdate *Autoupdate
}

// ID returns the id in the given key.
func (i RestrictedIDs) ID(ctx context.Context, key string) (int, error) {
	data, err := i.autoupdate.restrictedData(ctx, i.user, key)
	if err != nil {
		return 0, fmt.Errorf("get value for key `%s`: %w", key, err)
	}

	id, err := strconv.Atoi(data[key])
	if err != nil {
		return 0, ValueError{key: key}
	}

	return id, nil
}

// IDList returns the a list of ids in the key.
func (i RestrictedIDs) IDList(ctx context.Context, key string) ([]int, error) {
	data, err := i.autoupdate.restrictedData(ctx, i.user, key)
	if err != nil {
		return nil, fmt.Errorf("get value for key `%s`: %w", key, err)
	}

	var value []int
	if err := json.Unmarshal([]byte(data[key]), &value); err != nil {
		return nil, ValueError{key: key}
	}
	return value, nil
}

// GenericID returns a collection-id tuple.
func (i RestrictedIDs) GenericID(ctx context.Context, key string) (string, error) {
	data, err := i.autoupdate.restrictedData(ctx, i.user, key)
	if err != nil {
		return "", fmt.Errorf("get value for key `%s`: %w", key, err)
	}

	var value string
	if err := json.Unmarshal([]byte(data[key]), &value); err != nil {
		return "", ValueError{key: key}
	}
	return value, nil
}

// GenericIDs returns a list of collection-id tuples.
func (i RestrictedIDs) GenericIDs(ctx context.Context, key string) ([]string, error) {
	data, err := i.autoupdate.restrictedData(ctx, i.user, key)
	if err != nil {
		return nil, fmt.Errorf("get value for key `%s`: %w", key, err)
	}

	var values []string
	if err := json.Unmarshal([]byte(data[key]), &values); err != nil {
		return nil, ValueError{key: key}
	}
	return values, nil
}

// Template returns the strings from a template field.
func (i RestrictedIDs) Template(ctx context.Context, key string) ([]string, error) {
	return i.GenericIDs(ctx, key)
}
