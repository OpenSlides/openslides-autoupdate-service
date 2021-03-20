package slide

import "github.com/openslides/openslides-autoupdate-service/internal/projector"

// Slides returns all OpenSlides-Slides.
func Slides() *projector.SlideStore {
	s := new(projector.SlideStore)
	User(s)
	ListOfSpeaker(s)
	return s
}
