package restrict_test

import (
	"context"
	"reflect"
	"testing"

	restrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func TestRestritGet(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		user/1/group_$_ids: ["5"]
		user/1/group_$5_ids: [7]
		agenda_item:
			1:
				meeting_id: 5
				weight: 50
				duration: test
				comment: my agenda item
		meeting/5:
			group_ids: [7]
			committee_id: 2
		group/7:
			permissions:
			- agenda_item.can_see
			meeting_id: 5
			user_ids: [1]
		`),
		dsmock.NewCounter,
	)
	counter := flow.Middlewares()[0].(*dsmock.Counter)

	restricter := restrict.New(flow)

	waiter := make(chan struct{}, 1)
	go restricter.Update(ctx, func(map[dskey.Key][]byte, error) {
		waiter <- struct{}{}
	})

	ctx, userRestricter := restricter.ForUser(ctx, 1)
	keys := []dskey.Key{
		dskey.MustKey("agenda_item/1/weight"),
		dskey.MustKey("agenda_item/1/duration"),
		dskey.MustKey("agenda_item/1/comment"),
		dskey.MustKey("agenda_item/1/does_not_exist"),
		dskey.MustKey("meeting/5/group_ids"), // Old restricter
	}

	got, err := userRestricter.Get(ctx, keys...)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := map[dskey.Key][]byte{
		dskey.MustKey("agenda_item/1/weight"):         []byte("50"),
		dskey.MustKey("agenda_item/1/duration"):       nil,
		dskey.MustKey("agenda_item/1/comment"):        nil,
		dskey.MustKey("agenda_item/1/does_not_exist"): nil,
		dskey.MustKey("meeting/5/group_ids"):          []byte("[7]"),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got\n%v\n\nexpected\n%v", got, expect)
	}

	t.Run("Update another key", func(t *testing.T) {
		counter.Reset()
		flow.Send(map[dskey.Key][]byte{dskey.MustKey("agenda_item/1/other"): []byte(`"some value"`)})
		<-waiter

		if counter.Count() != 0 {
			t.Errorf("update an irrelevant key made an update")
		}
	})

	t.Run("update an not hot key", func(t *testing.T) {
		counter.Reset()
		flow.Send(map[dskey.Key][]byte{dskey.MustKey("agenda_item/1/weight"): []byte(`"some value"`)})
		<-waiter

		if counter.Count() != 0 {
			t.Errorf("update an irrelevant key made an update")
		}
	})

	t.Run("update an hot key", func(t *testing.T) {
		counter.Reset()
		flow.Send(map[dskey.Key][]byte{dskey.MustKey("agenda_item/1/is_internal"): []byte(`true`)})
		<-waiter

		if counter.Count() == 0 {
			t.Errorf("No update")
		}
	})

	t.Run("get a key again", func(t *testing.T) {
		counter.Reset()
		if _, err := userRestricter.Get(ctx, dskey.MustKey("agenda_item/1/weight")); err != nil {
			t.Fatalf("Get: %v", err)
		}

		if counter.Count() != 6 {
			t.Errorf("Got %d requests to the db. Expected 6 (to get the weight key and to build perm.Permission): %v", counter.Count(), counter.Requests())
		}
	})
}