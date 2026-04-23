package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestPollConfigStvScottishModeA(t *testing.T) {
	mode := collection.PollConfigStvScottish{}.Modes("A")

	testCase(
		"no perms",
		t,
		mode,
		false,
		`---
		poll/3:
			meeting_id: 30
			content_object_id: topic/5

		poll_config_stv_scottish/1/poll_id: 3

		topic/5:
			meeting_id: 30
			agenda_item_id: 3

		agenda_item/3/meeting_id: 30
		`,
	)

	testCase(
		"can see poll",
		t,
		mode,
		true,
		`---
		poll/3:
			meeting_id: 30
			content_object_id: topic/5

		poll_config_stv_scottish/1/poll_id: 3

		topic/5:
			meeting_id: 30
			agenda_item_id: 3

		agenda_item/3/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)
}
