package main

import (
	"strings"
	"testing"
)

func TestCollection(t *testing.T) {
	const yaml = `
user:
  id: number
  user_name: string
  active: boolean
`
	// In this template @ is replaced by `
	const expected = `
type User struct {
	ID       int    @json:"id"@
	UserName string @json:"user_name"@
	Active   bool   @json:"active"@
}

func LoadUser(ctx context.Context, ds Getter, id int) (*User, error) {
	fields := []string{
		"id",
		"user_name",
		"active",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("user/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c User
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.UserName); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Active); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	return &c, nil
}`

	collections, err := parseModelsYML(strings.NewReader(yaml))
	if err != nil {
		t.Fatalf("Building collections: %v", err)
	}

	got, err := collections[0].template()
	if err != nil {
		t.Fatalf("Parsing template: %v", err)
	}

	got = strings.TrimSpace(got)

	if e := strings.TrimSpace(strings.ReplaceAll(expected, "@", "`")); got != e {
		t.Errorf("Got %d bytes:\n%s\nexpected %d bytes:\n%s", len(got), got, len(e), e)
	}
}

func TestTemplateField(t *testing.T) {
	const yaml = `
model:
  field_$:
    type: template
    fields: string
  relation_$_ids:
    type: template
    fields:
      type: relation-list
`
	// In this template @ is replaced by `
	const expected = `
type Model struct {
	Field       map[string]string @json:"field_$"@
	RelationIDs map[string][]int  @json:"relation_$_ids"@
}

func LoadModel(ctx context.Context, ds Getter, id int) (*Model, error) {
	fields := []string{
		"field_$",
		"relation_$_ids",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("model/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Model
	if values[0] != nil {
		var repl []string
		if err := json.Unmarshal(values[0], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[0], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[0], err)
		}

		data := make(map[string]string, len(repl))
		for i, r := range repl {
			var decoded string
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.Field = data
	}
	if values[1] != nil {
		var repl []string
		if err := json.Unmarshal(values[1], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[1], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[1], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.RelationIDs = data
	}
	return &c, nil
}`

	collections, err := parseModelsYML(strings.NewReader(yaml))
	if err != nil {
		t.Fatalf("Building collections: %v", err)
	}

	got, err := collections[0].template()
	if err != nil {
		t.Fatalf("Parsing template: %v", err)
	}

	got = strings.TrimSpace(got)

	if e := strings.TrimSpace(strings.ReplaceAll(expected, "@", "`")); got != e {
		t.Errorf("Got %d bytes:\n%s\nexpected %d bytes:\n%s", len(got), got, len(e), e)
	}
}
