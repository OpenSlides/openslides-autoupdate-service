package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// User handels the restrictions for the user collection.
//
// Y is the request user and X the user, that is requested.
//
// The user Y can see a user X, if at least one condition is true:
//     Y==X
//     Y has the OML can_manage_users or higher.
//     There exists a committee where Y has the CML can_manage and X is in committee/user_ids.
//     X is in a group of a meeting where Y has user.can_see.
//     There exists a meeting where Y has the CML can_manage for the meeting's committee X is in meeting/user_ids.
//     There is a related object:
//         There exists a motion which Y can see and X is a submitter/supporter.
//         There exists an option which Y can see and X is the linked content object.
//         There exists an assignment candidate which Y can see and X is the linked user.
//         There exists a speaker which Y can see and X is the linked user.
//         There exists a poll where Y can see the poll/voted_ids and X is part of that list.
//         There exists a vote which Y can see and X is linked in user_id or delegated_user_id.
//         There exists a chat_message which Y can see and X has sent it (specified by chat_message/user_id).
//     X is linked in one of the relations vote_delegated_$_to_id or vote_delegations_$_from_ids of Y.
//
// Mode A: Y can see X.
//
// Mode B: Y==X.
//
// Mode D: Y can see these fields if at least one condition is true:
//     Y has the OML can_manage_users or higher.
//     X is in a group of a meeting where Y has user.can_manage.
//
// Mode E: Y can see these fields if at least one condition is true:
//     Y has the OML can_manage_users or higher.
//     There exists a committee where Y has the CML can_manage and X is in committee/user_ids.
//     X is in a group of a meeting where Y has user.can_manage.
//     Y==X.
//
// Mode F: Y has the OML can_manage_users or higher or Y==X.
//
// Mode G: No one. Not even the superadmin.
type User struct{}

// Modes returns the field restriction for each mode.
func (u User) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return u.see
	case "B":
		return u.modeB
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

