package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestHistoryEntry(t *testing.T) {
	mode := collection.HistoryEntry{}.Modes("A")

	testCase(
		"no perms",
		t,
		mode,
		false,
		`---
		history_entry/1/position_id: 5
		`,
	)

	testCase(
		"Superadmin",
		t,
		mode,
		true,
		`---
		history_entry/1/position_id: 5
		user/1/organization_management_level: superadmin
		`,
	)
}
