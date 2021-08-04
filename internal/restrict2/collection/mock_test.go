package collection_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/test"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

type testData struct {
	name   string
	data   string
	perms  []perm.TPermission
	expect bool
}

func (tt testData) test(t *testing.T, f collection.FieldRestricter) {
	t.Helper()

	t.Run(tt.name, func(t *testing.T) {
		fetch := datastore.NewFetcher(dsmock.Stub(dsmock.YAMLData(tt.data)))
		perms := test.MeetingPermissionStub{UID: 1, Permissions: map[int][]perm.TPermission{
			1: tt.perms,
		}}

		got, err := f(context.Background(), fetch, perms, 1)

		if err != nil {
			t.Fatalf("See returned unexpected error: %v", err)
		}

		if got != tt.expect {
			t.Errorf("See() returned %t, expected %t", got, tt.expect)
		}
	})
}
