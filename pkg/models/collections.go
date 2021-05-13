// Code generated with go generate DO NOT EDIT.
package models

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type Organisation struct {
	ID                         int             `json:"id"`
	Name                       string          `json:"name"`
	Description                string          `json:"description"`
	LegalNotice                string          `json:"legal_notice"`
	PrivacyPolicy              string          `json:"privacy_policy"`
	LoginText                  string          `json:"login_text"`
	Theme                      string          `json:"theme"`
	CustomTranslations         json.RawMessage `json:"custom_translations"`
	ResetPasswordVerboseErrors bool            `json:"reset_password_verbose_errors"`
	EnableElectronicVoting     bool            `json:"enable_electronic_voting"`
	CommitteeIDs               []int           `json:"committee_ids"`
	ResourceIDs                []int           `json:"resource_ids"`
	OrganisationTagIDs         []int           `json:"organisation_tag_ids"`
}

func LoadOrganisation(ctx context.Context, ds Getter, id int) (*Organisation, error) {
	fields := []string{
		"id",
		"name",
		"description",
		"legal_notice",
		"privacy_policy",
		"login_text",
		"theme",
		"custom_translations",
		"reset_password_verbose_errors",
		"enable_electronic_voting",
		"committee_ids",
		"resource_ids",
		"organisation_tag_ids",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("organisation/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Organisation
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Description); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.LegalNotice); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.PrivacyPolicy); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.LoginText); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.Theme); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.CustomTranslations); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.ResetPasswordVerboseErrors); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.EnableElectronicVoting); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.CommitteeIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.ResourceIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.OrganisationTagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	return &c, nil
}

type User struct {
	ID                          int               `json:"id"`
	Username                    string            `json:"username"`
	Title                       string            `json:"title"`
	FirstName                   string            `json:"first_name"`
	LastName                    string            `json:"last_name"`
	IsActive                    bool              `json:"is_active"`
	IsPhysicalPerson            bool              `json:"is_physical_person"`
	Password                    string            `json:"password"`
	DefaultPassword             string            `json:"default_password"`
	Gender                      string            `json:"gender"`
	Email                       string            `json:"email"`
	DefaultNumber               string            `json:"default_number"`
	DefaultStructureLevel       string            `json:"default_structure_level"`
	DefaultVoteWeight           int               `json:"default_vote_weight"`
	LastEmailSend               int               `json:"last_email_send"`
	IsDemoUser                  bool              `json:"is_demo_user"`
	OrganisationManagementLevel string            `json:"organisation_management_level"`
	IsPresentInMeetingIDs       []int             `json:"is_present_in_meeting_ids"`
	CommitteeIDs                []int             `json:"committee_ids"`
	CommitteeManagementLevel    map[string]string `json:"committee_$_management_level"`
	Comment                     map[string]string `json:"comment_$"`
	Number                      map[string]string `json:"number_$"`
	StructureLevel              map[string]string `json:"structure_level_$"`
	AboutMe                     map[string]string `json:"about_me_$"`
	VoteWeight                  map[string]int    `json:"vote_weight_$"`
	GroupIDs                    map[string][]int  `json:"group_$_ids"`
	SpeakerIDs                  map[string][]int  `json:"speaker_$_ids"`
	PersonalNoteIDs             map[string][]int  `json:"personal_note_$_ids"`
	SupportedMotionIDs          map[string][]int  `json:"supported_motion_$_ids"`
	SubmittedMotionIDs          map[string][]int  `json:"submitted_motion_$_ids"`
	PollVotedIDs                map[string][]int  `json:"poll_voted_$_ids"`
	OptionIDs                   map[string][]int  `json:"option_$_ids"`
	VoteIDs                     map[string][]int  `json:"vote_$_ids"`
	VoteDelegatedVoteIDs        map[string][]int  `json:"vote_delegated_vote_$_ids"`
	AssignmentCandidateIDs      map[string][]int  `json:"assignment_candidate_$_ids"`
	ProjectionIDs               map[string][]int  `json:"projection_$_ids"`
	VoteDelegatedToID           map[string]int    `json:"vote_delegated_$_to_id"`
	VoteDelegationsFromIDs      map[string][]int  `json:"vote_delegations_$_from_ids"`
	MeetingIDs                  []int             `json:"meeting_ids"`
}

