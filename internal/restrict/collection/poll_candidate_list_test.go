package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestPollCandidateListModeA(t *testing.T) {
	mode := collection.PollCandidateList{}.Modes("A")

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
		topic/4/meeting_id: 30
		`,
		withElementID(23),
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
		topic/4/meeting_id: 30
		`,
		withElementID(23),
		withPerms(30, perm.AgendaItemCanSee),
	)
}
