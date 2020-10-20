package allowed

import (
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

type IsAllowedContext struct {
	UserId       int
	Data         definitions.FqfieldData
	DataProvider dataprovider.DataProvider
}

type IsAllowed = func(ctx *IsAllowedContext) (bool, map[string]interface{}, error)
