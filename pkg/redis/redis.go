// Package redis connects to a redis database to fetch database updates and
// logout events.
package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

const (
	// maxMessages desides how many messages are read at once from the stream.
	maxMessages = "10"

	// fieldChangedTopic is the redis key name of the autoupdate stream.
	fieldChangedTopic = "ModifiedFields"

	// logoutTopic is the redis key name of the logout stream.
	logoutTopic = "logout"

	// lastLogoutDuration decides how many old logout messages are received.
	lastLogoutDuration = 15 * time.Minute
)

// Connection is the raw connection to a redis server.
type Connection interface {
	XREAD(count, stream, lastID string) (interface{}, error)
}

// Redis holds the state of the redis receiver.
type Redis struct {
	Conn             Connection
	lastAutoupdateID string
	lastLogoutID     string
}

// Update is a blocking function that returns, when there is new data.
func (r *Redis) Update(ctx context.Context) (map[string][]byte, error) {
	id := r.lastAutoupdateID
	if id == "" {
		id = "$"
	}

	var data map[string][]byte
	err := contextFunc(ctx, func() error {
		newID, d, err := autoupdateStream(r.Conn.XREAD(maxMessages, fieldChangedTopic, id))
		if err != nil {
			return err
		}
		id = newID
		data = d
		return nil
	})

	if err != nil {
		if err == errNil {
			// No new data
			return nil, nil
		}
		return nil, fmt.Errorf("get xread data from redis: %w", err)
	}
	if id != "" {
		r.lastAutoupdateID = id
	}
	return data, nil
}

// LogoutEvent is a blocking function that returns, when a session was revoked.
func (r *Redis) LogoutEvent(ctx context.Context) ([]string, error) {
	id := r.lastLogoutID
	if id == "" {
		// Generate an redis ID to get the logout events from the since `lastLogoutDuration`.
		id = strconv.FormatInt(time.Now().Add(-lastLogoutDuration).Unix(), 10)
	}

	var sessionIDs []string
	err := contextFunc(ctx, func() error {
		newID, sIDs, err := logoutStream(r.Conn.XREAD(maxMessages, logoutTopic, id))
		if err != nil {
			return err
		}
		id = newID
		sessionIDs = sIDs
		return nil
	})

	if err != nil {
		if err == errNil {
			// No new data
			return nil, nil
		}
		return nil, fmt.Errorf("get xread data from redis: %w", err)
	}
	if id != "" {
		r.lastLogoutID = id
	}
	return sessionIDs, nil
}

// contextFunc calls f in a separat goroutine. If the given context is done,
// return and leave the goroutine behind.
//
// This is usefull if a blocking function does not support context and has no
// other way to cancel it.
//
// The returned error is either the error returned by f() or the context.
func contextFunc(ctx context.Context, f func() error) error {
	done := make(chan error)
	go func() {
		done <- f()
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
