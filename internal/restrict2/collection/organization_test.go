package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestOrganizationModeA(t *testing.T) {
	f := collection.Organization{}.Modes("A")

	testCase(
		"no perms",
		t,
		f,
		true,
		``,
	)
}

func TestOrganizationModeB(t *testing.T) {
	f := collection.Organization{}.Modes("B")

	testCase(
		"anonymous",
		t,
		f,
		false,
		``,
		withRequestUser(0),
	)

	testCase(
		"logged in",
		t,
		f,
		true,
		``,
		withRequestUser(1),
	)
}

func TestOrganizationModeC(t *testing.T) {
	f := collection.Organization{}.Modes("C")

	testCase(
		"anonymous",
		t,
		f,
		false,
		``,
		withRequestUser(0),
	)

	testCase(
		"logged in",
		t,
		f,
		false,
		``,
		withRequestUser(1),
	)

	testCase(
		"OML can manage users",
		t,
		f,
		true,
		`---
		user/1/organization_management_level: can_manage_users
		`,
		withRequestUser(1),
	)
}
