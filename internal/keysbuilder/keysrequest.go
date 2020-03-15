package keysbuilder

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

const (
	relationIdentifier            = "relation"
	relationListIdentifier        = "relation-list"
	genericRelationIdentifier     = "generic-relation"
	genericRelationListIdentifier = "generic-relation-list"
	templateIdentifier            = "template"
)

// body holds the information what keys are requested by the client.
type body struct {
	ids        []int
	collection string
	fieldsMap
}

// UnmarshallJSON builds a body object from json. It looks for the type
// argument in the fields and decodes the fields accorently.
func (b *body) UnmarshalJSON(data []byte) error {
	var field struct {
		IDs        []int     `json:"ids"`
		Collection string    `json:"collection"`
		Fields     fieldsMap `json:"fields"`
	}
	if err := json.Unmarshal(data, &field); err != nil {
		return err
	}
	if len(field.IDs) == 0 {
		return ErrInvalid{msg: "no ids"}
	}
	if field.Collection == "" {
		return ErrInvalid{msg: "no collection"}
	}
	if len(field.Fields.fields) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	b.ids = field.IDs
	b.collection = field.Collection
	b.fieldsMap = field.Fields
	return nil
}

func (b body) build(ctx context.Context, builder *Builder, keys chan<- string, errs chan<- error) {
	var wg sync.WaitGroup
	for _, id := range b.ids {
		for name, description := range b.fields {
			key := buildKey(b.collection, id, name)
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
		return ErrInvalid{msg: "no collection"}
	}
	if len(field.Fields.fields) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	r.collection = field.Collection
	r.fieldsMap = field.Fields
	return nil
}

func (r relationField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v := builder.cache.getOrSet(key, func() interface{} {
		id, err := builder.ider.ID(ctx, key)
		if err != nil {
			errs <- err
			return nil
		}
		return id
	})
	id, ok := v.(int)
	if !ok {
		errs <- fmt.Errorf("invalid value type in keybuilder cache: %v", v)
		return
	}
	if id == 0 {
		return
	}
	var wg sync.WaitGroup
	for name, description := range r.fields {
		key := buildKey(r.collection, id, name)
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

// relationListField is a fieldtype like relation, but redirects to a list of objects.
type relationListField struct {
	relationField
}

func (r relationListField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v := builder.cache.getOrSet(key, func() interface{} {
		ids, err := builder.ider.IDList(ctx, key)
		if err != nil {
			errs <- err
			return nil
		}
		return ids
	})
	ids, ok := v.([]int)
	if !ok {
		errs <- fmt.Errorf("invalid value type in keybuilder cache: %v", v)
		return
	}
	var wg sync.WaitGroup
	for _, id := range ids {
		for name, description := range r.fields {
			key := buildKey(r.collection, id, name)
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
	if len(field.Fields.fields) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	g.fieldsMap = field.Fields
	return nil
}

func (g genericRelationField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v := builder.cache.getOrSet(key, func() interface{} {
		gid, err := builder.ider.GenericID(ctx, key)
		if err != nil {
			errs <- err
			return nil
		}
		return gid
	})
	gid, ok := v.(string)
	if !ok {
		errs <- fmt.Errorf("invalid value type in keybuilder cache: %v", v)
		return
	}

	if gid == "" {
		return
	}
	var wg sync.WaitGroup
	for name, description := range g.fields {
		key := buildGenericKey(gid, name)
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

// genericRelationListField is like a genericRelationField but with a list of relations.
type genericRelationListField struct {
	genericRelationField
}

func (g genericRelationListField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v := builder.cache.getOrSet(key, func() interface{} {
		gids, err := builder.ider.GenericIDs(ctx, key)
		if err != nil {
			errs <- err
			return nil
		}
		return gids
	})
	gids, ok := v.([]string)
	if !ok {
		errs <- fmt.Errorf("invalid value type in keybuilder cache: %v", v)
		return
	}

	var wg sync.WaitGroup
	for _, gid := range gids {
		for name, description := range g.fields {
			key := buildGenericKey(gid, name)
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
	}
	wg.Wait()
}

// templateField requests a list of fields from a template.
type templateField struct {
	sub fieldDescription
}

func (t *templateField) UnmarshalJSON(data []byte) error {
	var jsonTemplate struct {
		Sub json.RawMessage `json:"sub"`
	}
	if err := json.Unmarshal(data, &jsonTemplate); err != nil {
		return err
	}
	sub, err := unmarshalField(jsonTemplate.Sub)
	if err != nil {
		if sub, ok := err.(ErrInvalid); ok {
			return ErrInvalid{sub: &sub, msg: "Error in template sub", field: "template"}
		}
		return err
	}
	t.sub = sub
	return nil
}

func (t templateField) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
	v := builder.cache.getOrSet(key, func() interface{} {
		values, err := builder.ider.Template(ctx, key)
		if err != nil {
			errs <- err
			return nil
		}
		return values
	})
	values, ok := v.([]string)
	if !ok {
		errs <- fmt.Errorf("invalid value type in keybuilder cache: %v", v)
		return
	}

	var wg sync.WaitGroup
	for _, value := range values {
		newKey := strings.Replace(key, "$", value, 1)
		keys <- newKey
		if t.sub == nil {
			continue
		}
		wg.Add(1)
		go func() {
			t.sub.build(ctx, builder, newKey, keys, errs)
			wg.Done()
		}()
	}
	wg.Wait()
}

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
	var err error
	switch t.Type {
	case relationIdentifier:
		var r relationField
		if err = json.Unmarshal(data, &r); err != nil {
			return nil, err
		}
		return r, nil
	case relationListIdentifier:
		var r relationListField
		if err = json.Unmarshal(data, &r); err != nil {
			return nil, err
		}
		return r, nil
	case genericRelationIdentifier:
		var r genericRelationField
		if err = json.Unmarshal(data, &r); err != nil {
			return nil, err
		}
		return r, nil
	case genericRelationListIdentifier:
		var r genericRelationListField
		if err = json.Unmarshal(data, &r); err != nil {
			return nil, err
		}
		return r, nil
	case templateIdentifier:
		var r templateField
		if err = json.Unmarshal(data, &r); err != nil {
			return nil, err
		}
		return r, nil
	case "":
		return nil, ErrInvalid{msg: "no type"}
	default:
		return nil, ErrInvalid{msg: fmt.Sprintf("unknown type %s", t.Type)}
	}
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
			if sub, ok := err.(ErrInvalid); ok {
				return ErrInvalid{sub: &sub, msg: "Error on field", field: name}
			}
			return err
		}
		f.fields[name] = fd
	}
	return nil
}
