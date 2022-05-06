package restrict_test

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"

	restrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func TestRestrict(t *testing.T) {
	ds := dsmock.Stub(dsmock.YAMLData(`---
	meeting/30/enable_anonymous: true
	meeting/2:
		enable_anonymous: false
		committee_id: 404

	user/1:
		group_$_ids: ["30","2"]
		group_$30_ids: [10]
		group_$2_ids: [2]
	group:
		1:
			meeting_id: 30
		2:
			meeting_id: 2
		10:
			meeting_id: 30
			permissions:
			- agenda_item.can_manage
			- motion.can_see
	agenda_item:
		1:
			meeting_id: 30
			item_number: five
			unknown_field: unknown
			tag_ids: [1,2]
		2:
			meeting_id: 30
			content_object_id: topic/1
			parent_id: 1
		10:
			meeting_id: 2
			item_number: six
	motion/1:
		id: 1
		meeting_id: 30
		origin_id: null
		state_id: 1
	motion_state/1/id: 1
	tag:
		1:
			meeting_id: 30
			tagged_ids: ["agenda_item/1","agenda_item/10"]
		2:
			meeting_id: 2
	
	topic/1:
		id: 1
		meeting_id: 30

	unknown_collection/1/field: 404
	`))

	restricter := restrict.Middleware(ds, 1)

	keys := []string{
		"agenda_item/1/item_number",
		"agenda_item/1/unknown_field",
		"agenda_item/1/tag_ids",
		"agenda_item/10/item_number",
		"unknown_collection/1/field",
		"tag/1/tagged_ids",
		"user/1/group_$_ids",
		"user/1/group_$30_ids",
		"user/1/group_$2_ids",
		"agenda_item/2/content_object_id",
		"agenda_item/2/parent_id",
		"motion/1/origin_id",
	}

	data, err := restricter.Get(context.Background(), keys...)

	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if data["agenda_item/1/item_number"] == nil {
		t.Errorf("agenda_item/1/item_number was removed")
	}

	if data["agenda_item/1/unknown_field"] != nil {
		t.Errorf("agenda_item/1/item_number was not removed")
	}

	if data["agenda_item/10/item_number"] != nil {
		t.Errorf("agenda_item/10/item_number was not removed")
	}

	if data["unknown_collection/1/field"] != nil {
		t.Errorf("unknown_collection/1/field was not removed")
	}

	if got := string(data["tag/1/tagged_ids"]); got != `["agenda_item/1"]` {
		t.Errorf("tag/1/tagged_ids was restricted to %q, expedted %q", got, `["agenda_item/1"]`)
	}

	if got := string(data["agenda_item/1/tag_ids"]); got != `[1]` {
		t.Errorf("agenda_item/1/tag_ids was restricted to %q, expedted %q", got, `[1]`)
	}

	// This should change in the future. meeting 2 is not visible
	if got := string(data["user/1/group_$_ids"]); got != `["30","2"]` {
		t.Errorf("user/1/group_$_ids was restricted to %q, did not expect it", got)
	}

	if got := string(data["user/1/group_$30_ids"]); got != `[10]` {
		t.Errorf("user/1/group_$30_ids was restricted to %q, did not expect it", got)
	}

	if got := string(data["user/1/group_$2_ids"]); got != `[]` {
		t.Errorf("user/1/group_$2_ids is %q, expected a empty list", got)
	}
}

func TestRestrictSuperAdmin(t *testing.T) {
	ds := dsmock.Stub(dsmock.YAMLData(`---
	user/1/organization_management_level: superadmin
	personal_note/1/user_id: 1
	personal_note/2/user_id: 2

	unknown_collection/404/field: 404
	user/404/unknown_field: blub
	`))

	restricter := restrict.Middleware(ds, 1)

	keys := []string{
		"unknown_collection/404/field",
		"personal_note/1/id",
		"personal_note/2/id",
		"user/404/unknown_field",
	}

	got, err := restricter.Get(context.Background(), keys...)
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

func TestCorruptedDatastore(t *testing.T) {
	ds := dsmock.Stub(dsmock.YAMLData(`---
	projector/13:
		meeting_id: 30
		current_projection_ids: [404]

	meeting/30/id: 30
	user/1:
		group_$_ids: ["30"]
		group_$30_ids: [10]
	group:
		10:
			meeting_id: 30
			permissions:
			- projector.can_see
	`))

	restricter := restrict.Middleware(ds, 1)

	var buf bytes.Buffer
	log.SetOutput(&buf)

	got, err := restricter.Get(context.Background(), "projector/13/current_projection_ids")
	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if string(got["projector/13/current_projection_ids"]) != `[]` {
		t.Errorf("projector/13/current_projection_ids == %s, expected an empty list", got["projector/13/current_projection_ids"])
	}

	if !strings.Contains(buf.String(), "Warning") {
		t.Errorf("no warning logged, got: %s", buf.String())
	}
}
