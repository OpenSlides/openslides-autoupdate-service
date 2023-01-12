package perm

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// MeetingPermission is a cache for different Permission objects for each
// meeting.
//
// Can be used if fields from different meetings are checked.
type MeetingPermission struct {
	forMeetingID map[int]map[TPermission]set.Set[int]
}

// NewMeetingPermission initializes a new MeetingPermission.
func NewMeetingPermission() MeetingPermission {
	return MeetingPermission{
		forMeetingID: make(map[int]map[TPermission]set.Set[int]),
	}
}

// Meeting returns the permission object for the meeting.
func (p MeetingPermission) Meeting(ctx context.Context, ds *dsfetch.Fetch, meetingID int) (map[TPermission]set.Set[int], error) {
	perms, ok := p.forMeetingID[meetingID]
	if ok {
		return perms, nil
	}

	perms, err := GroupByPerm(ctx, ds, meetingID)
	if err != nil {
		return nil, fmt.Errorf("group by perm: %w", err)
	}

	return perms, nil
}
