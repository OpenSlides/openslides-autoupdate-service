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
func (ps PermissionService) IsAllowed(ctx context.Context, name string, userID int, dataList []definitions.FqfieldData) ([]definitions.Addition, error, int) {
	var handler func(*allowed.IsAllowedParams) (definitions.Addition, error)
	var ok bool
	if handler, ok = Queries[name]; !ok {
		return nil, clientError{fmt.Sprintf("no such query: \"%s\"", name)}, -1
	}

	dp := dataprovider.NewDataProvider(ctx, ps.externalDataprovider)

	additions := make([]definitions.Addition, len(dataList))
	for i, data := range dataList {
		params := &allowed.IsAllowedParams{UserID: userID, Data: data, DataProvider: dp}
		addition, err := handler(params)

		if err != nil {
			err = fmt.Errorf("%s: %w", name, err)
			return nil, err, i
		}

		additions[i] = addition
	}

	return additions, nil, -1
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
