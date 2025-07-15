package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestHistoryPosition(t *testing.T) {
	mode := collection.HistoryPosition{}.Modes("A")

	testCase(
		"Can not see entry",
		t,
		mode,
		false,
		`---
		history_position/1/entry_ids: [5]
		history_entry/5/model_id: user/7
		`,
	)

	testCase(
		"Can see entry",
		t,
		mode,
		true,
		`---
		history_position/1/entry_ids: [5]
		history_entry/5/model_id: user/7

		user/1/organization_management_level: can_manage_organization
		`,
	)

	testCase(
		"Can see one",
		t,
		mode,
		true,
		`---
		history_position/1/entry_ids: [5, 6]
		history_entry:
			5:
				model_id: user/7
			6:
				model_id: motion/8

		motion/8/meeting_id: 77

		user/1/organization_management_level: can_manage_organization
		`,
		withPerms(77, perm.MeetingCanSeeHistory),
	)
}
