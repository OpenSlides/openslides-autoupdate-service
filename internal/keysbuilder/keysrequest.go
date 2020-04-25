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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
)

const (
	relationIdentifier            = "relation"
	relationListIdentifier        = "relation-list"
	genericRelationIdentifier     = "generic-relation"
	genericRelationListIdentifier = "generic-relation-list"
	templateIdentifier            = "template"
)

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
	if field.Collection == "" {
		return InvalidError{msg: "no collection"}
	}
	if field.Fields.fields == nil {
		return InvalidError{msg: "no fields"}
	}

	// Set the body fields.
	b.ids = field.IDs
	b.collection = field.Collection
	b.fieldsMap = field.Fields
	return nil
}

func (b body) build(ctx context.Context, builder *Builder, keys chan<- string, errs chan<- error) {
	var wg sync.WaitGroup
	for _, id := range b.ids {
		wg.Add(1)
		go func(id int) {
			b.fieldsMap.build(ctx, buildCollectionID(b.collection, id), builder, keys, errs)
			wg.Done()
		}(id)
	}
	wg.Wait()
}

// relationField is a fieldtype that redirects to one other collection
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
	r.collection = field.Collection
	r.fieldsMap = field.Fields
	return nil
}

func (r relationField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v, err := builder.cache.getOrSet(key, func() (interface{}, error) {
		return builder.ider.ID(ctx, key)
	})
	if err != nil {
		if !errors.Is(err, autoupdate.ErrUnknownKey) {
			errs <- fmt.Errorf("can not use value of key %s: %w", key, err)
		}
		return
	}
	id, ok := v.(int)
	if !ok {
		errs <- fmt.Errorf("invalid value type %T in keysbuilder cache, expected int, got: %v", v, v)
		return
	}
	r.fieldsMap.build(ctx, buildCollectionID(r.collection, id), builder, keys, errs)
}

// relationListField is a fieldtype like relation, but redirects to a list of objects.
type relationListField struct {
	relationField
}

func (r relationListField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v, err := builder.cache.getOrSet(key, func() (interface{}, error) {
		return builder.ider.IDList(ctx, key)
	})
	if err != nil {
		if !errors.Is(err, autoupdate.ErrUnknownKey) {
			errs <- fmt.Errorf("can not use value of key %s: %w", key, err)
		}
		return
	}
	ids, ok := v.([]int)
	if !ok {
		errs <- fmt.Errorf("invalid value type %T in keysbuilder cache, expected []int, got: %v", v, v)
		return
	}
	var wg sync.WaitGroup
	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			r.fieldsMap.build(ctx, buildCollectionID(r.collection, id), builder, keys, errs)
			wg.Done()
		}(id)
	}
	wg.Wait()
}

// genericRelationField is like a relationField but the collection is given from the restricter.
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

func (g genericRelationField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v, err := builder.cache.getOrSet(key, func() (interface{}, error) {
		return builder.ider.GenericID(ctx, key)
	})
	if err != nil {
		if !errors.Is(err, autoupdate.ErrUnknownKey) {
			errs <- fmt.Errorf("can not use value of key %s: %w", key, err)
		}
		return
	}
	gid, ok := v.(string)
	if !ok {
		errs <- fmt.Errorf("invalid value type %T in keysbuilder cache, expected []string, got: %v", v, v)
		return
	}
	g.fieldsMap.build(ctx, gid, builder, keys, errs)
}

// genericRelationListField is like a genericRelationField but with a list of relations.
type genericRelationListField struct {
	genericRelationField
}

func (g genericRelationListField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v, err := builder.cache.getOrSet(key, func() (interface{}, error) {
		return builder.ider.GenericIDs(ctx, key)
	})
	if err != nil {
		if !errors.Is(err, autoupdate.ErrUnknownKey) {
			errs <- fmt.Errorf("can not use value of key %s: %w", key, err)
		}
		return
	}
	gids, ok := v.([]string)
	if !ok {
		errs <- fmt.Errorf("invalid value type %T in keysbuilder cache, expected []string, got: %v", v, v)
		return
	}

	var wg sync.WaitGroup
	for _, gid := range gids {
		wg.Add(1)
		go func(gid string) {
			g.fieldsMap.build(ctx, gid, builder, keys, errs)
			wg.Done()
		}(gid)
	}
	wg.Wait()
}

// templateField requests a list of fields from a template.
type templateField struct {
	values fieldDescription
}

func (t *templateField) UnmarshalJSON(data []byte) error {
	var field struct {
		Values json.RawMessage `json:"values"`
	}
	if err := json.Unmarshal(data, &field); err != nil {
		return fmt.Errorf("can not decode template field: %w", err)
	}
	if len(field.Values) == 0 {
		return nil
	}

	values, err := unmarshalField(field.Values)
	if err != nil {
		if sub, ok := err.(InvalidError); ok {
			return InvalidError{sub: &sub, msg: "Error in template sub", field: "template"}
		}
		return fmt.Errorf("can not decode sub attribute of template field: %w", err)
	}
	t.values = values
	return nil
}

func (t templateField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v, err := builder.cache.getOrSet(key, func() (interface{}, error) {
		return builder.ider.Template(ctx, key)
	})
	if err != nil {
		if !errors.Is(err, autoupdate.ErrUnknownKey) {
			errs <- fmt.Errorf("can not use value of key %s: %w", key, err)
		}
		return
	}

	values, ok := v.([]string)
	if !ok {
		errs <- fmt.Errorf("invalid value type %T in keysbuilder cache, expected []string, got: %v", v, v)
		return
	}

	var wg sync.WaitGroup
	for _, value := range values {
		newKey := strings.Replace(key, "$", value, 1)
		keys <- newKey
		if t.values == nil {
			continue
		}
		wg.Add(1)
		go func() {
			t.values.build(ctx, builder, newKey, keys, errs)
			wg.Done()
		}()
	}
	wg.Wait()
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
	case relationIdentifier:
		r = new(relationField)

	case relationListIdentifier:
		r = new(relationListField)

	case genericRelationIdentifier:
		r = new(genericRelationField)

	case genericRelationListIdentifier:
		r = new(genericRelationListField)

	case templateIdentifier:
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

// fieldsMap describes in a abstract way the fieldsMap of a collection.
type fieldsMap struct {
	fields map[string]fieldDescription
}

func (f *fieldsMap) UnmarshalJSON(data []byte) error {
	var fm map[string]json.RawMessage
	if err := json.Unmarshal(data, &fm); err != nil {
		return fmt.Errorf("can not decode fields: %w", err)
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

// build calls the build method for all fields in the fieldsMap.
func (f *fieldsMap) build(ctx context.Context, collectionID string, builder *Builder, keys chan<- string, errs chan<- error) {
	var wg sync.WaitGroup
	for name, description := range f.fields {
		key := buildGenericKey(collectionID, name)
		keys <- key
		if description == nil {
			continue
		}
		wg.Add(1)
		go func(description fieldDescription) {
			description.build(ctx, builder, key, keys, errs)
			wg.Done()
		}(description)
	}
	wg.Wait()
}
