package agenda

import (
	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// ListOfSpeaker is holds the permission logic for list of speakers.
type ListOfSpeaker struct {
	dp dataprovider.DataProvider
}

// NewListOfSpeaker creates a new ListOfSpeakers object.
func NewListOfSpeaker(dp dataprovider.DataProvider) *ListOfSpeaker {
	return &ListOfSpeaker{
		dp: dp,
	}
}

// Connect connects the list_of_speakers routes.
func (l *ListOfSpeaker) Connect(s collection.HandlerStore) {
	writePerm := "agenda.can_manage_list_of_speakers"
	readPerm := "agenda.can_see"
	col := "list_of_speakers"

	s.RegisterWriteHandler("list_of_speakers.update", collection.Create(l.dp, writePerm, col))
	s.RegisterWriteHandler("list_of_speakers.delete_all_speakers", collection.Modify(l.dp, writePerm, col))
	s.RegisterWriteHandler("list_of_speakers.re_add_last", collection.Modify(l.dp, writePerm, col))

	s.RegisterReadHandler("list_of_speakers", collection.Restrict(l.dp, readPerm, col))
}
