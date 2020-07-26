package keysbuilder_test

import (
	"context"
	"encoding/json"
	"sort"
	"time"
)

type mockDataProvider struct {
	err          error
	data         map[string]json.RawMessage
	sleep        time.Duration
	requestCount int
}

func (r *mockDataProvider) RestrictedData(ctx context.Context, uid int, keys ...string) (map[string]json.RawMessage, error) {
	time.Sleep(r.sleep)
	if r.err != nil {
		return nil, r.err
	}

	r.requestCount++

	data := make(map[string]json.RawMessage, len(keys))
	for _, key := range keys {
		v, ok := r.data[key]
		if !ok {
			data[key] = nil
			continue
		}
		data[key] = v
	}

	return data, nil
}

func cmpSlice(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}

func cmpSet(one, two map[string]bool) []string {
	var out []string

	for key := range one {
		if !two[key] {
			out = append(out, "-"+key)
		}
	}
	for key := range two {
		if !one[key] {
			out = append(out, "+"+key)
		}
	}
	if len(out) == 0 {
		return nil
	}
	sort.Strings(out)
	return out
}

func set(keys ...string) map[string]bool {
	out := make(map[string]bool)
	for _, key := range keys {
		out[key] = true
	}
	return out
}

func mapKeys(m map[string][]int) []string {
	out := make([]string, 0, len(m))
	for key := range m {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

func strs(str ...string) []string { return str }
func ids(ids ...int) []int        { return ids }
