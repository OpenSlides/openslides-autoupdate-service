package perm

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// GroupByPerm returns a list of groupIDs from permission. Includes groups that
// derivit that perm.
func GroupByPerm(ctx context.Context, ds *dsfetch.Fetch, meetingID int) (map[TPermission][]int, error) {
	var groupIDs []int
	ds.Meeting_GroupIDs(meetingID).Lazy(&groupIDs)

	var adminGroup int
	ds.Meeting_AdminGroupID(meetingID).Lazy(&adminGroup)

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting groupIDs and admin group of meeting %d: %w", meetingID, err)
	}

	// Fetch all permissions
	groupToPermission := make(map[int]*[]string)
	for _, groupID := range groupIDs {
		var permissions []string
		ds.Group_Permissions(groupID).Lazy(&permissions)
		groupToPermission[groupID] = &permissions
	}

	if err := ds.Execute(ctx); err != nil {
		return nil, fmt.Errorf("get all permissions: %w", err)
	}

	// Map permission to group ids
	permissionMap := make(map[TPermission][]int, len(derivatePerms))

	if adminGroup > 0 {
		for perm := range derivatePerms {
			permissionMap[perm] = []int{adminGroup}
		}
	}

	for groupID, permissionStrList := range groupToPermission {
		for _, permissionStr := range *permissionStrList {
			tperm := TPermission(permissionStr)

			permissionMap[tperm] = append(permissionMap[tperm], groupID)
			for _, derived := range derivatePerms[tperm] {
				permissionMap[derived] = append(permissionMap[derived], groupID)
			}
		}
	}

	return permissionMap, nil
}

type groupForMeeting struct {
	forMeetingID map[int]map[TPermission][]int
}

func newGroupForMeeting() *groupForMeeting {
	return &groupForMeeting{forMeetingID: make(map[int]map[TPermission][]int)}
}

func (p *groupForMeeting) MeetingGroupMap(ctx context.Context, fetcher *dsfetch.Fetch, meetingID int) (map[TPermission][]int, error) {
	perms, ok := p.forMeetingID[meetingID]
	if ok {
		return perms, nil
	}

	perms, err := GroupByPerm(ctx, fetcher, meetingID)
	if err != nil {
		return nil, fmt.Errorf("group by perm: %w", err)
	}
	p.forMeetingID[meetingID] = perms
	return perms, nil
}

func GroupMapFromContext(ctx context.Context, ds *dsfetch.Fetch, meetingID int) (map[TPermission][]int, error) {
	v := ctx.Value(groupCacheKey)
	if v == nil {
		return nil, fmt.Errorf("context does not contain a groupForMeeting cache. Make sure to create the context with 'ContextWithGroupCache'")
	}

	meetingPermission, ok := v.(*groupForMeeting)
	if !ok {
		return nil, fmt.Errorf("meeting permission has wrong type: %T", v)
	}

	return meetingPermission.MeetingGroupMap(ctx, ds, meetingID)
}

// ContextWithGroupMap creates a context with the group cache.
func ContextWithGroupMap(ctx context.Context) context.Context {
	return context.WithValue(ctx, groupCacheKey, newGroupForMeeting())
}
