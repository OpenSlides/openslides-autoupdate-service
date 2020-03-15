package autoupdate

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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

// Template returns the strings from a template field.
func (i RestrictedIDs) Template(ctx context.Context, key string) ([]string, error) {
	data, err := i.r.Restrict(ctx, i.user, []string{key})
	if err != nil {
		return nil, fmt.Errorf("can not restrict key %s: %w", key, err)
	}

	var values []string
	if err := json.Unmarshal([]byte(data[key]), &values); err != nil {
		return nil, fmt.Errorf("can not decode template value from restricter: %w", err)
	}
	return values, nil
}

// ids returns ids for a key.
func (i RestrictedIDs) ids(ctx context.Context, key string, multi bool) ([]int, error) {
	data, err := i.r.Restrict(ctx, i.user, []string{key})
	if err != nil {
		return nil, fmt.Errorf("can not restrict key %s: %w", key, err)
	}

	rawIDs, ok := data[key]
	if !ok {
		return []int{}, nil
	}

	if multi {
		return decodeNumberList(rawIDs)
	}

	id, err := strconv.Atoi(rawIDs)
	if err != nil {
		return nil, fmt.Errorf("value in key %s is not an int, got: %s", key, rawIDs)
	}
	return []int{id}, nil
}

func decodeNumberList(value string) ([]int, error) {
	if len(value) < 3 {
		return nil, fmt.Errorf("invalid value, expect list of ints")
	}
	if value[0] != '[' || value[len(value)-1] != ']' {
		return nil, fmt.Errorf("expected first and last byte to be [ and ]")
	}
	var out []int
	value = value[1:]
	var idx int
	for {
		idx = strings.IndexByte(value, ',')
		if idx == -1 {
			break
		}
		id, err := strconv.Atoi(value[:idx])
		if err != nil {
			return nil, fmt.Errorf("can not convert value `%s` to int", value[:idx])
		}
		out = append(out, id)
		value = value[idx+1:]
	}
	id, err := strconv.Atoi(value[:len(value)-1])
	out = append(out, id)
	return out, err
}
