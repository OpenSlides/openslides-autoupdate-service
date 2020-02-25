package keysbuilder_test

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysbuilder"
	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
)

func TestKeys(t *testing.T) {
	for _, tt := range []struct {
		name     string
		request  keysrequest.Body
		keys     []string
		reqCount int
	}{
		{
			"One Field",
			keysrequest.Body{IDs: ids(1), Fields: simple("user", "name")},
			keys("user/1/name"),
			0,
		},
		{
			"Many Fields",
			keysrequest.Body{IDs: ids(1), Fields: simple("user", "first", "last")},
			keys("user/1/first", "user/1/last"),
			0,
		},
		{
			"Many IDs Many Fields",
			keysrequest.Body{IDs: ids(1, 2), Fields: simple("user", "first", "last")},
			keys("user/1/first", "user/1/last", "user/2/first", "user/2/last"),
			0,
		},
		{
			"Redirect Once id",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", simple("note", "important")})},
			keys("user/1/note_id", "note/1/important"),
			1,
		},
		{
			"Redirect Once ids",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"group_ids", simple("group", "admin")})},
			keys("user/1/group_ids", "group/1/admin", "group/2/admin"),
			1,
		},
		{
			"Redirect twice id",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", komplex("note", entry{"motion_id", simple("motion", "name")})})},
			keys("user/1/note_id", "note/1/motion_id", "motion/1/name"),
			2,
		},
		{
			"Redirect twice ids",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"group_ids", komplex("group", entry{"perm_ids", simple("perm", "name")})})},
			keys("user/1/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name"),
			3,
		},
		{
			"Request _id without redirect",
			keysrequest.Body{IDs: ids(1), Fields: simple("user", "note_id")},
			keys("user/1/note_id"),
			0,
		},
		{
			"Redirect id not exist",
			keysrequest.Body{IDs: ids(1), Fields: komplex("not_exist", entry{"note_id", simple("note", "important")})},
			keys("not_exist/1/note_id"),
			1,
		},
		{
			"Redirect ids not exist",
			keysrequest.Body{IDs: ids(1), Fields: komplex("not_exist", entry{"group_ids", simple("note", "important")})},
			keys("not_exist/1/group_ids"),
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ider := &mockIDer{}
			b, err := keysbuilder.New(context.Background(), ider, tt.request)
			if err != nil {
				t.Fatalf("Expected New() not to return an error, got: %v", err)
			}

			keys := b.Keys()

			if diff := cmpSet(set(tt.keys...), set(keys...)); diff != nil {
				t.Errorf("Expected %v, got: %v", tt.keys, diff)
			}
			if len(ider.reqLog) != tt.reqCount {
				t.Errorf("Expected %d requests to the restricter, got: %d: %v", tt.reqCount, len(ider.reqLog), ider.reqLog)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	for _, tt := range []struct {
		name    string
		request keysrequest.Body
		newDB   map[string][]int
		got     []string
		count   int
	}{
		{
			"One relation",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", simple("note", "important")})},
			map[string][]int{"user/1/note_id": ids(2)},
			keys("user/1/note_id", "note/2/important"),
			1,
		},
		{
			"One relation no change",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", simple("note", "important")})},
			map[string][]int{},
			keys("user/1/note_id", "note/1/important"),
			0,
		},
		{
			"Two ids one change",
			keysrequest.Body{IDs: ids(1, 2), Fields: komplex("user", entry{"note_id", simple("note", "important")})},
			map[string][]int{"user/1/note_id": ids(2)},
			keys("user/1/note_id", "user/2/note_id", "note/1/important", "note/2/important"),
			1,
		},
		{
			"Two relation one change",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", simple("note", "important")}, entry{"group_ids", simple("group", "admin")})},
			map[string][]int{"user/1/note_id": ids(2)},
			keys("user/1/note_id", "user/1/group_ids", "note/2/important", "group/1/admin", "group/2/admin"),
			1,
		},
		{
			"Two relation two changes",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", simple("note", "important")}, entry{"group_ids", simple("group", "admin")})},
			map[string][]int{"user/1/note_id": ids(2), "user/1/group_ids": ids(2)},
			keys("user/1/note_id", "note/2/important", "user/1/group_ids", "group/2/admin"),
			2,
		},
		{
			"Tree levels out changes",
			keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"group_ids", komplex("group", entry{"perm_ids", simple("perm", "name")})})},
			map[string][]int{"user/1/group_ids": ids(2)},
			keys("user/1/group_ids", "group/2/perm_ids", "perm/2/name", "perm/1/name"),
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			restr := mockIDer{}
			b, err := keysbuilder.New(context.Background(), &restr, tt.request)
			if err != nil {
				t.Fatalf("Expect Keys() not to return an error, got: %v", err)
			}
			restr.data = tt.newDB
			restr.reqLog = make([]string, 0)

			if err := b.Update(mapKeys(tt.newDB)); err != nil {
				t.Errorf("Expect Update() not to return an error, got: %v", err)
			}

			if diff := cmpSet(set(tt.got...), set(b.Keys()...)); diff != nil {
				t.Errorf("Expected %v, got: %v", b.Keys(), diff)
			}
			if tt.count != len(restr.reqLog) {
				t.Errorf("Expected %d requests to the restricter, got: %d: %v", tt.count, len(restr.reqLog), restr.reqLog)
			}
		})
	}
}

