package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestTopicModeA(t *testing.T) {
	f := collection.Topic{}.Modes("A")

	testCase(
		"no perm",
		t,
		f,
		false,
		`---
		topic/1:
			meeting_id: 2
			agenda_item_id: 7
		agenda_item/7/meeting_id: 30
		`,
	)

	testCase(
		"see agenda item",
		t,
		f,
		true,
		`---
		topic/1:
			meeting_id: 30
			agenda_item_id: 3

		agenda_item/3:
			is_internal: true
			meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"can not see agenda_item",
		t,
		f,
		false,
		`---
		topic/1:
			meeting_id: 30
			agenda_item_id: 3

		agenda_item/3:
			is_internal: true
			meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)
}
