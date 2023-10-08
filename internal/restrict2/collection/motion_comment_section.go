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
		fetcher.MotionCommentSection_ReadGroupIDs(id).Lazy(&readGroupIDs[i])
		fetcher.MotionCommentSection_WriteGroupIDs(id).Lazy(&writeGroupIDs[i])
		fetcher.MotionCommentSection_SubmitterCanWrite(id).Lazy(&submitterCanWrite[i])
		fetcher.MotionCommentSection_MeetingID(id).Lazy(&meetingID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching motion comment section data: %w", err)
	}

	out := make([]attribute.Func, len(motionCommentSectionIDs))
	for i := range motionCommentSectionIDs {
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

// SeeAs checks if the request user can see a comment section. Returns 1 the
// user can be seen the object because of a read or write group. Returns 2 if
// the user can see the object because of the submitter rule.
func (m MotionCommentSection) seeAs(ctx context.Context, ds *dsfetch.Fetch, perms *perm.Permission, motionCommentSectionID int) (int, error) {
	readGroups, err := ds.MotionCommentSection_ReadGroupIDs(motionCommentSectionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting read_groups: %w", err)
	}

	writeGroups, err := ds.MotionCommentSection_WriteGroupIDs(motionCommentSectionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting write_groups: %w", err)
	}

	for _, gid := range append(readGroups, writeGroups...) {
		if perms.InGroup(gid) {
			return 1, nil
		}
	}

	submitterCanWrite, err := ds.MotionCommentSection_SubmitterCanWrite(motionCommentSectionID).Value(ctx)
	if err != nil {
		return 0, fmt.Errorf("getting submitter_can_write: %w", err)
	}

	if submitterCanWrite {
		return 2, nil
	}

	return 0, nil
}
