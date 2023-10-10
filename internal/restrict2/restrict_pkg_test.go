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

func TestRestrictModeForAll(t *testing.T) {
	ctx := collection.ContextWithRestrictCache(context.Background())

	for collectionField := range restrictionModes {
		collection, field, found := strings.Cut(collectionField, "/")
		if !found {
			t.Fatalf("invalid field %s, expected one /", collectionField)
		}

		fieldMode, err := restrictModeName(collection, field)
		if err != nil {
			t.Fatalf("building field mode: %v", err)
		}

		if _, err := restrictModefunc(ctx, collection, fieldMode); err != nil {
			t.Errorf("restrictMode(%s, %s) returned: %v", collection, field, err)
		}
	}
}

// restrictModeName returns the restriction mode for a collection and field.
//
// This is a string like "A" or "B" or any other name of a restriction mode.
func restrictModeName(collection, field string) (string, error) {
	fieldMode, ok := restrictionModes[collection+"/"+field]
	if !ok {
		// TODO LAST ERROR
		return "", fmt.Errorf("fqfield %q is unknown, maybe run go generate ./... to fetch all fields from the models.yml", collection+"/"+field)
	}
	return fieldMode, nil
}

// restrictModefunc returns the field restricter function to use.
func restrictModefunc(ctx context.Context, collectionName, fieldMode string) (collection.FieldRestricter, error) {
	restricter := collection.FromName(ctx, collectionName)
	if _, ok := restricter.(collection.Unknown); ok {
		// TODO LAST ERROR
		return nil, fmt.Errorf("collection %q is not implemented, maybe run go generate ./... to fetch all fields from the models.yml", collectionName)
	}

	modefunc := restricter.Modes(fieldMode)
	if modefunc == nil {
		// TODO LAST ERROR
		return nil, fmt.Errorf("mode %q of models %q is not implemented", fieldMode, collectionName)
	}

	return modefunc, nil
}
