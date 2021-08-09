package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, id int) (bool, error)

// Allways is a restricter func that just returns true.
func Allways(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, elementID int) (bool, error) {
	return true, nil
}

// Restricter returns a fieldRestricter for a restriction_mode.
//
// The FieldRestricter is a function that tells, if a user can see fields in
// that mode.
type Restricter interface {
	Modes(mode string) FieldRestricter
}

// Collection returns the restricter for a collection
func Collection(collection string) Restricter {
	switch collection {
	case "agenda_item":
		return AgendaItem{}
	case "assignment":
		return Assignment{}
	case "assignment_candidate":
		return AssignmentCandidate{}
	case "list_of_speakers":
		return ListOfSpeakers{}
	case "chat_group":
		return ChatGroup{}
	case "committee":
		return Committee{}
	case "group":
		return Group{}
	case "mediafile":
		return Mediafile{}
	case "meeting":
		return Meeting{}
	case "personal_note":
		return PersonalNote{}
	default:
		return nil
	}
}