func LoadUser(ctx context.Context, ds Getter, id int) (*User, error) {
	fields := []string{
		"id",
		"username",
		"title",
		"first_name",
		"last_name",
		"is_active",
		"is_physical_person",
		"password",
		"default_password",
		"gender",
		"email",
		"default_number",
		"default_structure_level",
		"default_vote_weight",
		"last_email_send",
		"is_demo_user",
		"organisation_management_level",
		"is_present_in_meeting_ids",
		"committee_ids",
		"committee_$_management_level",
		"comment_$",
		"number_$",
		"structure_level_$",
		"about_me_$",
		"vote_weight_$",
		"group_$_ids",
		"speaker_$_ids",
		"personal_note_$_ids",
		"supported_motion_$_ids",
		"submitted_motion_$_ids",
		"poll_voted_$_ids",
		"option_$_ids",
		"vote_$_ids",
		"vote_delegated_vote_$_ids",
		"assignment_candidate_$_ids",
		"projection_$_ids",
		"vote_delegated_$_to_id",
		"vote_delegations_$_from_ids",
		"meeting_ids",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("user/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c User
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Username); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.FirstName); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.LastName); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.IsActive); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.IsPhysicalPerson); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.Password); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.DefaultPassword); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.Gender); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.Email); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.DefaultNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.DefaultStructureLevel); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.DefaultVoteWeight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.LastEmailSend); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.IsDemoUser); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	if values[16] != nil {
		if err := json.Unmarshal(values[16], &c.OrganisationManagementLevel); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[16], err)
		}
	}
	if values[17] != nil {
		if err := json.Unmarshal(values[17], &c.IsPresentInMeetingIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[17], err)
		}
	}
	if values[18] != nil {
		if err := json.Unmarshal(values[18], &c.CommitteeIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[18], err)
		}
	}
	if values[19] != nil {
		var repl []string
		if err := json.Unmarshal(values[19], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[19], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[19], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[19], err)
		}

		data := make(map[string]string, len(repl))
		for i, r := range repl {
			var decoded string
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.CommitteeManagementLevel = data
	}
	if values[20] != nil {
		var repl []string
		if err := json.Unmarshal(values[20], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[20], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[20], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[20], err)
		}

		data := make(map[string]string, len(repl))
		for i, r := range repl {
			var decoded string
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.Comment = data
	}
	if values[21] != nil {
		var repl []string
		if err := json.Unmarshal(values[21], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[21], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[21], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[21], err)
		}

		data := make(map[string]string, len(repl))
		for i, r := range repl {
			var decoded string
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.Number = data
	}
	if values[22] != nil {
		var repl []string
		if err := json.Unmarshal(values[22], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[22], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[22], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[22], err)
		}

		data := make(map[string]string, len(repl))
		for i, r := range repl {
			var decoded string
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.StructureLevel = data
	}
	if values[23] != nil {
		var repl []string
		if err := json.Unmarshal(values[23], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[23], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[23], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[23], err)
		}

		data := make(map[string]string, len(repl))
		for i, r := range repl {
			var decoded string
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.AboutMe = data
	}
	if values[24] != nil {
		var repl []string
		if err := json.Unmarshal(values[24], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[24], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[24], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[24], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.VoteWeight = data
	}
	if values[25] != nil {
		var repl []string
		if err := json.Unmarshal(values[25], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[25], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[25], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[25], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.GroupIDs = data
	}
	if values[26] != nil {
		var repl []string
		if err := json.Unmarshal(values[26], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[26], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[26], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[26], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.SpeakerIDs = data
	}
	if values[27] != nil {
		var repl []string
		if err := json.Unmarshal(values[27], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[27], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[27], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[27], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.PersonalNoteIDs = data
	}
	if values[28] != nil {
		var repl []string
		if err := json.Unmarshal(values[28], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[28], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[28], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[28], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.SupportedMotionIDs = data
	}
	if values[29] != nil {
		var repl []string
		if err := json.Unmarshal(values[29], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[29], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[29], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[29], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.SubmittedMotionIDs = data
	}
	if values[30] != nil {
		var repl []string
		if err := json.Unmarshal(values[30], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[30], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[30], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[30], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.PollVotedIDs = data
	}
	if values[31] != nil {
		var repl []string
		if err := json.Unmarshal(values[31], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[31], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[31], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[31], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.OptionIDs = data
	}
	if values[32] != nil {
		var repl []string
		if err := json.Unmarshal(values[32], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[32], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[32], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[32], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.VoteIDs = data
	}
	if values[33] != nil {
		var repl []string
		if err := json.Unmarshal(values[33], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[33], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[33], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[33], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.VoteDelegatedVoteIDs = data
	}
	if values[34] != nil {
		var repl []string
		if err := json.Unmarshal(values[34], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[34], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[34], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[34], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.AssignmentCandidateIDs = data
	}
	if values[35] != nil {
		var repl []string
		if err := json.Unmarshal(values[35], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[35], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[35], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[35], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.ProjectionIDs = data
	}
	if values[36] != nil {
		var repl []string
		if err := json.Unmarshal(values[36], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[36], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[36], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[36], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.VoteDelegatedToID = data
	}
	if values[37] != nil {
		var repl []string
		if err := json.Unmarshal(values[37], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[37], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[37], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[37], err)
		}

		data := make(map[string][]int, len(repl))
		for i, r := range repl {
			var decoded []int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.VoteDelegationsFromIDs = data
	}
	if values[38] != nil {
		if err := json.Unmarshal(values[38], &c.MeetingIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[38], err)
		}
	}
	return &c, nil
}

type Resource struct {
	ID             int    `json:"id"`
	Token          string `json:"token"`
	Filesize       int    `json:"filesize"`
	Mimetype       string `json:"mimetype"`
	OrganisationID int    `json:"organisation_id"`
}

func LoadResource(ctx context.Context, ds Getter, id int) (*Resource, error) {
	fields := []string{
		"id",
		"token",
		"filesize",
		"mimetype",
		"organisation_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("resource/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Resource
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Token); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Filesize); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Mimetype); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.OrganisationID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	return &c, nil
}

type OrganisationTag struct {
	ID             int      `json:"id"`
	Name           string   `json:"name"`
	Color          string   `json:"color"`
	TaggedIDs      []string `json:"tagged_ids"`
	OrganisationID int      `json:"organisation_id"`
}

func LoadOrganisationTag(ctx context.Context, ds Getter, id int) (*OrganisationTag, error) {
	fields := []string{
		"id",
		"name",
		"color",
		"tagged_ids",
		"organisation_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("organisation_tag/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c OrganisationTag
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Color); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.TaggedIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.OrganisationID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	return &c, nil
}

type Committee struct {
	ID                                 int    `json:"id"`
	Name                               string `json:"name"`
	Description                        string `json:"description"`
	MeetingIDs                         []int  `json:"meeting_ids"`
	TemplateMeetingID                  int    `json:"template_meeting_id"`
	DefaultMeetingID                   int    `json:"default_meeting_id"`
	UserIDs                            []int  `json:"user_ids"`
	ForwardToCommitteeIDs              []int  `json:"forward_to_committee_ids"`
	ReceiveForwardingsFromCommitteeIDs []int  `json:"receive_forwardings_from_committee_ids"`
	OrganisationTagIDs                 []int  `json:"organisation_tag_ids"`
	OrganisationID                     int    `json:"organisation_id"`
}

func LoadCommittee(ctx context.Context, ds Getter, id int) (*Committee, error) {
	fields := []string{
		"id",
		"name",
		"description",
		"meeting_ids",
		"template_meeting_id",
		"default_meeting_id",
		"user_ids",
		"forward_to_committee_ids",
		"receive_forwardings_from_committee_ids",
		"organisation_tag_ids",
		"organisation_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("committee/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Committee
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Description); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.MeetingIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.TemplateMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.DefaultMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.UserIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.ForwardToCommitteeIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.ReceiveForwardingsFromCommitteeIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.OrganisationTagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.OrganisationID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	return &c, nil
}

type Meeting struct {
	ID                                          int            `json:"id"`
	WelcomeTitle                                string         `json:"welcome_title"`
	WelcomeText                                 string         `json:"welcome_text"`
	Name                                        string         `json:"name"`
	Description                                 string         `json:"description"`
	Location                                    string         `json:"location"`
	StartTime                                   int            `json:"start_time"`
	EndTime                                     int            `json:"end_time"`
	JitsiDomain                                 string         `json:"jitsi_domain"`
	JitsiRoomName                               string         `json:"jitsi_room_name"`
	JitsiRoomPassword                           string         `json:"jitsi_room_password"`
	UrlName                                     string         `json:"url_name"`
	TemplateForCommitteeID                      int            `json:"template_for_committee_id"`
	EnableAnonymous                             bool           `json:"enable_anonymous"`
	ConferenceShow                              bool           `json:"conference_show"`
	ConferenceAutoConnect                       bool           `json:"conference_auto_connect"`
	ConferenceLosRestriction                    bool           `json:"conference_los_restriction"`
	ConferenceStreamUrl                         string         `json:"conference_stream_url"`
	ConferenceStreamPosterUrl                   string         `json:"conference_stream_poster_url"`
	ConferenceOpenMicrophone                    bool           `json:"conference_open_microphone"`
	ConferenceOpenVideo                         bool           `json:"conference_open_video"`
	ConferenceAutoConnectNextSpeakers           int            `json:"conference_auto_connect_next_speakers"`
	ProjectorCountdownDefaultTime               int            `json:"projector_countdown_default_time"`
	ProjectorCountdownWarningTime               int            `json:"projector_countdown_warning_time"`
	ExportCsvEncoding                           string         `json:"export_csv_encoding"`
	ExportCsvSeparator                          string         `json:"export_csv_separator"`
	ExportPdfPagenumberAlignment                string         `json:"export_pdf_pagenumber_alignment"`
	ExportPdfFontsize                           int            `json:"export_pdf_fontsize"`
	ExportPdfPagesize                           string         `json:"export_pdf_pagesize"`
	AgendaShowSubtitles                         bool           `json:"agenda_show_subtitles"`
	AgendaEnableNumbering                       bool           `json:"agenda_enable_numbering"`
	AgendaNumberPrefix                          string         `json:"agenda_number_prefix"`
	AgendaNumeralSystem                         string         `json:"agenda_numeral_system"`
	AgendaItemCreation                          string         `json:"agenda_item_creation"`
	AgendaNewItemsDefaultVisibility             string         `json:"agenda_new_items_default_visibility"`
	AgendaShowInternalItemsOnProjector          bool           `json:"agenda_show_internal_items_on_projector"`
	ListOfSpeakersAmountLastOnProjector         int            `json:"list_of_speakers_amount_last_on_projector"`
	ListOfSpeakersAmountNextOnProjector         int            `json:"list_of_speakers_amount_next_on_projector"`
	ListOfSpeakersCoupleCountdown               bool           `json:"list_of_speakers_couple_countdown"`
	ListOfSpeakersShowAmountOfSpeakersOnSlide   bool           `json:"list_of_speakers_show_amount_of_speakers_on_slide"`
	ListOfSpeakersPresentUsersOnly              bool           `json:"list_of_speakers_present_users_only"`
	ListOfSpeakersShowFirstContribution         bool           `json:"list_of_speakers_show_first_contribution"`
	ListOfSpeakersEnablePointOfOrderSpeakers    bool           `json:"list_of_speakers_enable_point_of_order_speakers"`
	ListOfSpeakersEnableProContraSpeech         bool           `json:"list_of_speakers_enable_pro_contra_speech"`
	ListOfSpeakersCanSetContributionSelf        bool           `json:"list_of_speakers_can_set_contribution_self"`
	ListOfSpeakersSpeakerNoteForEveryone        bool           `json:"list_of_speakers_speaker_note_for_everyone"`
	ListOfSpeakersInitiallyClosed               bool           `json:"list_of_speakers_initially_closed"`
	MotionsDefaultWorkflowID                    int            `json:"motions_default_workflow_id"`
	MotionsDefaultAmendmentWorkflowID           int            `json:"motions_default_amendment_workflow_id"`
	MotionsDefaultStatuteAmendmentWorkflowID    int            `json:"motions_default_statute_amendment_workflow_id"`
	MotionsPreamble                             string         `json:"motions_preamble"`
	MotionsDefaultLineNumbering                 string         `json:"motions_default_line_numbering"`
	MotionsLineLength                           int            `json:"motions_line_length"`
	MotionsReasonRequired                       bool           `json:"motions_reason_required"`
	MotionsEnableTextOnProjector                bool           `json:"motions_enable_text_on_projector"`
	MotionsEnableReasonOnProjector              bool           `json:"motions_enable_reason_on_projector"`
	MotionsEnableSideboxOnProjector             bool           `json:"motions_enable_sidebox_on_projector"`
	MotionsEnableRecommendationOnProjector      bool           `json:"motions_enable_recommendation_on_projector"`
	MotionsShowReferringMotions                 bool           `json:"motions_show_referring_motions"`
	MotionsShowSequentialNumber                 bool           `json:"motions_show_sequential_number"`
	MotionsRecommendationsBy                    string         `json:"motions_recommendations_by"`
	MotionsStatuteRecommendationsBy             string         `json:"motions_statute_recommendations_by"`
	MotionsRecommendationTextMode               string         `json:"motions_recommendation_text_mode"`
	MotionsDefaultSorting                       string         `json:"motions_default_sorting"`
	MotionsNumberType                           string         `json:"motions_number_type"`
	MotionsNumberMinDigits                      int            `json:"motions_number_min_digits"`
	MotionsNumberWithBlank                      bool           `json:"motions_number_with_blank"`
	MotionsStatutesEnabled                      bool           `json:"motions_statutes_enabled"`
	MotionsAmendmentsEnabled                    bool           `json:"motions_amendments_enabled"`
	MotionsAmendmentsInMainList                 bool           `json:"motions_amendments_in_main_list"`
	MotionsAmendmentsOfAmendments               bool           `json:"motions_amendments_of_amendments"`
	MotionsAmendmentsPrefix                     string         `json:"motions_amendments_prefix"`
	MotionsAmendmentsTextMode                   string         `json:"motions_amendments_text_mode"`
	MotionsAmendmentsMultipleParagraphs         bool           `json:"motions_amendments_multiple_paragraphs"`
	MotionsSupportersMinAmount                  int            `json:"motions_supporters_min_amount"`
	MotionsExportTitle                          string         `json:"motions_export_title"`
	MotionsExportPreamble                       string         `json:"motions_export_preamble"`
	MotionsExportSubmitterRecommendation        bool           `json:"motions_export_submitter_recommendation"`
	MotionsExportFollowRecommendation           bool           `json:"motions_export_follow_recommendation"`
	MotionPollBallotPaperSelection              string         `json:"motion_poll_ballot_paper_selection"`
	MotionPollBallotPaperNumber                 int            `json:"motion_poll_ballot_paper_number"`
	MotionPollDefaultType                       string         `json:"motion_poll_default_type"`
	MotionPollDefault100PercentBase             string         `json:"motion_poll_default_100_percent_base"`
	MotionPollDefaultMajorityMethod             string         `json:"motion_poll_default_majority_method"`
	MotionPollDefaultGroupIDs                   []int          `json:"motion_poll_default_group_ids"`
	UsersSortBy                                 string         `json:"users_sort_by"`
	UsersEnablePresenceView                     bool           `json:"users_enable_presence_view"`
	UsersEnableVoteWeight                       bool           `json:"users_enable_vote_weight"`
	UsersAllowSelfSetPresent                    bool           `json:"users_allow_self_set_present"`
	UsersPdfWelcometitle                        string         `json:"users_pdf_welcometitle"`
	UsersPdfWelcometext                         string         `json:"users_pdf_welcometext"`
	UsersPdfUrl                                 string         `json:"users_pdf_url"`
	UsersPdfWlanSsid                            string         `json:"users_pdf_wlan_ssid"`
	UsersPdfWlanPassword                        string         `json:"users_pdf_wlan_password"`
	UsersPdfWlanEncryption                      string         `json:"users_pdf_wlan_encryption"`
	UsersEmailSender                            string         `json:"users_email_sender"`
	UsersEmailReplyto                           string         `json:"users_email_replyto"`
	UsersEmailSubject                           string         `json:"users_email_subject"`
	UsersEmailBody                              string         `json:"users_email_body"`
	AssignmentsExportTitle                      string         `json:"assignments_export_title"`
	AssignmentsExportPreamble                   string         `json:"assignments_export_preamble"`
	AssignmentPollBallotPaperSelection          string         `json:"assignment_poll_ballot_paper_selection"`
	AssignmentPollBallotPaperNumber             int            `json:"assignment_poll_ballot_paper_number"`
	AssignmentPollAddCandidatesToListOfSpeakers bool           `json:"assignment_poll_add_candidates_to_list_of_speakers"`
	AssignmentPollSortPollResultByVotes         bool           `json:"assignment_poll_sort_poll_result_by_votes"`
	AssignmentPollDefaultType                   string         `json:"assignment_poll_default_type"`
	AssignmentPollDefaultMethod                 string         `json:"assignment_poll_default_method"`
	AssignmentPollDefault100PercentBase         string         `json:"assignment_poll_default_100_percent_base"`
	AssignmentPollDefaultMajorityMethod         string         `json:"assignment_poll_default_majority_method"`
	AssignmentPollDefaultGroupIDs               []int          `json:"assignment_poll_default_group_ids"`
	PollBallotPaperSelection                    string         `json:"poll_ballot_paper_selection"`
	PollBallotPaperNumber                       int            `json:"poll_ballot_paper_number"`
	PollSortPollResultByVotes                   bool           `json:"poll_sort_poll_result_by_votes"`
	PollDefaultType                             string         `json:"poll_default_type"`
	PollDefaultMethod                           string         `json:"poll_default_method"`
	PollDefault100PercentBase                   string         `json:"poll_default_100_percent_base"`
	PollDefaultMajorityMethod                   string         `json:"poll_default_majority_method"`
	PollDefaultGroupIDs                         []int          `json:"poll_default_group_ids"`
	PollCoupleCountdown                         bool           `json:"poll_couple_countdown"`
	ProjectorIDs                                []int          `json:"projector_ids"`
	AllProjectionIDs                            []int          `json:"all_projection_ids"`
	ProjectorMessageIDs                         []int          `json:"projector_message_ids"`
	ProjectorCountdownIDs                       []int          `json:"projector_countdown_ids"`
	TagIDs                                      []int          `json:"tag_ids"`
	AgendaItemIDs                               []int          `json:"agenda_item_ids"`
	ListOfSpeakersIDs                           []int          `json:"list_of_speakers_ids"`
	SpeakerIDs                                  []int          `json:"speaker_ids"`
	TopicIDs                                    []int          `json:"topic_ids"`
	GroupIDs                                    []int          `json:"group_ids"`
	MediafileIDs                                []int          `json:"mediafile_ids"`
	MotionIDs                                   []int          `json:"motion_ids"`
	MotionCommentSectionIDs                     []int          `json:"motion_comment_section_ids"`
	MotionCategoryIDs                           []int          `json:"motion_category_ids"`
	MotionBlockIDs                              []int          `json:"motion_block_ids"`
	MotionWorkflowIDs                           []int          `json:"motion_workflow_ids"`
	MotionStatuteParagraphIDs                   []int          `json:"motion_statute_paragraph_ids"`
	MotionCommentIDs                            []int          `json:"motion_comment_ids"`
	MotionSubmitterIDs                          []int          `json:"motion_submitter_ids"`
	MotionChangeRecommendationIDs               []int          `json:"motion_change_recommendation_ids"`
	MotionStateIDs                              []int          `json:"motion_state_ids"`
	PollIDs                                     []int          `json:"poll_ids"`
	OptionIDs                                   []int          `json:"option_ids"`
	VoteIDs                                     []int          `json:"vote_ids"`
	AssignmentIDs                               []int          `json:"assignment_ids"`
	AssignmentCandidateIDs                      []int          `json:"assignment_candidate_ids"`
	PersonalNoteIDs                             []int          `json:"personal_note_ids"`
	LogoID                                      map[string]int `json:"logo_$_id"`
	FontID                                      map[string]int `json:"font_$_id"`
	CommitteeID                                 int            `json:"committee_id"`
	DefaultMeetingForCommitteeID                int            `json:"default_meeting_for_committee_id"`
	OrganisationTagIDs                          []int          `json:"organisation_tag_ids"`
	PresentUserIDs                              []int          `json:"present_user_ids"`
	UserIDs                                     []int          `json:"user_ids"`
	ReferenceProjectorID                        int            `json:"reference_projector_id"`
	ListOfSpeakersCountdownID                   int            `json:"list_of_speakers_countdown_id"`
	PollCountdownID                             int            `json:"poll_countdown_id"`
	DefaultProjectorID                          map[string]int `json:"default_projector_$_id"`
	ProjectionIDs                               []int          `json:"projection_ids"`
	DefaultGroupID                              int            `json:"default_group_id"`
	AdminGroupID                                int            `json:"admin_group_id"`
}

func LoadMeeting(ctx context.Context, ds Getter, id int) (*Meeting, error) {
	fields := []string{
		"id",
		"welcome_title",
		"welcome_text",
		"name",
		"description",
		"location",
		"start_time",
		"end_time",
		"jitsi_domain",
		"jitsi_room_name",
		"jitsi_room_password",
		"url_name",
		"template_for_committee_id",
		"enable_anonymous",
		"conference_show",
		"conference_auto_connect",
		"conference_los_restriction",
		"conference_stream_url",
		"conference_stream_poster_url",
		"conference_open_microphone",
		"conference_open_video",
		"conference_auto_connect_next_speakers",
		"projector_countdown_default_time",
		"projector_countdown_warning_time",
		"export_csv_encoding",
		"export_csv_separator",
		"export_pdf_pagenumber_alignment",
		"export_pdf_fontsize",
		"export_pdf_pagesize",
		"agenda_show_subtitles",
		"agenda_enable_numbering",
		"agenda_number_prefix",
		"agenda_numeral_system",
		"agenda_item_creation",
		"agenda_new_items_default_visibility",
		"agenda_show_internal_items_on_projector",
		"list_of_speakers_amount_last_on_projector",
		"list_of_speakers_amount_next_on_projector",
		"list_of_speakers_couple_countdown",
		"list_of_speakers_show_amount_of_speakers_on_slide",
		"list_of_speakers_present_users_only",
		"list_of_speakers_show_first_contribution",
		"list_of_speakers_enable_point_of_order_speakers",
		"list_of_speakers_enable_pro_contra_speech",
		"list_of_speakers_can_set_contribution_self",
		"list_of_speakers_speaker_note_for_everyone",
		"list_of_speakers_initially_closed",
		"motions_default_workflow_id",
		"motions_default_amendment_workflow_id",
		"motions_default_statute_amendment_workflow_id",
		"motions_preamble",
		"motions_default_line_numbering",
		"motions_line_length",
		"motions_reason_required",
		"motions_enable_text_on_projector",
		"motions_enable_reason_on_projector",
		"motions_enable_sidebox_on_projector",
		"motions_enable_recommendation_on_projector",
		"motions_show_referring_motions",
		"motions_show_sequential_number",
		"motions_recommendations_by",
		"motions_statute_recommendations_by",
		"motions_recommendation_text_mode",
		"motions_default_sorting",
		"motions_number_type",
		"motions_number_min_digits",
		"motions_number_with_blank",
		"motions_statutes_enabled",
		"motions_amendments_enabled",
		"motions_amendments_in_main_list",
		"motions_amendments_of_amendments",
		"motions_amendments_prefix",
		"motions_amendments_text_mode",
		"motions_amendments_multiple_paragraphs",
		"motions_supporters_min_amount",
		"motions_export_title",
		"motions_export_preamble",
		"motions_export_submitter_recommendation",
		"motions_export_follow_recommendation",
		"motion_poll_ballot_paper_selection",
		"motion_poll_ballot_paper_number",
		"motion_poll_default_type",
		"motion_poll_default_100_percent_base",
		"motion_poll_default_majority_method",
		"motion_poll_default_group_ids",
		"users_sort_by",
		"users_enable_presence_view",
		"users_enable_vote_weight",
		"users_allow_self_set_present",
		"users_pdf_welcometitle",
		"users_pdf_welcometext",
		"users_pdf_url",
		"users_pdf_wlan_ssid",
		"users_pdf_wlan_password",
		"users_pdf_wlan_encryption",
		"users_email_sender",
		"users_email_replyto",
		"users_email_subject",
		"users_email_body",
		"assignments_export_title",
		"assignments_export_preamble",
		"assignment_poll_ballot_paper_selection",
		"assignment_poll_ballot_paper_number",
		"assignment_poll_add_candidates_to_list_of_speakers",
		"assignment_poll_sort_poll_result_by_votes",
		"assignment_poll_default_type",
		"assignment_poll_default_method",
		"assignment_poll_default_100_percent_base",
		"assignment_poll_default_majority_method",
		"assignment_poll_default_group_ids",
		"poll_ballot_paper_selection",
		"poll_ballot_paper_number",
		"poll_sort_poll_result_by_votes",
		"poll_default_type",
		"poll_default_method",
		"poll_default_100_percent_base",
		"poll_default_majority_method",
		"poll_default_group_ids",
		"poll_couple_countdown",
		"projector_ids",
		"all_projection_ids",
		"projector_message_ids",
		"projector_countdown_ids",
		"tag_ids",
		"agenda_item_ids",
		"list_of_speakers_ids",
		"speaker_ids",
		"topic_ids",
		"group_ids",
		"mediafile_ids",
		"motion_ids",
		"motion_comment_section_ids",
		"motion_category_ids",
		"motion_block_ids",
		"motion_workflow_ids",
		"motion_statute_paragraph_ids",
		"motion_comment_ids",
		"motion_submitter_ids",
		"motion_change_recommendation_ids",
		"motion_state_ids",
		"poll_ids",
		"option_ids",
		"vote_ids",
		"assignment_ids",
		"assignment_candidate_ids",
		"personal_note_ids",
		"logo_$_id",
		"font_$_id",
		"committee_id",
		"default_meeting_for_committee_id",
		"organisation_tag_ids",
		"present_user_ids",
		"user_ids",
		"reference_projector_id",
		"list_of_speakers_countdown_id",
		"poll_countdown_id",
		"default_projector_$_id",
		"projection_ids",
		"default_group_id",
		"admin_group_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("meeting/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Meeting
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.WelcomeTitle); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.WelcomeText); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Description); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.Location); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.StartTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.EndTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.JitsiDomain); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.JitsiRoomName); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.JitsiRoomPassword); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.UrlName); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.TemplateForCommitteeID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.EnableAnonymous); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.ConferenceShow); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.ConferenceAutoConnect); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	if values[16] != nil {
		if err := json.Unmarshal(values[16], &c.ConferenceLosRestriction); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[16], err)
		}
	}
	if values[17] != nil {
		if err := json.Unmarshal(values[17], &c.ConferenceStreamUrl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[17], err)
		}
	}
	if values[18] != nil {
		if err := json.Unmarshal(values[18], &c.ConferenceStreamPosterUrl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[18], err)
		}
	}
	if values[19] != nil {
		if err := json.Unmarshal(values[19], &c.ConferenceOpenMicrophone); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[19], err)
		}
	}
	if values[20] != nil {
		if err := json.Unmarshal(values[20], &c.ConferenceOpenVideo); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[20], err)
		}
	}
	if values[21] != nil {
		if err := json.Unmarshal(values[21], &c.ConferenceAutoConnectNextSpeakers); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[21], err)
		}
	}
	if values[22] != nil {
		if err := json.Unmarshal(values[22], &c.ProjectorCountdownDefaultTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[22], err)
		}
	}
	if values[23] != nil {
		if err := json.Unmarshal(values[23], &c.ProjectorCountdownWarningTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[23], err)
		}
	}
	if values[24] != nil {
		if err := json.Unmarshal(values[24], &c.ExportCsvEncoding); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[24], err)
		}
	}
	if values[25] != nil {
		if err := json.Unmarshal(values[25], &c.ExportCsvSeparator); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[25], err)
		}
	}
	if values[26] != nil {
		if err := json.Unmarshal(values[26], &c.ExportPdfPagenumberAlignment); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[26], err)
		}
	}
	if values[27] != nil {
		if err := json.Unmarshal(values[27], &c.ExportPdfFontsize); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[27], err)
		}
	}
	if values[28] != nil {
		if err := json.Unmarshal(values[28], &c.ExportPdfPagesize); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[28], err)
		}
	}
	if values[29] != nil {
		if err := json.Unmarshal(values[29], &c.AgendaShowSubtitles); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[29], err)
		}
	}
	if values[30] != nil {
		if err := json.Unmarshal(values[30], &c.AgendaEnableNumbering); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[30], err)
		}
	}
	if values[31] != nil {
		if err := json.Unmarshal(values[31], &c.AgendaNumberPrefix); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[31], err)
		}
	}
	if values[32] != nil {
		if err := json.Unmarshal(values[32], &c.AgendaNumeralSystem); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[32], err)
		}
	}
	if values[33] != nil {
		if err := json.Unmarshal(values[33], &c.AgendaItemCreation); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[33], err)
		}
	}
	if values[34] != nil {
		if err := json.Unmarshal(values[34], &c.AgendaNewItemsDefaultVisibility); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[34], err)
		}
	}
	if values[35] != nil {
		if err := json.Unmarshal(values[35], &c.AgendaShowInternalItemsOnProjector); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[35], err)
		}
	}
	if values[36] != nil {
		if err := json.Unmarshal(values[36], &c.ListOfSpeakersAmountLastOnProjector); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[36], err)
		}
	}
	if values[37] != nil {
		if err := json.Unmarshal(values[37], &c.ListOfSpeakersAmountNextOnProjector); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[37], err)
		}
	}
	if values[38] != nil {
		if err := json.Unmarshal(values[38], &c.ListOfSpeakersCoupleCountdown); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[38], err)
		}
	}
	if values[39] != nil {
		if err := json.Unmarshal(values[39], &c.ListOfSpeakersShowAmountOfSpeakersOnSlide); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[39], err)
		}
	}
	if values[40] != nil {
		if err := json.Unmarshal(values[40], &c.ListOfSpeakersPresentUsersOnly); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[40], err)
		}
	}
	if values[41] != nil {
		if err := json.Unmarshal(values[41], &c.ListOfSpeakersShowFirstContribution); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[41], err)
		}
	}
	if values[42] != nil {
		if err := json.Unmarshal(values[42], &c.ListOfSpeakersEnablePointOfOrderSpeakers); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[42], err)
		}
	}
	if values[43] != nil {
		if err := json.Unmarshal(values[43], &c.ListOfSpeakersEnableProContraSpeech); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[43], err)
		}
	}
	if values[44] != nil {
		if err := json.Unmarshal(values[44], &c.ListOfSpeakersCanSetContributionSelf); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[44], err)
		}
	}
	if values[45] != nil {
		if err := json.Unmarshal(values[45], &c.ListOfSpeakersSpeakerNoteForEveryone); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[45], err)
		}
	}
	if values[46] != nil {
		if err := json.Unmarshal(values[46], &c.ListOfSpeakersInitiallyClosed); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[46], err)
		}
	}
	if values[47] != nil {
		if err := json.Unmarshal(values[47], &c.MotionsDefaultWorkflowID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[47], err)
		}
	}
	if values[48] != nil {
		if err := json.Unmarshal(values[48], &c.MotionsDefaultAmendmentWorkflowID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[48], err)
		}
	}
	if values[49] != nil {
		if err := json.Unmarshal(values[49], &c.MotionsDefaultStatuteAmendmentWorkflowID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[49], err)
		}
	}
	if values[50] != nil {
		if err := json.Unmarshal(values[50], &c.MotionsPreamble); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[50], err)
		}
	}
	if values[51] != nil {
		if err := json.Unmarshal(values[51], &c.MotionsDefaultLineNumbering); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[51], err)
		}
	}
	if values[52] != nil {
		if err := json.Unmarshal(values[52], &c.MotionsLineLength); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[52], err)
		}
	}
	if values[53] != nil {
		if err := json.Unmarshal(values[53], &c.MotionsReasonRequired); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[53], err)
		}
	}
	if values[54] != nil {
		if err := json.Unmarshal(values[54], &c.MotionsEnableTextOnProjector); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[54], err)
		}
	}
	if values[55] != nil {
		if err := json.Unmarshal(values[55], &c.MotionsEnableReasonOnProjector); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[55], err)
		}
	}
	if values[56] != nil {
		if err := json.Unmarshal(values[56], &c.MotionsEnableSideboxOnProjector); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[56], err)
		}
	}
	if values[57] != nil {
		if err := json.Unmarshal(values[57], &c.MotionsEnableRecommendationOnProjector); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[57], err)
		}
	}
	if values[58] != nil {
		if err := json.Unmarshal(values[58], &c.MotionsShowReferringMotions); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[58], err)
		}
	}
	if values[59] != nil {
		if err := json.Unmarshal(values[59], &c.MotionsShowSequentialNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[59], err)
		}
	}
	if values[60] != nil {
		if err := json.Unmarshal(values[60], &c.MotionsRecommendationsBy); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[60], err)
		}
	}
	if values[61] != nil {
		if err := json.Unmarshal(values[61], &c.MotionsStatuteRecommendationsBy); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[61], err)
		}
	}
	if values[62] != nil {
		if err := json.Unmarshal(values[62], &c.MotionsRecommendationTextMode); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[62], err)
		}
	}
	if values[63] != nil {
		if err := json.Unmarshal(values[63], &c.MotionsDefaultSorting); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[63], err)
		}
	}
	if values[64] != nil {
		if err := json.Unmarshal(values[64], &c.MotionsNumberType); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[64], err)
		}
	}
	if values[65] != nil {
		if err := json.Unmarshal(values[65], &c.MotionsNumberMinDigits); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[65], err)
		}
	}
	if values[66] != nil {
		if err := json.Unmarshal(values[66], &c.MotionsNumberWithBlank); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[66], err)
		}
	}
	if values[67] != nil {
		if err := json.Unmarshal(values[67], &c.MotionsStatutesEnabled); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[67], err)
		}
	}
	if values[68] != nil {
		if err := json.Unmarshal(values[68], &c.MotionsAmendmentsEnabled); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[68], err)
		}
	}
	if values[69] != nil {
		if err := json.Unmarshal(values[69], &c.MotionsAmendmentsInMainList); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[69], err)
		}
	}
	if values[70] != nil {
		if err := json.Unmarshal(values[70], &c.MotionsAmendmentsOfAmendments); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[70], err)
		}
	}
	if values[71] != nil {
		if err := json.Unmarshal(values[71], &c.MotionsAmendmentsPrefix); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[71], err)
		}
	}
	if values[72] != nil {
		if err := json.Unmarshal(values[72], &c.MotionsAmendmentsTextMode); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[72], err)
		}
	}
	if values[73] != nil {
		if err := json.Unmarshal(values[73], &c.MotionsAmendmentsMultipleParagraphs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[73], err)
		}
	}
	if values[74] != nil {
		if err := json.Unmarshal(values[74], &c.MotionsSupportersMinAmount); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[74], err)
		}
	}
	if values[75] != nil {
		if err := json.Unmarshal(values[75], &c.MotionsExportTitle); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[75], err)
		}
	}
	if values[76] != nil {
		if err := json.Unmarshal(values[76], &c.MotionsExportPreamble); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[76], err)
		}
	}
	if values[77] != nil {
		if err := json.Unmarshal(values[77], &c.MotionsExportSubmitterRecommendation); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[77], err)
		}
	}
	if values[78] != nil {
		if err := json.Unmarshal(values[78], &c.MotionsExportFollowRecommendation); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[78], err)
		}
	}
	if values[79] != nil {
		if err := json.Unmarshal(values[79], &c.MotionPollBallotPaperSelection); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[79], err)
		}
	}
	if values[80] != nil {
		if err := json.Unmarshal(values[80], &c.MotionPollBallotPaperNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[80], err)
		}
	}
	if values[81] != nil {
		if err := json.Unmarshal(values[81], &c.MotionPollDefaultType); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[81], err)
		}
	}
	if values[82] != nil {
		if err := json.Unmarshal(values[82], &c.MotionPollDefault100PercentBase); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[82], err)
		}
	}
	if values[83] != nil {
		if err := json.Unmarshal(values[83], &c.MotionPollDefaultMajorityMethod); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[83], err)
		}
	}
	if values[84] != nil {
		if err := json.Unmarshal(values[84], &c.MotionPollDefaultGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[84], err)
		}
	}
	if values[85] != nil {
		if err := json.Unmarshal(values[85], &c.UsersSortBy); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[85], err)
		}
	}
	if values[86] != nil {
		if err := json.Unmarshal(values[86], &c.UsersEnablePresenceView); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[86], err)
		}
	}
	if values[87] != nil {
		if err := json.Unmarshal(values[87], &c.UsersEnableVoteWeight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[87], err)
		}
	}
	if values[88] != nil {
		if err := json.Unmarshal(values[88], &c.UsersAllowSelfSetPresent); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[88], err)
		}
	}
	if values[89] != nil {
		if err := json.Unmarshal(values[89], &c.UsersPdfWelcometitle); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[89], err)
		}
	}
	if values[90] != nil {
		if err := json.Unmarshal(values[90], &c.UsersPdfWelcometext); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[90], err)
		}
	}
	if values[91] != nil {
		if err := json.Unmarshal(values[91], &c.UsersPdfUrl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[91], err)
		}
	}
	if values[92] != nil {
		if err := json.Unmarshal(values[92], &c.UsersPdfWlanSsid); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[92], err)
		}
	}
	if values[93] != nil {
		if err := json.Unmarshal(values[93], &c.UsersPdfWlanPassword); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[93], err)
		}
	}
	if values[94] != nil {
		if err := json.Unmarshal(values[94], &c.UsersPdfWlanEncryption); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[94], err)
		}
	}
	if values[95] != nil {
		if err := json.Unmarshal(values[95], &c.UsersEmailSender); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[95], err)
		}
	}
	if values[96] != nil {
		if err := json.Unmarshal(values[96], &c.UsersEmailReplyto); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[96], err)
		}
	}
	if values[97] != nil {
		if err := json.Unmarshal(values[97], &c.UsersEmailSubject); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[97], err)
		}
	}
	if values[98] != nil {
		if err := json.Unmarshal(values[98], &c.UsersEmailBody); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[98], err)
		}
	}
	if values[99] != nil {
		if err := json.Unmarshal(values[99], &c.AssignmentsExportTitle); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[99], err)
		}
	}
	if values[100] != nil {
		if err := json.Unmarshal(values[100], &c.AssignmentsExportPreamble); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[100], err)
		}
	}
	if values[101] != nil {
		if err := json.Unmarshal(values[101], &c.AssignmentPollBallotPaperSelection); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[101], err)
		}
	}
	if values[102] != nil {
		if err := json.Unmarshal(values[102], &c.AssignmentPollBallotPaperNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[102], err)
		}
	}
	if values[103] != nil {
		if err := json.Unmarshal(values[103], &c.AssignmentPollAddCandidatesToListOfSpeakers); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[103], err)
		}
	}
	if values[104] != nil {
		if err := json.Unmarshal(values[104], &c.AssignmentPollSortPollResultByVotes); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[104], err)
		}
	}
	if values[105] != nil {
		if err := json.Unmarshal(values[105], &c.AssignmentPollDefaultType); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[105], err)
		}
	}
	if values[106] != nil {
		if err := json.Unmarshal(values[106], &c.AssignmentPollDefaultMethod); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[106], err)
		}
	}
	if values[107] != nil {
		if err := json.Unmarshal(values[107], &c.AssignmentPollDefault100PercentBase); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[107], err)
		}
	}
	if values[108] != nil {
		if err := json.Unmarshal(values[108], &c.AssignmentPollDefaultMajorityMethod); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[108], err)
		}
	}
	if values[109] != nil {
		if err := json.Unmarshal(values[109], &c.AssignmentPollDefaultGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[109], err)
		}
	}
	if values[110] != nil {
		if err := json.Unmarshal(values[110], &c.PollBallotPaperSelection); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[110], err)
		}
	}
	if values[111] != nil {
		if err := json.Unmarshal(values[111], &c.PollBallotPaperNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[111], err)
		}
	}
	if values[112] != nil {
		if err := json.Unmarshal(values[112], &c.PollSortPollResultByVotes); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[112], err)
		}
	}
	if values[113] != nil {
		if err := json.Unmarshal(values[113], &c.PollDefaultType); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[113], err)
		}
	}
	if values[114] != nil {
		if err := json.Unmarshal(values[114], &c.PollDefaultMethod); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[114], err)
		}
	}
	if values[115] != nil {
		if err := json.Unmarshal(values[115], &c.PollDefault100PercentBase); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[115], err)
		}
	}
	if values[116] != nil {
		if err := json.Unmarshal(values[116], &c.PollDefaultMajorityMethod); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[116], err)
		}
	}
	if values[117] != nil {
		if err := json.Unmarshal(values[117], &c.PollDefaultGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[117], err)
		}
	}
	if values[118] != nil {
		if err := json.Unmarshal(values[118], &c.PollCoupleCountdown); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[118], err)
		}
	}
	if values[119] != nil {
		if err := json.Unmarshal(values[119], &c.ProjectorIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[119], err)
		}
	}
	if values[120] != nil {
		if err := json.Unmarshal(values[120], &c.AllProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[120], err)
		}
	}
	if values[121] != nil {
		if err := json.Unmarshal(values[121], &c.ProjectorMessageIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[121], err)
		}
	}
	if values[122] != nil {
		if err := json.Unmarshal(values[122], &c.ProjectorCountdownIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[122], err)
		}
	}
	if values[123] != nil {
		if err := json.Unmarshal(values[123], &c.TagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[123], err)
		}
	}
	if values[124] != nil {
		if err := json.Unmarshal(values[124], &c.AgendaItemIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[124], err)
		}
	}
	if values[125] != nil {
		if err := json.Unmarshal(values[125], &c.ListOfSpeakersIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[125], err)
		}
	}
	if values[126] != nil {
		if err := json.Unmarshal(values[126], &c.SpeakerIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[126], err)
		}
	}
	if values[127] != nil {
		if err := json.Unmarshal(values[127], &c.TopicIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[127], err)
		}
	}
	if values[128] != nil {
		if err := json.Unmarshal(values[128], &c.GroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[128], err)
		}
	}
	if values[129] != nil {
		if err := json.Unmarshal(values[129], &c.MediafileIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[129], err)
		}
	}
	if values[130] != nil {
		if err := json.Unmarshal(values[130], &c.MotionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[130], err)
		}
	}
	if values[131] != nil {
		if err := json.Unmarshal(values[131], &c.MotionCommentSectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[131], err)
		}
	}
	if values[132] != nil {
		if err := json.Unmarshal(values[132], &c.MotionCategoryIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[132], err)
		}
	}
	if values[133] != nil {
		if err := json.Unmarshal(values[133], &c.MotionBlockIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[133], err)
		}
	}
	if values[134] != nil {
		if err := json.Unmarshal(values[134], &c.MotionWorkflowIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[134], err)
		}
	}
	if values[135] != nil {
		if err := json.Unmarshal(values[135], &c.MotionStatuteParagraphIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[135], err)
		}
	}
	if values[136] != nil {
		if err := json.Unmarshal(values[136], &c.MotionCommentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[136], err)
		}
	}
	if values[137] != nil {
		if err := json.Unmarshal(values[137], &c.MotionSubmitterIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[137], err)
		}
	}
	if values[138] != nil {
		if err := json.Unmarshal(values[138], &c.MotionChangeRecommendationIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[138], err)
		}
	}
	if values[139] != nil {
		if err := json.Unmarshal(values[139], &c.MotionStateIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[139], err)
		}
	}
	if values[140] != nil {
		if err := json.Unmarshal(values[140], &c.PollIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[140], err)
		}
	}
	if values[141] != nil {
		if err := json.Unmarshal(values[141], &c.OptionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[141], err)
		}
	}
	if values[142] != nil {
		if err := json.Unmarshal(values[142], &c.VoteIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[142], err)
		}
	}
	if values[143] != nil {
		if err := json.Unmarshal(values[143], &c.AssignmentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[143], err)
		}
	}
	if values[144] != nil {
		if err := json.Unmarshal(values[144], &c.AssignmentCandidateIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[144], err)
		}
	}
	if values[145] != nil {
		if err := json.Unmarshal(values[145], &c.PersonalNoteIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[145], err)
		}
	}
	if values[146] != nil {
		var repl []string
		if err := json.Unmarshal(values[146], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[146], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[146], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[146], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.LogoID = data
	}
	if values[147] != nil {
		var repl []string
		if err := json.Unmarshal(values[147], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[147], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[147], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[147], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.FontID = data
	}
	if values[148] != nil {
		if err := json.Unmarshal(values[148], &c.CommitteeID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[148], err)
		}
	}
	if values[149] != nil {
		if err := json.Unmarshal(values[149], &c.DefaultMeetingForCommitteeID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[149], err)
		}
	}
	if values[150] != nil {
		if err := json.Unmarshal(values[150], &c.OrganisationTagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[150], err)
		}
	}
	if values[151] != nil {
		if err := json.Unmarshal(values[151], &c.PresentUserIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[151], err)
		}
	}
	if values[152] != nil {
		if err := json.Unmarshal(values[152], &c.UserIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[152], err)
		}
	}
	if values[153] != nil {
		if err := json.Unmarshal(values[153], &c.ReferenceProjectorID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[153], err)
		}
	}
	if values[154] != nil {
		if err := json.Unmarshal(values[154], &c.ListOfSpeakersCountdownID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[154], err)
		}
	}
	if values[155] != nil {
		if err := json.Unmarshal(values[155], &c.PollCountdownID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[155], err)
		}
	}
	if values[156] != nil {
		var repl []string
		if err := json.Unmarshal(values[156], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[156], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[156], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[156], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.DefaultProjectorID = data
	}
	if values[157] != nil {
		if err := json.Unmarshal(values[157], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[157], err)
		}
	}
	if values[158] != nil {
		if err := json.Unmarshal(values[158], &c.DefaultGroupID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[158], err)
		}
	}
	if values[159] != nil {
		if err := json.Unmarshal(values[159], &c.AdminGroupID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[159], err)
		}
	}
	return &c, nil
}

