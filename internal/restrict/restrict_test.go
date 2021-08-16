package restrict_test

import (
	"context"
	"testing"

	restrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func TestRestrict(t *testing.T) {
	fetch := datastore.NewFetcher(dsmock.Stub(dsmock.YAMLData(`---
	meeting/1/enable_anonymous: true
	meeting/2/enable_anonymous: false

	user/1:
		group_$_ids: ["1"]
		group_$1_ids: [10]
	group:
		1:
			meeting_id: 1
		2:
			meeting_id: 2
		10:
			meeting_id: 1
			permissions:
			- agenda_item.can_manage
	agenda_item:
		1:
			meeting_id: 1
		10:
			meeting_id: 2
	tag:
		1:
			meeting_id: 1
		2:
			meeting_id: 2
	`)))

	data := map[string]string{
		"agenda_item/1/item_number":   `"numberOne"`,
		"agenda_item/1/unknown_field": `"numberOne"`,
		"agenda_item/1/tag_ids":       `[1,2]`,
		"agenda_item/404/item_number": `"numberA"`,
		"agenda_item/10/item_number":  `"numberB"`,
		"unknown_collection/1/field":  "404",
		"tag/1/tagged_ids":            `["agenda_item/1","agenda_item/10"]`,
		"user/1/group_$_ids":          `["1","2"]`,
		"user/1/group_$1_ids":         `[1]`,
		"user/1/group_$2_ids":         `[2]`,
	}

	got := make(map[string][]byte, len(data))
	for k, v := range data {
		got[k] = []byte(v)
	}

	err := restrict.Restrict(context.Background(), fetch, 1, got)

	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if got["agenda_item/1/item_number"] == nil {
		t.Errorf("agenda_item/1/item_number was removed")
	}

	if got["agenda_item/1/unknown_field"] != nil {
		t.Errorf("agenda_item/1/item_number was not removed")
	}

	if got["agenda_item/404/item_number"] != nil {
		t.Errorf("agenda_item/404/item_number was not removed")
	}

	if got["agenda_item/10/item_number"] != nil {
		t.Errorf("agenda_item/404/item_number was not removed")
	}

	if got["unknown_collection/1/field"] != nil {
		t.Errorf("unknown_collection/1/field was not removed")
	}

	if got := string(got["tag/1/tagged_ids"]); got != `["agenda_item/1"]` {
		t.Errorf("tag/1/tagged_ids was restricted to %q, expedted %q", got, `["agenda_item/1"]`)
	}

	if got := string(got["agenda_item/1/tag_ids"]); got != `[1]` {
		t.Errorf("agenda_item/1/tag_ids was restricted to %q, expedted %q", got, `[1]`)
	}

	// This should change in the future. meeting 2 is not visible
	if got := string(got["user/1/group_$_ids"]); got != `["1","2"]` {
		t.Errorf("user/1/group_$_ids was restricted to %q, did not expect it", got)
	}

	if got := string(got["user/1/group_$1_ids"]); got != `[1]` {
		t.Errorf("user/1/group_$1_ids was restricted to %q, did not expect it", got)
	}

	if got := string(got["user/1/group_$2_ids"]); got != `null` {
		t.Errorf("user/1/group_$2_ids is %q, expected a empty list", got)
	}
}

func TestRestrictSuperAdmin(t *testing.T) {
	fetch := datastore.NewFetcher(dsmock.Stub(dsmock.YAMLData(`---
	user/1/organization_management_level: superadmin
	personal_note/1/user_id: 1
	personal_note/2/user_id: 2
	`)))

	data := map[string]string{
		"unknown_collection/404/field": "404",
		"user/404/unknown_field":       "404",
		"personal_note/1/id":           "1",
		"personal_note/2/id":           "2",
	}

	got := make(map[string][]byte, len(data))
	for k, v := range data {
		got[k] = []byte(v)
	}

	err := restrict.Restrict(context.Background(), fetch, 1, got)

	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if got["unknown_collection/404/field"] == nil {
		t.Errorf("unknown_collection/404/field was restricted")
	}

	if got["user/404/unknown_field"] == nil {
		t.Errorf("user/404/unknown_field was restricted")
	}

	if got["personal_note/1/id"] == nil {
		t.Errorf("personal_note/1/id got restricted")
	}

	if got["personal_note/2/id"] != nil {
		t.Errorf("personal_note/2/id got not restricted")
	}
}
