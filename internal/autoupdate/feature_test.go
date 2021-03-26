package autoupdate_test

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

var dataSet = map[string]string{
	"A/1/a":                   `"a1"`,
	"A/1/title":               `"a1"`,
	"A/1/B_id":                `1`,
	"A/1/C_ids":               `[]`,
	"A/1/G1_ids":              `[1,2]`,
	"A/2/a":                   `"a2"`,
	"A/2/title":               `"a2"`,
	"A/2/C_ids":               `[1,2]`,
	"A/2/G1_ids":              `[]`,
	"B/1/b":                   `"b1"`,
	"B/1/title":               `"b1"`,
	"B/1/A_id":                `1`,
	"B/1/C_ids":               `[1]`,
	"B/1/G2_id":               `1`,
	"B/1/B_children_ids":      `[2]`,
	"B/1/D_ids":               `[1]`,
	"B/2/b":                   `"b2"`,
	"B/2/title":               `"b2"`,
	"B/2/C_ids":               `[1,2]`,
	"B/2/B_parent_id":         `1`,
	"B/2/B_children_ids":      `[]`,
	"B/2/D_ids":               `[1,2]`,
	"C/1/c":                   `"c1"`,
	"C/1/title":               `"c1"`,
	"C/1/A_id":                `2`,
	"C/1/B_ids":               `[1,2]`,
	"C/1/G1_ids":              `[2,3]`,
	"C/2/c":                   `"c2"`,
	"C/2/title":               `"c2"`,
	"C/2/A_id":                `2`,
	"C/2/B_ids":               `[2]`,
	"C/2/G1_ids":              `[2,3]`,
	"D/1/d":                   `"d1"`,
	"D/1/B_$_ids":             `["1","2","3"]`,
	"D/1/B_$1_ids":            `[1,2]`,
	"D/1/B_$2_ids":            `[1]`,
	"D/1/B_$3_ids":            `[]`,
	"D/2/d":                   `"d2"`,
	"D/2/B_$_ids":             `["1","4"]`,
	"D/2/B_$1_ids":            `[]`,
	"D/2/B_$4_ids":            `[2]`,
	"G1/1/g1":                 `"g1.1"`,
	"G1/1/content_object_ids": `["A/1"]`,
	"G1/2/g1":                 `"g1.2"`,
	"G1/2/content_object_ids": `["A/1","C/1","C/2"]`,
	"G1/3/g1":                 `"g1.3"`,
	"G1/3/content_object_ids": `["C/1","C/2"]`,
	"G2/1/g2":                 `"g2.1"`,
	"G2/1/content_object_id":  `"B/1"`,
}

