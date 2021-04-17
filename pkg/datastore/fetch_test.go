package datastore_test

import (
	"context"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/openslides/openslides-autoupdate-service/pkg/datastore"
	"github.com/stretchr/testify/assert"
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
	keys, err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 1, testModel.ID, "testModel.ID")
	assert.Equal(t, "my text", testModel.Text, "testModel.Text")
	assert.Equal(t, []int{1, 2, 3}, testModel.Friends, "testModel.Friends")
	assert.ElementsMatch(t, []string{"testmodel/1/id", "testmodel/1/text", "testmodel/1/friend_ids"}, keys)
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
	keys, err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 1, testModel.ID, "testModel.ID")
	assert.Equal(t, "", testModel.Other, "testModel.Other")
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, keys)
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
	keys, err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 1, testModel.ID, "testModel.ID")
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, keys)
}

func TestGetObjectFieldDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := test.NewMockDatastore(closed, map[string]string{})

	var testModel struct {
		ID int `json:"id"`
	}
	keys, err := datastore.GetObject(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 0, testModel.ID, "testModel.ID")
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, keys)
}
