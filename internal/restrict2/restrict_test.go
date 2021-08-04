package restrict_test

import (
	"context"
	"testing"

	restrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func TestRestrict(t *testing.T) {
	fetch := datastore.NewFetcher(dsmock.Stub(dsmock.YAMLData(`---
	meeting/1/id: 1
	user/1:
		group_$_ids: ["1"]
		group_$1_ids: [10]
	group/10:
		permissions:
		- agenda_item.can_manage
	agenda_item:
		1:
			meeting_id: 1
		10:
			meeting_id: 2
	`)))

	data := map[string]string{
		"agenda_item/1/item_number":   `"numberOne"`,
		"agenda_item/1/unknown_field": `"numberOne"`,
		"agenda_item/404/item_number": `"numberA"`,
		"agenda_item/10/item_number":  `"numberB"`,
		"unknown_collection/1/field":  "404",
	}

	got := make(map[string][]byte, len(data))
	for k, v := range data {
		got[k] = []byte(v)
	}

	err := restrict.Restrict(context.Background(), fetch, 1, got)

	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if got["agenda_item/1/item_number"] == nil {
		t.Errorf("agenda_item/1/item_number was removed")
	}

	if got["agenda_item/1/unknown_field"] != nil {
		t.Errorf("agenda_item/1/item_number was not removed")
	}

	if got["agenda_item/404/item_number"] != nil {
		t.Errorf("agenda_item/404/item_number was not removed")
	}

	if got["agenda_item/10/item_number"] != nil {
		t.Errorf("agenda_item/404/item_number was not removed")
	}

	if got["unknown_collection/1/field"] != nil {
		t.Errorf("unknown_collection/1/field was not removed")
	}

}