func TestFeatures(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	datastore := test.NewMockDatastore(closed, dataSet)
	s := autoupdate.New(datastore, new(test.MockRestricter), test.UserUpdater{}, closed)

	for _, tt := range []struct {
		name string

		// request is an example request to the autoupdate service. The
		// http-server expects not one request, but a list of requests.
		request string

		// result is the data returned for the request. The http-server returns
		// the same json encoded data but its format differs. It only returns
		// one line without spaces.
		result string
	}{
		{
			"Basic",
			`{
				"collection": "A",
				"ids": [
					1,
					2
				],
				"fields": {
					"a": null,
					"C_ids": {
						"type": "relation-list",
						"collection": "C",
						"fields": {
							"c": null,
							"G1_ids": {
								"type": "relation-list",
								"collection": "G1",
								"fields": {
									"g1": null
								}
							}
						}
					},
					"B_id": {
						"type": "relation",
						"collection": "B",
						"fields": {}
					},
					"G1_ids": {
						"type": "relation-list",
						"collection": "G1",
						"fields": {
							"g1": null
						}
					}
				}
			}`,
			`{
				"A/1/a":      "a1",
				"A/1/C_ids":  [],
				"A/1/B_id":   1,
				"A/1/G1_ids": [1,2],
				"A/2/a":      "a2",
				"A/2/C_ids":  [1,2],
				"A/2/G1_ids": [],
				"C/1/c":      "c1",
				"C/1/G1_ids": [2,3],
				"C/2/c":      "c2",
				"C/2/G1_ids": [2,3],
				"G1/1/g1":    "g1.1",
				"G1/2/g1":    "g1.2",
				"G1/3/g1":    "g1.3"
			}`,
		},
		{
			"Partial merged fields, generic lookup",
			`{
				"collection": "G2",
				"ids": [1],
				"fields": {
					"content_object_id": {
						"type": "generic-relation",
						"fields": {
							"B_children_ids": {
								"type": "relation-list",
								"collection": "B",
								"fields": {
									"C_ids": {
										"type": "relation-list",
										"collection": "C",
										"fields": {
											"c": null
										}
									},
									"B_parent_id": null
								}
							},
							"C_ids": {
								"type": "relation-list",
								"collection": "C",
								"fields": {
									"c": null,
									"title": null
								}
							},
							"G2_id": null
						}
					}
				}
			}`,
			`{
				"B/1/B_children_ids":     [2],
				"B/1/C_ids":              [1],
				"B/1/G2_id":              1,
				"B/2/C_ids":              [1,2],
				"B/2/B_parent_id":        1,
				"G2/1/content_object_id": "B/1",
				"C/1/c":                  "c1",
				"C/1/title":              "c1",
				"C/2/c":                  "c2"
			}`,
		},
		{
			"non-existent ids, fields, fqids, references, generic relations and fields without a relation",
			`{
				"collection": "G1",
				"ids": [2,4],
				"fields": {
					"content_object_ids": {
						"type": "generic-relation-list",
						"fields": {
							"a": null,
							"b": null,
							"not_existent": {
								"type": "generic-relation",
								"fields": {"key": null}
							},
							"title": null,
							"G1_ids": null,
							"A_id": null
						}
					}
				}
			}`,
			`{
				"G1/2/content_object_ids": ["A/1","C/1","C/2"],
				"A/1/a":                   "a1",
				"A/1/title":               "a1",
				"A/1/G1_ids":              [1,2],
				"C/1/title":               "c1",
				"C/1/A_id":                2,
				"C/1/G1_ids":              [2,3],
				"C/2/title":               "c2",
				"C/2/A_id":                2,
				"C/2/G1_ids":              [2,3]
			}`,
		},
		{
			"template fields",
			`{
				"collection": "D",
				"ids": [1,2],
				"fields": {
					"d": null,
					"B_$_ids": null
				}
			}`,
			`{
				"D/1/d":       "d1",
				"D/1/B_$_ids": ["1","2","3"],
				"D/2/d":       "d2",
				"D/2/B_$_ids": ["1","4"]
			}`,
		},
		{
			"structured fields without references",
			`{
				"collection": "D",
				"ids": [1,2],
				"fields": {
					"d": null,
					"B_$_ids": {
						"type": "template"
					}
				}
			}`,
			`{
				"D/1/d":       "d1",
				"D/1/B_$_ids": ["1","2","3"],
				"D/1/B_$1_ids": [1,2],
				"D/1/B_$2_ids": [1],
				"D/1/B_$3_ids": [],
				"D/2/d":       "d2",
				"D/2/B_$_ids": ["1","4"],
				"D/2/B_$1_ids": [],
				"D/2/B_$4_ids": [2]
			}`,
		},
		{
			"structed references",
			`{
				"collection": "D",
				"ids": [1,2],
				"fields": {
					"B_$_ids": {
						"type": "template",
						"values": {
							"type": "relation-list",
							"collection": "B",
							"fields": {
								"b": null
							}
						}
					},
					"B_$4_ids": {
						"type": "relation-list",
						"collection": "B",
						"fields": {
							"title": null
						}
					}
				}
			}`,
			`{
				"D/1/B_$_ids": ["1","2","3"],
				"D/1/B_$1_ids": [1,2],
				"D/1/B_$2_ids": [1],
				"D/1/B_$3_ids": [],
				"D/2/B_$_ids": ["1","4"],
				"D/2/B_$1_ids": [],
				"D/2/B_$4_ids": [2],
				"B/1/b":       "b1",
				"B/2/b":       "b2",
				"B/2/title":   "b2"
			}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			b, err := keysbuilder.FromJSON(strings.NewReader(tt.request), s, 1)
			if err != nil {
				t.Fatalf("FromJSON() returned an unexpected error: %v", err)
			}
			c := s.Connect(1, b)
			data, err := c.Next(context.Background())
			if err != nil {
				t.Fatalf("Can not get data: %v", err)
			}

			var expect map[string]json.RawMessage
			if err := json.Unmarshal([]byte(tt.result), &expect); err != nil {
				t.Fatalf("Can not decode keys from expected data: %v", err)
			}

			cmpMap(t, data, expect)
		})
	}
}

func cmpMap(t *testing.T, got, expect map[string]json.RawMessage) {
	v1, _ := json.Marshal(got)
	v2, _ := json.Marshal(expect)
	if string(v1) != string(v2) {
		t.Errorf("got %s, expected %s", string(v1), string(v2))
	}
}
