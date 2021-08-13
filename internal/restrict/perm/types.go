package perm

import (
	"fmt"
	"strconv"
	"strings"
)

// FQField contains all parts of a fqfield.
type FQField struct {
	Collection string
	ID         int
	Field      string
}

// ParseFQField creates an FQField object from a fqfield string.
func ParseFQField(fqfield string) (FQField, error) {
	parts := strings.Split(fqfield, "/")
	if len(parts) != 3 {
		return FQField{}, fmt.Errorf("invalid fqfield '%s'", fqfield)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return FQField{}, fmt.Errorf("invalid fqfield '%s': %w", fqfield, err)
	}

	return FQField{
		Collection: parts[0],
		ID:         id,
		Field:      parts[2],
	}, nil
}

func (fqfield FQField) String() string {
	return fmt.Sprintf("%s/%d/%s", fqfield.Collection, fqfield.ID, fqfield.Field)
}

// FQID returns the fqid representation of the fqfiedl.
func (fqfield FQField) FQID() string {
	return fmt.Sprintf("%s/%d", fqfield.Collection, fqfield.ID)
}

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
