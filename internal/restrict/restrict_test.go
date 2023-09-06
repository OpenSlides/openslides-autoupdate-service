package restrict_test

import (
	"context"
	"testing"

	restrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func TestRestrict(t *testing.T) {
	ctx := context.Background()
	ds := dsmock.Stub(dsmock.YAMLData(`---
	meeting:
		30:
			enable_anonymous: true
		2:
			enable_anonymous: false
			committee_id: 404
		22:
			enable_anonymous: false
			admin_group_id: 32
			committee_id: 404

	user/1:
		meeting_user_ids: [10,11]
	
	meeting_user:
		10:
			meeting_id: 30
			group_ids: [10]
			user_id: 1
		11:
			meeting_id: 2
			group_ids: [2]
			user_id: 1

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
		32:
			meeting_id: 22

	agenda_item:
		1:
			meeting_id: 30
			item_number: five
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
		agenda_item_id: 1
	`))

	ctx, restricter := restrict.Middleware(ctx, ds, 1)

	keys := []dskey.Key{
		dskey.MustKey("agenda_item/1/item_number"),
		dskey.MustKey("agenda_item/1/tag_ids"),
		dskey.MustKey("agenda_item/10/item_number"),
		dskey.MustKey("tag/1/tagged_ids"),
		dskey.MustKey("user/1/meeting_user_ids"),
		dskey.MustKey("meeting_user/10/group_ids"),
		dskey.MustKey("meeting_user/11/group_ids"),
		dskey.MustKey("agenda_item/2/content_object_id"),
		dskey.MustKey("agenda_item/2/parent_id"),
		dskey.MustKey("motion/1/origin_id"),
		dskey.MustKey("meeting/22/admin_group_id"),
	}

	data, err := restricter.Get(ctx, keys...)
	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if data[dskey.MustKey("agenda_item/1/item_number")] == nil {
		t.Errorf("agenda_item/1/item_number was removed")
	}

	if data[dskey.MustKey("agenda_item/10/item_number")] != nil {
		t.Errorf("agenda_item/10/item_number was not removed")
	}

	if got := string(data[dskey.MustKey("tag/1/tagged_ids")]); got != `["agenda_item/1"]` {
		t.Errorf("tag/1/tagged_ids was restricted to %q, expedted %q", got, `["agenda_item/1"]`)
	}

	if got := string(data[dskey.MustKey("agenda_item/1/tag_ids")]); got != `[1]` {
		t.Errorf("agenda_item/1/tag_ids was restricted to %q, expedted %q", got, `[1]`)
	}

	if got := string(data[dskey.MustKey("user/1/meeting_user_ids")]); got != `[10,11]` {
		t.Errorf("user/1/meeting_user_ids was restricted to %q, did not expect it", got)
	}
}

func TestRestrictSuperAdmin(t *testing.T) {
	ctx := context.Background()
	ds := dsmock.Stub(dsmock.YAMLData(`---
	user/1:
		organization_management_level: superadmin
		meeting_user_ids: [10]
	personal_note:
		1:
			meeting_user_id: 10
		2:
			meeting_user_id: 20
	meeting_user:
		10:
			user_id: 1
		20:
			user_id: 200
	`))

	ctx, restricter := restrict.Middleware(ctx, ds, 1)

	keys := []dskey.Key{
		dskey.MustKey("personal_note/1/id"),
		dskey.MustKey("personal_note/2/id"),
	}

	got, err := restricter.Get(ctx, keys...)
	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if got[dskey.MustKey("personal_note/1/id")] == nil {
		t.Errorf("personal_note/1/id got restricted")
	}

	if got[dskey.MustKey("personal_note/2/id")] != nil {
		t.Errorf("personal_note/2/id got not restricted")
	}
}
