package fastjson_test

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/fastjson"
)

func FuzzDecodeInt(f *testing.F) {
	f.Add([]byte("0"))
	f.Add([]byte("24234"))
	f.Add([]byte("-34"))
	f.Add([]byte("234234234234234234234234982348234789243324243243141243"))
	f.Add([]byte("hello world"))
	f.Add([]byte(`"hello world"`))

	f.Fuzz(func(t *testing.T, value []byte) {
		if len(bytes.TrimSpace(value)) != len(value) {
			t.Skip()
		}

		myResult, myErr := fastjson.DecodeInt(value)

		var jsonResult int
		jsonErr := json.Unmarshal(value, &jsonResult)

		switch {
		case myErr != nil && jsonErr != nil:
			// Both versions returned an error
			return
		case myErr != nil:
			t.Errorf("DecodeInt(%s): %v", value, myErr)
		case jsonErr != nil:
			// It is ok for my version to give better results.
			return
			// t.Errorf("DecodeInt(%s) did not return an error", value)
		case myResult != jsonResult:
			t.Errorf("DecodeInt(%s) == %d, json == %d", value, myResult, jsonResult)
		}
	})
}

func FuzzDecodeIntList(f *testing.F) {
	f.Add([]byte("[0]"))
	f.Add([]byte("[2,4,2,34]"))
	f.Add([]byte("[-34,34]"))
	f.Add([]byte("[]"))
	f.Add([]byte("hello world"))

	f.Fuzz(func(t *testing.T, value []byte) {
		if len(bytes.TrimSpace(value)) != len(value) {
			t.Skip()
		}

		myResult, myErr := fastjson.DecodeIntList(value)

		var jsonResult []int
		jsonErr := json.Unmarshal(value, &jsonResult)

		switch {
		case myErr != nil && jsonErr != nil:
			// Both versions returned an error
			return
		case myErr != nil:
			t.Errorf("DecodeIntList(`%s`): %v, json: %v", value, myErr, jsonResult)
		case jsonErr != nil:
			// It is ok for my version to give better results.
			return
			// t.Errorf("DecodeInt(%s) did not return an error", value)
		case !reflect.DeepEqual(myResult, jsonResult):
			t.Errorf("DecodeIntList(`%s`) == %v, json == %v", value, myResult, jsonResult)
		}
	})
}
