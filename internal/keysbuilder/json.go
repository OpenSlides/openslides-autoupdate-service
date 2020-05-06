package keysbuilder

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// FromJSON creates a Keysbuilder from json.
func FromJSON(ctx context.Context, r io.Reader, valuer Valuer, uid int) (*Builder, error) {
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

	kb, err := newBuilder(ctx, valuer, uid, b)
	if err != nil {
		return nil, fmt.Errorf("build keys: %w", err)
	}
	return kb, nil
}

// ManyFromJSON creates a list of Keysbuilder objects from a json list.
func ManyFromJSON(ctx context.Context, r io.Reader, valuer Valuer, uid int) (*Builder, error) {
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
			return nil, InvalidError{msg: fmt.Sprintf("wrong format at byte %d", jerr.Offset)}
		}
		return nil, fmt.Errorf("decode keysrequest: %w", err)
	}

	if len(bs) == 0 {
		return nil, InvalidError{msg: "No data"}
	}

	kb, err := newBuilder(ctx, valuer, uid, bs...)
	if err != nil {
		return nil, fmt.Errorf("build keys: %w", err)
	}
	return kb, nil
}
