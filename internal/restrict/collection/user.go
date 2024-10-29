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
//	There is a related object:
//	    There exists a motion which Y can see and X is a submitter/supporter.
//	    There exists an option which Y can see and X is the linked content object.
//	    There exists an assignment candidate which Y can see and X is the linked user.
//	    There exists a speaker which Y can see and X is the linked user.
//	    There exists a poll where Y can see the poll/voted_ids and X is part of that list.
//	    There exists a vote which Y can see and X is linked in user_id or delegated_user_id.
//	    There exists a chat_message which Y can see and X has sent it (specified by chat_message/user_id).
//	X is linked in one of the relations vote_delegated_to_id or vote_delegations_from_ids of Y.
//
// Mode A: Y can see X.
//
// Mode B:
//
//	Y==X
//	Y has the OML can_manage_users or higher.
//	There exists a committee where Y has the CML can_manage and X is in committee/user_ids.
//	X is in a group of a meeting where Y has user.can_see_sensitive_data.
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
// Mode H: The fields can be seen if one of the following conditions is true:
//   - Any of the conditions in D or
//   - There exists a committee where Y has the CML can_manage and X is in committee/user_ids
//     But the fields are never visible, if the request has a lower organization management level than the requested user.
type User struct{}

// Name returns the collection name.
func (u User) Name() string {
	return "user"
}

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

func (u User) see(ctx context.Context, ds *dsfetch.Fetch, userIDs ...int) ([]int, error) {
	requestUserID, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	isUserManager, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUserID, perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("check organization management level: %w", err)
	}

	if isUserManager {
		return userIDs, nil
	}

	// Precalculated list of userIDs, that the user can see.
	allowedUserIDs := set.New[int]()
	if requestUserID != 0 {
		allowedUserIDs.Add(requestUserID)

		// Get all userIDs of committees, where the request user is manager.
		commiteeIDs, err := perm.ManagementLevelCommittees(ctx, ds, requestUserID)
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

		meetingUserIDs, err := ds.User_MeetingUserIDs(requestUserID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting meeting user: %w", err)
		}

		for _, meetingUserID := range meetingUserIDs {
			// Getting users where the request users delegated his vote to.
			delegatedToMeetingUserID, err := ds.MeetingUser_VoteDelegatedToID(meetingUserID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting 'vote delegated to' for meeting_user %d: %w", meetingUserID, err)
			}

			if id, ok := delegatedToMeetingUserID.Value(); ok {
				delegatedUser, err := ds.MeetingUser_UserID(id).Value(ctx)
				if err != nil {
					return nil, fmt.Errorf("getting delegated user: %w", err)
				}

				allowedUserIDs.Add(delegatedUser)
			}

			// Getting users, that delegated his vote to the request user.
			delegationsFromMeetingUserID, err := ds.MeetingUser_VoteDelegationsFromIDs(meetingUserID).Value(ctx)
			if err != nil {
				return nil, fmt.Errorf("getting 'vote delegations from' for meeting_user %d: %w", meetingUserID, err)
			}

			for _, delegateMeetingUserID := range delegationsFromMeetingUserID {
				delegationUserID, err := ds.MeetingUser_UserID(delegateMeetingUserID).Value(ctx)
				if err != nil {
					return nil, fmt.Errorf("getting delegation user id: %w", err)
				}

				allowedUserIDs.Add(delegationUserID)
			}
		}
	}

	return eachCondition(userIDs, func(otherUserID int) (bool, error) {
		if allowedUserIDs.Has(otherUserID) {
			return true, nil
		}

		// Check if the user is in a meeting, where the request user can
		// user.can_see.
		otherUserMeetingUserIDs, err := ds.User_MeetingUserIDs(otherUserID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("fetch meeting ids from requested user %d: %w", otherUserID, err)
		}

		for _, meetingUserID := range otherUserMeetingUserIDs {
			meetingID, err := ds.MeetingUser_MeetingID(meetingUserID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting meeting: %w", err)
			}

			perms, err := perm.FromContext(ctx, meetingID)
			if err != nil {
				return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.UserCanSee) {
				return true, nil
			}
		}

		meetingUserIDs, err := ds.User_MeetingUserIDs(otherUserID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("fetching meeting_user of %d: %w", otherUserID, err)
		}

		for i, meetingUserID := range meetingUserIDs {
			for _, r := range u.RequiredObjects(ctx, ds) {
				id := meetingUserID
				if r.OnUser {
					if i > 0 {
						// Some requiredObjects have the realtion directly on
						// the user object. They only have to be checked once in
						// this loop
						continue
					}
					id = otherUserID
				}

				ids, err := r.ElemFunc(id).Value(ctx)
				if err != nil {
					return false, fmt.Errorf("getting ids for %s: %w", r.Name, err)
				}

				allowedIDs, err := r.SeeFunc(ctx, ds, ids...)
				if err != nil {
					meetingUserOrUser := "meetingUserID"
					if r.OnUser {
						meetingUserOrUser = "user"
					}
					return false, fmt.Errorf("checking required object %s on %s %d: %w", r.Name, meetingUserOrUser, id, err)
				}
				if len(allowedIDs) > 0 {
					return true, nil
				}

			}
		}

		return false, nil
	})
}

