package http_test

import (
	"net/http"
	"sort"
)

func mustRequest(r *http.Request, err error) *http.Request {
	if err != nil {
		panic(err)
	}
	return r
}

func keys(ks ...string) []string {
	return ks
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
