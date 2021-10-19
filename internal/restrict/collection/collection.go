package collection

import (
	"context"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, id int) (bool, error)

// Allways is a restricter func that just returns true.
func Allways(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, elementID int) (bool, error) {
	return true, nil
}

func loggedIn(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, elementID int) (bool, error) {
	return mperms.UserID() != 0, nil
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
	case "motion":
		return Motion{}
	case "motion_block":
		return MotionBlock{}
	case "motion_category":
		return MotionCategory{}
	case "motion_change_recommendation":
		return MotionChangeRecommendation{}
	case "motion_state":
		return MotionState{}
	case "motion_statute_paragraph":
		return MotionStatuteParagraph{}
	case "motion_comment":
		return MotionComment{}
	case "motion_comment_section":
		return MotionCommentSection{}
	case "motion_submitter":
		return MotionSubmitter{}
	case "motion_workflow":
		return MotionWorkflow{}
	case "option":
		return Option{}
	case "organization":
		return Organization{}
	case "organization_tag":
		return OrganizationTag{}
	case "personal_note":
		return PersonalNote{}
	case "poll":
		return Poll{}
	case "projection":
		return Projection{}
	case "projector":
		return Projector{}
	case "projector_countdown":
		return ProjectorCountdown{}
	case "projector_message":
		return ProjectorMessage{}
	case "resource":
		return Resource{}
	case "speaker":
		return Speaker{}
	case "tag":
		return Tag{}
	case "theme":
		return Theme{}
	case "topic":
		return Topic{}
	case "user":
		return User{}
	case "vote":
		return Vote{}

	default:
		return nil
	}
}
