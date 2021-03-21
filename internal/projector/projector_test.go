package projector_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/openslides/openslides-autoupdate-service/internal/projector"
	"github.com/openslides/openslides-autoupdate-service/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectionDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := test.NewMockDatastore(closed, nil)
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	assert.Nil(t, fields[0], "Get content for nonexisting projection should not exist")
}

func TestProjectionFromContentObject(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := test.NewMockDatastore(closed, map[string]string{
		"projection/1/content_object_id": `"test_model/1"`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"test_model"` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func TestProjectionFromType(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := test.NewMockDatastore(closed, map[string]string{
		"projection/1/type": `"test1"`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"abc"` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func TestProjectionUpdateProjection(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := test.NewMockDatastore(closed, map[string]string{
		"projection/1/type": `"test1"`,
	})
	projector.Register(ds, testSlides())

	ds.Send(map[string]string{
		"projection/1/type":              "",
		"projection/1/content_object_id": `"test_model/1"`,
	})

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"test_model"` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func TestProjectionUpdateSlide(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := test.NewMockDatastore(closed, map[string]string{
		"projection/1/type": `"test_model"`,
	})
	projector.Register(ds, testSlides())

	// Call once to add field to cache.
	ds.Get(context.Background(), "projection/1/content")

	// Register a listener that tells, when cache is updated.
	done := make(chan struct{})
	ds.RegisterChangeListener(func(data map[string]json.RawMessage) error {
		close(done)
		return nil
	})

	ds.Send(map[string]string{
		"test_model/1/field": `"new value"`,
	})
	<-done

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"calculated with new value"` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func TestProjectionUpdateOtherKey(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := test.NewMockDatastore(closed, map[string]string{
		"projection/1/type": `"test_model"`,
	})
	projector.Register(ds, testSlides())

	// Call once to add field to cache.
	ds.Get(context.Background(), "projection/1/content")

	// Register a listener that tells, when cache is updated.
	done := make(chan struct{})
	ds.RegisterChangeListener(func(data map[string]json.RawMessage) error {
		close(done)
		return nil
	})

	ds.Send(map[string]string{
		"some_other/1/field": `"new value"`,
	})
	<-done

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"test_model"` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func testSlides() *projector.SlideStore {
	s := new(projector.SlideStore)
	s.AddFunc("test1", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"abc"`), nil, nil
	})
	s.AddFunc("test_model", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		field, err := ds.Get(ctx, "test_model/1/field")
		if field[0] == nil {
			return []byte(`"test_model"`), []string{"test_model/1/field"}, nil
		}
		return []byte(fmt.Sprintf(`"calculated with %s"`, string(field[0][1:len(field[0])-1]))), []string{"test_model/1/field"}, nil
	})
	return s
}
