package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestMeetingPollSettingModeA(t *testing.T) {
	m := collection.MeetingPollSetting{}.Modes("A")
	testCase(
		"No perms",
		t,
		m,
		false,
		`---
		meeting/30:
			id: 1
			committee_id: 300

		meeting_poll_setting/1:
			meeting_id: 30
		`,
	)

	testCase(
		"user in meeting",
		t,
		m,
		true,
		`---
		meeting/30:
			group_ids: [7]
			committee_id: 2

		group/7/meeting_user_ids: [10]
		meeting_user/10:
			user_id: 1
			meeting_id: 30
		user/1/meeting_user_ids: [10]

		meeting_poll_setting/1:
			meeting_id: 30
		`,
	)
}
