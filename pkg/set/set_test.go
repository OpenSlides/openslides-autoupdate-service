package set_test

import (
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

func TestLen(t *testing.T) {
	s := set.New(1, 2, 3, 4, 5)
	if got := s.Len(); got != 5 {
		t.Errorf("set.Len() == %d, expected 5", got)
	}
}
