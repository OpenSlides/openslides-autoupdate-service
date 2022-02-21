package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestMotionBlockModeA(t *testing.T) {
	f := collection.MotionBlock{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		motion_block/1:
			id: 1
			meeting_id: 2
		`,
	)

	testCase(
		"can manage",
		t,
		f,
		true,
		`---
		motion_block/1/meeting_id: 30
		`,
		withPerms(30, perm.MotionCanManage),
	)

	testCase(
		"see agenda item",
		t,
		f,
		true,
		`---
		motion_block/1:
			meeting_id: 30
			agenda_item_id: 3
		
		agenda_item/3/meeting_id: 2
		`,
		withPerms(2, perm.AgendaItemCanSee),
	)

	testCase(
		"not see agenda item",
		t,
		f,
		false,
		`---
		motion_block/1:
			meeting_id: 30
			agenda_item_id: 3
		
		agenda_item/3/meeting_id: 2
		`,
	)
}

func TestMotionBlockModeB(t *testing.T) {
	f := collection.MotionBlock{}.Modes("B")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		motion_block/1:
			id: 1
			meeting_id: 30
		`,
	)

	testCase(
		"can manage",
		t,
		f,
		true,
		`---
		motion_block/1/meeting_id: 30
		`,
		withPerms(30, perm.MotionCanManage),
	)

	testCase(
		"can see not internal",
		t,
		f,
		true,
		`---
		motion_block/1:
			meeting_id: 30
			internal: false
		`,
		withPerms(30, perm.MotionCanSee),
	)

	testCase(
		"can see internal",
		t,
		f,
		false,
		`---
		motion_block/1:
			meeting_id: 30
			internal: true
		`,
		withPerms(30, perm.MotionCanSee),
	)
}
