package restrict

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
	"github.com/OpenSlides/openslides-go/datastore/dskey"
	"github.com/OpenSlides/openslides-go/datastore/flow"
	"github.com/OpenSlides/openslides-go/fastjson"
	"github.com/OpenSlides/openslides-go/metagen"
	"github.com/OpenSlides/openslides-go/oserror"
	"github.com/OpenSlides/openslides-go/perm"
	"github.com/OpenSlides/openslides-go/set"
)

// Middleware can be used as a flow.Getter that restrict the data for a
// user.
//
// It also initializes a ctx that has to be used in the future getter calls.
func Middleware(ctx context.Context, getter flow.Getter, uid int) (context.Context, flow.Getter) {
	ctx = contextWithCache(ctx, getter, uid)
	return ctx, restricter{
		getter: getter,
		uid:    uid,
	}
}

// contextWithCache adds some restrictor caches to the context.
func contextWithCache(ctx context.Context, getter flow.Getter, uid int) context.Context {
	ctx = collection.ContextWithRestrictCache(ctx)
	ctx = perm.ContextWithPermissionCache(ctx, getter, uid)
	return ctx
}

type restricter struct {
	getter flow.Getter
	uid    int
}

// Get returns restricted data.
func (r restricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	data, err := r.getter.Get(ctx, keys...)
	if err != nil {
		return nil, fmt.Errorf("getting data: %w", err)
	}

	start := time.Now()
	times, err := restrict(ctx, r.getter, data)
	if err != nil {
		return nil, fmt.Errorf("restricting data: %w", err)
	}

	duration := time.Since(start)

	if times != nil && (duration > slowCalls || oserror.HasTagFromContext(ctx, "profile_restrict")) {
		body, ok := oserror.BodyFromContext(ctx)
		if !ok {
			body = "unknown body, probably simple request"
		}
		profile(body, duration, times)
	}

	return data, nil
}

// restrict changes the keys and values in data for the user with the given user
// id.
func restrict(ctx context.Context, getter flow.Getter, data map[dskey.Key][]byte) (map[string]timeCount, error) {
	ds := dsfetch.New(getter)

	// Get all required collections with there ids.
	restrictModeIDs := make(map[collection.CM]set.Set[int])
	for key := range data {
		if data[key] == nil {
			continue
		}

		if err := groupKeysByCollection(key, data[key], restrictModeIDs); err != nil {
			return nil, fmt.Errorf("grouping keys by collection: %w", err)
		}
	}

	if len(restrictModeIDs) == 0 {
		return nil, nil
	}

	// Call restrict Mode function for each collection.
	times := make(map[string]timeCount, len(restrictModeIDs))
	orderedCMs := sortRestrictModeIDs(restrictModeIDs)
	allowedMods := make(map[collection.CM]set.Set[int])
	for _, cm := range orderedCMs {
		ids := restrictModeIDs[cm].List()
		start := time.Now()

		modeFunc, err := restrictModefunc(ctx, cm.Collection, cm.Mode)
		if err != nil {
			return nil, fmt.Errorf("getting restiction mode for %s/%s: %w", cm.Collection, cm.Mode, err)
		}

		allowedIDs, err := modeFunc(ctx, ds, ids...)
		if err != nil {
			return nil, fmt.Errorf("calling collection %s modefunc %s with ids %v: %w", cm.Collection, cm.Mode, ids, err)
		}
		allowedMods[cm] = set.New(allowedIDs...)

		duration := time.Since(start)
		times[cm.Collection+"/"+cm.Mode] = timeCount{time: duration, count: len(ids)}
	}

	// Remove restricted keys.
	for key := range data {
		if data[key] == nil {
			continue
		}

		restrictionMode, err := restrictModeName(key.Collection(), key.Field())
		if err != nil {
			return nil, fmt.Errorf("getting restriction Mode for %s: %w", key, err)
		}

		cm := collection.CM{Collection: key.Collection(), Mode: restrictionMode}
		if !allowedMods[cm].Has(key.ID()) {
			data[key] = nil
			continue
		}

		newValue, ok, err := manipulateRelations(key, data[key], allowedMods)
		if err != nil {
			return nil, fmt.Errorf("new value for relation key %s: %w", key, err)
		}

		if ok {
			data[key] = newValue
		}
	}

	return times, nil
}

