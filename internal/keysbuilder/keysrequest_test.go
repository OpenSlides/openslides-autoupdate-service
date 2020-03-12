package keysbuilder_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

func TestJSONValid(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"motion_ids": {
				"type": "relation-list",
				"collection": "motion",
				"fields": {"name": null}
			}
		}
	}
	`)
	if _, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{}); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestJSONInvalid(t *testing.T) {
	json := strings.NewReader(`{5`)
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var errJSON keysbuilder.ErrJSON
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
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var errJSON keysbuilder.ErrJSON
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
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
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
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
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
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "no ids"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
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
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}

func TestJSONRelationNoCollection(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_id": {
				"type": "relation",
				"fields": {"name": null}
			}
		}
	}
	`)
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
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

func TestJSONRelationNoFields(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_id": {
				"type": "relation",
				"collection": "group"
			}
		}
	}
	`)
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
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

func TestJSONNoType(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_id": {
				"collection": "group",
				"fields": {"name": null}
			}
		}
	}
	`)
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "field \"group_id\": no type"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
	if fields := kErr.Fields(); len(fields) == 0 || fields[0] != "group_id" {
		t.Errorf("Expected error to be on field \"name\"")
	}
}

func TestJSONUnknwonType(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_id": {
				"type": "invalid-type",
				"collection": "group",
				"fields": {"name": null}
			}
		}
	}
	`)
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected err to be %T, got: %v", kErr, err)
	}
	expect := "field \"group_id\": unknown type invalid-type"
	if got := kErr.Error(); got != expect {
		t.Errorf("Expected error message \"%s\", got: \"%s\"", expect, got)
	}
	if fields := kErr.Fields(); len(fields) == 0 || fields[0] != "group_id" {
		t.Errorf("Expected error to be on field \"name\"")
	}
}

func TestJSONRelationTwiceNoFields(t *testing.T) {
	json := strings.NewReader(`
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": {
				"type": "relation-list",
				"collection": "group",
				"fields": {
					"perm_ids": {
						"type": "relation-list",
						"collection": "perm"
					}
				}
			}
		}
	}`)
	_, err := keysbuilder.FromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	var kErr keysbuilder.ErrInvalid
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
				"type": "relation-list",
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
	_, err := keysbuilder.ManyFromJSON(context.Background(), json, &mockIDer{})
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
				"type": "relation-list",
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
	_, err := keysbuilder.ManyFromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Error("Expected ManyFromJSON() to return an error, got not")
	}
	var errJSON keysbuilder.ErrJSON
	if !errors.As(err, &errJSON) {
		t.Errorf("Expected error to be of type ErrJSON, got: %v", err)
	}
}

func TestManyFromJSONInvalidInput(t *testing.T) {
	json := strings.NewReader(`[
	{
		"ids": [5],
		"collection": "user",
		"fields": {
			"group_ids": {
				"type": "relation-list",
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
				"type": "relation-list",
				"fields": {
					"name": null
				}
			}
		}
	}]`)
	_, err := keysbuilder.ManyFromJSON(context.Background(), json, &mockIDer{})
	if err == nil {
		t.Error("Expected ManyFromJSON() to return an error, got not")
	}
	var kErr keysbuilder.ErrInvalid
	if !errors.As(err, &kErr) {
		t.Errorf("Expected error to be of type ErrInvalid, got: %v", err)
	}
}
