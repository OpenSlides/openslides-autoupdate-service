package collection_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
)

func TestAssignmentSee(t *testing.T) {
	var a collection.Assignment

	testData{
		"Without perms",
		``,
		nil,
		false,
	}.test(t, a.See)
	testData{
		"Can see",
		``,
		[]perm.TPermission{perm.AssignmentCanSee},
		false,
	}.test(t, a.See)
}

func TestAssignmentModeA(t *testing.T) {

}
