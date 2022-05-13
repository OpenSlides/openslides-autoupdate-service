package autoupdate_test

import (
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func getConnection(closed <-chan struct{}) (autoupdate.DataProvider, *dsmock.MockDatastore) {
	datastore := dsmock.NewMockDatastore(closed, map[datastore.Key][]byte{
		userNameKey: []byte(`"Hello World"`),
	})
	s := autoupdate.New(datastore, test.RestrictAllowed, "")
	kb, _ := keysbuilder.FromKeys(userNameKey.String())
	next := s.Connect(1, kb)

	return next, datastore
}

func blocking(f func()) bool {
	return blockingTime(10*time.Millisecond, f)
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
