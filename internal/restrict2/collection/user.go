package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
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
//	X is linked in one of the relations vote_delegated_$_to_id or vote_delegations_$_from_ids of Y.
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
// Mode H: Like D but the fields are not visible, if the request has a lower
// organization management level then the requested user.
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

func (u User) see(ctx context.Context, fetcher *dsfetch.Fetch, userIDs []int) ([]attribute.Func, error) {
	userManager := attribute.FuncOrgaLevel(perm.OMLCanManageUsers)

	inCommitteeList := make([][]int, len(userIDs))
	inMeetingList := make([][]int, len(userIDs))
	for i, id := range userIDs {
		if id == 0 {
			continue
		}

		fetcher.User_CommitteeIDs(id).Lazy(&inCommitteeList[i])
		fetcher.User_MeetingIDs(id).Lazy(&inMeetingList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching in committee and in meeting for every requested user: %w", err)
	}

	committeeIDs := set.New[int]() // IDs of all committees where the requested users are part of
	for _, cids := range inCommitteeList {
		if cids == nil {
			continue
		}

		committeeIDs.Add(cids...)
	}

	committeeManagers, err := fetchCommitteeManagers(ctx, fetcher, committeeIDs.List())
	if err != nil {
		return nil, fmt.Errorf("calculating committee managers: %w", err)
	}

	result := make([]attribute.Func, len(userIDs))
	for i, userID := range userIDs {
		if userID == 0 {
			continue
		}

		var usersCommitteeManagers []int
		for _, committeeID := range inCommitteeList[i] {
			usersCommitteeManagers = append(usersCommitteeManagers, committeeManagers[committeeID]...)
		}

		var canSeeGroups []int
		for _, meetingID := range inMeetingList[i] {
			groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
			if err != nil {
				return nil, fmt.Errorf("group map for user %d in meeting %d: %w", userID, meetingID, err)
			}

			canSeeGroups = append(canSeeGroups, groupMap[perm.UserCanSee]...)
		}

		result[i] = attribute.FuncOr(
			attribute.FuncUserIDs([]int{userID}),
			userManager,
			attribute.FuncUserIDs(usersCommitteeManagers),
			attribute.FuncInGroup(canSeeGroups),
			// TODO: There is a related object...
		)
	}
	return result, nil
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
			Collection(ctx, MotionSubmitter{}).Modes("A"),
			false,
		},

		{
			"motion supporter",
			ds.MeetingUser_SupportedMotionIDs,
			Collection(ctx, Motion{}).Modes("C"),
			false,
		},

		// {
		// 	"option",
		// 	ds.User_OptionIDs,
		// 	Collection(ctx, Option{}.Name()).Modes("A"),
		// 	true,
		// },

		// {
		// 	"assignment candidate",
		// 	ds.MeetingUser_AssignmentCandidateIDs,
		// 	Collection(ctx, AssignmentCandidate{}.Name()).Modes("A"),
		// 	false,
		// },

		// {
		// 	"speaker",
		// 	ds.MeetingUser_SpeakerIDs,
		// 	Collection(ctx, Speaker{}.Name()).Modes("A"),
		// 	false,
		// },

		// {
		// 	"poll voted",
		// 	ds.User_PollVotedIDs,
		// 	Collection(ctx, Poll{}.Name()).Modes("A"),
		// 	true,
		// },

		// {
		// 	"vote user",
		// 	ds.User_VoteIDs,
		// 	Collection(ctx, Vote{}.Name()).Modes("A"),
		// 	true,
		// },

		// {
		// 	"vote delegated user",
		// 	ds.MeetingUser_VoteDelegationsFromIDs,
		// 	Collection(ctx, Vote{}.Name()).Modes("A"),
		// 	false,
		// },

		// {
		// 	"chat messages",
		// 	ds.MeetingUser_ChatMessageIDs,
		// 	Collection(ctx, ChatMessage{}.Name()).Modes("A"),
		// 	false,
		// },
	}
}

