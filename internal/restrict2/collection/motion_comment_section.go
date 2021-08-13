package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
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

func (m MotionCommentSection) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionCommentSectionID int) (bool, error) {
	meetingID := fetch.Field().MotionCommentSection_MeetingID(ctx, motionCommentSectionID)
	if err := fetch.Err(); err != nil {
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

	readGroups := fetch.Field().MotionCommentSection_ReadGroupIDs(ctx, motionCommentSectionID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting readGroups: %w", err)
	}

	for _, gid := range readGroups {
		if perms.InGroup(gid) {
			return true, nil
		}
	}

	return false, nil
}
