package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestPointOfOrderCategoryModeA(t *testing.T) {
	var p collection.PointOfOrderCategory

	testCase(
		"Can not see meeting",
		t,
		p.Modes("A"),
		false,
		`---
		point_of_order_category/1/meeting_id: 5
		meeting/5:
			committee_id: 404
		`,
	)

	testCase(
		"Can see the meeting",
		t,
		p.Modes("A"),
		true,
		`---
		organization/1/enable_anonymous: true
		point_of_order_category/1/meeting_id: 5
		meeting/5:
			enable_anonymous: true
			committee_id: 404
		`,
	)
}
