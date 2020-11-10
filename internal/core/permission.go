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
func NewPermissionService(edp dataprovider.ExternalDataProvider) *PermissionService {
	return &PermissionService{edp}
}

// IsAllowed tells, if something is allowed.
func (ps PermissionService) IsAllowed(ctx context.Context, name string, userID int, data definitions.FqfieldData) (bool, map[string]interface{}, error) {
	var handler func(*allowed.IsAllowedParams) (map[string]interface{}, error)
	var ok bool
	if handler, ok = Queries[name]; !ok {
		return false, nil, clientError{fmt.Sprintf("no such query: \"%s\"", name)}
	}

	dp := dataprovider.NewDataProvider(ctx, ps.externalDataprovider)
	params := &allowed.IsAllowedParams{UserID: userID, Data: data, DataProvider: dp}
	addition, err := handler(params)

	isAllowed := err == nil

	// Wrap the query name around the error
	if err != nil {
		err = fmt.Errorf("%s: %w", name, err)
	}

	return isAllowed, addition, err
}

// RestrictFQIDs does currently nothing.
func (ps PermissionService) RestrictFQIDs(ctx context.Context, userID int, fqids []definitions.Fqid) (map[definitions.Fqid]bool, error) {
	r := make(map[definitions.Fqid]bool, len(fqids))
	for _, v := range fqids {
		r[v] = true
	}
	return r, nil
}

// RestrictFQFields does currently nothing.
func (ps PermissionService) RestrictFQFields(ctx context.Context, userID int, fqfields []definitions.Fqfield) (map[definitions.Fqfield]bool, error) {
	return ps.RestrictFQIDs(ctx, userID, fqfields)
}

// AdditionalUpdate does ...
func (ps PermissionService) AdditionalUpdate(ctx context.Context, updated definitions.FqfieldData) ([]definitions.Id, error) {
	return []definitions.Id{}, nil
}
