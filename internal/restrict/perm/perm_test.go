package perm_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

func TestHasSuperAdmin(t *testing.T) {
	ctx := context.Background()
	ds, _ := dsmock.NewMockDatastore(dsmock.YAMLData(`---
		user/1/organization_management_level: superadmin

		meeting/3/id: 3
	`))

	p, err := perm.New(ctx, dsfetch.New(ds), 1, 3)
	if err != nil {
		t.Fatalf("perm.New(): %v", err)
	}

	if !p.Has(perm.AgendaItemCanSee) {
		t.Errorf("p.Has(perm.AgendaItemCanSee) returned false, expected true for any perm")
	}
}

func TestPermissionMap(t *testing.T) {
	ctx := context.Background()
	ds := dsmock.Stub(dsmock.YAMLData(`---
		meeting:
			7:
				group_ids: [100,101,102,103]
				admin_group_id: 103
		
		group:
			100:
				permissions:
				- agenda_item.can_manage
			101:
				permissions:
				- agenda_item.can_see
			102:
				permissions: []
			103:
				permissions: []
	`))

	got, err := perm.GroupByPerm(ctx, dsfetch.New(ds), 7)
	if err != nil {
		t.Fatalf("GroupByPerm: %v", err)
	}

	expect := map[perm.TPermission]set.Set[int]{
		perm.AgendaItemCanManage:      set.New(100, 103),
		perm.AgendaItemCanSeeInternal: set.New(100, 103),
		perm.AgendaItemCanSee:         set.New(100, 101, 103),
	}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got:\n%v\n\nexpected:\n%v", got, expect)
	}
}
