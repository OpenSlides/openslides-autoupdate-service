package attribute

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/flow"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// UserAttributes are some values from a user.
type UserAttributes struct {
	UserID            int
	GroupIDs          set.Set[int]
	OrgaLevel         perm.OrganizationManagementLevel
	IsCommitteManager bool
}

// NewUserAttributes initializes a UserAttributes object.
func NewUserAttributes(ctx context.Context, getter flow.Getter, userID int) (UserAttributes, error) {
	var zero UserAttributes
	fetcher := dsfetch.New(getter)

	if userID == 0 {
		return zero, nil
	}

	var meetingUserIDs []int
	var globalLevelStr string
	var committeeAsManager []int
	fetcher.User_OrganizationManagementLevel(userID).Lazy(&globalLevelStr)
	fetcher.User_MeetingUserIDs(userID).Lazy(&meetingUserIDs)
	fetcher.User_CommitteeManagementIDs(userID).Lazy(&committeeAsManager)

	if err := fetcher.Execute(ctx); err != nil {
		return zero, fmt.Errorf("getting meeting ids,  global and committee level for user %d: %w", userID, err)
	}

	groupIDList := make([][]int, len(meetingUserIDs))
	for i, muid := range meetingUserIDs {
		fetcher.MeetingUser_GroupIDs(muid).Lazy(&groupIDList[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return zero, fmt.Errorf("getting group IDs %d: %w", userID, err)
	}

	// TODO: Test if a sorted slice would be faster.
	groupIDs := set.New[int]()
	for _, idList := range groupIDList {
		groupIDs.Add(idList...)
	}

	return UserAttributes{
		UserID:            userID,
		GroupIDs:          groupIDs,
		OrgaLevel:         perm.OrganizationManagementLevel(globalLevelStr),
		IsCommitteManager: len(committeeAsManager) > 0,
	}, nil
}
