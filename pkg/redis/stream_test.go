package redis

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestStream(t *testing.T) {
	var data interface{}
	err := json.Unmarshal([]byte(`
	[
		[
			"stream1",
			[
				[
					"12345-0",
					["user/1/name", "Helga", "user/2/name", "Isolde"]
				],
				[
					"12346-0",
					["user/1/name", "Hubert", "user/3/name", "Igor"]
				]
			]
		]
	]`), &data)
	if err != nil {
		t.Fatalf("Data is invalid json: %v", err)
	}

	id, retData, err := autoupdateStream(data, nil)
	if err != nil {
		t.Errorf("Returned unexpected error %v", err)
	}
	expect := map[string]json.RawMessage{
		"user/1/name": []byte("Hubert"),
		"user/2/name": []byte("Isolde"),
		"user/3/name": []byte("Igor"),
	}
	if !cmpMap(retData, expect) {
		t.Errorf("Got %v, expected %v", retData, expect)
	}
	if id != "12346-0" {
		t.Errorf("Expected id to be 12346-0, got: %v", id)
	}
}

func TestStreamInvalidData(t *testing.T) {
	td := []struct {
		name string
		json string
		err  string
	}{
		{"Outer list", `"data"`, "invalid input. Data has to be a list"},
		{"One stream", `[]`, "invalid input. No stream in data"},
		{"Stream no list", `["data"]`, "invalid input. Stream has to be a two-tuple"},
		{"Stream no elements", `[[]]`, "invalid input. Stream has to be a two-tuple"},
		{"Stream one element", `[["one"]]`, "invalid input. Stream has to be a two-tuple"},
		{"Stream tree elements", `[["one", "two", "tree"]]`, "invalid input. Stream has to be a two-tuple"},
		{"Stream data no list", `[["one", "two"]]`, "invalid input. Stream data has to be a list"},
		{"Stream element no list", `[["one", ["data"]]]`, "invalid input. Stream element has to be a two-tuple"},
		{"Stream element no elements", `[["one", [[]]]]`, "invalid input. Stream element has to be a two-tuple"},
		{"Stream element one element", `[["one", [["one"]]]]`, "invalid input. Stream element has to be a two-tuple"},
		{"Stream element tree elements", `[["one", [["one", "two", "tree"]]]]`, "invalid input. Stream element has to be a two-tuple"},
		{"id no string", `[["one", [[123, ["data"]]]]]`, "invalid input. Stream ID has to be a string"},
		{"key-value no string list", `[["one", [["123", "data"]]]]`, "invalid input. Key values has to be a list of strings"},
		{"Odd key value", `[["one", [["123", ["1"]]]]]`, "invalid input. Odd number of key value pairs"},
		{"Key no string", `[["one", [["123", [1, "2"]]]]]`, "invalid input. Key has to be a string"},
	}
	for _, tt := range td {
		t.Run(tt.name, func(t *testing.T) {
			var data interface{}
			err := json.Unmarshal([]byte(tt.json), &data)
			if err != nil {
				t.Fatalf("Data is invalid json: %v", err)
			}

			_, _, err = autoupdateStream(data, nil)
			if err == nil {
				t.Fatalf("Expected an error, got none")
			}
			if got := err.Error(); !strings.HasPrefix(got, tt.err) {
				t.Errorf("Expect error message to be \"%s\", got: %v", tt.err, got)
			}
		})
	}
}

func cmpMap(one, two map[string]json.RawMessage) bool {
	if len(one) != len(two) {
		return false
	}

	for key := range one {
		if bytes.Compare(one[key], two[key]) != 0 {
			return false
		}
	}
	return true
}
