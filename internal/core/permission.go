package core

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
)

// PermissionService impelements the permission.Permission interface.
type PermissionService struct {
	externalDataprovider dataprovider.ExternalDataProvider
}

// NewPermissionService returns a new permission service.
func NewPermissionService(externalDataprovider dataprovider.ExternalDataProvider) *PermissionService {
	return &PermissionService{externalDataprovider}
}

// IsAllowed tells, if something is allowed.
func (permissionService PermissionService) IsAllowed(ctx context.Context, name string, userID int, data definitions.FqfieldData) (bool, map[string]interface{}, error) {
	dp := dataprovider.NewDataProvider(ctx, permissionService.externalDataprovider)

	if val, ok := Queries[name]; ok {
		context := &allowed.IsAllowedParams{UserID: userID, Data: data, DataProvider: dp}
		return val(context)
	}

	return false, nil, fmt.Errorf("no such query: \"%s\"", name)
}

// RestrictFQIDs does currently nothing.
func (permissionService PermissionService) RestrictFQIDs(ctx context.Context, userID int, fqids []definitions.Fqid) (map[definitions.Fqid]bool, error) {
	r := make(map[definitions.Fqid]bool, len(fqids))
	for _, v := range fqids {
		r[v] = true
	}
	return r, nil
}

// RestrictFQFields does currently nothing.
func (permissionService PermissionService) RestrictFQFields(ctx context.Context, userID int, fqfields []definitions.Fqfield) (map[definitions.Fqfield]bool, error) {
	return permissionService.RestrictFQIDs(ctx, userID, fqfields)
}

// AdditionalUpdate does ...
func (permissionService PermissionService) AdditionalUpdate(ctx context.Context, updated definitions.FqfieldData) ([]definitions.Id, error) {
	return []definitions.Id{}, nil
}
