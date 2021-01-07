package collection_test

import (
	"context"
	"testing"

	user "github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestAgendaRead(t *testing.T) {
	tdp := tests.NewTestDataProvider()
	tdp.AddUserToMeeting(1, 1)

	// Normal
	tdp.AddBasicModel("agenda_item", 1)
	tdp.Set("agenda_item/1/is_hidden", "false")
	tdp.Set("agenda_item/1/is_internal", "false")

	// Hidden
	tdp.AddBasicModel("agenda_item", 2)
	tdp.Set("agenda_item/2/is_hidden", "true")
	tdp.Set("agenda_item/2/is_internal", "false")

	// Internal
	tdp.AddBasicModel("agenda_item", 3)
	tdp.Set("agenda_item/3/is_hidden", "false")
	tdp.Set("agenda_item/3/is_internal", "true")

	// Hidden and Internal
	tdp.AddBasicModel("agenda_item", 4)
	tdp.Set("agenda_item/4/is_hidden", "true")
	tdp.Set("agenda_item/4/is_internal", "true")

	dp := dataprovider.DataProvider{External: tdp}
	a := user.NewAgendaItem(dp)
	hs := new(tests.HandlerStoreMock)
	a.Connect(hs)
	read := hs.ReadHandler["agenda_item"]

	fqfields := mustFQfields(
		"agenda_item/1/id",
		"agenda_item/1/duration",
		"agenda_item/1/comment",
		"agenda_item/2/id",
		"agenda_item/3/id",
		"agenda_item/4/id",
	)

	t.Run("unprivileged user", func(t *testing.T) {
		tdp.Set("group/1/permissions", "[]")
		r := make(map[string]bool)

		if err := read.RestrictFQFields(context.Background(), 1, fqfields, r); err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}

		checkRead(t, r)
	})

	t.Run("can_see", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["agenda_item.can_see"]`)
		r := make(map[string]bool)

		if err := read.RestrictFQFields(context.Background(), 1, fqfields, r); err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}

		checkRead(t, r, "agenda_item/1/id")
	})

	t.Run("can_see_internal", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["agenda_item.can_see_internal"]`)
		r := make(map[string]bool)

		if err := read.RestrictFQFields(context.Background(), 1, fqfields, r); err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}

		checkRead(t, r, "agenda_item/1/id", "agenda_item/1/duration", "agenda_item/3/id")
	})

	t.Run("can_manage", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["agenda_item.can_manage"]`)
		r := make(map[string]bool)

		if err := read.RestrictFQFields(context.Background(), 1, fqfields, r); err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}

		checkRead(t, r, "agenda_item/1/id", "agenda_item/1/duration", "agenda_item/1/comment", "agenda_item/2/id", "agenda_item/3/id", "agenda_item/4/id")
	})
}
