package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Projection handels the restriction for the projection collection.
//
// The user can see a projection, if the user has projector.can_see.
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

func (p Projection) see(ctx context.Context, fetcher *dsfetch.Fetch, projectionIDs []int) ([]attribute.Func, error) {
	return meetingPerm(ctx, fetcher, p, projectionIDs, perm.ProjectorCanSee)
}
