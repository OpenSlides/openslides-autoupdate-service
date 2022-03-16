package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
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
