package keysbuilder_test

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func FuzzFromJSON(f *testing.F) {
	f.Add(`{
		"ids": [1],
		"collection": "user",
		"fields": {"name": null}
	}`)
	f.Add(`{
		"ids": [1],
		"collection": "user",
		"fields": {
			"first": null,
			"last": null
		}
	}`)
	f.Add(`{
		"ids": [1, 2],
		"collection": "user",
		"fields": {
			"first": null,
			"last": null
		}
	}`)
	f.Add(`{
		"ids": [1],
		"collection": "user",
		"fields": {
			"note_id": {
				"type": "relation",
				"collection": "note",
				"fields": {"important": null}
			}
		}
	}`)
	f.Add(`{
		"ids": [1],
		"collection": "user",
		"fields": {
			"group_ids": {
				"type": "relation-list",
				"collection": "group",
				"fields": {"admin": null}
			}
		}
	}`)
	f.Add(`{
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
	}`)
	f.Add(`{
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
	}`)
	f.Add(`{
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
	}`)
	f.Add(`{
		"ids": [1],
		"collection": "user",
		"fields": {
			"like": {
				"type": "generic-relation",
				"fields": {"name": null}
			}
		}
	}`)
	f.Add(`{
		"ids": [1],
		"collection": "user",
		"fields": {
			"like": {
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
	}`)
	f.Add(`{
		"ids": [1],
		"collection": "user",
		"fields": {
			"likes": {
				"type": "generic-relation-list",
				"fields": {"name": null}
			}
		}
	}`)

	ds := dsmock.Stub(map[dskey.Key][]byte{
		dskey.MustKey("user/1/note_id"):       []byte(`1`),
		dskey.MustKey("user/1/group_ids"):     []byte(`[1,2]`),
		dskey.MustKey("note/1/motion_id"):     []byte(`1`),
		dskey.MustKey("group/1/perm_ids"):     []byte(`[1,2]`),
		dskey.MustKey("group/2/perm_ids"):     []byte(`[1,2]`),
		dskey.MustKey("user/1/group_$_ids"):   []byte(`["1","2"]`),
		dskey.MustKey("user/1/group_$1_ids"):  []byte(`[1,2]`),
		dskey.MustKey("user/1/group_$_2_ids"): []byte(`[1,2]`),
		dskey.MustKey("user/1/like"):          []byte(`"topic/1"`),
		dskey.MustKey("user/1/likes"):         []byte(`["topic/1","agenda/1"]`),
		dskey.MustKey("topic/1/tag_ids"):      []byte(`[1,2]`),
		dskey.MustKey("agenda/1/tag_ids"):     []byte(`[1,2]`),
	})

	f.Fuzz(func(t *testing.T, query string) {
		if !json.Valid([]byte(query)) {
			t.Skip("invalid JSON")
		}

		kb, err := keysbuilder.FromJSON(strings.NewReader(query))
		if err != nil {
			var typedErr interface {
				Type() string
			}
			if errors.As(err, &typedErr) {
				t.Skip()
			}
			t.Fatalf("building keysbuilder:\n%s\n%v", query, err)
		}

		if _, err := kb.Update(context.Background(), ds); err != nil {
			var typedErr interface {
				Type() string
			}
			if errors.As(err, &typedErr) {
				t.Skip()
			}
			t.Fatalf("Updating keybuilder: %v", err)
		}
	})

}
