package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

// CollectionMap is an index from collection name to its restricter.
var CollectionMap = map[string]Restricter{
	// ActionWorker{}.Name():               ActionWorker{},
	AgendaItem{}.Name():          AgendaItem{},
	Assignment{}.Name():          Assignment{},
	AssignmentCandidate{}.Name(): AssignmentCandidate{},
	ListOfSpeakers{}.Name():      ListOfSpeakers{},
	// ChatGroup{}.Name():                  ChatGroup{},
	// ChatMessage{}.Name():                ChatMessage{},
	// Committee{}.Name():                  Committee{},
	Group{}.Name():       Group{},
	Mediafile{}.Name():   Mediafile{},
	Meeting{}.Name():     Meeting{},
	MeetingUser{}.Name(): MeetingUser{},
	Motion{}.Name():      Motion{},
	// MotionBlock{}.Name():                MotionBlock{},
	// MotionCategory{}.Name():             MotionCategory{},
	// MotionChangeRecommendation{}.Name(): MotionChangeRecommendation{},
	MotionState{}.Name(): MotionState{},
	// MotionStatuteParagraph{}.Name():     MotionStatuteParagraph{},
	MotionComment{}.Name():        MotionComment{},
	MotionCommentSection{}.Name(): MotionCommentSection{},
	MotionSubmitter{}.Name():      MotionSubmitter{},
	MotionWorkflow{}.Name():       MotionWorkflow{},
	Option{}.Name():               Option{},
	Organization{}.Name():         Organization{},
	// OrganizationTag{}.Name():            OrganizationTag{},
	// PersonalNote{}.Name():               PersonalNote{},
	PointOfOrderCategory{}.Name(): PointOfOrderCategory{},
	Poll{}.Name():                 Poll{},
	PollCandidate{}.Name():        PollCandidate{},
	PollCandidateList{}.Name():    PollCandidateList{},
	// Projection{}.Name():                 Projection{},
	Projector{}.Name(): Projector{},
	// ProjectorCountdown{}.Name():         ProjectorCountdown{},
	// ProjectorMessage{}.Name():           ProjectorMessage{},
	// Speaker{}.Name():                    Speaker{},
	Tag{}.Name():   Tag{},
	Theme{}.Name(): Theme{},
	Topic{}.Name(): Topic{},
	User{}.Name():  User{},
	// Vote{}.Name():                       Vote{},
}

// FromName returns a restricter for a collection from its name.
func FromName(ctx context.Context, name string) Restricter {
	r, ok := CollectionMap[name]
	if !ok {
		return Unknown{name}
	}

	// TODO: Fixme for superadmin. It needs the restrict superadmin method
	return withRestrictCache(ctx, r)
}

// Collection returns the restricter for a collection
func Collection(ctx context.Context, collection Restricter) Restricter {
	r, ok := CollectionMap[collection.Name()]
	if !ok {
		panic(fmt.Sprintf("collection %s is not in collection.collectionMap", collection.Name()))
	}

	// TODO: Fixme for superadmin. It needs the restrict superadmin method
	return withRestrictCache(ctx, r)
}

// Unknown is a collection that does not exist in the models.yml
type Unknown struct {
	name string
}

