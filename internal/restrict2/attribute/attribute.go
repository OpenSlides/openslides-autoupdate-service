package attribute

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
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

	// GroupIDs are groups, that can see the field. Groups are meeting specific.
	GroupIDs set.Set[int]

	// UserIDs are list from users that can see the field but do not have the
	// globalPermission or are not in the groups.
	UserIDs set.Set[int]

	// HotKeys are all the keys that where needed to calculate the Attribute
	HotKeys set.Set[dskey.Key]
}

type GlobalPermission byte

const (
	GlobalAll GlobalPermission = iota
	GlobalSuperadmin
	GlobalCanManageOrganization
	GlobalCanManageUsers
	GlobalLoggedIn
	GlobalNobody
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