// UserRequiredObject represents the reference from a user to other objects.
type UserRequiredObject struct {
	Name     string
	ElemFunc func(int) *dsfetch.ValueIntSlice
	SeeFunc  FieldRestricter
	OnUser   bool // Tells, if the relation is via meeting_user_id or user_id
}

// RequiredObjects returns all references to other objects from the user.
func (User) RequiredObjects(ctx context.Context, ds *dsfetch.Fetch) []UserRequiredObject {
	return []UserRequiredObject{
		{
			"motion submitter",
			ds.MeetingUser_MotionSubmitterIDs,
			Collection(ctx, MotionSubmitter{}.Name()).Modes("A"),
			false,
		},

		{
			"motion supporter",
			ds.MeetingUser_SupportedMotionIDs,
			Collection(ctx, Motion{}.Name()).Modes("C"),
			false,
		},

		{
			"option",
			ds.User_OptionIDs,
			Collection(ctx, Option{}.Name()).Modes("A"),
			true,
		},

		{
			"assignment candidate",
			ds.MeetingUser_AssignmentCandidateIDs,
			Collection(ctx, AssignmentCandidate{}.Name()).Modes("A"),
			false,
		},

		{
			"speaker",
			ds.MeetingUser_SpeakerIDs,
			Collection(ctx, Speaker{}.Name()).Modes("A"),
			false,
		},

		{
			"poll voted",
			ds.User_PollVotedIDs,
			Collection(ctx, Poll{}.Name()).Modes("A"),
			true,
		},

		{
			"vote user",
			ds.User_VoteIDs,
			Collection(ctx, Vote{}.Name()).Modes("A"),
			true,
		},

		{
			"chat messages",
			ds.MeetingUser_ChatMessageIDs,
			Collection(ctx, ChatMessage{}.Name()).Modes("A"),
			false,
		},
	}
}

func (u User) modeB(ctx context.Context, ds *dsfetch.Fetch, userIDs ...int) ([]int, error) {
	requestUserID, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	isUserManager, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUserID, perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("check organization management level: %w", err)
	}

	if isUserManager {
		return userIDs, nil
	}

	// Precalculated list of userIDs, that the user can see.
	allowedUserIDs := set.New[int]()
	if requestUserID != 0 {
		allowedUserIDs.Add(requestUserID)

		// Get all userIDs of committees, where the request user is manager.
		commiteeIDs, err := perm.ManagementLevelCommittees(ctx, ds, requestUserID)
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

	}

	return eachCondition(userIDs, func(otherUserID int) (bool, error) {
		if allowedUserIDs.Has(otherUserID) {
			return true, nil
		}

		// Check if the user is in a meeting, where the request user can
		// user.can_see_sensitive_data.
		otherUserMeetingUserIDs, err := ds.User_MeetingUserIDs(otherUserID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("fetch meeting ids from requested user %d: %w", otherUserID, err)
		}

		for _, meetingUserID := range otherUserMeetingUserIDs {
			meetingID, err := ds.MeetingUser_MeetingID(meetingUserID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting meeting: %w", err)
			}

			perms, err := perm.FromContext(ctx, meetingID)
			if err != nil {
				return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.UserCanSeeSensitiveData) {
				return true, nil
			}
		}

		return false, nil
	})
}

