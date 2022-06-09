// Package autoupdate allows clients to request keys and get updates when the
// keys changes.
//
// To register to the autoupdate serive, a client has to call Connect() to get a
// connection object. It is not necessary and therefore not possible to close
// the connection.
package autoupdate

import (
	"context"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/ostcar/topic"
)

const (
	// pruneTime defines how long the data in the topic will be valid. If a
	// client needs more time to process the data, it will get an error and has
	// to reconnect. A higher value means, that more memory is used.
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

// Datastore is the source for the data.
type Datastore interface {
	Get(ctx context.Context, keys ...datastore.Key) (map[datastore.Key][]byte, error)
	GetPosition(ctx context.Context, position int, keys ...datastore.Key) (map[datastore.Key][]byte, error)
	RegisterChangeListener(f func(map[datastore.Key][]byte) error)
	ResetCache()
	RegisterCalculatedField(
		field string,
		f func(ctx context.Context, key datastore.Key, changed map[datastore.Key][]byte) ([]byte, error),
	)
	HistoryInformation(ctx context.Context, fqid string, w io.Writer) error
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update(ctx context.Context, ds datastore.Getter) error
	Keys() []datastore.Key
}

// RestrictMiddleware is a function that can restrict data.
type RestrictMiddleware func(getter datastore.Getter, uid int) datastore.Getter

// Autoupdate holds the state of the autoupdate service. It has to be initialized
// with autoupdate.New().
type Autoupdate struct {
	datastore  Datastore
	topic      *topic.Topic[datastore.Key]
	restricter RestrictMiddleware
}

// New creates a new autoupdate service.
//
// You should call `go a.PruneOldData()` and `go a.ResetCache()` after creating
// the service.
func New(ds Datastore, restricter RestrictMiddleware) *Autoupdate {
	a := &Autoupdate{
		datastore:  ds,
		topic:      topic.New[datastore.Key](),
		restricter: restricter,
	}

	// Update the topic when an data update is received.
	a.datastore.RegisterChangeListener(func(data map[datastore.Key][]byte) error {
		keys := make([]datastore.Key, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}

		a.topic.Publish(keys...)
		return nil
	})

	return a
}

// DataProvider is a function that returns the next data for a user.
type DataProvider func(ctx context.Context) (map[datastore.Key][]byte, error)

// Connect has to be called by a client to register to the service. The method
// returns a Connection object, that can be used to receive the data.
//
// There is no need to "close" the returned DataProvider.
func (a *Autoupdate) Connect(userID int, kb KeysBuilder) DataProvider {
	c := &connection{
		autoupdate: a,
		uid:        userID,
		kb:         kb,
	}

	return c.Next
}

// SingleData returns the data for the given keysbuilder without autoupdates.
//
// The attribute position can be used to get data from the history.
func (a *Autoupdate) SingleData(ctx context.Context, userID int, kb KeysBuilder, position int) (map[datastore.Key][]byte, error) {
	var restricter datastore.Getter = a.restricter(a.datastore, userID)

	//restricter = datastore.NewRecorder(restricter)

	if position != 0 {
		getter := datastore.NewGetPosition(a.datastore, position)
		restricter = restrict.NewHistory(a.datastore, getter, userID)
	}

	if err := kb.Update(ctx, restricter); err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := restricter.Get(ctx, kb.Keys()...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}

	// recorder := restricter.(*datastore.Recorder)
	// db, err := recorder.DB()
	// if err != nil {
	// 	return nil, fmt.Errorf("creating db: %w", err)
	// }
	// fmt.Println(string(db))

	for k, v := range data {
		if len(v) == 0 {
			delete(data, k)
		}
	}

	return data, nil
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
		}
	}
}

var reValidKeys = regexp.MustCompile(`^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*`)

// HistoryInformation writes the history information for an fqid.
func (a *Autoupdate) HistoryInformation(ctx context.Context, uid int, fqid string, w io.Writer) error {
	if !reValidKeys.MatchString(fqid) {
		// TODO Client Error
		return invalidInputError{fmt.Sprintf("fqid %s is invalid", fqid)}
	}

	coll, rawID, _ := strings.Cut(fqid, "/")
	id, _ := strconv.Atoi(rawID)

	ds := dsfetch.New(a.datastore)

	meetingID, hasMeeting, err := collection.Collection(coll).MeetingID(ctx, ds, id)
	if err != nil {
		var errNotExist dsfetch.DoesNotExistError
		if errors.As(err, &errNotExist) {
			// TODO Client Error
			return notExistError{datastore.Key(errNotExist)}
		}
		return fmt.Errorf("getting meeting id for collection %s id %d: %w", coll, id, err)
	}

	if !hasMeeting {
		hasOML, err := perm.HasOrganizationManagementLevel(ctx, ds, uid, perm.OMLCanManageOrganization)
		if err != nil {
			return fmt.Errorf("getting organization management level: %w", err)
		}

		if !hasOML {
			// TODO Client Error
			return permissionDeniedError{fmt.Errorf("you are not allowed to use history information on %s", fqid)}
		}
	} else {
		p, err := perm.New(ctx, ds, uid, meetingID)
		if err != nil {
			return fmt.Errorf("getting meeting permissions: %w", err)
		}

		if !p.Has(perm.MeetingCanSeeHistory) {
			// TODO Client Error
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
	key datastore.Key
}

func (e notExistError) Error() string {
	return fmt.Sprintf("%s does not exist", e.key)
}

func (e notExistError) Type() string {
	return "not_exist"
}

type invalidInputError struct {
	msg string
}

func (e invalidInputError) Error() string {
	return e.msg
}

func (e invalidInputError) Type() string {
	return "invalid_input"
}
