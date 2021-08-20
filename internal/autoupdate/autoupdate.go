// Package autoupdate allows clients to request keys and get updates when the
// keys changes.
//
// To register to the autoupdate serive, a client has to receive a connection
// object by calling the Connect()-method. It is not necessary and therefore not
// possible to close a connection. The client can just stop listening.
package autoupdate

import (
	"context"
	"fmt"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/ostcar/topic"
)

const (
	// pruneTime defines how long a topic id will be valid. If a client needs
	// more time to process the data, it will get an error and has to reconnect.
	// A higher value means, that more memory is used.
	pruneTime = 10 * time.Minute

	// cacheResetTime defines when the cache should be reseted.
	//
	// When the datastore runs for a long time, its cache grows bigger and more
	// calculated keys have to be calculated. A reset means, that everything
	// gets cleaned.
	//
	// A high value means more memory and cpu usage after some time. A lower
	// value means more Requests to the Datastore Service and therefore a slower
	// responce time for the clients.
	datastoreCacheResetTime = 24 * time.Hour
)

// Format of keys in the topic that shows, that a full update is necessary. It
// is in the same namespace then model names. So make sure, there is no model
// with this name.
const fullUpdateFormat = "fullupdate/%d"

// Autoupdate holds the state of the autoupdate service. It has to be initialized
// with autoupdate.New().
type Autoupdate struct {
	datastore  Datastore
	topic      *topic.Topic
	restricter RestrictMiddleware
}

// RestrictMiddleware is a function that can restrict data.
type RestrictMiddleware func(getter datastore.Getter, uid int) datastore.Getter

// New creates a new autoupdate service.
func New(datastore Datastore, restricter RestrictMiddleware, closed <-chan struct{}) *Autoupdate {
	a := &Autoupdate{
		datastore:  datastore,
		topic:      topic.New(topic.WithClosed(closed)),
		restricter: restricter,
	}

	// Update the topic when an data update is received.
	a.datastore.RegisterChangeListener(func(data map[string][]byte) error {
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}

		a.topic.Publish(keys...)
		return nil
	})

	go a.pruneTopic(closed)
	go a.resetCache(closed)

	return a
}

// DataProvider is a function that returns the next data for a user.
type DataProvider func(ctx context.Context) (map[string][]byte, error)

// Connect has to be called by a client to register to the service. The method
// returns a Connection object, that can be used to receive the data.
//
// There is no need to "close" the Connection object.
func (a *Autoupdate) Connect(userID int, kb KeysBuilder) DataProvider {
	c := &Connection{
		autoupdate: a,
		uid:        userID,
		kb:         kb,
	}

	return c.Next
}

// LastID returns the id of the last data update.
func (a *Autoupdate) LastID() uint64 {
	return a.topic.LastID()
}

// pruneTopic removes old data from the topic. Blocks until the service is
// closed.
func (a *Autoupdate) pruneTopic(closed <-chan struct{}) {
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()

	for {
		select {
		case <-closed:
			return
		case <-tick.C:
			a.topic.Prune(time.Now().Add(-pruneTime))
		}
	}
}

// resetCache runs in the background and cleans the cache from time to time.
// Blocks until the service is closed.
func (a *Autoupdate) resetCache(closed <-chan struct{}) {
	tick := time.NewTicker(datastoreCacheResetTime)
	defer tick.Stop()

	for {
		select {
		case <-closed:
			return
		case <-tick.C:
			a.datastore.ResetCache()
			// After the cache was updated, every connection has to be recalculated.
			a.topic.Publish(fmt.Sprintf(fullUpdateFormat, -1))
		}
	}
}
