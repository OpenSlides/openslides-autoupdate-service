package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// CM is the name of a collection and a mode.
type CM struct {
	Collection string
	Mode       string
}

func (cm CM) String() string {
	return cm.Collection + "/" + cm.Mode
}

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, ds *dsfetch.Fetch, id ...int) ([]int, error)

type singleFieldRestricter func(ctx context.Context, ds *dsfetch.Fetch, id int) (bool, error)

// Allways is a restricter func that just returns true.
func Allways(ctx context.Context, ds *dsfetch.Fetch, elementIDs ...int) ([]int, error) {
	return elementIDs, nil
}

func loggedIn(ctx context.Context, ds *dsfetch.Fetch, elementIDs ...int) ([]int, error) {
	uid, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	if uid != 0 {
		return elementIDs, nil
	}
	return nil, nil
}

func never(ctx context.Context, ds *dsfetch.Fetch, elementIDs ...int) ([]int, error) {
	return nil, nil
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

// type contextKeyType string

// var contextKey contextKeyType = "context_key"

// func ContextWithCollectionCache(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, contextKey, make(map[dskey.Key]bool))
// }

// func CollectionCacheFromCache(ctx context.Context, key dskey.Key, allowed bool) map[dskey.Key]bool {
// 	v := ctx.Value(contextKey)
// 	if v == nil {
// 		return nil
// 	}

// 	cache, ok := v.(map[dskey.Key]bool)
// 	if !ok {
// 		return nil
// 	}

// 	return cache
// }

// func cacheWrapper(ctx context.Context, collection string, mode string, testIDs, allowedIDs []int, err error) ([]int, error) {
// 	if err != nil {
// 		return nil, err
// 	}

// 	allowedSet := set.New(allowedIDs...)
// 	for _, id := range testIDs {
// 		allowedSet.Has(id)
// 	}
// }

type restrictCache struct {
	cache map[dskey.Key]bool
}

func newRestrictCache() *restrictCache {
	return &restrictCache{
		cache: make(map[dskey.Key]bool),
	}
}

func (r *restrictCache) middleware(fn FieldRestricter) FieldRestricter

var collections = []Restricter{
	ActionWorker{},
	AgendaItem{},
	Assignment{},
	AssignmentCandidate{},
	ListOfSpeakers{},
	ChatGroup{},
	ChatMessage{},
	Committee{},
	Group{},
	Mediafile{},
	Meeting{},
	Motion{},
	MotionBlock{},
	MotionCategory{},
	MotionChangeRecommendation{},
	MotionState{},
	MotionStatuteParagraph{},
	MotionComment{},
	MotionCommentSection{},
	MotionSubmitter{},
	MotionWorkflow{},
	Option{},
	Organization{},
	OrganizationTag{},
	PersonalNote{},
	Poll{},
	Projection{},
	Projector{},
	ProjectorCountdown{},
	ProjectorMessage{},
	Speaker{},
	Tag{},
	Theme{},
	Topic{},
	User{},
	Vote{},
}

// Collection returns the restricter for a collection
func Collection(collection string) Restricter {
	for _, c := range collections {
		if c.Name() == collection {
			return c
		}
	}
	return Unknown{collection}
}

// Unknown is a collection that does not exist in the models.yml
type Unknown struct {
	name string
}

// Modes on an unknown field can not be seen.
func (u Unknown) Modes(string) FieldRestricter {
	return never
}

// MeetingID is not a thing on a unknown meeting
func (u Unknown) MeetingID(context.Context, *dsfetch.Fetch, int) (int, bool, error) {
	return 0, false, nil
}

// Name returns the collection name.
func (u Unknown) Name() string {
	return u.name
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
		if meetingID == 0 {
			// TODO Last Error
			return nil, fmt.Errorf("element with id %d has no meeting", id)
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

func meetingPerm(ctx context.Context, ds *dsfetch.Fetch, r Restricter, ids []int, permission perm.TPermission) ([]int, error) {
	return eachMeeting(ctx, ds, r, ids, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permission: %w", err)
		}

		if perms.Has(permission) {
			return ids, nil
		}
		return nil, nil
	})
}

func eachRelationField(ctx context.Context, toField func(int) *dsfetch.ValueInt, ids []int, f func(id int, ids []int) ([]int, error)) ([]int, error) {
	filteredIDs := make(map[int][]int)
	for _, id := range ids {
		fieldID, err := toField(id).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting id for element %d: %w", id, err)
		}
		if fieldID == 0 {
			// TODO Last Error
			return nil, fmt.Errorf("element with id %d has no relation", id)
		}
		filteredIDs[fieldID] = append(filteredIDs[fieldID], id)
	}

	allAllowed := make([]int, 0, len(ids))
	for fieldID, ids := range filteredIDs {
		allowed, err := f(fieldID, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for element %d: %w", fieldID, err)
		}

		allAllowed = append(allAllowed, allowed...)
	}

	return allAllowed, nil
}

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
func eachContentObjectCollection(ctx context.Context, toField func(int) *dsfetch.ValueString, ids []int, f func(collection string, id int, ids []int) ([]int, error)) ([]int, error) {
	filteredIDs := make(map[string][]int)
	for _, id := range ids {
		contentObjectID, err := toField(id).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting id for element %d: %w", id, err)
		}

		filteredIDs[contentObjectID] = append(filteredIDs[contentObjectID], id)
	}

	allAllowed := make([]int, 0, len(ids))
	for contentObjectID, ids := range filteredIDs {
		collection, objectID, found := strings.Cut(contentObjectID, "/")
		if !found {
			return nil, fmt.Errorf("content object_id has to have exacly one /, got %q", contentObjectID)
		}

		id, err := strconv.Atoi(objectID)
		if err != nil {
			return nil, fmt.Errorf("second part of content_object_id has to be int, got %q", objectID)
		}

		allowed, err := f(collection, id, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for element %s: %w", contentObjectID, err)
		}

		allAllowed = append(allAllowed, allowed...)
	}

	return allAllowed, nil
}

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
