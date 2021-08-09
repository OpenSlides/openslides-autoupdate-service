package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestPersonalNoteModeA(t *testing.T) {
	var p collection.PersonalNote
	ds := `personal_note/1/user_id: 1`

	testCase(
		"own note",
		true,
		ds,
		withRequestUser(1),
	).test(t, p.Modes("A"))

	testCase(
		"Other note",
		false,
		ds,
		withRequestUser(2),
	).test(t, p.Modes("A"))
}

func TestPersonalNoteSuperAdminModeA(t *testing.T) {
	var p collection.PersonalNote
	ds := `personal_note/1/user_id: 1`

	testCase(
		"Other note",
		false,
		ds,
		withRequestUser(2),
	).test(t, p.SuperAdmin("A"))
}
