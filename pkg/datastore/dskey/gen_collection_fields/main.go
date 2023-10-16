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

	collectionFields, collectionModes, cf2cm, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, collectionFields, collectionModes, cf2cm); err != nil {
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

func parse(r io.Reader) ([]collectionField, []collectionField, []int, error) {
	inData, err := models.Unmarshal(r)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	var cfs []collectionField

	modefields := set.New[collectionField]()
	fieldToMode := make(map[collectionField]collectionField)

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
	for i, cf := range cfs {
		idx := slices.Index(mfs, fieldToMode[cf])
		if idx == -1 {
			return nil, nil, nil, fmt.Errorf("can not find mode for %s", cf)
		}
		cf2cm[i] = idx + 1
	}

	return cfs, mfs, cf2cm, nil
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

var collectionFieldToMode = [...]int{
	-1,
	-1,
	{{- range $idx := .ModeIndex}}
		{{$idx}},
	{{- end}}
}
`

func writeFile(w io.Writer, collectionFields []collectionField, collectionModes []collectionField, cf2cm []int) error {
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
	}{
		CollectionFields: collectionFields,
		CollectionModes:  collectionModes,
		ModeIndex:        cf2cm,
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
