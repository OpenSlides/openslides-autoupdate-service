package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
)

func TestOrganizationTagModeA(t *testing.T) {
	f := collection.OrganizationTag{}.Modes("A")

	testCase(
		"Public access",
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
