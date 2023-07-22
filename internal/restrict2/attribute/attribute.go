package attribute

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

// Attribute is are flags that each field has. A user is allowed to see a
// field, if he has one of the attribute-fields.
type Attribute struct {
	// GlobalPermission is like the orga permission:
	//
	// 0: All user can see this
	// 1: Only superadmin can see this
	// 2: Global managers can see this
	// 3: Global user managers can see this
	// 4: Logged in users can see this
	// 5: Nobody can see this (not even the superadmin)
	GlobalPermission GlobalPermission

	// Permissions is a list of lists of permissions.
	//
	// The outer list meens AND and the inner list meens OR
	//
	// [[perm1, perm2], [perm3, perm4]]
	// meens ((perm1 or perm2) and (perm3 or perm4))
	//
	// TODO: Make TPermission an int or byte so it takes less memory (or is a string a pointer anyway?)
	MeetingID   int
	Permissions [][]perm.TPermission

	// UserIDs are list from users that can see the field but do not have the
	// globalPermission or are not in the groups.
	UserIDs []int
}

type GlobalPermission byte

const (
	GlobalNobody GlobalPermission = iota
	GlobalSuperadmin
	GlobalCanManageOrganization
	GlobalCanManageUsers
	GlobalLoggedIn
	GlobalAll
)

func GlobalFromPerm(p perm.OrganizationManagementLevel) GlobalPermission {
	switch p {
	case perm.OMLSuperadmin:
		return GlobalSuperadmin
	case perm.OMLCanManageOrganization:
		return GlobalCanManageOrganization
	case perm.OMLCanManageUsers:
		return GlobalCanManageUsers
	default:
		return GlobalNobody
	}
}