func (u User) modeB(ctx context.Context, fetcher *dsfetch.Fetch, userIDs []int) ([]attribute.Func, error) {
	userManager := attribute.FuncOrgaLevel(perm.OMLCanManageUsers)

	inCommitteeList := make([][]int, len(userIDs))
	inMeetingList := make([][]int, len(userIDs))
	for i, id := range userIDs {
		if id == 0 {
			continue
		}

		fetcher.User_CommitteeIDs(id).Lazy(&inCommitteeList[i])
		fetcher.User_MeetingIDs(id).Lazy(&inMeetingList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching in committee and in meeting for every requested user: %w", err)
	}

	committeeIDs := set.New[int]() // IDs of all committees where the requested users are part of
	for _, cids := range inCommitteeList {
		if cids == nil {
			continue
		}

		committeeIDs.Add(cids...)
	}

	committeeManagers, err := fetchCommitteeManagers(ctx, fetcher, committeeIDs.List())
	if err != nil {
		return nil, fmt.Errorf("calculating committee managers: %w", err)
	}

	result := make([]attribute.Func, len(userIDs))
	for i, userID := range userIDs {
		if userID == 0 {
			continue
		}

		var usersCommitteeManagers []int
		for _, committeeID := range inCommitteeList[i] {
			usersCommitteeManagers = append(usersCommitteeManagers, committeeManagers[committeeID]...)
		}

		var canSeeGroups []int
		for _, meetingID := range inMeetingList[i] {
			groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
			if err != nil {
				return nil, fmt.Errorf("group map for user %d in meeting %d: %w", userID, meetingID, err)
			}

			canSeeGroups = append(canSeeGroups, groupMap[perm.UserCanSeeSensitiveData]...)
		}

		result[i] = attribute.FuncOr(
			attribute.FuncUserIDs([]int{userID}),
			userManager,
			attribute.FuncUserIDs(usersCommitteeManagers),
			attribute.FuncInGroup(canSeeGroups),
		)
	}
	return result, nil
}

func (u User) modeD(ctx context.Context, fetcher *dsfetch.Fetch, userIDs []int) ([]attribute.Func, error) {
	userManager := attribute.FuncOrgaLevel(perm.OMLCanManageUsers)

	inMeetingList := make([][]int, len(userIDs))
	for i, id := range userIDs {
		if id == 0 {
			continue
		}

		fetcher.User_MeetingIDs(id).Lazy(&inMeetingList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user data: %w", err)
	}

	result := make([]attribute.Func, len(userIDs))
	for i, userID := range userIDs {
		if userID == 0 {
			continue
		}

		var canSeeGroups []int
		for _, meetingID := range inMeetingList[i] {
			groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
			if err != nil {
				return nil, fmt.Errorf("group map for user %d in meeting %d: %w", userID, meetingID, err)
			}

			canSeeGroups = append(canSeeGroups, groupMap[perm.UserCanManage]...)
		}

		result[i] = attribute.FuncOr(
			userManager,
			attribute.FuncInGroup(canSeeGroups),
		)
	}
	return result, nil
}

func (u User) modeE(ctx context.Context, fetcher *dsfetch.Fetch, userIDs []int) ([]attribute.Func, error) {
	userManager := attribute.FuncOrgaLevel(perm.OMLCanManageUsers)

	inCommitteeList := make([][]int, len(userIDs))
	inMeetingList := make([][]int, len(userIDs))
	for i, id := range userIDs {
		if id == 0 {
			continue
		}

		fetcher.User_CommitteeIDs(id).Lazy(&inCommitteeList[i])
		fetcher.User_MeetingIDs(id).Lazy(&inMeetingList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user data: %w", err)
	}

	committeeIDs := set.New[int]() // IDs of all committees where the requested users are part of
	for _, cids := range inCommitteeList {
		if cids == nil {
			continue
		}
		committeeIDs.Add(cids...)
	}

	committeeManagers, err := fetchCommitteeManagers(ctx, fetcher, committeeIDs.List())
	if err != nil {
		return nil, fmt.Errorf("calculating committee managers: %w", err)
	}

	result := make([]attribute.Func, len(userIDs))
	for i, userID := range userIDs {
		if userID == 0 {
			continue
		}

		var usersCommitteeManagers []int
		for _, committeeID := range inCommitteeList[i] {
			usersCommitteeManagers = append(usersCommitteeManagers, committeeManagers[committeeID]...)
		}

		var canSeeGroups []int
		for _, meetingID := range inMeetingList[i] {
			groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
			if err != nil {
				return nil, fmt.Errorf("group map for user %d in meeting %d: %w", userID, meetingID, err)
			}

			canSeeGroups = append(canSeeGroups, groupMap[perm.UserCanSee]...)
		}

		result[i] = attribute.FuncOr(
			attribute.FuncUserIDs([]int{userID}),
			userManager,
			attribute.FuncUserIDs(usersCommitteeManagers),
			attribute.FuncInGroup(canSeeGroups),
		)
	}
	return result, nil
}

func (u User) modeF(ctx context.Context, fetcher *dsfetch.Fetch, userIDs []int) ([]attribute.Func, error) {
	userManager := attribute.FuncOrgaLevel(perm.OMLCanManageUsers)
	return attributeFuncList(userIDs, userManager), nil
}

func (u User) modeH(ctx context.Context, fetcher *dsfetch.Fetch, userIDs []int) ([]attribute.Func, error) {
	userOrgaLevel := make([]string, len(userIDs))
	for i, id := range userIDs {
		if id == 0 {
			continue
		}
		fetcher.User_OrganizationManagementLevel(id).Lazy(&userOrgaLevel[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching orga levels: %w", err)
	}

	result, err := Collection(ctx, u).Modes("D")(ctx, fetcher, userIDs)
	if err != nil {
		return nil, fmt.Errorf("check like D: %w", err)
	}

	for i, id := range userIDs {
		if id == 0 {
			continue
		}
		result[i] = attribute.FuncAnd(
			attribute.FuncOrgaLevel(perm.OrganizationManagementLevel(userOrgaLevel[i])),
			result[i],
		)
	}
	return result, nil
}

// fetchCommitteeManagers returns for a list of committeeIDs the userIDs of its
// managers.
func fetchCommitteeManagers(ctx context.Context, fetcher *dsfetch.Fetch, committeeIDs []int) (map[int][]int, error) {
	managers := make([][]int, len(committeeIDs))
	for i, committeeID := range committeeIDs {
		fetcher.Committee_ManagerIDs(committeeID).Lazy(&managers[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting committee manager ids: %w", err)
	}

	out := make(map[int][]int, len(committeeIDs))
	for i, userIDs := range managers {
		out[committeeIDs[i]] = userIDs
	}

	return out, nil
}
