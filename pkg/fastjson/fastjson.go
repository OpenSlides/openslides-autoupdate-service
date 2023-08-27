// Package fastjson provides some function for hast json decoding for special
// types.
package fastjson

import (
	"bytes"
	"fmt"
	"strconv"
)

// DecodeInt decodes a json int value to an int type.
func DecodeInt(bs []byte) (int, error) {
	return strconv.Atoi(string(bs))
}

// DecodeIntList decodes a json List[int] value to an []int type.
func DecodeIntList(bs []byte) ([]int, error) {
	// Remove [ and ]
	if len(bs) < 2 {
		return nil, fmt.Errorf("invalid int list: %s", bs)
	}
	bs = bs[1 : len(bs)-1]

	numbers := bytes.Split(bs, []byte(","))

	out := make([]int, 0, len(numbers))
	for i, n := range numbers {
		n = bytes.TrimSpace(n)
		if len(n) == 0 {
			continue
		}

		v, err := DecodeInt(n)
		if err != nil {
			return nil, fmt.Errorf("%dth value, `%s`,  is not a number: %w", i, n, err)
		}

		out = append(out, v)
	}

	return out, nil
}
