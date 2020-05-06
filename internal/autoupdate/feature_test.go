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

var dataSet = map[string]json.RawMessage{
	"A/1/a":                   []byte(`"a1"`),
	"A/1/title":               []byte(`"a1"`),
	"A/1/B_id":                []byte(`1`),
	"A/1/C_ids":               []byte(`[]`),
	"A/1/G1_ids":              []byte(`[1,2]`),
	"A/2/a":                   []byte(`"a2"`),
	"A/2/title":               []byte(`"a2"`),
	"A/2/C_ids":               []byte(`[1,2]`),
	"A/2/G1_ids":              []byte(`[]`),
	"B/1/b":                   []byte(`"b1"`),
	"B/1/title":               []byte(`"b1"`),
	"B/1/A_id":                []byte(`1`),
	"B/1/C_ids":               []byte(`[1]`),
	"B/1/G2_id":               []byte(`1`),
	"B/1/B_children_ids":      []byte(`[2]`),
	"B/1/D_ids":               []byte(`[1]`),
	"B/2/b":                   []byte(`"b2"`),
	"B/2/title":               []byte(`"b2"`),
	"B/2/C_ids":               []byte(`[1,2]`),
	"B/2/B_parent_id":         []byte(`1`),
	"B/2/B_children_ids":      []byte(`[]`),
	"B/2/D_ids":               []byte(`[1,2]`),
	"C/1/c":                   []byte(`"c1"`),
	"C/1/title":               []byte(`"c1"`),
	"C/1/A_id":                []byte(`2`),
	"C/1/B_ids":               []byte(`[1,2]`),
	"C/1/G1_ids":              []byte(`[2,3]`),
	"C/2/c":                   []byte(`"c2"`),
	"C/2/title":               []byte(`"c2"`),
	"C/2/A_id":                []byte(`2`),
	"C/2/B_ids":               []byte(`[2]`),
	"C/2/G1_ids":              []byte(`[2,3]`),
	"D/1/d":                   []byte(`"d1"`),
	"D/1/B_$_ids":             []byte(`["1","2","3"]`),
	"D/1/B_1_ids":             []byte(`[1,2]`),
	"D/1/B_2_ids":             []byte(`[1]`),
	"D/1/B_3_ids":             []byte(`[]`),
	"D/2/d":                   []byte(`"d2"`),
	"D/2/B_$_ids":             []byte(`["1","4"]`),
	"D/2/B_1_ids":             []byte(`[]`),
	"D/2/B_4_ids":             []byte(`[2]`),
	"G1/1/g1":                 []byte(`"g1.1"`),
	"G1/1/content_object_ids": []byte(`["A/1"]`),
	"G1/2/g1":                 []byte(`"g1.2"`),
	"G1/2/content_object_ids": []byte(`["A/1","C/1","C/2"]`),
	"G1/3/g1":                 []byte(`"g1.3"`),
	"G1/3/content_object_ids": []byte(`["C/1","C/2"]`),
	"G2/1/g2":                 []byte(`"g2.1"`),
	"G2/1/content_object_id":  []byte(`"B/1"`),
}

func TestFeatures(t *testing.T) {
	//TODO
	t.Skip("Talk with Finn, if its ok to return a dict again.")

	datastore := test.NewMockDatastore()
	datastore.Data = dataSet
	datastore.OnlyData = true
	s := autoupdate.New(datastore, new(test.MockRestricter))
	defer s.Close()

	for _, tt := range []struct {
		name    string
		request string
		result  string
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
				"A": {
					"1": {
						"a": "a1",
						"C_ids": [],
						"B_id": 1,
						"G1_ids": [1,2]
					},
					"2": {
						"a": "a2",
						"C_ids": [1,2],
						"G1_ids": []
					}
				},
				"C": {
					"1": {
						"c": "c1",
						"G1_ids": [2,3]
					},
					"2": {
						"c": "c2",
						"G1_ids": [2,3]
					}
				},
				"G1": {
					"1": {
						"g1": "g1.1"
					},
					"2": {
						"g1": "g1.2"
					},
					"3": {
						"g1": "g1.3"
					}
				}
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
				"B": {
					"1": {
						"B_children_ids": [2],
						"C_ids": [1],
						"G2_id": 1
					},
					"2" : {
						"C_ids": [1,2],
						"B_parent_id": 1
					}
				},
				"G2": {
					"1": {
						"content_object_id": "B/1"
					}
				},
				"C": {
					"1": {
						"c": "c1",
						"title": "c1"
					},
					"2": {
						"c": "c2"
					}
				}
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
				"G1": {
					"2": {
						"content_object_ids": ["A/1","C/1","C/2"]
					}
				},
				"A": {
					"1": {
						"a": "a1",
						"title": "a1",
						"G1_ids": [1,2]
					}
				},
				"C": {
					"1": {
						"title": "c1",
						"A_id": 2,
						"G1_ids": [2,3]
					},
					"2": {
						"title": "c2",
						"A_id": 2,
						"G1_ids": [2,3]
					}
				}
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
				"D": {
					"1": {
						"d": "d1",
						"B_$_ids": ["1","2","3"]
					},
					"2": {
						"d": "d2",
						"B_$_ids": ["1","4"]
					}
				}
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
				"D": {
					"1":{
						"d": "d1",
						"B_$_ids": ["1","2","3"],
						"B_1_ids": [1,2],
						"B_2_ids": [1],
						"B_3_ids": []
					},
					"2": {
						"d":       "d2",
						"B_$_ids": ["1","4"],
						"B_1_ids": [],
						"B_4_ids": [2]
					}
				}
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
					"B_4_ids": {
						"type": "relation-list",
						"collection": "B",
						"fields": {
							"title": null
						}
					}
				}
			}`,
			`{
				"D": {
					"1": {
						"B_$_ids": ["1","2","3"],
						"B_1_ids": [1,2],
						"B_2_ids": [1],
						"B_3_ids": []
					},
					"2": {
						"B_$_ids": ["1","4"],
						"B_1_ids": [],
						"B_4_ids": [2]
					}
				},
				"B": {
					"1": {
						"b": "b1"
					},
					"2": {
						"b": "b2",
						"title": "b2"
					}
				}
			}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			b, err := keysbuilder.FromJSON(context.Background(), strings.NewReader(tt.request), s.IDer(1))
			if err != nil {
				t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
			}
			c := s.Connect(context.Background(), 1, b)
			data, err := c.Next()
			if err != nil {
				t.Fatalf("Can not get data: %v", err)
			}

			var expect map[string]map[string]map[string]json.RawMessage
			if err := json.Unmarshal([]byte(tt.result), &expect); err != nil {
				t.Fatalf("Can not decode keys from expected data: %v", err)
			}

			cmpMap(t, data, expect)
		})
	}
}

func cmpMap(t *testing.T, got map[string]json.RawMessage, expect map[string]map[string]map[string]json.RawMessage) {
	v1, _ := json.Marshal(got)
	v2, _ := json.Marshal(expect)
	if string(v1) != string(v2) {
		t.Errorf("got %s, expected %s", string(v1), string(v2))
	}
}
