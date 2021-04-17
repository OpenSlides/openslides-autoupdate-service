package slide

import "github.com/OpenSlides/openslides-autoupdate-service/internal/projector"

// Slides returns all OpenSlides-Slides.
func Slides() *projector.SlideStore {
	s := new(projector.SlideStore)
	AgendaItem(s)
	AgendaItemList(s)
	Assignment(s)
	ListOfSpeaker(s)
	CurrentListOfSpeakers(s)
	CurrentSpeakerChyron(s)
	Mediafile(s)
	Motion(s)
	MotionBlock(s)
	Poll(s)
	ProjectorCountdown(s)
	ProjectorMessage(s)
	Topic(s)
	User(s)
	return s
}
