/*
This file defines the different field types. The body (the "root" field) is also
defined.

Each fieldtype implements the fieldDescription interface by implementing the
build method.

Each field is processed by creating a field object from json. Therefore most
field types also implement the json.Unmarshaler interface.

Most fields have sub-keys. Together the different fields create a tree like
object starting from the body.

After the tree is parsed, the build-methods are used to receive the acutal keys.

Each build-call parses its branches concurrently. But it is important that the
build method only returns, when all sub-jobs are done.
*/

package keysbuilder

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

const (
	ftRelation            = "relation"
	ftRelationList        = "relation-list"
	ftGenericRelation     = "generic-relation"
	ftGenericRelationList = "generic-relation-list"
	ftTemplate            = "template"
)

var (
	reCollection = regexp.MustCompile(`^([a-z]+|[a-z][a-z_]*[a-z])$`)
	reField      = regexp.MustCompile(`^[a-z][a-z0-9_]*\$?[a-z0-9_]*$`)
)

type fieldDescription interface {
	keys(key datastore.Key, value json.RawMessage, data map[datastore.Key]fieldDescription) error
}

// body holds the information which keys are requested by the client.
type body struct {
	ids        []int
	collection string
	fieldsMap
}

// UnmarshallJSON builds a body object from json. It looks for the type argument
// in the fields and decodes the fields accorently.
func (b *body) UnmarshalJSON(data []byte) error {
	var field struct {
		IDs        []int     `json:"ids"`
		Collection string    `json:"collection"`
		Fields     fieldsMap `json:"fields"`
	}

	// Read and validate the data.
	if err := json.Unmarshal(data, &field); err != nil {
		return err
	}

	if len(field.IDs) == 0 {
		return InvalidError{msg: "no ids"}
	}
	for _, id := range field.IDs {
		if id <= 0 {
			return InvalidError{msg: "id has to be a positive number"}
		}
	}

	if field.Collection == "" {
		return InvalidError{msg: "attribute collection is missing"}
	}
	if field.Fields.fields == nil {
		return InvalidError{msg: "attribte fields is missing"}
	}
	if !reCollection.MatchString(field.Collection) {
		return InvalidError{msg: "invalid collection name"}
	}

	// Set the body fields.
	b.ids = field.IDs
	b.collection = field.Collection
	b.fieldsMap = field.Fields
	return nil
}

func (b *body) keys(data map[datastore.Key]fieldDescription) {
	for _, id := range b.ids {
		cid := buildCollectionID(b.collection, id)
		b.fieldsMap.keys(cid, data)
	}
}

// relationField is a fieldtype that redirects to one other collection.
//
// {
//	"ids": [1],
//	"collection": "user",
//	"fields": {
//		"note_id": {
//			"type": "relation",
//			"collection": "note",
//			"fields": {"important": null}
//		}
//	}
// }
type relationField struct {
	collection string
	fieldsMap
}

func (r *relationField) UnmarshalJSON(data []byte) error {
	var field struct {
		Collection string    `json:"collection"`
		Fields     fieldsMap `json:"fields"`
	}
	if err := json.Unmarshal(data, &field); err != nil {
		return err
	}
	if field.Collection == "" {
		return InvalidError{msg: "no collection"}
	}
	if field.Fields.fields == nil {
		return InvalidError{msg: "no fields"}
	}
	if !reCollection.MatchString(field.Collection) {
		return InvalidError{msg: "invalid collection name"}
	}
	r.collection = field.Collection
	r.fieldsMap = field.Fields
	return nil
}

