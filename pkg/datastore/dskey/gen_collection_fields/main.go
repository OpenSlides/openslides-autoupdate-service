package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
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

	collectionFields, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, collectionFields); err != nil {
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

func parse(r io.Reader) ([]collectionField, error) {
	inData, err := models.Unmarshal(r)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	var result []collectionField

	modefields := set.New[collectionField]()

	for collection, collInfo := range inData {
		for field, fieldInfo := range collInfo.Fields {
			result = append(result, collectionField{
				Collection: collection,
				Field:      field,
			})

			modefields.Add(collectionField{
				Collection: collection,
				Field:      fieldInfo.RestrictionMode(),
			})
		}
	}

	// TODO: Save mode fields in own datastructure
	for _, cf := range modefields.List() {
		result = append(result, cf)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Collection == result[j].Collection {
			return result[i].Field < result[j].Field
		}
		return result[i].Collection < result[j].Collection
	})

	return result, nil
}

const tpl = `// Code generated with models.yml DO NOT EDIT.
package dskey

var collectionFields = [...]collectionField{
	{"invalid", "key"},
	{{- range $cf := .CollectionFields}}
		{"{{$cf.Collection}}", "{{$cf.Field}}"},
	{{- end}}
}

func collectionFieldToID(cf string) int{
	switch cf{
	{{- range $idx, $cf := .CollectionFields}}
	case "{{$cf.Collection}}/{{$cf.Field}}":
		return {{add1 $idx}}
	{{- end}}
	default: 
		return -1
	}
}
`

func writeFile(w io.Writer, collectionFields []collectionField) error {
	t := template.New("t").Funcs(template.FuncMap{
		"add1": func(num int) int {
			return num + 1
		},
	})
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)

	templateData := struct {
		CollectionFields []collectionField
	}{
		CollectionFields: collectionFields,
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
