package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestPointOfOrderCategoryModeA(t *testing.T) {
	t.Parallel()
	var p collection.PointOfOrderCategory

	testCase(
		"Can not see meeting",
		t,
		p.Modes("A"),
		false,
		`---
		point_of_order_category/1/meeting_id: 5
		meeting/5:
			enable_anonymous: false
			committee_id: 40
		committee/40/id: 40
		`,
	)

	testCase(
		"Can see the meeting",
		t,
		p.Modes("A"),
		true,
		`---
		point_of_order_category/1/meeting_id: 5
		meeting/5:
			enable_anonymous: true
			committee_id: 404
		`,
	)
}