type Group struct {
	ID                               int      `json:"id"`
	Name                             string   `json:"name"`
	Permissions                      []string `json:"permissions"`
	UserIDs                          []int    `json:"user_ids"`
	DefaultGroupForMeetingID         int      `json:"default_group_for_meeting_id"`
	AdminGroupForMeetingID           int      `json:"admin_group_for_meeting_id"`
	MediafileAccessGroupIDs          []int    `json:"mediafile_access_group_ids"`
	MediafileInheritedAccessGroupIDs []int    `json:"mediafile_inherited_access_group_ids"`
	ReadCommentSectionIDs            []int    `json:"read_comment_section_ids"`
	WriteCommentSectionIDs           []int    `json:"write_comment_section_ids"`
	PollIDs                          []int    `json:"poll_ids"`
	UsedAsMotionPollDefaultID        int      `json:"used_as_motion_poll_default_id"`
	UsedAsAssignmentPollDefaultID    int      `json:"used_as_assignment_poll_default_id"`
	UsedAsPollDefaultID              int      `json:"used_as_poll_default_id"`
	MeetingID                        int      `json:"meeting_id"`
}

func LoadGroup(ctx context.Context, ds Getter, id int) (*Group, error) {
	fields := []string{
		"id",
		"name",
		"permissions",
		"user_ids",
		"default_group_for_meeting_id",
		"admin_group_for_meeting_id",
		"mediafile_access_group_ids",
		"mediafile_inherited_access_group_ids",
		"read_comment_section_ids",
		"write_comment_section_ids",
		"poll_ids",
		"used_as_motion_poll_default_id",
		"used_as_assignment_poll_default_id",
		"used_as_poll_default_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("group/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Group
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Permissions); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.UserIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.DefaultGroupForMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.AdminGroupForMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.MediafileAccessGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.MediafileInheritedAccessGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.ReadCommentSectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.WriteCommentSectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.PollIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.UsedAsMotionPollDefaultID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.UsedAsAssignmentPollDefaultID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.UsedAsPollDefaultID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	return &c, nil
}

