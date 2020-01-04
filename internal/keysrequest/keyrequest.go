// Package keysrequest holds the keyrequest object and functions
// to generate it. Currently, it can only be generated from json.
package keysrequest

import (
	"strings"
)

// FieldDescription describes in a abstract way fields of a collection.
// It can also map to related keys.
type FieldDescription struct {
	Collection string                      `json:"collection"`
	Fields     map[string]FieldDescription `json:"fields"`
}

// Null returns true if fieldDescription is empty (null in json)
func (fd *FieldDescription) Null() bool {
	return fd.Collection == "" && len(fd.Fields) == 0
}

func (fd *FieldDescription) validate() error {
	if len(fd.Fields) == 0 {
		return ErrInvalid{msg: "no fields"}
	}
	if fd.Collection == "" {
		return ErrInvalid{msg: "no collection"}
	}
	for name, description := range fd.Fields {
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

// KeysRequest holds the information what keys are requested by the client
// meetingID is ignored, if ids != nil
type KeysRequest struct {
	IDs       []int `json:"ids"`
	MeetingID int   `json:"meeting_id"`
	FieldDescription
}

// Validate maks sure the KeysRequest is valid. Returns an ErrInvalid if not.
func (kr *KeysRequest) Validate() error {
	return kr.FieldDescription.validate()
}
