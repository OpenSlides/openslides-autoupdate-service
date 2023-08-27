package restrict

import (
	"context"
	"strings"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

func TestRestrictModeForAll(t *testing.T) {
	for collectionField := range restrictionModes {
		collection, field, found := strings.Cut(collectionField, "/")
		if !found {
			t.Fatalf("invalid field %s, expected one /", collectionField)
		}

		fieldMode, err := restrictModeName(collection, field)
		if err != nil {
			t.Fatalf("building field mode: %v", err)
		}

		if _, err := restrictModefunc(context.Background(), collection, fieldMode); err != nil {
			t.Errorf("restrictMode(%s, %s) returned: %v", collection, field, err)
		}
	}
}

func TestCollectionOrderContainsAll(t *testing.T) {
	allCollection := set.New[string]()
	for collectionField := range restrictionModes {
		collection, _, found := strings.Cut(collectionField, "/")
		if !found {
			t.Fatalf("invalid field %s, expected one /", collectionField)
		}
		allCollection.Add(collection)
	}

	collectionOrderKeys := set.New[string]()
	for key := range collectionOrder {
		collection, _, found := strings.Cut(key, "/")
		if !found {
			collection = key
		}

		collectionOrderKeys.Add(collection)
	}

	if !set.Equal(allCollection, collectionOrderKeys) {
		t.Errorf("CollectionOrder is incorrect")
	}
}
