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
	var jsonBody struct {
		IDs        []int     `json:"ids"`
		Collection string    `json:"collection"`
		Fields     fieldsMap `json:"fields"`
	}
	if err := json.Unmarshal(data, &jsonBody); err != nil {
		return err
	}
	b.ids = jsonBody.IDs
	b.collection = jsonBody.Collection
	b.fieldsMap = jsonBody.Fields
	return nil
}

// validate makes sure the body is valid. Returns an ErrInvalid if not.
func (b body) validate() error {
	if len(b.ids) == 0 {
		return ErrInvalid{msg: "no ids"}
	}
	if b.collection == "" {
		return ErrInvalid{msg: "no collection"}
	}
	return b.fieldsMap.validate()
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
	var jsonRelation struct {
		Collection string    `json:"collection"`
		Fields     fieldsMap `json:"fields"`
	}
	if err := json.Unmarshal(data, &jsonRelation); err != nil {
		return fmt.Errorf("can not decode id and collection: %w", err)
	}
	r.collection = jsonRelation.Collection
	r.fieldsMap = jsonRelation.Fields
	return nil
}

func (r relationField) validate() error {
	if len(r.fields) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	if r.collection == "" {
		return ErrInvalid{msg: "no collection"}
	}
	return r.fieldsMap.validate()
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
		return fmt.Errorf("can not decode id and collection: %w", err)
	}
	g.fieldsMap = field.Fields
	return nil
}

func (g genericRelationField) validate() error {
	return g.fieldsMap.validate()
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
		return fmt.Errorf("can not decode template field: %w", err)
	}
	fd, err := unmarshalFieldDescription(jsonTemplate.Sub)
	if err != nil {
		return ErrInvalid{msg: err.Error(), field: "template"}
	}
	t.sub = fd
	return nil
}

func (t templateField) validate() error {
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

func unmarshalFieldDescription(data []byte) (fieldDescription, error) {
	// TODO: handle json errors
	var t *struct {
		Type string `json:"type"`
	}
	json.Unmarshal(data, &t)
	if t == nil {
		return nil, nil
	}
	switch t.Type {
	case relationIdentifier:
		var r relationField
		json.Unmarshal(data, &r)
		return r, nil
	case relationListIdentifier:
		var r relationListField
		json.Unmarshal(data, &r)
		return r, nil
	case genericRelationIdentifier:
		var r genericRelationField
		json.Unmarshal(data, &r)
		return r, nil
	case genericRelationListIdentifier:
		var r genericRelationListField
		json.Unmarshal(data, &r)
		return r, nil
	case templateIdentifier:
		var r templateField
		json.Unmarshal(data, &r)
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
	f.fields = make(map[string]fieldDescription)
	for name, description := range fm {
		fd, err := unmarshalFieldDescription(description)
		if err != nil {
			return ErrInvalid{msg: err.Error(), field: name}
		}
		f.fields[name] = fd
	}
	return nil
}

// validate maks sure the fields are valid. Returns an ErrInvalid if not.
func (f fieldsMap) validate() error {
	if len(f.fields) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	for name, description := range f.fields {
		if description == nil {
			continue
		}
		if err := description.validate(); err != nil {
			sub := err.(ErrInvalid)
			return ErrInvalid{sub: &sub, field: name, msg: "Error on field"}
		}
	}
	return nil
}
