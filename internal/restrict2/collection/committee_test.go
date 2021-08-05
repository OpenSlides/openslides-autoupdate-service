package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestCommitteeModeA(t *testing.T) {
	var c collection.Committee

	for _, tt := range []testData{
		testCase(
			"No perms",
			false,
			`committee/1/id: 1`,
		),

		testCase(
			"In committee/user_ids",
			true,
			`---
			committee/1/user_ids: [1]
			`,
		),

		testCase(
			"OML can_manage_users",
			true,
			`---
			committee/1/id: 1
			user/1/organization_management_level: can_manage_organization
			`,
		),
	} {
		tt.test(t, c.Modes("A"))
	}

}
