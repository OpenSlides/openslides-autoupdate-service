package keysrequest

import (
	"encoding/json"
	"io"
)

// FromJSON creates a KeysRequest from json
func FromJSON(r io.Reader) (Body, error) {
	var kr Body
	if err := json.NewDecoder(r).Decode(&kr); err != nil {
		return kr, ErrJSON{msg: "can not decode keysrequest from json", err: err}
	}
	return kr, kr.validate()
}

// ManyFromJSON creates a list of KeysRequest from json
func ManyFromJSON(r io.Reader) ([]Body, error) {
	var krs []Body
	if err := json.NewDecoder(r).Decode(&krs); err != nil {
		return nil, ErrJSON{msg: "can not decode many keysrequest from json", err: err}
	}
	for _, kr := range krs {
		if err := kr.validate(); err != nil {
			return nil, err
		}
	}
	return krs, nil
}
