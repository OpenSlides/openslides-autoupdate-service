package models

import (
	"bufio"
	_ "embed" // needed for embeding
	"fmt"
	"io"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed models-version
var version string

func getVersion() string {
	scanner := bufio.NewScanner(strings.NewReader(version))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "//") || line == "" {
			continue
		}
		return line
	}
	if err := scanner.Err(); err != nil {
		log.Printf("invalid version %s", err)
		return "invalid"
	}
	return "no-value"
}

const (
	repo            = "https://raw.githubusercontent.com/OpenSlides/openslides-backend/"
	modelsPath      = "/global/meta/models.yml"
	exampleDataPath = "/global/data/example-data.json"
	permissionPath  = "/global/meta/permission.yml"
)

// URLModelsYML return the url to the models.yml.
func URLModelsYML() string {
	return repo + getVersion() + modelsPath
}

// URLExampleData returns the url to the example-data file.
func URLExampleData() string {
	return repo + getVersion() + exampleDataPath
}

// URLPermission returns the url to the example-data file.
func URLPermission() string {
	return repo + getVersion() + permissionPath
}

// Unmarshal parses the content of models.yml to a datastruct.q
func Unmarshal(r io.Reader) (map[string]Model, error) {
	var m map[string]Model
	if err := yaml.NewDecoder(r).Decode(&m); err != nil {
		return nil, fmt.Errorf("decoding models: %w", err)
	}
	return m, nil
}

// Model replresents one model from models.yml.
type Model struct {
	Fields map[string]*Field
}

// UnmarshalYAML decodes a yaml model to models.Model.
func (m *Model) UnmarshalYAML(node *yaml.Node) error {
	return node.Decode(&m.Fields)
}

// Field of a model.
type Field struct {
	Type            string
	restrictionMode string
	relation        Relation
	Required        bool
}

// Relation returns the relation object if the Field is a relation. In other
// cases, it returns nil.
func (f *Field) Relation() Relation {
	if f.relation != nil {
		return f.relation
	}

	return nil
}

// RestrictionMode returns the restriction mode the field belongs to.
func (f *Field) RestrictionMode() string {
	return f.restrictionMode
}

// UnmarshalYAML decodes a model attribute from yaml.
func (f *Field) UnmarshalYAML(value *yaml.Node) error {
	var typer struct {
		Type            string `yaml:"type"`
		RestrictionMode string `yaml:"restriction_mode"`
		Required        bool   `yaml:"required"`
	}
	if err := value.Decode(&typer); err != nil {
		return fmt.Errorf("field object without type: %w", err)
	}

	f.Type = typer.Type
	f.restrictionMode = typer.RestrictionMode
	f.Required = typer.Required

	var list bool
	switch typer.Type {
	case "relation-list":
		list = true
		fallthrough

	case "relation":
		var relation AttributeRelation
		if err := value.Decode(&relation); err != nil {
			return fmt.Errorf("invalid object of type %s at line %d object: %w", typer.Type, value.Line, err)
		}
		relation.list = list
		f.relation = &relation

	case "generic-relation-list":
		list = true
		fallthrough

	case "generic-relation":
		var relation AttributeGenericRelation
		if err := value.Decode(&relation); err != nil {
			return fmt.Errorf("invalid object of type %s at line %d object: %w", typer.Type, value.Line, err)
		}
		relation.list = list
		f.relation = &relation

	}
	return nil
}

// Relation represents some kind of relation between fields.
type Relation interface {
	ToCollections() []ToCollectionField
	List() bool
}

// ToCollectionField represents a field and a collection
type ToCollectionField struct {
	Collection string  `yaml:"collection"`
	ToField    ToField `yaml:"field"`
}

