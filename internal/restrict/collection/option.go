package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Option handels restrictions of the collection option.
type Option struct{}

// Modes returns the restrictions modes for the meeting collection.
func (o Option) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return o.see
	case "B":
		return o.modeB
	}
	return nil
}

func (o Option) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, optionID int) (bool, error) {
	pollID, err := pollID(ctx, ds, optionID)
	if err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}

	see, err := Poll{}.see(ctx, ds, mperms, pollID)
	if err != nil {
		return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
	}

	return see, nil
}

func (o Option) modeB(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, optionID int) (bool, error) {
	pollID, err := pollID(ctx, ds, optionID)
	if err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}

	see, err := Poll{}.see(ctx, ds, mperms, pollID)
	if err != nil {
		return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
	}

	if !see {
		return false, nil
	}

	canManage, err := Poll{}.manage(ctx, ds, mperms, pollID)
	if err != nil {
		return false, fmt.Errorf("checking see poll %d: %w", pollID, err)
	}

	if canManage {
		return true, nil
	}

	pollState, err := ds.Poll_State(pollID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting poll state: %w", err)
	}

	return pollState == "published", nil
}

func pollID(ctx context.Context, ds *datastore.Request, optionID int) (int, error) {
	pollID, exist, err := ds.Option_PollID(optionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting poll id from field poll_id: %w", err)
	}

	if exist {
		return pollID, nil
	}

	pollID, exist, err = ds.Option_UsedAsGlobalOptionInPollID(optionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting used as global option id in poll id: %w", err)
	}

	if exist {
		return pollID, nil
	}

	return 0, fmt.Errorf("database seems corrupted. option %d has no poll id", optionID)
}
