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
		if sub, ok := err.(ErrInvalid); ok {
			return nil, sub
		}
		return nil, ErrJSON{msg: "can not decode keysbuilder from json", err: err}
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
		if sub, ok := err.(ErrInvalid); ok {
			return nil, sub
		}
		return nil, ErrJSON{msg: "can not decode many keysbuilder from json", err: err}
	}

	kb, err := newBuilder(ctx, ider, bs...)
	if err != nil {
		return nil, fmt.Errorf("can not build keys: %w", err)
	}
	return kb, nil
}
