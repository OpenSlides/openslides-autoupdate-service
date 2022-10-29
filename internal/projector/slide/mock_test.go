package slide_test

import (
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

func convertData(data map[string]string) map[datastore.Key][]byte {
	converted := make(map[datastore.Key][]byte, len(data))
	for k, v := range data {
		key := dskey.MustKey(k)
		converted[key] = []byte(v)
	}
	return converted
}
