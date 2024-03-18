package dsfetch

import "encoding/json"

// Maybe holds a type or null.
type Maybe[T any] struct {
	hasValue bool
	value    T
}

// MaybeValue initializes a Maybe with a value.
func MaybeValue[T any](v T) Maybe[T] {
	return Maybe[T]{
		hasValue: true,
		value:    v,
	}
}

func (m Maybe[T]) Value() (T, bool) {
	if m.hasValue {
		return m.value, true
	}
	var zero T
	return zero, false
}

func (m Maybe[T]) Null() bool {
	return !m.hasValue
}

func (m *Maybe[T]) Set(v T) {
	m.value = v
	m.hasValue = true
}

func (m *Maybe[T]) SetNull() {
	m.hasValue = false
}

func (m *Maybe[T]) UnmarshalJSON(bs []byte) error {
	if string(bs) == "null" {
		m.SetNull()
		return nil
	}

	var v T
	if err := json.Unmarshal(bs, &v); err != nil {
		return err
	}
	m.Set(v)
	return nil
}
