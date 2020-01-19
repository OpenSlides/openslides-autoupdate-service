package keysrequest_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate/keysrequest"
)

func TestJSONValid(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"motion_ids": {
				"collection": "motion",
				"fields": {"name": null}
			}
		}
	}
	`)
	if _, err := keysrequest.FromJSON(json); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestJSONInvalid(t *testing.T) {
	json := strings.NewReader(`{5`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var errJSON keysrequest.ErrJSON
	if !errors.As(err, &errJSON) {
		t.Errorf("Expected error to be of type ErrJSON, got: %v", err)
	}
}

func TestJSONSingleID(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": 5,
		"collection": "user",
		"fields": {"name": null}
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var errJSON keysrequest.ErrJSON
	if !errors.As(err, &errJSON) {
		t.Errorf("Expected error to be of type ErrJSON, got: %v", err)
	}
}

func TestJSONNoField(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user"
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "no fields"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
}

func TestJSONNoCollection(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"fields": {"name": null}
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "no collection"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
}

func TestJSONNoIDs(t *testing.T) {
	json := strings.NewReader(`
	{
		"fields": {"name": null},
		"collection": "user"
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "no ids"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
}

func TestJSONNoSuffix(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"name": {
				"collection": "username",
				"fields": {"inner_name": null}
			}
		}
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "field \"name\": relation but no _id or _ids suffix"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
	if fields := kErr.Fields(); len(fields) == 0 || fields[0] != "name" {
		t.Errorf("Expected error to be on field name, got: %v", fields)
	}
}

func TestJSONSuffixNoFields(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": null,
			"note_id": null
		}
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestJSONInnerNoCollection(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_id": {
				"fields": {"name": null}
			}
		}
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "field \"group_id\": no collection"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
	if fields := kErr.Fields(); len(fields) == 0 || fields[0] != "group_id" {
		t.Errorf("Expected error to be on field \"name\"")
	}
}

func TestJSONInnerNoFields(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_id": {
				"collection": "group"
			}
		}
	}
	`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "field \"group_id\": no fields"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
	if fields := kErr.Fields(); len(fields) == 0 || fields[0] != "group_id" {
		t.Errorf("Expected error to be on field \"name\"")
	}
}

func TestJSONInnerTwiceNoFields(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": {
				"collection": "group",
				"fields": {
					"perm_ids": {
						"collection": "perm"
					}
				}
			}
		}
	}`)
	_, err := keysrequest.FromJSON(json)
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "field \"group_ids.perm_ids\": no fields"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
	if fields := kErr.Fields(); len(fields) != 2 || fields[0] != "group_ids" || fields[1] != "perm_ids" {
		t.Errorf("Expected error to be on field \"group_ids.perm_ids\", got: %v", fields)
	}
}

func TestManyFromJSON(t *testing.T) {
	json := strings.NewReader(`[
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": {
				"collection": "group",
				"fields": {
					"name": null
				}
			}
		}
	},
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"name": null
		}
	}]`)
	_, err := keysrequest.ManyFromJSON(json)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestManyFromJSONInvalidJSON(t *testing.T) {
	json := strings.NewReader(`[
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": {
				"collection": "group",
				"fields": {
					"name": null
				}
			}
		}
	},
	{
		"ids": [5],
		"collection": "user",
		"fi
	}]`)
	_, err := keysrequest.ManyFromJSON(json)
	if err == nil {
		t.Error("Expected ManyFromJSON() to return an error, got not")
	}
	var errJSON keysrequest.ErrJSON
	if !errors.As(err, &errJSON) {
		t.Errorf("Expected error to be of type ErrJSON, got: %v", err)
	}
}

func TestManyFromJSONInvalidKeysRequest(t *testing.T) {
	json := strings.NewReader(`[
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": {
				"collection": "group",
				"fields": {
					"name": null
				}
			}
		}
	},
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": {
				"fields": {
					"name": null
				}
			}
		}
	}]`)
	_, err := keysrequest.ManyFromJSON(json)
	if err == nil {
		t.Error("Expected ManyFromJSON() to return an error, got not")
	}
	var kErr keysrequest.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected error to be of type ErrInvalid, got: %v", err)
	}
}
