package listofspeakers

import "github.com/OpenSlides/openslides-permission-service/internal/allowed"

// Update TODO
var Update = allowed.BuildModify([]string{
	"id",
	"closed",
}, "list_of_speakers", "agenda.can_manage_list_of_speakers")

// DeleteAllSpeakers TODO
var DeleteAllSpeakers = allowed.BuildModify([]string{"id"}, "list_of_speakers", "agenda.can_manage_list_of_speakers")

// ReAddLast TODO
var ReAddLast = allowed.BuildModify([]string{"id"}, "list_of_speakers", "agenda.can_manage_list_of_speakers")
