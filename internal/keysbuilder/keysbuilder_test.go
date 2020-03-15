package keysbuilder_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/openslides/openslides-autoupdate-service/internal/keysbuilder"
)

func TestKeys(t *testing.T) {
	for _, tt := range []struct {
		name     string
		request  string
		keys     []string
		reqCount int
	}{
		{
			"One Field",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {"name": null}
			}`,
			strs("user/1/name"),
			0,
		},
		{
			"Many Fields",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"first": null,
					"last": null
				}
			}`,
			strs("user/1/first", "user/1/last"),
			0,
		},
		{
			"Many IDs Many Fields",
			`{
				"ids": [1, 2],
				"collection": "user",
				"fields": {
					"first": null,
					"last": null
				}
			}`,
			strs("user/1/first", "user/1/last", "user/2/first", "user/2/last"),
			0,
		},
		{
			"Redirect Once id",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {"important": null}
					}
				}
			}`,
			strs("user/1/note_id", "note/1/important"),
			1,
		},
		{
			"Redirect Once ids",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"group_ids": {
						"type": "relation-list",
						"collection": "group",
						"fields": {"admin": null}
					}
				}
			}`,
			strs("user/1/group_ids", "group/1/admin", "group/2/admin"),
			1,
		},
		{
			"Redirect twice id",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {
							"motion_id": {
								"type": "relation",
								"collection": "motion",
								"fields": {"name": null}
							}
						}
					}
				}
			}`,
			strs("user/1/note_id", "note/1/motion_id", "motion/1/name"),
			2,
		},
		{
			"Redirect twice ids",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"group_ids": {
						"type": "relation-list",
						"collection": "group",
						"fields": {
							"perm_ids": {
								"type": "relation-list",
								"collection": "perm",
								"fields": {"name": null}
							}
						}
					}
				}
			}`,
			strs("user/1/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name"),
			3,
		},
		{
			"Request _id without redirect",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {"note_id": null}
			}`,
			strs("user/1/note_id"),
			0,
		},
		{
			"Redirect id not exist",
			`{
				"ids": [1],
				"collection": "not_exist",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {"important": null}
					}
				}
			}`, // "not_exist" is a magic string in the ider mock
			strs("not_exist/1/note_id"),
			1,
		},
		{
			"Redirect ids not exist",
			`{
				"ids": [1],
				"collection": "not_exist",
				"fields": {
					"group_ids": {
						"type": "relation-list",
						"collection": "group",
						"fields": {"name": null}
					}
				}
			}`, // "not_exist" is a magic string in the ider mock
			strs("not_exist/1/group_ids"),
			1,
		},
		{
			"Template field",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"group_$_ids": {
						"type": "template",
						"sub": {
							"type": "relation-list",
							"collection": "group",
							"fields": {"name": null}
						}
					}
				}
			}`,
			strs("user/1/group_$_ids", "user/1/group_1_ids", "user/1/group_2_ids", "group/1/name", "group/2/name"),
			3,
		},
		{
			"Generic field",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"likes": {
						"type": "generic-relation",
						"fields": {"name": null}
					}
				}
			}`,
			strs("user/1/likes", "other/1/name"),
			1,
		},
		{
			"Generic field with sub fields",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"likes": {
						"type": "generic-relation",
						"fields": {
							"tag_ids": {
								"type": "relation-list",
								"collection": "tag",
								"fields": {"name": null}
							}
						}
					}
				}
			}`,
			strs("user/1/likes", "other/1/tag_ids", "tag/1/name", "tag/2/name"),
			2,
		},
		{
			"Generic list field",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"likes": {
						"type": "generic-relation-list",
						"fields": {"name": null}
					}
				}
			}`,
			strs("user/1/likes", "other/1/name", "other/2/name"),
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ider := &mockIDer{}
			b, err := keysbuilder.FromJSON(context.Background(), strings.NewReader(tt.request), ider)
			if err != nil {
				t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
			}

			keys := b.Keys()

			if diff := cmpSet(set(tt.keys...), set(keys...)); diff != nil {
				t.Errorf("Expected %v, got: %v", tt.keys, diff)
			}
			if len(ider.reqLog) != tt.reqCount {
				t.Errorf("Expected %d requests to the restricter, got: %d: %v", tt.reqCount, len(ider.reqLog), ider.reqLog)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	for _, tt := range []struct {
		name    string
		request string
		newDB   map[string][]int
		got     []string
		count   int
	}{
		{
			"One relation",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {"important": null}
					}
				}
			}`,
			map[string][]int{"user/1/note_id": ids(2)},
			strs("user/1/note_id", "note/2/important"),
			1,
		},
		{
			"One relation no change",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {"important": null}
					}
				}
			}`,
			map[string][]int{},
			strs("user/1/note_id", "note/1/important"),
			0,
		},
		{
			"Two ids one change",
			`{
				"ids": [1, 2],
				"collection": "user",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {"important": null}
					}
				}
			}`,
			map[string][]int{"user/1/note_id": ids(2)},
			strs("user/1/note_id", "user/2/note_id", "note/1/important", "note/2/important"),
			1,
		},
		{
			"Two relation one change",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {"important": null}
					},
					"group_ids": {
						"type": "relation-list",
						"collection": "group",
						"fields": {"admin": null}
					}
				}
			}`,
			map[string][]int{"user/1/note_id": ids(2)},
			strs("user/1/note_id", "user/1/group_ids", "note/2/important", "group/1/admin", "group/2/admin"),
			1,
		},
		{
			"Two relation two changes",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"note_id": {
						"type": "relation",
						"collection": "note",
						"fields": {"important": null}
					},
					"group_ids": {
						"type": "relation-list",
						"collection": "group",
						"fields": {"admin": null}
					}
				}
			}`,
			map[string][]int{"user/1/note_id": ids(2), "user/1/group_ids": ids(2)},
			strs("user/1/note_id", "note/2/important", "user/1/group_ids", "group/2/admin"),
			2,
		},
		{
			"Tree levels out changes",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"group_ids": {
						"type": "relation-list",
						"collection": "group",
						"fields": {
							"perm_ids": {
								"type": "relation-list",
								"collection": "perm",
								"fields": {"name": null}
							}
						}
					}
				}
			}`,
			map[string][]int{"user/1/group_ids": ids(2)},
			strs("user/1/group_ids", "group/2/perm_ids", "perm/2/name", "perm/1/name"),
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ider := &mockIDer{}
			b, err := keysbuilder.FromJSON(context.Background(), strings.NewReader(tt.request), ider)
			if err != nil {
				t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
			}
			ider.data = tt.newDB
			ider.reqLog = ider.reqLog[:0]

			if err := b.Update(mapKeys(tt.newDB)); err != nil {
				t.Errorf("Expect Update() not to return an error, got: %v", err)
			}

			if diff := cmpSet(set(tt.got...), set(b.Keys()...)); diff != nil {
				t.Errorf("Expected %v, got: %v", b.Keys(), diff)
			}
			if tt.count != len(ider.reqLog) {
				t.Errorf("Expected %d requests to the restricter, got: %d: %v", tt.count, len(ider.reqLog), ider.reqLog)
			}
		})
	}
}

