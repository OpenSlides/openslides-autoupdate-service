package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestAssignmentModeB(t *testing.T) {
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
		withPerms(30),
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