// groupKeysByCollection groups all the keys in data by there collection.
func groupKeysByCollection(key dskey.Key, value []byte, restrictModeIDs map[collection.CM]set.Set[int]) error {
	restrictionMode, err := restrictModeName(key.Collection(), key.Field())
	if err != nil {
		return fmt.Errorf("getting restriction Mode for %s: %w", key, err)
	}

	cm := collection.CM{Collection: key.Collection(), Mode: restrictionMode}
	if restrictModeIDs[cm].IsNotInitialized() {
		restrictModeIDs[cm] = set.New[int]()
	}
	restrictModeIDs[cm].Add(key.ID())

	if err := addRelationToRestrictModeIDs(key, value, restrictModeIDs); err != nil {
		return fmt.Errorf("check %s for relation: %w", key, err)
	}

	return nil
}

func addRelationToRestrictModeIDs(key dskey.Key, value []byte, restrictModeIDs map[collection.CM]set.Set[int]) error {
	collectionField := key.CollectionField()

	cm, id, ok, err := isRelation(collectionField, value)
	if err != nil {
		return fmt.Errorf("checking for relation: %w", err)
	}

	if ok {
		if restrictModeIDs[cm].IsNotInitialized() {
			restrictModeIDs[cm] = set.New[int]()
		}
		restrictModeIDs[cm].Add(id)
		return nil
	}

	cm, ids, ok, err := isRelationList(collectionField, value)
	if err != nil {
		return fmt.Errorf("checking for relation-list: %w", err)
	}

	if ok {
		if restrictModeIDs[cm].IsNotInitialized() {
			restrictModeIDs[cm] = set.New[int]()
		}
		restrictModeIDs[cm].Add(ids...)
		return nil
	}

	cm, id, ok, err = isGenericRelation(collectionField, value)
	if err != nil {
		return fmt.Errorf("checking for generic-relation: %w", err)
	}

	if ok {
		if restrictModeIDs[cm].IsNotInitialized() {
			restrictModeIDs[cm] = set.New[int]()
		}
		restrictModeIDs[cm].Add(id)
		return nil
	}

	mcm, _, ok, err := isGenericRelationList(collectionField, value)
	if err != nil {
		return fmt.Errorf("checking for generic-relation-list: %w", err)
	}

	if ok {
		for _, cmID := range mcm {
			if restrictModeIDs[cmID.cm].IsNotInitialized() {
				restrictModeIDs[cmID.cm] = set.New[int]()
			}
			restrictModeIDs[cmID.cm].Add(cmID.id)
		}
	}

	return nil
}

// manipulateRelations changes the value of relation fields.
//
// The first return value is the new value. The second is, if the value was
// manipulated.q
func manipulateRelations(key dskey.Key, value []byte, allowedRestrictions map[collection.CM]set.Set[int]) ([]byte, bool, error) {
	collectionField := key.CollectionField()

	cm, id, ok, err := isRelation(collectionField, value)
	if err != nil {
		return nil, false, fmt.Errorf("checking %s for relation: %w", key, err)
	}

	if ok {
		return nil, !allowedRestrictions[cm].Has(id), nil
	}

	cm, ids, ok, err := isRelationList(collectionField, value)
	if err != nil {
		return nil, false, fmt.Errorf("checking %s for relation-list: %w", key, err)
	}

	if ok {
		allowed := make([]int, 0, len(ids))
		for _, id := range ids {
			if allowedRestrictions[cm].Has(id) {
				allowed = append(allowed, id)
			}
		}

		if len(allowed) != len(ids) {
			newValue, err := json.Marshal(allowed)
			if err != nil {
				return nil, false, fmt.Errorf("marshal new value for key %s: %w", key, err)
			}
			return newValue, true, nil
		}
		return nil, false, nil
	}

	cm, id, ok, err = isGenericRelation(collectionField, value)
	if err != nil {
		return nil, false, fmt.Errorf("checking %s for generic-relation: %w", key, err)
	}

	if ok {
		return nil, !allowedRestrictions[cm].Has(id), nil
	}

	mcm, genericIDs, ok, err := isGenericRelationList(collectionField, value)
	if err != nil {
		return nil, false, fmt.Errorf("checking %s for generic-relation-list: %w", key, err)
	}

	if ok {
		allowed := make([]string, 0, len(genericIDs))
		for genericID, cmID := range mcm {
			if allowedRestrictions[cmID.cm].Has(cmID.id) {
				allowed = append(allowed, genericID)
			}
		}

		if len(allowed) != len(genericIDs) {
			newValue, err := json.Marshal(allowed)
			if err != nil {
				return nil, false, fmt.Errorf("marshal new value for key %s: %w", key, err)
			}
			return newValue, true, nil
		}

		return nil, false, nil
	}

	return nil, false, nil
}

