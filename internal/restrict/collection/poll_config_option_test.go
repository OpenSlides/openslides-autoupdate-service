package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/perm"
)

func TestPollConfigOptionModeA(t *testing.T) {
	mode := collection.PollConfigOption{}.Modes("A")

	testCase(
		"no perms",
		t,
		mode,
		false,
		`---
		poll/3:
			meeting_id: 30
			content_object_id: topic/5

		poll_config_selection/2/poll_id: 3

		poll_config_option/1/poll_config_id: poll_config_selection/2

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

		poll_config_selection/2/poll_id: 3
		poll_config_option/1/poll_config_id: poll_config_selection/2

		topic/5:
			meeting_id: 30
			agenda_item_id: 3

		agenda_item/3/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)
}
