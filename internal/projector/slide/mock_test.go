package slide_test

import (
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

func convertData(data map[string]string) map[datastore.Key][]byte {
	converted := make(map[datastore.Key][]byte, len(data))
	for k, v := range data {
		key, err := datastore.KeyFromString(k)
		if err != nil {
			panic(fmt.Errorf("invalid key: %s", k))
		}
		converted[key] = []byte(v)
	}
	return converted
}
