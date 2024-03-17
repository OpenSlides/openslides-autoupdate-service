package keysbuilder_test

import (
	"sort"

	"github.com/OpenSlides/openslides-go/datastore/dskey"
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

func cmpSet(one, two map[dskey.Key]bool) []string {
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

func set(keys ...dskey.Key) map[dskey.Key]bool {
	out := make(map[dskey.Key]bool)
	for _, key := range keys {
		out[key] = true
	}
	return out
}

func mustKeys(ks ...string) []dskey.Key {
	keys := make([]dskey.Key, len(ks))
	for i, k := range ks {
		key := dskey.MustKey(k)
		keys[i] = key
	}
	return keys
}
