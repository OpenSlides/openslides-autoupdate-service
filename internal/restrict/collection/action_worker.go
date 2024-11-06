package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// ActionWorker handles permission for action_worker.
//
// A user can see an action worker, if he is the user from action_worker/user_id.
type ActionWorker struct{}

// Name returns the collection name.
func (a ActionWorker) Name() string {
	return "action_worker"
}

// MeetingID returns no meeting.
func (a ActionWorker) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the restrictions modes for the action_worker collection.
func (a ActionWorker) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	}
	return nil
}

func (a ActionWorker) see(ctx context.Context, ds *dsfetch.Fetch, actionWorkerIDs ...int) ([]int, error) {
	userIDs := make([]int, len(actionWorkerIDs))
	for i, id := range actionWorkerIDs {
		ds.ActionWorker_UserID(id).Lazy(&userIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting user ids: %w", err)
	}

	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	var allowed []int
	for i, id := range actionWorkerIDs {
		if userIDs[i] == requestUser {
			allowed = append(allowed, id)
		}
	}

	return allowed, nil

}
