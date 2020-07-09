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
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ostcar/topic"
)

// pruneTime defines how long a topic id will be valid. If a client needs more
// time to process the data, it will get an error and has to reconnect. A higher
// value means, that more memory is used.
const pruneTime = time.Minute

// Autoupdate holds the state of the autoupdate service. It has to be initialized
// with autoupdate.New().
//
// The service updates its data in the background. To stop this background job,
// the service has to be closed in the end with the Close()-method.
type Autoupdate struct {
	datastore  Datastore
	restricter Restricter
	closed     chan struct{}
	topic      *topic.Topic
}

// New creates a new autoupdate service.
//
// After the service is not needed anymore, it has to be closed with s.Close().
func New(datastore Datastore, restricter Restricter) *Autoupdate {
	s := &Autoupdate{
		datastore:  datastore,
		restricter: restricter,
		closed:     make(chan struct{}),
	}
	s.topic = topic.New(topic.WithClosed(s.closed))

	go s.receiveKeyChanges()
	go s.pruneTopic()

	return s
}

// Close calls the shutdown logic of the service. This method is not save for
// concourent use. It not allowed to call it more then once. Only the caller of
// New() should call Close().
func (a *Autoupdate) Close() {
	close(a.closed)
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

// Value decodes the restricted value for the given key.
func (a *Autoupdate) Value(ctx context.Context, uid int, key string, value interface{}) error {
	data, err := a.restrictedData(ctx, uid, key)
	if err != nil {
		return fmt.Errorf("get restricted value for key %s: %w", key, err)
	}

	if len(data[key]) == 0 {
		// No value for key.
		return NotExistError{Key: key}
	}

	if err := json.Unmarshal(data[key], value); err != nil {
		var invalidErr *json.UnmarshalTypeError
		if match := errors.As(err, &invalidErr); match {
			// value has wrong type.
			return ValueError{key: key, err: err}
		}
		return fmt.Errorf("decode value of key %s: %w", key, err)
	}
	return nil
}

// LastID returns the last id of the last data update.
func (a *Autoupdate) LastID() uint64 {
	return a.topic.LastID()
}

// pruneTopic removes old data from the topic. Blocks until the service is
// closed.
func (a *Autoupdate) pruneTopic() {
	tick := time.NewTicker(time.Second)
	defer tick.Stop()

	for {
		select {
		case <-a.closed:
			return
		case <-tick.C:
			a.topic.Prune(time.Now().Add(-pruneTime))
		}
	}
}

// receiveKeyChanges listens for updates and saves then into the topic. This
// function blocks until the service is closed.
func (a *Autoupdate) receiveKeyChanges() {
	for {
		select {
		case <-a.closed:
			return
		default:
		}

		keys, err := a.datastore.KeysChanged()
		if err != nil {
			log.Printf("Could not update keys: %v\n", err)
			time.Sleep(time.Second)
			continue
		}

		a.topic.Publish(keys...)
	}
}

// restrictedData returns a map containing the restricted values for the given
// keys.
func (a *Autoupdate) restrictedData(ctx context.Context, uid int, keys ...string) (map[string]json.RawMessage, error) {
	values, err := a.datastore.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("get values for keys `%v` from datastore: %w", keys, err)
	}

	data := make(map[string]json.RawMessage, len(keys))
	for i, key := range keys {
		data[key] = values[i]
	}

	a.restricter.Restrict(uid, data)
	return data, nil
}
