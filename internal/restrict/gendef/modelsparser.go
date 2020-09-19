package main

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type mField struct {
	Name string
	Type string
}

func (t *mField) UnmarshalYAML(value *yaml.Node) error {
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

// type mFieldStructuredRelation struct {
// 	mField
// 	Replacement string `yaml:"replacement"`
// }

// type mFieldStructuredTag struct {
// 	mField
// 	Replacement string `yaml:"replacement"`
// }

type mTo struct {
	Collection string
	Field      mField
}

func (t mTo) modelField() string {
	return fmt.Sprintf("%s/%s", t.Collection, t.Field.Name)
}

func (t *mTo) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err == nil {
		cf := strings.Split(s, "/")
		if len(cf) != 2 {
			return fmt.Errorf("invalid value of `to` in line %d, expected one `/`: %s", value.Line, s)
		}
		t.Collection = cf[0]
		t.Field.Name = cf[1]
		return nil
	}

	var d struct {
		Collection string `yaml:"collection"`
		Field      mField `yaml:"field"`
	}
	if err := value.Decode(&d); err != nil {
		return fmt.Errorf("decoding to field at line %d: %w", value.Line, err)
	}

	t.Collection = d.Collection
	t.Field = d.Field
	return nil
}

type mToGeneric struct {
	Collection []string
	Field      mField
}

func (t mToGeneric) modelField() string {
	return fmt.Sprintf("*/%s", t.Field.Name)
}

type toCollectioner interface {
	toCollection() string
}

type mValue struct {
	Type     string
	relation toCollectioner
	template *mValueTemplate
}

type mValueRelation struct {
	To mTo `yaml:"to"`
}

func (r mValueRelation) toCollection() string {
	return r.To.Collection
}

type mValueGenericRelation struct {
	To mToGeneric `yaml:"to"`
}

func (r mValueGenericRelation) toCollection() string {
	return "*"
}

type mValueTemplate struct {
	Replacement string `yaml:"replacement"`
	Fields      mValue `yaml:"fields"`
}

func (v *mValue) UnmarshalYAML(value *yaml.Node) error {
	var s string
	if err := value.Decode(&s); err == nil {
		v.Type = s
		return nil
	}

	var typer struct {
		Type string `yaml:"type"`
	}
	if err := value.Decode(&typer); err != nil {
		return fmt.Errorf("field object without type: %w", err)
	}

	v.Type = typer.Type
	switch typer.Type {
	case "relation":
		fallthrough
	case "relation-list":
		var relation mValueRelation
		if err := value.Decode(&relation); err != nil {
			return fmt.Errorf("invalid object of type %s at line %d object: %w", typer.Type, value.Line, err)
		}
		v.relation = &relation
	case "generic-relation":
		fallthrough
	case "generic-relation-list":
		var relation mValueGenericRelation
		if err := value.Decode(&relation); err != nil {
			return fmt.Errorf("invalid object of type %s at line %d object: %w", typer.Type, value.Line, err)
		}
		v.relation = &relation
	case "template":
		var template mValueTemplate
		if err := value.Decode(&template); err != nil {
			return fmt.Errorf("invalid object of type template object in line %d: %w", value.Line, err)
		}
		v.template = &template
	default:
		return fmt.Errorf("unknown type `%s` in line %d", typer.Type, value.Line)
	}
	return nil
}
