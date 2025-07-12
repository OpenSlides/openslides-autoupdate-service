package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestHistoryPosition(t *testing.T) {
	mode := collection.HistoryPosition{}.Modes("A")

	testCase(
		"no perms",
		t,
		mode,
		false,
		`---
		`,
	)

	testCase(
		"Superadmin",
		t,
		mode,
		true,
		`---
		user/1/organization_management_level: superadmin
		`,
	)
}
