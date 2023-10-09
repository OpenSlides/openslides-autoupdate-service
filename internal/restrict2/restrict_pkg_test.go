package restrict

import (
	"context"
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func TestRestricterWithID_0(t *testing.T) {
	ctx := collection.ContextWithRestrictCache(context.Background())
	fetcher := dsfetch.New(dsmock.Stub(nil))

	allModes := make(map[string]string)
	for fieldName, mode := range restrictionModes {
		collectionName, _, found := strings.Cut(fieldName, "/")
		if !found {
			t.Fatalf("invalid field %s, expected one /", fieldName)
		}

		allModes[collectionName] = mode
	}

	keys := make([]string, 0, len(allModes))
	for k := range allModes {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, collectionName := range keys {
		modeName := allModes[collectionName]
		fr := collection.FromName(ctx, collectionName).Modes(modeName)

		t.Run(fmt.Sprintf("%s/%s", collectionName, modeName), func(t *testing.T) {
			attr, err := fr(ctx, fetcher, []int{0, 0, 0})
			if err != nil {
				t.Errorf("field restricter: %v", err)
			}

			for i := 0; i < 3; i++ {
				if attr[i] != nil {
					t.Errorf("attr %d is not nil", i)
				}
			}
		})
	}
}
