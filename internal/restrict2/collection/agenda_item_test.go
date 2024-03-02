package collection_test

import (
	"context"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

func TestAgendaModeA(t *testing.T) {
	t.Parallel()
	var a collection.AgendaItem

	testCase(
		"No permission",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1/meeting_id: 30
		`,
	)

	testCase(
		"manager",
		t,
		a.Modes("A"),
		true,
		`---
		agenda_item/1/meeting_id: 30
		`,
		withPerms(30, perm.AgendaItemCanManage),
	)

	testCase(
		"Can see internal with hidden",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: true
		`,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can see internal not hidden",
		t,
		a.Modes("A"),
		true,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: false
		`,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can see with hidden and internal",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: true
			is_internal: true
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can see not hidden but internal",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: false
			is_internal: true
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can see with hidden but not internal",
		t,
		a.Modes("A"),
		false,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: true
			is_internal: false
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can see not hidden and not internal",
		t,
		a.Modes("A"),
		true,
		`---
		agenda_item/1:
			meeting_id: 30
			is_hidden: false
			is_internal: false
		`,
		withPerms(30, perm.AgendaItemCanSee),
	)
}

func TestAgendaModeB(t *testing.T) {
	t.Parallel()
	var a collection.AgendaItem
	ds := `---
	agenda_item/1/meeting_id: 30
	`

	testCase(
		"Can see internal",
		t,
		a.Modes("B"),
		true,
		ds,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can not see internal",
		t,
		a.Modes("B"),
		false,
		ds,
	)
}

func TestAgendaModeC(t *testing.T) {
	t.Parallel()
	var a collection.AgendaItem
	ds := `---
	agenda_item/1/meeting_id: 30
	`

	testCase(
		"Can see internal",
		t,
		a.Modes("C"),
		false,
		ds,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can see",
		t,
		a.Modes("C"),
		false,
		ds,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can manage",
		t,
		a.Modes("C"),
		true,
		ds,
		withPerms(30, perm.AgendaItemCanManage),
	)

	testCase(
		"No perm",
		t,
		a.Modes("C"),
		false,
		ds,
	)
}

func TestAgendaModeD(t *testing.T) {
	var a collection.AgendaItem
	mode := a.Modes("D")
	ds := `---
	agenda_item/1/meeting_id: 30
	`

	testCase(
		"Can see internal",
		t,
		mode,
		false,
		ds,
		withPerms(30, perm.AgendaItemCanSeeInternal),
	)

	testCase(
		"Can see",
		t,
		mode,
		false,
		ds,
		withPerms(30, perm.AgendaItemCanSee),
	)

	testCase(
		"Can manage",
		t,
		mode,
		false,
		ds,
		withPerms(30, perm.AgendaItemCanManage),
	)

	testCase(
		"No perm",
		t,
		mode,
		false,
		ds,
	)
}

func TestAgendaItemMeetingID_0_returns_an_InvalidData_error(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	fetcher := dsfetch.New(dsmock.Stub(dsmock.YAMLData(`---
	agenda_item/1/meeting_id: 0
	`)))

	_, _, err := collection.AgendaItem{}.MeetingID(ctx, fetcher, 1)
	if err == nil {
		t.Fatalf("MeetingID did not return an error.")
	}

	var errInvalidData datastore.InvalidDataError
	if !errors.As(err, &errInvalidData) {
		t.Fatalf("MeetingID() == '%s', expected an InvalidDataError", err)
	}

	if got := errInvalidData.Key; got != dskey.MustKey("agenda_item/1/meeting_id") {
		t.Fatalf("errInvalidData.Key == %s, expected key(agenda_item/1/meeting_id)", got)
	}

	if got := errInvalidData.Value; string(got) != "0" {
		t.Fatalf("errInvalidData.Value == %s, expected 0", got)
	}
}