type PersonalNote struct {
	ID              int    `json:"id"`
	Note            string `json:"note"`
	Star            bool   `json:"star"`
	UserID          int    `json:"user_id"`
	ContentObjectID string `json:"content_object_id"`
	MeetingID       int    `json:"meeting_id"`
}

func LoadPersonalNote(ctx context.Context, ds Getter, id int) (*PersonalNote, error) {
	fields := []string{
		"id",
		"note",
		"star",
		"user_id",
		"content_object_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("personal_note/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c PersonalNote
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Note); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Star); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.UserID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.ContentObjectID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	return &c, nil
}

type Tag struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	TaggedIDs []string `json:"tagged_ids"`
	MeetingID int      `json:"meeting_id"`
}

func LoadTag(ctx context.Context, ds Getter, id int) (*Tag, error) {
	fields := []string{
		"id",
		"name",
		"tagged_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("tag/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Tag
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.TaggedIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	return &c, nil
}

type AgendaItem struct {
	ID              int    `json:"id"`
	ItemNumber      string `json:"item_number"`
	Comment         string `json:"comment"`
	Closed          bool   `json:"closed"`
	Type            string `json:"type"`
	Duration        int    `json:"duration"`
	IsInternal      bool   `json:"is_internal"`
	IsHidden        bool   `json:"is_hidden"`
	Level           int    `json:"level"`
	Weight          int    `json:"weight"`
	ContentObjectID string `json:"content_object_id"`
	ParentID        int    `json:"parent_id"`
	ChildIDs        []int  `json:"child_ids"`
	TagIDs          []int  `json:"tag_ids"`
	ProjectionIDs   []int  `json:"projection_ids"`
	MeetingID       int    `json:"meeting_id"`
}

func LoadAgendaItem(ctx context.Context, ds Getter, id int) (*AgendaItem, error) {
	fields := []string{
		"id",
		"item_number",
		"comment",
		"closed",
		"type",
		"duration",
		"is_internal",
		"is_hidden",
		"level",
		"weight",
		"content_object_id",
		"parent_id",
		"child_ids",
		"tag_ids",
		"projection_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("agenda_item/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c AgendaItem
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.ItemNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Comment); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Closed); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Type); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.Duration); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.IsInternal); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.IsHidden); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.Level); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.ContentObjectID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.ParentID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.ChildIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.TagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	return &c, nil
}

