package dsfetch_test

import (
	"context"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
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

	ds := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
	user/1/username: max
	`)))

	var username1 string
	var username2 string
	ds.User_Username(1).Lazy(&username1)
	ds.User_Username(1).Lazy(&username2)

	ctx := context.Background()

	if err := ds.Execute(ctx); err != nil {
		t.Fatalf("Execute: %v", err)
	}

	if username1 != "max" || username2 != "max" {
		t.Errorf("usernames: '%s', '%s', expected two times max", username1, username2)
	}
}
