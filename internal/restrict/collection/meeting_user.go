package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MeetingUser handels permissions for the collection meeting_user.
// Y = the request user; X = the requested user.
//
// A User can see a meeting_user, if
//
//	the request user is the related user,
//	the request user has user.can_see,
//	X is linked in one of the relations vote_delegated_to_id or vote_delegations_from_ids of Y or
//	there is a related object:
//	  There exists a motion which Y can see and X is a submitter/supporter.
//	  There exists an option which Y can see and X is the linked content object.
//	  There exists an assignment candidate which Y can see and X is the linked user.
//	  There exists a speaker which Y can see and X is the linked user.
//	  There exists a poll where Y can see the poll/voted_ids and X is part of that list.
//	  There exists a vote which Y can see and X is linked in user_id or delegated_user_id.
//	  There exists a chat_message which Y can see and X has sent it (specified by chat_message/user_id).
//
// Mode A: Can see.
//
// Mode B: The request user is the related user.
//
// Mode D: Y can see these fields if
//   - the request user has the OML can_manage_users or higher or
//   - the request user has user.can_manage in the meeting.
//
// Mode E: Y can see these fields if
//   - Y has the permissoin can_see_sensible_data or
//   - Y is the related user.
type MeetingUser struct{}

// Name returns the collection name.
func (m MeetingUser) Name() string {
	return "meeting_user"
}

// MeetingID returns the meetingID for the object.
func (m MeetingUser) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	mid, err := ds.MeetingUser_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meeting_user id: %w", err)
	}
	return mid, true, nil
}

// Modes returns the field modes for the collection mediafile.
func (m MeetingUser) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see

	case "B":
		return m.modeB

	case "D":
		return m.modeD

	case "E":
		return m.modeE

	}
	return nil
}

func (m MeetingUser) see(ctx context.Context, ds *dsfetch.Fetch, meetingUserIDs ...int) ([]int, error) {
	requestUserID, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	meetingToMeetingUser := make(map[int]int)
	if requestUserID != 0 {
		requestMeetingUserIDs, err := ds.User_MeetingUserIDs(requestUserID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting meeting_user for request user: %w", err)
		}

		meetingUserMeetingIDs := make([]int, len(requestMeetingUserIDs))
		for i, meetingUserID := range requestMeetingUserIDs {
			ds.MeetingUser_MeetingID(meetingUserID).Lazy(&meetingUserMeetingIDs[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return nil, fmt.Errorf("fetching meeting ids of request users meeting user: %w", err)
		}

		meetingToMeetingUser = make(map[int]int, len(meetingUserMeetingIDs))
		for i, meetingID := range meetingUserMeetingIDs {
			meetingToMeetingUser[meetingID] = requestMeetingUserIDs[i]
		}
	}

	return eachMeeting(ctx, ds, m, meetingUserIDs, func(meetingID int, meetingUserIDs []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
		}

		if perms.Has(perm.UserCanSee) {
			return meetingUserIDs, nil
		}

		return eachCondition(meetingUserIDs, func(meetingUserID int) (bool, error) {
			userID, err := ds.MeetingUser_UserID(meetingUserID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("fetching user id: %w", err)
			}

			if userID == requestUserID {
				return true, nil
			}

			delegatedToMeetingUserID, err := ds.MeetingUser_VoteDelegatedToID(meetingUserID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting 'vote delegated to' for meeting_user %d: %w", meetingUserID, err)
			}

			if id, ok := delegatedToMeetingUserID.Value(); ok {
				if meetingToMeetingUser[meetingID] == id {
					return true, nil
				}
			}

			// Getting users, that delegated his vote to the request user.
			delegationsFromMeetingUserID, err := ds.MeetingUser_VoteDelegationsFromIDs(meetingUserID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting 'vote delegations from' for meeting_user %d: %w", meetingUserID, err)
			}

			for _, delegateMeetingUserID := range delegationsFromMeetingUserID {
				if meetingToMeetingUser[meetingID] == delegateMeetingUserID {
					return true, nil
				}
			}

			var u User // TODO: Remove me

			for _, r := range u.RequiredObjects(ctx, ds) {
				id := meetingUserID
				if r.OnUser {
					id = userID
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

			return false, nil
		})
	})
}

func (MeetingUser) modeB(ctx context.Context, ds *dsfetch.Fetch, meetingUserIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachRelationField(ctx, ds.MeetingUser_UserID, meetingUserIDs, func(userID int, ids []int) ([]int, error) {
		if userID == requestUser {
			return ids, nil
		}

		return nil, nil
	})
}

func (m MeetingUser) modeD(ctx context.Context, ds *dsfetch.Fetch, meetingUserIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	canManage, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageUsers)
	if err != nil {
		return nil, fmt.Errorf("cheching oml: %w", err)
	}

	if canManage {
		return meetingUserIDs, nil
	}

	return meetingPerm(ctx, ds, m, meetingUserIDs, perm.UserCanManage)
}

func (m MeetingUser) modeE(ctx context.Context, ds *dsfetch.Fetch, meetingUserIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachMeeting(ctx, ds, m, meetingUserIDs, func(meetingID int, idsInMeeting []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permission: %w", err)
		}

		if perms.Has(perm.UserCanSeeSensitiveData) {
			return idsInMeeting, nil
		}

		return eachRelationField(ctx, ds.MeetingUser_UserID, idsInMeeting, func(userID int, ids []int) ([]int, error) {
			if userID == requestUser {
				return ids, nil
			}

			return nil, nil
		})
	})
}
