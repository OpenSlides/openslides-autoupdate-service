package attribute

import "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"

// Func tells for user attributes, if the user can see the object.
type Func func(user UserAttributes) bool

// FuncAnd combines Funcs. Returns true if all functions return true.
func FuncAnd(fn ...Func) Func {
	return func(user UserAttributes) bool {
		for _, f := range fn {
			if !f(user) {
				return false
			}
		}
		return true
	}
}

// FuncOr combines Funcs. Returns true if one functions return true.
func FuncOr(fn ...Func) Func {
	return func(user UserAttributes) bool {
		for _, f := range fn {
			if f(user) {
				return true
			}
		}
		return false
	}
}

// FuncOrgaLevel returns true, if the user has the organization level or a
// higher one.
func FuncOrgaLevel(oml perm.OrganizationManagementLevel) Func {
	return func(user UserAttributes) bool {
		switch user.OrgaLevel {
		case perm.OMLSuperadmin:
			return true

		case perm.OMLCanManageOrganization:
			return oml == perm.OMLCanManageOrganization || oml == perm.OMLCanManageUsers

		case perm.OMLCanManageUsers:
			return oml == perm.OMLCanManageUsers

		default:
			return false
		}
	}
}

// FuncInGroup returns true if the user is in one of the groups.
func FuncInGroup(groupIDs []int) Func {
	return func(user UserAttributes) bool {
		for _, id := range groupIDs {
			if user.GroupIDs.Has(id) {
				return true
			}
		}

		return false
	}
}

// FuncUserIDs returns true, if the user has one of the user ids.
func FuncUserIDs(userIDs []int) Func {
	return func(user UserAttributes) bool {
		for _, needUserID := range userIDs {
			if user.UserID == needUserID {
				return true
			}
		}
		return false
	}
}

// FuncAllowed returns true.
func FuncAllowed(UserAttributes) bool {
	return true
}

// FuncNotAllowed returns false.
func FuncNotAllowed(UserAttributes) bool {
	return false
}

// FuncIsCommitteeManager returns true if the user is a committee manager in one
// committee.
func FuncIsCommitteeManager(user UserAttributes) bool {
	return user.IsCommitteManager
}

// FuncLoggedIn returns true for logged in users.
func FuncLoggedIn(user UserAttributes) bool {
	return user.UserID > 0
}
