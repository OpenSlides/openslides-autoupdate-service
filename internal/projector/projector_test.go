package projector_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectionDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, nil)
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	assert.Nil(t, fields[0], "Get content for nonexisting projection should not exist")
}

func TestProjectionFromContentObject(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
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

	ds := dsmock.NewMockDatastore(closed, map[string]string{
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

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/type": `"test1"`,
	})
	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string]json.RawMessage) error {
		close(done)
		return nil
	})

	ds.Send(map[string]string{
		"projection/1/type":              "",
		"projection/1/content_object_id": `"test_model/1"`,
	})
	<-done

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"test_model"` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func TestProjectionUpdateProjectionMetaData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/type":       `"projection"`,
		"projection/1/meeting_id": `1`,
	})
	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string]json.RawMessage) error {
		close(done)
		return nil
	})

	ds.Send(map[string]string{
		"projection/1/stable": "true",
	})
	<-done

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"id": 0, "content_object_id": "", "type":"projection", "meeting_id": 1, "options": {"only_main_items": false}}` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func TestProjectionWithOptionsData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/type":       `"projection"`,
		"projection/1/meeting_id": `1`,
		"projection/1/options":    `{"only_main_items": true, "unused": "not in ProjectionOptions type"}`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"id": 0, "content_object_id": "", "type":"projection", "meeting_id": 1, "options": {"only_main_items": true}}` + "\n"
	assert.JSONEq(t, expect, string(fields[0]))
}

func TestProjectionUpdateSlide(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/type": `"test_model"`,
	})
	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")

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

	ds := dsmock.NewMockDatastore(closed, map[string]string{
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

func TestProjectionTypeDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/type": `"unexistingTestSlide"`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	if err != nil {
		t.Fatalf("Get returned unexpected error: %v", err)
	}

	var content struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal(fields[0], &content); err != nil {
		t.Fatalf("Can not unmarshal field[0] `%s`: %v", fields[0], err)
	}

	if content.Error == "" {
		t.Errorf("Field has not error")
	}
}

func testSlides() *projector.SlideStore {
	s := new(projector.SlideStore)
	s.RegisterSlideFunc("test1", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"abc"`), nil, nil
	})
	s.RegisterSlideFunc("test_model", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		field, err := ds.Get(ctx, "test_model/1/field")
		if field[0] == nil {
			return []byte(`"test_model"`), []string{"test_model/1/field"}, nil
		}
		return []byte(fmt.Sprintf(`"calculated with %s"`, string(field[0][1:len(field[0])-1]))), []string{"test_model/1/field"}, nil
	})
	s.RegisterSlideFunc("projection", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		bs, err := json.Marshal(p7on)
		return bs, nil, err
	})
	return s
}
