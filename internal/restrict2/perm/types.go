package perm

// TPermission is a type of all valid permission strings.
type TPermission string

// OrganizationManagementLevel are possible values from user/organization_management_level
type OrganizationManagementLevel string

// Organization management levels
const (
	OMLNone                  OrganizationManagementLevel = ""
	OMLSuperadmin            OrganizationManagementLevel = "superadmin"
	OMLCanManageOrganization OrganizationManagementLevel = "can_manage_organization"
	OMLCanManageUsers        OrganizationManagementLevel = "can_manage_users"
)
