package datastore_test

import (
	"context"
	"encoding/json"
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

	object, keys, err := datastore.Object(context.Background(), ds, "testmodel/1", []string{"id", "text", "friend_ids"})
	require.NoError(t, err, "Get returned unexpected error")

	assert.Equal(t, json.RawMessage([]byte("1")), object["id"])
	assert.Equal(t, json.RawMessage([]byte(`"my text"`)), object["text"])
	assert.Equal(t, json.RawMessage([]byte("[1,2,3]")), object["friend_ids"])
	assert.ElementsMatch(t, []string{"testmodel/1/id", "testmodel/1/text", "testmodel/1/friend_ids"}, keys)
}

func TestObjectFieldDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)
	ds := dsmock.NewMockDatastore(closed, map[string]string{})

	object, keys, err := datastore.Object(context.Background(), ds, "testmodel/1", []string{"id"})
	require.NoError(t, err, "Get returned unexpected error")

	require.Equal(t, 1, len(object))
	require.Nil(t, object["id"])
	assert.ElementsMatch(t, []string{"testmodel/1/id"}, keys)
}
