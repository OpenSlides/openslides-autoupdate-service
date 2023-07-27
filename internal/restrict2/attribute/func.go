package attribute

import (
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

type UserAttributes struct {
	UserID            int
	GroupIDs          set.Set[int]
	OrgaLevel         perm.OrganizationManagementLevel
	IsCommitteManager bool
}

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
			return user.OrgaLevel == perm.OMLCanManageOrganization || user.OrgaLevel == perm.OMLCanManageUsers

		case perm.OMLCanManageUsers:
			return user.OrgaLevel == perm.OMLCanManageUsers

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

func FuncAllow() Func {
	return func(UserAttributes) bool {
		return true
	}
}

func FuncNotAllowed() Func {
	return func(UserAttributes) bool {
		return false
	}
}

func FuncIsCommitteeManager() Func {
	return func(user UserAttributes) bool {
		return user.IsCommitteManager
	}
}

func FuncLoggedIn() Func {
	return func(user UserAttributes) bool {
		return user.UserID > 0
	}
}
