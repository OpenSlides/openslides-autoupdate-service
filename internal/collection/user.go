package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// User handels the permissions of user-actions and the user collection.
func User(dp dataprovider.DataProvider) perm.ConnecterFunc {
	u := &user{dp: dp}
	return func(s perm.HandlerStore) {
		s.RegisterWriteHandler("user.create", perm.WriteCheckerFunc(u.create))
		s.RegisterWriteHandler("user.update_self", perm.WriteCheckerFunc(u.updateSelf))
		s.RegisterWriteHandler("user.update", perm.WriteCheckerFunc(u.update))
		s.RegisterWriteHandler("user.set_password_self", perm.WriteCheckerFunc(u.passwordSelf))
		s.RegisterWriteHandler("user.reset_password_to_default", perm.WriteCheckerFunc(u.update))
		s.RegisterWriteHandler("user.generate_new_password", perm.WriteCheckerFunc(u.update))
		s.RegisterWriteHandler("user.set_password", perm.WriteCheckerFunc(u.update))
		s.RegisterWriteHandler("user.delete", perm.WriteCheckerFunc(u.update))

		s.RegisterReadHandler("user", perm.ReadCheckerFunc(u.read))
	}
}

type user struct {
	dp dataprovider.DataProvider
}

func (u *user) create(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var orgaLevel string
	if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/organisation_management_level", userID), &orgaLevel); err != nil {
		return false, fmt.Errorf("getting organisation level: %w", err)
	}
	switch orgaLevel {
	case "can_manage_organisation", "can_manage_users":
		return true, nil
	default:
		return false, nil
	}
}

func (u *user) updateSelf(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	return userID != 0, nil
}

func (u *user) update(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var orgaLevel string
	if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/organisation_management_level", userID), &orgaLevel); err != nil {
		return false, fmt.Errorf("getting organisation level: %w", err)
	}

	return orgaLevel != "", nil
}

func (u *user) passwordSelf(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	if userID == 0 {
		return false, nil
	}

	var meetingID int
	if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/meeting_id", userID), &meetingID); err != nil {
		return false, fmt.Errorf("getting meeting id: %w", err)
	}

	if meetingID == 0 {
		// Not a temporary user.
		return true, nil
	}

	b, err := perm.HasPerm(ctx, u.dp, userID, meetingID, "user.can_change_own_password")
	if err != nil {
		return false, fmt.Errorf("getting perm: %w", err)
	}

	return b, nil
}

func (u *user) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	var orgaLevel string
	if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/organisation_management_level", userID), &orgaLevel); err != nil {
		return fmt.Errorf("getting organisation level: %w", err)
	}

	meetingFields := make(map[int]map[string]bool)

	grouped := groupByID(fqfields)
	for _, fqfields := range grouped {
		seeFields := make(map[string]bool)

		if orgaLevel != "" {
			addSlice(seeFields, canSeeFields[3])
		}
		if fqfields[0].ID == userID {
			addSlice(seeFields, canSeeFields[4])
		}

		var meetingIDsStr []string
		if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/group_$_ids", fqfields[0].ID), &meetingIDsStr); err != nil {
			return fmt.Errorf("getting meeting ids: %w", err)
		}

		meetingIDs := make([]int, len(meetingIDsStr))
		for i, midS := range meetingIDsStr {
			mid, err := strconv.Atoi(midS)
			if err != nil {
				return fmt.Errorf("invalid meetingid: %s", midS)
			}
			meetingIDs[i] = mid
		}

		for _, meetingID := range meetingIDs {
			fields, ok := meetingFields[meetingID]
			if !ok {
				fields = make(map[string]bool)
				perms, err := perm.New(ctx, u.dp, userID, meetingID)
				if err != nil {
					return fmt.Errorf("getting perms for user %d in meeting %d: %w", userID, meetingID, err)
				}
				if perms.Has("user.can_see") {
					addSlice(fields, canSeeFields[0])
				}
				if perms.Has("user.can_see_extra_data") {
					addSlice(fields, canSeeFields[1])
				}
				if perms.Has("user.can_manage") {
					addSlice(fields, canSeeFields[2])
				}
				meetingFields[meetingID] = fields
			}

			addMap(seeFields, fields)
		}

		if len(seeFields) == 0 {
			r, err := isRequired(ctx, u.dp, userID, fqfields[0].ID, meetingIDs)
			if err != nil {
				return err
			}
			if r {
				addSlice(seeFields, canSeeFields[0])
			}
		}

		for _, f := range fqfields {
			if !seeFields[templateFieldPrefix(f)] {
				continue
			}

			if mid := meetingFilter(f); mid != 0 {
				if !meetingFields[mid][templateFieldPrefix(f)] {
					continue
				}
			}
			result[f.String()] = true
		}
	}
	return nil
}

