// Package redis connects to a redis database to fetch database updates and
// logout events.
package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	// maxMessages desides how many messages are read at once from the stream.
	maxMessages = "10"

	// fieldChangedTopic is the redis key name of the autoupdate stream.
	fieldChangedTopic = "ModifiedFields"

	// logoutTopic is the redis key name of the logout stream.
	logoutTopic = "logout"

	// requestMetricKey is the key to save request metrics.
	requestMetricKey = "autoupdate_metric_request"

	// lastLogoutDuration decides how many old logout messages are received.
	lastLogoutDuration = 15 * time.Minute
)

// Connection is the raw connection to a redis server.
type Connection interface {
	XREAD(count, stream, lastID string) (interface{}, error)
	ZINCR(key string, value []byte) error
	ZRANGE(key string) (interface{}, error)
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

// RequestMeticSave saves how often a request was send.
func (r *Redis) RequestMeticSave(request []byte) error {
	normalized, err := normalizeRequest(request)
	if err != nil {
		return fmt.Errorf("normalize request: %w", err)
	}

	if err := r.Conn.ZINCR(requestMetricKey, normalized); err != nil {
		return fmt.Errorf("saving metric: %w", err)
	}

	return nil
}

// RequestMetricGet writes all request with there count as json.
func (r *Redis) RequestMetricGet(w io.Writer) error {
	values, err := redis.IntMap(r.Conn.ZRANGE(requestMetricKey))
	if err != nil {
		return fmt.Errorf("reading data: %w", err)
	}

	if err := json.NewEncoder(w).Encode(values); err != nil {
		return fmt.Errorf("encoding and sending data: %w", err)
	}

	return nil
}

// normalizeRequest takes json and returns the same output when the input has
// the same keys but in the different order.
func normalizeRequest(request []byte) ([]byte, error) {
	var bodies []struct {
		Collection string          `json:"collection"`
		Fields     json.RawMessage `json:"fields"`
	}
	err := json.Unmarshal(request, &bodies)
	if err != nil {
		return nil, err
	}
	output, err := json.Marshal(bodies)
	if err != nil {
		return nil, err
	}
	return output, nil
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
