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
			b, err := keysbuilder.New(1, &mockRestricter{}, kr)
			if err != nil {
				t.Fatalf("Expected New() not to return an error, got: %v", err)
			}

			keys := b.Keys()
			if !cmpSlice(tt.keys, keys) {
				t.Errorf("Expected %v, got: %v", tt.keys, keys)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	req := `{"ids":[1],"collection":"user","fields":{"note_id":{"collection":"note","fields":{"important":null}}}}`
	kr, err := keysrequest.FromJSON((strings.NewReader(req)))
	if err != nil {
		t.Fatalf("Did not expect an error, got :%v", err)
	}
	restr := mockRestricter{nextIDs: []int{1}}
	b, err := keysbuilder.New(1, &restr, kr)
	if err != nil {
		t.Errorf("Expect Keys() not to return an error, got: %v", err)
	}

	keys := b.Keys()
	expect := []string{"user/1/note_id", "note/1/important"}
	if !cmpSlice(expect, keys) {
		t.Errorf("Expected %v, got: %v", expect, keys)
	}

	restr.nextIDs = []int{2}
	if err := b.Update([]string{"user/1/note_id"}); err != nil {
		t.Errorf("Expect Update() not to return an error, got: %v", err)
	}

	keys = b.Keys()
	expect = []string{"user/1/note_id", "note/2/important"}
	if !cmpSlice(expect, keys) {
		t.Errorf("Expected %v, got: %v", expect, keys)
	}
}

func TestUpdateRequestCount(t *testing.T) {
	j1 := `{
		"ids": [1],
		"collection": "user",
		"fields": {
			"note_id": {
				"collection": "note",
				"fields": {"important": null}
			}
		}
	}`
	j2 := `{
		"ids": [1, 2],
		"collection": "user",
		"fields": {
			"note_id": {
				"collection": "note",
				"fields": {"important": null}
			}
		}
	}`
	j3 := `{
		"ids": [1],
		"collection": "user",
		"fields": {
			"note_id": {
				"collection": "note",
				"fields": {"important": null}
			},
			"group_ids": {
				"collection": "group",
				"fields": {"admin": null}
			}
		}
	}`
	j4 := `{
		"ids": [1],
		"collection": "user",
		"fields": {
			"group_ids": {
				"collection": "group",
				"fields": {
					"perm_ids": {
						"collection": "permission",
						"fields": {"can_test": null}
					}
				}
			}
		}
	}`

	tc := []struct {
		name    string
		request string
		change  []string
		count   int
	}{
		{"One relation", j1, []string{"user/1/note_id"}, 1},
		{"One relation no change", j1, []string{"user/1/name"}, 0},
		{"Two ids one change", j2, []string{"user/1/note_id"}, 1},
		{"Two relation one change", j3, []string{"user/1/note_id"}, 1},
		{"Two relation two changes", j3, []string{"user/1/note_id", "user/1/group_ids"}, 2},
		{"Tree levels out changes", j4, []string{"user/1/group_ids"}, 2},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			kr, err := keysrequest.FromJSON((strings.NewReader(tt.request)))
			if err != nil {
				t.Fatalf("Expect FromJSON not to return an error, got: %v", err)
			}
			restr := mockRestricter{}
			b, err := keysbuilder.New(1, &restr, kr)
			if err != nil {
				t.Errorf("Expect Keys() not to return an error, got: %v", err)
			}
			restr.nextIDs = []int{2}
			restr.reqCount = 0
			if err := b.Update(tt.change); err != nil {
				t.Errorf("Expect Update() not to return an error, got: %v", err)
			}
			if tt.count != restr.reqCount {
				t.Errorf("Expected %d requests to the restricter, got: %d", tt.count, restr.reqCount)
			}
		})
	}

}

type mockRestricter struct {
	nextIDs  []int
	reqCount int
}

func (r *mockRestricter) Restrict(ctx context.Context, uid int, keys []string) (map[string][]byte, error) {
	r.reqCount++
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

func (r *mockRestricter) IDsFromKey(ctx context.Context, uid int, mid int, key string) ([]int, error) {
	r.reqCount++
	if len(r.nextIDs) > 0 {
		return r.nextIDs, nil
	}
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

func (r *mockRestricter) IDsFromCollection(ctx context.Context, uid int, mid int, collection string) ([]int, error) {
	r.reqCount++
	if len(r.nextIDs) > 0 {
		return r.nextIDs, nil
	}
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