func isRequired(ctx context.Context, dp dataprovider.DataProvider, userID int, otherUserID int, meetingIDs []int) (bool, error) {
	var ids []int
	for _, mid := range meetingIDs {
		p, err := perm.New(ctx, dp, userID, mid)
		if err != nil {
			return false, fmt.Errorf("getting perms: %w", err)
		}
		if p == nil {
			continue
		}

		// Speaker
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/speaker_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting speaker ids: %w", err)
		}
		if len(ids) > 0 && canSeeSpeaker(p) {
			return true, nil
		}

		// Motion Supporter
		ids = nil
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/supported_motion_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting supporter ids: %w", err)
		}
		if len(ids) > 0 {
			b, err := canSeeMotionSupporter(ctx, dp, userID, p, ids)
			if err != nil {
				return false, err
			}
			if b {
				return true, nil
			}
		}

		// Motion Submitter
		ids = nil
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/submitted_motion_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting submitter ids: %w", err)
		}
		if len(ids) > 0 {
			b, err := canSeeMotionSubmitter(ctx, dp, userID, p, ids)
			if err != nil {
				return false, err
			}
			if b {
				return true, nil
			}
		}

		// Poll voted
		ids = nil
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/poll_voted_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting poll ids: %w", err)
		}
		if len(ids) > 0 {
			b, err := canSeePolls(ctx, dp, p, userID, ids)
			if err != nil {
				return false, err
			}
			if b {
				return true, nil
			}
		}

		// Poll option
		ids = nil
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/option_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting option ids: %w", err)
		}
		if len(ids) > 0 {
			b, err := canSeePollOptions(ctx, dp, p, userID, ids)
			if err != nil {
				return false, err
			}
			if b {
				return true, nil
			}
		}

		// Assignment Candidate
		ids = nil
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/assignment_candidate_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting assignment candidates ids: %w", err)
		}
		if len(ids) > 0 && canSeeAssignmentCandidate(p) {
			return true, nil
		}

		// Projection
		ids = nil
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/projection_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting projection ids: %w", err)
		}
		if len(ids) > 0 && canSeeProjection(p) {
			return true, nil
		}

		// Projection
		ids = nil
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/current_projector_$%d_ids", otherUserID, mid), &ids); err != nil {
			return false, fmt.Errorf("getting current projector ids: %w", err)
		}
		if len(ids) > 0 && canSeeProjector(p) {
			return true, nil
		}
	}
	return false, nil
}

// groupByID groups a list of fqfields by there id part.
//
// It expects the input to be sorted.
func groupByID(fqfields []perm.FQField) map[int][]perm.FQField {
	grouped := make(map[int][]perm.FQField)
	for _, f := range fqfields {
		grouped[f.ID] = append(grouped[f.ID], f)
	}
	return grouped
}

func addSlice(data map[string]bool, slice []string) {
	for _, v := range slice {
		data[v] = true
	}
}

func addMap(data map[string]bool, m map[string]bool) {
	for k, v := range m {
		if v {
			data[k] = true
		}
	}
}

func templateFieldPrefix(fqfield perm.FQField) string {
	i := strings.IndexByte(fqfield.Field, '$')
	if i < 0 {
		return fqfield.Field
	}
	return fqfield.Field[:i+1]
}

// meetingFilter acts on speciel fields containing a meeting id. For this
// fields, it returns the meeting id. for other fields, it returns 0.
func meetingFilter(fqfield perm.FQField) int {
	p := templateFieldPrefix(fqfield)
	switch p {
	case "number_$", "structure_level_$", "about_me_$", "vote_weight_$":
		if len(p) == len(fqfield.Field) {
			// fqfield is the structure field
			break
		}
		var mid int
		fmt.Sscanf(fqfield.Field, p+"%d", &mid)
		return mid
	}
	return 0
}

// canSeeFields list all fields of the user ordered by permission levels.
//
// Structured fields contain only the prefix, ending with the $.
var canSeeFields = [...][]string{
	{ // can_see
		"id",
		"username",
		"title",
		"first_name",
		"last_name",
		"is_physical_person",
		"gender",
		"default_number",
		"default_structure_level",
		"is_demo_user",
		"is_present_in_meeting_ids",
		"number_$",
		"structure_level_$",
		"about_me_$",
		"vote_weight_$",
		"group_$",
		"speaker_$",
		"supported_motion_$",
		"submitted_motion_$",
		"poll_voted_$",
		"option_$",
		"vote_$",
		"vote_delegated_vote_$",
		"assignment_candidate_$",
		"projection_$",
		"current_projector_$",
	},
	{ // can_see_extra
		"is_active",
		"email",
		"last_email_send",
		"meeting_id",
		"guest_meeting_ids",
		"comment_$",
		"vote_delegated_$",
		"vote_delegations_$",
		"default_vote_weight",
	},
	{ // can_manage
		"default_password",
	},
	{ // orga can_manage_user
		"id",
		"username",
		"title",
		"first_name",
		"last_name",
		"is_physical_person",
		"gender",
		"default_number",
		"default_structure_level",
		"is_demo_user",
		"number_$",
		"structure_level_$",
		"vote_weight_$",
		"organisation_management_level",
		"committee_as_member_ids",
		"email",
		"last_email_send",
		"committee_as_manager_ids",
		"is_active",
		"guest_meeting_ids",
		"default_password",
		"default_vote_weight",
	},
	{ // own user
		"id",
		"username",
		"title",
		"first_name",
		"last_name",
		"is_physical_person",
		"gender",
		"default_number",
		"default_structure_level",
		"is_demo_user",
		"is_present_in_meeting_ids",
		"number_$",
		"structure_level_$",
		"about_me_$",
		"vote_weight_$",
		"group_$",
		"speaker_$",
		"supported_motion_$",
		"submitted_motion_$",
		"poll_voted_$",
		"option_$",
		"vote_$",
		"vote_delegated_vote_$",
		"assignment_candidate_$",
		"projection_$",
		"current_projector_$",
		"is_active",
		"email",
		"last_email_send",
		"meeting_id",
		"guest_meeting_ids",
		"comment_$",
		"vote_delegated_$",
		"vote_delegations_$",
		"default_password",
		"organisation_management_level",
		"personal_note_$",
		"committee_as_member_ids",
		"committee_as_manager_ids",
		"default_vote_weight",
	},
}
