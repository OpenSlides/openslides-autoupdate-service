package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
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

func (m MeetingUser) see(ctx context.Context, fetcher *dsfetch.Fetch, meetingUserIDs []int) ([]attribute.Func, error) {
	return canSeeRelatedCollection(ctx, fetcher, fetcher.MeetingUser_UserID, Collection(ctx, User{}).Modes("A"), meetingUserIDs)
}

func (MeetingUser) modeB(ctx context.Context, fetcher *dsfetch.Fetch, meetingUserIDs []int) ([]attribute.Func, error) {
	userIDs := make([]int, len(meetingUserIDs))
	for i, id := range meetingUserIDs {
		if id == 0 {
			continue
		}
		fetcher.MeetingUser_UserID(id).Lazy(&userIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user ids: %w", err)
	}

	out := make([]attribute.Func, len(meetingUserIDs))
	for i, id := range meetingUserIDs {
		if id == 0 {
			continue
		}
		out[i] = attribute.FuncUserIDs([]int{userIDs[i]})
	}

	return out, nil
}

func (m MeetingUser) modeD(ctx context.Context, fetcher *dsfetch.Fetch, meetingUserIDs []int) ([]attribute.Func, error) {
	oml := attribute.FuncGlobalLevel(perm.OMLCanManageUsers)

	canManageUser, err := meetingPerm(ctx, fetcher, m, meetingUserIDs, perm.UserCanManage)
	if err != nil {
		return nil, fmt.Errorf("checking meeting perm user can manage: %w", err)
	}

	out := make([]attribute.Func, len(meetingUserIDs))
	for i := range meetingUserIDs {
		out[i] = attribute.FuncOr(
			oml,
			canManageUser[i],
		)
	}
	return out, nil
}