type ListOfSpeakers struct {
	ID              int    `json:"id"`
	Closed          bool   `json:"closed"`
	ContentObjectID string `json:"content_object_id"`
	SpeakerIDs      []int  `json:"speaker_ids"`
	ProjectionIDs   []int  `json:"projection_ids"`
	MeetingID       int    `json:"meeting_id"`
}

func LoadListOfSpeakers(ctx context.Context, ds Getter, id int) (*ListOfSpeakers, error) {
	fields := []string{
		"id",
		"closed",
		"content_object_id",
		"speaker_ids",
		"projection_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("list_of_speakers/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c ListOfSpeakers
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Closed); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.ContentObjectID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.SpeakerIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	return &c, nil
}

type Speaker struct {
	ID               int    `json:"id"`
	BeginTime        int    `json:"begin_time"`
	EndTime          int    `json:"end_time"`
	Weight           int    `json:"weight"`
	SpeechState      string `json:"speech_state"`
	Note             string `json:"note"`
	PointOfOrder     bool   `json:"point_of_order"`
	ListOfSpeakersID int    `json:"list_of_speakers_id"`
	UserID           int    `json:"user_id"`
	MeetingID        int    `json:"meeting_id"`
}

func LoadSpeaker(ctx context.Context, ds Getter, id int) (*Speaker, error) {
	fields := []string{
		"id",
		"begin_time",
		"end_time",
		"weight",
		"speech_state",
		"note",
		"point_of_order",
		"list_of_speakers_id",
		"user_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("speaker/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Speaker
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.BeginTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.EndTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.SpeechState); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.Note); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.PointOfOrder); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.ListOfSpeakersID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.UserID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	return &c, nil
}

type Topic struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Text             string `json:"text"`
	AttachmentIDs    []int  `json:"attachment_ids"`
	AgendaItemID     int    `json:"agenda_item_id"`
	ListOfSpeakersID int    `json:"list_of_speakers_id"`
	OptionIDs        []int  `json:"option_ids"`
	TagIDs           []int  `json:"tag_ids"`
	ProjectionIDs    []int  `json:"projection_ids"`
	MeetingID        int    `json:"meeting_id"`
}

func LoadTopic(ctx context.Context, ds Getter, id int) (*Topic, error) {
	fields := []string{
		"id",
		"title",
		"text",
		"attachment_ids",
		"agenda_item_id",
		"list_of_speakers_id",
		"option_ids",
		"tag_ids",
		"projection_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("topic/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Topic
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Text); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.AttachmentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.AgendaItemID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.ListOfSpeakersID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.OptionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.TagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	return &c, nil
}

