package keysbuilder_test

import (
	"sort"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

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

func cmpSet(one, two map[datastore.Key]bool) []string {
	var out []string

	for key := range one {
		if !two[key] {
			out = append(out, "-"+key.String())
		}
	}
	for key := range two {
		if !one[key] {
			out = append(out, "+"+key.String())
		}
	}
	if len(out) == 0 {
		return nil
	}
	sort.Strings(out)
	return out
}

func set(keys ...datastore.Key) map[datastore.Key]bool {
	out := make(map[datastore.Key]bool)
	for _, key := range keys {
		out[key] = true
	}
	return out
}

func keys(ks ...string) []datastore.Key {
	keys := make([]datastore.Key, len(ks))
	for i, k := range ks {
		key, err := datastore.KeyFromString(k)
		if err != nil {
			panic(err)
		}

		keys[i] = key
	}
	return keys
}
