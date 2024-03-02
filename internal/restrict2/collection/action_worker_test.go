package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestActionWorkerModeA(t *testing.T) {
	var a collection.ActionWorker

	testCase(
		"No permission",
		t,
		a.Modes("A"),
		true,
		`---
		action_worker/1/id: 30
		`,
	)
}