type Motion struct {
	ID                                           int               `json:"id"`
	Number                                       string            `json:"number"`
	NumberValue                                  int               `json:"number_value"`
	SequentialNumber                             int               `json:"sequential_number"`
	Title                                        string            `json:"title"`
	Text                                         string            `json:"text"`
	AmendmentParagraph                           map[string]string `json:"amendment_paragraph_$"`
	ModifiedFinalVersion                         string            `json:"modified_final_version"`
	Reason                                       string            `json:"reason"`
	CategoryWeight                               int               `json:"category_weight"`
	StateExtension                               string            `json:"state_extension"`
	RecommendationExtension                      string            `json:"recommendation_extension"`
	SortWeight                                   int               `json:"sort_weight"`
	Created                                      int               `json:"created"`
	LastModified                                 int               `json:"last_modified"`
	LeadMotionID                                 int               `json:"lead_motion_id"`
	AmendmentIDs                                 []int             `json:"amendment_ids"`
	SortParentID                                 int               `json:"sort_parent_id"`
	SortChildIDs                                 []int             `json:"sort_child_ids"`
	OriginID                                     int               `json:"origin_id"`
	DerivedMotionIDs                             []int             `json:"derived_motion_ids"`
	ForwardingTreeMotionIDs                      []int             `json:"forwarding_tree_motion_ids"`
	StateID                                      int               `json:"state_id"`
	RecommendationID                             int               `json:"recommendation_id"`
	RecommendationExtensionReferenceIDs          []string          `json:"recommendation_extension_reference_ids"`
	ReferencedInMotionRecommendationExtensionIDs []int             `json:"referenced_in_motion_recommendation_extension_ids"`
	CategoryID                                   int               `json:"category_id"`
	BlockID                                      int               `json:"block_id"`
	SubmitterIDs                                 []int             `json:"submitter_ids"`
	SupporterIDs                                 []int             `json:"supporter_ids"`
	PollIDs                                      []int             `json:"poll_ids"`
	OptionIDs                                    []int             `json:"option_ids"`
	ChangeRecommendationIDs                      []int             `json:"change_recommendation_ids"`
	StatuteParagraphID                           int               `json:"statute_paragraph_id"`
	CommentIDs                                   []int             `json:"comment_ids"`
	AgendaItemID                                 int               `json:"agenda_item_id"`
	ListOfSpeakersID                             int               `json:"list_of_speakers_id"`
	TagIDs                                       []int             `json:"tag_ids"`
	AttachmentIDs                                []int             `json:"attachment_ids"`
	ProjectionIDs                                []int             `json:"projection_ids"`
	PersonalNoteIDs                              []int             `json:"personal_note_ids"`
	MeetingID                                    int               `json:"meeting_id"`
}

func LoadMotion(ctx context.Context, ds Getter, id int) (*Motion, error) {
	fields := []string{
		"id",
		"number",
		"number_value",
		"sequential_number",
		"title",
		"text",
		"amendment_paragraph_$",
		"modified_final_version",
		"reason",
		"category_weight",
		"state_extension",
		"recommendation_extension",
		"sort_weight",
		"created",
		"last_modified",
		"lead_motion_id",
		"amendment_ids",
		"sort_parent_id",
		"sort_child_ids",
		"origin_id",
		"derived_motion_ids",
		"forwarding_tree_motion_ids",
		"state_id",
		"recommendation_id",
		"recommendation_extension_reference_ids",
		"referenced_in_motion_recommendation_extension_ids",
		"category_id",
		"block_id",
		"submitter_ids",
		"supporter_ids",
		"poll_ids",
		"option_ids",
		"change_recommendation_ids",
		"statute_paragraph_id",
		"comment_ids",
		"agenda_item_id",
		"list_of_speakers_id",
		"tag_ids",
		"attachment_ids",
		"projection_ids",
		"personal_note_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Motion
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Number); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.NumberValue); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.SequentialNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.Text); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		var repl []string
		if err := json.Unmarshal(values[6], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[6], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[6], err)
		}

		data := make(map[string]string, len(repl))
		for i, r := range repl {
			var decoded string
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.AmendmentParagraph = data
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.ModifiedFinalVersion); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.Reason); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.CategoryWeight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.StateExtension); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.RecommendationExtension); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.SortWeight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.Created); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.LastModified); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.LeadMotionID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	if values[16] != nil {
		if err := json.Unmarshal(values[16], &c.AmendmentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[16], err)
		}
	}
	if values[17] != nil {
		if err := json.Unmarshal(values[17], &c.SortParentID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[17], err)
		}
	}
	if values[18] != nil {
		if err := json.Unmarshal(values[18], &c.SortChildIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[18], err)
		}
	}
	if values[19] != nil {
		if err := json.Unmarshal(values[19], &c.OriginID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[19], err)
		}
	}
	if values[20] != nil {
		if err := json.Unmarshal(values[20], &c.DerivedMotionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[20], err)
		}
	}
	if values[21] != nil {
		if err := json.Unmarshal(values[21], &c.ForwardingTreeMotionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[21], err)
		}
	}
	if values[22] != nil {
		if err := json.Unmarshal(values[22], &c.StateID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[22], err)
		}
	}
	if values[23] != nil {
		if err := json.Unmarshal(values[23], &c.RecommendationID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[23], err)
		}
	}
	if values[24] != nil {
		if err := json.Unmarshal(values[24], &c.RecommendationExtensionReferenceIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[24], err)
		}
	}
	if values[25] != nil {
		if err := json.Unmarshal(values[25], &c.ReferencedInMotionRecommendationExtensionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[25], err)
		}
	}
	if values[26] != nil {
		if err := json.Unmarshal(values[26], &c.CategoryID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[26], err)
		}
	}
	if values[27] != nil {
		if err := json.Unmarshal(values[27], &c.BlockID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[27], err)
		}
	}
	if values[28] != nil {
		if err := json.Unmarshal(values[28], &c.SubmitterIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[28], err)
		}
	}
	if values[29] != nil {
		if err := json.Unmarshal(values[29], &c.SupporterIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[29], err)
		}
	}
	if values[30] != nil {
		if err := json.Unmarshal(values[30], &c.PollIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[30], err)
		}
	}
	if values[31] != nil {
		if err := json.Unmarshal(values[31], &c.OptionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[31], err)
		}
	}
	if values[32] != nil {
		if err := json.Unmarshal(values[32], &c.ChangeRecommendationIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[32], err)
		}
	}
	if values[33] != nil {
		if err := json.Unmarshal(values[33], &c.StatuteParagraphID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[33], err)
		}
	}
	if values[34] != nil {
		if err := json.Unmarshal(values[34], &c.CommentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[34], err)
		}
	}
	if values[35] != nil {
		if err := json.Unmarshal(values[35], &c.AgendaItemID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[35], err)
		}
	}
	if values[36] != nil {
		if err := json.Unmarshal(values[36], &c.ListOfSpeakersID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[36], err)
		}
	}
	if values[37] != nil {
		if err := json.Unmarshal(values[37], &c.TagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[37], err)
		}
	}
	if values[38] != nil {
		if err := json.Unmarshal(values[38], &c.AttachmentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[38], err)
		}
	}
	if values[39] != nil {
		if err := json.Unmarshal(values[39], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[39], err)
		}
	}
	if values[40] != nil {
		if err := json.Unmarshal(values[40], &c.PersonalNoteIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[40], err)
		}
	}
	if values[41] != nil {
		if err := json.Unmarshal(values[41], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[41], err)
		}
	}
	return &c, nil
}

type MotionSubmitter struct {
	ID        int `json:"id"`
	Weight    int `json:"weight"`
	UserID    int `json:"user_id"`
	MotionID  int `json:"motion_id"`
	MeetingID int `json:"meeting_id"`
}

func LoadMotionSubmitter(ctx context.Context, ds Getter, id int) (*MotionSubmitter, error) {
	fields := []string{
		"id",
		"weight",
		"user_id",
		"motion_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_submitter/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionSubmitter
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.UserID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.MotionID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	return &c, nil
}

type MotionComment struct {
	ID        int    `json:"id"`
	Comment   string `json:"comment"`
	MotionID  int    `json:"motion_id"`
	SectionID int    `json:"section_id"`
	MeetingID int    `json:"meeting_id"`
}

func LoadMotionComment(ctx context.Context, ds Getter, id int) (*MotionComment, error) {
	fields := []string{
		"id",
		"comment",
		"motion_id",
		"section_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_comment/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionComment
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Comment); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.MotionID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.SectionID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	return &c, nil
}

type MotionCommentSection struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Weight        int    `json:"weight"`
	CommentIDs    []int  `json:"comment_ids"`
	ReadGroupIDs  []int  `json:"read_group_ids"`
	WriteGroupIDs []int  `json:"write_group_ids"`
	MeetingID     int    `json:"meeting_id"`
}

func LoadMotionCommentSection(ctx context.Context, ds Getter, id int) (*MotionCommentSection, error) {
	fields := []string{
		"id",
		"name",
		"weight",
		"comment_ids",
		"read_group_ids",
		"write_group_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_comment_section/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionCommentSection
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.CommentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.ReadGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.WriteGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	return &c, nil
}

type MotionCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Prefix    string `json:"prefix"`
	Weight    int    `json:"weight"`
	Level     int    `json:"level"`
	ParentID  int    `json:"parent_id"`
	ChildIDs  []int  `json:"child_ids"`
	MotionIDs []int  `json:"motion_ids"`
	MeetingID int    `json:"meeting_id"`
}

func LoadMotionCategory(ctx context.Context, ds Getter, id int) (*MotionCategory, error) {
	fields := []string{
		"id",
		"name",
		"prefix",
		"weight",
		"level",
		"parent_id",
		"child_ids",
		"motion_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_category/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionCategory
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Prefix); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Level); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.ParentID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.ChildIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.MotionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	return &c, nil
}

type MotionBlock struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Internal         bool   `json:"internal"`
	MotionIDs        []int  `json:"motion_ids"`
	AgendaItemID     int    `json:"agenda_item_id"`
	ListOfSpeakersID int    `json:"list_of_speakers_id"`
	ProjectionIDs    []int  `json:"projection_ids"`
	MeetingID        int    `json:"meeting_id"`
}

func LoadMotionBlock(ctx context.Context, ds Getter, id int) (*MotionBlock, error) {
	fields := []string{
		"id",
		"title",
		"internal",
		"motion_ids",
		"agenda_item_id",
		"list_of_speakers_id",
		"projection_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_block/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionBlock
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Internal); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.MotionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.AgendaItemID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.ListOfSpeakersID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	return &c, nil
}

type MotionChangeRecommendation struct {
	ID               int    `json:"id"`
	Rejected         bool   `json:"rejected"`
	Internal         bool   `json:"internal"`
	Type             string `json:"type"`
	OtherDescription string `json:"other_description"`
	LineFrom         int    `json:"line_from"`
	LineTo           int    `json:"line_to"`
	Text             string `json:"text"`
	CreationTime     int    `json:"creation_time"`
	MotionID         int    `json:"motion_id"`
	MeetingID        int    `json:"meeting_id"`
}

func LoadMotionChangeRecommendation(ctx context.Context, ds Getter, id int) (*MotionChangeRecommendation, error) {
	fields := []string{
		"id",
		"rejected",
		"internal",
		"type",
		"other_description",
		"line_from",
		"line_to",
		"text",
		"creation_time",
		"motion_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_change_recommendation/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionChangeRecommendation
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Rejected); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Internal); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Type); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.OtherDescription); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.LineFrom); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.LineTo); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.Text); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.CreationTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.MotionID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	return &c, nil
}

