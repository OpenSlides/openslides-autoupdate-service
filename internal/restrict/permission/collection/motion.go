package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Motion initializes a motion.
func Motion(dp dataprovider.DataProvider) perm.ConnecterFunc {
	m := &motion{dp}
	return func(s perm.HandlerStore) {
		s.RegisterAction("motion.delete", m.modify(perm.MotionCanManage))
		s.RegisterAction("motion.set_state", m.modify(perm.MotionCanManageMetadata))
		s.RegisterAction("motion.create", m.create())
		s.RegisterAction("motion_submitter.create", m.submitterCreate())
		s.RegisterAction("motion.update", m.modify(perm.MotionCanManage))
		s.RegisterAction("motion_comment.delete", perm.ActionFunc(m.commentModify))
		s.RegisterAction("motion_comment.update", perm.ActionFunc(m.commentModify))
		s.RegisterAction("motion_comment.create", perm.ActionFunc(m.commentCreate))

		s.RegisterRestricter("motion", perm.CollectionFunc(m.readMotion))
		s.RegisterRestricter("motion_submitter", perm.CollectionFunc(m.readSubmitter))
		s.RegisterRestricter("motion_block", m.readBlock())
		s.RegisterRestricter("motion_change_recommendation", m.readChangeRecommendation())
		s.RegisterRestricter("motion_comment_section", perm.CollectionFunc(m.readCommentSection))
		s.RegisterRestricter("motion_comment", perm.CollectionFunc(m.readComment))
	}
}

type motion struct {
	dp dataprovider.DataProvider
}

func (m *motion) create() perm.ActionFunc {
	allowList := map[string]bool{
		"title":                true,
		"text":                 true,
		"reason":               true,
		"category_id":          true,
		"statute_paragraph_id": true,
		"workflow_id":          true,
		"meeting_id":           true,
	}

	allowListAmendment := map[string]bool{
		"parent_id":            true,
		"amendment_paragraphs": true,
	}

	for k := range allowList {
		allowListAmendment[k] = true
	}

	return func(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
		meetingID, err := strconv.Atoi(string(payload["meeting_id"]))
		if err != nil {
			return false, fmt.Errorf("invalid field meeting_id in payload: %w", err)
		}

		perms, err := perm.New(ctx, m.dp, userID, meetingID)
		if err != nil {
			return false, fmt.Errorf("fetching perms: %w", err)
		}

		if perms.Has(perm.MotionCanManage) {
			return true, nil
		}

		requiredPerm := perm.MotionCanCreate
		aList := allowList
		if _, ok := payload["parent_id"]; ok {
			requiredPerm = perm.MotionCanCreateAmendments
			aList = allowListAmendment
		}

		if !perms.Has(requiredPerm) {
			perm.LogNotAllowedf("User %d does not have permission %s", userID, requiredPerm)
			return false, nil
		}

		for e := range payload {
			if !aList[string(e)] {
				perm.LogNotAllowedf("Field `%s` is forbidden for non manager.", e)
				return false, nil
			}
		}

		return true, nil
	}
}

