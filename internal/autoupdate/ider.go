package autoupdate

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// RestrictedIDs implements the IDer interface by using a restricer.
type RestrictedIDs struct {
	user int
	r    Restricter
}

// ID returns the id in the key.
func (i RestrictedIDs) ID(ctx context.Context, key string) (int, error) {
	ids, err := i.ids(ctx, key, false)
	if err != nil {
		return 0, err
	}
	if len(ids) == 0 {
		return 0, nil
	}
	return ids[0], nil
}

// IDList returns the ids in the key.
func (i RestrictedIDs) IDList(ctx context.Context, key string) ([]int, error) {
	ids, err := i.ids(ctx, key, true)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// GenericID returns a collection-id tuple.
func (i RestrictedIDs) GenericID(ctx context.Context, key string) (string, error) {
	data, err := i.r.Restrict(ctx, i.user, []string{key})
	if err != nil {
		return "", fmt.Errorf("can not restrict key %s: %w", key, err)
	}

	if _, ok := data[key]; !ok {
		return "", nil
	}

	var value string
	if err := json.Unmarshal([]byte(data[key]), &value); err != nil {
		return "", fmt.Errorf("can not decode generic value from restricter: %w", err)
	}
	return value, nil
}

// GenericIDs returns a list of collection-id tuple.
func (i RestrictedIDs) GenericIDs(ctx context.Context, key string) ([]string, error) {
	data, err := i.r.Restrict(ctx, i.user, []string{key})
	if err != nil {
		return nil, fmt.Errorf("can not restrict key %s: %w", key, err)
	}

	if _, ok := data[key]; !ok {
		return nil, nil
	}

	var values []string
	if err := json.Unmarshal([]byte(data[key]), &values); err != nil {
		return nil, fmt.Errorf("can not decode generic-list value from restricter: %w", err)
	}
	return values, nil
}

// Template returns the strings from a template field.
func (i RestrictedIDs) Template(ctx context.Context, key string) ([]string, error) {
	return i.GenericIDs(ctx, key)
}

// ids returns ids for a key.
func (i RestrictedIDs) ids(ctx context.Context, key string, multi bool) ([]int, error) {
	data, err := i.r.Restrict(ctx, i.user, []string{key})
	if err != nil {
		return nil, fmt.Errorf("can not restrict key %s: %w", key, err)
	}

	rawIDs, ok := data[key]
	if !ok {
		return nil, nil
	}

	if multi {
		var value []int
		if err := json.Unmarshal([]byte(rawIDs), &value); err != nil {
			return nil, fmt.Errorf("can not read ids from restricter: %w", err)
		}
		return value, nil
		//return decodeNumberList(rawIDs)
	}

	id, err := strconv.Atoi(rawIDs)
	if err != nil {
		return nil, fmt.Errorf("value in key %s is not an int, got: %s", key, rawIDs)
	}
	return []int{id}, nil
}
