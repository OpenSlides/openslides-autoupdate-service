package autoupdate_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/autoupdate"
	"github.com/OpenSlides/openslides-autoupdate-service/keysbuilder"
	"github.com/OpenSlides/openslides-go/datastore/dsmock"
	"github.com/OpenSlides/openslides-go/environment"
)

func TestSingleDataEmptyValues(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flow := dsmock.NewFlow(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin
	`))
	s, _, _ := autoupdate.New(environment.ForTests{}, flow, RestrictAllowed)

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
