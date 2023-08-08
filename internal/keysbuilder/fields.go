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
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/fastjson"
)

const (
	ftRelation            = "relation"
	ftRelationList        = "relation-list"
	ftGenericRelation     = "generic-relation"
	ftGenericRelationList = "generic-relation-list"
)

// keyDescription combines a key and a fieldDescription.
//
// This is used, when a list of key-description combinations are needed.
type keyDescription struct {
	key         dskey.Key
	description fieldDescription
}

// fieldDescription is an interface that appends keys.
//
// The different field-types (relation, relation-list, etc.) implement this
// interface and return all keys, they represent.
type fieldDescription interface {
	appendKeys(key dskey.Key, value json.RawMessage, data []keyDescription) ([]keyDescription, error)
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

	for fieldName := range field.Fields.fields {
		if !dskey.ValidateCollectionField(field.Collection, fieldName) {
			return InvalidError{
				msg:   fmt.Sprintf("%s/%s does not exist", field.Collection, fieldName),
				field: fieldName,
			}
		}
	}

	// Set the body fields.
	b.ids = field.IDs
	b.collection = field.Collection
	b.fieldsMap = field.Fields
	return nil
}

// appendKeys appends all body-keys with there descriptions.
//
// It is simular to the fieldDescription interface. But it requires other
// arguments.
func (b *body) appendKeys(data []keyDescription) ([]keyDescription, error) {
	var err error
	for _, id := range b.ids {
		data, err = b.fieldsMap.appendKeys(b.collection, id, data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

// relationField is a fieldtype that redirects to one other collection.
//
//	{
//		"ids": [1],
//		"collection": "user",
//		"fields": {
//			"note_id": {
//				"type": "relation",
//				"collection": "note",
//				"fields": {"important": null}
//			}
//		}
//	}
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

	for fieldName := range field.Fields.fields {
		if !dskey.ValidateCollectionField(field.Collection, fieldName) {
			return InvalidError{
				msg:   fmt.Sprintf("%s/%s does not exist", field.Collection, fieldName),
				field: fieldName,
			}
		}
	}

	r.collection = field.Collection
	r.fieldsMap = field.Fields
	return nil
}

func (r *relationField) appendKeys(key dskey.Key, value json.RawMessage, data []keyDescription) ([]keyDescription, error) {
	id, err := fastjson.DecodeInt(value)
	if err != nil {
		return nil, fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	if id <= 0 {
		// TODO: This should be return an error on required fields
		return data, nil
	}

	return r.fieldsMap.appendKeys(r.collection, id, data)
}

// relationListField is a fieldtype like relation, but redirects to a list of objects.
//
//	{
//		"ids": [1],
//		"collection": "user",
//		"fields": {
//			"group_ids": {
//				"type": "relation-list",
//				"collection": "group",
//				"fields": {"name": null}
//			}
//		}
//	}
type relationListField struct {
	relationField
}

func (r *relationListField) appendKeys(key dskey.Key, value json.RawMessage, data []keyDescription) ([]keyDescription, error) {
	ids, err := fastjson.DecodeIntList(value)
	if err != nil {
		return nil, fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	for _, id := range ids {
		for field, description := range r.fields {
			key, err := dskey.FromParts(r.collection, id, field)
			if err != nil {
				return nil, fmt.Errorf("invalid key: %w", err)
			}
			data = append(data, keyDescription{key: key, description: description})
		}
	}
	return data, nil
}

// genericRelationField is like a relationField but the collection is given from the restricter.
//
//	{
//		"ids": [1],
//		"collection": "user",
//		"fields": {
//			"most_seen": {
//				"type": "generic-relation",
//				"fields": {"name": null}
//			}
//		}
//	}
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

func (g *genericRelationField) appendKeys(key dskey.Key, value json.RawMessage, data []keyDescription) ([]keyDescription, error) {
	var fqID string
	if err := json.Unmarshal(value, &fqID); err != nil {
		return nil, fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	collection, rawID, found := strings.Cut(fqID, "/")
	if !found {
		return nil, fmt.Errorf("invalid collection id: %s", fqID)
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		return nil, fmt.Errorf("invalid collection id: %s", fqID)
	}

	for fieldName := range g.fieldsMap.fields {
		if !dskey.ValidateCollectionField(collection, fieldName) {
			delete(g.fieldsMap.fields, fieldName)
		}
	}

	return g.fieldsMap.appendKeys(collection, id, data)
}

// genericRelationListField is like a genericRelationField but with a list of relations.
//
//	{
//		"ids": [1],
//		"collection": "user",
//		"fields": {
//			"seen": {
//				"type": "generic-relation-list",
//				"fields": {"name": null}
//			}
//		}
//	}
type genericRelationListField struct {
	genericRelationField
}

func (g *genericRelationListField) appendKeys(key dskey.Key, value json.RawMessage, data []keyDescription) ([]keyDescription, error) {
	var fqIDs []string
	if err := json.Unmarshal(value, &fqIDs); err != nil {
		return nil, fmt.Errorf("decoding value for key %s: %w", key, err)
	}

	for _, fqID := range fqIDs {
		collection, rawID, found := strings.Cut(fqID, "/")
		if !found {
			return nil, fmt.Errorf("invalid collection id: %s", fqID)
		}

		id, err := strconv.Atoi(rawID)
		if err != nil {
			return nil, fmt.Errorf("invalid collection id: %s", fqID)
		}

		for fieldName := range g.fieldsMap.fields {
			if !dskey.ValidateCollectionField(collection, fieldName) {
				delete(g.fieldsMap.fields, fieldName)
			}
		}

		data, err = g.fieldsMap.appendKeys(collection, id, data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
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

// appendKeys appends its fields to data.
//
// It is like the fieldDescription interface. But it requires other arguments.
func (f *fieldsMap) appendKeys(collection string, id int, data []keyDescription) ([]keyDescription, error) {
	for field, description := range f.fields {
		key, err := dskey.FromParts(collection, id, field)
		if err != nil {
			return nil, err
		}
		data = append(data, keyDescription{key: key, description: description})
	}
	return data, nil
}
