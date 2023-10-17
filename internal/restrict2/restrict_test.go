package restrict_test

import (
	"context"
	"reflect"
	"testing"

	restrict "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

// TODO Fix me
func TestRestritGet(t *testing.T) {
	ctx, cancel := context.WithCancel(collection.ContextWithRestrictCache(context.Background()))
	defer cancel()

	flow := dsmock.NewFlow(
		dsmock.YAMLData(`---
		user/1/meeting_user_ids: [10]
		meeting_user/10:
			group_ids: [7]
			meeting_id: 5

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
			- motion.can_see
			meeting_id: 5
			meeting_user_ids: [10]
		motion/13:
			title: foobar
			meeting_id: 5
			state_id: 15
		motion_state/15/id: 15

		committee/2/manager_ids: []
		`),
		dsmock.NewCounter,
	)
	counter := flow.Middlewares()[0].(*dsmock.Counter)

	restricter := restrict.New(flow)

	waiter := make(chan struct{}, 1)
	go restricter.Update(ctx, func(map[dskey.Key][]byte, error) {
		waiter <- struct{}{}
	})

	ctx, userRestricter, _ := restricter.ForUser(ctx, 1)
	keys := []dskey.Key{
		dskey.MustKey("agenda_item/1/weight"),
		dskey.MustKey("agenda_item/1/duration"),
		dskey.MustKey("agenda_item/1/comment"),
		dskey.MustKey("meeting/5/group_ids"),
		dskey.MustKey("motion/13/title"),
	}

	got, err := userRestricter.Get(ctx, keys...)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := map[dskey.Key][]byte{
		dskey.MustKey("agenda_item/1/weight"):   []byte("50"),
		dskey.MustKey("agenda_item/1/duration"): nil,
		dskey.MustKey("agenda_item/1/comment"):  nil,
		dskey.MustKey("meeting/5/group_ids"):    []byte("[7]"),
		dskey.MustKey("motion/13/title"):        []byte(`"foobar"`),
	}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got\n%v\n\nexpected\n%v", got, expect)
	}

	t.Run("Update another key", func(t *testing.T) {
		counter.Reset()
		flow.Send(map[dskey.Key][]byte{dskey.MustKey("agenda_item/1/item_number"): []byte(`"some value"`)})
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

		if counter.Count() != 3 {
			t.Errorf("Got %d requests to the db. Expected 3 (to get the user attributes and the weight key): %v", counter.Count(), counter.Requests())
		}
	})
}

func TestRestrict(t *testing.T) {
	ctx, cancel := context.WithCancel(collection.ContextWithRestrictCache(context.Background()))
	defer cancel()

	flow := dsmock.NewFlow(dsmock.YAMLData(`---
	meeting:
		30:
			enable_anonymous: true
			group_ids: [1,10]
		2:
			enable_anonymous: false
			committee_id: 404
			group_ids: [2]
		22:
			enable_anonymous: false
			admin_group_id: 32
			committee_id: 404
			group_ids: [32]
		404:
			enable_anonymous: false
			committee_id: 404

	user/1:
		meeting_user_ids: [10,11]
	
	meeting_user:
		10:
			meeting_id: 30
			group_ids: [10]
			user_id: 1
		11:
			meeting_id: 2
			group_ids: [2]
			user_id: 1

	group:
		1:
			meeting_id: 30
		2:
			meeting_id: 2
			meeting_user_ids: [2]
		10:
			meeting_id: 30
			meeting_user_ids: [10]
			permissions:
			- agenda_item.can_manage
			- motion.can_see
		32:
			meeting_id: 22

	agenda_item:
		1:
			meeting_id: 30
			item_number: five
			tag_ids: [1,2]
			content_object_id: assignment/7
		2:
			meeting_id: 30
			content_object_id: topic/1
			parent_id: 1
		10:
			meeting_id: 2
			item_number: six
	motion/1:
		id: 1
		meeting_id: 30
		origin_id: null
		state_id: 1
	motion_state/1/id: 1
	assignment/7/meeting_id: 30
	tag:
		1:
			meeting_id: 30
			tagged_ids: ["agenda_item/1","agenda_item/10"]
		2:
			meeting_id: 404
	
	topic/1:
		id: 1
		meeting_id: 30
		agenda_item_id: 1
	
	committee/404/id: 404
	`))

	restricter := restrict.New(flow)

	keys := []dskey.Key{
		dskey.MustKey("agenda_item/1/item_number"),
		dskey.MustKey("agenda_item/1/tag_ids"),
		dskey.MustKey("agenda_item/1/content_object_id"),
		dskey.MustKey("agenda_item/10/item_number"),
		dskey.MustKey("tag/1/tagged_ids"),
		dskey.MustKey("user/1/meeting_user_ids"),
		dskey.MustKey("meeting_user/10/group_ids"),
		dskey.MustKey("meeting_user/11/group_ids"),
		dskey.MustKey("agenda_item/2/content_object_id"),
		dskey.MustKey("agenda_item/2/parent_id"),
		dskey.MustKey("motion/1/origin_id"),
		dskey.MustKey("meeting/22/admin_group_id"),
	}

	ctx, forUser, _ := restricter.ForUser(ctx, 1)

	data, err := forUser.Get(ctx, keys...)
	if err != nil {
		t.Fatalf("Restrict returned: %v", err)
	}

	if data[dskey.MustKey("agenda_item/1/item_number")] == nil {
		t.Errorf("agenda_item/1/item_number was removed")
	}

	if data[dskey.MustKey("agenda_item/1/content_object_id")] != nil {
		t.Errorf("agenda_item/1/content_object_id was not removed")
	}

	if data[dskey.MustKey("agenda_item/10/item_number")] != nil {
		t.Errorf("agenda_item/10/item_number was not removed")
	}

	if got := string(data[dskey.MustKey("tag/1/tagged_ids")]); got != `["agenda_item/1"]` {
		t.Errorf("tag/1/tagged_ids was restricted to `%s`, expected `%s`", got, `["agenda_item/1"]`)
	}

	if got := string(data[dskey.MustKey("agenda_item/1/tag_ids")]); got != `[1]` {
		t.Errorf("agenda_item/1/tag_ids was restricted to `%s`, expected `%s`", got, `[1]`)
	}

	if got := string(data[dskey.MustKey("user/1/meeting_user_ids")]); got != `[10,11]` {
		t.Errorf("user/1/meeting_user_ids was restricted to `%s`, did not expect it", got)
	}
}
