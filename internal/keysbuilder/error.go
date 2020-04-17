package keysbuilder

import (
	"fmt"
	"strings"
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
	return fmt.Sprintf("field \"%s\": %s", strings.Join(fields, "."), last.msg)
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
	if e.sub == nil {
		return []string{e.field}, &e
	}
	var fields []string
	last := &e
	for last.sub != nil {
		fields = append(fields, last.field)
		last = last.sub
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
