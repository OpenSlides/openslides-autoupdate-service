package flow_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

func TestCombine(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userFlow := newMockFlow(dsmock.YAMLData(`---
	user/1/username: max
	`))

	groupFlow := newMockFlow(dsmock.YAMLData(`---
	group/1/name: admin
	`))

	userKey := dskey.MustKey("user/1/username")
	groupKey := dskey.MustKey("group/1/name")

	combined := flow.Combine(userFlow, map[string]flow.Flow{"group/name": groupFlow})
	var _ flow.Flow = combined

	t.Run("get both keys", func(t *testing.T) {
		got, err := combined.Get(ctx, userKey, groupKey)
		if err != nil {
			t.Fatalf("Get: %v", err)
		}

		expect := map[dskey.Key][]byte{
			userKey:  []byte(`"max"`),
			groupKey: []byte(`"admin"`),
		}

		if !reflect.DeepEqual(got, expect) {
			t.Errorf("got %v, expected %v", got, expect)
		}
	})

	t.Run("one has error", func(t *testing.T) {
		userFlow.err = errors.New("some error")
		defer func() { userFlow.err = nil }()

		_, err := combined.Get(ctx, userKey, groupKey)
		if err == nil {
			t.Errorf("did not get an error, expected one.")
		}
	})

	t.Run("update user", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer userFlow.ClearUpdate()
		defer groupFlow.ClearUpdate()

		errCh := make(chan error, 1)
		go combined.Update(
			ctx,
			func(got map[dskey.Key][]byte, err error) {
				if err != nil {
					errCh <- fmt.Errorf("Update: %v", err)
					return
				}

				expect := map[dskey.Key][]byte{userKey: []byte(`"new name"`)}
				if !reflect.DeepEqual(got, expect) {
					errCh <- fmt.Errorf("Got %v, expected %v", got, expect)
					return
				}

				errCh <- nil
			},
		)

		<-userFlow.Registered()
		userFlow.SendUpdate(map[dskey.Key][]byte{userKey: []byte(`"new name"`)}, nil)

		if err := <-errCh; err != nil {
			t.Error(err)
		}
	})

	t.Run("update group", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer userFlow.ClearUpdate()
		defer groupFlow.ClearUpdate()

		errCh := make(chan error, 1)
		go combined.Update(
			ctx,
			func(got map[dskey.Key][]byte, err error) {
				if err != nil {
					errCh <- fmt.Errorf("Update: %v", err)
					return
				}

				expect := map[dskey.Key][]byte{groupKey: []byte(`"new name"`)}
				if !reflect.DeepEqual(got, expect) {
					errCh <- fmt.Errorf("Got %v, expected %v", got, expect)
					return
				}

				errCh <- nil
			},
		)

		<-groupFlow.Registered()
		groupFlow.SendUpdate(map[dskey.Key][]byte{groupKey: []byte(`"new name"`)}, nil)

		if err := <-errCh; err != nil {
			t.Error(err)
		}
	})
}
