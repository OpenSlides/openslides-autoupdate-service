package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestActionWorkerModeA(t *testing.T) {
	var a collection.ActionWorker

	testCase(
		"Other User",
		t,
		a.Modes("A"),
		false,
		`---
		action_worker/1/user_id: 5
		`,
		withRequestUser(1),
	)

	testCase(
		"Same User",
		t,
		a.Modes("A"),
		true,
		`---
		action_worker/1/user_id: 5
		`,
		withRequestUser(5),
	)
}
