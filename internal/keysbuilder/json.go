package keysbuilder

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

// FromJSON creates a Keysbuilder from json.
func FromJSON(r io.Reader) (*Builder, error) {
	var b body
	if err := json.NewDecoder(r).Decode(&b); err != nil {
		if err == io.EOF {
			return nil, InvalidError{msg: "No data"}
		}
		if sub, ok := err.(InvalidError); ok {
			return nil, sub
		}
		return nil, JSONError{err}
	}

	kb := &Builder{
		bodies: []body{b},
	}
	return kb, nil
}

// ManyFromJSON creates a list of Keysbuilder objects from a json list.
func ManyFromJSON(r io.Reader) (*Builder, error) {
	var bs []body
	if err := json.NewDecoder(r).Decode(&bs); err != nil {
		if err == io.EOF {
			return nil, InvalidError{msg: "No data"}
		}
		if sub, ok := err.(InvalidError); ok {
			return nil, sub
		}
		if jerr, ok := err.(*json.SyntaxError); ok {
			return nil, JSONError{jerr}
		}
		if jerr, ok := err.(*json.UnmarshalTypeError); ok {
			var expectType string
			switch jerr.Type.Kind() {
			case reflect.Struct:
				expectType = "object"
			case reflect.Slice:
				expectType = "list"
			case reflect.Int:
				expectType = "number"
			default:
				expectType = jerr.Type.Kind().String()
			}

			return nil, InvalidError{msg: fmt.Sprintf("wrong type at field `%s`. Got %s, expected %v", jerr.Field, jerr.Value, expectType)}
		}
		return nil, fmt.Errorf("decode keysrequest: %w", err)
	}

	if len(bs) == 0 {
		return nil, InvalidError{msg: "No data"}
	}

	kb := &Builder{
		bodies: bs,
	}
	return kb, nil
}
