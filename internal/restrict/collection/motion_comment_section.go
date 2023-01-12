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
//	The user has motion.can_see and
//		has at least one group in common with motion_comment_section/read_group_ids or
//		has at least one group in common with motion_comment_section/write_group_ids or
//		submitter_can_write is set to true for this comment section
//
//	The user has motion.can_manage.
//
// Mode A: The user can see the motion comment section.
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
	// TODO: Implement me
	return Allways(m.Name(), mode)

	// switch mode {
	// case "A":
	// 	return m.see
	// }
	// return nil
}

//func (m MotionCommentSection) see(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, motionCommentSectionIDs ...int) error {
// return eachMeeting(ctx, ds, m, motionCommentSectionIDs, func(meetingID int, ids []int) error {
// 	groupMap, err := mperms.Meeting(ctx, ds, meetingID)
// 	if err != nil {
// 		return fmt.Errorf("groupMap: %w", err)
// 	}

// 	for _, motionCommentSectionID := range motionCommentSectionIDs {
// 		var readGroupIDs []int
// 		var writeGroupIDs []int
// 		var submitterCanWrite bool
// 		ds.MotionCommentSection_ReadGroupIDs(motionCommentSectionID).Lazy(&readGroupIDs)
// 		ds.MotionCommentSection_WriteGroupIDs(motionCommentSectionID).Lazy(&writeGroupIDs)
// 		ds.MotionCommentSection_SubmitterCanWrite(motionCommentSectionID).Lazy(&submitterCanWrite)
// 		if err := ds.Execute(ctx); err != nil {
// 			return fmt.Errorf("getting motion commect section data for id %d: %w", motionCommentSectionID, err)
// 		}

// 		if
// 	}

// 	attrInternal := Attributes{
// 		GlobalPermission: byte(perm.OMLSuperadmin),
// 		GroupIDs:         groupMap[perm.MotionCanManage],
// 	}

// 	attrPublic := Attributes{
// 		GlobalPermission: byte(perm.OMLSuperadmin),
// 		GroupIDs:         groupMap[perm.MotionCanSee],
// 	}

// 	perms, err := mperms.Meeting(ctx, meetingID)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting permissions: %w", err)
// 	}

// 	if perms.Has(perm.MotionCanManage) {
// 		return ids, nil
// 	}

// 	if !perms.Has(perm.MotionCanSee) {
// 		return nil, nil
// 	}

// 	allowed, err := eachCondition(ids, func(motionCommentSectionID int) (bool, error) {
// 		seeAs, err := m.seeAs(ctx, ds, perms, motionCommentSectionID)
// 		if err != nil {
// 			return false, err
// 		}

// 		return seeAs > 0, nil
// 	})

// 	if err != nil {
// 		return nil, fmt.Errorf("checking if user is in read group: %w", err)
// 	}

// 	return allowed, nil
// })
//}

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