func (User) modeD(ctx context.Context, ds *dsfetch.Fetch, userIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return userIDs, nil
	}

	return eachCondition(userIDs, func(otherUserID int) (bool, error) {
		otherMeetingUserIDs, err := ds.User_MeetingUserIDs(otherUserID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("get meeting ids: %w", err)
		}

		for _, muid := range otherMeetingUserIDs {
			meetingID, err := ds.MeetingUser_MeetingID(muid).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting meeting id: %w", err)
			}

			perms, err := perm.FromContext(ctx, meetingID)
			if err != nil {
				return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.UserCanManage) {
				return true, nil
			}
		}

		return false, nil
	})
}

func (User) modeE(ctx context.Context, ds *dsfetch.Fetch, userIDs ...int) ([]int, error) {
	requestUserID, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	if requestUserID == 0 {
		return nil, nil
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUserID, perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return userIDs, nil
	}

	return eachCondition(userIDs, func(otherUserID int) (bool, error) {
		if requestUserID == otherUserID {
			return true, nil
		}

		commiteeIDs, err := perm.ManagementLevelCommittees(ctx, ds, requestUserID)
		if err != nil {
			return false, fmt.Errorf("getting committee ids: %w", err)
		}

		for _, committeeID := range commiteeIDs {
			userIDs, err := ds.Committee_UserIDs(committeeID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting users of committee: %w", err)
			}

			for _, uid := range userIDs {
				if otherUserID == uid {
					return true, nil
				}
			}
		}

		otherUserMeetingUserIDs, err := ds.User_MeetingUserIDs(otherUserID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting meeting user ids: %w", err)
		}

		for _, otherUserMeetingUserID := range otherUserMeetingUserIDs {
			meetingID, err := ds.MeetingUser_MeetingID(otherUserMeetingUserID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting meeting ID: %w", err)
			}

			perms, err := perm.FromContext(ctx, meetingID)
			if err != nil {
				return false, fmt.Errorf("checking permissions of meeting %d: %w", meetingID, err)
			}

			if perms.Has(perm.UserCanManage) {
				return true, nil
			}
		}

		return false, nil
	})
}

func (User) modeF(ctx context.Context, ds *dsfetch.Fetch, userIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	isUserManager, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageUsers)
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

func (u User) modeH(ctx context.Context, ds *dsfetch.Fetch, userIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	if requestUser == 0 {
		return nil, nil
	}

	ownOrgaManagementLevel, err := ds.User_OrganizationManagementLevel(requestUser).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting own managament: %w", err)
	}

	ownLevel := perm.OrganizationManagementLevel(ownOrgaManagementLevel)

	otherLevelStr := make([]string, len(userIDs))
	for i, userID := range userIDs {
		ds.User_OrganizationManagementLevel(userID).Lazy(&otherLevelStr[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting organization management level of other users: %w", err)
	}

	otherLevel := make([]perm.OrganizationManagementLevel, len(userIDs))
	for i, str := range otherLevelStr {
		otherLevel[i] = perm.OrganizationManagementLevel(str)
	}

	fromD, err := u.modeD(ctx, ds, userIDs...)
	if err != nil {
		return nil, fmt.Errorf("restriction with mode d: %w", err)
	}

	// Get all userIDs of committees, where the request user is manager.
	commiteeIDs, err := perm.ManagementLevelCommittees(ctx, ds, requestUser)
	if err != nil {
		return nil, fmt.Errorf("getting committee ids: %w", err)
	}

	committeeUsers := make([][]int, len(commiteeIDs))
	for i, commiteeID := range commiteeIDs {
		ds.Committee_UserIDs(commiteeID).Lazy(&committeeUsers[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting committee user ids: %w", err)
	}

	allUsers := set.NewWithSize(len(userIDs), fromD...)
	for _, fromCommittee := range committeeUsers {
		allUsers.Add(fromCommittee...)
	}

	// Do the logic
	allowed := make([]int, 0, allUsers.Len())
	for i, userID := range userIDs {
		if !higherThenOrgaManagement(ownLevel, otherLevel[i]) {
			continue
		}

		if !allUsers.Has(userID) {
			continue
		}

		allowed = append(allowed, userID)
	}

	return allowed, nil
}
