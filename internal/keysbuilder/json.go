package keysbuilder

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
)

// FromJSON creates a Keysbuilder from json.
func FromJSON(ctx context.Context, r io.Reader, ider IDer) (*Builder, error) {
	var b body
	if err := json.NewDecoder(r).Decode(&b); err != nil {
		if err == io.EOF {
			return nil, ErrInvalid{msg: "No data"}
		}
		if sub, ok := err.(ErrInvalid); ok {
			return nil, sub
		}
		return nil, ErrJSON{err}
	}

	kb, err := newBuilder(ctx, ider, b)
	if err != nil {
		return nil, fmt.Errorf("can not build keys: %w", err)
	}
	return kb, nil
}

// ManyFromJSON creates a Keysbuilder from a json list.
func ManyFromJSON(ctx context.Context, r io.Reader, ider IDer) (*Builder, error) {
	var bs []body
	if err := json.NewDecoder(r).Decode(&bs); err != nil {
		if err == io.EOF {
			return nil, ErrInvalid{msg: "No data"}
		}
		if sub, ok := err.(ErrInvalid); ok {
			return nil, sub
		}
		if jerr, ok := err.(*json.SyntaxError); ok {
			return nil, ErrJSON{jerr}
		}
		if jerr, ok := err.(*json.UnmarshalTypeError); ok {
			return nil, ErrInvalid{msg: fmt.Sprintf("wrong format at byte %d", jerr.Offset)}
		}
		return nil, fmt.Errorf("can not decode keysrequest: %v", err)
	}

	if len(bs) == 0 {
		return nil, ErrInvalid{msg: "No data"}
	}

	kb, err := newBuilder(ctx, ider, bs...)
	if err != nil {
		return nil, fmt.Errorf("can not build keys: %w", err)
	}
	return kb, nil
}
