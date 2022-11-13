package pendingmap_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/pendingmap"
)

func TestGet_When_a_key_gets_unmarked_while_waiting_an_error_is_returned(t *testing.T) {
	ctx := context.Background()
	pm := pendingmap.New()

	k1, k2 := dskey.MustKey("user/1/username"), dskey.MustKey("user/1/first_name")

	pm.MarkPending(k1, k2)

	// Reader requests 1 and 2
	type Done struct {
		data [][]byte
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

func BenchmarkGet(b *testing.B) {
	ctx := context.Background()

	for _, bt := range []struct {
		keysAmount int
		valueSize  int
	}{
		{100, 100},
		{1_000, 100},
		{10_000, 100},
		{20_000, 100},
		{40_000, 100},
		{60_000, 100},
		{80_000, 100},
		{85_000, 100},
		{90_000, 100},
		{100_000, 100},
		{100, 100_000},
	} {
		b.Run(fmt.Sprintf("%dK %dB", bt.keysAmount, bt.valueSize), func(b *testing.B) {
			pm := pendingmap.New()
			keys := make([]dskey.Key, bt.keysAmount)
			data := make(map[dskey.Key][]byte, bt.keysAmount)
			for i := range keys {
				k := dskey.Key{Collection: "collection", ID: i + 1, Field: "field"}
				keys[i] = k
				data[k] = bytes.Repeat([]byte("X"), bt.valueSize)
			}

			pm.MarkPending(keys...)

			pm.SetIfPending(data)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				pm.Get(ctx, keys...)
			}
		})
		b.Run(fmt.Sprintf("%dK %dB Only V", bt.keysAmount, bt.valueSize), func(b *testing.B) {
			pm := pendingmap.New()
			keys := make([]dskey.Key, bt.keysAmount)
			data := make(map[dskey.Key][]byte, bt.keysAmount)
			for i := range keys {
				k := dskey.Key{Collection: "collection", ID: i + 1, Field: "field"}
				keys[i] = k
				data[k] = bytes.Repeat([]byte("X"), bt.valueSize)
			}

			pm.MarkPending(keys...)

			pm.SetIfPending(data)

			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				pm.Get(ctx, keys...)
			}
		})
	}
}
