package datastore_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObject(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id":         "1",
		"testmodel/1/text":       `"my text"`,
		"testmodel/1/friend_ids": "[1,2,3]",
	})

	var testModel struct {
		ID      int    `json:"id"`
		Text    string `json:"text"`
		Friends []int  `json:"friend_ids"`
	}
	keys, err := datastore.Object(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 1, testModel.ID, "testModel.ID")
	assert.Equal(t, "my text", testModel.Text, "testModel.Text")
	assert.Equal(t, []int{1, 2, 3}, testModel.Friends, "testModel.Friends")
	assert.ElementsMatch(t, []string{"testmodel/1/id", "testmodel/1/text", "testmodel/1/friend_ids"}, keys)
}

func TestObjectOtherFields(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id": "1",
	})

	var testModel struct {
		ID    int `json:"id"`
		Other string
	}
	keys, err := datastore.Object(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 1, testModel.ID, "testModel.ID")
	assert.Equal(t, "", testModel.Other, "testModel.Other")
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, keys)
}

func TestObjectOptions(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/id": "1",
	})

	var testModel struct {
		ID int `json:"id,omitempty"`
	}
	keys, err := datastore.Object(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 1, testModel.ID, "testModel.ID")
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, keys)
}

func TestObjectFieldDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{})

	var testModel struct {
		ID int `json:"id"`
	}
	keys, err := datastore.Object(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Get returned unexpected error")
	assert.Equal(t, 0, testModel.ID, "testModel.ID")
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, keys)
}

func TestObjectTemplate(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/field_$":       `["fast","long"]`,
		"testmodel/1/field_$fast":   `"value1"`,
		"testmodel/1/field_$long":   `"value2"`,
		"testmodel/1/number_$_ids":  `["1","2"]`,
		"testmodel/1/number_$1_ids": `[1,2,3]`,
		"testmodel/1/number_$2_ids": `[4,5,6]`,
	})

	var testModel struct {
		Fields  map[string]string `json:"field_$"`
		Numbers map[int][]int     `json:"number_$_ids"`
	}
	keys, err := datastore.Object(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Object returned unexpected error")
	assert.Equal(t, testModel.Fields, map[string]string{
		"fast": "value1",
		"long": "value2",
	})

	assert.Equal(t, testModel.Numbers, map[int][]int{
		1: {1, 2, 3},
		2: {4, 5, 6},
	})

	expectKeys := []string{
		"testmodel/1/field_$",
		"testmodel/1/field_$fast",
		"testmodel/1/field_$long",
		"testmodel/1/number_$_ids",
		"testmodel/1/number_$1_ids",
		"testmodel/1/number_$2_ids",
	}
	assert.ElementsMatch(t, expectKeys, keys)
}

func TestObjectTemplateFieldDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/field_$":     `["fast","long"]`,
		"testmodel/1/field_$fast": `"value1"`,
	})

	var testModel struct {
		Fields map[string]string `json:"field_$"`
	}
	keys, err := datastore.Object(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Object returned unexpected error")
	assert.Equal(t, testModel.Fields, map[string]string{
		"fast": "value1",
	})

	expectKeys := []string{
		"testmodel/1/field_$",
		"testmodel/1/field_$fast",
		"testmodel/1/field_$long",
	}
	assert.ElementsMatch(t, expectKeys, keys)
}

func TestObjectTemplateDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"testmodel/1/field_$fast": `"value1"`,
	})

	var testModel struct {
		Fields map[string]string `json:"field_$"`
	}
	keys, err := datastore.Object(context.Background(), ds, "testmodel/1", &testModel)
	require.NoError(t, err, "Object returned unexpected error")
	assert.Nil(t, testModel.Fields)

	expectKeys := []string{
		"testmodel/1/field_$",
	}
	assert.ElementsMatch(t, expectKeys, keys)
}
