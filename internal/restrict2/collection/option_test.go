package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestOptionModeA(t *testing.T) {
	f := collection.Option{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		option/1/poll_id: 3
		poll/3/meeting_id: 1
		meeting/1/enable_anonymous: false
		`,
	)

	testCase(
		"can see poll",
		t,
		f,
		true,
		`---
		option/1/poll_id: 3
		poll/3/meeting_id: 1
		meeting/1/enable_anonymous: true
		`,
	)
}

func TestOptionModeB(t *testing.T) {
	f := collection.Option{}.Modes("B")

	testCase(
		"no perms",
		t,
		f,
		false,
		`---
		option/1/poll_id: 3
		poll/3/meeting_id: 1
		`,
	)

	testCase(
		"can see poll",
		t,
		f,
		true,
		`---
		option/1/poll_id: 3
		poll/3/meeting_id: 1
		`,
		withPerms(1, perm.PollCanManage),
	)
}
