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
		history_entry/1/meeting_id: 5
		`,
	)

	testCase(
		"orga admin on meeting collection",
		t,
		mode,
		true,
		`---
		history_entry/1/meeting_id: 5
		user/1/organization_management_level: can_manage_organization
		`,
	)

	testCase(
		"orga admin on non meeting collection",
		t,
		mode,
		true,
		`---
		history_entry/1/meeting_id: null
		user/1/organization_management_level: can_manage_organization
		`,
	)

	testCase(
		"with perm on meeting collection",
		t,
		mode,
		true,
		`---
		history_entry/1/meeting_id: 5
		`,
		withPerms(5, perm.MeetingCanSeeHistory),
	)
}
