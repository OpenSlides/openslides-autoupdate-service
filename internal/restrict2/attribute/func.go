package attribute

import "github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"

type UserAttributes struct {
	UserID       int
	MeetingPerms map[int]*perm.Permission
	OrgaLevel    perm.OrganizationManagementLevel
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
		switch oml {
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

func FuncPerm(meetingID int, p perm.TPermission) Func {
	return func(user UserAttributes) bool {
		perms, ok := user.MeetingPerms[meetingID]
		if !ok {
			return false
		}

		return perms.Has(p)
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
