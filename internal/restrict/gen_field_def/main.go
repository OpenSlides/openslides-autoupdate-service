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
	"strings"
	"text/template"

	models "github.com/OpenSlides/openslides-models-to-go"
)

const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/master/docs/models.yml"

func main() {
	r, err := loadDefition()
	if err != nil {
		log.Fatalf("Can not load models defition: %v", err)
	}
	defer r.Close()

	td, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, td); err != nil {
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

type templateData struct {
	Relation            map[string]string
	RelationList        map[string]string
	GenericRelation     map[string]string
	GenericRelationList map[string]string
	Restrictions        map[string]string
}

// parse returns all relation-list and generic-relation-list fields and where
// they point to.
func parse(r io.Reader) (td templateData, err error) {
	inData, err := models.Unmarshal(r)
	if err != nil {
		return td, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	td.Relation = make(map[string]string)
	td.RelationList = make(map[string]string)
	td.GenericRelation = make(map[string]string)
	td.GenericRelationList = make(map[string]string)
	td.Restrictions = make(map[string]string)
	for modelName, model := range inData {
		for fieldName, field := range model.Fields {
			collectionField := fmt.Sprintf("%s/%s", modelName, fieldName)
			reducedKey := reduceKey(collectionField)
			td.Restrictions[reducedKey] = field.RestrictionMode()
			td.Restrictions[collectionField] = field.RestrictionMode()

			r := field.Relation()

			if r == nil {
				continue
			}

			switch v := r.(type) {
			case *models.AttributeRelation:
				to := v.ToCollections()[0].Collection + "/" + v.ToCollections()[0].ToField.Name
				if r.List() {
					td.RelationList[reducedKey] = to
				} else {
					td.Relation[reducedKey] = to
				}

			case *models.AttributeGenericRelation:
				to := v.ToCollections()[0].ToField.Name
				if r.List() {
					td.GenericRelationList[reducedKey] = to
				} else {
					td.GenericRelation[reducedKey] = to
				}

			default:
				return td, fmt.Errorf("unknown type %t for field.Relation", v)
			}

		}
	}

	return td, nil
}

// reduceKey returns key for normal keys and the prefix for template keys.
func reduceKey(key string) string {
	i := strings.IndexByte(key, '$')
	if i < 0 || i == len(key)-1 {
		// Normal field or $ at the end
		return key
	}

	return key[:i+1]
}

const tpl = `// Code generated with models.yml DO NOT EDIT.
package restrict

var relationFields = map[string]string{
	{{- range $key, $value := .Relation}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}

var relationListFields = map[string]string{
	{{- range $key, $value := .RelationList}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}

var genericRelationFields = map[string]string{
	{{- range $key, $value := .GenericRelation}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}

var genericRelationListFields = map[string]string{
	{{- range $key, $value := .GenericRelationList}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}

// restrictionModes are all fields to there restriction_mode.
var restrictionModes = map[string]string{
	{{- range $key, $value := .Restrictions}}
	"{{$key}}": "{{$value}}",
	{{- end}}
}
`

func writeFile(w io.Writer, td templateData) error {
	t := template.New("t")
	t, err := t.Parse(tpl)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)

	if err := t.Execute(buf, td); err != nil {
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