func isRelation(collectionField string, value []byte) (collection.CM, int, bool, error) {
	toCollectionfield, ok := metagen.RelationFields[collectionField]
	if !ok {
		return collection.CM{}, 0, false, nil
	}

	id, err := fastjson.DecodeInt(value)
	if err != nil {
		return collection.CM{}, 0, false, fmt.Errorf("decoding %q (`%s`): %w", collectionField, value, err)
	}

	coll, field, _ := strings.Cut(toCollectionfield, "/")
	fieldMode, err := restrictModeName(coll, field)
	if err != nil {
		return collection.CM{}, 0, false, fmt.Errorf("building relation field field mode: %w", err)
	}

	return collection.CM{Collection: coll, Mode: fieldMode}, id, true, nil
}

func isRelationList(keyPrefix string, value []byte) (collection.CM, []int, bool, error) {
	toCollectionfield, ok := metagen.RelationListFields[keyPrefix]
	if !ok {
		return collection.CM{}, nil, false, nil
	}

	ids, err := fastjson.DecodeIntList(value)
	if err != nil {
		return collection.CM{}, nil, false, fmt.Errorf("decoding value (size: %d): %w", len(value), err)
	}

	coll, field, _ := strings.Cut(toCollectionfield, "/")
	fieldMode, err := restrictModeName(coll, field)
	if err != nil {
		return collection.CM{}, nil, false, fmt.Errorf("building relation field field mode: %w", err)
	}

	return collection.CM{Collection: coll, Mode: fieldMode}, ids, true, nil
}

func isGenericRelation(keyPrefix string, value []byte) (collection.CM, int, bool, error) {
	toCollectionfield, ok := metagen.GenericRelationFields[keyPrefix]
	if !ok {
		return collection.CM{}, 0, false, nil
	}

	var genericID string
	if err := json.Unmarshal(value, &genericID); err != nil {
		return collection.CM{}, 0, false, fmt.Errorf("decoding %q: %w", keyPrefix, err)
	}

	cm, id, err := genericKeyToCollectionMode(genericID, toCollectionfield)
	if err != nil {
		return collection.CM{}, 0, false, fmt.Errorf("parsing generic key: %w", err)
	}

	return cm, id, true, nil
}

type collectionModeID struct {
	cm collection.CM
	id int
}

func isGenericRelationList(keyPrefix string, value []byte) (map[string]collectionModeID, []string, bool, error) {
	toCollectionfield, ok := metagen.GenericRelationListFields[keyPrefix]
	if !ok {
		return nil, nil, false, nil
	}

	var genericIDs []string
	if err := json.Unmarshal(value, &genericIDs); err != nil {
		return nil, nil, false, fmt.Errorf("decoding %q: %w", keyPrefix, err)
	}

	mcm := make(map[string]collectionModeID, len(genericIDs))
	for _, genericID := range genericIDs {
		cm, id, err := genericKeyToCollectionMode(genericID, toCollectionfield)
		if err != nil {
			return nil, nil, false, fmt.Errorf("parsing generic key: %w", err)
		}
		mcm[genericID] = collectionModeID{cm, id}
	}

	return mcm, genericIDs, true, nil
}

