package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// CM is the name of a collection and a mode.
type CM struct {
	Collection string
	Mode       string
}

func (cm CM) String() string {
	return cm.Collection + "/" + cm.Mode
}

// Attributes is are flags that each field has. A user is allowed to see a
// field, if he has one of the attribute-fields.
type Attributes struct {
	// GlobalPermission is like the orga permission:
	//
	// 0: All user can see this
	// 1: Only superadmin can see this
	// 2: Global managers can see this
	// 3: Global user managers can see this
	// 254: Logged in users can see this
	// 255: Nobody can see this (not even the superadmin)
	GlobalPermission byte

	// GroupIDs are groups, that can see the field. Groups are meeting specific.
	GroupIDs set.Set[int]

	// UserIDs are list from users that can see the field but do not have the
	// globalPermission or are not in the groups.
	UserIDs set.Set[int]

	// GroupAnd is another attribute. If the user can see the Attribute because
	// of the Group, he also needs the GroupAnd attribute.
	GroupAnd *Attributes
}

// AttributeMap is like restrict.AttributeMap
type AttributeMap interface {
	Add(modeKey dskey.Key, value *Attributes)
	Get(ctx context.Context, fetch *dsfetch.Fetch, mperms perm.MeetingPermission, modeKEy dskey.Key) (*Attributes, error)
	SameAs(ctx context.Context, fetch *dsfetch.Fetch, mperms perm.MeetingPermission, toModeKey, fromModeKey dskey.Key) error
}

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, ids ...int) error

var allwaysAttr = Attributes{
	GlobalPermission: 0,
}

var loggedInAttr = Attributes{
	GlobalPermission: 254,
}

var neverAttr = Attributes{
	GlobalPermission: 0,
}

// Allways is a restricter func that just returns true.
func Allways(collection string, mode string) FieldRestricter {
	return func(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, elementIDs ...int) error {
		for _, id := range elementIDs {
			attrMap.Add(dskey.Key{Collection: collection, ID: id, Field: mode}, &allwaysAttr)
		}
		return nil
	}
}

func loggedIn(collection string, mode string) FieldRestricter {
	return func(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, elementIDs ...int) error {
		for _, id := range elementIDs {
			attrMap.Add(dskey.Key{Collection: collection, ID: id, Field: mode}, &loggedInAttr)
		}
		return nil
	}
}

func never(collection string, mode string) FieldRestricter {
	return func(ctx context.Context, ds *dsfetch.Fetch, mperms perm.MeetingPermission, attrMap AttributeMap, elementIDs ...int) error {
		for _, id := range elementIDs {
			attrMap.Add(dskey.Key{Collection: collection, ID: id, Field: mode}, &neverAttr)
		}
		return nil
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

	Name() string
}

// Collection returns the restricter for a collection
func Collection(collection string) Restricter {
	switch collection {
	case "action_worker":
		return ActionWorker{name: collection}
	case "agenda_item":
		return AgendaItem{name: collection}
	case "assignment":
		return Assignment{name: collection}
	case "assignment_candidate":
		return AssignmentCandidate{name: collection}
	case "list_of_speakers":
		return ListOfSpeakers{name: collection}
	case "chat_group":
		return ChatGroup{name: collection}
	case "chat_message":
		return ChatMessage{name: collection}
	case "committee":
		return Committee{name: collection}
	case "group":
		return Group{name: collection}
	case "mediafile":
		return Mediafile{name: collection}
	case "meeting":
		return Meeting{name: collection}
	case "motion":
		return Motion{name: collection}
	case "motion_block":
		return MotionBlock{name: collection}
	case "motion_category":
		return MotionCategory{name: collection}
	case "motion_change_recommendation":
		return MotionChangeRecommendation{name: collection}
	case "motion_state":
		return MotionState{name: collection}
	case "motion_statute_paragraph":
		return MotionStatuteParagraph{name: collection}
	case "motion_comment":
		return MotionComment{name: collection}
	case "motion_comment_section":
		return MotionCommentSection{name: collection}
	case "motion_submitter":
		return MotionSubmitter{name: collection}
	case "motion_workflow":
		return MotionWorkflow{name: collection}
	case "option":
		return Option{name: collection}
	case "organization":
		return Organization{name: collection}
	case "organization_tag":
		return OrganizationTag{name: collection}
	case "personal_note":
		return PersonalNote{name: collection}
	case "poll":
		return Poll{name: collection}
	case "projection":
		return Projection{name: collection}
	case "projector":
		return Projector{name: collection}
	case "projector_countdown":
		return ProjectorCountdown{name: collection}
	case "projector_message":
		return ProjectorMessage{name: collection}
	case "speaker":
		return Speaker{name: collection}
	case "tag":
		return Tag{name: collection}
	case "theme":
		return Theme{name: collection}
	case "topic":
		return Topic{name: collection}
	case "user":
		return User{name: collection}
	case "vote":
		return Vote{name: collection}

	default:
		return Unknown{name: collection}
	}
}

// Unknown is a collection that does not exist in the models.yml
type Unknown struct {
	name string
}

// Name returns the collection name.
func (u Unknown) Name() string {
	return u.name
}

// Modes on an unknown field can not be seen.
func (u Unknown) Modes(string) FieldRestricter {
	return never(u.name, "any")
}

// MeetingID is not a thing on a unknown meeting
func (u Unknown) MeetingID(context.Context, *dsfetch.Fetch, int) (int, bool, error) {
	return 0, false, nil
}

func eachMeeting(ctx context.Context, ds *dsfetch.Fetch, r Restricter, ids []int, f func(meetingID int, ids []int) error) error {
	meetingToIDs := make(map[int][]int)
	for _, id := range ids {
		meetingID, hasMeeting, err := r.MeetingID(ctx, ds, id)
		if err != nil {
			return fmt.Errorf("getting meeting id of element %d: %w", id, err)
		}
		if !hasMeeting || meetingID == 0 {
			return fmt.Errorf("element with id %d has no meeting", id)
		}

		meetingToIDs[meetingID] = append(meetingToIDs[meetingID], id)
	}

	for meetingID, ids := range meetingToIDs {
		if err := f(meetingID, ids); err != nil {
			return fmt.Errorf("restricting for meeting %d: %w", meetingID, err)
		}
	}

	return nil
}

func meetingPerm(ctx context.Context, ds *dsfetch.Fetch, r Restricter, mode string, ids []int, mperms perm.MeetingPermission, permission perm.TPermission, attrMap AttributeMap) error {
	return eachMeeting(ctx, ds, r, ids, func(meetingID int, ids []int) error {
		groupMap, err := mperms.Meeting(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("get groups with permission %s: %w", permission, err)
		}

		attr := Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         groupMap[permission],
		}

		for _, id := range ids {
			attrMap.Add(dskey.Key{Collection: r.Name(), ID: id, Field: mode}, &attr)
		}
		return nil
	})
}

