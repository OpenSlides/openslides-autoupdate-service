package restrict_test

import (
	"encoding/json"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestRestrict(t *testing.T) {
	perms := new(test.MockPermission)
	perms.Data = map[string]bool{
		"user/1/name":     true,
		"user/1/password": false,
	}
	r := restrict.New(perms, new(test.MockDatastore), nil)
	data := map[string]json.RawMessage{
		"user/1/name":     []byte("uwe"),
		"user/1/password": []byte("easy"),
	}
	if err := r.Restrict(1, data); err != nil {
		t.Errorf("Restrict returned unexpected error: %v", err)
	}

	if got := string(data["user/1/name"]); got != "uwe" {
		t.Errorf("data[user/1/name] = `%s`, expected `uwe`", got)
	}

	if got := data["user/1/password"]; got != nil {
		t.Errorf("data[user/1/password] = `%s`, expected nil", got)
	}
}

func TestChecker(t *testing.T) {
	perms := new(test.MockPermission)
	perms.Data = map[string]bool{
		"user/1/name":     true,
		"user/1/password": false,
	}

	called := make(map[string]bool)
	checker := map[string]restrict.Checker{
		"user/name": restrict.CheckerFunc(func(uid int, key string, value json.RawMessage) (json.RawMessage, error) {
			called[key] = true
			return []byte("touched"), nil
		}),
		"user/password": restrict.CheckerFunc(func(uid int, key string, value json.RawMessage) (json.RawMessage, error) {
			called[key] = true
			return []byte("touched"), nil
		}),
	}

	r := restrict.New(perms, new(test.MockDatastore), checker)
	data := map[string]json.RawMessage{
		"user/1/name":     []byte("uwe"),
		"user/1/password": []byte("easy"),
	}
	if err := r.Restrict(1, data); err != nil {
		t.Errorf("Restrict returned unexpected error: %v", err)
	}

	if got := string(data["user/1/name"]); got != "touched" {
		t.Errorf("data[user/1/name] = `%s`, expected `touched`", got)
	}

	if got := data["user/1/password"]; got != nil {
		t.Errorf("data[user/1/password] = `%s`, expected nil", got)
	}

	if !called["user/1/name"] {
		t.Errorf("checker for key user/1/name was not called")
	}

	if called["user/1/password"] {
		t.Errorf("checker for key user/1/password was called")
	}
}
