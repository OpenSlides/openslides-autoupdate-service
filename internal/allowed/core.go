package allowed

import (
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

// IsAllowedParams does ...
type IsAllowedParams struct {
	UserID       int
	Data         definitions.FqfieldData
	DataProvider dataprovider.DataProvider
}

// IsAllowed does ...
type IsAllowed = func(params *IsAllowedParams) (map[string]interface{}, error)