func TestConcurency(t *testing.T) {
	json := `
	{
		"ids": [1, 2, 3],
		"collection": "user",
		"fields": {
			"group_ids": {
				"type": "relation-list",
				"collection": "group",
				"fields": {
					"perm_ids": {
						"type": "relation-list",
						"collection": "perm",
						"fields": {"name": null}
					}
				}
			}
		}

	}`
	ider := &mockIDer{sleep: 10 * time.Millisecond}
	start := time.Now()
	b, err := keysbuilder.FromJSON(context.Background(), strings.NewReader(json), ider)
	if err != nil {
		t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
	}
	finished := time.Since(start)

	if finished > 30*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 30 Milliseconds, got: %v", finished)
	}
	expect := strs("user/1/group_ids", "user/2/group_ids", "user/3/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name")
	if diff := cmpSet(set(expect...), set(b.Keys()...)); diff != nil {
		t.Errorf("Expected %v, got: %v", expect, diff)
	}
	if len(ider.reqLog) != 5 {
		t.Errorf("Expected %d requests to the restricter.IDsFromKey, got: %d: %v", 5, len(ider.reqLog), ider.reqLog)
	}
}

func TestManyRequests(t *testing.T) {
	json := `
	[
		{
			"ids": [1],
			"collection": "user",
			"fields": {
				"note_id": {
					"type": "relation",
					"collection": "note",
					"fields": {"important": null}
				}
			}
		}, {
			"ids": [1],
			"collection": "motion",
			"fields": {"name": null}
		}, {
			"ids": [2],
			"collection": "user",
			"fields": {
				"note_id": {
					"type": "relation",
					"collection": "note",
					"fields": {"important": null}
				}
			}
		}
	]`
	ider := &mockIDer{sleep: 10 * time.Millisecond}
	start := time.Now()
	b, err := keysbuilder.ManyFromJSON(context.Background(), strings.NewReader(json), ider)
	if err != nil {
		t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
	}

	finished := time.Since(start)
	if finished > 20*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 20 Milliseconds, got: %v", finished)
	}

	expect := strs("user/1/note_id", "user/2/note_id", "motion/1/name", "note/1/important")
	if diff := cmpSet(set(expect...), set(b.Keys()...)); diff != nil {
		t.Errorf("Expected %v, got: %v", expect, diff)
	}
	if len(ider.reqLog) != 2 {
		t.Errorf("Expected %d requests to the restricter.IDsFromKey, got: %d: %v", 2, len(ider.reqLog), ider.reqLog)
	}
}

func TestError(t *testing.T) {
	json := `
	{
		"ids": [1],
		"collection": "user",
		"fields": {
			"note_id": {
				"type": "relation",
				"collection": "note",
				"fields": {"important": null}
			}
		}
	}`
	ider := &mockIDer{err: errors.New("Some Error"), sleep: 10 * time.Millisecond}

	start := time.Now()
	_, err := keysbuilder.FromJSON(context.Background(), strings.NewReader(json), ider)
	if err == nil {
		t.Fatalf("Expected FromJSON() to return an error, got none")
	}
	finished := time.Since(start)

	if finished > 20*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 20 Milliseconds, got: %v", finished)
	}
}
