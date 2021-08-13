package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
)

func TestOrganizationTagModeA(t *testing.T) {
	f := collection.OrganizationTag{}.Modes("A")

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
