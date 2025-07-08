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

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/environment"
	"github.com/OpenSlides/openslides-go/oserror"
	"github.com/OpenSlides/openslides-go/perm"
	"github.com/OpenSlides/openslides-go/set"
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

// Connect has to be called by a client to register to the service. The method
// returns a Connection object, that can be used to receive the data.
//
// There is no need to "close" the returned DataProvider.
func (a *Autoupdate) Connect(ctx context.Context, userID int, kb KeysBuilder) (Connection, error) {
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

	return c, nil
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
		groupIDs, err := ds.MeetingUser_GroupIDs(muid).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting groupIDs of user: %w", err)
		}

		adminGroups := make([]dsfetch.Maybe[int], len(groupIDs))
		for i := 0; i < len(groupIDs); i++ {
			ds.Group_AdminGroupForMeetingID(groupIDs[i]).Lazy(&adminGroups[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return false, fmt.Errorf("checking for admin groups: %w", err)
		}

		for _, isAdmin := range adminGroups {
			if _, ok := isAdmin.Value(); ok {
				return true, nil
			}
		}

	}

	return false, nil
}

// CanSeeConnectionCount returns, if the user can see the connection count.
//
// If the second value is not empty, the user is only allowed for meetings in that list.
func (a *Autoupdate) CanSeeConnectionCount(ctx context.Context, userID int) (bool, []int, error) {
	if userID == 0 {
		return false, nil, nil
	}

	ds := dsfetch.New(a.flow)

	hasOML, err := perm.HasOrganizationManagementLevel(ctx, ds, userID, perm.OMLCanManageOrganization)
	if err != nil {
		return false, nil, fmt.Errorf("getting organization management level: %w", err)
	}

	if hasOML {
		return true, nil, nil
	}

	meetingUserIDs, err := ds.User_MeetingUserIDs(userID).Value(ctx)
	if err != nil {
		return false, nil, fmt.Errorf("getting meeting_user objects: %w", err)
	}

	groupPerMeetingIDs := make([][]int, len(meetingUserIDs))
	for i, muID := range meetingUserIDs {
		ds.MeetingUser_GroupIDs(muID).Lazy(&groupPerMeetingIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return false, nil, fmt.Errorf("getting all group ids: %w", err)
	}

	var groupIDs []int
	for i := range groupPerMeetingIDs {
		groupIDs = append(groupIDs, groupPerMeetingIDs[i]...)
	}

	isAdminGroup := make([]dsfetch.Maybe[int], len(groupIDs))
	for i, groupID := range groupIDs {
		ds.Group_AdminGroupForMeetingID(groupID).Lazy(&isAdminGroup[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return false, nil, fmt.Errorf("getting admin flag of groups: %w", err)
	}

	var meetingAdmin []int
	for _, meetingID := range isAdminGroup {
		if id, ok := meetingID.Value(); ok {
			meetingAdmin = append(meetingAdmin, id)
		}
	}

	return len(meetingAdmin) > 0, meetingAdmin, nil
}

// FilterConnectionCount removes users from the count where the user is not in
// one of the given meetings.
func (a *Autoupdate) FilterConnectionCount(ctx context.Context, meetingIDs []int, count map[int]int) error {
	if len(meetingIDs) == 0 {
		return nil
	}

	ds := dsfetch.New(a.flow)

	meetingUserIDs := make([][]int, len(meetingIDs))
	for i, meetingID := range meetingIDs {
		ds.Meeting_MeetingUserIDs(meetingID).Lazy(&meetingUserIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return fmt.Errorf("getting meeting user ids: %w", err)
	}

	userCount := 0
	userIDs := make([][]int, len(meetingIDs))
	for i := range meetingUserIDs {
		userCount += len(meetingUserIDs[i])
		userIDs[i] = make([]int, len(meetingUserIDs[i]))
		for j, muID := range meetingUserIDs[i] {
			ds.MeetingUser_UserID(muID).Lazy(&userIDs[i][j])
		}
	}

	if err := ds.Execute(ctx); err != nil {
		return fmt.Errorf("getting user ids: %w", err)
	}

	userSet := set.NewWithSize[int](userCount)
	for i := range userIDs {
		userSet.Add(userIDs[i]...)
	}

	for userID := range count {
		if !userSet.Has(userID) {
			delete(count, userID)
		}
	}

	return nil
}

var reValidKeys = regexp.MustCompile(`^([a-z]+|[a-z][a-z_]*[a-z])/[1-9][0-9]*`)

// HistoryInformation returns the histrory information for an fqid.
func (a *Autoupdate) HistoryInformation(ctx context.Context, uid int, fqid string, w io.Writer) error {
	type History interface {
		historyInformation(ctx context.Context, fqid string, w io.Writer) error
	}
	hi, ok := a.flow.(History)
	if !ok {
		return fmt.Errorf("history not supported")
	}

	if !reValidKeys.MatchString(fqid) {
		// TODO Client Error
		return invalidInputError{fmt.Sprintf("fqid %s is invalid", fqid)}
	}

	coll, rawID, _ := strings.Cut(fqid, "/")
	id, _ := strconv.Atoi(rawID)

	ds := dsfetch.New(a.flow)

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

	if err := hi.historyInformation(ctx, fqid, w); err != nil {
		return fmt.Errorf("getting history information: %w", err)
	}

	fmt.Fprintln(w)

	return nil
}
