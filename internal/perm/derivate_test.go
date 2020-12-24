package perm_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestDerivatePerm(t *testing.T) {
	tdp := tests.NewTestDataProvider()
	tdp.AddUser(1)
	tdp.AddUserToMeeting(1, 1)
	tdp.AddUserToGroup(1, 1, 2)
	dp := dataprovider.DataProvider{External: tdp}

	p, err := perm.Perms(context.Background(), 1, 1, dp)

	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if !p.HasOne("motion.can_see") {
		t.Errorf("User does not have can_see permission")
	}
}
