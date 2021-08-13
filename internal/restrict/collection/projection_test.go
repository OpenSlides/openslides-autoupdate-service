package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
)

func TestProjectionModeA(t *testing.T) {
	f := collection.Projection{}.Modes("A")

	testCase(
		"can see",
		t,
		f,
		true,
		"projection/1/meeting_id: 1",
		withPerms(1, perm.ProjectorCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"projection/1/meeting_id: 1",
	)
}
