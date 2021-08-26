package autoupdate_test

import (
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func getConnection(closed <-chan struct{}) (autoupdate.DataProvider, *dsmock.MockDatastore) {
	datastore := dsmock.NewMockDatastore(closed, map[string][]byte{
		"user/1/name": []byte(`"Hello World"`),
	})
	s := autoupdate.New(datastore, test.RestrictAllowed, closed)
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