func eachRelationField(ctx context.Context, toField func(int) *dsfetch.ValueInt, ids []int, f func(id int, ids []int) error) error {
	filteredIDs := make(map[int][]int)
	for _, id := range ids {
		fieldID, err := toField(id).Value(ctx)
		if err != nil {
			return fmt.Errorf("getting id for element %d: %w", id, err)
		}
		if fieldID == 0 {
			// TODO Last Error
			return fmt.Errorf("element with id %d has no relation", id)
		}
		filteredIDs[fieldID] = append(filteredIDs[fieldID], id)
	}

	for fieldID, ids := range filteredIDs {
		err := f(fieldID, ids)
		if err != nil {
			return fmt.Errorf("restricting for element %d: %w", fieldID, err)
		}
	}

	return nil
}

// TODO: Can probably be removed
func eachStringField(ctx context.Context, toField func(int) *dsfetch.ValueString, ids []int, f func(value string, ids []int) ([]int, error)) ([]int, error) {
	filteredIDs := make(map[string][]int)
	for _, id := range ids {
		value, err := toField(id).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting value for element %d: %w", id, err)
		}
		filteredIDs[value] = append(filteredIDs[value], id)
	}

	allAllowed := make([]int, 0, len(ids))
	for value, ids := range filteredIDs {
		allowed, err := f(value, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for element %s: %w", value, err)
		}

		allAllowed = append(allAllowed, allowed...)
	}

	return allAllowed, nil
}

// TODO: currently, this calls the function with the same collectionObject
// (motion/5, motion/5), but it should bundle it by collection (motion/1,
// motion/2).
func eachContentObjectCollection(ctx context.Context, toField func(int) *dsfetch.ValueString, ids []int, f func(collection string, id int, ids []int) error) error {
	filteredIDs := make(map[string][]int)
	for _, id := range ids {
		contentObjectID, err := toField(id).Value(ctx)
		if err != nil {
			return fmt.Errorf("getting id for element %d: %w", id, err)
		}

		filteredIDs[contentObjectID] = append(filteredIDs[contentObjectID], id)
	}

	for contentObjectID, ids := range filteredIDs {
		collection, objectID, found := strings.Cut(contentObjectID, "/")
		if !found {
			return fmt.Errorf("content object_id has to have exacly one /, got %q", contentObjectID)
		}

		id, err := strconv.Atoi(objectID)
		if err != nil {
			return fmt.Errorf("second part of content_object_id has to be int, got %q", objectID)
		}

		if err := f(collection, id, ids); err != nil {
			return fmt.Errorf("restricting for element %s: %w", contentObjectID, err)
		}
	}

	return nil
}

// TODO: Can probably be removed
func eachCondition(ids []int, f func(id int) (bool, error)) ([]int, error) {
	allowed := make([]int, 0, len(ids))
	for _, id := range ids {
		ok, err := f(id)
		if err != nil {
			return nil, fmt.Errorf("checking for element %d: %w", id, err)
		}

		if ok {
			allowed = append(allowed, id)
		}
	}
	return allowed, nil
}
