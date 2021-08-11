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

func (u User) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	if mperms.UserID() == UserID {
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
			if UserID == uid {
				return true, nil
			}
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking committee management level: %w", err)
	}

	meetingIDs := fetch.Field().User_GroupIDsTmpl(ctx, UserID)
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
		if delegated == UserID {
			return true, nil
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking vote deleted to: %w", err)
	}

	for _, meetingID := range fetch.Field().User_VoteDelegationsFromIDsTmpl(ctx, mperms.UserID()) {
		delegations := fetch.Field().User_VoteDelegationsFromIDs(ctx, mperms.UserID(), meetingID)
		for _, uid := range delegations {
			if uid == UserID {
				return true, nil
			}
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking vote delegations form: %w", err)
	}

	// TODO: required user to see:
	// There is a related object:

	//     There exists a motion which Y can see and X is a submitter/supporter.
	//     There exists an option which Y can see and X is the linked content object.
	//     There exists an assignment candidate which Y can see and X is the linked user.
	//     There exists a speaker which Y can see and X is the linked user.
	//     There exists a poll where Y can see the poll/voted_ids and X is part of that list.
	//     There exists a vote which Y can see and X is linked in user_id or delegated_user_id.

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
