package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionComment handels restrictions of the collection motion_comment.
//
// The user can see a motion comment if
//
//	The user can see the linked motion and
//
//		has at least one group in common with motion_comment_section/read_group_ids, or
//		has at least one group in common with motion_comment_section/write_group_ids, or
//		submitter_can_write is set to true for the corresponding comment section AND the user is the submitter of this motion
//
//	Or the user has motion.can_manage.
//
// Mode A: The user can see the motion comment.
type MotionComment struct{}

// Name returns the collection name.
func (m MotionComment) Name() string {
	return "motion_comment"
}

// MeetingID returns the meetingID for the object.
func (m MotionComment) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	motionID, err := ds.MotionComment_MotionID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting motionID: %w", err)
	}

	return Motion{}.MeetingID(ctx, ds, motionID)
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionComment) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionComment) see(ctx context.Context, ds *dsfetch.Fetch, motionCommentIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachRelationField(ctx, ds.MotionComment_SectionID, motionCommentIDs, func(commentSectionID int, ids []int) ([]int, error) {
		commentSectionMeetingID, err := ds.MotionCommentSection_MeetingID(commentSectionID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("get meeting id from comment section %d: %w", commentSectionID, err)
		}

		perms, err := perm.FromContext(ctx, commentSectionMeetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions: %w", err)
		}

		seeSectionAs, err := MotionCommentSection{}.seeAs(ctx, ds, perms, commentSectionID)
		if err != nil {
			return nil, fmt.Errorf("checking motion comment section %d can see: %w", commentSectionID, err)
		}

		if seeSectionAs == 0 {
			return nil, nil
		}

		allowed, err := eachCondition(ids, func(motionCommentID int) (bool, error) {
			motionID, err := ds.MotionComment_MotionID(motionCommentID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting motion id from comment %d: %w", motionCommentID, err)
			}

			// TODO: Do this outside of section
			seeMotion, err := Collection(ctx, Motion{}.Name()).Modes("C")(ctx, ds, motionID)
			if err != nil {
				return false, fmt.Errorf("checking motion %d can see: %w", motionID, err)
			}

			if len(seeMotion) == 0 {
				return false, nil
			}

			if seeSectionAs == 1 {
				return true, nil
			}

			submitterIDs, err := ds.Motion_SubmitterIDs(motionID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting motion submitter ids for motion %d: %w", motionID, err)
			}

			for _, submitterID := range submitterIDs {
				meetingUser := ds.MotionSubmitter_MeetingUserID(submitterID).ErrorLater(ctx)
				userID, err := ds.MeetingUser_UserID(meetingUser).Value(ctx)
				if err != nil {
					return false, fmt.Errorf("getting user id for submitter %d: %w", submitterID, err)
				}

				if userID == requestUser {
					return true, nil
				}
			}

			return false, nil
		})
		if err != nil {
			return nil, fmt.Errorf("checking motion can see: %w", err)
		}

		return allowed, nil
	})
}
