package projector_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
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
	if fields["projection/1/content"] != nil {
		t.Errorf("Content was calculated, should be nil, got: %q", fields["projection/1/content"])
	}
}

func TestProjectionFromContentObject(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/content_object_id": `"test_model/1"`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"test_model"` + "\n"
	assert.JSONEq(t, expect, string(fields["projection/1/content"]))
}

func TestProjectionFromType(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/content_object_id": `"meeting/1"`,
		"projection/1/type":              `"test1"`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `"abc"` + "\n"
	assert.JSONEq(t, expect, string(fields["projection/1/content"]))
}

func TestProjectionUpdateProjection(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/content_object_id": `"meeting/1"`,
		"projection/1/type":              `"test1"`,
	})
	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string][]byte) error {
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
	assert.JSONEq(t, expect, string(fields["projection/1/content"]))
}

func TestProjectionUpdateProjectionMetaData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/type":              `"projection"`,
		"projection/1/content_object_id": `"meeting/1"`,
	})
	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[string][]byte) error {
		close(done)
		return nil
	})

	ds.Send(map[string]string{
		"projection/1/stable": "true",
	})
	<-done

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"id": 1, "content_object_id": "meeting/1", "meeting_id":0, "type":"projection", "options": null}` + "\n"
	assert.JSONEq(t, expect, string(fields["projection/1/content"]))
}

func TestProjectionWithOptionsData(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/content_object_id": `"meeting/6"`,
		"projection/1/type":              `"projection"`,
		"projection/1/meeting_id":        `1`,
		"projection/1/options":           `{"only_main_items": true}`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"id": 1, "content_object_id": "meeting/6", "type":"projection", "meeting_id": 1, "options": {"only_main_items": true}}` + "\n"
	assert.JSONEq(t, expect, string(fields["projection/1/content"]))
}

func TestProjectionUpdateSlide(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/content_object_id": `"meeting/6"`,
		"projection/1/type":              `"test_model"`,
	})
	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), "projection/1/content")
	require.NoError(t, err, "Get returned unexpected error")

	// Register a listener that tells, when cache is updated.
	done := make(chan struct{})
	ds.RegisterChangeListener(func(data map[string][]byte) error {
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
	assert.JSONEq(t, expect, string(fields["projection/1/content"]))
}

func TestProjectionUpdateOtherKey(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/content_object_id": `"meeting/1"`,
		"projection/1/type":              `"test_model"`,
	})
	projector.Register(ds, testSlides())

	// Call once to add field to cache.
	ds.Get(context.Background(), "projection/1/content")

	// Register a listener that tells, when cache is updated.
	done := make(chan struct{})
	ds.RegisterChangeListener(func(data map[string][]byte) error {
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
	assert.JSONEq(t, expect, string(fields["projection/1/content"]))
}

func TestProjectionTypeDoesNotExist(t *testing.T) {
	closed := make(chan struct{})
	defer close(closed)

	ds := dsmock.NewMockDatastore(closed, map[string]string{
		"projection/1/id":                "1",
		"projection/1/content_object_id": `"meeting/1"`,
		"projection/1/type":              `"unexistingTestSlide"`,
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), "projection/1/content")
	if err != nil {
		t.Fatalf("Get returned unexpected error: %v", err)
	}

	var content struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal(fields["projection/1/content"], &content); err != nil {
		t.Fatalf("Can not unmarshal field projection/1/content `%s`: %v", fields["projection/1/content"], err)
	}

	if content.Error == "" {
		t.Errorf("Field has not error")
	}
}

func testSlides() *projector.SlideStore {
	s := new(projector.SlideStore)
	s.RegisterSliderFunc("test1", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		return []byte(`"abc"`), nil
	})

	s.RegisterSliderFunc("test_model", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		var field json.RawMessage
		fetch.Fetch(ctx, &field, "test_model/1/field")
		if field == nil {
			return []byte(`"test_model"`), nil
		}
		return []byte(fmt.Sprintf(`"calculated with %s"`, string(field[1:len(field)-1]))), nil
	})

	s.RegisterSliderFunc("projection", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		bs, err := json.Marshal(p7on)
		return bs, err
	})
	return s
}
