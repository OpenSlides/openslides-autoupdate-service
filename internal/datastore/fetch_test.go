package datastore_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/require"
)

func TestGetObject(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := test.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id":         "1",
		"testmodel/1/text":       `"my text"`,
		"testmodel/1/friend_ids": "[1,2,3]",
	})

	var testModel struct {
		ID      int    `json:"id"`
		Text    string `json:"text"`
		Friends []int  `json:"friend_ids"`
	}
	if err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
		t.Fatalf("dataGetObject returned unexpected error: %v", err)
	}

	if testModel.ID != 1 {
		t.Errorf("testModel.ID == %d, expected 1", testModel.ID)
	}
	if testModel.Text != "my text" {
		t.Errorf("testModel.Text == `%s`, expected `my text`", testModel.Text)
	}
	if len(testModel.Friends) != 3 || testModel.Friends[0] != 1 || testModel.Friends[1] != 2 || testModel.Friends[2] != 3 {
		t.Errorf("testModel.Friends == `%v`, expected `[1 2 3]`", testModel.Friends)
	}
}

func TestGetObjectOtherFields(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := test.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id": "1",
	})

	var testModel struct {
		ID    int `json:"id"`
		Other string
	}
	if err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
		t.Fatalf("dataGetObject returned unexpected error: %v", err)
	}

	if testModel.ID != 1 {
		t.Errorf("testModel.ID == %d, expected 1", testModel.ID)
	}
	if testModel.Other != "" {
		t.Errorf("testModel.Text == `%s`, expected ``", testModel.Other)
	}
}

func TestGetObjectOptions(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := test.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id": "1",
	})

	var testModel struct {
		ID int `json:"id,omitempty"`
	}
	if err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
		t.Fatalf("dataGetObject returned unexpected error: %v", err)
	}

	if testModel.ID != 1 {
		t.Errorf("testModel.ID == %d, expected 1", testModel.ID)
	}
}

func TestGetObjectFieldDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := test.NewMockDatastore(closed, map[string]string{})

	var testModel struct {
		ID int `json:"id"`
	}
	if err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
		t.Fatalf("dataGetObject returned unexpected error: %v", err)
	}

	if testModel.ID != 0 {
		t.Errorf("testModel.ID == %d, expected 0", testModel.ID)
	}
}

func TestObjectKeys(t *testing.T) {
	var testModel struct {
		ID      int    `json:"id"`
		Text    string `json:"text"`
		Friends []int  `json:"friend_ids"`
	}

	keys := datastore.ObjectKeys("testmodel/1", &testModel)

	require.ElementsMatch(t, keys, []string{"testmodel/1/id", "testmodel/1/text", "testmodel/1/friend_ids"})
}