func (m *motion) modify(managePerm perm.TPermission) perm.ActionFunc {
	return func(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
		motionFQID := fmt.Sprintf("motion/%s", payload["id"])
		meetingID, err := m.dp.MeetingFromModel(ctx, motionFQID)
		if err != nil {
			return false, fmt.Errorf("getting meeting for %s: %w", motionFQID, err)
		}

		perms, err := perm.New(ctx, m.dp, userID, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting perms: %w", err)
		}

		if perms.Has(managePerm) {
			return true, nil
		}

		// Non managers can only edit some fields
		for k := range payload {
			switch k {
			case "id", "title", "text", "reason", "amendment_paragraphs":
				continue
			default:
				perm.LogNotAllowedf("Non managers can not modify field %s", k)
				return false, nil
			}
		}

		motionID, err := strconv.Atoi(string(payload["id"]))
		if err != nil {
			return false, fmt.Errorf("invalid payload: %w", err)
		}

		b, err := canSeeMotion(ctx, m.dp, userID, motionID, perms)
		if err != nil {
			return false, fmt.Errorf("getting canSee: %w", err)
		}

		if !b {
			perm.LogNotAllowedf("User %d can not see the motion", userID)
			return false, nil
		}

		var submitterIDs []int
		if err := m.dp.GetIfExist(ctx, motionFQID+"/submitter_ids", &submitterIDs); err != nil {
			return false, fmt.Errorf("getting submitter ids: %w", err)
		}

		var isSubmitter bool
		for _, sid := range submitterIDs {
			var sUserID int
			if err := m.dp.Get(ctx, fmt.Sprintf("motion_submitter/%d/user_id", sid), &sUserID); err != nil {
				return false, fmt.Errorf("getting userid of sumitter %d: %w", sid, err)
			}
			if sUserID == userID {
				isSubmitter = true
				break
			}
		}

		if !isSubmitter {
			perm.LogNotAllowedf("User %d is not a manager and not a submitter of %s", userID, motionFQID)
			return false, nil
		}

		var stateID int
		if err := m.dp.Get(ctx, motionFQID+"/state_id", &stateID); err != nil {
			return false, fmt.Errorf("getting stateID: %w", err)
		}

		var allowSubmitterEdit bool
		if err := m.dp.GetIfExist(ctx, fmt.Sprintf("motion_state/%d/allow_submitter_edit", stateID), &allowSubmitterEdit); err != nil {
			return false, fmt.Errorf("getting allow_submitter_edit: %w", err)
		}

		if !allowSubmitterEdit {
			perm.LogNotAllowedf("Motion state does not allow submitter edites")
			return false, nil
		}

		return true, nil
	}
}

