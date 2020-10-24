package http_test

import (
	"context"
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

func (a mockAuth) Authenticate(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	return r.Context(), nil
}

func (a mockAuth) FromContext(ctx context.Context) int {
	return a.uid
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
