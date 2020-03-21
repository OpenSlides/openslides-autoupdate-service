package keysbuilder

import (
	"fmt"
	"strings"
)

// ErrInvalid is an error that happens on an invalid request.
type ErrInvalid struct {
	msg   string
	field string
	sub   *ErrInvalid
}

func (e ErrInvalid) Error() string {
	fields, last := e.fields()
	if fields == nil {
		return e.msg
	}
	return fmt.Sprintf("field \"%s\": %s", strings.Join(fields, "."), last.msg)
}

// Type returns the name of the error.
func (e ErrInvalid) Type() string {
	return "SyntaxError"
}

// Fields returns a list of field names from the parent to this error.
func (e ErrInvalid) Fields() []string {
	fields, _ := e.fields()
	return fields
}

func (e ErrInvalid) fields() ([]string, *ErrInvalid) {
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

// ErrJSON is returned when invalid json is parsed or the json can not
// be decoded to a keysbuilder.
type ErrJSON struct {
	err error
}

func (e ErrJSON) Error() string {
	return e.err.Error()
}

// Unwrap returns the thrown error.
func (e ErrJSON) Unwrap() error {
	return e.err
}

// Type returns the name of the error.
func (e ErrJSON) Type() string {
	return "JsonError"
}
