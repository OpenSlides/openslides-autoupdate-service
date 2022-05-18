// Package redis connects to a redis database to fetch database updates and
// logout events.
package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
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
	XREAD(ctx context.Context, count, stream, lastID string) (interface{}, error)
}

// Redis holds the state of the redis receiver.
type Redis struct {
	Conn             Connection
	lastAutoupdateID string
	lastLogoutID     string
}

// Update is a blocking function that returns, when there is new data.
func (r *Redis) Update(ctx context.Context) (map[datastore.Key][]byte, error) {
	id := r.lastAutoupdateID
	if id == "" {
		id = "$"
	}

	id, data, err := autoupdateStream(r.Conn.XREAD(ctx, maxMessages, fieldChangedTopic, id))
	if err != nil {
		if err == errNil {
			// No new data
			return nil, nil
		}
		// TODO External Error
		return nil, fmt.Errorf("get xread data from redis: %w", err)
	}

	if id != "" {
		r.lastAutoupdateID = id
	}

	converted := make(map[datastore.Key][]byte, len(data))
	for k, v := range data {
		key, err := datastore.KeyFromString(k)
		if err != nil {
			// TODO End Error
			return nil, fmt.Errorf("invalid key: %s", k)
		}
		converted[key] = v
	}

	return converted, nil
}

// LogoutEvent is a blocking function that returns, when a session was revoked.
func (r *Redis) LogoutEvent(ctx context.Context) ([]string, error) {
	id := r.lastLogoutID
	if id == "" {
		// Generate an redis ID to get the logout events from the since `lastLogoutDuration`.
		id = strconv.FormatInt(time.Now().Add(-lastLogoutDuration).Unix(), 10)
	}

	id, sessionIDs, err := logoutStream(r.Conn.XREAD(ctx, maxMessages, logoutTopic, id))
	if err != nil {
		if err == errNil {
			// No new data
			return nil, nil
		}
		// TODO External Error
		return nil, fmt.Errorf("get xread data from redis: %w", err)
	}
	if id != "" {
		r.lastLogoutID = id
	}
	return sessionIDs, nil
}
