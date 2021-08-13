package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// User handels the restrictions for the user collection.
type User struct{}

// Modes returns the field restriction for each mode.
func (u User) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return u.see
	case "B":
		return u.modeB
	case "C":
		return u.modeC
	case "D":
		return u.modeD
	case "E":
		return u.modeE
	case "F":
		return u.modeF
	case "G":
		return u.modeG
	}
	return nil
}

// SuperAdmin restricts the super admin.
func (u User) SuperAdmin(mode string) FieldRestricter {
	if mode == "G" {
		return u.modeG
	}
	return Allways
}

func (u User) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, userID int) (bool, error) {
	if mperms.UserID() == userID {
		return true, nil
	}

	canManageUsers, err := perm.HasOrganizationManagementLevel(ctx, fetch, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("get organization level: %w", err)
	}

	if canManageUsers {
		return true, nil
	}

	committeeManager := make(map[int]bool)
	for _, committeeID := range fetch.Field().User_CommitteeManagementLevelTmpl(ctx, mperms.UserID()) {
		committeeManagementLevel := fetch.Field().User_CommitteeManagementLevel(ctx, mperms.UserID(), committeeID)
		if committeeManagementLevel != "can_manage" {
			continue
		}
		committeeManager[committeeID] = true

		userIDs := fetch.Field().Committee_UserIDs(ctx, committeeID)
		for _, uid := range userIDs {
			if userID == uid {
				return true, nil
			}
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking committee management level: %w", err)
	}

	meetingIDs := fetch.Field().User_GroupIDsTmpl(ctx, userID)
	for _, meetingID := range meetingIDs {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.UserCanSee) {
			return true, nil
		}

		cid := fetch.Field().Meeting_CommitteeID(ctx, meetingID)
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("getting committee id of meeting %d: %w", meetingID, err)
		}

		if committeeManager[cid] {
			return true, nil
		}
	}

	for _, meetingID := range fetch.Field().User_VoteDelegatedToIDTmpl(ctx, mperms.UserID()) {
		delegated := fetch.Field().User_VoteDelegatedToID(ctx, mperms.UserID(), meetingID)
		if delegated == userID {
			return true, nil
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking vote deleted to: %w", err)
	}

	for _, meetingID := range fetch.Field().User_VoteDelegationsFromIDsTmpl(ctx, mperms.UserID()) {
		delegations := fetch.Field().User_VoteDelegationsFromIDs(ctx, mperms.UserID(), meetingID)
		for _, uid := range delegations {
			if uid == userID {
				return true, nil
			}
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking vote delegations form: %w", err)
	}

	requiredObjects := []struct {
		name     string
		tmplFunc func(context.Context, int) []int
		elemFunc func(context.Context, int, int) []int
		seeFunc  FieldRestricter
	}{
		{
			"motion submitter",
			fetch.Field().User_SubmittedMotionIDsTmpl,
			fetch.Field().User_SubmittedMotionIDs,
			Motion{}.see,
		},

		{
			"motion supporter",
			fetch.Field().User_SupportedMotionIDsTmpl,
			fetch.Field().User_SupportedMotionIDs,
			Motion{}.see,
		},

		{
			"option",
			fetch.Field().User_OptionIDsTmpl,
			fetch.Field().User_OptionIDs,
			Option{}.see,
		},

		{
			"assignment candidate",
			fetch.Field().User_AssignmentCandidateIDsTmpl,
			fetch.Field().User_AssignmentCandidateIDs,
			AssignmentCandidate{}.see,
		},

		{
			"speaker",
			fetch.Field().User_SpeakerIDsTmpl,
			fetch.Field().User_SpeakerIDs,
			Speaker{}.see,
		},

		{
			"poll voted",
			fetch.Field().User_PollVotedIDsTmpl,
			fetch.Field().User_PollVotedIDs,
			Poll{}.modeB, // Checking field poll/voted_ids that is in modeB and not in see.
		},

		{
			"vote user",
			fetch.Field().User_VoteIDsTmpl,
			fetch.Field().User_VoteIDs,
			Vote{}.see,
		},

		{
			"vote delegated user",
			fetch.Field().User_VoteDelegatedVoteIDsTmpl,
			fetch.Field().User_VoteDelegatedVoteIDs,
			Vote{}.see,
		},
	}

	for _, r := range requiredObjects {
		for _, meetingID := range r.tmplFunc(ctx, userID) {
			for _, elementID := range r.elemFunc(ctx, userID, meetingID) {
				see, err := r.seeFunc(ctx, fetch, mperms, elementID)
				if err != nil {
					return false, fmt.Errorf("checking required object %q: %w", r.name, err)
				}

				if see {
					return true, nil
				}
			}
		}
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("getting object %q: %w", r.name, err)
		}
	}

	return false, nil
}

func (u User) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	return mperms.UserID() == UserID, nil
}

func (u User) modeC(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	if mperms.UserID() == UserID {
		return true, nil
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, fetch, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return true, nil
	}

	meetingIDs := fetch.Field().User_GroupIDsTmpl(ctx, UserID)
	for _, meetingID := range meetingIDs {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.UserCanSeeExtraData) {
			return true, nil
		}
	}

	return false, nil
}

func (u User) modeD(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	canManage, err := perm.HasOrganizationManagementLevel(ctx, fetch, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return true, nil
	}

	meetingIDs := fetch.Field().User_GroupIDsTmpl(ctx, UserID)
	for _, meetingID := range meetingIDs {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.UserCanManage) {
			return true, nil
		}
	}

	return false, nil
}

func (u User) modeE(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	if mperms.UserID() == UserID {
		return true, nil
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, fetch, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return true, nil
	}

	for _, committeeID := range fetch.Field().User_CommitteeManagementLevelTmpl(ctx, mperms.UserID()) {
		committeeManagementLevel := fetch.Field().User_CommitteeManagementLevel(ctx, mperms.UserID(), committeeID)
		if committeeManagementLevel != "can_manage" {
			continue
		}

		userIDs := fetch.Field().Committee_UserIDs(ctx, committeeID)
		for _, uid := range userIDs {
			if UserID == uid {
				return true, nil
			}
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking committee management level: %w", err)
	}

	return false, nil
}

func (u User) modeF(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	if mperms.UserID() == UserID {
		return true, nil
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, fetch, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return true, nil
	}

	return false, nil
}

func (u User) modeG(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	return false, nil
}
