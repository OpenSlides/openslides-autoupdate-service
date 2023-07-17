package perm

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// GroupByPerm returns a map from permission to all groups, that have it.
// Includes groups that derivit that perm.
func GroupByPerm(ctx context.Context, ds *dsfetch.Fetch, meetingID int) (map[TPermission]set.Set[int], error) {
	groupIDs, err := ds.Meeting_GroupIDs(meetingID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch group ids: %w", err)
	}

	adminGroup, adminGroupExists, err := ds.Meeting_AdminGroupID(meetingID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch admin group id: %w", err)
	}

	var adminGroupList []int
	if adminGroupExists {
		adminGroupList = []int{adminGroup}
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
	permissionMap := make(map[TPermission]set.Set[int], len(derivatePerms))
	for groupID, permissionStrList := range groupToPermission {
		for _, permissionStr := range *permissionStrList {
			tperm := TPermission(permissionStr)

			if permissionMap[tperm].IsNotInitialized() {
				permissionMap[tperm] = set.New(adminGroupList...)
			}

			permissionMap[tperm].Add(groupID)
			for _, derived := range derivatePerms[tperm] {
				if permissionMap[derived].IsNotInitialized() {
					permissionMap[derived] = set.New(adminGroupList...)
				}

				permissionMap[derived].Add(groupID)
			}
		}
	}

	return permissionMap, nil
}