// genericKeyToCollectionMode calls f for each collection mode.
func genericKeyToCollectionMode(genericID string, toCollectionFieldMap map[string]string) (collection.CM, int, error) {
	coll, rawID, found := strings.Cut(genericID, "/")
	if !found {
		// TODO LAST ERROR
		return collection.CM{}, 0, fmt.Errorf("invalid generic relation: %s", genericID)
	}

	id, err := strconv.Atoi(rawID)
	if err != nil {
		// TODO LAST ERROR
		return collection.CM{}, 0, fmt.Errorf("invalid generic relation, no id: %s", genericID)
	}

	toField := toCollectionFieldMap[coll]
	if toField == "" {
		// TODO LAST ERROR
		return collection.CM{}, 0, fmt.Errorf("unknown generic relation: %s", coll)
	}

	fieldMode, err := restrictModeName(coll, toField)
	if err != nil {
		return collection.CM{}, 0, fmt.Errorf("building generic relation field mode: %w", err)
	}

	return collection.CM{Collection: coll, Mode: fieldMode}, id, nil
}

// restrictModeName returns the restriction mode for a collection and field.
//
// This is a string like "A" or "B" or any other name of a restriction mode.
func restrictModeName(collection, field string) (string, error) {
	fieldMode, ok := metagen.RestrictionModes[collection+"/"+field]
	if !ok {
		// TODO LAST ERROR
		return "", fmt.Errorf("fqfield %q is unknown, maybe run go generate ./... in openlsides-go to fetch all fields from the models.yml", collection+"/"+field)
	}
	return fieldMode, nil
}

// restrictModefunc returns the field restricter function to use.
func restrictModefunc(ctx context.Context, collectionName, fieldMode string) (collection.FieldRestricter, error) {
	restricter := collection.Collection(ctx, collectionName)
	if _, ok := restricter.(collection.Unknown); ok {
		// TODO LAST ERROR
		return nil, fmt.Errorf("collection %q is not implemented, maybe run go generate ./... in openslides-go to fetch all fields from the models.yml", collectionName)
	}

	modefunc := restricter.Modes(fieldMode)
	if modefunc == nil {
		// TODO LAST ERROR
		return nil, fmt.Errorf("mode %q of models %q is not implemented", fieldMode, collectionName)
	}

	return modefunc, nil
}

func sortRestrictModeIDs(data map[collection.CM]set.Set[int]) []collection.CM {
	keys := make([]collection.CM, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(a, b int) bool {
		var aid, bid int
		if id, ok := collectionOrder[keys[a].String()]; ok {
			aid = id
		} else if id, ok := collectionOrder[keys[a].Collection]; ok {
			aid = id
		}

		if id, ok := collectionOrder[keys[b].String()]; ok {
			bid = id
		} else if id, ok := collectionOrder[keys[b].Collection]; ok {
			bid = id
		}

		return aid < bid
	})

	return keys
}

var collectionOrder = map[string]int{
	"agenda_item":                      1,
	"assignment":                       2,
	"assignment_candidate":             3,
	"chat_group":                       4,
	"chat_message":                     5,
	"committee":                        6,
	"meeting/A":                        7,
	"meeting/E":                        8,
	"point_of_order_category":          8,
	"group":                            9,
	"projector":                        10,
	"projection":                       11,
	"projector_countdown":              12,
	"projector_message":                13,
	"meeting_mediafile":                14,
	"mediafile":                        15,
	"tag":                              16,
	"motion/C":                         17,
	"motion/B":                         18,
	"motion_block":                     19,
	"motion_category":                  20,
	"motion_change_recommendation":     21,
	"motion_comment_section":           22,
	"motion_comment":                   23,
	"motion_state":                     24,
	"motion_submitter":                 26,
	"motion_workflow":                  27,
	"poll":                             28,
	"option":                           29,
	"poll_candidate_list":              30,
	"poll_candidate":                   41,
	"vote":                             42,
	"organization":                     43,
	"organization_tag":                 44,
	"personal_note":                    45,
	"theme":                            46,
	"topic":                            47,
	"list_of_speakers":                 48,
	"speaker":                          49,
	"user":                             50,
	"meeting_user":                     51,
	"action_worker":                    52,
	"import_preview":                   53,
	"structure_level":                  54,
	"structure_level_list_of_speakers": 55,
	"motion_working_group_speaker":     56,
	"motion_editor":                    57,
	"gender":                           58,
}
