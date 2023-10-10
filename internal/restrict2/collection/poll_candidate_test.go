package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestPollCandidateModeA(t *testing.T) {
	t.Parallel()
	mode := collection.PollCandidate{}.Modes("A")

	testCase(
		"no permission",
		t,
		mode,
		false,
		`---
		poll/3:
			content_object_id: topic/4
		option/5/poll_id: 3	
		poll_candidate_list/23:
			option_id: 5
		poll_candidate/42:
			poll_candidate_list_id: 23
		topic/4:
			meeting_id: 30
			agenda_item_id: 7
		agenda_item/7/meeting_id: 30
		`,
		withElementID(42),
	)

	testCase(
		"can see",
		t,
		mode,
		true,
		`---
		poll/3:
			content_object_id: topic/4
		option/5/poll_id: 3	
		poll_candidate_list/23:
			option_id: 5
		poll_candidate/42:
			poll_candidate_list_id: 23
		topic/4:
			meeting_id: 30
			agenda_item_id: 7
		agenda_item/7/meeting_id: 30
		`,
		withElementID(42),
		withPerms(30, perm.AgendaItemCanSee),
	)
}
