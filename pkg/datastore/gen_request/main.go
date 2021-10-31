// This tool generates the code needed for the request object.
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

	"github.com/OpenSlides/openslides-autoupdate-service/internal/models"
)

//go:embed value.go.tmpl
var tmplValue string

//go:embed header.go.tmpl
var tmplHeader string

//go:embed field.go.tmpl
var tmplField string

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run(w io.Writer) error {
	buf := new(bytes.Buffer)

	if err := genHeader(buf); err != nil {
		return fmt.Errorf("generate file header: %w", err)
	}

	if err := genValueTypes(buf); err != nil {
		return fmt.Errorf("generate value types: %w", err)
	}

	if err := genFieldMethods(buf); err != nil {
		return fmt.Errorf("generate field methods: %w", err)
	}

	formated, err := format.Source(buf.Bytes())
	if err != nil {
		if _, err := w.Write(buf.Bytes()); err != nil {
			return fmt.Errorf("writing output: %w", err)
		}
		return fmt.Errorf("formating code: %w", err)
	}

	if _, err := w.Write(formated); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}
	return nil
}

func genHeader(buf *bytes.Buffer) error {
	tmpl, err := template.New("header.go").Parse(tmplHeader)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	if err := tmpl.Execute(buf, nil); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	return nil
}

func genValueTypes(buf *bytes.Buffer) error {
	typesToGo := map[string]string{
		"ValueInt":         "int",
		"ValueString":      "string",
		"ValueBool":        "bool",
		"ValueFloat":       "float32",
		"ValueJSON":        "json.RawMessage",
		"ValueIntSlice":    "[]int",
		"ValueStringSlice": "[]string",
		"ValueIDSlice":     "[]int",
	}

	var types []string
	for k := range typesToGo {
		types = append(types, k)
	}
	sort.Strings(types)

	tmpl, err := template.New("value.go").Parse(tmplValue)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	for _, name := range types {
		data := struct {
			TypeName string
			GoType   string
			Zero     string
		}{
			name,
			typesToGo[name],
			zeroValue(typesToGo[name]),
		}
		if err := tmpl.Execute(buf, data); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}
	}
	return nil
}

// zeroValue returns the zero value for a go type
func zeroValue(t string) string {
	switch t {
	case "int", "float32":
		return "0"
	case "string":
		return `""`
	case "bool":
		return "false"
	case "json.RawMessage", "[]int", "[]string":
		return "nil"
	}
	return "unknown type " + t
}

func genFieldMethods(buf *bytes.Buffer) error {
	r, err := loadDefition()
	if err != nil {
		log.Fatalf("Can not load models defition: %v", err)
	}
	defer r.Close()

	fields, err := parse(r)
	if err != nil {
		log.Fatalf("Can not parse model definition: %v", err)
	}

	tmpl, err := template.New("field.go").Parse(tmplField)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	for _, field := range fields {
		if err := tmpl.Execute(buf, field); err != nil {
			return fmt.Errorf("executing template: %w", err)
		}
	}

	return nil
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
			f.ValueType = valueType(modelField.Type)
			f.Collection = firstLower(goName(collectionName))
			f.FQField = collectionName + "/%d/" + fieldName
			f.Required = modelField.Required

			if modelField.Type == "relation" || modelField.Type == "generic-relation" {
				f.SingleRelation = true
			}

			if strings.Contains(fieldName, "$") {
				f.TemplateAttr = "replacement"
				f.TemplateAttrType = "string"
				f.TemplateFQField = collectionName + "/%d/" + strings.Replace(fieldName, "$", "$%s", 1)
				f.ValueType = valueType(modelField.Template.Fields.Type)

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
	ValueType        string
	Collection       string
	FQField          string
	TemplateFQField  string
	TemplateAttr     string
	TemplateAttrType string
	Required         bool
	SingleRelation   bool
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

func valueType(modelsType string) string {
	switch modelsType {
	case "number", "relation", "timestamp":
		return "ValueInt"

	case "string", "HTMLStrict", "color", "HTMLPermissive", "generic-relation", "template", "decimal(6)":
		return "ValueString"

	case "boolean":
		return "ValueBool"

	case "float":
		return "ValueFloat"

	case "relation-list", "number[]":
		return "ValueIntSlice"

	case "JSON":
		return "ValueJSON"

	case "string[]", "generic-relation-list":
		return "ValueStringSlice"

	default:
		panic(fmt.Sprintf("Unknown type %q", modelsType))
	}
}
