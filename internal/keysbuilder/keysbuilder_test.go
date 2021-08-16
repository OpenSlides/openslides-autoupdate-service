package keysbuilder_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
)

func TestKeys(t *testing.T) {
	for _, tt := range []struct {
		name    string
		request string
		data    map[string][]byte
		keys    []string
	}{
		{
			"One Field",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {"name": null}
			}`,
			nil,
			strs("user/1/name"),
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
			nil,
			strs("user/1/first", "user/1/last"),
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
			nil,
			strs("user/1/first", "user/1/last", "user/2/first", "user/2/last"),
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
			map[string][]byte{"user/1/note_id": []byte("1")},
			strs("user/1/note_id", "note/1/important"),
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
			map[string][]byte{"user/1/group_ids": []byte("[1,2]")},
			strs("user/1/group_ids", "group/1/admin", "group/2/admin"),
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
			map[string][]byte{
				"user/1/note_id":   []byte("1"),
				"note/1/motion_id": []byte("1"),
			},
			strs("user/1/note_id", "note/1/motion_id", "motion/1/name"),
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
			map[string][]byte{
				"user/1/group_ids": []byte("[1,2]"),
				"group/1/perm_ids": []byte("[1,2]"),
				"group/2/perm_ids": []byte("[1,2]"),
			},
			strs("user/1/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name"),
		},
		{
			"Request _id without redirect",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {"note_id": null}
			}`,
			nil,
			strs("user/1/note_id"),
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
			}`,
			nil,
			strs("not_exist/1/note_id"),
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
			}`,
			nil,
			strs("not_exist/1/group_ids"),
		},
		{
			"Template field",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"group_$_ids": {
						"type": "template",
						"values": {
							"type": "relation-list",
							"collection": "group",
							"fields": {"name": null}
						}
					}
				}
			}`,
			map[string][]byte{
				"user/1/group_$_ids":  []byte(`["1","2"]`),
				"user/1/group_$1_ids": []byte("[1,2]"),
				"user/1/group_$2_ids": []byte("[1,2]"),
			},
			strs("user/1/group_$_ids", "user/1/group_$1_ids", "user/1/group_$2_ids", "group/1/name", "group/2/name"),
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
			map[string][]byte{
				"user/1/likes": []byte(`"other/1"`),
			},
			strs("user/1/likes", "other/1/name"),
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
			map[string][]byte{
				"user/1/likes":    []byte(`"other/1"`),
				"other/1/tag_ids": []byte("[1,2]"),
			},
			strs("user/1/likes", "other/1/tag_ids", "tag/1/name", "tag/2/name"),
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
			map[string][]byte{
				"user/1/likes": []byte(`["other/1","other/2"]`),
			},
			strs("user/1/likes", "other/1/name", "other/2/name"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dataProvider := &test.DataProvider{Data: tt.data}
			b, err := keysbuilder.FromJSON(strings.NewReader(tt.request), dataProvider, 1)
			if err != nil {
				t.Fatalf("FromJSON returned the unexpected error: %v", err)
			}

			if err := b.Update(context.Background()); err != nil {
				t.Fatalf("Building keys: %v", err)
			}

			keys := b.Keys()

			if diff := cmpSet(set(tt.keys...), set(keys...)); diff != nil {
				t.Errorf("Got keys %v, expected %v", diff, tt.keys)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	for _, tt := range []struct {
		name    string
		request string
		data    map[string][]byte
		newData map[string][]byte
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
			map[string][]byte{"user/1/note_id": []byte("1")},
			map[string][]byte{"user/1/note_id": []byte("2")},
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
			map[string][]byte{"user/1/note_id": []byte("1")},
			map[string][]byte{"user/1/note_id": []byte("1")},
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
			map[string][]byte{"user/1/note_id": []byte("1"), "user/2/note_id": []byte("1")},
			map[string][]byte{"user/1/note_id": []byte("2"), "user/2/note_id": []byte("1")},
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
			map[string][]byte{
				"user/1/note_id":   []byte("1"),
				"user/1/group_ids": []byte("[1,2]"),
				"user/2/group_ids": []byte("[1,2]"),
			},
			map[string][]byte{
				"user/1/note_id":   []byte("2"),
				"user/1/group_ids": []byte("[1,2]"),
				"user/2/group_ids": []byte("[1,2]"),
			},
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
			map[string][]byte{"user/1/note_id": []byte("1"), "group_ids": []byte("[1,2]")},
			map[string][]byte{"user/1/note_id": []byte("2"), "user/1/group_ids": []byte("[2]")},
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
			map[string][]byte{
				"user/1/group_ids": []byte("[1,2]"),
				"group/1/perm_ids": []byte("[1,2]"),
				"group/2/perm_ids": []byte("[1,2]"),
			},
			map[string][]byte{
				"user/1/group_ids": []byte("[2]"),
				"group/1/perm_ids": []byte("[1,2]"),
				"group/2/perm_ids": []byte("[1,2]"),
			},
			strs("user/1/group_ids", "group/2/perm_ids", "perm/2/name", "perm/1/name"),
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dataProvider := &test.DataProvider{Data: tt.data}
			b, err := keysbuilder.FromJSON(strings.NewReader(tt.request), dataProvider, 1)
			if err != nil {
				t.Fatalf("FromJSON() returned an unexpected error: %v", err)
			}
			dataProvider.Data = tt.newData

			if err := b.Update(context.Background()); err != nil {
				t.Errorf("Update() returned an unexpect error: %v", err)
			}

			if diff := cmpSet(set(tt.got...), set(b.Keys()...)); diff != nil {
				t.Errorf("Update() returned %v, expected %v", diff, b.Keys())
			}
		})
	}
}

func TestConcurency(t *testing.T) {
	jsonData := `
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
	data := map[string][]byte{
		"user/1/group_ids": []byte("[1,2]"),
		"user/2/group_ids": []byte("[1,2]"),
		"user/3/group_ids": []byte("[1,2]"),
		"group/1/perm_ids": []byte("[1,2]"),
		"group/2/perm_ids": []byte("[1,2]"),
	}
	dataProvider := &test.DataProvider{Data: data, Sleep: 10 * time.Millisecond}
	start := time.Now()
	b, err := keysbuilder.FromJSON(strings.NewReader(jsonData), dataProvider, 1)
	if err != nil {
		t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
	}
	if err := b.Update(context.Background()); err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	finished := time.Since(start)

	if finished > 30*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 30 Milliseconds, got: %v", finished)
	}
	expect := strs("user/1/group_ids", "user/2/group_ids", "user/3/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name")
	if diff := cmpSet(set(expect...), set(b.Keys()...)); diff != nil {
		t.Errorf("Expected %v, got: %v", expect, diff)
	}
}

