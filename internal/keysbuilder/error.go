package keysbuilder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// InvalidError is an error that happens on an invalid request.
type InvalidError struct {
	msg   string
	field string
	sub   *InvalidError
}

func (e InvalidError) Error() string {
	fields, last := e.fields()
	if fields == nil {
		return e.msg
	}
	return fmt.Sprintf("field %q: %s", strings.Join(fields, "."), last.msg)
}

// Type returns the name of the error.
func (e InvalidError) Type() string {
	return "SyntaxError"
}

// Fields returns a list of field names from the parent to this error.
func (e InvalidError) Fields() []string {
	fields, _ := e.fields()
	return fields
}

func (e InvalidError) fields() ([]string, *InvalidError) {
	if e.field == "" {
		return nil, nil
	}
	fields := []string{e.field}
	last := &e
	for last.sub != nil {
		last = last.sub
		if last.field != "" {
			fields = append(fields, last.field)
		}
	}
	return fields, last
}

// JSONError is returned when invalid json is parsed or the json can not be
// decoded as a keysbuilder.
type JSONError struct {
	err error
}

func (e JSONError) Error() string {
	return e.err.Error()
}

// Unwrap returns the thrown error.
func (e JSONError) Unwrap() error {
	return e.err
}

// Type returns the name of the error.
func (e JSONError) Type() string {
	return "JsonError"
}

// ValueError in returned by keysbuilder.Update(), when the value of a key has
// not the expected format.
type ValueError struct {
	key        datastore.Key
	gotType    string
	expectType reflect.Type
	err        error
}

func (e ValueError) Error() string {
	return fmt.Sprintf("invalid value in key %s. Got %s, expected %s", e.key, e.gotType, e.expectType)
}

// Type returns the name of the error.
func (e ValueError) Type() string {
	return "ValueError"
}

// Unwrap returns the thrown error.
func (e ValueError) Unwrap() error {
	return e.err
}
