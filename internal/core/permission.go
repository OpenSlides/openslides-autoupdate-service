package core

import (
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// PermissionService impelements the permission.Permission interface.
type PermissionService struct {
	dataprovider dataprovider.DataProvider
}

// NewPermissionService returns a new permission service.
func NewPermissionService(externalDataprovider dataprovider.ExternalDataProvider) *PermissionService {
	dp := dataprovider.NewDataProvider(externalDataprovider)
	return &PermissionService{dp}
}

// IsAllowed tells, if something is allowed.
func (permissionService PermissionService) IsAllowed(name string, userID int, data definitions.FqfieldData) (bool, map[string]interface{}, error) {
	if val, ok := Queries[name]; ok {
		context := &allowed.IsAllowedParams{UserID: userID, Data: data, DataProvider: permissionService.dataprovider}
		return val(context)
	}

	return false, nil, fmt.Errorf("no such query: \"%s\"", name)
}

// RestrictFQIDs does currently nothing.
func (permissionService PermissionService) RestrictFQIDs(userID int, fqids []definitions.Fqid) (map[definitions.Fqid]bool, error) {
	r := make(map[definitions.Fqid]bool, len(fqids))
	for _, v := range fqids {
		r[v] = true
	}
	return r, nil
}

// RestrictFQFields does currently nothing.
func (permissionService PermissionService) RestrictFQFields(userID int, fqfields []definitions.Fqfield) (map[definitions.Fqfield]bool, error) {
	return permissionService.RestrictFQIDs(userID, fqfields)
}

// AdditionalUpdate does ...
func (permissionService PermissionService) AdditionalUpdate(updated definitions.FqfieldData) ([]definitions.Id, error) {
	return []definitions.Id{}, nil
}
