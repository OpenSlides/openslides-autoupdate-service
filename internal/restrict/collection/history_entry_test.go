package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestHistoryEntry(t *testing.T) {
	mode := collection.HistoryEntry{}.Modes("A")

	testCase(
		"no perms",
		t,
		mode,
		false,
		`---
		history_entry/1:
			model_id: user/5
		`,
	)

	testCase(
		"orga admin",
		t,
		mode,
		true,
		`---
		history_entry/1:
			model_id: user/5

		user/1/organization_management_level: can_manage_organization
		`,
	)

	testCase(
		"meeting specific collection with perm",
		t,
		mode,
		true,
		`---
		history_entry/1:
			model_id: motion/5
		motion/5/meeting_id: 77
		`,
		withPerms(77, perm.MeetingCanSeeHistory),
	)
}