func canSeeMotion(ctx context.Context, dp dataprovider.DataProvider, userID int, motionID int, perms *perm.Permission) (bool, error) {
	if perms.Has(perm.MotionCanManage) {
		return true, nil
	}

	if !perms.Has(perm.MotionCanSee) {
		return false, nil
	}

	motionFQID := fmt.Sprintf("motion/%d", motionID)

	var stateID int
	if err := dp.Get(ctx, motionFQID+"/state_id", &stateID); err != nil {
		return false, fmt.Errorf("getting motion state: %w", err)
	}

	var restriction []string
	field := fmt.Sprintf("motion_state/%d/restrictions", stateID)
	if err := dp.GetIfExist(ctx, field, &restriction); err != nil {
		return false, fmt.Errorf("getting field %s: %w", field, err)
	}

	if len(restriction) == 0 {
		return true, nil
	}

	for _, r := range restriction {
		switch r {
		case "motion.can_see_internal", "motion.can_manage_metadata", "motion.can_manage":
			if perms.Has(perm.TPermission(r)) {
				return true, nil
			}

		case "is_submitter":
			var submitterIDs []int
			if err := dp.GetIfExist(ctx, motionFQID+"/submitter_ids", &submitterIDs); err != nil {
				return false, fmt.Errorf("getting field %s/submitter_ids: %w", motionFQID, err)
			}

			for _, sid := range submitterIDs {
				var uid int
				f := fmt.Sprintf("motion_submitter/%d/user_id", sid)
				if err := dp.Get(ctx, f, &uid); err != nil {
					return false, fmt.Errorf("getting field %s: %w", f, err)
				}
				if uid == userID {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func (m *motion) readMotion(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		meetingID, err := m.dp.MeetingFromModel(ctx, fmt.Sprintf("motion/%d", fqfield.ID))
		if err != nil {
			return false, fmt.Errorf("getting meetingID from motion: %w", err)
		}

		perms, err := perm.New(ctx, m.dp, userID, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting user permissions: %w", err)
		}

		return canSeeMotion(ctx, m.dp, userID, fqfield.ID, perms)
	})
}

func (m *motion) submitterCreate() perm.ActionFunc {
	return func(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
		motionFQID := fmt.Sprintf("motion/%s", payload["motion_id"])
		meetingID, err := m.dp.MeetingFromModel(ctx, motionFQID)
		if err != nil {
			return false, fmt.Errorf("getting meeting for %s: %w", motionFQID, err)
		}

		p, err := perm.HasPerm(ctx, m.dp, userID, meetingID, perm.MotionCanManageMetadata)
		if err != nil {
			return false, fmt.Errorf("getting perm: %w", err)
		}
		return p, nil
	}
}

func (m *motion) readSubmitter(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		var motionID int
		if err := m.dp.Get(ctx, fmt.Sprintf("motion_submitter/%d/motion_id", fqfield.ID), &motionID); err != nil {
			return false, fmt.Errorf("getting motionID: %w", err)
		}

		meetingID, err := m.dp.MeetingFromModel(ctx, fmt.Sprintf("motion/%d", motionID))
		if err != nil {
			return false, fmt.Errorf("getting meetingID from motion: %w", err)
		}

		perms, err := perm.New(ctx, m.dp, userID, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting user permissions: %w", err)
		}
		return canSeeMotion(ctx, m.dp, userID, motionID, perms)
	})
}

func (m *motion) readBlock() perm.CollectionFunc {
	return func(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
		return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
			fqid := fmt.Sprintf("motion_block/%d", fqfield.ID)
			meetingID, err := m.dp.MeetingFromModel(ctx, fqid)
			if err != nil {
				return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
			}

			perms, err := perm.New(ctx, m.dp, userID, meetingID)
			if err != nil {
				return false, fmt.Errorf("getting user permissions: %w", err)
			}

			if perms.Has(perm.MotionCanManage) {
				return true, nil
			}

			if !perms.Has(perm.MotionCanSee) {
				return false, nil
			}

			var internal bool
			if err := m.dp.GetIfExist(ctx, fqid+"/internal", &internal); err != nil {
				return false, fmt.Errorf("get /internal: %w", err)
			}

			if !internal {
				return true, nil
			}

			return false, nil
		})
	}
}

func (m *motion) readChangeRecommendation() perm.CollectionFunc {
	return func(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
		return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
			fqid := fmt.Sprintf("motion_change_recommendation/%d", fqfield.ID)
			meetingID, err := m.dp.MeetingFromModel(ctx, fqid)
			if err != nil {
				return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
			}

			perms, err := perm.New(ctx, m.dp, userID, meetingID)
			if err != nil {
				return false, fmt.Errorf("getting user permissions: %w", err)
			}

			if perms.Has(perm.MotionCanManage) {
				return true, nil
			}

			var motionID int
			if err := m.dp.Get(ctx, fqid+"/motion_id", &motionID); err != nil {
				return false, fmt.Errorf("getting motion id: %w", err)
			}

			motionOK, err := canSeeMotion(ctx, m.dp, userID, motionID, perms)
			if err != nil {
				return false, fmt.Errorf("checking permission for motion: %w", err)
			}
			if !motionOK {
				return false, nil
			}

			var internal bool
			if err := m.dp.GetIfExist(ctx, fqid+"/internal", &internal); err != nil {
				return false, fmt.Errorf("get /internal: %w", err)
			}

			if !internal {
				return true, nil
			}

			return perms.Has(perm.MotionCanManage), nil
		})
	}
}

