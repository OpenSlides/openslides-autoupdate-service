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
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
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

var envConcurentWorker = environment.NewVariable("CONCURENT_WORKER", "0", "Amount of clients that calculate there values at the same time. Default to GOMAXPROCS.")

// Datastore is the source for the data.
type Datastore interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
	GetPosition(ctx context.Context, position int, keys ...dskey.Key) (map[dskey.Key][]byte, error)
	RegisterChangeListener(f func(map[dskey.Key][]byte) error)
	ResetCache()
	RegisterCalculatedField(
		field string,
		f func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, error),
	)
	HistoryInformation(ctx context.Context, fqid string, w io.Writer) error
}

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update(ctx context.Context, ds datastore.Getter) ([]dskey.Key, error)
}

// RestrictMiddleware is a function that can restrict data.
type RestrictMiddleware func(ctx context.Context, getter datastore.Getter, uid int) (context.Context, datastore.Getter)

// Autoupdate holds the state of the autoupdate service. It has to be initialized
// with autoupdate.New().
type Autoupdate struct {
	datastore  Datastore
	topic      *topic.Topic[dskey.Key]
	restricter RestrictMiddleware
	pool       *workPool
}

// New creates a new autoupdate service.
//
// You should call `go a.PruneOldData()` and `go a.ResetCache()` after creating
// the service.
func New(lookup environment.Environmenter, ds Datastore, restricter RestrictMiddleware) (*Autoupdate, func(context.Context, func(error)), error) {
	workers, err := strconv.Atoi(envConcurentWorker.Value(lookup))
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for %s: %w", envConcurentWorker.Key, err)
	}

	if workers == 0 {
		workers = runtime.GOMAXPROCS(0)
	}

	a := &Autoupdate{
		datastore:  ds,
		topic:      topic.New[dskey.Key](),
		restricter: restricter,
		pool:       newWorkPool(workers),
	}

	// Update the topic when an data update is received.
	a.datastore.RegisterChangeListener(func(data map[dskey.Key][]byte) error {
		keys := make([]dskey.Key, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}

		a.topic.Publish(keys...)
		return nil
	})

	background := func(ctx context.Context, errorHandler func(error)) {
		go a.pruneOldData(ctx)
		go a.resetCache(ctx)
	}

	return a, background, nil
}

// DataProvider is a function that returns the next data for a user.
type DataProvider func() (func(ctx context.Context) (map[dskey.Key][]byte, error), bool)

// Connect has to be called by a client to register to the service. The method
// returns a Connection object, that can be used to receive the data.
//
// There is no need to "close" the returned DataProvider.
func (a *Autoupdate) Connect(ctx context.Context, userID int, kb KeysBuilder) (DataProvider, error) {
	skipWorkpool, err := a.skipWorkpool(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("check if workpool should be used: %w", err)
	}

	c := &connection{
		autoupdate:   a,
		uid:          userID,
		kb:           kb,
		skipWorkpool: skipWorkpool,
	}

	return c.Next, nil
}

// SingleData returns the data for the given keysbuilder without autoupdates.
//
// The attribute position can be used to get data from the history.
func (a *Autoupdate) SingleData(ctx context.Context, userID int, kb KeysBuilder, position int) (map[dskey.Key][]byte, error) {
	var restricter datastore.Getter

	ctx, restricter = a.restricter(ctx, a.datastore, userID)

	if position != 0 {
		getter := datastore.NewGetPosition(a.datastore, position)
		restricter = restrict.NewHistory(a.datastore, getter, userID)
	}

	keys, err := kb.Update(ctx, restricter)
	if err != nil {
		return nil, fmt.Errorf("create keys for keysbuilder: %w", err)
	}

	data, err := restricter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("get restricted data: %w", err)
	}

	for k, v := range data {
		if len(v) == 0 {
			delete(data, k)
		}
	}

	return data, nil
}

// pruneOldData removes old data from the topic. Blocks until the service is
// closed.
func (a *Autoupdate) pruneOldData(ctx context.Context) {
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

// resetCache runs in the background and cleans the cache from time to time.
// Blocks until the service is closed.
func (a *Autoupdate) resetCache(ctx context.Context) {
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

	meetingID, hasMeeting, err := collection.Collection(ctx, coll).MeetingID(ctx, ds, id)
	if err != nil {
		var errNotExist dsfetch.DoesNotExistError
		if errors.As(err, &errNotExist) {
			// TODO Client Error
			return notExistError{dskey.Key(errNotExist)}
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

// RestrictFQIDs returns the full collections, restricted for the user for a
// list of fqids.
// In requestedFields one can specify which fields per collection should be
// returned if not specified all available fields will be included.
//
// The return format is a map from fqid to an object as map from field to value.
func (a *Autoupdate) RestrictFQIDs(ctx context.Context, userID int, fqids []string, requestedFields map[string][]string) (map[string]map[string][]byte, error) {
	requestedFieldsMap := make(map[string]set.Set[string], len(requestedFields))
	for col, val := range requestedFields {
		requestedFieldsMap[col] = set.New(val...)
	}

	var keys []dskey.Key
	for _, fqid := range fqids {
		collection, rawID, found := strings.Cut(fqid, "/")
		if !found {
			return nil, fmt.Errorf("invalid fqid %s, expected one /", fqid)
		}

		id, err := strconv.Atoi(rawID)
		if err != nil {
			return nil, fmt.Errorf("invalid fqid %s, second part has to be an nummber", fqid)
		}

		fields := restrict.FieldsForCollection(collection)
		if fields == nil {
			return nil, fmt.Errorf("unknown collection in fqid %s", fqid)
		}

		for _, field := range fields {
			if _, ok := requestedFields[collection]; !ok || requestedFieldsMap[collection].Has(field) {
				key := dskey.Key{Collection: collection, ID: id, Field: field}
				keys = append(keys, key)
			}
		}
	}

	ctx, restricter := a.restricter(ctx, a.datastore, userID)

	values, err := restricter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("getting data: %w", err)
	}

	result := make(map[string]map[string][]byte, len(fqids))
	for key, value := range values {
		fqid := key.FQID()
		if _, ok := result[fqid]; !ok {
			result[fqid] = make(map[string][]byte)
		}

		if value == nil {
			continue
		}

		result[fqid][key.Field] = value
	}

	return result, nil
}

// skipWorkpool desides, if a connection is allowed to skip the workpool.
//
// The current implementation returns true, if the user is a meeting admin in
// one meeting.
func (a *Autoupdate) skipWorkpool(ctx context.Context, userID int) (bool, error) {
	if userID == 0 {
		return false, nil
	}

	ds := dsfetch.New(a.datastore)

	meetingIDs := ds.User_GroupIDsTmpl(userID).ErrorLater(ctx)

	for _, mid := range meetingIDs {
		gids := ds.User_GroupIDs(userID, mid).ErrorLater(ctx)
		for _, gid := range gids {
			if _, ok := ds.Group_AdminGroupForMeetingID(gid).ErrorLater(ctx); ok {
				return true, nil
			}
		}
	}
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("check if user %d is a meeting admin: %w", userID, err)
	}

	return false, nil
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
	key dskey.Key
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
