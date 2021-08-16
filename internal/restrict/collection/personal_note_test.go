package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestPersonalNoteModeA(t *testing.T) {
	var p collection.PersonalNote
	ds := `personal_note/1/user_id: 1`

	testCase(
		"own note",
		t,
		p.Modes("A"),
		true,
		ds,
		withRequestUser(1),
	)

	testCase(
		"Other note",
		t,
		p.Modes("A"),
		false,
		ds,
		withRequestUser(2),
	)
}

func TestPersonalNoteSuperAdminModeA(t *testing.T) {
	var p collection.PersonalNote
	ds := `personal_note/1/user_id: 1`

	testCase(
		"Other note",
		t,
		p.SuperAdmin("A"),
		false,
		ds,
		withRequestUser(2),
	)
}
