package keysbuilder_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func TestKeys(t *testing.T) {
	for _, tt := range []struct {
		name    string
		request string
		data    string
		keys    []dskey.Key
	}{
		{
			"One Field",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {"name": null}
			}`,
			"",
			mustKeys("user/1/name"),
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
			"",
			mustKeys("user/1/first", "user/1/last"),
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
			"",
			mustKeys("user/1/first", "user/1/last", "user/2/first", "user/2/last"),
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
			"user/1/note_id: 1",
			mustKeys("user/1/note_id", "note/1/important"),
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
			"user/1/group_ids: [1,2]",
			mustKeys("user/1/group_ids", "group/1/admin", "group/2/admin"),
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
			`---
			user/1/note_id: 1
			note/1/motion_id: 1
			`,
			mustKeys("user/1/note_id", "note/1/motion_id", "motion/1/name"),
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
			`---
			user/1/group_ids: [1,2]
			group/1/perm_ids: [1,2]
			group/2/perm_ids: [1,2]
			`,
			mustKeys("user/1/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name"),
		},
		{
			"Request _id without redirect",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {"note_id": null}
			}`,
			"",
			mustKeys("user/1/note_id"),
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
			"",
			mustKeys("not_exist/1/note_id"),
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
			"",
			mustKeys("not_exist/1/group_ids"),
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
			`---
			user/1:
				group_$_ids:  ["1","2"]
				group_$1_ids: [1,2]
				group_$2_ids: [1,2]
			`,
			mustKeys("user/1/group_$_ids", "user/1/group_$1_ids", "user/1/group_$2_ids", "group/1/name", "group/2/name"),
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
			"user/1/likes: other/1",
			mustKeys("user/1/likes", "other/1/name"),
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
			`---
			user/1/likes:    other/1
			other/1/tag_ids: [1,2]
			`,
			mustKeys("user/1/likes", "other/1/tag_ids", "tag/1/name", "tag/2/name"),
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
			`user/1/likes: ["other/1","other/2"]`,
			mustKeys("user/1/likes", "other/1/name", "other/2/name"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(dsmock.YAMLData(tt.data))
			b, err := keysbuilder.FromJSON(strings.NewReader(tt.request))
			if err != nil {
				t.Fatalf("FromJSON returned the unexpected error: %v", err)
			}

			keys, err := b.Update(context.Background(), ds)
			if err != nil {
				t.Fatalf("Building keys: %v", err)
			}

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
		data    string
		newData string
		got     []dskey.Key
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
			"user/1/note_id: 1",
			"user/1/note_id: 2",
			mustKeys("user/1/note_id", "note/2/important"),
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
			"user/1/note_id: 1",
			"user/1/note_id: 1",
			mustKeys("user/1/note_id", "note/1/important"),
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
			`---
			user/1/note_id: 1
			user/2/note_id: 1
			`,
			`---
			user/1/note_id: 2
			user/2/note_id: 1
			`,
			mustKeys("user/1/note_id", "user/2/note_id", "note/1/important", "note/2/important"),
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
			`---
			user/1/note_id:   1
			user/1/group_ids: [1,2]
			user/2/group_ids: [1,2]
			`,
			`---
			user/1/note_id:   2
			user/1/group_ids: [1,2]
			user/2/group_ids: [1,2]
			`,
			mustKeys("user/1/note_id", "user/1/group_ids", "note/2/important", "group/1/admin", "group/2/admin"),
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
			`---
			user/1/note_id: 1
			user/1/group_ids: [1,2]
			`,
			`---
			user/1/note_id: 2
			user/1/group_ids: [2]
			`,
			mustKeys("user/1/note_id", "note/2/important", "user/1/group_ids", "group/2/admin"),
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
			`---
			user/1/group_ids: [1,2]
			group/1/perm_ids: [1,2]
			group/2/perm_ids: [1,2]
			`,
			`---
			user/1/group_ids: [2]
			group/1/perm_ids: [1,2]
			group/2/perm_ids: [1,2]
			`,
			mustKeys("user/1/group_ids", "group/2/perm_ids", "perm/2/name", "perm/1/name"),
			1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(dsmock.YAMLData(tt.data))
			b, err := keysbuilder.FromJSON(strings.NewReader(tt.request))
			if err != nil {
				t.Fatalf("FromJSON() returned an unexpected error: %v", err)
			}
			if _, err := b.Update(context.Background(), ds); err != nil {
				t.Errorf("Update() returned an unexpect error: %v", err)
			}

			ds = dsmock.Stub(dsmock.YAMLData(tt.newData))

			keys, err := b.Update(context.Background(), ds)
			if err != nil {
				t.Errorf("Update() returned an unexpect error: %v", err)
			}

			if diff := cmpSet(set(tt.got...), set(keys...)); diff != nil {
				t.Errorf("Update() returned %v, expected %v", diff, keys)
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

	ds, _ := dsmock.NewMockDatastore(dsmock.YAMLData(`---
	user/1/group_ids: [1,2]
	user/2/group_ids: [1,2]
	user/3/group_ids: [1,2]
	group/1/perm_ids: [1,2]
	group/2/perm_ids: [1,2]
	`))

	b, err := keysbuilder.FromJSON(strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
	}

	keys, err := b.Update(context.Background(), ds)
	if err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	if got := len(ds.Requests()); got != 2 {
		t.Errorf("Got %d requests to the datastore, expected 2: %v", got, ds.Requests())
	}

	expect := mustKeys("user/1/group_ids", "user/2/group_ids", "user/3/group_ids", "group/1/perm_ids", "group/2/perm_ids", "perm/1/name", "perm/2/name")
	if diff := cmpSet(set(expect...), set(keys...)); diff != nil {
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

	ds, _ := dsmock.NewMockDatastore(dsmock.YAMLData(`---
	user/1/note_id: 1
	user/2/note_id: 1
	`))

	b, err := keysbuilder.ManyFromJSON(strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("FromJSON() returned an unexpected error: %v", err)
	}

	keys, err := b.Update(context.Background(), ds)
	if err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	if got := len(ds.Requests()); got != 1 {
		t.Errorf("Got %d requests, expected 1: %v", got, ds.Requests())
	}

	expect := mustKeys("user/1/note_id", "user/2/note_id", "motion/1/name", "note/1/important")
	if diff := cmpSet(set(expect...), set(keys...)); diff != nil {
		t.Errorf("Got %v, expected %v", diff, expect)
	}
}

func TestManyRequestsSameKeys(t *testing.T) {
	ctx := context.Background()

	jsonData := `
	[
		{
			"ids": [1],
			"collection": "user",
			"fields": {
				"note_id": {
					"type": "relation",
					"collection": "note",
					"fields": {
						"important": null
					}
				}
			}
		}, {
			"ids": [1],
			"collection": "user",
			"fields": {
				"note_id": null
			}
		}
	]`

	ds := dsmock.Stub(dsmock.YAMLData(`---
	user/1/note_id: 1
	`))

	b, err := keysbuilder.ManyFromJSON(strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("FromJSON(): %v", err)
	}

	keys, err := b.Update(ctx, ds)
	if err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	expect := mustKeys(
		"user/1/note_id",
		"note/1/important",
	)
	if diff := cmpSet(set(expect...), set(keys...)); diff != nil {
		t.Errorf("Got %v, expected %v", diff, expect)
	}
}

func TestError(t *testing.T) {
	ds, _ := dsmock.NewMockDatastore(nil)
	ds.InjectError(errors.New("Some Error"))
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

	b, err := keysbuilder.FromJSON(strings.NewReader(json))
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if _, err := b.Update(context.Background(), ds); err == nil {
		t.Fatalf("Expected Update() to return an error, got none")
	}

	if got := len(ds.Requests()); got != 0 {
		t.Errorf("Got %d requests, expected 0: %v", got, ds.Requests())
	}
}

func TestRequestCount(t *testing.T) {
	ds, _ := dsmock.NewMockDatastore(nil)
	json := `{
		"ids": [1],
		"collection": "user",
		"fields": {
			"name": null,
			"goodlooking": null,
			"note_ids": {
				"type": "relation-list",
				"collection": "note",
				"fields": {
					"important": null,
					"text": null
				}
			},
			"main_group": {
				"type": "relation",
				"collection": "note",
				"fields": {
					"name": null,
					"permissions": null
				}
			}
		}
	}`
	b, err := keysbuilder.FromJSON(strings.NewReader(json))
	if err != nil {
		t.Fatalf("FromJSON returned unexpected error: %v", err)
	}
	if _, err := b.Update(context.Background(), ds); err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	if got := len(ds.Requests()); got != 1 {
		t.Errorf("Updated() did %d requests, expected 1", got)
	}
}
