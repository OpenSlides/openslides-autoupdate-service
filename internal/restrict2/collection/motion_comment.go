package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
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

func (m MotionComment) see(ctx context.Context, fetcher *dsfetch.Fetch, motionCommentIDs []int) ([]attribute.Func, error) {
	meetingID := make([]int, len(motionCommentIDs))
	commentSectionID := make([]int, len(motionCommentIDs))
	motionID := make([]int, len(motionCommentIDs))
	for i, id := range motionCommentIDs {
		if id == 0 {
			continue
		}

		fetcher.MotionComment_MeetingID(id).Lazy(&meetingID[i])
		fetcher.MotionComment_MotionID(id).Lazy(&motionID[i])
		fetcher.MotionComment_SectionID(id).Lazy(&commentSectionID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion comment data: %w", err)
	}

	readGroupIDs := make([][]int, len(motionCommentIDs))
	writeGroupIDs := make([][]int, len(motionCommentIDs))
	submitterCanWrite := make([]bool, len(motionCommentIDs))
	submitterIDs := make([][]int, len(motionCommentIDs))
	for i, id := range motionCommentIDs {
		if id == 0 {
			continue
		}

		fetcher.MotionCommentSection_ReadGroupIDs(commentSectionID[i]).Lazy(&readGroupIDs[i])
		fetcher.MotionCommentSection_WriteGroupIDs(commentSectionID[i]).Lazy(&writeGroupIDs[i])
		fetcher.MotionCommentSection_SubmitterCanWrite(commentSectionID[i]).Lazy(&submitterCanWrite[i])

		fetcher.Motion_SubmitterIDs(motionID[i]).Lazy(&submitterIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion comment section and motion data: %w", err)
	}

	submitterMeetingUserIDs := make([][]int, len(motionCommentIDs))
	for i, id := range motionCommentIDs {
		if id == 0 {
			continue
		}

		submitterMeetingUserIDs[i] = make([]int, len(submitterIDs[i]))
		for j := range submitterMeetingUserIDs[i] {
			fetcher.MotionSubmitter_MeetingUserID(submitterIDs[i][j]).Lazy(&submitterMeetingUserIDs[i][j])
		}
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching meeting user from motion submitter: %w", err)
	}

	submitterUserIDs := make([][]int, len(motionCommentIDs))
	for i := range motionCommentIDs {
		submitterUserIDs[i] = make([]int, len(submitterMeetingUserIDs[i]))
		for j := range submitterMeetingUserIDs[i] {
			fetcher.MeetingUser_UserID(submitterMeetingUserIDs[i][j]).Lazy(&submitterUserIDs[i][j])
		}
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user from meeting user: %w", err)
	}

	canSeeMotion, err := canSeeRelatedCollection(ctx, fetcher, fetcher.MotionComment_MotionID, Collection(ctx, Motion{}).Modes("C"), motionCommentIDs)
	if err != nil {
		return nil, fmt.Errorf("checking motion can see: %w", err)
	}

	out := make([]attribute.Func, len(motionCommentIDs))
	for i, id := range motionCommentIDs {
		if id == 0 {
			continue
		}

		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		var funcs attribute.Func
		if submitterCanWrite[i] {
			funcs = attribute.FuncAnd(
				canSeeMotion[i],
				attribute.FuncUserIDs(submitterUserIDs[i]),
			)
		} else {
			funcs = attribute.FuncAnd(
				canSeeMotion[i],
				attribute.FuncOr(
					attribute.FuncInGroup(readGroupIDs[i]),
					attribute.FuncInGroup(writeGroupIDs[i]),
				),
			)
		}

		out[i] = attribute.FuncOr(
			attribute.FuncInGroup(groupMap[perm.MotionCanManage]),
			funcs,
		)
	}

	return out, nil
}