func (m *motion) canSeeCommentSection(ctx context.Context, userID, id int) (bool, error) {
	fqid := fmt.Sprintf("motion_comment_section/%d", id)
	meetingID, err := m.dp.MeetingFromModel(ctx, fqid)
	if err != nil {
		return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
	}

	perms, err := perm.New(ctx, m.dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting user permissions: %w", err)
	}

	if perms.Has(perm.MotionCanManage) {
		return true, nil
	}

	var motionID int
	if err := m.dp.Get(ctx, fqid+"/motion_id", &motionID); err != nil {
		return false, fmt.Errorf("getting motion id: %w", err)
	}

	motionOK, err := canSeeMotion(ctx, m.dp, userID, motionID, perms)
	if err != nil {
		return false, fmt.Errorf("checking permission for motion: %w", err)
	}
	if !motionOK {
		return false, nil
	}

	var readGroupIDs []int
	if err := m.dp.GetIfExist(ctx, fqid+"/read_group_ids", &readGroupIDs); err != nil {
		return false, fmt.Errorf("getting read groups: %w", err)
	}
	for _, gid := range readGroupIDs {
		if perms.InGroup(gid) {
			return true, nil
		}
	}
	return false, nil
}

func (m *motion) readCommentSection(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		return m.canSeeCommentSection(ctx, userID, fqfield.ID)
	})
}

func (m *motion) commentAction(ctx context.Context, userID int, sectionID int) (bool, error) {
	fqid := "motion_comment_section/" + strconv.Itoa(sectionID)
	meetingID, err := m.dp.MeetingFromModel(ctx, fqid)
	if err != nil {
		return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
	}

	perms, err := perm.New(ctx, m.dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms: %w", err)
	}

	if perms.Has(perm.MotionCanManage) {
		return true, nil
	}

	for _, field := range []string{"/read_group_ids", "/write_group_ids"} {
		var groupIDs []int
		if err := m.dp.GetIfExist(ctx, fqid+field, &groupIDs); err != nil {
			return false, fmt.Errorf("getting groups from %s: %w", field, err)
		}

		var inGroup bool
		for _, gid := range groupIDs {
			if perms.InGroup(gid) {
				inGroup = true
			}
		}

		if !inGroup {
			return false, nil
		}
	}
	return true, nil
}

func (m *motion) commentModify(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var sectionID int
	if err := m.dp.Get(ctx, fmt.Sprintf("motion_comment/%s/section_id", payload["id"]), &sectionID); err != nil {
		return false, fmt.Errorf("getting section id: %w", err)
	}

	return m.commentAction(ctx, userID, sectionID)
}

func (m *motion) commentCreate(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	sectionID, err := strconv.Atoi(string(payload["section_id"]))
	if err != nil {
		return false, fmt.Errorf("invalid section_id: %s", payload["section_id"])
	}

	return m.commentAction(ctx, userID, sectionID)
}

func (m *motion) readComment(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		var sectionID int
		if err := m.dp.Get(ctx, fmt.Sprintf("motion_comment/%d/section_id", fqfield.ID), &sectionID); err != nil {
			return false, fmt.Errorf("getting section id: %w", err)
		}
		return m.canSeeCommentSection(ctx, userID, sectionID)
	})
}

func canSeeMotionSupporter(ctx context.Context, dp dataprovider.DataProvider, userID int, p *perm.Permission, ids []int) (bool, error) {
	for _, id := range ids {
		b, err := canSeeMotion(ctx, dp, userID, id, p)
		if err != nil {
			return false, err
		}
		if b {
			return true, nil
		}
	}
	return false, nil
}

func canSeeMotionSubmitter(ctx context.Context, dp dataprovider.DataProvider, userID int, p *perm.Permission, ids []int) (bool, error) {
	for _, id := range ids {
		var motionID int
		if err := dp.Get(ctx, fmt.Sprintf("motion_submitter/%d/motion_id", id), &motionID); err != nil {
			return false, fmt.Errorf("getting motion id: %w", err)
		}

		b, err := canSeeMotion(ctx, dp, userID, id, p)
		if err != nil {
			return false, fmt.Errorf("can see motion %d: %w", motionID, err)
		}
		if b {
			return true, nil
		}
	}
	return false, nil
}
