package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestPollEntitledUserModeA(t *testing.T) {
	var a collection.PollEntitledUser
	ds := `---
	poll_entitled_user/1/poll_id: 7

	poll/7:
		meeting_id: 30
		content_object_id: topic/3

	topic/3:
		meeting_id: 30
		agenda_item_id: 4

	agenda_item/4:
		meeting_id: 30
	`

	testCase(
		"Can see polls",
		t,
		a.Modes("A"),
		true,
		ds,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"No Perm",
		t,
		a.Modes("A"),
		false,
		ds,
	)
}
