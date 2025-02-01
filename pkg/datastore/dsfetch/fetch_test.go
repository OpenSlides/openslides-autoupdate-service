package dsfetch_test

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsrecorder"
)

func TestRequestTwoWithErrors(t *testing.T) {
	ds := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
	topic/1/title: foo
	`)))

	_, err := ds.Topic_Title(2).Value(context.Background())
	if err == nil {
		t.Fatalf("Title2 returned no error, expected DoesNotExist")
	}

	_, err = ds.Topic_Title(1).Value(context.Background())
	if err != nil {
		t.Errorf("Title1 returned unexpected error: %v", err)
	}
}

func TestRequestTwoWithErrorsOften(t *testing.T) {
	for i := 0; i < 100; i++ {
		TestRequestTwoWithErrors(t)
	}
}

func TestRequestEmpty(t *testing.T) {
	counter := dsmock.NewCounter(dsmock.Stub(nil)).(*dsmock.Counter)
	ds := dsfetch.New(counter)

	if err := ds.Execute(context.Background()); err != nil {
		t.Fatalf("Execute returned: %v", err)
	}

	if got := counter.Count(); got != 0 {
		t.Errorf("Got %d requests, expected 0: %v", got, counter.Requests())
	}
}

func TestFetch_required_field_when_object_does_not_exist(t *testing.T) {
	ds := dsfetch.New(dsmock.Stub(dsmock.YAMLData(``)))

	var username string
	ds.User_Username(404).Lazy(&username)

	ctx := context.Background()
	err := ds.Execute(ctx)

	var errDoesNotExist dsfetch.DoesNotExistError
	if !errors.As(err, &errDoesNotExist) {
		t.Errorf("got err `%v`, expected DoesNotExistError", err)
	}
}

func TestFech_two_times_lazy(t *testing.T) {
	ctx := context.Background()

	ds := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
	user/1/username: max
	`)))

	var username1 string
	var username2 string
	ds.User_Username(1).Lazy(&username1)
	ds.User_Username(1).Lazy(&username2)

	if err := ds.Execute(ctx); err != nil {
		t.Fatalf("Execute: %v", err)
	}

	if username1 != "max" || username2 != "max" {
		t.Errorf("usernames: '%s', '%s', expected two times max", username1, username2)
	}
}

func TestFech_error_on_value_followed_by_execute(t *testing.T) {
	ctx := context.Background()

	ds := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
	user/1/username: max
	`)))

	if _, err := ds.User_Username(2).Value(ctx); err == nil {
		t.Fatalf("fetching non existing user did not return an error")
	}

	var username1 string
	ds.User_Username(1).Lazy(&username1)

	if err := ds.Execute(ctx); err != nil {
		t.Errorf("Execute: %v", err)
	}
}

func TestFech_lazy_after_value_should_not_fetch_value_again(t *testing.T) {
	ctx := context.Background()

	stub := dsmock.Stub(dsmock.YAMLData(`---
	user/1/username: max
	user/2/username: muster
	`))

	recorder := dsrecorder.New(stub)

	ds := dsfetch.New(recorder)

	if _, err := ds.User_Username(1).Value(ctx); err != nil {
		t.Fatalf("Error: %v", err)
	}

	recorder.Reset()

	var username2 string
	ds.User_Username(2).Lazy(&username2)

	if err := ds.Execute(ctx); err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(recorder.Keys()) != 2 {
		t.Errorf("Expecting 2 keys (user/2/username and user/2/id), got: %v", recorder.Keys())
	}
}

func TestCollection(t *testing.T) {
	ctx := context.Background()

	fetch := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
	tag/1:
		name: my-tag
		tagged_ids: ["motion/1", "assignment/2"]
		meeting_id: 30
	`)))

	tag, err := fetch.Tag(1).Value(ctx)
	if err != nil {
		t.Fatalf("loading Tag: %v", err)
	}

	expectedTag := dsfetch.Tag{
		ID:        1,
		Name:      "my-tag",
		TaggedIDs: []string{"motion/1", "assignment/2"},
		MeetingID: 30,
	}

	if !reflect.DeepEqual(tag, expectedTag) {
		t.Errorf("%v != %v", tag, expectedTag)
	}
}
