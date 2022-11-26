package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// User handels the restrictions for the user collection.
//
// Y is the request user and X the user, that is requested.
//
// The user Y can see a user X, if at least one condition is true:
//
//	Y==X
//	Y has the OML can_manage_users or higher.
//	There exists a committee where Y has the CML can_manage and X is in committee/user_ids.
//	X is in a group of a meeting where Y has user.can_see.
//	There exists a meeting where Y has the CML can_manage for the meeting's committee X is in meeting/user_ids.
//	There is a related object:
//	    There exists a motion which Y can see and X is a submitter/supporter.
//	    There exists an option which Y can see and X is the linked content object.
//	    There exists an assignment candidate which Y can see and X is the linked user.
//	    There exists a speaker which Y can see and X is the linked user.
//	    There exists a poll where Y can see the poll/voted_ids and X is part of that list.
//	    There exists a vote which Y can see and X is linked in user_id or delegated_user_id.
//	    There exists a chat_message which Y can see and X has sent it (specified by chat_message/user_id).
//	X is linked in one of the relations vote_delegated_$_to_id or vote_delegations_$_from_ids of Y.
//
// Mode A: Y can see X.
//
// Mode B: Y==X.
//
// Mode D: Y can see these fields if at least one condition is true:
//
//	Y has the OML can_manage_users or higher.
//	X is in a group of a meeting where Y has user.can_manage.
//
// Mode E: Y can see these fields if at least one condition is true:
//
//	Y has the OML can_manage_users or higher.
//	There exists a committee where Y has the CML can_manage and X is in committee/user_ids.
//	X is in a group of a meeting where Y has user.can_manage.
//	Y==X.
//
// Mode F: Y has the OML can_manage_users or higher.
//
// Mode G: No one. Not even the superadmin.
//
// Mode H: Like D but the fields are not visible, if the request has a lower
// organization management level then the requested user.
type User struct{}

// MeetingID returns the meetingID for the object.
func (User) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

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
		return never
	case "H":
		return u.modeH
	}
	return nil
}

// SuperAdmin restricts the super admin.
func (User) SuperAdmin(mode string) FieldRestricter {
	if mode == "G" {
		return never
	}
	return Allways
}

func (u User) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, userIDs ...int) ([]int, error) {
	isUserManager, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("check organization management level: %w", err)
	}

	if isUserManager {
		return userIDs, nil
	}

	// Precalculated list of userIDs, that the user can see.
	allowedUserIDs := set.New[int]()
	if mperms.UserID() != 0 {
		allowedUserIDs.Add(mperms.UserID())

		// Get all userIDs of committees, where the request user is manager.
		commiteeIDs, err := perm.ManagementLevelCommittees(ctx, ds, mperms.UserID())
		if err != nil {
			return nil, fmt.Errorf("getting committee ids: %w", err)
		}

		for _, committeeID := range commiteeIDs {
			userIDs, err := ds.Committee_UserIDs(committeeID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("fetching users from committee %d: %w", committeeID, err)
			}
			allowedUserIDs.Add(userIDs...)
		}

		// Getting users where the request users delegated his vote to.
		meetingWithDelegationTo, err := ds.User_VoteDelegatedToIDTmpl(mperms.UserID()).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting meeting ids with vote delegations: %w", err)
		}

		for _, meetingID := range meetingWithDelegationTo {
			delegated, err := ds.User_VoteDelegatedToID(mperms.UserID(), meetingID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting 'vote delegated to' in meeting %d: %w", meetingID, err)
			}
			allowedUserIDs.Add(delegated)
		}

		// Getting users, that delegated his vote to the request user.
		meetingWithDelegationFrom, err := ds.User_VoteDelegationsFromIDsTmpl(mperms.UserID()).Value(ctx)
		for _, meetingID := range meetingWithDelegationFrom {
			delegations, err := ds.User_VoteDelegationsFromIDs(mperms.UserID(), meetingID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting 'vote delegations from' in meeting %d: %w", meetingID, err)
			}
			allowedUserIDs.Add(delegations...)
		}
	}

	return eachCondition(userIDs, func(userID int) (bool, error) {
		if allowedUserIDs.Has(userID) {
			return true, nil
		}

		// Check if the user is in a meeting, where the request user can
		// user.can_see.
		meetingIDs, err := ds.User_GroupIDsTmpl(userID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("fetch meeting ids from requested user %d: %w", userID, err)
		}

		for _, meetingID := range meetingIDs {
			perms, err := mperms.Meeting(ctx, meetingID)
			if err != nil {
				return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.UserCanSee) {
				return true, nil
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
	})
}

// UserRequiredObject represents the reference from a user to other objects.
type UserRequiredObject struct {
	Name     string
	TmplFunc func(int) *dsfetch.ValueIDSlice
	ElemFunc func(int, int) *dsfetch.ValueIntSlice
	SeeFunc  FieldRestricter
}

// RequiredObjects returns all references to other objects from the user.
func (User) RequiredObjects(ds *dsfetch.Fetch) []UserRequiredObject {
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
			Poll{}.see,
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
}

func (User) modeB(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, userIDs ...int) ([]int, error) {
	for _, userID := range userIDs {
		if userID == mperms.UserID() {
			return []int{userID}, nil
		}
	}
	return nil, nil
}

func (User) modeD(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, userIDs ...int) ([]int, error) {
	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return userIDs, nil
	}

	// TODO: group by many meeting.
	return eachCondition(userIDs, func(userID int) (bool, error) {
		meetingIDs := ds.User_GroupIDsTmpl(userID).ErrorLater(ctx)
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
	})
}

