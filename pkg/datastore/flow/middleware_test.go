package flow_test

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

// add100Middleware for testing the API.
// Manipulates int value by adding 100
type add100Middleware struct {
	sub flow.Flow
}

func add100(v []byte) []byte {
	var decoded int
	if err := json.Unmarshal(v, &decoded); err != nil {
		return v
	}

	return []byte(strconv.Itoa(decoded + 100))
}

func (m *add100Middleware) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	values, err := m.sub.Get(ctx, keys...)
	if err != nil {
		return nil, err
	}

	for k, v := range values {
		values[k] = add100(v)
	}
	return values, nil
}

func (m *add100Middleware) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	m.sub.Update(ctx, func(values map[dskey.Key][]byte, err error) {
		for k, v := range values {
			values[k] = add100(v)
		}
		updateFn(values, err)
	})
}

func TestMiddleware(t *testing.T) {
	ctx := context.Background()
	subFlow := newMockFlow(dsmock.YAMLData(`---
	motion/1/start_line_number: 5
	`))
	likesKey := dskey.MustKey("motion/1/start_line_number")

	middleware := add100Middleware{subFlow}

	t.Run("get", func(t *testing.T) {
		got, err := middleware.Get(ctx, likesKey)
		if err != nil {
			t.Errorf("Get: %v", err)
		}

		if got := string(got[likesKey]); got != "105" {
			t.Errorf("got %v, expected 105", got)
		}
	})

	t.Run("update", func(t *testing.T) {
		updateCtx, cancel := context.WithCancel(ctx)
		defer cancel()

		errCh := make(chan error, 1)
		go middleware.Update(
			updateCtx,
			func(got map[dskey.Key][]byte, err error) {
				if err != nil {
					errCh <- fmt.Errorf("update: %v", err)
					return
				}

				expect := map[dskey.Key][]byte{likesKey: []byte("205")}
				if !reflect.DeepEqual(got, expect) {
					errCh <- fmt.Errorf("Got %v, expected %v", got, expect)
					return
				}

				errCh <- nil
			},
		)

		<-subFlow.Registered()
		subFlow.SendUpdate(map[dskey.Key][]byte{likesKey: []byte("105")}, nil)

		if err := <-errCh; err != nil {
			t.Error(err)
		}
	})
}
