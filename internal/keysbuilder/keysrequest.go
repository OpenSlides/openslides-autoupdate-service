package keysbuilder

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
)

const (
	relationIdentifier     = "relation"
	relationListIdentifier = "relation-list"
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
		return fmt.Errorf("can not decode id and collection: %w", err)
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

// relation is a fieldtype that redirects to one other collection
type relation struct {
	collection string
	fieldsMap
}

func (r *relation) UnmarshalJSON(data []byte) error {
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

func (r relation) validate() error {
	if len(r.fields) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	if r.collection == "" {
		return ErrInvalid{msg: "no collection"}
	}
	return r.fieldsMap.validate()
}

func (r relation) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
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

// relationList is a fieldtype like relation, but redirects to a list of objects.
type relationList struct {
	relation
}

func (r relationList) build(ctx context.Context, builder *Builder, key string, keys chan<- string, errs chan<- error) {
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

// fieldsMap describes in a abstract way the fieldsMap of a collection.
type fieldsMap struct {
	fields map[string]fieldDescription
}

func (f *fieldsMap) UnmarshalJSON(data []byte) error {
	s := string(data)
	_ = s
	var fm map[string]json.RawMessage
	if err := json.Unmarshal(data, &fm); err != nil {
		return fmt.Errorf("can not decode fields: %w", err)
	}
	var t *struct {
		Type string `json:"type"`
	}
	f.fields = make(map[string]fieldDescription)
	for name, description := range fm {
		t = nil
		json.Unmarshal(description, &t)
		if t == nil {
			f.fields[name] = nil
			continue
		}
		switch t.Type {
		case relationIdentifier:
			var r relation
			json.Unmarshal(description, &r)
			f.fields[name] = r
		case relationListIdentifier:
			var r relationList
			json.Unmarshal(description, &r)
			f.fields[name] = r
		case "":
			return ErrInvalid{msg: "no type", field: name}
		default:
			return ErrInvalid{msg: fmt.Sprintf("unknown type %s", t.Type), field: name}
		}
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
