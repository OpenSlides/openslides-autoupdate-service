package autoupdate

import (
	"context"
	"fmt"
	"strconv"
	"strings"
)

// RestrictedIDs returns the ids of a field by using a restricter
type RestrictedIDs struct {
	user int
	r    Restricter
}

// IDs returns ids for a key
func (i RestrictedIDs) IDs(ctx context.Context, key string) ([]int, error) {
	var multi bool
	switch {
	case strings.HasSuffix(key, "_id"):
	case strings.HasSuffix(key, "_ids"):
		multi = true
	default:
		return nil, fmt.Errorf("key %s can not be a reference; expected suffix _id or _ids", key)
	}

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