func TestManyRequests(t *testing.T) {
	jsonData := `
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
	data := map[string][]byte{
		"user/1/note_id": []byte("1"),
		"user/2/note_id": []byte("1"),
	}
	dataProvider := &test.DataProvider{Data: data, Sleep: 10 * time.Millisecond}
	start := time.Now()
	b, err := keysbuilder.ManyFromJSON(strings.NewReader(jsonData), dataProvider, 1)
	if err != nil {
		t.Fatalf("FromJSON() returned an unexpected error: %v", err)
	}
	if err := b.Update(context.Background()); err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	finished := time.Since(start)
	if finished > 20*time.Millisecond {
		t.Errorf("ManyFromJON() took %v, expected less then 20 Milliseconds", finished)
	}

	expect := strs("user/1/note_id", "user/2/note_id", "motion/1/name", "note/1/important")
	if diff := cmpSet(set(expect...), set(b.Keys()...)); diff != nil {
		t.Errorf("Got %v, expected %v", diff, expect)
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
	dataProvider := &test.DataProvider{Err: errors.New("Some Error"), Sleep: 10 * time.Millisecond}

	start := time.Now()
	b, err := keysbuilder.FromJSON(strings.NewReader(json), dataProvider, 1)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)

	}
	if err := b.Update(context.Background()); err == nil {
		t.Fatalf("Expected Update() to return an error, got none")
	}
	finished := time.Since(start)

	if finished > 20*time.Millisecond {
		t.Errorf("Expect keysbuilder to run in less then 20 Milliseconds, got: %v", finished)
	}
}

func TestRequestCount(t *testing.T) {
	dataProvider := new(test.DataProvider)
	json := `{
		"ids": [1],
		"collection": "user",
		"fields": {
			"name": null,
			"goodLocking": null,
			"note_ids": {
				"type": "relation-list",
				"collection": "note",
				"fields": {
					"important": null,
					"text": null
				}
			},
			"main-group": {
				"type": "relation",
				"collection": "note",
				"fields": {
					"name": null,
					"permissions": null
				}
			}
		}
	}`
	b, err := keysbuilder.FromJSON(strings.NewReader(json), dataProvider, 1)
	if err != nil {
		t.Fatalf("FromJSON returned unexpected error: %v", err)
	}
	if err := b.Update(context.Background()); err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	if dataProvider.RequestCount != 1 {
		t.Errorf("Updated() did %d requests, expected 1", dataProvider.RequestCount)
	}
}

func TestLazy(t *testing.T) {
	dataProvider := new(test.DataProvider)
	dataProvider.Data = map[string][]byte{
		"user/1/note_id": []byte("1"),
	}

	jsonData := `{
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

	b, err := keysbuilder.FromJSON(strings.NewReader(jsonData), dataProvider, 1)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	// Change data after kb was created
	dataProvider.Data = map[string][]byte{
		"user/1/note_id": []byte("2"),
	}

	if err := b.Update(context.Background()); err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	expect := "note/2/important"
	if got := b.Keys()[1]; got != expect {
		t.Errorf("Got %s, expected %s", got, expect)
	}
}
