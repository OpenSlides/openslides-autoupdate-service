package autoupdate

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
)

// restrictedIDs returns the ids of a field by using a restricter
type restrictedIDs struct {
	user int
	r    Restricter
}

func (i restrictedIDs) IDs(ctx context.Context, key string) ([]int, error) {
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

	id, err := strconv.Atoi(string(rawIDs))
	if err != nil {
		return nil, fmt.Errorf("value in key %s is not an int, got: %s", key, rawIDs)
	}
	return []int{id}, nil
}

func decodeNumberList(buf []byte) ([]int, error) {
	if len(buf) < 3 {
		return nil, fmt.Errorf("invalid value, expect list of ints")
	}
	if buf[0] != '[' || buf[len(buf)-1] != ']' {
		return nil, fmt.Errorf("expected first and last byte to be [ and ]")
	}
	var out []int
	buf = buf[1:]
	var idx int
	for {
		idx = bytes.IndexByte(buf, ',')
		if idx == -1 {
			break
		}
		id, err := strconv.Atoi(string(buf[:idx]))
		if err != nil {
			return nil, fmt.Errorf("can not convert value `%s` to int", buf[:idx])
		}
		out = append(out, id)
		buf = buf[idx+1:]
	}
	id, err := strconv.Atoi(string(buf[:len(buf)-1]))
	out = append(out, id)
	return out, err
}
