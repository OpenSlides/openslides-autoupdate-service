package attribute

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

type Func func(user UserAttributes) bool

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

func FuncGlobalLevel(oml perm.OrganizationManagementLevel) Func {
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

func FuncAllow(UserAttributes) bool {
	return true
}

func FuncNotAllowed(UserAttributes) bool {
	return false
}

func FuncIsCommitteeManager(user UserAttributes) bool {
	return user.IsCommitteManager
}

func FuncLoggedIn(user UserAttributes) bool {
	return user.UserID > 0
}
