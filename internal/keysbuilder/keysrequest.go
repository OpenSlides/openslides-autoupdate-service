package keysbuilder

import (
	"strings"
)

// body holds the information what keys are requested by the client.
type body struct {
	IDs []int `json:"ids"`
	fields
}

// validate maks sure the body is valid. Returns an ErrInvalid if not.
func (kr *body) validate() error {
	if len(kr.IDs) == 0 {
		return ErrInvalid{msg: "no ids"}
	}
	return kr.fields.validate()
}

// fields describes in a abstract way fields of a collection.
// It can map to related keys.
type fields struct {
	Collection string            `json:"collection"`
	Names      map[string]fields `json:"fields"`
}

// Null returns true if fieldDescription is empty (null in json)
func (fd *fields) null() bool {
	return fd.Collection == "" && len(fd.Names) == 0
}

// validate maks sure the fields are valid. Returns an ErrInvalid if not.
func (fd *fields) validate() error {
	if len(fd.Names) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	if fd.Collection == "" {
		return ErrInvalid{msg: "no collection"}
	}
	for name, description := range fd.Names {
		if description.null() {
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
