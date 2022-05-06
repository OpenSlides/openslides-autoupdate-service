// Package autoupdate allows clients to request keys and get updates when the
// keys changes.
//
// To register to the autoupdate serive, a client has to receive a connection
// object by calling the Connect()-method. It is not necessary and therefore not
// possible to close a connection. The client can just stop listening.
package autoupdate

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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
	// response time for the clients.
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
	topic      *topic.Topic[string]
	restricter RestrictMiddleware
	voteAddr   string
}

// RestrictMiddleware is a function that can restrict data.
type RestrictMiddleware func(getter datastore.Getter, uid int) datastore.Getter

// New creates a new autoupdate service.
//
// The attribute closed is a channel that should be closed when the server shuts
// down. In this case, all connections get closed.
func New(datastore Datastore, restricter RestrictMiddleware, voteAddr string) *Autoupdate {
	a := &Autoupdate{
		datastore:  datastore,
		topic:      topic.New[string](),
		restricter: restricter,
		voteAddr:   voteAddr,
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

// SingleData returns the data for the kb. It is the same as calling Connect and
// then Next for the first time.
func (a *Autoupdate) SingleData(ctx context.Context, userID int, kb KeysBuilder, position int) (map[string][]byte, error) {
	var getter datastore.Getter = a.datastore
	var restricter datastore.Getter = a.restricter(getter, userID)

	if position != 0 {
		getter = datastore.NewGetPosition(a.datastore, position)
		restricter = restrict.NewHistory(userID, a.datastore, getter)
	}

	if err := kb.Update(ctx, restricter); err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := restricter.Get(ctx, kb.Keys()...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}

	var f filter
	f.filter(data)

	return data, nil
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

// HistoryInformation writes the history information for an fqid.
func (a *Autoupdate) HistoryInformation(ctx context.Context, uid int, fqid string, w io.Writer) error {
	coll, rawID, found := strings.Cut(fqid, "/")
	if !found {
		return fmt.Errorf("invalid fqid")
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		return fmt.Errorf("invalid fqid. ID part is not an int")
	}

	ds := datastore.NewRequest(a.datastore)

	meetingID, hasMeeting, err := collection.Collection(coll).MeetingID(ctx, ds, id)
	if err != nil {
		var errNotExist datastore.DoesNotExistError
		if errors.As(err, &errNotExist) {
			return notExistError{string(errNotExist)}
		}
		return fmt.Errorf("getting meeting id for collection %s id %d: %w", coll, id, err)
	}

	if !hasMeeting {
		hasOML, err := perm.HasOrganizationManagementLevel(ctx, ds, uid, perm.OMLCanManageOrganization)
		if err != nil {
			return fmt.Errorf("getting organization management level: %w", err)
		}

		if !hasOML {
			return permissionDeniedError{fmt.Errorf("you are not allowed to use history information on %s", fqid)}
		}
	} else {
		p, err := perm.New(ctx, ds, uid, meetingID)
		if err != nil {
			return fmt.Errorf("getting meeting permissions: %w", err)
		}

		if !p.Has(perm.MeetingCanSeeHistory) {
			return permissionDeniedError{fmt.Errorf("you are not allowed to use history information on %s", fqid)}
		}
	}

	if err := a.datastore.HistoryInformation(ctx, fqid, w); err != nil {
		return fmt.Errorf("getting history information: %w", err)
	}

	fmt.Fprintln(w)

	return nil
}

type permissionDeniedError struct {
	err error
}

func (e permissionDeniedError) Error() string {
	return fmt.Sprintf("permissoin denied: %v", e.err)
}

func (e permissionDeniedError) Type() string {
	return "permission_denied"
}

type notExistError struct {
	fqid string
}

func (e notExistError) Error() string {
	return fmt.Sprintf("%s does not exist", e.fqid)
}

func (e notExistError) Type() string {
	return "ont_exist"
}