// UnmarshalYAML decodes the models.yml to a To object.
func (t *ToCollectionField) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err == nil {
		cf := strings.Split(s, "/")
		if len(cf) != 2 {
			return fmt.Errorf("invalid value of `to` in line %d, expected one `/`: %s", value.Line, s)
		}
		t.Collection = cf[0]
		t.ToField.Name = cf[1]
		return nil
	}

	var d struct {
		Collection string  `yaml:"collection"`
		Field      ToField `yaml:"field"`
	}
	if err := value.Decode(&d); err != nil {
		return fmt.Errorf("decoding to collection field at line %d: %w", value.Line, err)
	}
	t.Collection = d.Collection
	t.ToField = d.Field
	return nil
}

// ToField is
type ToField struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

// UnmarshalYAML decodes the models.yml to a ToField object.
func (t *ToField) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err == nil {
		t.Name = s
		t.Type = "normal"
		return nil
	}

	var d struct {
		Name string `yaml:"name"`
		Type string `yaml:"type"`
	}
	if err := value.Decode(&d); err != nil {
		return fmt.Errorf("decoding to field at line %d: %w", value.Line, err)
	}
	t.Name = d.Name
	t.Type = d.Type
	return nil
}

// AttributeRelation is a relation or relation-list field.
type AttributeRelation struct {
	To   To `yaml:"to"`
	list bool
}

// ToCollections returns the names of the collections there the attribute points
// to. It is allways a slice with one element.
func (r AttributeRelation) ToCollections() []ToCollectionField {
	return []ToCollectionField{r.To.CollectionField}
}

// List returns true, if object is an attribute-relation-list
func (r AttributeRelation) List() bool {
	return r.list
}

// To is shows a Relation where to point to.
type To struct {
	CollectionField ToCollectionField
}

// UnmarshalYAML decodes the models.yml to a To object.
func (t *To) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err == nil {
		cf := strings.Split(s, "/")
		if len(cf) != 2 {
			return fmt.Errorf("invalid value of `to` in line %d, expected one `/`: %s", value.Line, s)
		}
		t.CollectionField.Collection = cf[0]
		t.CollectionField.ToField.Name = cf[1]
		return nil
	}

	if err := value.Decode(&(t.CollectionField)); err != nil {
		return fmt.Errorf("decoding to field at line %d: %w", value.Line, err)
	}
	return nil
}

// AttributeGenericRelation is a generic-relation or generic-relation-list field.
type AttributeGenericRelation struct {
	To   ToGeneric `yaml:"to"`
	list bool
}

// List tells, if the object is a generic-relation-list.
func (r AttributeGenericRelation) List() bool {
	return r.list
}

// ToCollections returns all collection, where the generic field could point to.
func (r AttributeGenericRelation) ToCollections() []ToCollectionField {
	return r.To.CollectionFields
}

// ToGeneric is like a To object, but for generic relations.
type ToGeneric struct {
	CollectionFields []ToCollectionField
}

// UnmarshalYAML unmarshalls data to a ToGeneric object.
func (t *ToGeneric) UnmarshalYAML(value *yaml.Node) error {
	var d struct {
		Collections []string `yaml:"collections"`
		Field       ToField  `yaml:"field"`
	}
	if err := value.Decode(&d); err == nil {
		t.CollectionFields = make([]ToCollectionField, len(d.Collections))
		for i, collection := range d.Collections {
			t.CollectionFields[i].Collection = collection
			t.CollectionFields[i].ToField = d.Field
		}
		return nil
	}

	var e []string
	if err := value.Decode(&e); err != nil {
		return fmt.Errorf("decoding to generic field at line %d: %w", value.Line, err)
	}
	t.CollectionFields = make([]ToCollectionField, len(e))
	for i, collectionfield := range e {
		cf := strings.Split(collectionfield, "/")
		if len(cf) != 2 {
			return fmt.Errorf("invalid value of `to` in line %d, expected one `/`: %s", value.Line, collectionfield)
		}
		t.CollectionFields[i].Collection = cf[0]
		t.CollectionFields[i].ToField.Name = cf[1]
	}
	return nil
}
