package projector_test

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectionDoesNotExist(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds, bg := dsmock.NewMockDatastore(nil)
	go bg(shutdownCtx, oserror.Handle)

	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	if fields[dskey.MustKey("projection/1/content")] != nil {
		t.Errorf("Content was calculated, should be nil, got: %q", fields[dskey.MustKey("projection/1/content")])
	}
}

func TestProjectionFromContentObject(t *testing.T) {
	ds, _ := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"test_model/1"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"collection":"test_model","value":"test_model"}` + "\n"
	assert.JSONEq(t, expect, string(fields[dskey.MustKey("projection/1/content")]))
}

func TestProjectionFromContentObjectIfNotOnProjector(t *testing.T) {
	ds, _ := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"test_model/1"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`null`),
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	if got := fields[dskey.MustKey("projection/1/content")]; got != nil {
		t.Errorf("got %s, expected nil", got)
	}
}

func TestProjectionFromType(t *testing.T) {
	ds, _ := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"meeting/1"`),
		dskey.MustKey("projection/1/type"):                 []byte(`"test1"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"collection":"test1","value":"abc"}` + "\n"
	assert.JSONEq(t, expect, string(fields[dskey.MustKey("projection/1/content")]))
}

func TestProjectionUpdateProjection(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds, bg := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"meeting/1"`),
		dskey.MustKey("projection/1/type"):                 []byte(`"test1"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	go bg(shutdownCtx, oserror.Handle)

	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[dskey.Key][]byte) error {
		close(done)
		return nil
	})

	ds.Send(dsmock.YAMLData(`---
	projection/1/type: null
	projection/1/content_object_id: test_model/1
	`))
	<-done

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"collection":"test_model","value":"test_model"}` + "\n"
	assert.JSONEq(t, expect, string(fields[dskey.MustKey("projection/1/content")]))
}

func TestProjectionUpdateProjectionMetaData(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds, bg := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/type"):                 []byte(`"projection"`),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"meeting/1"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	go bg(shutdownCtx, oserror.Handle)

	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")

	done := make(chan struct{})
	ds.RegisterChangeListener(func(map[dskey.Key][]byte) error {
		close(done)
		return nil
	})

	ds.Send(dsmock.YAMLData("projection/1/stable: true"))
	<-done

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"collection":"projection","id": 1, "content_object_id": "meeting/1", "meeting_id":0, "type":"projection", "options": null,"current_projector_id":1}` + "\n"
	assert.JSONEq(t, expect, string(fields[dskey.MustKey("projection/1/content")]))
}

func TestProjectionWithOptionsData(t *testing.T) {
	ds, _ := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"meeting/6"`),
		dskey.MustKey("projection/1/type"):                 []byte(`"projection"`),
		dskey.MustKey("projection/1/meeting_id"):           []byte(`1`),
		dskey.MustKey("projection/1/options"):              []byte(`{"only_main_items": true}`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"collection":"projection","id": 1,"current_projector_id":1, "content_object_id": "meeting/6", "type":"projection", "meeting_id": 1, "options": {"only_main_items": true}}` + "\n"
	assert.JSONEq(t, expect, string(fields[dskey.MustKey("projection/1/content")]))
}

func TestProjectionUpdateSlide(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds, bg := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"meeting/6"`),
		dskey.MustKey("projection/1/type"):                 []byte(`"test_model"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	go bg(shutdownCtx, oserror.Handle)

	projector.Register(ds, testSlides())

	// Fetch data once to fill the test.
	_, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")

	// Register a listener that tells, when cache is updated.
	done := make(chan struct{})
	ds.RegisterChangeListener(func(data map[dskey.Key][]byte) error {
		close(done)
		return nil
	})

	ds.Send(dsmock.YAMLData("test_model/1/field: new value"))
	<-done

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"collection":"test_model","value":"calculated with new value"}` + "\n"
	assert.JSONEq(t, expect, string(fields[dskey.MustKey("projection/1/content")]))
}

func TestProjectionUpdateOtherKey(t *testing.T) {
	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds, bg := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"meeting/1"`),
		dskey.MustKey("projection/1/type"):                 []byte(`"test_model"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	go bg(shutdownCtx, oserror.Handle)

	projector.Register(ds, testSlides())

	// Call once to add field to cache.
	ds.Get(context.Background(), dskey.MustKey("projection/1/content"))

	// Register a listener that tells, when cache is updated.
	done := make(chan struct{})
	ds.RegisterChangeListener(func(data map[dskey.Key][]byte) error {
		close(done)
		return nil
	})

	ds.Send(dsmock.YAMLData("some_other/1/field: new value"))
	<-done

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	require.NoError(t, err, "Get returned unexpected error")
	expect := `{"collection":"test_model","value":"test_model"}` + "\n"
	assert.JSONEq(t, expect, string(fields[dskey.MustKey("projection/1/content")]))
}

func TestProjectionTypeDoesNotExist(t *testing.T) {
	ds, _ := dsmock.NewMockDatastore(map[dskey.Key][]byte{
		dskey.MustKey("projection/1/id"):                   []byte("1"),
		dskey.MustKey("projection/1/content_object_id"):    []byte(`"meeting/1"`),
		dskey.MustKey("projection/1/type"):                 []byte(`"unexistingTestSlide"`),
		dskey.MustKey("projection/1/current_projector_id"): []byte(`1`),
	})
	projector.Register(ds, testSlides())

	fields, err := ds.Get(context.Background(), dskey.MustKey("projection/1/content"))
	if err != nil {
		t.Fatalf("Get returned unexpected error: %v", err)
	}

	var content struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal(fields[dskey.MustKey("projection/1/content")], &content); err != nil {
		t.Fatalf("Can not unmarshal field projection/1/content `%s`: %v", fields[dskey.MustKey("projection/1/content")], err)
	}

	if content.Error == "" {
		t.Errorf("Field has not error")
	}
}

func TestOnTwoProjections(t *testing.T) {
	// Test that when reading two different projections at the same time in
	// different goroutines, there is no race condition.
	//
	// This test is only usefull, when the race detector is enabled.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	key1 := dskey.MustKey("projection/1/content")
	key2 := dskey.MustKey("projection/2/content")

	ds, bg := dsmock.NewMockDatastore(dsmock.YAMLData(`---
	projection:
		1:
			content_object_id: meeting/1
			type: test_model

		2:
			content_object_id: meeting/1
			type: test_model
	`))
	go bg(ctx, oserror.Handle)

	projector.Register(ds, testSlides())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		_, err := ds.Get(ctx, key1)
		if err != nil {
			t.Errorf("Get returned unexpected error: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		_, err := ds.Get(ctx, key2)
		if err != nil {
			t.Errorf("Get returned unexpected error: %v", err)
		}
	}()

	wg.Wait()
}

func testSlides() *projector.SlideStore {
	s := new(projector.SlideStore)
	s.RegisterSliderFunc("test1", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		return []byte(`{"value":"abc"}`), nil
	})

	s.RegisterSliderFunc("test_model", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		var field json.RawMessage
		fetch.Fetch(ctx, &field, "test_model/1/field")
		if field == nil {
			return []byte(`{"value":"test_model"}`), nil
		}
		return []byte(fmt.Sprintf(`{"value":"calculated with %s"}`, string(field[1:len(field)-1]))), nil
	})

	s.RegisterSliderFunc("projection", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		bs, err := json.Marshal(p7on)
		return bs, err
	})
	return s
}