func (u User) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, userID int) (bool, error) {
	if mperms.UserID() == userID {
		return true, nil
	}

	committeeManager := make(map[int]bool)
	if mperms.UserID() != 0 {
		canManageUsers, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
		if err != nil {
			return false, fmt.Errorf("get organization level: %w", err)
		}

		if canManageUsers {
			return true, nil
		}

		for _, committeeID := range ds.User_CommitteeManagementLevelTmpl(mperms.UserID()).ErrorLater(ctx) {
			committeeManagementLevel := ds.User_CommitteeManagementLevel(mperms.UserID(), committeeID).ErrorLater(ctx)
			if committeeManagementLevel != "can_manage" {
				continue
			}
			committeeManager[committeeID] = true

			userIDs := ds.Committee_UserIDs(committeeID).ErrorLater(ctx)
			for _, uid := range userIDs {
				if userID == uid {
					return true, nil
				}
			}
		}
		if err := ds.Err(); err != nil {
			return false, fmt.Errorf("checking committee management level: %w", err)
		}
	}

	meetingIDs := ds.User_GroupIDsTmpl(userID).ErrorLater(ctx)
	for _, meetingID := range meetingIDs {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.UserCanSee) {
			return true, nil
		}

		cid, err := ds.Meeting_CommitteeID(meetingID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting committee id of meeting %d: %w", meetingID, err)
		}

		if committeeManager[cid] {
			return true, nil
		}
	}

	if mperms.UserID() != 0 {
		for _, meetingID := range ds.User_VoteDelegatedToIDTmpl(mperms.UserID()).ErrorLater(ctx) {
			delegated := ds.User_VoteDelegatedToID(mperms.UserID(), meetingID).ErrorLater(ctx)
			if delegated == userID {
				return true, nil
			}
		}
		if err := ds.Err(); err != nil {
			return false, fmt.Errorf("checking vote deleted to: %w", err)
		}

		for _, meetingID := range ds.User_VoteDelegationsFromIDsTmpl(mperms.UserID()).ErrorLater(ctx) {
			delegations := ds.User_VoteDelegationsFromIDs(mperms.UserID(), meetingID).ErrorLater(ctx)
			for _, uid := range delegations {
				if uid == userID {
					return true, nil
				}
			}
		}
		if err := ds.Err(); err != nil {
			return false, fmt.Errorf("checking vote delegations form: %w", err)
		}
	}

	requiredObjects := []struct {
		name     string
		tmplFunc func(int) *datastore.ValueIDSlice
		elemFunc func(int, int) *datastore.ValueIntSlice
		seeFunc  FieldRestricter
	}{
		{
			"motion submitter",
			ds.User_SubmittedMotionIDsTmpl,
			ds.User_SubmittedMotionIDs,
			Motion{}.see,
		},

		{
			"motion supporter",
			ds.User_SupportedMotionIDsTmpl,
			ds.User_SupportedMotionIDs,
			Motion{}.see,
		},

		{
			"option",
			ds.User_OptionIDsTmpl,
			ds.User_OptionIDs,
			Option{}.see,
		},

		{
			"assignment candidate",
			ds.User_AssignmentCandidateIDsTmpl,
			ds.User_AssignmentCandidateIDs,
			AssignmentCandidate{}.see,
		},

		{
			"speaker",
			ds.User_SpeakerIDsTmpl,
			ds.User_SpeakerIDs,
			Speaker{}.see,
		},

		{
			"poll voted",
			ds.User_PollVotedIDsTmpl,
			ds.User_PollVotedIDs,
			Poll{}.modeB, // Checking field poll/voted_ids that is in modeB and not in see.
		},

		{
			"vote user",
			ds.User_VoteIDsTmpl,
			ds.User_VoteIDs,
			Vote{}.see,
		},

		{
			"vote delegated user",
			ds.User_VoteDelegatedVoteIDsTmpl,
			ds.User_VoteDelegatedVoteIDs,
			Vote{}.see,
		},

		{
			"chat messages",
			ds.User_ChatMessageIDsTmpl,
			ds.User_ChatMessageIDs,
			ChatMessage{}.see,
		},
	}

	for _, r := range requiredObjects {
		for _, meetingID := range r.tmplFunc(userID).ErrorLater(ctx) {
			for _, elementID := range r.elemFunc(userID, meetingID).ErrorLater(ctx) {
				see, err := r.seeFunc(ctx, ds, mperms, elementID)
				if err != nil {
					return false, fmt.Errorf("checking required object %q: %w", r.name, err)
				}

				if see {
					return true, nil
				}
			}
		}
		if err := ds.Err(); err != nil {
			return false, fmt.Errorf("getting object %q: %w", r.name, err)
		}
	}

	return false, nil
}

func (u User) modeB(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	return mperms.UserID() == UserID, nil
}

func (u User) modeD(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return true, nil
	}

	meetingIDs := ds.User_GroupIDsTmpl(UserID).ErrorLater(ctx)
	for _, meetingID := range meetingIDs {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.UserCanManage) {
			return true, nil
		}
	}
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("checking manage in any meeting: %w", err)
	}

	return false, nil
}

func (u User) modeE(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	if mperms.UserID() == UserID {
		return true, nil
	}

	if mperms.UserID() == 0 {
		return false, nil
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return true, nil
	}

	for _, committeeID := range ds.User_CommitteeManagementLevelTmpl(mperms.UserID()).ErrorLater(ctx) {
		committeeManagementLevel := ds.User_CommitteeManagementLevel(mperms.UserID(), committeeID).ErrorLater(ctx)
		if committeeManagementLevel != "can_manage" {
			continue
		}

		userIDs := ds.Committee_UserIDs(committeeID).ErrorLater(ctx)
		for _, uid := range userIDs {
			if UserID == uid {
				return true, nil
			}
		}
	}
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("checking committee management level: %w", err)
	}

	meetingIDs := ds.User_GroupIDsTmpl(UserID).ErrorLater(ctx)
	for _, meetingID := range meetingIDs {
		perms, err := mperms.Meeting(ctx, meetingID)
		if err != nil {
			return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.UserCanManage) {
			return true, nil
		}
	}
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("checking manage in any meeting: %w", err)
	}

	return false, nil
}

func (u User) modeF(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	if mperms.UserID() == UserID {
		return true, nil
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return false, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return true, nil
	}

	return false, nil
}

func (u User) modeG(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	return false, nil
}
