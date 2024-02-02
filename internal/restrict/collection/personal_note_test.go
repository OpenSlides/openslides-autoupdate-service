package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestPersonalNoteModeA(t *testing.T) {
	var p collection.PersonalNote

	testCase(
		"own note",
		t,
		p.Modes("A"),
		true,
		`---
		personal_note/1/meeting_user_id: 5
		meeting_user/5/user_id: 1
		user/1/meeting_user_ids: [5]
		`,
		withRequestUser(1),
	)

	testCase(
		"Other note",
		t,
		p.Modes("A"),
		false,
		`---
		personal_note/1/meeting_user_id: 5
		user/2/meeting_user_ids: [4]
		`,
		withRequestUser(2),
	)
}

func TestPersonalNoteSuperAdminModeA(t *testing.T) {
	var p collection.PersonalNote
	ds := `---
	personal_note/1/meeting_user_id: 5
	meeting_user/5/user_id: 2
	`

	testCase(
		"Other note",
		t,
		p.SuperAdmin("A"),
		false,
		ds,
		withRequestUser(2),
	)
}
