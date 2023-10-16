package main

import (
	"bytes"
	"cmp"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"text/template"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/models"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

func main() {
	r, err := loadDefition()
	if err != nil {
		log.Fatalf("Can not load models defition: %v", err)
	}
	defer r.Close()

	p, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, p); err != nil {
		log.Fatalf("Can not write result: %v", err)
	}
}

func loadDefition() (io.ReadCloser, error) {
	r, err := http.Get(models.URLModelsYML())
	if err != nil {
		return nil, fmt.Errorf("request defition: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned status %s", r.Status)
	}
	return r.Body, nil
}

type collectionField struct {
	Collection string
	Field      string
}

type parseResult struct {
	collectionFields      []collectionField
	collectionModes       []collectionField
	collectionFieldToMode []int
	relationType          []int
	relationTo            []int
}

func parse(r io.Reader) (parseResult, error) {
	inData, err := models.Unmarshal(r)
	if err != nil {
		return parseResult{}, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	var cfs []collectionField

	modefields := set.New[collectionField]()
	fieldToMode := make(map[collectionField]collectionField)
	fieldRelationType := make(map[collectionField]int)
	fieldRelationTo := make(map[collectionField]collectionField)

	for collection, collInfo := range inData {
		for field, fieldInfo := range collInfo.Fields {
			cf := collectionField{
				Collection: collection,
				Field:      field,
			}
			cm := collectionField{
				Collection: collection,
				Field:      fieldInfo.RestrictionMode(),
			}

			cfs = append(cfs, cf)

			modefields.Add(cm)
			fieldToMode[cf] = cm

			relation := fieldInfo.Relation()

			if relation == nil {
				continue
			}

			switch v := relation.(type) {
			case *models.AttributeRelation:

				fieldRelationTo[cf] = collectionField{
					Collection: v.ToCollections()[0].Collection,
					Field:      v.ToCollections()[0].ToField.Name,
				}

				if relation.List() {
					fieldRelationType[cf] = 2
				} else {
					fieldRelationType[cf] = 1
				}

			case *models.AttributeGenericRelation:
				if relation.List() {
					fieldRelationType[cf] = 4
				} else {
					fieldRelationType[cf] = 3
				}

			default:
				return parseResult{}, fmt.Errorf("unknown type %t for field.Relation", v)
			}

		}
	}

	// Add internal fields that are required by the restricter.
	modefields.Add(collectionField{
		Collection: "poll",
		Field:      "MANAGE",
	})

	slices.SortFunc(cfs, func(a, b collectionField) int {
		if x := cmp.Compare(a.Collection, b.Collection); x != 0 {
			return x
		}
		return cmp.Compare(a.Field, b.Field)
	})

	mfs := modefields.List()
	slices.SortFunc(mfs, func(a, b collectionField) int {
		if x := cmp.Compare(a.Collection, b.Collection); x != 0 {
			return x
		}
		return cmp.Compare(a.Field, b.Field)
	})

	cf2cm := make([]int, len(cfs))
	rt := make([]int, len(cfs))
	r2 := make([]int, len(cfs))
	for i, cf := range cfs {
		idx := slices.Index(mfs, fieldToMode[cf])
		if idx == -1 {
			return parseResult{}, fmt.Errorf("can not find mode for %s", cf)
		}
		cf2cm[i] = idx + 1

		rt[i] = fieldRelationType[cf]

		idx = slices.Index(cfs, fieldRelationTo[cf])
		if idx != -1 {
			r2[i] = idx + 2
		}
	}

	return parseResult{
		collectionFields:      cfs,
		collectionModes:       mfs,
		collectionFieldToMode: cf2cm,
		relationType:          rt,
		relationTo:            r2,
	}, nil
}

const tpl = `// Code generated with models.yml DO NOT EDIT.
package dskey

var collectionFields = [...]collectionField{
	{"invalid", "key"},
	{"_meta", "update"},
	{{- range $cf := .CollectionFields}}
		{"{{$cf.Collection}}", "{{$cf.Field}}"},
	{{- end}}
}

func collectionFieldToID(cf string) int{
	switch cf{
	case "_meta/update":
		return 1
	{{- range $idx, $cf := .CollectionFields}}
	case "{{$cf.Collection}}/{{$cf.Field}}":
		return {{add2 $idx}}
	{{- end}}
	default: 
		return -1
	}
}

var collectionFieldToMode = [...]int{
	0,
	0,
	{{- range $idx := .ModeIndex}}
		{{$idx}},
	{{- end}}
}

var relationType = [...]Relation{
	RelationNone,
	RelationNone,
	{{- range $type := .RelationType}}
		{{- if eq $type 0}}
			RelationNone,
		{{- else if eq $type 1}}
			RelationSingle,
		{{- else if eq $type 2}}
			RelationList,
		{{- else if eq $type 3}}
			RelationGenericSingle,
		{{- else if eq $type 4}}
			RelationGenericList,
		{{- end}}
	{{- end}}
}

var relationTo = [...]int{
	0,
	0,
	{{- range $idx := .RelationTo}}
		{{$idx}},
	{{- end}}
}

var collectionModeFields = [...]collectionMode{
	{"invalid", "mode"},
	{{- range $cf := .CollectionModes}}
		{"{{$cf.Collection}}", "{{$cf.Field}}"},
	{{- end}}
}

func collectionModeToID(cf string) int{
	switch cf{
	{{- range $idx, $cf := .CollectionModes}}
	case "{{$cf.Collection}}/{{$cf.Field}}":
		return {{add1 $idx}}
	{{- end}}
	default: 
		return -1
	}
}
`

func writeFile(w io.Writer, p parseResult) error {
	t := template.New("t").Funcs(template.FuncMap{
		"add1": func(num int) int {
			return num + 1
		},

		"add2": func(num int) int {
			return num + 2
		},
	})
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)

	templateData := struct {
		CollectionFields []collectionField
		CollectionModes  []collectionField
		ModeIndex        []int
		RelationType     []int
		RelationTo       []int
	}{
		CollectionFields: p.collectionFields,
		CollectionModes:  p.collectionModes,
		ModeIndex:        p.collectionFieldToMode,
		RelationType:     p.relationType,
		RelationTo:       p.relationTo,
	}

	if err := t.Execute(buf, templateData); err != nil {
		return fmt.Errorf("writing template: %w", err)
	}

	formated, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formating code: %w", err)
	}

	if _, err := w.Write(formated); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	return nil
}
