package keysbuilder_test

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

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
			if diff := cmpSet(set(tt.keys...), set(keys...)); diff != nil {
				t.Errorf("Expected %v, got: %v", tt.keys, diff)
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
	restr := mockRestricter{}
	b, err := keysbuilder.New(1, &restr, kr)
	if err != nil {
		t.Errorf("Expect Keys() not to return an error, got: %v", err)
	}

	keys := b.Keys()
	expect := set("user/1/note_id", "note/1/important")
	if cmpSet(expect, set(keys...)) != nil {
		t.Errorf("Expected %v, got: %v", expect, keys)
	}
	if len(restr.reqLog) != 1 {
		t.Errorf("Expected %d requests to the restricter, got: %d: %v", 1, len(restr.reqLog), restr.reqLog)
	}

	restr.reqLog = make([]string, 0)
	restr.data = map[string][]int{"user/1/note_id": []int{2}}
	if err := b.Update([]string{"user/1/note_id"}); err != nil {
		t.Errorf("Expect Update() not to return an error, got: %v", err)
	}

	keys = b.Keys()
	expect = set("user/1/note_id", "note/2/important")
	if diff := cmpSet(expect, set(keys...)); diff != nil {
		t.Errorf("Expected %v, got: %v", mapKeys(expect), keys)
	}
	if len(restr.reqLog) != 1 {
		t.Errorf("Expected %d requests to the restricter, got: %d: %v", 1, len(restr.reqLog), restr.reqLog)
	}
}

