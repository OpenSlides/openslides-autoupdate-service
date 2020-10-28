package allowed

import (
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

type IsAllowedParams struct {
	UserId       int
	Data         definitions.FqfieldData
	DataProvider dataprovider.DataProvider
}

type IsAllowed = func(params *IsAllowedParams) (bool, map[string]interface{}, error)
