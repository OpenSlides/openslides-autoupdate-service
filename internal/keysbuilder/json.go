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
		return nil, ErrJSON{msg: "can not decode keysbuilder from json", err: err}
	}
	if err := b.validate(); err != nil {
		return nil, err
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
		return nil, ErrJSON{msg: "can not decode many keysbuilder from json", err: err}
	}
	for _, b := range bs {
		if err := b.validate(); err != nil {
			return nil, err
		}
	}

	kb, err := newBuilder(ctx, ider, bs...)
	if err != nil {
		return nil, fmt.Errorf("can not build keys: %w", err)
	}
	return kb, nil
}
