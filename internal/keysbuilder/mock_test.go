package keysbuilder_test

import (
	"context"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
)

type mockIDer struct {
	err   error
	data  map[string][]int
	sleep time.Duration
}

func (r *mockIDer) ID(ctx context.Context, key string) (int, error) {
	time.Sleep(r.sleep)
	if r.err != nil {
		return 0, r.err
	}

	if ids, ok := r.data[key]; ok {
		return ids[0], nil
	}
	if strings.HasPrefix(key, "not_exist") {
		return 0, autoupdate.ErrUnknownKey
	}
	return 1, nil
}

func (r *mockIDer) IDList(ctx context.Context, key string) ([]int, error) {
	time.Sleep(r.sleep)
	if r.err != nil {
		return nil, r.err
	}

	if ids, ok := r.data[key]; ok {
		return ids, nil
	}
	if strings.HasPrefix(key, "not_exist") {
		return nil, autoupdate.ErrUnknownKey
	}
	return ids(1, 2), nil
}

func (r *mockIDer) GenericID(ctx context.Context, key string) (string, error) {
	time.Sleep(r.sleep)
	if r.err != nil {
		return "", r.err
	}

	if strings.HasPrefix(key, "not_exist") {
		return "", autoupdate.ErrUnknownKey
	}
	return "other/1", nil
}

func (r *mockIDer) GenericIDs(ctx context.Context, key string) ([]string, error) {
	time.Sleep(r.sleep)
	if r.err != nil {
		return nil, r.err
	}

	if strings.HasPrefix(key, "not_exist") {
		return nil, autoupdate.ErrUnknownKey
	}
	return strs("other/1", "other/2"), nil
}

func (r *mockIDer) Template(ctx context.Context, key string) ([]string, error) {
	time.Sleep(r.sleep)
	if r.err != nil {
		return nil, r.err
	}

	if ids, ok := r.data[key]; ok {
		var out []string
		for _, id := range ids {
			out = append(out, strconv.Itoa(id))
		}
		return out, nil
	}
	if strings.HasPrefix(key, "not_exist") {
		return nil, nil
	}
	return strs("1", "2"), nil

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
