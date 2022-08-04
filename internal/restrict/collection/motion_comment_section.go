package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionCommentSection handels restrictions of the collection motion_comment_section.
//
// The user can see a motion comment section if any of:
//
//	The user has motion.can_see and has at least one group in common with motion_comment_section/read_group_ids
//	The user has motion.can_manage.
//
// The user can see the motion comment section.
type MotionCommentSection struct{}

// MeetingID returns the meetingID for the object.
func (m MotionCommentSection) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionCommentSection_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionCommentSection) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionCommentSection) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, motionCommentSectionIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, m, motionCommentSectionIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		if perms.Has(perm.MotionCanManage) {
			return ids, nil
		}

		if !perms.Has(perm.MotionCanSee) {
			return nil, nil
		}

		allowed, err := eachCondition(ids, func(motionCommentSectionID int) (bool, error) {
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
		})

		if err != nil {
			return nil, fmt.Errorf("checking if user is in read group: %w", err)
		}

		return allowed, nil
	})
}
