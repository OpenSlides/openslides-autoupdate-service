package projector_test

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/nsf/jsondiff"
)

func TestProjectionDoesNotExist(t *testing.T) {
	ctx := context.Background()

	flow := dsmock.NewFlow(nil)
	myKey := dskey.MustKey("projection/1/content")

	p := projector.NewProjector(flow, testSlides())

	got, err := p.Get(ctx, myKey)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := map[dskey.Key][]byte{myKey: nil}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("Got %v, expected %v", got, expect)
	}
}

func TestProjectionFromContentObject(t *testing.T) {
	ctx := context.Background()
	flow := dsmock.NewFlow(dsmock.YAMLData(`---
	projection/1:
		content_object_id:    user/1
		current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := []byte(`{"collection":"user","value":"user"}` + "\n")

	if equal, explain := cmpJson(got[key], expect); !equal {
		t.Errorf("got != expect: %s", explain)
	}
}

func TestProjectionFromContentObjectIfNotOnProjector(t *testing.T) {
	ctx := context.Background()
	flow := dsmock.NewFlow(dsmock.YAMLData(`---
	projection/1:
		content_object_id:    user/1
		current_projector_id: null
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	if got[key] != nil {
		t.Errorf("got %v, expected nil", got)
	}
}

func TestProjectionFromType(t *testing.T) {
	ctx := context.Background()
	flow := dsmock.NewFlow(dsmock.YAMLData(`
		projection/1:
			content_object_id:    meeting/1
			type:                 test1
			current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := []byte(`{"collection":"test1","value":"abc"}` + "\n")
	if equal, explain := cmpJson(got[key], expect); !equal {
		t.Errorf("got != expect: %s", explain)
	}
}

func TestProjectionWithOptionsData(t *testing.T) {
	ctx := context.Background()
	flow := dsmock.NewFlow(dsmock.YAMLData(`
		projection/1:
			content_object_id:    "meeting/6"
			type:                 "projection"
			meeting_id:           1
			options:              {"only_main_items": true}
			current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	expect := []byte(`{"collection":"projection","id": 1,"current_projector_id":1, "content_object_id": "meeting/6", "type":"projection", "meeting_id": 1, "options": {"only_main_items": true}}` + "\n")
	if equal, expain := cmpJson(got[key], expect); !equal {
		t.Errorf("got != expect: %s", expain)
	}
}

func TestProjectionTypeDoesNotExist(t *testing.T) {
	ctx := context.Background()
	flow := dsmock.NewFlow(dsmock.YAMLData(`
		projection/1:
			content_object_id:    meeting/1
			type:                 unexistingTestSlide

			current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	var content struct {
		Error string `json:"error"`
	}
	if err := json.Unmarshal(got[key], &content); err != nil {
		t.Fatalf("Can not unmarshal field projection/1/content `%s`: %v", got[key], err)
	}

	if content.Error == "" {
		t.Errorf("Field has not error")
	}
}

func TestProjectionUpdateProjection(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(dsmock.YAMLData(`---
		projection/1:
			content_object_id:    meeting/1
			type:                 test1
			current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	done := make(chan struct{})
	go p.Update(ctx, func(map[dskey.Key][]byte, error) {
		close(done)
	})

	// Fetch data once to fill the test.
	if _, err := p.Get(ctx, key); err != nil {
		t.Fatalf("Get: %v", err)
	}

	flow.Send(dsmock.YAMLData(`---
	projection/1:
		type: null
		content_object_id: user/1
	`))
	<-done

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Fatalf("Second Get: %v", err)
	}

	expect := []byte(`{"collection":"user","value":"user"}` + "\n")
	if equal, expain := cmpJson(got[key], expect); !equal {
		t.Errorf("got != expect: %s", expain)
	}
}

func TestProjectionUpdateProjectionMetaData(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(dsmock.YAMLData(`---
		projection/1:
			type:                 projection
			content_object_id:    meeting/1
			current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	done := make(chan struct{})
	go p.Update(ctx, func(map[dskey.Key][]byte, error) {
		close(done)
	})

	// Fetch data once to fill the hot keys.
	if _, err := p.Get(ctx, key); err != nil {
		t.Fatalf("Get: %v", err)
	}

	flow.Send(dsmock.YAMLData("projection/1/stable: true"))
	<-done

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Fatalf("Second get: %v", err)
	}

	expect := []byte(`{"collection":"projection","id": 1, "content_object_id": "meeting/1", "meeting_id":0, "type":"projection", "options": null,"current_projector_id":1}` + "\n")
	if equal, expain := cmpJson(got[key], expect); !equal {
		t.Errorf("got != expect: %s", expain)
	}
}

func TestProjectionUpdateSlide(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(dsmock.YAMLData(`---
		projection/1:
			type:                 user
			content_object_id:    meeting/6
			current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	done := make(chan struct{})
	go p.Update(ctx, func(map[dskey.Key][]byte, error) {
		close(done)
	})

	// Fetch data once to fill the hot keys.
	if _, err := p.Get(ctx, key); err != nil {
		t.Fatalf("Get: %v", err)
	}

	flow.Send(dsmock.YAMLData("user/1/username: new value"))
	<-done

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Errorf("second Get: %v", err)
	}

	expect := []byte(`{"collection":"user","value":"calculated with new value"}` + "\n")
	if equal, expain := cmpJson(got[key], expect); !equal {
		t.Errorf("got != expect: %s", expain)
	}
}

func TestProjectionUpdateOtherKey(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(dsmock.YAMLData(`---
		projection/1:
			type:                 user
			content_object_id:    meeting/1
			current_projector_id: 1
	`))
	key := dskey.MustKey("projection/1/content")
	p := projector.NewProjector(flow, testSlides())

	done := make(chan struct{})
	go p.Update(ctx, func(map[dskey.Key][]byte, error) {
		close(done)
	})

	// Fetch data once to fill the hot keys.
	if _, err := p.Get(ctx, key); err != nil {
		t.Fatalf("Get: %v", err)
	}

	flow.Send(dsmock.YAMLData("group/1/name: new value"))
	<-done

	got, err := p.Get(ctx, key)
	if err != nil {
		t.Errorf("second Get: %v", err)
	}

	expect := []byte(`{"collection":"user","value":"user"}` + "\n")
	if equal, expain := cmpJson(got[key], expect); !equal {
		t.Errorf("got != expect: %s", expain)
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

	ds := dsmock.NewFlow(dsmock.YAMLData(`---
	projection:
		1:
			content_object_id: meeting/1
			type: user

		2:
			content_object_id: meeting/1
			type: user
	`))

	p := projector.NewProjector(ds, testSlides())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		if _, err := p.Get(ctx, key1); err != nil {
			t.Errorf("Get returned unexpected error: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		if _, err := p.Get(ctx, key2); err != nil {
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

	s.RegisterSliderFunc("user", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		var field json.RawMessage
		fetch.Fetch(ctx, &field, "user/1/username")
		if field == nil {
			return []byte(`{"value":"user"}`), nil
		}
		return []byte(fmt.Sprintf(`{"value":"calculated with %s"}`, string(field[1:len(field)-1]))), nil
	})

	s.RegisterSliderFunc("projection", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		bs, err := json.Marshal(p7on)
		return bs, err
	})
	return s
}

func cmpJson(got, expect []byte) (bool, string) {
	options := jsondiff.DefaultJSONOptions()
	if cmp, explain := jsondiff.Compare(got, []byte(expect), &options); cmp != jsondiff.FullMatch {
		return false, explain
	}
	return true, ""
}