func (User) modeE(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, userIDs ...int) ([]int, error) {
	if mperms.UserID() == 0 {
		return nil, nil
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return userIDs, nil
	}

	// TODO: optimize
	return eachCondition(userIDs, func(userID int) (bool, error) {
		if mperms.UserID() == userID {
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

		meetingIDs := ds.User_GroupIDsTmpl(userID).ErrorLater(ctx)
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
	})
}

func (User) modeF(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, userIDs ...int) ([]int, error) {
	isUserManager, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("check organization management level: %w", err)
	}

	if isUserManager {
		return userIDs, nil
	}

	return nil, nil
}

// higherThenOrgaManagement returns true if request equal or higher  then
// request.
//
// An empty string is a valid organization management level for this function
// that has the lowest value.
func higherThenOrgaManagement(request, requested perm.OrganizationManagementLevel) bool {
	toNum := func(level perm.OrganizationManagementLevel) int {
		switch level {
		case perm.OMLNone:
			return 0
		case perm.OMLCanManageUsers:
			return 1
		case perm.OMLCanManageOrganization:
			return 2
		case perm.OMLSuperadmin:
			return 3
		default:
			return 4
		}
	}

	return toNum(request) >= toNum(requested)
}

func (u User) modeH(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, userIDs ...int) ([]int, error) {
	ownOrgaManagementLevel, err := ds.User_OrganizationManagementLevel(mperms.UserID()).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting own managament: %w", err)
	}

	ownLevel := perm.OrganizationManagementLevel(ownOrgaManagementLevel)

	fromD, err := u.modeD(ctx, ds, mperms, userIDs...)
	if err != nil {
		return nil, fmt.Errorf("restriction with mode d: %w", err)
	}

	allowed := make([]int, 0, len(fromD))
	for _, userID := range fromD {
		requestedOrgaManagementLevel, err := ds.User_OrganizationManagementLevel(userID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting orga managament level for user %d: %w", userID, err)
		}

		otherLevel := perm.OrganizationManagementLevel(requestedOrgaManagementLevel)

		if higherThenOrgaManagement(ownLevel, otherLevel) {
			allowed = append(allowed, userID)
		}
	}

	return allowed, nil
}
