package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// MeetingUser handels permissions for the collection meeting_user.
//
// A User can see a MeetingUser if the request user is the related user or the
// request user has user.can_see.
//
// Mode A: The user can see the meeting_user.
//
// Mode B: The request user is the related user.
//
// Mode D: Y can see these fields if at least one condition is true:
//
//	The request user has the OML can_manage_users or higher.
//	The request user has user.can_manage in the meeting
//
// Mode E: Y can see these fields if at least one condition is true:
//
//	Y has the permissoin can_see_sensible_data.
//	Y is the related user.
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
	allowed, err := meetingPerm(ctx, ds, m, meetingUserIDs, perm.UserCanSee)
	if err != nil {
		return nil, fmt.Errorf("checking permission: %w", err)
	}

	if len(meetingUserIDs) == len(allowed) {
		// Fast exit if all requested users are allowed.
		return allowed, nil
	}

	allowedSet := set.New(allowed...)

	userIDs := make([]int, len(meetingUserIDs))
	for i, meetingUserID := range meetingUserIDs {
		ds.MeetingUser_UserID(meetingUserID).Lazy(&userIDs[i])
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user ids: %w", err)
	}

	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	for i, meetingUserID := range meetingUserIDs {
		if allowedSet.Has(meetingUserID) {
			continue
		}

		if userIDs[i] == requestUser {
			allowed = append(allowed, meetingUserID)
		}
	}

	return allowed, nil
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
