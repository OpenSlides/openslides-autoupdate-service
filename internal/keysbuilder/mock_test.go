package keysbuilder_test

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"
)

type mockIDer struct {
	err   error
	data  map[string][]int
	sleep time.Duration

	reqLogMu sync.Mutex
	reqLog   []string
}

func (r *mockIDer) IDs(ctx context.Context, key string) ([]int, error) {
	time.Sleep(r.sleep)
	if r.err != nil {
		return nil, r.err
	}

	r.reqLogMu.Lock()
	r.reqLog = append(r.reqLog, key)
	r.reqLogMu.Unlock()

	if ids, ok := r.data[key]; ok {
		return ids, nil
	}
	if strings.HasPrefix(key, "not_exist") {
		return nil, nil
	}
	if strings.HasSuffix(key, "_id") {
		return []int{1}, nil
	}
	if !strings.HasSuffix(key, "_ids") {
		return nil, fmt.Errorf("key %s can not be a reference; expected suffex _id or _ids", key)
	}
	return []int{1, 2}, nil
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

func keys(keys ...string) []string { return keys }
func ids(ids ...int) []int         { return ids }
