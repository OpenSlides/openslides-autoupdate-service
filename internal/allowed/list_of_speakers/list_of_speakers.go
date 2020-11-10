package list_of_speakers

import "github.com/OpenSlides/openslides-permission-service/internal/allowed"

var Update = allowed.BuildModify([]string{
	"id",
	"closed",
}, "list_of_speakers", "agenda.can_manage_list_of_speakers")

var DeleteAllSpeakers = allowed.BuildModify([]string{"id"}, "list_of_speakers", "agenda.can_manage_list_of_speakers")

var ReAddLast = allowed.BuildModify([]string{"id"}, "list_of_speakers", "agenda.can_manage_list_of_speakers")
