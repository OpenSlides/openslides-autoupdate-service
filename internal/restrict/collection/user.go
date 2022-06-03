package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
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

// MeetingID returns the meetingID for the object.
func (u User) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns the field restriction for each mode.
func (u User) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return todoToSingle(u.see)
	case "B":
		return todoToSingle(u.modeB)
	case "D":
		return todoToSingle(u.modeD)
	case "E":
		return todoToSingle(u.modeE)
	case "G":
		return never
	}
	return nil
}

// SuperAdmin restricts the super admin.
func (u User) SuperAdmin(mode string) FieldRestricter {
	if mode == "G" {
		return never
	}
	return Allways
}

func (u User) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, userID int) (bool, error) {
	if mperms.UserID() == userID {
		return true, nil
	}

	if mperms.UserID() != 0 {
		canManageUsers, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
		if err != nil {
			return false, fmt.Errorf("get organization level: %w", err)
		}

		if canManageUsers {
			return true, nil
		}

		commiteeIDs, err := perm.ManagementLevelCommittees(ctx, ds, mperms.UserID())
		if err != nil {
			return false, fmt.Errorf("getting committee ids: %w", err)
		}

		for _, committeeID := range commiteeIDs {
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

		committeeManager, err := perm.HasCommitteeManagementLevel(ctx, ds, mperms.UserID(), cid)
		if err != nil {
			return false, fmt.Errorf("getting committee management level: %w", err)
		}

		if committeeManager {
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

	for _, r := range u.RequiredObjects(ds) {
		for _, meetingID := range r.TmplFunc(userID).ErrorLater(ctx) {
			ids := r.ElemFunc(userID, meetingID).ErrorLater(ctx)

			if len(ids) == 0 {
				continue
			}

			allowedIDs, err := r.SeeFunc(ctx, ds, mperms, ids...)
			if err != nil {
				return false, fmt.Errorf("checking required object %q: %w", r.Name, err)
			}

			if len(allowedIDs) > 0 {
				return true, nil
			}

		}
		if err := ds.Err(); err != nil {
			return false, fmt.Errorf("getting object %q: %w", r.Name, err)
		}
	}

	return false, nil
}

// UserRequiredObject represents the reference from a user to other objects.
type UserRequiredObject struct {
	Name     string
	TmplFunc func(int) *dsfetch.ValueIDSlice
	ElemFunc func(int, int) *dsfetch.ValueIntSlice
	SeeFunc  FieldRestricter
}

// RequiredObjects returns all references to other objects from the user.
func (u User) RequiredObjects(ds *dsfetch.Fetch) []UserRequiredObject {
	return []UserRequiredObject{
		{
			"motion submitter",
			ds.User_SubmittedMotionIDsTmpl,
			ds.User_SubmittedMotionIDs,
			MotionSubmitter{}.see,
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
			todoToSingle(Poll{}.see),
		},

		{
			"vote user",
			ds.User_VoteIDsTmpl,
			ds.User_VoteIDs,
			todoToSingle(Vote{}.see),
		},

		{
			"vote delegated user",
			ds.User_VoteDelegatedVoteIDsTmpl,
			ds.User_VoteDelegatedVoteIDs,
			todoToSingle(Vote{}.see),
		},

		{
			"chat messages",
			ds.User_ChatMessageIDsTmpl,
			ds.User_ChatMessageIDs,
			ChatMessage{}.see,
		},
	}
}

func (u User) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	return mperms.UserID() == UserID, nil
}

func (u User) modeD(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, UserID int) (bool, error) {
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

func (u User) modeE(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, UserID int) (bool, error) {
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

	commiteeIDs, err := perm.ManagementLevelCommittees(ctx, ds, mperms.UserID())
	if err != nil {
		return false, fmt.Errorf("getting committee ids: %w", err)
	}

	for _, committeeID := range commiteeIDs {
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
