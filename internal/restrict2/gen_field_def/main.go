// This tool generates the list of related fields in the file def.go.
// To call it, just call "go generate ./..." in the root folder of the repository
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"text/template"

	models "github.com/OpenSlides/openslides-models-to-go"
)

const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/models.yml"

func main() {
	r, err := loadDefition()
	if err != nil {
		log.Fatalf("Can not load models defition: %v", err)
	}
	defer r.Close()

	relations, restrictions, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, relations, restrictions); err != nil {
		log.Fatalf("Can not write result: %v", err)
	}
}

func loadDefition() (io.ReadCloser, error) {
	r, err := http.Get(defURL)
	if err != nil {
		return nil, fmt.Errorf("request defition: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned status %s", r.Status)
	}
	return r.Body, nil
}

// parse returns all relation-list and generic-relation-list fields and where
// they point to.
func parse(r io.Reader) (relations map[string]string, restrictions map[string]string, err error) {
	inData, err := models.Unmarshal(r)
	if err != nil {
		return nil, nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	relations = make(map[string]string)
	restrictions = make(map[string]string)
	for modelName, model := range inData {
		for fieldName, field := range model.Fields {
			fqfield := fmt.Sprintf("%s/%s", modelName, fieldName)
			restrictions[fqfield] = field.RestrictionMode()

			r := field.Relation()

			if r == nil || !r.List() {
				continue
			}

			collection := "*"
			if _, ok := r.(*models.AttributeRelation); ok {
				collection = r.ToCollections()[0].Collection
			}

			relations[fqfield] = collection
		}
	}

	return relations, restrictions, nil
}

const tpl = `// Code generated with models.txt DO NOT EDIT.
package restrict


// RelationLists is list from all relation-list and generic-relation-list the
// model where it directs to. generic-relation-list habe '*' als value. The list
// contains also all template-fields that contain relation-list and
// geneeric-relation-lists.
//
// The map is automaticly created from the models.yml file.
var RelationLists = map[string]string{
	{{- range $key, $value := .Relations}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}

// RestrictionModes are all fields to there restriction_mode.
var RestrictionModes = map[string]string{
	{{- range $key, $value := .Restrictions}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}
`

func writeFile(w io.Writer, relations map[string]string, restrictions map[string]string) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	data := map[string]interface{}{
		"Relations":    relations,
		"Restrictions": restrictions,
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, data); err != nil {
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
