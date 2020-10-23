package restrict_test

import (
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/restrict"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestRelationChecker(t *testing.T) {
	rl := map[string]string{
		"model/relation_ids":           "otherModel",
		"model/generic_relation_ids":   "*",
		"model/template_$_ids":         "otherModel",
		"model/generic_template_$_ids": "*",
	}

	permer := new(test.MockPermission)

	checker := restrict.RelationChecker(rl, permer)

	t.Run("Check keys", func(t *testing.T) {
		if len(checker) != 6 {
			t.Errorf("Got %d checker, expected 6. (Template fields create two checkers.)", len(checker))
		}

		keys := []string{
			"model/relation_ids",
			"model/generic_relation_ids",
			"model/template_$_ids",
			"model/generic_template_$_ids",
			"model/template_",
			"model/generic_template_",
		}
		for _, key := range keys {
			if _, ok := checker[key]; !ok {
				t.Fatalf("checker does not contain key %s", key)
			}
		}
	})

	t.Run("relation list", func(t *testing.T) {
		permer.Data = map[string]bool{
			"otherModel/1": true,
			"otherModel/2": false,
		}

		v, err := checker["model/relation_ids"].Check(1, "model/1/relation_ids", []byte("[1,2]"))

		if err != nil {
			t.Fatalf("Check returned an error: %v", err)
		}

		if got := string(v); got != "[1]" {
			t.Fatalf("Check returned `%s`, expected `[1]`", got)
		}
	})

	t.Run("generic relation list", func(t *testing.T) {
		permer.Data = map[string]bool{
			"foo/1":       true,
			"other_foo/2": false,
		}

		v, err := checker["model/generic_relation_ids"].Check(1, "model/1/generic_relation_ids", []byte(`["foo/1","other_foo/2"]`))

		if err != nil {
			t.Errorf("Check returned an error: %v", err)
		}

		if got := string(v); got != `["foo/1"]` {
			t.Errorf("Check returned `%s`, expected `[\"foo/1\"]`", got)
		}
	})

	t.Run("template field", func(t *testing.T) {
		permer.Data = map[string]bool{
			"model/1/template_$1_ids": true,
			"model/1/template_$2_ids": false,
		}

		v, err := checker["model/template_$_ids"].Check(1, "model/1/template_$_ids", []byte(`["1","2"]`))

		if err != nil {
			t.Errorf("Check returned an error: %v", err)
		}

		if got := string(v); got != `["1"]` {
			t.Errorf("Check returned `%s`, expected `[\"1\"]`", got)
		}
	})

	t.Run("template containing relation field", func(t *testing.T) {
		permer.Data = map[string]bool{
			"otherModel/1": true,
			"otherModel/2": false,
		}

		v, err := checker["model/template_"].Check(1, "model/1/template_$1_ids", []byte(`[1,2]`))

		if err != nil {
			t.Errorf("Check returned an error: %v", err)
		}

		if got := string(v); got != "[1]" {
			t.Fatalf("Check returned `%s`, expected `[1]`", got)
		}
	})

	t.Run("template containing generic relation field", func(t *testing.T) {
		permer.Data = map[string]bool{
			"foo/1":       true,
			"other_foo/2": false,
		}

		v, err := checker["model/generic_template_"].Check(1, "model/1/generic_template_$1_ids", []byte(`["foo/1","other_foo/2"]`))

		if err != nil {
			t.Errorf("Check returned an error: %v", err)
		}

		if got := string(v); got != `["foo/1"]` {
			t.Errorf("Check returned `%s`, expected `[\"foo/1\"]`", got)
		}
	})
}
