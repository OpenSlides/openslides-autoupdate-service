package pendingmap_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/pendingmap"
)

func TestGet_When_a_key_gets_unmarked_while_waiting_an_error_is_returned(t *testing.T) {
	ctx := context.Background()
	pm := pendingmap.New[int, int]()

	pm.MarkPending(1, 2)

	// Reader requests 1 and 2
	type Done struct {
		data map[int]int
		err  error
	}
	done := make(chan Done)
	go func() {
		got, err := pm.Get(ctx, 1, 2)
		done <- Done{got, err}
	}()

	time.Sleep(time.Millisecond)

	// Reader is waiting for 1
	pm.UnMarkPending(1)
	// Reader is waiting for 2
	pm.MarkPending(1)

	pm.SetIfPending(map[int]int{2: 99})

	result := <-done
	if !errors.Is(result.err, pendingmap.ErrNotExist) {
		t.Errorf("got error: %v, expected %v", result.err, pendingmap.ErrNotExist)
	}

	if result.data != nil {
		t.Errorf("got %v, expected nil", result.data)
	}
}
