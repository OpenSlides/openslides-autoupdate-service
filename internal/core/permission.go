package core

import (
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

type permissionService struct {
	dataprovider dataprovider.DataProvider
}

func NewPermissionService(externalDataprovider definitions.ExternalDataProvider) definitions.Permission {
	dp := dataprovider.NewDataProvider(externalDataprovider)
	return &permissionService{dp}
}

func (permissionService permissionService) IsAllowed(name string, userId int, data definitions.FqfieldData) (bool, map[string]interface{}, error) {
	if val, ok := Queries[name]; ok {
		context := &allowed.IsAllowedContext{UserId: userId, Data: data, DataProvider: permissionService.dataprovider}
		return val(context)
	} else {
		return false, nil, fmt.Errorf("no such query: \"%s\"", name)
	}
}

func (permissionService permissionService) RestrictFqIds(fqids map[definitions.Fqid]bool, userId int) (map[definitions.Fqid]bool, error) {
	return fqids, nil
}

func (permissionService permissionService) RestrictFqields(fqfields map[definitions.Fqfield]bool, userId int) (map[definitions.Fqfield]bool, error) {
	return fqfields, nil
}

func (permissionService permissionService) AdditionalUpdate(updated definitions.FqfieldData) ([]definitions.Id, error) {
	return []definitions.Id{}, nil
}
