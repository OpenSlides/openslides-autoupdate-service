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

func (u *user) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	var orgaLevel string
	if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/organisation_management_level", userID), &orgaLevel); err != nil {
		return fmt.Errorf("getting organisation level: %w", err)
	}

	meetingFields := make(map[int]map[string]bool)

	grouped := groupByID(fqfields)
	for _, fqfields := range grouped {
		var meetingIDs []string
		if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/group_$_ids", fqfields[0].ID), &meetingIDs); err != nil {
			return fmt.Errorf("getting meeting ids: %w", err)
		}

		seeFields := make(map[string]bool)

		if orgaLevel != "" {
			addSlice(seeFields, canSeeFields[3])
		}
		if fqfields[0].ID == userID {
			addSlice(seeFields, canSeeFields[4])
		}

		for _, meetingID := range meetingIDs {
			mid, err := strconv.Atoi(meetingID)
			if err != nil {
				return fmt.Errorf("invalid meetingid: %s", meetingID)
			}

			fields, ok := meetingFields[mid]
			if !ok {
				fields = make(map[string]bool)
				perms, err := perm.New(ctx, u.dp, userID, mid)
				if err != nil {
					return fmt.Errorf("getting perms for user %d in meeting %d: %w", userID, mid, err)
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
				meetingFields[mid] = fields
			}

			addMap(seeFields, fields)
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
	},
}
