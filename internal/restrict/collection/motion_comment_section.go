package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionCommentSection handels restrictions of the collection motion_comment_section.
type MotionCommentSection struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionCommentSection) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionCommentSection) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionCommentSectionID int) (bool, error) {
	meetingID, err := ds.MotionCommentSection_MeetingID(motionCommentSectionID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if perms.Has(perm.MotionCanManage) {
		return true, nil
	}

	if !perms.Has(perm.MotionCanSee) {
		return false, nil
	}

	readGroups, err := ds.MotionCommentSection_ReadGroupIDs(motionCommentSectionID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting readGroups: %w", err)
	}

	for _, gid := range readGroups {
		if perms.InGroup(gid) {
			return true, nil
		}
	}

	return false, nil
}
