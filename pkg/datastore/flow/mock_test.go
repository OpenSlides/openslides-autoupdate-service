package flow_test

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
)

type mockFlow struct {
	data map[dskey.Key][]byte
	err  error

	mu         sync.Mutex
	updateFn   func(map[dskey.Key][]byte, error)
	registerCh chan struct{}
}

func newMockFlow(data map[dskey.Key][]byte) *mockFlow {
	if data == nil {
		data = make(map[dskey.Key][]byte)
	}

	return &mockFlow{
		data:       data,
		registerCh: make(chan struct{}),
	}
}

func (m *mockFlow) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	r := make(map[dskey.Key][]byte, len(keys))
	for _, key := range keys {
		r[key] = m.data[key]
	}
	return r, m.err
}

func (m *mockFlow) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	m.mu.Lock()
	m.updateFn = updateFn
	close(m.registerCh)
	m.mu.Unlock()

	<-ctx.Done()

	return
}

func (m *mockFlow) Registered() <-chan struct{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.registerCh
}

// ClearUpdate waits until the update function is clean again.
func (m *mockFlow) ClearUpdate() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.updateFn = nil
	m.registerCh = make(chan struct{})
}

func (m *mockFlow) SendUpdate(data map[dskey.Key][]byte, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if fn := m.updateFn; fn != nil {
		fn(data, err)
	}
}

func TestFlowMock(t *testing.T) {
	ctx := context.Background()
	flw := newMockFlow(dsmock.YAMLData(`---
	user/1/username: max
	`))

	var _ flow.Flow = flw

	got, err := flw.Get(ctx, dskey.MustKey("user/1/username"))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := map[dskey.Key][]byte{dskey.MustKey("user/1/username"): []byte(`"max"`)}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got %v, expected %v", got, expect)
	}

	updateCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	errCh := make(chan error, 1)
	go flw.Update(
		updateCtx,
		func(got map[dskey.Key][]byte, err error) {
			if err != nil {
				errCh <- fmt.Errorf("update: %v", err)
				return
			}

			expect := map[dskey.Key][]byte{dskey.MustKey("user/1/username"): []byte(`"new name"`)}
			if !reflect.DeepEqual(got, expect) {
				errCh <- fmt.Errorf("Got %v, expected %v", got, expect)
			}

			errCh <- nil
		},
	)

	<-flw.Registered()
	flw.SendUpdate(map[dskey.Key][]byte{dskey.MustKey("user/1/username"): []byte(`"new name"`)}, nil)

	if err := <-errCh; err != nil {
		t.Error(err)
	}
}
