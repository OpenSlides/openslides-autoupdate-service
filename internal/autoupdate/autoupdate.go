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
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
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

// KeysBuilder holds the keys that are requested by a user.
type KeysBuilder interface {
	Update(ctx context.Context, ds flow.Getter) ([]dskey.Key, error)
}

// RestrictMiddleware is a function that can restrict data.
type RestrictMiddleware func(ctx context.Context, getter flow.Getter, uid int) (context.Context, flow.Getter)

// Autoupdate holds the state of the autoupdate service. It has to be initialized
// with autoupdate.New().
type Autoupdate struct {
	flow       flow.Flow
	topic      *topic.Topic[dskey.Key]
	restricter RestrictMiddleware
	pool       *workPool

	cacheReset time.Duration
}

// New creates a new autoupdate service.
//
// You should call `go a.PruneOldData()` and `go a.ResetCache()` after creating
// the service.
func New(lookup environment.Environmenter, flow flow.Flow, restricter RestrictMiddleware) (*Autoupdate, func(context.Context, func(error)), error) {
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
		flow:       flow,
		topic:      topic.New[dskey.Key](),
		restricter: restricter,
		pool:       newWorkPool(workers),
		cacheReset: cacheResetTime,
	}

	background := func(ctx context.Context, errorHandler func(error)) {
		go a.pruneOldData(ctx)
		go a.resetCache(ctx)
		go a.flow.Update(ctx, func(data map[dskey.Key][]byte, err error) {
			if err != nil {
				oserror.Handle(err)
				// Continue. The update function can return an error and data.
			}

			keys := make([]dskey.Key, 0, len(data))
			for k := range data {
				keys = append(keys, k)
			}

			a.topic.Publish(keys...)
		})
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
func (a *Autoupdate) SingleData(ctx context.Context, userID int, kb KeysBuilder) (map[dskey.Key][]byte, error) {
	ctx, restricter := a.restricter(ctx, a.flow, userID)

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
	type resetter interface {
		ResetCache()
	}
	reset, ok := a.flow.(resetter)
	if !ok {
		return
	}

	tick := time.NewTicker(a.cacheReset)
	defer tick.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-tick.C:
			reset.ResetCache()
		}
	}
}

// skipWorkpool desides, if a connection is allowed to skip the workpool.
//
// The current implementation returns true, if the user is a meeting admin in
// one meeting.
func (a *Autoupdate) skipWorkpool(ctx context.Context, userID int) (bool, error) {
	if userID == 0 {
		return false, nil
	}

	ds := dsfetch.New(a.flow)

	meetingUserIDs, err := ds.User_MeetingUserIDs(userID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting meeting_user objects: %w", err)
	}

	for _, muid := range meetingUserIDs {
		groupIDs := ds.MeetingUser_GroupIDs(muid).ErrorLater(ctx)
		for _, gid := range groupIDs {
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

	ds := dsfetch.New(a.flow)

	hasOML, err := perm.HasOrganizationManagementLevel(ctx, ds, userID, perm.OMLCanManageOrganization)
	if err != nil {
		return false, fmt.Errorf("getting organization management level: %w", err)
	}

	return hasOML, nil
}
