package datastore_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/datastore"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestDatastore(t *testing.T) {
	dbServer := tests.NewDatastoreServer()
	defer dbServer.TS.Close()

	db := &datastore.Datastore{Addr: dbServer.TS.URL}

	result, err := db.Get(context.Background(), "role/1/name", "unknown/2/field")
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("Got %d values, expected 2", len(result))
	}

	if string(result[0]) != `"Superadmin role"` {
		t.Errorf("Got first value `%s`, expected `\"Superadmin role\"`", result[0])
	}

	if result[1] != nil {
		t.Errorf("Got second value `%s`, expected nil", result[1])
	}
}

func TestDataStoreNull(t *testing.T) {
	dbServer := tests.NewDatastoreServer()
	defer dbServer.TS.Close()

	db := &datastore.Datastore{Addr: dbServer.TS.URL}

	result, err := db.Get(context.Background(), "group/1/superadmin_group_for_meeting_id")
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("Got %d values, expected 1", len(result))
	}

	if result[0] != nil {
		t.Errorf("Got first value `%s`, expected nil", result[0])
	}
}
