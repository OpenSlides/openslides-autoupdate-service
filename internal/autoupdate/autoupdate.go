// Package autoupdate allows clients to request keys and get updates when the
// keys changes.
//
// To register to the autoupdate serive, a client has to receive a connection
// object by calling the Connect()-method. It is not necessary and therefore not
// possible to close a connection. The client can just stop listening.
package autoupdate

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ostcar/topic"
)

// pruneTime defines how long a topic id will be valid. If a client needs more
// time to process the data, it will get an error and has to reconnect. A higher
// value means, that more memory is used.
const pruneTime = 10 * time.Minute

// Autoupdate holds the state of the autoupdate service. It has to be initialized
// with autoupdate.New().
type Autoupdate struct {
	datastore  Datastore
	restricter Restricter
	topic      *topic.Topic
}

// New creates a new autoupdate service.
func New(datastore Datastore, restricter Restricter, closed <-chan struct{}) *Autoupdate {
	a := &Autoupdate{
		datastore:  datastore,
		restricter: restricter,
		topic:      topic.New(topic.WithClosed(closed)),
	}

	// Update the topic when an data update is received.
	a.datastore.RegisterChangeListener(func(data map[string]json.RawMessage) error {
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}
		a.topic.Publish(keys...)
		return nil
	})

	go a.pruneTopic(closed)

	return a
}

// Connect has to be called by a client to register to the service. The method
// returns a Connection object, that can be used to receive the data.
//
// There is no need to "close" the Connection object.
func (a *Autoupdate) Connect(userID int, kb KeysBuilder, tid uint64) *Connection {
	return &Connection{
		autoupdate: a,
		uid:        userID,
		kb:         kb,
		tid:        tid,
	}
}

// LastID returns the last id of the last data update.
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

// RestrictedData returns a map containing the restricted values for the given
// keys. If a key does not exist or the user has not the permission to see it,
// the value in the returned map is nil.
func (a *Autoupdate) RestrictedData(ctx context.Context, uid int, keys ...string) (map[string]json.RawMessage, error) {
	values, err := a.datastore.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("get values for keys `%v` from datastore: %w", keys, err)
	}

	data := make(map[string]json.RawMessage, len(keys))
	for i, key := range keys {
		data[key] = values[i]
	}

	if err := a.restricter.Restrict(uid, data); err != nil {
		return nil, fmt.Errorf("restrict data: %w", err)
	}
	return data, nil
}
