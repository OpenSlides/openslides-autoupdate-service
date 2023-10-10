package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestPersonalNoteModeA(t *testing.T) {
	t.Parallel()
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
		personal_note/1/meeting_user_id: 50
		meeting_user/50/user_id : 5
		`,
		withRequestUser(2),
	)
}
