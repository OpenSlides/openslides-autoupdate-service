package keysrequest

import (
	"encoding/json"
	"io"
)

// FromJSON creates a KeysRequest from json
func FromJSON(r io.Reader) (KeysRequest, error) {
	var kr KeysRequest
	if err := json.NewDecoder(r).Decode(&kr); err != nil {
		return kr, ErrJSON{msg: "can not decode key request from json", err: err}
	}
	return kr, kr.Validate()
}
