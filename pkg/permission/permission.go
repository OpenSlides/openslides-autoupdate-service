package permission

import (
	"github.com/OpenSlides/openslides-permission-service/internal/core"
	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

func New(dataprovider definitions.ExternalDataProvider) definitions.Permission {
	return core.NewPermissionService(dataprovider)
}
