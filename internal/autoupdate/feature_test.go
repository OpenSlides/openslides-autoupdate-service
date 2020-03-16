package autoupdate_test

import (
	"context"
	"strings"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

var dataSet = map[string]string{
	"A/1/a":                   `"a1"`,
	"A/1/title":               `"a1"`,
	"A/1/B_id":                `1`,
	"A/1/C_ids":               `[]`,
	"A/1/G1_ids":              `[1, 2]`,
	"A/2/a":                   `"a2"`,
	"A/2/title":               `"a2"`,
	"A/2/C_ids":               `[1, 2]`,
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
	"B/2/C_ids":               `[1, 2]`,
	"B/2/B_parent_id":         `1`,
	"B/2/B_children_ids":      `[]`,
	"B/2/D_ids":               `[1, 2]`,
	"C/1/c":                   `"c1"`,
	"C/1/title":               `"c1"`,
	"C/1/A_id":                `2`,
	"C/1/B_ids":               `[1, 2]`,
	"C/1/G1_ids":              `[2, 3]`,
	"C/2/c":                   `"c2"`,
	"C/2/title":               `"c2"`,
	"C/2/A_id":                `2`,
	"C/2/B_ids":               `[2]`,
	"C/2/G1_ids":              `[2, 3]`,
	"D/1/d":                   `"d1"`,
	"D/1/B_$_ids":             `["1", "2", "3"]`,
	"D/1/B_1_ids":             `[1, 2]`,
	"D/1/B_2_ids":             `[1]`,
	"D/1/B_3_ids":             `[]`,
	"D/2/d":                   `"d2"`,
	"D/2/B_$_ids":             `["1", "4"]`,
	"D/2/B_1_ids":             `[]`,
	"D/2/B_4_ids":             `[2]`,
	"G1/1/g1":                 `"g1.1"`,
	"G1/1/content_object_ids": `["A/1"]`,
	"G1/2/g1":                 `"g1.2"`,
	"G1/2/content_object_ids": `["A/1", "C/1", "C/2"]`,
	"G1/3/g1":                 `"g1.3"`,
	"G1/3/content_object_ids": `["C/1", "C/2"]`,
	"G2/1/g2":                 `"g2.1"`,
	"G2/1/content_object_id":  `"B/1"`,
}

func TestFeatures(t *testing.T) {
	m := &datasetMock{}
	s := autoupdate.New(m, m)
	defer s.Close()

	for _, tt := range []struct {
		name    string
		request string
		data    map[string]string
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
			map[string]string{
				"A/1/a":      `"a1"`,
				"A/1/C_ids":  `[]`,
				"A/1/B_id":   `1`,
				"A/1/G1_ids": `[1, 2]`,
				"A/2/a":      `"a2"`,
				"A/2/C_ids":  `[1, 2]`,
				"A/2/G1_ids": `[]`,
				"C/1/c":      `"c1"`,
				"C/1/G1_ids": `[2, 3]`,
				"C/2/c":      `"c2"`,
				"C/2/G1_ids": `[2, 3]`,
				"G1/1/g1":    `"g1.1"`,
				"G1/2/g1":    `"g1.2"`,
				"G1/3/g1":    `"g1.3"`,
			},
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
			map[string]string{
				"B/1/B_children_ids":     `[2]`,
				"B/1/C_ids":              `[1]`,
				"B/1/G2_id":              `1`,
				"B/2/C_ids":              `[1, 2]`,
				"B/2/B_parent_id":        `1`,
				"G2/1/content_object_id": `"B/1"`,
				"C/1/c":                  `"c1"`,
				"C/1/title":              `"c1"`,
				"C/2/c":                  `"c2"`,
			},
		},
		{
			"non-existent ids, fields, fqids, references, generic relations and fields without a relation",
			`{
				"collection": "G1",
				"ids": [2, 4],
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
			map[string]string{
				"G1/2/content_object_ids": `["A/1", "C/1", "C/2"]`,
				"A/1/a":                   `"a1"`,
				"A/1/title":               `"a1"`,
				"A/1/G1_ids":              `[1, 2]`,
				"C/1/title":               `"c1"`,
				"C/1/A_id":                `2`,
				"C/1/G1_ids":              `[2, 3]`,
				"C/2/title":               `"c2"`,
				"C/2/A_id":                `2`,
				"C/2/G1_ids":              `[2, 3]`,
			},
		},
		{
			"template fields",
			`{
				"collection": "D",
				"ids": [1, 2],
				"fields": {
					"d": null,
					"B_$_ids": null
				}
			}`,
			map[string]string{
				"D/1/d":       `"d1"`,
				"D/1/B_$_ids": `["1", "2", "3"]`,
				"D/2/d":       `"d2"`,
				"D/2/B_$_ids": `["1", "4"]`,
			},
		},
		{
			"structured fields without references",
			`{
				"collection": "D",
				"ids": [1, 2],
				"fields": {
					"d": null,
					"B_$_ids": {
						"type": "template"
					}
				}
			}`,
			map[string]string{
				"D/1/d":       `"d1"`,
				"D/1/B_$_ids": `["1", "2", "3"]`,
				"D/1/B_1_ids": `[1, 2]`,
				"D/1/B_2_ids": `[1]`,
				"D/1/B_3_ids": `[]`,
				"D/2/d":       `"d2"`,
				"D/2/B_$_ids": `["1", "4"]`,
				"D/2/B_1_ids": `[]`,
				"D/2/B_4_ids": `[2]`,
			},
		},
		{
			"structed references",
			`{
				"collection": "D",
				"ids": [1, 2],
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
					"B_4_ids": {
						"type": "relation-list",
						"collection": "B",
						"fields": {
							"title": null
						}
					}
				}
			}`,
			map[string]string{
				"D/1/B_$_ids": `["1", "2", "3"]`,
				"D/1/B_1_ids": `[1, 2]`,
				"D/1/B_2_ids": `[1]`,
				"D/1/B_3_ids": `[]`,
				"D/2/B_$_ids": `["1", "4"]`,
				"D/2/B_1_ids": `[]`,
				"D/2/B_4_ids": `[2]`,
				"B/1/b":       `"b1"`,
				"B/2/b":       `"b2"`,
				"B/2/title":   `"b2"`,
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			b, err := keysbuilder.FromJSON(context.Background(), strings.NewReader(tt.request), s.IDer(1))
			if err != nil {
				t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
			}
			c := s.Connect(context.Background(), 1, b)
			if next := c.Next(); !next {
				t.Fatalf("Expected connection to have data, got err: %v", c.Err())
			}
			data := c.Data()

			if diff := cmpMap(data, tt.data); !diff {
				t.Errorf("Expected %v, got: %v", tt.data, data)
			}
		})
	}
}

type datasetMock struct{}

func (d *datasetMock) Restrict(ctx context.Context, uid int, keys []string) (map[string]string, error) {
	out := make(map[string]string)
	for _, key := range keys {
		value, ok := dataSet[key]
		if !ok {
			continue
		}
		out[key] = value
	}
	return out, nil
}

func (d *datasetMock) KeysChanged() ([]string, error) {
	select {}
}

func cmpMap(one, two map[string]string) bool {
	if len(one) != len(two) {
		return false
	}
	for key := range one {
		tv, ok := two[key]
		if !ok {
			return false
		}
		if one[key] != tv {
			return false
		}
	}
	return true
}
