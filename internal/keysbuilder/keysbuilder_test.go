package keysbuilder_test

import (
	"context"
	"fmt"
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
				"fields": {"username": null}
			}`,
			"",
			mustKeys("user/1/username"),
		},
		{
			"Many Fields",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"first_name": null,
					"last_name": null
				}
			}`,
			"",
			mustKeys("user/1/first_name", "user/1/last_name"),
		},
		{
			"Many IDs Many Fields",
			`{
				"ids": [1, 2],
				"collection": "user",
				"fields": {
					"first_name": null,
					"last_name": null
				}
			}`,
			"",
			mustKeys("user/1/first_name", "user/1/last_name", "user/2/first_name", "user/2/last_name"),
		},
		{
			"Redirect Once id",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"organization_id": {
						"type": "relation",
						"collection": "organization",
						"fields": {"name": null}
					}
				}
			}`,
			"user/1/organization_id: 1",
			mustKeys("user/1/organization_id", "organization/1/name"),
		},
		{
			"Redirect Once ids",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"meeting_user_ids": {
						"type": "relation-list",
						"collection": "meeting_user",
						"fields": {"comment": null}
					}
				}
			}`,
			"user/1/meeting_user_ids: [1,2]",
			mustKeys("user/1/meeting_user_ids", "meeting_user/1/comment", "meeting_user/2/comment"),
		},
		{
			"Redirect twice id",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"organization_id": {
						"type": "relation",
						"collection": "organization",
						"fields": {
							"theme_id": {
								"type": "relation",
								"collection": "theme",
								"fields": {"name": null}
							}
						}
					}
				}
			}`,
			`---
			user/1/organization_id: 1
			organization/1/theme_id: 1
			`,
			mustKeys("user/1/organization_id", "organization/1/theme_id", "theme/1/name"),
		},
		{
			"Redirect twice ids",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"meeting_user_ids": {
						"type": "relation-list",
						"collection": "meeting_user",
						"fields": {
							"group_ids": {
								"type": "relation-list",
								"collection": "group",
								"fields": {"name": null}
							}
						}
					}
				}
			}`,
			`---
			user/1/meeting_user_ids: [1,2]
			meeting_user/1/group_ids: [1,2]
			meeting_user/2/group_ids: [1,2]
			`,
			mustKeys("user/1/meeting_user_ids", "meeting_user/1/group_ids", "meeting_user/2/group_ids", "group/1/name", "group/2/name"),
		},
		{
			"Request _id without redirect",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {"organization_id": null}
			}`,
			"",
			mustKeys("user/1/organization_id"),
		},
		{
			"Generic field",
			`{
				"ids": [1],
				"collection": "personal_note",
				"fields": {
					"content_object_id": {
						"type": "generic-relation",
						"fields": {"title": null}
					}
				}
			}`,
			"personal_note/1/content_object_id: motion/1",
			mustKeys("personal_note/1/content_object_id", "motion/1/title"),
		},
		{
			"Generic field with sub fields",
			`{
				"ids": [1],
				"collection": "personal_note",
				"fields": {
					"content_object_id": {
						"type": "generic-relation",
						"fields": {
							"amendment_ids": {
								"type": "relation-list",
								"collection": "motion",
								"fields": {"title": null}
							}
						}
					}
				}
			}`,
			`---
			personal_note/1/content_object_id:    motion/1
			motion/1/amendment_ids: [1,2]
			`,
			mustKeys("personal_note/1/content_object_id", "motion/1/amendment_ids", "motion/1/title", "motion/2/title"),
		},
		{
			"Generic list field",
			`{
				"ids": [1],
				"collection": "organization_tag",
				"fields": {
					"tagged_ids": {
						"type": "generic-relation-list",
						"fields": {"name": null}
					}
				}
			}`,
			`organization_tag/1/tagged_ids: ["meeting/1","meeting/2"]`,
			mustKeys("organization_tag/1/tagged_ids", "meeting/1/name", "meeting/2/name"),
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
					"organization_id": {
						"type": "relation",
						"collection": "organization",
						"fields": {"login_text": null}
					}
				}
			}`,
			"user/1/organization_id: 1",
			"user/1/organization_id: 2",
			mustKeys("user/1/organization_id", "organization/2/login_text"),
			1,
		},
		{
			"One relation no change",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"organization_id": {
						"type": "relation",
						"collection": "organization",
						"fields": {"login_text": null}
					}
				}
			}`,
			"user/1/organization_id: 1",
			"user/1/organization_id: 1",
			mustKeys("user/1/organization_id", "organization/1/login_text"),
			0,
		},
		{
			"Two ids one change",
			`{
				"ids": [1, 2],
				"collection": "user",
				"fields": {
					"organization_id": {
						"type": "relation",
						"collection": "organization",
						"fields": {"login_text": null}
					}
				}
			}`,
			`---
			user/1/organization_id: 1
			user/2/organization_id: 1
			`,
			`---
			user/1/organization_id: 2
			user/2/organization_id: 1
			`,
			mustKeys("user/1/organization_id", "user/2/organization_id", "organization/1/login_text", "organization/2/login_text"),
			1,
		},
		{
			"Two relation one change",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"organization_id": {
						"type": "relation",
						"collection": "organization",
						"fields": {"login_text": null}
					},
					"meeting_user_ids": {
						"type": "relation-list",
						"collection": "meeting_user",
						"fields": {"comment": null}
					}
				}
			}`,
			`---
			user/1/organization_id:   1
			user/1/meeting_user_ids: [1,2]
			user/2/meeting_user_ids: [1,2]
			`,
			`---
			user/1/organization_id:   2
			user/1/meeting_user_ids: [1,2]
			user/2/meeting_user_ids: [1,2]
			`,
			mustKeys("user/1/organization_id", "user/1/meeting_user_ids", "organization/2/login_text", "meeting_user/1/comment", "meeting_user/2/comment"),
			1,
		},
		{
			"Two relation two changes",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"organization_id": {
						"type": "relation",
						"collection": "organization",
						"fields": {"login_text": null}
					},
					"meeting_user_ids": {
						"type": "relation-list",
						"collection": "meeting_user",
						"fields": {"comment": null}
					}
				}
			}`,
			`---
			user/1/organization_id: 1
			user/1/meeting_user_ids: [1,2]
			`,
			`---
			user/1/organization_id: 2
			user/1/meeting_user_ids: [2]
			`,
			mustKeys("user/1/organization_id", "organization/2/login_text", "user/1/meeting_user_ids", "meeting_user/2/comment"),
			2,
		},
		{
			"Tree levels out changes",
			`{
				"ids": [1],
				"collection": "user",
				"fields": {
					"meeting_user_ids": {
						"type": "relation-list",
						"collection": "meeting_user",
						"fields": {
							"personal_note_ids": {
								"type": "relation-list",
								"collection": "personal_note",
								"fields": {"note": null}
							}
						}
					}
				}
			}`,
			`---
			user/1/meeting_user_ids: [1,2]
			meeting_user/1/personal_note_ids: [1,2]
			meeting_user/2/personal_note_ids: [1,2]
			`,
			`---
			user/1/meeting_user_ids: [2]
			meeting_user/1/personal_note_ids: [1,2]
			meeting_user/2/personal_note_ids: [1,2]
			`,
			mustKeys("user/1/meeting_user_ids", "meeting_user/2/personal_note_ids", "personal_note/2/note", "personal_note/1/note"),
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
	ctx := context.Background()
	jsonData := `
	{
		"ids": [1, 2, 3],
		"collection": "user",
		"fields": {
			"meeting_user_ids": {
				"type": "relation-list",
				"collection": "meeting_user",
				"fields": {
					"personal_note_ids": {
						"type": "relation-list",
						"collection": "personal_note",
						"fields": {"note": null}
					}
				}
			}
		}

	}`

	ds := dsmock.NewFlow(
		dsmock.YAMLData(`---
		user/1/meeting_user_ids: [1,2]
		user/2/meeting_user_ids: [1,2]
		user/3/meeting_user_ids: [1,2]
		meeting_user/1/personal_note_ids: [1,2]
		meeting_user/2/personal_note_ids: [1,2]
		`),
		dsmock.NewCounter,
	)
	counter := ds.Middlewares()[0].(*dsmock.Counter)

	b, err := keysbuilder.FromJSON(strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("Expected FromJSON() not to return an error, got: %v", err)
	}

	keys, err := b.Update(ctx, ds)
	if err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	if got := counter.Count(); got != 2 {
		t.Errorf("Got %d requests to the datastore, expected 2: %v", got, counter.Requests())
	}

	expect := mustKeys(
		"user/1/meeting_user_ids",
		"user/2/meeting_user_ids",
		"user/3/meeting_user_ids",
		"meeting_user/1/personal_note_ids",
		"meeting_user/2/personal_note_ids",
		"personal_note/1/note",
		"personal_note/2/note",
	)
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
				"organization_id": {
					"type": "relation",
					"collection": "organization",
					"fields": {"login_text": null}
				}
			}
		}, {
			"ids": [1],
			"collection": "motion",
			"fields": {"title": null}
		}, {
			"ids": [2],
			"collection": "user",
			"fields": {
				"organization_id": {
					"type": "relation",
					"collection": "organization",
					"fields": {"login_text": null}
				}
			}
		}
	]`

	ds := dsmock.NewFlow(
		dsmock.YAMLData(`---
		user/1/organization_id: 1
		user/2/organization_id: 1
		`),
		dsmock.NewCounter,
	)
	counter := ds.Middlewares()[0].(*dsmock.Counter)

	b, err := keysbuilder.ManyFromJSON(strings.NewReader(jsonData))
	if err != nil {
		t.Fatalf("FromJSON() returned an unexpected error: %v", err)
	}

	keys, err := b.Update(context.Background(), ds)
	if err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	if got := counter.Count(); got != 1 {
		t.Errorf("Got %d requests, expected 1: %v", got, counter.Requests())
	}

	expect := mustKeys("user/1/organization_id", "user/2/organization_id", "motion/1/title", "organization/1/login_text")
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
				"organization_id": {
					"type": "relation",
					"collection": "organization",
					"fields": {
						"login_text": null
					}
				}
			}
		}, {
			"ids": [1],
			"collection": "user",
			"fields": {
				"organization_id": null
			}
		}
	]`

	ds := dsmock.Stub(dsmock.YAMLData(`---
	user/1/organization_id: 1
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
		"user/1/organization_id",
		"organization/1/login_text",
	)
	if diff := cmpSet(set(expect...), set(keys...)); diff != nil {
		t.Errorf("Got %v, expected %v", diff, expect)
	}
}

