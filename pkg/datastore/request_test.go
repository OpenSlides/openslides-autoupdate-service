package datastore_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func TestRequestTwoWithErrors(t *testing.T) {
	ds := datastore.NewRequest(dsmock.Stub(dsmock.YAMLData(`---
	topic/1/title: foo
	`)))

	_, err := ds.Topic_Title(2).Value(context.Background())
	if err == nil {
		t.Fatalf("Title2 returned no error, expected DoesNotExist")
	}

	_, err = ds.Topic_Title(1).Value(context.Background())
	if err != nil {
		t.Errorf("Title1 returned unexpected error: %v", err)
	}
}
