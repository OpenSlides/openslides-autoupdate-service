package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Projector handels the restriction for the projector collection.
//
// The user can see a projector, if the user has projector.can_see.
//
// Mode A: The user can see the projector.
type Projector struct {
	name string
}

// Name returns the collection name.
func (p Projector) Name() string {
	return p.name
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

func (p Projector) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, projectorIDs ...int) error {
	return meetingPerm(ctx, ds, p, "A", projectorIDs, mperms, perm.ProjectorCanSee, attrMap)
}
