package perm

// TPermission is a type of all valid permission strings.
type TPermission string

// OrganizationManagementLevel are possible values from user/organization_management_level
type OrganizationManagementLevel byte

// Organization management levels
const (
	OMLNone OrganizationManagementLevel = iota
	OMLSuperadmin
	OMLCanManageOrganization
	OMLCanManageUsers
)

// OrganizationManagementFromString returns the OrganizationManagementLevel type
// from a string.
func OrganizationManagementFromString(s string) OrganizationManagementLevel {
	switch s {
	case "superadmin":
		return OMLSuperadmin
	case "can_manage_organization":
		return OMLCanManageOrganization
	case "can_manage_users":
		return OMLCanManageUsers
	default:
		return OMLNone
	}
}
