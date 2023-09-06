package pendingmap_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/cache/pendingmap"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

func TestGet_When_a_key_gets_unmarked_while_waiting_an_error_is_returned(t *testing.T) {
	ctx := context.Background()
	pm := pendingmap.New()

	k1, k2 := dskey.MustKey("user/1/username"), dskey.MustKey("user/1/first_name")

	pm.MarkPending(k1, k2)

	// Reader requests 1 and 2
	type Done struct {
		data map[dskey.Key][]byte
		err  error
	}
	done := make(chan Done)
	go func() {
		got, err := pm.Get(ctx, k1, k2)
		done <- Done{got, err}
	}()

	time.Sleep(time.Millisecond)

	// Reader is waiting for 1
	pm.UnMarkPending(k1)
	// Reader is waiting for 2
	pm.MarkPending(k1)

	pm.SetIfPending(map[dskey.Key][]byte{k2: []byte("new")})

	result := <-done
	if !errors.Is(result.err, pendingmap.ErrNotExist) {
		t.Errorf("got error: %v, expected %v", result.err, pendingmap.ErrNotExist)
	}

	if result.data != nil {
		t.Errorf("got %v, expected nil", result.data)
	}
}
