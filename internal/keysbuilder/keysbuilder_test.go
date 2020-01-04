package keysbuilder_test

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/openslides/openslides-autoupdate-service/internal/keysrequest"
)

func TestKeys(t *testing.T) {

	tc := []struct {
		name    string
		request string
		keys    []string
	}{
		{"One Field", `{"ids":[1],"collection":"user","fields":{"name":null}}`, []string{"user/1/name"}},
		{"Many Fields", `{"ids":[1],"collection":"user","fields":{"first":null,"last":null}}`, []string{"user/1/first", "user/1/last"}},
		{"Many IDs Many Fields", `{"ids":[1,2],"collection":"user","fields":{"first":null,"last":null}}`, []string{"user/1/first", "user/1/last", "user/2/first", "user/2/last"}},
		{"All IDs", `{"collection":"user","fields":{"name":null}}`, []string{"user/1/name", "user/2/name"}},
		{"All IDs one meeting", `{"meeting_id": 1,"collection":"user","fields":{"name":null}}`, []string{"user/1/name"}},
		{"Redirect Once id", `{"ids":[1],"collection":"user","fields":{"note_id":{"collection":"note","fields":{"important":null}}}}`, []string{"user/1/note_id", "note/1/important"}},
		{"Redirect Once ids", `{"ids":[1],"collection":"user","fields":{"group_ids":{"collection":"group","fields":{"admin":null}}}}`, []string{"user/1/group_ids", "group/1/admin", "group/2/admin"}},
		{"Redirect Once ids one meeting", `{"ids":[1], "meeting_id": 1, "collection":"user","fields":{"group_ids":{"collection":"group","fields":{"admin":null}}}}`, []string{"user/1/group_ids", "group/1/admin"}},
		{"Redirect twice id", `{"ids":[1],"collection":"user","fields":{"note_id":{"collection":"note","fields":{"motion_id":{"collection":"motion","fields":{"name":null}}}}}}`, []string{"user/1/note_id", "note/1/motion_id", "motion/1/name"}},
		{"Request _id without redirect", `{"ids":[1],"collection":"user","fields":{"note_id":null}}`, []string{"user/1/note_id"}},
		{"Redirect id not exist", `{"ids":[1],"collection":"not_exist","fields":{"note_id":{"collection":"note","fields":{"important":null}}}}`, []string{"not_exist/1/note_id"}},
		{"Redirect ids not exist", `{"ids":[1],"collection":"not_exist","fields":{"group_ids":{"collection":"group","fields":{"name":null}}}}`, []string{"not_exist/1/group_ids"}},
	}
	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			kr, err := keysrequest.FromJSON(strings.NewReader(tt.request))
			if err != nil {
				t.Fatalf("Did not expect an error, got: %v", err)
			}
			b := keysbuilder.Builder{
				User:    1,
				Restr:   mockRestricter{},
				Request: kr,
			}
			keys, err := b.Keys()
			if err != nil {
				t.Fatalf("Expected no error, got: %v", err)
			}
			if !cmpSlice(tt.keys, keys) {
				t.Errorf("Expected %v, got: %v", tt.keys, keys)
			}
		})
	}
}

type mockRestricter struct{}

func (r mockRestricter) Restrict(ctx context.Context, uid int, keys []string) (map[string][]byte, error) {
	out := make(map[string][]byte, len(keys))
	for _, key := range keys {
		switch {
		case strings.HasSuffix(key, "_id"):
			out[key] = []byte("1")
		case strings.HasSuffix(key, "_ids"):
			out[key] = []byte("[1,2]")
		default:
			out[key] = []byte("some value")
		}
	}
	return out, nil
}

func (r mockRestricter) IDsFromKey(ctx context.Context, uid int, mid int, key string) ([]int, error) {
	if strings.HasPrefix(key, "not_exist") {
		return nil, nil
	}
	if strings.HasSuffix(key, "_id") {
		return []int{1}, nil
	}
	if !strings.HasSuffix(key, "_ids") {
		return nil, fmt.Errorf("Key %s can not be a reference; expected suffex _id or _ids", key)
	}
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
}

func (r mockRestricter) IDsFromCollection(ctx context.Context, uid int, mid int, collection string) ([]int, error) {
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
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
