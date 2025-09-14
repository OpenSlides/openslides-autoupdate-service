package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/perm"
)

// Projection handels the restriction for the projection collection.
//
// The user can see a projection,
// * if he has projector.can_manage or
// * he can see the projector linked in projection/current_projector_id.
//
// Mode A: The user can see the projection.
type Projection struct{}

// Name returns the collection name.
func (p Projection) Name() string {
	return "projection"
}

// MeetingID returns the meetingID for the object.
func (p Projection) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Projection_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p Projection) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p Projection) see(ctx context.Context, ds *dsfetch.Fetch, projectionIDs ...int) ([]int, error) {
	projectorRestrictor := Collection(ctx, "projector").Modes("A")

	return eachMeeting(ctx, ds, p, projectionIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permission: %w", err)
		}

		if perms.Has(perm.ProjectorCanManage) {
			return ids, nil
		}

		currentProjector := make([]dsfetch.Maybe[int], len(ids))
		for i, id := range ids {
			ds.Projection_CurrentProjectorID(id).Lazy(&currentProjector[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return nil, fmt.Errorf("reading current_projector_id: %w", err)
		}

		var allowed []int
		for i, maybeID := range currentProjector {
			projectorID, hasCurrent := maybeID.Value()
			if !hasCurrent {
				continue
			}

			// This chekcs each projector by its own. But the result should be
			// in the cache anyway. So this should be more performent, then
			// putting many projector-ids in a set.
			canSeeProjector, err := projectorRestrictor(ctx, ds, projectorID)
			if err != nil {
				return nil, fmt.Errorf("checking projector restrictor: %w", err)
			}

			if len(canSeeProjector) == 0 {
				continue
			}

			allowed = append(allowed, ids[i])
		}

		return allowed, nil
	})
}
