package main

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		name   string
		models string
		expect map[string]string
	}{
		{
			"simple",
			modelsSimple,
			map[string]string{},
		},
		{
			"relation",
			modelsRelationList,
			map[string]string{
				"model/other_ids": "other",
			},
		},
		{
			"generic-relation",
			modelsGenericRelationList,
			map[string]string{
				"model/tag_id":   "tag",
				"other/tag_id":   "tag",
				"tag/tagged_ids": "*",
			},
		},
		{
			"template relation-list",
			modelsTemplateRelationList,
			map[string]string{
				"model/other_$_ids": "other",
			},
		},
		{
			"template relation",
			modelsTemplateRelation,
			map[string]string{},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			yml := strings.ReplaceAll(tt.models, "\t", "  ")
			got, err := parse(strings.NewReader(yml))
			if err != nil {
				t.Fatalf("Got unexpected error: %v", err)
			}

			if len(got.RelationList) != len(tt.expect) {
				t.Errorf("Got %d fields, expected %d", len(got.RelationList), len(tt.expect))
			}

			for k, v := range tt.expect {
				if got.RelationList[k] != v {
					t.Errorf("got[%s] == `%s`, expected `%s`", k, got.RelationList[k], v)
				}
			}
		})
	}
}

const modelsSimple = `
model:
	id: 
		type: number
		restriction_mode: A
	name: 
		type: string
		restriction_mode: A
`

const modelsRelationList = `
model:
	other_ids:
		type: relation-list
		to: other/model_id
		restriction_mode: A

other:
	model_id:
		type: relation
		to: model/other_ids
		restriction_mode: A
`

const modelsGenericRelationList = `
model:
	tag_id:
		type: relation-list
		to: tag/tagged_ids
		restriction_mode: A

other:
	tag_id:
		type: relation-list
		to: tag/tagged_ids
		restriction_mode: A

tag:
	tagged_ids:
		type: generic-relation-list
		to:
			collection:
			- model
			- other
			field: tag_id
		restriction_mode: A
	default_id:
		type: generic-relation
		to:
			collection:
			- model
			- other
			field: tag_id
		restriction_mode: A
`

const modelsTemplateRelationList = `
model:
	other_$_ids:
		type: template
		fields:
			type: relation-list
			to: other/model_ids
		restriction_mode: A

other:
	model_ids:
		type: relation
		to: model/other_$_ids
		restriction_mode: A
`

const modelsTemplateRelation = `
model:
	other_$_id:
		type: template
		fields:
			type: relation
			to: other/model_id
		restriction_mode: A

other:
	model_id:
		type: relation
		to: model/other_$_id
		restriction_mode: A
`
