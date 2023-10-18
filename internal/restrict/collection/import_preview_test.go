package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestImportPreviewModeA(t *testing.T) {
	var i collection.ImportPreview

	testCase(
		"no perms",
		t,
		i.Modes("A"),
		false,
		``,
	)

	testCase(
		"organization manager",
		t,
		i.Modes("A"),
		false,
		`---
		user/1/organization_management_level: can_manage_organization
		`,
	)
}