func TestConcurency(t *testing.T) {
	kr := keysrequest.Body{IDs: ids(1, 2, 3), Fields: komplex("user", entry{"group_ids", komplex("group", entry{"perm_ids", simple("perm", "name")})})}
	restr := mockIDer{sleep: 10 * time.Millisecond}

	start := time.Now()
	b, err := keysbuilder.New(context.Background(), &restr, kr)
	if err != nil {
		t.Errorf("Expect Keys() not to return an error, got: %v", err)
	}
	finished := time.Since(start)

	if finished > 30*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 30 Milliseconds, got: %v", finished)
	}
	expect := keys("user/1/group_ids", "user/2/group_ids", "user/3/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name")
	if diff := cmpSet(set(expect...), set(b.Keys()...)); diff != nil {
		t.Errorf("Expected %v, got: %v", expect, diff)
	}
	if len(restr.reqLog) != 5 {
		t.Errorf("Expected %d requests to the restricter.IDsFromKey, got: %d: %v", 5, len(restr.reqLog), restr.reqLog)
	}
}

func TestManyRequests(t *testing.T) {
	krs := []keysrequest.Body{
		keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", simple("note", "important")})},
		keysrequest.Body{IDs: ids(1), Fields: simple("motion", "name")},
		keysrequest.Body{IDs: ids(2), Fields: komplex("user", entry{"note_id", simple("note", "important")})},
	}
	restr := mockIDer{sleep: 10 * time.Millisecond}

	start := time.Now()
	b, err := keysbuilder.New(context.Background(), &restr, krs...)
	if err != nil {
		t.Errorf("Expect Keys() not to return an error, got: %v", err)
	}

	finished := time.Since(start)
	if finished > 20*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 20 Milliseconds, got: %v", finished)
	}

	expect := keys("user/1/note_id", "user/2/note_id", "motion/1/name", "note/1/important")
	if diff := cmpSet(set(expect...), set(b.Keys()...)); diff != nil {
		t.Errorf("Expected %v, got: %v", expect, diff)
	}
	if len(restr.reqLog) != 2 {
		t.Errorf("Expected %d requests to the restricter.IDsFromKey, got: %d: %v", 2, len(restr.reqLog), restr.reqLog)
	}
}

func TestError(t *testing.T) {
	kr := keysrequest.Body{IDs: ids(1), Fields: komplex("user", entry{"note_id", simple("note", "important")})}
	restr := mockIDer{err: errors.New("Some Error"), sleep: 10 * time.Millisecond}

	start := time.Now()
	if _, err := keysbuilder.New(context.Background(), &restr, kr); err == nil {
		t.Errorf("Expect Keys() to return an error, got none")
	}
	finished := time.Since(start)

	if finished > 20*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 20 Milliseconds, got: %v", finished)
	}
}

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
		return nil, fmt.Errorf("Key %s can not be a reference; expected suffex _id or _ids", key)
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

func simple(collection string, fields ...string) keysrequest.Fields {
	names := make(map[string]keysrequest.Fields, len(fields))
	for _, f := range fields {
		names[f] = keysrequest.Fields{}
	}
	return keysrequest.Fields{Collection: collection, Names: names}
}

type entry struct {
	property string
	fields   keysrequest.Fields
}

func komplex(collection string, fe ...entry) keysrequest.Fields {
	names := make(map[string]keysrequest.Fields)
	for _, f := range fe {
		names[f.property] = f.fields
	}
	return keysrequest.Fields{Collection: collection, Names: names}
}
