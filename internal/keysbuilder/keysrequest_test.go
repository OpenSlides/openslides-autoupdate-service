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

func TestRequestErrors(t *testing.T) {
	for _, tt := range []struct {
		name   string
		input  string
		msg    string
		fields []string
	}{
		{
			"NoField",
			`{
				"ids": [5],
				"collection": "user"
			}`,
			"no fields",
			keys(),
		},
		{
			"No Collection",
			`{
				"ids": [5],
				"fields": {"name": null}
			}`,
			"no collection",
			keys(),
		},
		{
			"no ids",
			`{
				"fields": {"name": null},
				"collection": "user"
			}`,
			"no ids",
			keys(),
		},
		{
			"Relation no collection",
			`{
				"ids": [5],
				"collection": "user",
				"fields": {
					"group_id": {
						"type": "relation",
						"fields": {"name": null}
					}
				}
			}`,
			`field "group_id": no collection`,
			keys("group_id"),
		},
		{
			"Relation No Fields",
			`	{
				"ids": [5],
				"collection": "user",
				"fields": {
					"group_id": {
						"type": "relation",
						"collection": "group"
					}
				}
			}`,
			`field "group_id": no fields`,
			keys("group_id"),
		},
		{
			"NoType",
			`	{
				"ids": [5],
				"collection": "user",
				"fields": {
					"group_id": {
						"collection": "group",
						"fields": {"name": null}
					}
				}
			}`,
			`field "group_id": no type`,
			keys("group_id"),
		},
		{
			"NoType sub",
			`	{
				"ids": [5],
				"collection": "user",
				"fields": {
					"group_id": {
						"type": "relation-list",
						"collection": "group",
						"fields": {
							"perm_ids": {
								"fields": {
									"collection": "perm",
									"fields": {"name": null}
								}
							}
						}
					}
				}
			}`,
			`field "group_id.perm_ids": no type`,
			keys("group_id", "perm_ids"),
		},
		{
			"Unknown Type",
			`	{
				"ids": [5],
				"collection": "user",
				"fields": {
					"group_id": {
						"type": "invalid-type",
						"collection": "group",
						"fields": {"name": null}
					}
				}
			}`,
			`field "group_id": unknown type invalid-type`,
			keys("group_id"),
		},
		{
			"Relation twice no fields",
			`{
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
			}`,
			`field "group_ids.perm_ids": no fields`,
			keys("group_ids", "perm_ids"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := keysbuilder.FromJSON(context.Background(), strings.NewReader(tt.input), &mockIDer{})
			if err == nil {
				t.Errorf("Expected an error, got none")
			}
			var kErr keysbuilder.ErrInvalid
			if !errors.As(err, &kErr) {
				t.Errorf("Expected err to be %T, got: %v", kErr, err)
			}
			if got := kErr.Error(); got != tt.msg {
				t.Errorf("Expected error message \"%s\", got: \"%s\"", tt.msg, got)
			}
			if fields := kErr.Fields(); !cmpSlice(fields, tt.fields) {
				t.Errorf("Expected error to be on field \"%v\", got %v", tt.fields, fields)
			}

		})
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