type MotionState struct {
	ID                               int      `json:"id"`
	Name                             string   `json:"name"`
	RecommendationLabel              string   `json:"recommendation_label"`
	CssClass                         string   `json:"css_class"`
	Restrictions                     []string `json:"restrictions"`
	AllowSupport                     bool     `json:"allow_support"`
	AllowCreatePoll                  bool     `json:"allow_create_poll"`
	AllowSubmitterEdit               bool     `json:"allow_submitter_edit"`
	SetNumber                        bool     `json:"set_number"`
	ShowStateExtensionField          bool     `json:"show_state_extension_field"`
	MergeAmendmentIntoFinal          string   `json:"merge_amendment_into_final"`
	ShowRecommendationExtensionField bool     `json:"show_recommendation_extension_field"`
	NextStateIDs                     []int    `json:"next_state_ids"`
	PreviousStateIDs                 []int    `json:"previous_state_ids"`
	MotionIDs                        []int    `json:"motion_ids"`
	MotionRecommendationIDs          []int    `json:"motion_recommendation_ids"`
	WorkflowID                       int      `json:"workflow_id"`
	FirstStateOfWorkflowID           int      `json:"first_state_of_workflow_id"`
	MeetingID                        int      `json:"meeting_id"`
}

func LoadMotionState(ctx context.Context, ds Getter, id int) (*MotionState, error) {
	fields := []string{
		"id",
		"name",
		"recommendation_label",
		"css_class",
		"restrictions",
		"allow_support",
		"allow_create_poll",
		"allow_submitter_edit",
		"set_number",
		"show_state_extension_field",
		"merge_amendment_into_final",
		"show_recommendation_extension_field",
		"next_state_ids",
		"previous_state_ids",
		"motion_ids",
		"motion_recommendation_ids",
		"workflow_id",
		"first_state_of_workflow_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_state/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionState
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.RecommendationLabel); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.CssClass); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Restrictions); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.AllowSupport); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.AllowCreatePoll); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.AllowSubmitterEdit); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.SetNumber); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.ShowStateExtensionField); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.MergeAmendmentIntoFinal); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.ShowRecommendationExtensionField); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.NextStateIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.PreviousStateIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.MotionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.MotionRecommendationIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	if values[16] != nil {
		if err := json.Unmarshal(values[16], &c.WorkflowID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[16], err)
		}
	}
	if values[17] != nil {
		if err := json.Unmarshal(values[17], &c.FirstStateOfWorkflowID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[17], err)
		}
	}
	if values[18] != nil {
		if err := json.Unmarshal(values[18], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[18], err)
		}
	}
	return &c, nil
}

type MotionWorkflow struct {
	ID                                       int    `json:"id"`
	Name                                     string `json:"name"`
	StateIDs                                 []int  `json:"state_ids"`
	FirstStateID                             int    `json:"first_state_id"`
	DefaultWorkflowMeetingID                 int    `json:"default_workflow_meeting_id"`
	DefaultAmendmentWorkflowMeetingID        int    `json:"default_amendment_workflow_meeting_id"`
	DefaultStatuteAmendmentWorkflowMeetingID int    `json:"default_statute_amendment_workflow_meeting_id"`
	MeetingID                                int    `json:"meeting_id"`
}

func LoadMotionWorkflow(ctx context.Context, ds Getter, id int) (*MotionWorkflow, error) {
	fields := []string{
		"id",
		"name",
		"state_ids",
		"first_state_id",
		"default_workflow_meeting_id",
		"default_amendment_workflow_meeting_id",
		"default_statute_amendment_workflow_meeting_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_workflow/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionWorkflow
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.StateIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.FirstStateID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.DefaultWorkflowMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.DefaultAmendmentWorkflowMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.DefaultStatuteAmendmentWorkflowMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	return &c, nil
}

type MotionStatuteParagraph struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Weight    int    `json:"weight"`
	MotionIDs []int  `json:"motion_ids"`
	MeetingID int    `json:"meeting_id"`
}

func LoadMotionStatuteParagraph(ctx context.Context, ds Getter, id int) (*MotionStatuteParagraph, error) {
	fields := []string{
		"id",
		"title",
		"text",
		"weight",
		"motion_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("motion_statute_paragraph/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c MotionStatuteParagraph
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Text); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.MotionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	return &c, nil
}

type Poll struct {
	ID                    int             `json:"id"`
	Description           string          `json:"description"`
	Title                 string          `json:"title"`
	Type                  string          `json:"type"`
	Backend               string          `json:"backend"`
	IsPseudoanonymized    bool            `json:"is_pseudoanonymized"`
	Pollmethod            string          `json:"pollmethod"`
	State                 string          `json:"state"`
	MinVotesAmount        int             `json:"min_votes_amount"`
	MaxVotesAmount        int             `json:"max_votes_amount"`
	GlobalYes             bool            `json:"global_yes"`
	GlobalNo              bool            `json:"global_no"`
	GlobalAbstain         bool            `json:"global_abstain"`
	OnehundredPercentBase string          `json:"onehundred_percent_base"`
	MajorityMethod        string          `json:"majority_method"`
	Votesvalid            int             `json:"votesvalid"`
	Votesinvalid          int             `json:"votesinvalid"`
	Votescast             int             `json:"votescast"`
	EntitledUsersAtStop   json.RawMessage `json:"entitled_users_at_stop"`
	ContentObjectID       string          `json:"content_object_id"`
	OptionIDs             []int           `json:"option_ids"`
	GlobalOptionID        int             `json:"global_option_id"`
	VotedIDs              []int           `json:"voted_ids"`
	EntitledGroupIDs      []int           `json:"entitled_group_ids"`
	ProjectionIDs         []int           `json:"projection_ids"`
	MeetingID             int             `json:"meeting_id"`
}

func LoadPoll(ctx context.Context, ds Getter, id int) (*Poll, error) {
	fields := []string{
		"id",
		"description",
		"title",
		"type",
		"backend",
		"is_pseudoanonymized",
		"pollmethod",
		"state",
		"min_votes_amount",
		"max_votes_amount",
		"global_yes",
		"global_no",
		"global_abstain",
		"onehundred_percent_base",
		"majority_method",
		"votesvalid",
		"votesinvalid",
		"votescast",
		"entitled_users_at_stop",
		"content_object_id",
		"option_ids",
		"global_option_id",
		"voted_ids",
		"entitled_group_ids",
		"projection_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("poll/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Poll
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Description); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Type); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Backend); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.IsPseudoanonymized); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.Pollmethod); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.State); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.MinVotesAmount); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.MaxVotesAmount); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.GlobalYes); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.GlobalNo); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.GlobalAbstain); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.OnehundredPercentBase); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.MajorityMethod); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.Votesvalid); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	if values[16] != nil {
		if err := json.Unmarshal(values[16], &c.Votesinvalid); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[16], err)
		}
	}
	if values[17] != nil {
		if err := json.Unmarshal(values[17], &c.Votescast); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[17], err)
		}
	}
	if values[18] != nil {
		if err := json.Unmarshal(values[18], &c.EntitledUsersAtStop); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[18], err)
		}
	}
	if values[19] != nil {
		if err := json.Unmarshal(values[19], &c.ContentObjectID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[19], err)
		}
	}
	if values[20] != nil {
		if err := json.Unmarshal(values[20], &c.OptionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[20], err)
		}
	}
	if values[21] != nil {
		if err := json.Unmarshal(values[21], &c.GlobalOptionID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[21], err)
		}
	}
	if values[22] != nil {
		if err := json.Unmarshal(values[22], &c.VotedIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[22], err)
		}
	}
	if values[23] != nil {
		if err := json.Unmarshal(values[23], &c.EntitledGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[23], err)
		}
	}
	if values[24] != nil {
		if err := json.Unmarshal(values[24], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[24], err)
		}
	}
	if values[25] != nil {
		if err := json.Unmarshal(values[25], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[25], err)
		}
	}
	return &c, nil
}

type Option struct {
	ID                         int    `json:"id"`
	Weight                     int    `json:"weight"`
	Text                       string `json:"text"`
	Yes                        int    `json:"yes"`
	No                         int    `json:"no"`
	Abstain                    int    `json:"abstain"`
	PollID                     int    `json:"poll_id"`
	UsedAsGlobalOptionInPollID int    `json:"used_as_global_option_in_poll_id"`
	VoteIDs                    []int  `json:"vote_ids"`
	ContentObjectID            string `json:"content_object_id"`
	MeetingID                  int    `json:"meeting_id"`
}

func LoadOption(ctx context.Context, ds Getter, id int) (*Option, error) {
	fields := []string{
		"id",
		"weight",
		"text",
		"yes",
		"no",
		"abstain",
		"poll_id",
		"used_as_global_option_in_poll_id",
		"vote_ids",
		"content_object_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("option/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Option
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Text); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Yes); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.No); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.Abstain); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.PollID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.UsedAsGlobalOptionInPollID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.VoteIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.ContentObjectID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	return &c, nil
}

type Vote struct {
	ID              int    `json:"id"`
	Weight          int    `json:"weight"`
	Value           string `json:"value"`
	UserToken       string `json:"user_token"`
	OptionID        int    `json:"option_id"`
	UserID          int    `json:"user_id"`
	DelegatedUserID int    `json:"delegated_user_id"`
	MeetingID       int    `json:"meeting_id"`
}

func LoadVote(ctx context.Context, ds Getter, id int) (*Vote, error) {
	fields := []string{
		"id",
		"weight",
		"value",
		"user_token",
		"option_id",
		"user_id",
		"delegated_user_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("vote/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Vote
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Value); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.UserToken); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.OptionID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.UserID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.DelegatedUserID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	return &c, nil
}

type Assignment struct {
	ID                     int    `json:"id"`
	Title                  string `json:"title"`
	Description            string `json:"description"`
	OpenPosts              int    `json:"open_posts"`
	Phase                  string `json:"phase"`
	DefaultPollDescription string `json:"default_poll_description"`
	NumberPollCandidates   bool   `json:"number_poll_candidates"`
	CandidateIDs           []int  `json:"candidate_ids"`
	PollIDs                []int  `json:"poll_ids"`
	AgendaItemID           int    `json:"agenda_item_id"`
	ListOfSpeakersID       int    `json:"list_of_speakers_id"`
	TagIDs                 []int  `json:"tag_ids"`
	AttachmentIDs          []int  `json:"attachment_ids"`
	ProjectionIDs          []int  `json:"projection_ids"`
	MeetingID              int    `json:"meeting_id"`
}

func LoadAssignment(ctx context.Context, ds Getter, id int) (*Assignment, error) {
	fields := []string{
		"id",
		"title",
		"description",
		"open_posts",
		"phase",
		"default_poll_description",
		"number_poll_candidates",
		"candidate_ids",
		"poll_ids",
		"agenda_item_id",
		"list_of_speakers_id",
		"tag_ids",
		"attachment_ids",
		"projection_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("assignment/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Assignment
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Description); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.OpenPosts); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Phase); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.DefaultPollDescription); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.NumberPollCandidates); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.CandidateIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.PollIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.AgendaItemID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.ListOfSpeakersID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.TagIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.AttachmentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	return &c, nil
}

type AssignmentCandidate struct {
	ID           int `json:"id"`
	Weight       int `json:"weight"`
	AssignmentID int `json:"assignment_id"`
	UserID       int `json:"user_id"`
	MeetingID    int `json:"meeting_id"`
}

