package redis

import (
	"encoding/json"
	"sort"
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
					["updated", "key1", "updated", "key2"]
				],
				[
					"12346-0",
					["updated", "key1", "updated", "key3"]
				]
			]
		]
	]`), &data)
	if err != nil {
		t.Fatalf("Data is invalid json: %v", err)
	}

	id, keys, err := stream(data, nil)
	if err != nil {
		t.Errorf("Did not expect an error, got: %v", err)
	}
	expect := []string{"key1", "key2", "key3"}
	if !cmpSlice(keys, expect) {
		t.Errorf("Expected %v, got %v", expect, keys)
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
		{"Outer list", `"data"`, "Invalid input. Data has to be a list"},
		{"One stream", `[]`, "Invalid input. No stream in data"},
		{"Stream no list", `["data"]`, "Invalid input. Stream has to be a two-tuple"},
		{"Stream no elements", `[[]]`, "Invalid input. Stream has to be a two-tuple"},
		{"Stream one element", `[["one"]]`, "Invalid input. Stream has to be a two-tuple"},
		{"Stream tree elements", `[["one", "two", "tree"]]`, "Invalid input. Stream has to be a two-tuple"},
		{"Stream data no list", `[["one", "two"]]`, "Invalid input. Stream data has to be a list"},
		{"Stream element no list", `[["one", ["data"]]]`, "Invalid input. Stream element has to be a two-tuple"},
		{"Stream element no elements", `[["one", [[]]]]`, "Invalid input. Stream element has to be a two-tuple"},
		{"Stream element one element", `[["one", [["one"]]]]`, "Invalid input. Stream element has to be a two-tuple"},
		{"Stream element tree elements", `[["one", [["one", "two", "tree"]]]]`, "Invalid input. Stream element has to be a two-tuple"},
		{"id no string", `[["one", [[123, ["data"]]]]]`, "Invalid input. Stream ID has to be a string"},
		{"key-value no string list", `[["one", [["123", "data"]]]]`, "Invalid input. Key values has to be a list of strings"},
		{"Odd key value", `[["one", [["123", ["1"]]]]]`, "Invalid input. Odd number of key value pairs"},
		{"Key no string", `[["one", [["123", [1, "2"]]]]]`, "Invalid input. Key has to be a string"},
		{"Value no string", `[["one", [["123", ["1", 2]]]]]`, "Invalid input. Values has to be a string"},
		{"unknown key", `[["one", [["123", ["data", "value"]]]]]`, "Invalid input. Unknown key \"data\""},
	}
	for _, tt := range td {
		t.Run(tt.name, func(t *testing.T) {
			var data interface{}
			err := json.Unmarshal([]byte(tt.json), &data)
			if err != nil {
				t.Fatalf("Data is invalid json: %v", err)
			}

			_, _, err = stream(data, nil)
			if err == nil {
				t.Fatalf("Expected an error, got none")
			}
			if got := err.Error(); !strings.HasPrefix(got, tt.err) {
				t.Errorf("Expect error message to be \"%s\", got: %v", tt.err, got)
			}
		})
	}

}

func cmpSlice(one, two []string) bool {
	if len(one) != len(two) {
		return false
	}

	sort.Strings(one)
	sort.Strings(two)
	for i := range one {
		if one[i] != two[i] {
			return false
		}
	}
	return true
}
