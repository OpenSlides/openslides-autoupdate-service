// This tool generates the list of fields.
//
// To call it, just call "go generate ./..." in the root folder of the repository
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"text/template"

	models "github.com/OpenSlides/openslides-models-to-go"
)

const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/models.yml"

//go:embed fields.go.tmpl
var tmplFields string

//go:embed normal_field.go.tmpl
var tmplNormalField string

//go:embed structured_field.go.tmpl
var tmplStructuredField string

func main() {
	r, err := loadDefition()
	if err != nil {
		log.Fatalf("Can not load models defition: %v", err)
	}
	defer r.Close()

	fields, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	if err := writeFile(os.Stdout, fields); err != nil {
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

// parse returns all fields from the models.yml with there go-type as string.
func parse(r io.Reader) ([]field, error) {
	inData, err := models.Unmarshal(r)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling models.yml: %w", err)
	}

	var fields []field
	for collectionName, collection := range inData {
		for fieldName, modelField := range collection.Fields {
			f := field{}
			f.Name = collectionName + "/" + fieldName
			f.GoName = goName(collectionName) + "_" + goName(fieldName)
			f.GoType = goType(modelField.Type)
			f.Collection = firstLower(goName(collectionName))
			f.FQField = collectionName + "/%d/" + fieldName

			if strings.Contains(fieldName, "$") {
				f.TemplateAttr = "replacement"
				f.TemplateAttrType = "string"
				f.TemplateFQField = collectionName + "/%d/" + strings.Replace(fieldName, "$", "$%s", 1)
				f.GoType = goType(modelField.Template.Fields.Type)

				if modelField.Template.Replacement != "" {
					f.TemplateAttr = modelField.Template.Replacement + "ID"
					f.TemplateAttrType = "int"
					f.TemplateFQField = collectionName + "/%d/" + strings.Replace(fieldName, "$", "$%d", 1)
				}
			}

			fields = append(fields, f)
		}
	}

	// TODO: fix models-to-go to return fields in input order.
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].GoName < fields[j].GoName
	})

	return fields, nil
}

type field struct {
	Name             string
	GoName           string
	GoType           string
	Collection       string
	FQField          string
	TemplateFQField  string
	TemplateAttr     string
	TemplateAttrType string
}

func (f field) template() (string, error) {
	tmpl, err := template.New("collection").Parse(tmplNormalField)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	if f.TemplateAttr != "" {
		tmpl, err = template.New("collection").Parse(tmplStructuredField)
		if err != nil {
			return "", fmt.Errorf("parsing template: %w", err)
		}
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, f); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return "", fmt.Errorf("formating code: %w", err)
	}
	return string(formatted), nil
}

func writeFile(w io.Writer, fields []field) error {
	tmpl, err := template.New("fields.go").Parse(tmplFields)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	var data struct {
		Fields []string
	}

	for _, f := range fields {
		t, err := f.template()
		if err != nil {
			return fmt.Errorf("template for field %q: %w", f.Name, err)
		}

		data.Fields = append(data.Fields, t)
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
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

func goName(name string) string {
	if name == "id" {
		return "ID"
	}

	name = strings.ReplaceAll(name, "_$", "")

	parts := strings.Split(name, "_")
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}
	name = strings.Join(parts, "")

	name = strings.ReplaceAll(name, "Id", "ID")
	return name
}

func firstLower(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

func goType(modelsType string) string {
	switch modelsType {
	case "number", "decimal(6)", "relation", "timestamp":
		return "int"

	case "string", "HTMLStrict", "color", "HTMLPermissive", "generic-relation", "template":
		return "string"

	case "boolean":
		return "bool"

	case "float":
		return "float32"

	case "relation-list", "number[]":
		return "[]int"

	case "JSON":
		return "json.RawMessage"

	case "string[]", "generic-relation-list":
		return "[]string"

	default:
		panic(fmt.Sprintf("Unknown type %q", modelsType))
	}
}