func TestError(t *testing.T) {
	ctx := context.Background()
	waiter := make(chan error, 1)
	ds := dsmock.NewFlow(
		nil,
		dsmock.NewWait(waiter),
		dsmock.NewCounter,
	)
	counter := ds.Middlewares()[1].(*dsmock.Counter)
	waiter <- fmt.Errorf("some error")

	json := `
	{
		"ids": [1],
		"collection": "user",
		"fields": {
			"organization_id": {
				"type": "relation",
				"collection": "organization",
				"fields": {"login_text": null}
			}
		}
	}`

	b, err := keysbuilder.FromJSON(strings.NewReader(json))
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if _, err := b.Update(ctx, ds); err == nil {
		t.Fatalf("Expected Update() to return an error, got none")
	}

	if got := counter.Count(); got != 1 {
		t.Errorf("Got %d requests, expected 1: %v", got, counter.Requests())
	}
}

func TestRequestCount(t *testing.T) {
	ctx := context.Background()
	ds := dsmock.NewFlow(
		nil,
		dsmock.NewCounter,
	)
	counter := ds.Middlewares()[0].(*dsmock.Counter)
	json := `{
		"ids": [1],
		"collection": "user",
		"fields": {
			"name": null,
			"goodlooking": null,
			"organization_ids": {
				"type": "relation-list",
				"collection": "organization",
				"fields": {
					"login_text": null,
					"text": null
				}
			},
			"main_group": {
				"type": "relation",
				"collection": "organization",
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
	if _, err := b.Update(ctx, ds); err != nil {
		t.Fatalf("Building keys: %v", err)
	}

	if got := counter.Count(); got != 1 {
		t.Errorf("Updated() did %d requests, expected 1", got)
	}
}
