package autoupdate_test

import (
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
)

func getConnection(closed <-chan struct{}) (*autoupdate.Connection, *test.MockDatastore) {
	datastore := test.NewMockDatastore(closed, map[string]string{
		"user/1/name": `"Hello World"`,
	})
	s := autoupdate.New(datastore, test.RestrictAllowed(), test.UserUpdater{}, closed)
	kb := test.KeysBuilder{K: test.Str("user/1/name")}
	c := s.Connect(1, kb)

	return c, datastore
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
