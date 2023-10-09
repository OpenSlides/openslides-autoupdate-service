package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionCommentSection handels restrictions of the collection motion_comment_section.
//
// The user can see a motion comment section if any of:
//
//	The user has motion.can_see and
//		has at least one group in common with motion_comment_section/read_group_ids or
//		has at least one group in common with motion_comment_section/write_group_ids or
//		submitter_can_write is set to true for this comment section
//
//	The user has motion.can_manage.
//
// The user can see the motion comment section.
type MotionCommentSection struct{}

// Name returns the collection name.
func (m MotionCommentSection) Name() string {
	return "motion_comment_section"
}

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

func (m MotionCommentSection) see(ctx context.Context, fetcher *dsfetch.Fetch, motionCommentSectionIDs []int) ([]attribute.Func, error) {
	readGroupIDs := make([][]int, len(motionCommentSectionIDs))
	writeGroupIDs := make([][]int, len(motionCommentSectionIDs))
	submitterCanWrite := make([]bool, len(motionCommentSectionIDs))
	meetingID := make([]int, len(motionCommentSectionIDs))
	for i, id := range motionCommentSectionIDs {
		if id == 0 {
			continue
		}

		fetcher.MotionCommentSection_ReadGroupIDs(id).Lazy(&readGroupIDs[i])
		fetcher.MotionCommentSection_WriteGroupIDs(id).Lazy(&writeGroupIDs[i])
		fetcher.MotionCommentSection_SubmitterCanWrite(id).Lazy(&submitterCanWrite[i])
		fetcher.MotionCommentSection_MeetingID(id).Lazy(&meetingID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion comment section data: %w", err)
	}

	out := make([]attribute.Func, len(motionCommentSectionIDs))
	for i, id := range motionCommentSectionIDs {
		if id == 0 {
			continue
		}

		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		funcs := attribute.FuncInGroup(groupMap[perm.MotionCanSee])
		if !submitterCanWrite[i] {
			funcs = attribute.FuncAnd(
				attribute.FuncInGroup(groupMap[perm.MotionCanSee]),
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
