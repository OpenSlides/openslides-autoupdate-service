package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

var collectionMap = map[string]Restricter{
	// ActionWorker{}.Name():               ActionWorker{},
	AgendaItem{}.Name(): AgendaItem{},
	// Assignment{}.Name():                 Assignment{},
	// AssignmentCandidate{}.Name():        AssignmentCandidate{},
	// ListOfSpeakers{}.Name():             ListOfSpeakers{},
	// ChatGroup{}.Name():                  ChatGroup{},
	// ChatMessage{}.Name():                ChatMessage{},
	// Committee{}.Name():                  Committee{},
	// Group{}.Name():                      Group{},
	// Mediafile{}.Name():                  Mediafile{},
	// Meeting{}.Name():                    Meeting{},
	Motion{}.Name(): Motion{},
	// MotionBlock{}.Name():                MotionBlock{},
	// MotionCategory{}.Name():             MotionCategory{},
	// MotionChangeRecommendation{}.Name(): MotionChangeRecommendation{},
	// MotionState{}.Name():                MotionState{},
	// MotionStatuteParagraph{}.Name():     MotionStatuteParagraph{},
	// MotionComment{}.Name():              MotionComment{},
	// MotionCommentSection{}.Name():       MotionCommentSection{},
	// MotionSubmitter{}.Name():            MotionSubmitter{},
	// MotionWorkflow{}.Name():             MotionWorkflow{},
	// Option{}.Name():                     Option{},
	// Organization{}.Name():               Organization{},
	// OrganizationTag{}.Name():            OrganizationTag{},
	// PersonalNote{}.Name():               PersonalNote{},
	// PointOfOrderCategory{}.Name():       PointOfOrderCategory{},
	// Poll{}.Name():                       Poll{},
	// PollCandidate{}.Name():              PollCandidate{},
	// PollCandidateList{}.Name():          PollCandidateList{},
	// Projection{}.Name():                 Projection{},
	// Projector{}.Name():                  Projector{},
	// ProjectorCountdown{}.Name():         ProjectorCountdown{},
	// ProjectorMessage{}.Name():           ProjectorMessage{},
	// Speaker{}.Name():                    Speaker{},
	// Tag{}.Name():                        Tag{},
	// Theme{}.Name():                      Theme{},
	// Topic{}.Name():                      Topic{},
	// User{}.Name():                       User{},
	// Vote{}.Name():                       Vote{},
}

// Collection returns the restricter for a collection
func Collection(ctx context.Context, collection string) Restricter {
	r, ok := collectionMap[collection]
	if !ok {
		return Unknown{collection}
	}

	// TODO: Fixme for superadmin. It needs the restrict superadmin method
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

// Tuple is a pair of modekey and attribute.func
//
// They can be merged into a map later.
type Tuple struct {
	Key   dskey.Key
	Value attribute.Func
}

// TupleFromModeKeys creates list of Tuples for a list of ids from the same
// collection mode.
func TupleFromModeKeys(r Restricter, ids []int, mode string, attr attribute.Func) []Tuple {
	result := make([]Tuple, len(ids))
	for i, id := range ids {
		result[i] = Tuple{modeKey(r, id, mode), attr}
	}
	return result
}

// FieldRestricter is a function to restrict fields of a collection.
type FieldRestricter func(ctx context.Context, fetcher *dsfetch.Fetch, ids []int) ([]Tuple, error)

func meetingPerm(ctx context.Context, fetcher *dsfetch.Fetch, r Restricter, mode string, ids []int, permission perm.TPermission) ([]Tuple, error) {
	return byMeeting(ctx, fetcher, r, ids, func(meetingID int, ids []int) ([]Tuple, error) {
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		attr := attribute.FuncOr(
			attribute.FuncGlobalLevel(perm.OMLSuperadmin),
			attribute.FuncInGroup(groupMap[permission]),
		)

		result := make([]Tuple, len(ids))
		for i, id := range ids {
			result[i] = Tuple{
				Key:   modeKey(r, id, mode),
				Value: attr,
			}
		}
		return result, nil
	})
}

func byMeeting(ctx context.Context, fetcher *dsfetch.Fetch, r Restricter, ids []int, fn func(meetingID int, ids []int) ([]Tuple, error)) ([]Tuple, error) {
	meetingToIDs := make(map[int][]int)
	for _, id := range ids {
		meetingID, hasMeeting, err := r.MeetingID(ctx, fetcher, id)
		if err != nil {
			return nil, fmt.Errorf("getting meeting id of element %d: %w", id, err)
		}

		if !hasMeeting || meetingID == 0 {
			return nil, fmt.Errorf("element with id %d has no meeting", id)
		}

		meetingToIDs[meetingID] = append(meetingToIDs[meetingID], id)
	}

	resultList := make([][]Tuple, 0, len(meetingToIDs))
	var count int
	for meetingID, ids := range meetingToIDs {
		result, err := fn(meetingID, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for meeting %d: %w", meetingID, err)
		}
		resultList = append(resultList, result)
		count += len(result)
	}

	combined := make([]Tuple, 0, count)
	for _, result := range resultList {
		combined = append(combined, result...)
	}

	return combined, nil
}

// TODO: byRelationField and byMeeting are very simular. Maybe write an abstract
// function that takes a filter function.
func byRelationField(ctx context.Context, fetcher *dsfetch.Fetch, r Restricter, toField func(int) *dsfetch.ValueInt, ids []int, fn func(relationID int, ids []int) ([]Tuple, error)) ([]Tuple, error) {
	filteredIDs := make(map[int][]int)
	for _, id := range ids {
		fieldID, err := toField(id).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting id for element %d: %w", id, err)
		}
		if fieldID == 0 {
			return nil, fmt.Errorf("element with id %d has no relation", id)
		}
		filteredIDs[fieldID] = append(filteredIDs[fieldID], id)
	}

	resultList := make([][]Tuple, 0, len(filteredIDs))
	var count int
	for meetingID, ids := range filteredIDs {
		result, err := fn(meetingID, ids)
		if err != nil {
			return nil, fmt.Errorf("restricting for meeting %d: %w", meetingID, err)
		}
		resultList = append(resultList, result)
		count += len(result)
	}

	combined := make([]Tuple, 0, count)
	for _, result := range resultList {
		combined = append(combined, result...)
	}

	return combined, nil
}

func never(ctx context.Context, ds *dsfetch.Fetch, agendaIDs []int) ([]Tuple, error) {
	return nil, nil
}

func modeKey(r Restricter, id int, mode string) dskey.Key {
	return dskey.Key{Collection: r.Name(), ID: id, Field: mode}
}
