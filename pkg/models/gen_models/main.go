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
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

const defURL = "https://raw.githubusercontent.com/OpenSlides/OpenSlides/openslides4-dev/docs/models.yml"

//go:embed collection.tmpl
var tmplCollection string

//go:embed all.tmpl
var tmplAll string

func main() {
	var to io.Writer = os.Stdout
	filename := outFileName(os.Args)
	if filename != "" {
		f, err := os.Create(filename)
		if err != nil {
			log.Printf("Creating file: %v", err)
			os.Exit(1)
		}
		defer f.Close()
		to = f
	}

	if err := run(to); err != nil {
		log.Printf("Error: %v", err)
		os.Exit(1)
	}
}

func outFileName(args []string) string {
	if len(args) == 1 {
		return ""
	}

	for _, a := range args[1:] {
		if a == "--" {
			continue
		}
		return a
	}
	return ""
}

func run(w io.Writer) error {
	r, err := loadDefition()
	if err != nil {
		return fmt.Errorf("loading definition: %w", err)
	}
	defer r.Close()

	collections, err := parseModelsYML(r)
	if err != nil {
		return fmt.Errorf("parsing models: %w", err)
	}

	tmpl, err := template.New("all").Parse(tmplAll)
	if err != nil {
		return fmt.Errorf("parsing template: %w", err)
	}

	var data struct {
		Collections []string
	}

	for _, c := range collections {
		rendered, err := c.template()
		if err != nil {
			return fmt.Errorf("rendering template for collection %s: %w", c.Name, err)
		}
		data.Collections = append(data.Collections, rendered)
	}

	buf := new(bytes.Buffer)

	if err := tmpl.Execute(buf, data); err != nil {
		return fmt.Errorf("executing template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("formatting output: %w", err)
	}

	if _, err := w.Write(formatted); err != nil {
		return fmt.Errorf("writing output: %w", err)
	}

	return nil
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

func parseModelsYML(r io.Reader) ([]collection, error) {
	var node yaml.Node
	if err := yaml.NewDecoder(r).Decode(&node); err != nil {
		return nil, fmt.Errorf("decoding models.yml: %w", err)
	}
	content := node.Content[0].Content

	var collections []collection
	for i := 0; i < len(content); i += 2 {
		var c collection
		if err := content[i+1].Decode(&c); err != nil {
			return nil, fmt.Errorf("decoding collection %d : %w", (i/2)+1, err)
		}
		c.Name = content[i].Value
		collections = append(collections, c)
	}
	return collections, nil
}

type collection struct {
	Name   string
	Fields []field
}

func (c *collection) GoName() string {
	return goName(c.Name)
}

func (c *collection) template() (string, error) {
	tmpl, err := template.New("collection").Parse(tmplCollection)
	if err != nil {
		return "", fmt.Errorf("parsing template: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, c); err != nil {
		return "", fmt.Errorf("executing template: %w", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return "", fmt.Errorf("formating code: %w", err)
	}
	return string(formatted), nil
}

func (c *collection) UnmarshalYAML(value *yaml.Node) error {
	for i := 0; i < len(value.Content); i += 2 {
		var f field
		if err := value.Content[i+1].Decode(&f); err != nil {
			return fmt.Errorf("decoding field %d of collection %s: %w", (i/2)+1, c.Name, err)
		}
		f.Name = value.Content[i].Value
		c.Fields = append(c.Fields, f)
	}
	return nil
}

type field struct {
	Name         string
	Type         string
	TemplateType string
}

func (f *field) GoName() string {
	return goName(f.Name)
}

func (f *field) GoType() string {
	if f.IsTemplate() {
		return "map[string]" + goType(f.TemplateType)
	}

	return goType(f.Type)
}

func (f *field) GoTemplateType() string {
	return goType(f.TemplateType)
}

func (f *field) IsTemplate() bool {
	return strings.Contains(f.Name, "$")
}

func (f *field) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		if err := value.Decode(&f.Type); err != nil {
			return fmt.Errorf("decoding field with scalar value at line %d: %w", value.Line, err)
		}
		return nil

	case yaml.MappingNode:
		var typer struct {
			Type string `yaml:"type"`
		}
		if err := value.Decode(&typer); err != nil {
			return fmt.Errorf("decoding field with map value at line %d: %w", value.Line, err)
		}
		f.Type = typer.Type

		if typer.Type == "template" {
			var fielder struct {
				Fields field `yaml:"fields"`
			}
			if err := value.Decode(&fielder); err != nil {
				return fmt.Errorf("decoding template field type: %w", err)
			}
			f.TemplateType = fielder.Fields.Type
		}
		return nil

	default:
		return fmt.Errorf("unknown value for field at line %d, expected string or mapping", value.Line)
	}
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

func goType(name string) string {
	switch name {
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
		panic(fmt.Sprintf("Unknown type: `%s`", name))
	}
}
