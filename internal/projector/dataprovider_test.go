package projector

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/test"
)

func TestDataGetObject(t *testing.T) {
	ds := test.NewMockDatastore(map[string]string{
		"testmodel/1/id":         "1",
		"testmodel/1/text":       `"my text"`,
		"testmodel/1/friend_ids": "[1,2,3]",
	})

	var testModel struct {
		ID      int    `json:"id"`
		Text    string `json:"text"`
		Friends []int  `json:"friend_ids"`
	}
	if err := dataGetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
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

func TestDataGetObjectOtherFields(t *testing.T) {
	ds := test.NewMockDatastore(map[string]string{
		"testmodel/1/id": "1",
	})

	var testModel struct {
		ID    int `json:"id"`
		Other string
	}
	if err := dataGetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
		t.Fatalf("dataGetObject returned unexpected error: %v", err)
	}

	if testModel.ID != 1 {
		t.Errorf("testModel.ID == %d, expected 1", testModel.ID)
	}
	if testModel.Other != "" {
		t.Errorf("testModel.Text == `%s`, expected ``", testModel.Other)
	}
}

func TestDataGetObjectOptions(t *testing.T) {
	ds := test.NewMockDatastore(map[string]string{
		"testmodel/1/id": "1",
	})

	var testModel struct {
		ID int `json:"id,omitempty"`
	}
	if err := dataGetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
		t.Fatalf("dataGetObject returned unexpected error: %v", err)
	}

	if testModel.ID != 1 {
		t.Errorf("testModel.ID == %d, expected 1", testModel.ID)
	}
}

func TestDataGetObjectFieldDoesNotExist(t *testing.T) {
	ds := test.NewMockDatastore(map[string]string{})

	var testModel struct {
		ID int `json:"id"`
	}
	if err := dataGetObject(context.Background(), ds, "testmodel/1", &testModel); err != nil {
		t.Fatalf("dataGetObject returned unexpected error: %v", err)
	}

	if testModel.ID != 0 {
		t.Errorf("testModel.ID == %d, expected 0", testModel.ID)
	}
}
