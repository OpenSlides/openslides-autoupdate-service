package redis

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

func TestStream(t *testing.T) {
	var data any
	err := json.Unmarshal([]byte(`
	[
		[
			"ModifiedFields",
			[
				[
					"12345-0",
					["user/1/username", "Helga", "user/2/username", "Isolde"]
				],
				[
					"12346-0",
					["user/1/username", "Hubert", "user/3/username", "Igor"]
				]
			]
		]
	]`), &data)
	if err != nil {
		t.Fatalf("Data is invalid json: %v", err)
	}

	id, got, err := parseMessageBus(data)
	if err != nil {
		t.Errorf("Returned unexpected error %v", err)
	}

	expect := map[dskey.Key][]byte{
		dskey.MustKey("user/1/username"): []byte("Hubert"),
		dskey.MustKey("user/2/username"): []byte("Isolde"),
		dskey.MustKey("user/3/username"): []byte("Igor"),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got %v, expected %v", got, expect)
	}

	if id != "12346-0" {
		t.Errorf("Expected id to be 12346-0, got: %v", id)
	}
}

func TestStreamInvalidData(t *testing.T) {
	td := []struct {
		name string
		json string
	}{
		{"Outer list", `"data"`},
		{"One stream", `[]`},
		{"Stream no list", `["data"]`},
		{"Stream no elements", `[[]]`},
		{"Stream one element", `[["one"]]`},
		{"Stream tree elements", `[["one", "two", "tree"]]`},
		{"Stream data no list", `[["one", "two"]]`},
		{"Stream element no list", `[["one", ["data"]]]`},
		{"Stream element no elements", `[["one", [[]]]]`},
		{"Stream element one element", `[["one", [["one"]]]]`},
		{"Stream element tree elements", `[["one", [["one", "two", "tree"]]]]`},
		{"id no string", `[["one", [[123, ["data"]]]]]`},
		{"key-value no string list", `[["one", [["123", "data"]]]]`},
		{"Odd key value", `[["one", [["123", ["1"]]]]]`},
		{"Key no string", `[["one", [["123", [1, "2"]]]]]`},
	}
	for _, tt := range td {
		t.Run(tt.name, func(t *testing.T) {
			var data any
			err := json.Unmarshal([]byte(tt.json), &data)
			if err != nil {
				t.Fatalf("Data is invalid json: %v", err)
			}

			_, _, err = parseMessageBus(data)
			if err == nil {
				t.Fatalf("Expected an error, got none")
			}
		})
	}
}
