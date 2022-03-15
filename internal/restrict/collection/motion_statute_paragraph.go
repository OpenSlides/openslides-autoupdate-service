package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionStatuteParagraph handels restrictions of the collection motion_statute_paragraph.
//
// The user can see a motion statute paragraph if the user has motion.can_see.
//
// Mode A: The user can see the motion statute paragraph.
type MotionStatuteParagraph struct{}

// MeetingID returns the meetingID for the object.
func (m MotionStatuteParagraph) MeetingID(ctx context.Context, ds *datastore.Request, id int) (int, bool, error) {
	meetingID, err := ds.MotionStatuteParagraph_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionStatuteParagraph) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionStatuteParagraph) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, MotionStatuteParagraphID int) (bool, error) {
	meetingID, err := ds.MotionStatuteParagraph_MeetingID(MotionStatuteParagraphID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permission: %w", err)
	}

	return perms.Has(perm.MotionCanSee), nil
}
