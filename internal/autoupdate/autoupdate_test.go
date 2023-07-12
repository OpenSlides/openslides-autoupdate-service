package autoupdate_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/keysbuilder"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

func TestSingleDataEmptyValues(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds, _ := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin
	`))
	s, _, _ := autoupdate.New(environment.ForTests{}, ds, RestrictAllowed)

	kb, err := keysbuilder.FromKeys("user/1/username")
	if err != nil {
		t.Fatalf("keysbuilder from keys: %v", err)
	}

	data, err := s.SingleData(ctx, 1, kb)
	if err != nil {
		t.Errorf("SingleData: %v", err)
	}

	if len(data) != 0 {
		t.Errorf("Got %v, expected empty dict", data)
	}
}

func TestRestrictFQIDs(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ds, _ := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1:
			username: superadmin
			first_name: kevin
			last_name: foo
	`))
	s, _, _ := autoupdate.New(environment.ForTests{}, ds, RestrictAllowed)

	got, err := s.RestrictFQIDs(ctx, 1, []string{"user/1"}, map[string][]string{"user": {"id", "username", "first_name"}})
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