func TestUpdateKomplex(t *testing.T) {
	req := `{
		"ids": [1],
		"collection": "user",
		"fields": {
			"group_ids": {
				"collection": "group",
				"fields": {
					"perm_ids": {
						"collection": "perm",
						"fields": {"name": null}
					}
				}
			}
		}
	}`
	dbInit := map[string][]int{
		"user/1/group_ids": []int{1, 2},
		"group/1/perm_ids": []int{1, 2},
		"group/2/perm_ids": []int{2, 3},
		"group/3/perm_ids": []int{3, 4},
	}
	dbChanged := map[string][]int{
		"user/1/group_ids": []int{2, 3},
		"group/1/perm_ids": []int{1, 2},
		"group/2/perm_ids": []int{2, 3},
		"group/3/perm_ids": []int{3, 4},
	}
	kr, err := keysrequest.FromJSON((strings.NewReader(req)))
	if err != nil {
		t.Fatalf("Did not expect an error, got :%v", err)
	}
	restr := mockRestricter{data: dbInit}
	b, err := keysbuilder.New(1, &restr, kr)
	if err != nil {
		t.Errorf("Expect Keys() not to return an error, got: %v", err)
	}

	keys := b.Keys()
	expect := set("user/1/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name", "perm/3/name")
	if cmpSet(expect, set(keys...)) != nil {
		t.Errorf("Expected %v, got: %v", expect, keys)
	}
	if len(restr.reqLog) != 3 {
		t.Errorf("Expected %d requests to the restricter, got: %d: %v", 3, len(restr.reqLog), restr.reqLog)
	}

	restr.data = dbChanged
	restr.reqLog = make([]string, 0)
	if err := b.Update([]string{"user/1/group_ids"}); err != nil {
		t.Errorf("Expect Update() not to return an error, got: %v", err)
	}

	keys = b.Keys()
	expect = set("user/1/group_ids", "group/2/perm_ids", "group/3/perm_ids", "perm/2/name", "perm/3/name", "perm/4/name")
	if diff := cmpSet(expect, set(keys...)); diff != nil {
		t.Errorf("Expected %v, got: %v", mapKeys(expect), diff)
	}
	if len(restr.reqLog) != 2 {
		t.Errorf("Expected %d requests to the restricter, got: %d: %v", 2, len(restr.reqLog), restr.reqLog)
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
						"collection": "perm",
						"fields": {"name": null}
					}
				}
			}
		}
	}`
	db1 := map[string][]int{"user/1/note_id": []int{2}, "user/1/group_ids": []int{2}}

	tc := []struct {
		name    string
		request string
		initDB  map[string][]int
		newDB   map[string][]int
		change  []string
		count   int
	}{
		{"One relation", j1, nil, db1, []string{"user/1/note_id"}, 1},
		{"One relation no change", j1, nil, db1, []string{"user/1/name"}, 0},
		{"Two ids one change", j2, nil, db1, []string{"user/1/note_id"}, 1},
		{"Two relation one change", j3, nil, db1, []string{"user/1/note_id"}, 1},
		{"Two relation two changes", j3, nil, db1, []string{"user/1/note_id", "user/1/group_ids"}, 2},
		{"Tree levels out changes", j4, nil, db1, []string{"user/1/group_ids"}, 1},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			kr, err := keysrequest.FromJSON((strings.NewReader(tt.request)))
			if err != nil {
				t.Fatalf("Expect FromJSON not to return an error, got: %v", err)
			}
			restr := mockRestricter{data: tt.initDB}
			b, err := keysbuilder.New(1, &restr, kr)
			if err != nil {
				t.Errorf("Expect Keys() not to return an error, got: %v", err)
			}
			restr.data = tt.newDB
			restr.reqLog = make([]string, 0)
			if err := b.Update(tt.change); err != nil {
				t.Errorf("Expect Update() not to return an error, got: %v", err)
			}
			if tt.count != len(restr.reqLog) {
				t.Errorf("Expected %d requests to the restricter, got: %d: %v", tt.count, len(restr.reqLog), restr.reqLog)
			}
		})
	}
}

func TestConcurency(t *testing.T) {
	t.Parallel()
	req := `{
		"ids": [1, 2, 3],
		"collection": "user",
		"fields": {
			"group_ids": {
				"collection": "group",
				"fields": {
					"perm_ids": {
						"collection": "perm",
						"fields": {"name": null}
					}
				}
			}
		}
	}`
	kr, err := keysrequest.FromJSON((strings.NewReader(req)))
	if err != nil {
		t.Fatalf("Did not expect an error, got :%v", err)
	}
	restr := mockRestricter{sleep: 10 * time.Millisecond}

	start := time.Now()
	b, err := keysbuilder.New(1, &restr, kr)
	if err != nil {
		t.Errorf("Expect Keys() not to return an error, got: %v", err)
	}

	keys := b.Keys()
	finished := time.Since(start)
	if finished > 30*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 30 Milliseconds, got: %v", finished)

	}
	expect := []string{"user/1/group_ids", "user/2/group_ids", "user/3/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name"}
	if diff := cmpSet(set(expect...), set(keys...)); diff != nil {
		t.Errorf("Expected %v, got: %v", expect, diff)
	}
	if len(restr.reqLog) != 5 {
		t.Errorf("Expected %d requests to the restricter.IDsFromKey, got: %d: %v", 3, len(restr.reqLog), restr.reqLog)
	}
}

type mockRestricter struct {
	data  map[string][]int
	sleep time.Duration

	reqLogMu sync.Mutex
	reqLog   []string
}

func (r *mockRestricter) Restrict(ctx context.Context, uid int, keys []string) (map[string][]byte, error) {
	time.Sleep(r.sleep)

	r.reqLogMu.Lock()
	r.reqLog = append(r.reqLog, strings.Join(keys, ", "))
	r.reqLogMu.Unlock()

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
	time.Sleep(r.sleep)

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
		return nil, fmt.Errorf("Key %s can not be a reference; expected suffex _id or _ids", key)
	}
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
}

func (r *mockRestricter) IDsFromCollection(ctx context.Context, uid int, mid int, collection string) ([]int, error) {
	time.Sleep(r.sleep)
	r.reqLog = append(r.reqLog, collection)
	if ids, ok := r.data[collection]; ok {
		return ids, nil
	}
	if mid == 1 {
		return []int{1}, nil
	}
	return []int{1, 2}, nil
}

func cmpSet(one, two map[string]bool) []string {
	out := make([]string, 0)

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

func mapKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for key := range m {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}