// Modes on an unknown field can not be seen.
func (u Unknown) Modes(mode string) FieldRestricter {
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

// Allways is a restricter func that just returns true.
func Allways(ctx context.Context, fetcher *dsfetch.Fetch, ids []int) ([]attribute.Func, error) {
	return attributeFuncList(ids, attribute.FuncAllow), nil
}

func loggedIn(ctx context.Context, fetcher *dsfetch.Fetch, ids []int) ([]attribute.Func, error) {
	return attributeFuncList(ids, attribute.FuncLoggedIn), nil
}

func never(ctx context.Context, ds *dsfetch.Fetch, ids []int) ([]attribute.Func, error) {
	result := make([]attribute.Func, len(ids))
	for i := range ids {
		result[i] = attribute.FuncNotAllowed
	}
	return result, nil
}

type restrictCache struct {
	cache map[dskey.Key]attribute.Func
	Restricter
}

type contextKeyType string

var contextKey contextKeyType = "restrict cache"

// ContextWithRestrictCache returns a context with restrict cache.
func ContextWithRestrictCache(ctx context.Context) context.Context {
	return context.WithValue(ctx, contextKey, make(map[dskey.Key]attribute.Func))
}

func withRestrictCache(ctx context.Context, sub Restricter) Restricter {
	v := ctx.Value(contextKey)
	if v == nil {
		panic("collection cache not initialized")
	}

	cache, ok := v.(map[dskey.Key]attribute.Func)
	if !ok {
		panic("collection cache is broken")
	}

	return &restrictCache{
		cache:      cache,
		Restricter: sub,
	}
}

func (r *restrictCache) Modes(mode string) FieldRestricter {
	return func(ctx context.Context, fetcher *dsfetch.Fetch, ids []int) ([]attribute.Func, error) {
		notFoundIDs := make([]int, len(ids))
		attrList := make([]attribute.Func, len(ids))
		foundAll := true
		for i, id := range ids {
			key, err := dskey.FromParts(r.Name(), id, mode)
			if err != nil {
				return nil, err
			}

			attrFunc, found := r.cache[key]
			if !found {
				notFoundIDs[i] = id
				foundAll = false
				continue
			}

			attrList[i] = attrFunc
		}

		if foundAll {
			return attrList, nil
		}

		newAttrList, err := r.Restricter.Modes(mode)(ctx, fetcher, notFoundIDs)
		if err != nil {
			idsString := "with same ids"
			if len(notFoundIDs) != len(ids) {
				idsString = fmt.Sprintf("with ids %v", notFoundIDs)
			}
			return nil, fmt.Errorf("calling restricter %s: %w", idsString, err)
		}

		for i := range ids {
			if newAttrList[i] != nil {
				attrList[i] = newAttrList[i]
			}
		}

		return attrList, nil
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

// FieldRestricter is a function to restrict fields of a collection.
//
// The ids can contain 0. In this case, the coresponding attribute.Func has to be nil
type FieldRestricter func(ctx context.Context, fetcher *dsfetch.Fetch, ids []int) ([]attribute.Func, error)

func meetingPerm(ctx context.Context, fetcher *dsfetch.Fetch, r Restricter, ids []int, permission perm.TPermission) ([]attribute.Func, error) {
	return byMeeting(ctx, fetcher, r, ids, func(meetingID int, ids []int) ([]attribute.Func, error) {
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		attr := attribute.FuncOr(
			attribute.FuncGlobalLevel(perm.OMLSuperadmin),
			attribute.FuncInGroup(groupMap[permission]),
		)

		result := make([]attribute.Func, len(ids))
		for i, id := range ids {
			if id == 0 {
				continue
			}
			result[i] = attr
		}
		return result, nil
	})
}

func byMeeting(ctx context.Context, fetcher *dsfetch.Fetch, r Restricter, ids []int, fn func(meetingID int, ids []int) ([]attribute.Func, error)) ([]attribute.Func, error) {
	meetingToIDs := make(map[int][]int)

	for i, id := range ids {
		if id == 0 {
			continue
		}

		meetingID, hasMeeting, err := r.MeetingID(ctx, fetcher, id)
		if err != nil {
			return nil, fmt.Errorf("getting meeting id of element %d: %w", id, err)
		}

		if !hasMeeting || meetingID == 0 {
			return nil, fmt.Errorf("element with id %d has no meeting", id)
		}

		if meetingToIDs[meetingID] == nil {
			meetingToIDs[meetingID] = make([]int, len(ids))
		}

		meetingToIDs[meetingID][i] = id
	}

	resultList := make([]attribute.Func, len(ids))
	for meetingID, ids := range meetingToIDs {
		attrList, err := fn(meetingID, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for meeting %d: %w", meetingID, err)
		}

		for i, attr := range attrList {
			if attr == nil {
				continue
			}
			resultList[i] = attr
		}
	}
	return resultList, nil
}

// TODO: byRelationField and byMeeting are very simular. Maybe write an abstract
// function that takes a filter function.
func byRelationField(ctx context.Context, toField func(int) *dsfetch.ValueInt, ids []int, fn func(relationID int, ids []int) ([]attribute.Func, error)) ([]attribute.Func, error) {
	filteredIDs := make(map[int][]int)
	for i, id := range ids {
		fieldID, err := toField(id).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting id for element %d: %w", id, err)
		}
		if fieldID == 0 {
			return nil, fmt.Errorf("element with id %d has no relation", id)
		}

		if filteredIDs[fieldID] == nil {
			filteredIDs[fieldID] = make([]int, len(ids))
		}

		filteredIDs[fieldID][i] = id
	}

	resultList := make([]attribute.Func, len(ids))
	for fieldID, ids := range filteredIDs {
		attrList, err := fn(fieldID, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for field %d: %w", fieldID, err)
		}

		for i, attr := range attrList {
			if attr == nil {
				continue
			}
			resultList[i] = attr
		}
	}

	return resultList, nil
}

func canSeeRelatedCollection(ctx context.Context, fetcher *dsfetch.Fetch, toField func(int) *dsfetch.ValueInt, mode FieldRestricter, ids []int) ([]attribute.Func, error) {
	relationIDs := make([]int, len(ids))
	for i, id := range ids {
		if id == 0 {
			continue
		}
		toField(id).Lazy(&relationIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching related objects: %w", err)
	}

	return mode(ctx, fetcher, relationIDs)
}

func attributeFuncList(ids []int, attr attribute.Func) []attribute.Func {
	result := make([]attribute.Func, len(ids))
	for i, id := range ids {
		if id == 0 {
			continue
		}
		result[i] = attr
	}

	return result
}
