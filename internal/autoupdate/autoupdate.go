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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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
	voteAddr   string
}

// RestrictMiddleware is a function that can restrict data.
type RestrictMiddleware func(getter datastore.Getter, uid int) datastore.Getter

// New creates a new autoupdate service.
//
// The attribute closed is a channel that should be closed when the server shuts
// down. In this case, all connections get closed.
func New(datastore Datastore, restricter RestrictMiddleware, voteAddr string, closed <-chan struct{}) *Autoupdate {
	a := &Autoupdate{
		datastore:  datastore,
		topic:      topic.New(topic.WithClosed(closed)),
		restricter: restricter,
		voteAddr:   voteAddr,
	}

	// Update the topic when an data update is received.
	a.datastore.RegisterChangeListener(func(ctx context.Context, data map[string][]byte) error {
		keys := make([]string, 0, len(data)+1)
		for k := range data {
			keys = append(keys, k)
		}

		_, span := otel.Tracer("autoupdate").Start(ctx, "autoupdate data update")
		defer span.End()

		kv := make([]string, 0, len(data))
		for k, v := range data {
			kv = append(kv, fmt.Sprintf("%s: %s", k, v))
		}
		span.SetAttributes(attribute.StringSlice("data", kv))

		keys = append(keys, "span:"+encodeSpanContext(span.SpanContext()))

		a.topic.Publish(keys...)
		return nil
	})

	// Register the calculated field for vote_count.
	a.datastore.RegisterCalculatedField("poll/vote_count", a.datastorePollVoteCount)

	return a
}

// DataProvider is a function that returns the next data for a user.
type DataProvider func(ctx context.Context) (map[string][]byte, error)

// Connect has to be called by a client to register to the service. The method
// returns a Connection object, that can be used to receive the data.
//
// There is no need to "close" the Connection object.
func (a *Autoupdate) Connect(userID int, kb KeysBuilder) DataProvider {
	c := &connection{
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

// PruneOldData removes old data from the topic. Blocks until the service is
// closed.
func (a *Autoupdate) PruneOldData(ctx context.Context) {
	tick := time.NewTicker(time.Minute)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			a.topic.Prune(time.Now().Add(-pruneTime))
		}
	}
}

// ResetCache runs in the background and cleans the cache from time to time.
// Blocks until the service is closed.
func (a *Autoupdate) ResetCache(ctx context.Context) {
	tick := time.NewTicker(datastoreCacheResetTime)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			a.datastore.ResetCache()
			// After the cache was updated, every connection has to be recalculated.
			a.topic.Publish(fmt.Sprintf(fullUpdateFormat, -1))
		}
	}
}
