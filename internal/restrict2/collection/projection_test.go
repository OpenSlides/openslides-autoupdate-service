package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestProjectionModeA(t *testing.T) {
	t.Parallel()
	f := collection.Projection{}.Modes("A")

	testCase(
		"can see",
		t,
		f,
		true,
		"projection/1/meeting_id: 30",
		withPerms(30, perm.ProjectorCanSee),
	)

	testCase(
		"no perm",
		t,
		f,
		false,
		"projection/1/meeting_id: 30",
	)
}
