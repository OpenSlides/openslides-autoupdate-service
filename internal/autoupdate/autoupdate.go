// Package autoupdate allows clients to request keys and get updates when the
// keys changes.
//
// To register to the autoupdate serive, a client has to call Connect() to get a
// connection object. It is not necessary and therefore not possible to close
// the connection.
package autoupdate

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
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
)

var (
	envConcurentWorker = environment.NewVariable("CONCURENT_WORKER", "0", "Amount of clients that calculate there values at the same time. Default to GOMAXPROCS.")
	envCacheReset      = environment.NewVariable("CACHE_RESET", "24h", "Time to reset the cache.")
)

// Datastore is the source for the data.
type Datastore interface {
	Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error)
	RegisterChangeListener(f func(map[dskey.Key][]byte) error)
	ResetCache()
	RegisterCalculatedField(
		field string,
		f func(ctx context.Context, key dskey.Key, changed map[dskey.Key][]byte) ([]byte, bool, error),
	)
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

	cacheReset time.Duration
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

	cacheResetTime, err := environment.ParseDuration(envCacheReset.Value(lookup))
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for `CACHE_RESET`, expected duration got %s: %w", envCacheReset.Value(lookup), err)
	}

	a := &Autoupdate{
		datastore:  ds,
		topic:      topic.New[dskey.Key](),
		restricter: restricter,
		pool:       newWorkPool(workers),
		cacheReset: cacheResetTime,
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
func (a *Autoupdate) SingleData(ctx context.Context, userID int, kb KeysBuilder) (map[dskey.Key][]byte, error) {
	ctx, restricter := a.restricter(ctx, a.datastore, userID)

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
	tick := time.NewTicker(a.cacheReset)
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

// CanSeeConnectionCount returns, if the user can see the connection counter.
func (a *Autoupdate) CanSeeConnectionCount(ctx context.Context, userID int) (bool, error) {
	if userID == 0 {
		return false, nil
	}

	ds := dsfetch.New(a.datastore)

	hasOML, err := perm.HasOrganizationManagementLevel(ctx, ds, userID, perm.OMLCanManageOrganization)
	if err != nil {
		return false, fmt.Errorf("getting organization management level: %w", err)
	}

	return hasOML, nil
}
