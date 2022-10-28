package autoupdate_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func TestSingleDataEmptyValues(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin
	`))
	s, _ := autoupdate.New(ds, RestrictAllowed)

	kb, err := keysbuilder.FromKeys("user/1/username")
	if err != nil {
		t.Fatalf("keysbuilder from keys: %v", err)
	}

	data, err := s.SingleData(ctx, 1, kb, 0)
	if err != nil {
		t.Errorf("SingleData: %v", err)
	}

	if len(data) != 0 {
		t.Errorf("Got %v, expected empty dict", data)
	}
}

func TestHistoryInformation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin
	`))
	s, _ := autoupdate.New(ds, RestrictAllowed)

	buf := new(bytes.Buffer)
	err := s.HistoryInformation(ctx, 1, "collection/1", buf)

	if err != nil {
		t.Fatalf("HistoryInformation: %v", err)
	}

	var information []interface{}
	if err := json.Unmarshal(buf.Bytes(), &information); err != nil {
		t.Fatalf("HistoryInformation returned invalid data `%v`: %v", buf.String(), err)
	}

	if len(information) == 0 {
		t.Errorf("No History returned")
	}
}

func TestHistoryInformationWrongFQID(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin
	`))
	s, _ := autoupdate.New(ds, RestrictAllowed)

	buf := new(bytes.Buffer)
	err := s.HistoryInformation(ctx, 1, "collection", buf)

	var errType interface {
		Type() string
	}
	if !errors.As(err, &errType) || errType.Type() != "invalid_input" {
		t.Errorf("Got error `%v`, expected error with type `invalid_input`", err)
	}

	if buf.Len() != 0 {
		t.Errorf("got %s, expected no output", buf)
	}
}

func TestHistoryInformationSuperAdminOnMeetingCollection(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin

		motion/5/meeting_id: 1
	`))
	s, _ := autoupdate.New(ds, RestrictAllowed)

	buf := new(bytes.Buffer)
	err := s.HistoryInformation(ctx, 1, "motion/5", buf)

	if err != nil {
		t.Fatalf("HistoryInformation: %v", err)
	}

	var information []interface{}
	if err := json.Unmarshal(buf.Bytes(), &information); err != nil {
		t.Fatalf("HistoryInformation returned invalid data `%v`: %v", buf.String(), err)
	}

	if len(information) == 0 {
		t.Errorf("No History returned")
	}
}

func TestRestrictFQIDs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1:
			username: superadmin
			first_name: kevin
	`))
	s, _ := autoupdate.New(ds, RestrictAllowed)

	got, err := s.RestrictFQIDs(ctx, 1, []string{"user/1"})
	if err != nil {
		t.Fatalf("RestrictFQIDs: %v", err)
	}

	expect := map[string]map[string][]byte{
		"user/1": {
			"id":         []byte("1"),
			"username":   []byte(`"superadmin"`),
			"first_name": []byte(`"kevin"`),
		},
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("\nGot\t\t\t%v\nexpected\t%v", got, expect)
	}
}
