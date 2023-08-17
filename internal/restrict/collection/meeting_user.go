package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MeetingUser handels permissions for the collection meeting_user.
//
// A User can see a MeetingUser if he can see the user.
//
// Mode A: The user can see the related user.
//
// Mode B: The request user is the related user.
//
// Mode D: Y can see these fields if at least one condition is true:
//
//	The request user has the OML can_manage_users or higher.
//	The request user has user.can_manage in the meeting
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

	}
	return nil
}

func (m MeetingUser) see(ctx context.Context, ds *dsfetch.Fetch, meetingUserIDs ...int) ([]int, error) {
	userToMeetingUsers := make(map[int][]int)
	for _, id := range meetingUserIDs {
		userID, err := ds.MeetingUser_UserID(id).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting user id for meeting_user %d: %w", id, err)
		}

		userToMeetingUsers[userID] = append(userToMeetingUsers[userID], id)
	}

	userIDs := make([]int, 0, len(userToMeetingUsers))
	for userID := range userToMeetingUsers {
		userIDs = append(userIDs, userID)
	}

	allowedUserIDs, err := Collection(ctx, User{}.Name()).Modes("A")(ctx, ds, userIDs...)
	if err != nil {
		return nil, fmt.Errorf("checking user restrictions: %w", err)
	}

	var allowedIDs []int
	for _, allowedUserID := range allowedUserIDs {
		allowedIDs = append(allowedIDs, userToMeetingUsers[allowedUserID]...)
	}

	return allowedIDs, nil
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
