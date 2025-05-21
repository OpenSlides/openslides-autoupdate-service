package perm_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dsmock"
)

func TestHasSuperAdmin(t *testing.T) {
	ctx := context.Background()
	ds := dsmock.Stub(dsmock.YAMLData(`---
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

func TestLockedOut(t *testing.T) {
	ctx := context.Background()
	ds := dsmock.Stub(dsmock.YAMLData(`---
		meeting/3/committee_id: 2
		user/1/meeting_user_ids: [10]
		meeting_user/10:
			meeting_id: 3
			locked_out: true
			group_ids: [30]
		group/30/permissions: ["agenda_item.can_see"]
	`))

	p, err := perm.New(ctx, dsfetch.New(ds), 1, 3)
	if err != nil {
		t.Fatalf("perm.New(): %v", err)
	}

	if p.Has(perm.AgendaItemCanSee) {
		t.Errorf("p.Has(perm.AgendaItemCanSee) returned true, expected false for any perm")
	}

	if p.InGroup(30) {
		t.Errorf("p.InGroup returned true, expected false for any group")
	}
}

func TestManagementLevelCommittees(t *testing.T) {
	ctx := t.Context()
	ds := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
		committee/1/all_child_ids: [3,4]
		user/4:
			committee_management_ids: [1]
	`)))

	got, err := perm.ManagementLevelCommittees(ctx, ds, 4)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	expect := []int{1, 3, 4}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("ManagmenentLevelCommittees() == %v, expected %v", got, expect)
	}
}

func TestCommitteeManagerPermission(t *testing.T) {
	ctx := t.Context()
	ds := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
		committee/1/all_child_ids: [3,4]
		user/4:
			committee_management_ids: [1]
		meeting/5/committee_id: 3
	`)))

	p, err := perm.New(ctx, ds, 4, 5)
	if err != nil {
		t.Fatalf("perm.New(): %v", err)
	}

	if !p.Has(perm.MotionCanManage) {
		t.Errorf("Committee Manager of higher committee does not have all permissions")
	}
}
