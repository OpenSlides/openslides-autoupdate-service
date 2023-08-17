package models_test

import (
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/models"
)

const yamlWithRelation = `---
model:
	other_id:
		type: relation
		to: not_existing/field
		restriction_mode: A
		required: true
	other_ids:
		type: relation-list
		to: other/name
		restriction_mode: B
other:
	name:
		type: string
		restriction_mode: A
`

func TestUnmarshal(t *testing.T) {
	for _, tt := range []struct {
		name string
		yaml string
	}{
		{
			"With Relation",
			yamlWithRelation,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			yml := strings.ReplaceAll(tt.yaml, "\t", " ")
			_, err := models.Unmarshal(strings.NewReader(yml))
			if err != nil {
				t.Errorf("Can not unmarshal yaml: %v", err)
			}
		})
	}
}

func TestRelation(t *testing.T) {
	yml := strings.ReplaceAll(yamlWithRelation, "\t", " ")
	got, err := models.Unmarshal(strings.NewReader(yml))
	if err != nil {
		t.Errorf("Can not unmarshal yml: %v", err)
	}

	if got["model"].Fields["other_id"].Relation().List() {
		t.Errorf("model/other_id is a list")
	}

	if !got["model"].Fields["other_ids"].Relation().List() {
		t.Errorf("model/other_ids is not a list")
	}
}

func TestRestrictionMode(t *testing.T) {
	yml := strings.ReplaceAll(yamlWithRelation, "\t", " ")
	got, err := models.Unmarshal(strings.NewReader(yml))
	if err != nil {
		t.Errorf("Can not unmarshal yml: %v", err)
	}

	if v := got["model"].Fields["other_id"].RestrictionMode(); v != "A" {
		t.Errorf("Fireld moddel/other_id.RestrictionMode == %q, expected \"A\"", v)
	}

	if v := got["model"].Fields["other_ids"].RestrictionMode(); v != "B" {
		t.Errorf("Fireld model/other_ids.RestrictionMode == %q, expected \"B\"", v)
	}

	if v := got["other"].Fields["name"].RestrictionMode(); v != "A" {
		t.Errorf("Fireld model/name.RestrictionMode == %q, expected \"A\"", v)
	}
}

func TestRequired(t *testing.T) {
	yml := strings.ReplaceAll(yamlWithRelation, "\t", " ")
	got, err := models.Unmarshal(strings.NewReader(yml))
	if err != nil {
		t.Errorf("Can not unmarshal yml: %v", err)
	}

	if v := got["model"].Fields["other_id"].Required; !v {
		t.Errorf("Field model/other_id not required, expected true")
	}

	if v := got["model"].Fields["other_ids"].Required; v {
		t.Errorf("Field model/other_id is required, expected false")
	}
}