func LoadAssignmentCandidate(ctx context.Context, ds Getter, id int) (*AssignmentCandidate, error) {
	fields := []string{
		"id",
		"weight",
		"assignment_id",
		"user_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("assignment_candidate/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c AssignmentCandidate
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.AssignmentID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.UserID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	return &c, nil
}

type Mediafile struct {
	ID                      int             `json:"id"`
	Title                   string          `json:"title"`
	IsDirectory             bool            `json:"is_directory"`
	Filesize                int             `json:"filesize"`
	Filename                string          `json:"filename"`
	Mimetype                string          `json:"mimetype"`
	PdfInformation          json.RawMessage `json:"pdf_information"`
	CreateTimestamp         int             `json:"create_timestamp"`
	IsPublic                bool            `json:"is_public"`
	InheritedAccessGroupIDs []int           `json:"inherited_access_group_ids"`
	AccessGroupIDs          []int           `json:"access_group_ids"`
	ParentID                int             `json:"parent_id"`
	ChildIDs                []int           `json:"child_ids"`
	ListOfSpeakersID        int             `json:"list_of_speakers_id"`
	ProjectionIDs           []int           `json:"projection_ids"`
	AttachmentIDs           []string        `json:"attachment_ids"`
	MeetingID               int             `json:"meeting_id"`
	UsedAsLogoInMeetingID   map[string]int  `json:"used_as_logo_$_in_meeting_id"`
	UsedAsFontInMeetingID   map[string]int  `json:"used_as_font_$_in_meeting_id"`
}

func LoadMediafile(ctx context.Context, ds Getter, id int) (*Mediafile, error) {
	fields := []string{
		"id",
		"title",
		"is_directory",
		"filesize",
		"filename",
		"mimetype",
		"pdf_information",
		"create_timestamp",
		"is_public",
		"inherited_access_group_ids",
		"access_group_ids",
		"parent_id",
		"child_ids",
		"list_of_speakers_id",
		"projection_ids",
		"attachment_ids",
		"meeting_id",
		"used_as_logo_$_in_meeting_id",
		"used_as_font_$_in_meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("mediafile/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Mediafile
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.IsDirectory); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Filesize); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Filename); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.Mimetype); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.PdfInformation); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.CreateTimestamp); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.IsPublic); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.InheritedAccessGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.AccessGroupIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.ParentID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.ChildIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.ListOfSpeakersID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.AttachmentIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	if values[16] != nil {
		if err := json.Unmarshal(values[16], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[16], err)
		}
	}
	if values[17] != nil {
		var repl []string
		if err := json.Unmarshal(values[17], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[17], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[17], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[17], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.UsedAsLogoInMeetingID = data
	}
	if values[18] != nil {
		var repl []string
		if err := json.Unmarshal(values[18], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[18], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[18], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[18], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.UsedAsFontInMeetingID = data
	}
	return &c, nil
}

type Projector struct {
	ID                                int            `json:"id"`
	Name                              string         `json:"name"`
	Scale                             int            `json:"scale"`
	Scroll                            int            `json:"scroll"`
	Width                             int            `json:"width"`
	AspectRatioNumerator              int            `json:"aspect_ratio_numerator"`
	AspectRatioDenominator            int            `json:"aspect_ratio_denominator"`
	Color                             string         `json:"color"`
	BackgroundColor                   string         `json:"background_color"`
	HeaderBackgroundColor             string         `json:"header_background_color"`
	HeaderFontColor                   string         `json:"header_font_color"`
	HeaderH1Color                     string         `json:"header_h1_color"`
	ChyronBackgroundColor             string         `json:"chyron_background_color"`
	ChyronFontColor                   string         `json:"chyron_font_color"`
	ShowHeaderFooter                  bool           `json:"show_header_footer"`
	ShowTitle                         bool           `json:"show_title"`
	ShowLogo                          bool           `json:"show_logo"`
	ShowClock                         bool           `json:"show_clock"`
	CurrentProjectionIDs              []int          `json:"current_projection_ids"`
	PreviewProjectionIDs              []int          `json:"preview_projection_ids"`
	HistoryProjectionIDs              []int          `json:"history_projection_ids"`
	UsedAsReferenceProjectorMeetingID int            `json:"used_as_reference_projector_meeting_id"`
	UsedAsDefaultInMeetingID          map[string]int `json:"used_as_default_$_in_meeting_id"`
	MeetingID                         int            `json:"meeting_id"`
}

func LoadProjector(ctx context.Context, ds Getter, id int) (*Projector, error) {
	fields := []string{
		"id",
		"name",
		"scale",
		"scroll",
		"width",
		"aspect_ratio_numerator",
		"aspect_ratio_denominator",
		"color",
		"background_color",
		"header_background_color",
		"header_font_color",
		"header_h1_color",
		"chyron_background_color",
		"chyron_font_color",
		"show_header_footer",
		"show_title",
		"show_logo",
		"show_clock",
		"current_projection_ids",
		"preview_projection_ids",
		"history_projection_ids",
		"used_as_reference_projector_meeting_id",
		"used_as_default_$_in_meeting_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("projector/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Projector
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Name); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Scale); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Scroll); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Width); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.AspectRatioNumerator); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.AspectRatioDenominator); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.Color); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.BackgroundColor); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.HeaderBackgroundColor); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	if values[10] != nil {
		if err := json.Unmarshal(values[10], &c.HeaderFontColor); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[10], err)
		}
	}
	if values[11] != nil {
		if err := json.Unmarshal(values[11], &c.HeaderH1Color); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[11], err)
		}
	}
	if values[12] != nil {
		if err := json.Unmarshal(values[12], &c.ChyronBackgroundColor); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[12], err)
		}
	}
	if values[13] != nil {
		if err := json.Unmarshal(values[13], &c.ChyronFontColor); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[13], err)
		}
	}
	if values[14] != nil {
		if err := json.Unmarshal(values[14], &c.ShowHeaderFooter); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[14], err)
		}
	}
	if values[15] != nil {
		if err := json.Unmarshal(values[15], &c.ShowTitle); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[15], err)
		}
	}
	if values[16] != nil {
		if err := json.Unmarshal(values[16], &c.ShowLogo); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[16], err)
		}
	}
	if values[17] != nil {
		if err := json.Unmarshal(values[17], &c.ShowClock); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[17], err)
		}
	}
	if values[18] != nil {
		if err := json.Unmarshal(values[18], &c.CurrentProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[18], err)
		}
	}
	if values[19] != nil {
		if err := json.Unmarshal(values[19], &c.PreviewProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[19], err)
		}
	}
	if values[20] != nil {
		if err := json.Unmarshal(values[20], &c.HistoryProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[20], err)
		}
	}
	if values[21] != nil {
		if err := json.Unmarshal(values[21], &c.UsedAsReferenceProjectorMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[21], err)
		}
	}
	if values[22] != nil {
		var repl []string
		if err := json.Unmarshal(values[22], &repl); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[22], err)
		}

		tkeys := make([]string, len(repl))
		for i, r := range repl {
			tkeys[i] = strings.Replace(keys[22], "$", "$"+r, 1)
		}

		values, err := ds.Get(ctx, tkeys...)
		if err != nil {
			return nil, fmt.Errorf("getting template keys of field %s: %w", fields[22], err)
		}

		data := make(map[string]int, len(repl))
		for i, r := range repl {
			var decoded int
			if err := json.Unmarshal(values[i], &decoded); err != nil {
				return nil, fmt.Errorf("decoding field %s: %w", tkeys[i], err)
			}
			data[r] = decoded
		}
		c.UsedAsDefaultInMeetingID = data
	}
	if values[23] != nil {
		if err := json.Unmarshal(values[23], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[23], err)
		}
	}
	return &c, nil
}

type Projection struct {
	ID                 int             `json:"id"`
	Options            json.RawMessage `json:"options"`
	Stable             bool            `json:"stable"`
	Weight             int             `json:"weight"`
	Type               string          `json:"type"`
	CurrentProjectorID int             `json:"current_projector_id"`
	PreviewProjectorID int             `json:"preview_projector_id"`
	HistoryProjectorID int             `json:"history_projector_id"`
	ContentObjectID    string          `json:"content_object_id"`
	MeetingID          int             `json:"meeting_id"`
}

func LoadProjection(ctx context.Context, ds Getter, id int) (*Projection, error) {
	fields := []string{
		"id",
		"options",
		"stable",
		"weight",
		"type",
		"current_projector_id",
		"preview_projector_id",
		"history_projector_id",
		"content_object_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("projection/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c Projection
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Options); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Stable); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.Weight); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.Type); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.CurrentProjectorID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.PreviewProjectorID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.HistoryProjectorID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.ContentObjectID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	return &c, nil
}

type ProjectorMessage struct {
	ID            int    `json:"id"`
	Message       string `json:"message"`
	ProjectionIDs []int  `json:"projection_ids"`
	MeetingID     int    `json:"meeting_id"`
}

func LoadProjectorMessage(ctx context.Context, ds Getter, id int) (*ProjectorMessage, error) {
	fields := []string{
		"id",
		"message",
		"projection_ids",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("projector_message/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c ProjectorMessage
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Message); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	return &c, nil
}

type ProjectorCountdown struct {
	ID                                    int     `json:"id"`
	Title                                 string  `json:"title"`
	Description                           string  `json:"description"`
	DefaultTime                           int     `json:"default_time"`
	CountdownTime                         float32 `json:"countdown_time"`
	Running                               bool    `json:"running"`
	ProjectionIDs                         []int   `json:"projection_ids"`
	UsedAsListOfSpeakerCountdownMeetingID int     `json:"used_as_list_of_speaker_countdown_meeting_id"`
	UsedAsPollCountdownMeetingID          int     `json:"used_as_poll_countdown_meeting_id"`
	MeetingID                             int     `json:"meeting_id"`
}

func LoadProjectorCountdown(ctx context.Context, ds Getter, id int) (*ProjectorCountdown, error) {
	fields := []string{
		"id",
		"title",
		"description",
		"default_time",
		"countdown_time",
		"running",
		"projection_ids",
		"used_as_list_of_speaker_countdown_meeting_id",
		"used_as_poll_countdown_meeting_id",
		"meeting_id",
	}

	keys := make([]string, len(fields))
	for i, f := range fields {
		keys[i] = fmt.Sprintf("projector_countdown/%d/%s", id, f)
	}

	values, err := ds.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("loading fields: %w", err)
	}

	var c ProjectorCountdown
	if values[0] != nil {
		if err := json.Unmarshal(values[0], &c.ID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[0], err)
		}
	}
	if values[1] != nil {
		if err := json.Unmarshal(values[1], &c.Title); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[1], err)
		}
	}
	if values[2] != nil {
		if err := json.Unmarshal(values[2], &c.Description); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[2], err)
		}
	}
	if values[3] != nil {
		if err := json.Unmarshal(values[3], &c.DefaultTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[3], err)
		}
	}
	if values[4] != nil {
		if err := json.Unmarshal(values[4], &c.CountdownTime); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[4], err)
		}
	}
	if values[5] != nil {
		if err := json.Unmarshal(values[5], &c.Running); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[5], err)
		}
	}
	if values[6] != nil {
		if err := json.Unmarshal(values[6], &c.ProjectionIDs); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[6], err)
		}
	}
	if values[7] != nil {
		if err := json.Unmarshal(values[7], &c.UsedAsListOfSpeakerCountdownMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[7], err)
		}
	}
	if values[8] != nil {
		if err := json.Unmarshal(values[8], &c.UsedAsPollCountdownMeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[8], err)
		}
	}
	if values[9] != nil {
		if err := json.Unmarshal(values[9], &c.MeetingID); err != nil {
			return nil, fmt.Errorf("decoding %s: %w", keys[9], err)
		}
	}
	return &c, nil
}
