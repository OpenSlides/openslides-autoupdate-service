package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestListOfSpeakersSee(t *testing.T) {
	var los collection.ListOfSpeakers
	ds := `---
	list_of_speakers/1/meeting_id: 1
	`

	testData{
		"Can see internal",
		ds,
		[]perm.TPermission{perm.ListOfSpeakersCanSee},
		true,
	}.test(t, los.See)

	testData{
		"Can not see internal",
		ds,
		nil,
		false,
	}.test(t, los.See)
}

func TestListOfSpeakersModeA(t *testing.T) {
	var los collection.ListOfSpeakers

	testData{
		"simple",
		``,
		nil,
		true,
	}.test(t, los.Modes("A"))
}
