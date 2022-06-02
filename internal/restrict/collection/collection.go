package collection

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, id ...int) ([]int, error)

type singleFieldRestricter func(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, id int) (bool, error)

// Allways is a restricter func that just returns true.
func Allways(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, elementIDs ...int) ([]int, error) {
	return elementIDs, nil
}

func loggedIn(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, elementIDs ...int) ([]int, error) {
	if mperms.UserID() != 0 {
		return elementIDs, nil
	}
	return nil, nil
}

func never(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, elementIDs ...int) ([]int, error) {
	return nil, nil
}

func todoToSingle(f singleFieldRestricter) FieldRestricter {
	return func(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, elementIDs ...int) ([]int, error) {
		allowedIDs := make([]int, 0, len(elementIDs))
		for _, id := range elementIDs {
			allowed, err := f(ctx, ds, mperms, id)
			if err != nil {
				var errDoesNotExist dsfetch.DoesNotExistError
				if !errors.As(err, &errDoesNotExist) {
					return nil, fmt.Errorf("restrict element %d: %w", id, err)
				}
			}

			if allowed {
				allowedIDs = append(allowedIDs, id)
			}
		}

		return allowedIDs, nil
	}
}

// Restricter returns a fieldRestricter for a restriction_mode.
//
// The FieldRestricter is a function that tells, if a user can see fields in
// that mode.
type Restricter interface {
	Modes(mode string) FieldRestricter

	// MeetingID returns the meeting id for an object. Returns hasMeeting=false,
	// if the object does not belong to a meeting.
	MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (meetingID int, hasMeeting bool, err error)
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
	case "chat_message":
		return ChatMessage{}
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
		return Unknown{}
	}
}

// Unknown is a collection that does not exist in the models.yml
type Unknown struct{}

// Modes on an unknown field can not be seen.
func (u Unknown) Modes(string) FieldRestricter {
	return never
}

// MeetingID is not a thing on a unknown meeting
func (u Unknown) MeetingID(context.Context, *dsfetch.Fetch, int) (int, bool, error) {
	return 0, false, nil
}

func eachMeeting(ctx context.Context, ds *dsfetch.Fetch, r Restricter, ids []int, f func(meetingID int, ids []int) ([]int, error)) ([]int, error) {
	meetingToIDs := make(map[int][]int)
	for _, id := range ids {
		meetingID, hasMeeting, err := r.MeetingID(ctx, ds, id)
		if err != nil {
			return nil, fmt.Errorf("getting meeting id of element %d: %w", id, err)
		}
		if !hasMeeting {
			return nil, fmt.Errorf("calling eachMeeting for object, that has no meeting")
		}
		meetingToIDs[meetingID] = append(meetingToIDs[meetingID], id)
	}

	var allAllowed []int
	for meetingID, ids := range meetingToIDs {
		allowed, err := f(meetingID, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for meeting %d: %w", meetingID, err)
		}

		allAllowed = append(allAllowed, allowed...)
	}

	return allAllowed, nil
}

func eachField(ctx context.Context, toField func(int) *dsfetch.ValueInt, ids []int, f func(id int, ids []int) ([]int, error)) ([]int, error) {
	filteredIDs := make(map[int][]int)
	for _, id := range ids {
		fieldID, err := toField(id).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting id for element %d: %w", id, err)
		}
		filteredIDs[fieldID] = append(filteredIDs[fieldID], id)
	}

	var allAllowed []int
	for filteredID, ids := range filteredIDs {
		allowed, err := f(filteredID, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for element %d: %w", filteredID, err)
		}

		allAllowed = append(allAllowed, allowed...)
	}

	return allAllowed, nil
}
