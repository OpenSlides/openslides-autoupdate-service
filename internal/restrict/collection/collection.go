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

type restrictCache struct {
	cache map[dskey.Key]bool
	Restricter
}

type contextKeyType string

var contextKey contextKeyType = "restrict cache"

// ContextWithRestrictCache returns a context with restrict cache.
func ContextWithRestrictCache(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey, make(map[dskey.Key]bool))
}

func withRestrictCache(ctx context.Context, sub Restricter) Restricter {
	v := ctx.Value(contextKey)
	if v == nil {
		return sub
	}

	cache, ok := v.(map[dskey.Key]bool)
	if !ok {
		return sub
	}

	return &restrictCache{
		cache:      cache,
		Restricter: sub,
	}
}

func (r *restrictCache) Modes(mode string) FieldRestricter {
	return func(ctx context.Context, ds *dsfetch.Fetch, ids ...int) ([]int, error) {
		notFound := make([]int, 0, len(ids))
		cachedAllowedIDs := make([]int, 0, len(ids))
		for _, id := range ids {
			key, err := dskey.FromParts(r.Name(), id, mode)
			if err != nil {
				return nil, err
			}

			allowed, found := r.cache[key]
			if !found {
				notFound = append(notFound, id)
				continue
			}

			if allowed {
				cachedAllowedIDs = append(cachedAllowedIDs, id)
			}
		}

		if len(notFound) == 0 {
			return cachedAllowedIDs, nil
		}

		newAllowedIDs, err := r.Restricter.Modes(mode)(ctx, ds, notFound...)
		if err != nil {
			return nil, fmt.Errorf("calling restricter: %w", err)
		}

		// Add all not Found keys to the cache as not allowed.
		for _, id := range notFound {
			key, err := dskey.FromParts(r.Name(), id, mode)
			if err != nil {
				return nil, err
			}
			r.cache[key] = false
		}

		// Set all new allowed ids to the cache as true.
		for _, id := range newAllowedIDs {
			key, err := dskey.FromParts(r.Name(), id, mode)
			if err != nil {
				return nil, err
			}
			r.cache[key] = true
		}

		return append(cachedAllowedIDs, newAllowedIDs...), nil
	}
}

func (r *restrictCache) SuperAdmin(mode string) FieldRestricter {
	type superRestricter interface {
		SuperAdmin(mode string) FieldRestricter
	}

	if sr, ok := r.Restricter.(superRestricter); ok {
		return sr.SuperAdmin(mode)
	}

	return nil
}

var collectionMap = map[string]Restricter{
	ActionWorker{}.Name():               ActionWorker{},
	AgendaItem{}.Name():                 AgendaItem{},
	Assignment{}.Name():                 Assignment{},
	AssignmentCandidate{}.Name():        AssignmentCandidate{},
	ListOfSpeakers{}.Name():             ListOfSpeakers{},
	ChatGroup{}.Name():                  ChatGroup{},
	ChatMessage{}.Name():                ChatMessage{},
	Committee{}.Name():                  Committee{},
	Group{}.Name():                      Group{},
	Mediafile{}.Name():                  Mediafile{},
	Meeting{}.Name():                    Meeting{},
	MeetingUser{}.Name():                MeetingUser{},
	Motion{}.Name():                     Motion{},
	MotionBlock{}.Name():                MotionBlock{},
	MotionCategory{}.Name():             MotionCategory{},
	MotionChangeRecommendation{}.Name(): MotionChangeRecommendation{},
	MotionState{}.Name():                MotionState{},
	MotionStatuteParagraph{}.Name():     MotionStatuteParagraph{},
	MotionComment{}.Name():              MotionComment{},
	MotionCommentSection{}.Name():       MotionCommentSection{},
	MotionSubmitter{}.Name():            MotionSubmitter{},
	MotionWorkflow{}.Name():             MotionWorkflow{},
	Option{}.Name():                     Option{},
	Organization{}.Name():               Organization{},
	OrganizationTag{}.Name():            OrganizationTag{},
	PersonalNote{}.Name():               PersonalNote{},
	PointOfOrderCategory{}.Name():       PointOfOrderCategory{},
	Poll{}.Name():                       Poll{},
	PollCandidate{}.Name():              PollCandidate{},
	PollCandidateList{}.Name():          PollCandidateList{},
	Projection{}.Name():                 Projection{},
	Projector{}.Name():                  Projector{},
	ProjectorCountdown{}.Name():         ProjectorCountdown{},
	ProjectorMessage{}.Name():           ProjectorMessage{},
	Speaker{}.Name():                    Speaker{},
	Tag{}.Name():                        Tag{},
	Theme{}.Name():                      Theme{},
	Topic{}.Name():                      Topic{},
	User{}.Name():                       User{},
	Vote{}.Name():                       Vote{},
}

// Collection returns the restricter for a collection
func Collection(ctx context.Context, collection string) Restricter {
	r, ok := collectionMap[collection]
	if !ok {
		return Unknown{collection}
	}

	return withRestrictCache(ctx, r)
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
