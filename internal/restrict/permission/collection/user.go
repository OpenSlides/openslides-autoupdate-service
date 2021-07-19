package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// User handels the permissions of user-actions and the user collection.
func User(dp dataprovider.DataProvider) perm.ConnecterFunc {
	u := &user{dp: dp}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("user", perm.CollectionFunc(u.read))
	}
}

type user struct {
	dp dataprovider.DataProvider
}

// committeeManagerMembers returns all userIDs as a set that the userID can
// manage as committee manager. (Member in a meeting from a committe the userID
// is manager.)
func committeeManagerMembers(ctx context.Context, dp dataprovider.DataProvider, userID int) (map[int]bool, error) {
	if userID == 0 {
		return nil, nil
	}

	var committeManagerLevel []string
	if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/committee_$_management_level", userID), &committeManagerLevel); err != nil {
		return nil, fmt.Errorf("getting committee manager level: %w", err)
	}

	members := make(map[int]bool)
	for _, id := range committeManagerLevel {
		var level string
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/committee_$%s_management_level", userID, id), &level); err != nil {
			return nil, fmt.Errorf("getting committee manager level for committee %s: %w", id, err)
		}

		if level != "can_manage" {
			continue
		}

		var membersIDs []int
		if err := dp.GetIfExist(ctx, fmt.Sprintf("committee/%s/user_ids", id), &membersIDs); err != nil {
			return nil, fmt.Errorf("getting members: %w", err)
		}

		for _, memberID := range membersIDs {
			members[memberID] = true
		}

		// This is O(n³). There has to be another way.
		var meetingIDs []int
		if err := dp.GetIfExist(ctx, fmt.Sprintf("committee/%s/meeting_ids", id), &meetingIDs); err != nil {
			return nil, fmt.Errorf("getting meetings: %w", err)
		}

		for _, meetingID := range meetingIDs {
			var membersIDs []int
			if err := dp.GetIfExist(ctx, fmt.Sprintf("meeting/%d/user_ids", meetingID), &membersIDs); err != nil {
				return nil, fmt.Errorf("getting meeting members: %w", err)
			}

			for _, memberID := range membersIDs {
				members[memberID] = true
			}
		}
	}
	return members, nil
}

// delegatedVoteFromAndTo returns all user ID that the userID has given his vote
// or is delegated an all meetings.
func delegatedVoteFromAndTo(ctx context.Context, dp dataprovider.DataProvider, userID int) (map[int]bool, error) {
	if userID == 0 {
		return nil, nil
	}

	users := make(map[int]bool)

	var delegatedTo []string
	if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/vote_delegated_$_to_id", userID), &delegatedTo); err != nil {
		return nil, fmt.Errorf("get vote_delegated_to_id: %w", err)
	}

	for _, e := range delegatedTo {
		var oterUserID int
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/vote_delegated_$%s_to_id", userID, e), &oterUserID); err != nil {
			return nil, fmt.Errorf("get vote delegated to for meeting %s: %w", e, err)
		}
		users[oterUserID] = true
	}

	var delegatedFrom []string
	if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/vote_delegations_$_from_ids", userID), &delegatedFrom); err != nil {
		return nil, fmt.Errorf("get vote_delegated_from_ids: %w", err)
	}

	for _, e := range delegatedFrom {
		var userIDs []int
		if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/vote_delegations_$%s_from_ids", userID, e), &userIDs); err != nil {
			return nil, fmt.Errorf("get vote delegated from for meeting %s: %w", e, err)
		}
		for _, userID := range userIDs {
			users[userID] = true
		}
	}
	return users, nil
}

func (u *user) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	orgaLevel, err := u.dp.OrgaLevel(ctx, userID)
	if err != nil {
		return fmt.Errorf("getting organization level: %w", err)
	}

	committeeManagerMembers, err := committeeManagerMembers(ctx, u.dp, userID)
	if err != nil {
		return fmt.Errorf("getting members of committee: %w", err)
	}

	delegated, err := delegatedVoteFromAndTo(ctx, u.dp, userID)
	if err != nil {
		return fmt.Errorf("getting delegated (from/to) user ids: %w", err)
	}

	meetingFields := make(map[int]map[string]bool)

	grouped := groupByID(fqfields)
	for _, fqfields := range grouped {
		seeFields := make(map[string]bool)

		if orgaLevel != "" {
			addSlice(seeFields, canSeeFields['A'])
			addSlice(seeFields, canSeeFields['C'])
			addSlice(seeFields, canSeeFields['D'])
			addSlice(seeFields, canSeeFields['E'])
			addSlice(seeFields, canSeeFields['F'])
		}
		if fqfields[0].ID == userID {
			addSlice(seeFields, canSeeFields['A'])
			addSlice(seeFields, canSeeFields['B'])
			addSlice(seeFields, canSeeFields['C'])
			addSlice(seeFields, canSeeFields['E'])
			addSlice(seeFields, canSeeFields['F'])
		}
		if committeeManagerMembers[fqfields[0].ID] || delegated[fqfields[0].ID] {
			addSlice(seeFields, canSeeFields['A'])
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
				if perms.Has(perm.UserCanSee) {
					addSlice(fields, canSeeFields['A'])
				}
				if perms.Has(perm.UserCanSeeExtraData) {
					addSlice(fields, canSeeFields['C'])
				}
				if perms.Has(perm.UserCanManage) {
					addSlice(fields, canSeeFields['D'])
					addSlice(fields, canSeeFields['E'])
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
				addSlice(seeFields, canSeeFields['A'])
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

// templateFieldPrefix returns the part of a templatefield until the $.
//
// If the field is not a template field, it returns the full field.
//
// Example:
// * user/1/speaker_$_ids -> user/1/speaker_$
// * user/1/speaker_$5_ids -> user/1/speaker_$
// * user/1/username -> user/1/username
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
var canSeeFields = map[byte][]string{
	'A': {
		"id",
		"username",
		"title",
		"first_name",
		"last_name",
		"is_physical_person",
		"gender",
		"default_number",
		"default_structure_level",
		"default_vote_weight",
		"is_demo_user",
		"is_present_in_meeting_ids",
		"number_$",
		"structure_level_$",
		"about_me_$",
		"vote_weight_$",
		"speaker_$",
		"supported_motion_$",
		"submitted_motion_$",
		"poll_voted_$",
		"option_$",
		"vote_$",
		"vote_delegated_vote_$",
		"assignment_candidate_$",
		"projection_$",
	},
	'B': {
		"personal_note_$",
	},
	'C': {
		"email",
		"vote_delegated_$",
		"vote_delegations_$",
		"group_$",
	},
	'D': {
		"last_email_send",
		"is_active",
		"comment_$",
		"default_password",
		"can_change_own_password",
	},
	'E': {
		"committee_ids",
		"committee_$",
		"meeting_ids",
	},
	'F': {
		"organization_management_level",
	},
}