func (r *relationField) keys(key datastore.Key, value json.RawMessage, data map[datastore.Key]fieldDescription) error {
	var id int
	if err := json.Unmarshal(value, &id); err != nil {
		return fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	cid := buildCollectionID(r.collection, id)
	r.fieldsMap.keys(cid, data)
	return nil
}

// relationListField is a fieldtype like relation, but redirects to a list of objects.
//
// {
//	"ids": [1],
//	"collection": "user",
//	"fields": {
//		"group_ids": {
//			"type": "relation-list",
//			"collection": "group",
//			"fields": {"name": null}
//		}
//	}
// }
type relationListField struct {
	relationField
}

func (r *relationListField) keys(key datastore.Key, value json.RawMessage, data map[datastore.Key]fieldDescription) error {
	var ids []int
	if err := json.Unmarshal(value, &ids); err != nil {
		return fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	for _, id := range ids {
		cid := buildCollectionID(r.collection, id)
		for field, description := range r.fields {
			data[buildGenericKey(cid, field)] = description
		}
	}
	return nil
}

// genericRelationField is like a relationField but the collection is given from the restricter.
//
//{
//	"ids": [1],
//	"collection": "user",
//	"fields": {
//		"most_seen": {
//			"type": "generic-relation",
//			"fields": {"name": null}
//		}
//	}
// }
type genericRelationField struct {
	fieldsMap
}

func (g *genericRelationField) UnmarshalJSON(data []byte) error {
	var field struct {
		Fields fieldsMap `json:"fields"`
	}
	if err := json.Unmarshal(data, &field); err != nil {
		return err
	}
	if field.Fields.fields == nil {
		return InvalidError{msg: "no fields"}
	}
	g.fieldsMap = field.Fields
	return nil
}

func (g *genericRelationField) keys(key datastore.Key, value json.RawMessage, data map[datastore.Key]fieldDescription) error {
	var cid string
	if err := json.Unmarshal(value, &cid); err != nil {
		return fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	g.fieldsMap.keys(cid, data)
	return nil
}

// genericRelationListField is like a genericRelationField but with a list of relations.
//
// {
//	"ids": [1],
//	"collection": "user",
//	"fields": {
//		"seen": {
//			"type": "generic-relation-list",
//			"fields": {"name": null}
//		}
//	}
// }
type genericRelationListField struct {
	genericRelationField
}

func (g *genericRelationListField) keys(key datastore.Key, value json.RawMessage, data map[datastore.Key]fieldDescription) error {
	var cids []string
	if err := json.Unmarshal(value, &cids); err != nil {
		return fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	for _, cid := range cids {
		g.fieldsMap.keys(cid, data)
	}
	return nil
}

// templateField requests a list of fields from a template.
//
// {
//	"ids": [1],
//	"collection": "user",
//	"fields": {
//		"group_$_ids": {
//			"type": "template",
//			"values": {
//				"type": "relation-list",
//				"collection": "group",
//				"fields": {"name": null}
//			}
//		}
//	}
// }
type templateField struct {
	values fieldDescription
}

func (t *templateField) UnmarshalJSON(data []byte) error {
	var field struct {
		Values json.RawMessage `json:"values"`
	}
	if err := json.Unmarshal(data, &field); err != nil {
		return fmt.Errorf("decode template field: %w", err)
	}
	if len(field.Values) == 0 {
		return nil
	}

	values, err := unmarshalField(field.Values)
	if err != nil {
		if sub, ok := err.(InvalidError); ok {
			return InvalidError{sub: &sub, msg: "Error in template sub", field: "template"}
		}
		return fmt.Errorf("decoding sub attribute of template field: %w", err)
	}
	t.values = values
	return nil
}

func (t *templateField) keys(key datastore.Key, value json.RawMessage, data map[datastore.Key]fieldDescription) error {
	var values []string
	if err := json.Unmarshal(value, &values); err != nil {
		return fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	for _, value := range values {
		newkey := key
		newkey.Field = strings.Replace(key.Field, "$", "$"+value, 1)
		data[newkey] = t.values
	}
	return nil
}

// unmarshalField uses the type-attribute in the json object get the field-type.
// Afterwards, the json is parsed as this field-type and returned.
func unmarshalField(data []byte) (fieldDescription, error) {
	var t *struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}
	if t == nil {
		return nil, nil
	}

	var r fieldDescription
	switch t.Type {
	case ftRelation:
		r = new(relationField)

	case ftRelationList:
		r = new(relationListField)

	case ftGenericRelation:
		r = new(genericRelationField)

	case ftGenericRelationList:
		r = new(genericRelationListField)

	case ftTemplate:
		r = new(templateField)

	case "":
		return nil, InvalidError{msg: "no type"}

	default:
		return nil, InvalidError{msg: fmt.Sprintf("unknown type %s", t.Type)}
	}

	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	return r, nil
}

// fieldsMap is a map from each field of another field descrption.
//
// For example, user_id to a relation-field and group_ids to a
// relation-list-field.
//
// A fieldsMap knows how to be decoded from json and how to build the keys from
// it.
type fieldsMap struct {
	fields map[string]fieldDescription
}

func (f *fieldsMap) UnmarshalJSON(data []byte) error {
	var fm map[string]json.RawMessage
	if err := json.Unmarshal(data, &fm); err != nil {
		return fmt.Errorf("decode fields: %w", err)
	}

	f.fields = make(map[string]fieldDescription, len(fm))
	for name, field := range fm {
		if !reField.MatchString(name) {
			return InvalidError{msg: fmt.Sprintf("fieldname %q is not a valid fieldname", name), field: name}
		}

		fd, err := unmarshalField(field)
		if err != nil {
			if sub, ok := err.(InvalidError); ok {
				return InvalidError{sub: &sub, msg: "Error on field", field: name}
			}
			return err
		}
		f.fields[name] = fd
	}
	return nil
}

func (f *fieldsMap) keys(cid string, data map[datastore.Key]fieldDescription) {
	for field, description := range f.fields {
		data[buildGenericKey(cid, field)] = description
	}
}
