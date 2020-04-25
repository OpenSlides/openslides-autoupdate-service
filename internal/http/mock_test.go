package http_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func mustRequest(r *http.Request, err error) *http.Request {
	if err != nil {
		panic(err)
	}
	return r
}

type mockAuth struct {
	uid int
}

func (a mockAuth) Authenticate(context.Context, *http.Request) (int, error) {
	return a.uid, nil
}

func keys(ks ...string) []string {
	return ks
}

func mapKeys(m map[string]map[string]map[string]json.RawMessage) []string {
	out := make([]string, 0, len(m))
	for collection, v := range m {
		for id, v := range v {
			for field := range v {
				out = append(out, fmt.Sprintf("%s/%s/%s", collection, id, field))
			}
		}

	}
	return out
}

func cmpSlice(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}
	sort.Strings(one)
	sort.Strings(two)
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
