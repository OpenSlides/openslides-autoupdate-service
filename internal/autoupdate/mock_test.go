package autoupdate_test

import (
	"context"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func getConnection(ctx context.Context) (autoupdate.DataProvider, *dsmock.MockDatastore) {
	datastore := dsmock.NewMockDatastore(ctx, map[string]string{
		"user/1/name": `"Hello World"`,
	})
	s := autoupdate.New(datastore, test.RestrictAllowed, ctx.Done())
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	next := s.Connect(1, kb)

	return next, datastore
}

func blocking(f func()) bool {
	return blockingTime(10*time.Millisecond, f)
}

// blockingDebug can be used in debug sessions.
func blockingDebug(f func()) bool {
	return blockingTime(time.Hour, f)
}

func blockingTime(wait time.Duration, f func()) bool {
	done := make(chan struct{})
	go func() {
		f()
		close(done)
	}()

	timer := time.NewTimer(wait)
	defer timer.Stop()
	select {
	case <-done:
		return false
	case <-timer.C:
		return true
	}
}
