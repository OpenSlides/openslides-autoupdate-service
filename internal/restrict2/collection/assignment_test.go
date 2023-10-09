package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestAssignmentModeB(t *testing.T) {
	t.Parallel()
	var a collection.Assignment

	testCase(
		"Without perms",
		t,
		a.Modes("A"),
		false,
		`---
		assignment/1:
			meeting_id: 30
		`,
	)

	testCase(
		"Can see",
		t,
		a.Modes("A"),
		true,
		`---
		assignment/1:
			meeting_id: 30
		`,
		withPerms(30, perm.AssignmentCanSee),
	)
}
