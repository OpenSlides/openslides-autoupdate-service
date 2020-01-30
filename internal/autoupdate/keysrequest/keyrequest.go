// Package keysrequest holds the keyrequest object and functions
// to generate it. Currently, it can only be generated from json.
package keysrequest

import (
	"strings"
)

// Body holds the information what keys are requested by the client
type Body struct {
	IDs []int `json:"ids"`
	Fields
}

// validate maks sure the KeysRequest is valid. Returns an ErrInvalid if not.
func (kr *Body) validate() error {
	if len(kr.IDs) == 0 {
		return ErrInvalid{msg: "no ids"}
	}
	return kr.Fields.validate()
}

// Fields describes in a abstract way fields of a collection.
// It can map to related keys.
type Fields struct {
	Collection string            `json:"collection"`
	Names      map[string]Fields `json:"fields"`
}

// Null returns true if fieldDescription is empty (null in json)
func (fd *Fields) Null() bool {
	return fd.Collection == "" && len(fd.Names) == 0
}

func (fd *Fields) validate() error {
	if len(fd.Names) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	if fd.Collection == "" {
		return ErrInvalid{msg: "no collection"}
	}
	for name, description := range fd.Names {
		if description.Null() {
			continue
		}
		if !(strings.HasSuffix(name, "_id") || strings.HasSuffix(name, "_ids")) {
			return ErrInvalid{msg: "relation but no _id or _ids suffix", field: name}
		}
		if err := description.validate(); err != nil {
			sub := err.(ErrInvalid)
			return ErrInvalid{sub: &sub, field: name, msg: "Error on field"}
		}
	}
	return nil
}
