package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Projector handels the restriction for the projector collection.
//
// The user can see a projector, if the user has projector.can_see.
//
// If the projector has internal=true, then the user needs projector.can_manage
//
// Mode A: The user can see the projector.
type Projector struct{}

// Name returns the collection name.
func (p Projector) Name() string {
	return "projector"
}

// MeetingID returns the meetingID for the object.
func (p Projector) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Projector_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p Projector) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p Projector) see(ctx context.Context, fetcher *dsfetch.Fetch, projectorIDs []int) ([]attribute.Func, error) {
	return byMeeting(ctx, fetcher, p, projectorIDs, func(meetingID int, projectorIDs []int) ([]attribute.Func, error) {
		internal := make([]bool, len(projectorIDs))
		for i, id := range projectorIDs {
			fetcher.Projector_IsInternal(id).Lazy(&internal[i])
		}

		if err := fetcher.Execute(ctx); err != nil {
			return nil, fmt.Errorf("getting internal flag from projectors: %w", err)
		}

		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		forInternal := attribute.FuncInGroup(groupMap[perm.ProjectorCanManage])
		forPublic := attribute.FuncInGroup(groupMap[perm.ProjectorCanSee])

		result := make([]attribute.Func, len(projectorIDs))
		for i := range projectorIDs {
			if internal[i] {
				result[i] = forInternal
				continue
			}
			result[i] = forPublic
		}

		return result, nil
	})
}
