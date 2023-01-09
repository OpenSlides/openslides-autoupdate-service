package restrict

//go:generate  sh -c "go run gen_field_def/main.go > field_def.go"

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Restricter holds attributes for each restriction mode.
type Restricter struct {
	mu           sync.RWMutex
	attributeMap collection.AttributeMap
}

// New initializes the restricter.
func New() *Restricter {
	return &Restricter{
		attributeMap: collection.NewAttributeMap(),
	}
}

// InsertFields adds new fields.
func (r *Restricter) InsertFields(ctx context.Context, ds datastore.Getter, values map[dskey.Key][]byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	restrictModeIDs, err := groupKeysByCollectionMode(values)
	if err != nil {
		return fmt.Errorf("group keys: %w", err)
	}

	if err := r.prepare(ctx, ds, restrictModeIDs); err != nil {
		return fmt.Errorf("prepare: %w", err)
	}

	return nil
}

// UpdateFields updates the attribute fields.
func (r *Restricter) UpdateFields(ctx context.Context, ds datastore.Getter, values map[dskey.Key][]byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: Handle deletion.
	// TODO: Only update on hot keys
	restrictModeIDs := r.attributeMap.RestrictModeIDs()

	if err := r.prepare(ctx, ds, restrictModeIDs); err != nil {
		return fmt.Errorf("prepare: %w", err)
	}

	return nil
}

func (r *Restricter) prepare(ctx context.Context, ds datastore.Getter, restrictModeIDs map[collection.CM]set.Set[int]) error {
	orderedCMs := sortRestrictModeIDs(restrictModeIDs)
	mperms := perm.NewMeetingPermission()
	fetcher := dsfetch.New(ds)

	for _, cm := range orderedCMs {
		ids := restrictModeIDs[cm]

		modeFunc, err := restrictModefunc(cm.Collection, cm.Mode)
		if err != nil {
			return fmt.Errorf("getting modeFunc for %s: %w", cm, err)
		}

		if err := modeFunc(ctx, fetcher, mperms, r.attributeMap, ids.List()...); err != nil {
			return fmt.Errorf("restrict mode %s", cm)
		}
	}

	return nil
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

// groupKeysByCollection groups all the keys in data by there collection.
func groupKeysByCollectionMode(values map[dskey.Key][]byte) (map[collection.CM]set.Set[int], error) {
	restrictModeIDs := make(map[collection.CM]set.Set[int], len(values)) // TODO: the len is proably smaller then values

	for key := range values {
		restrictionMode, err := restrictModeName(key.Collection, key.Field)
		if err != nil {
			return nil, fmt.Errorf("getting restriction Mode for %s: %w", key, err)
		}

		cm := collection.CM{Collection: key.Collection, Mode: restrictionMode}
		if restrictModeIDs[cm].IsNotInitialized() {
			restrictModeIDs[cm] = set.New[int]()
		}
		restrictModeIDs[cm].Add(key.ID)
	}

	return restrictModeIDs, nil
}

// Getter returns a datastore getter that returns restricted data.
func (r *Restricter) Getter(getter datastore.Getter, uid int) datastore.Getter {
	return datastore.GetterFunc(func(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
		r.mu.RLock()
		defer r.mu.RUnlock()

		data, err := getter.Get(ctx, keys...)
		if err != nil {
			return nil, fmt.Errorf("fetching unrestricted data: %w", err)
		}

		fetcher := dsfetch.New(getter)

		globalPermStr, err := fetcher.User_OrganizationManagementLevel(uid).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting orga level from user: %w", err)
		}

		groupIDs, err := userGroups(ctx, fetcher, uid)
		if err != nil {
			return nil, fmt.Errorf("getting group ids: %w", err)
		}

		orgaLevel := perm.OrganizationManagementFromString(globalPermStr)

		for key := range data {
			restrictionMode, err := restrictModeName(key.Collection, key.Field)
			if err != nil {
				return nil, fmt.Errorf("getting restriction Mode for %s: %w", key, err)
			}

			requiredAttr := r.attributeMap.Get(key.Collection, key.ID, restrictionMode)

			if !allowedByAttr(requiredAttr, uid, orgaLevel, groupIDs) {
				data[key] = nil
				continue
			}

			// TODO: relation fields

		}

		return data, nil

	})
}

// userGroups returns all groups from a user.
func userGroups(ctx context.Context, fetcher *dsfetch.Fetch, uid int) ([]int, error) {
	meetingIDs, err := fetcher.User_GroupIDsTmpl(uid).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting meetings of user: %w", err)
	}

	groupIDs := make([][]int, len(meetingIDs))
	for i := 0; i < len(meetingIDs); i++ {
		fetcher.User_GroupIDs(uid, meetingIDs[i]).Lazy(&groupIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("getting groupIDs: %w", err)
	}

	var result []int
	for _, ids := range groupIDs {
		result = append(result, ids...)
	}

	return result, nil
}

func allowedByAttr(requiredAttr *collection.Attributes, uid int, orgaLevel perm.OrganizationManagementLevel, groupIDs []int) bool {
	if requiredAttr.GlobalPermission == 255 {
		return false
	}

	if allowedByGlobalPerm(orgaLevel, uid, requiredAttr.GlobalPermission) {
		return true
	}

	if allowedByGroup(requiredAttr.GroupIDs, groupIDs) {
		if nextAttr := requiredAttr.GroupAnd; nextAttr != nil {
			return allowedByAttr(nextAttr, uid, orgaLevel, groupIDs)
		}
		return true
	}

	if allowedByUser(requiredAttr.UserIDs, uid) {
		return true
	}

	return false
}

// allowedByGlobalPerm tells, if the user has the correct global permission.
func allowedByGlobalPerm(orgaLevel perm.OrganizationManagementLevel, uid int, globalAttr byte) bool {
	return globalAttr == 0 ||
		globalAttr == 254 && uid > 0 ||
		globalAttr == 1 && orgaLevel == perm.OMLSuperadmin ||
		globalAttr == 2 && (orgaLevel == perm.OMLSuperadmin || orgaLevel == perm.OMLCanManageOrganization) ||
		globalAttr == 3 && (orgaLevel == perm.OMLSuperadmin || orgaLevel == perm.OMLCanManageOrganization || orgaLevel == perm.OMLCanManageUsers)
}

func allowedByGroup(requiredGroups set.Set[int], hasGroups []int) bool {
	for _, gid := range hasGroups {
		if requiredGroups.Has(gid) {
			return true
		}
	}

	return false
}

func allowedByUser(requiredIDs set.Set[int], uid int) bool {
	return requiredIDs.Has(uid)
}

// restrictModefunc returns the field restricter function to use.
func restrictModefunc(collectionName, fieldMode string) (collection.FieldRestricter, error) {
	restricter := collection.Collection(collectionName)
	if _, ok := restricter.(collection.Unknown); ok {
		// TODO LAST ERROR
		return nil, fmt.Errorf("collection %q is not implemented, maybe run go generate ./... to fetch all fields from the models.yml", collectionName)
	}

	modefunc := restricter.Modes(fieldMode)
	if modefunc == nil {
		// TODO LAST ERROR
		return nil, fmt.Errorf("mode %q of models %q is not implemented", fieldMode, collectionName)
	}

	return modefunc, nil
}

// restrictModeName returns the restriction mode for a collection and field.
//
// This is a string like "A" or "B" or any other name of a restriction mode.
func restrictModeName(collection, field string) (string, error) {
	fieldMode, ok := restrictionModes[templateKeyPrefix(collection+"/"+field)]
	if !ok {
		// TODO LAST ERROR
		return "", fmt.Errorf("fqfield %q is unknown, maybe run go generate ./... to fetch all fields from the models.yml", collection+"/"+field)
	}
	return fieldMode, nil
}

// templateKeyPrefix returns the index of the field list list. For template fields this is
// the key without the replacement.
func templateKeyPrefix(collectionField string) string {
	i := strings.IndexByte(collectionField, '$')
	if i < 0 || i == len(collectionField)-1 || collectionField[i+1] == '_' {
		// Normal field or $ at the end or $_
		return collectionField
	}

	return collectionField[:i+1]
}

var collectionOrder = map[string]int{
	"agenda_item":                  1,
	"assignment":                   2,
	"assignment_candidate":         3,
	"chat_group":                   4,
	"chat_message":                 5,
	"committee":                    6,
	"meeting":                      7,
	"group":                        8,
	"mediafile":                    9,
	"tag":                          10,
	"motion/C":                     11,
	"motion/B":                     12,
	"motion_block":                 13,
	"motion_category":              14,
	"motion_change_recommendation": 15,
	"motion_comment_section":       16,
	"motion_comment":               17,
	"motion_state":                 18,
	"motion_statute_paragraph":     19,
	"motion_submitter":             20,
	"motion_workflow":              21,
	"poll":                         22,
	"option":                       23,
	"vote":                         24,
	"organization":                 25,
	"organization_tag":             26,
	"personal_note":                27,
	"projection":                   28,
	"projector":                    29,
	"projector_countdown":          30,
	"projector_message":            31,
	"theme":                        32,
	"topic":                        33,
	"list_of_speakers":             34,
	"speaker":                      35,
	"user":                         36,
}

// type userRestricter struct {
// 	getter datastore.Getter
// 	uid    int
// }

// // Get returns restricted data.
// func (r userRestricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
// 	data, err := r.getter.Get(ctx, keys...)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting data: %w", err)
// 	}

// 	start := time.Now()
// 	times, err := restrict(ctx, r.getter, r.uid, data)
// 	if err != nil {
// 		return nil, fmt.Errorf("restricting data: %w", err)
// 	}

// 	duration := time.Since(start)

// 	if times != nil && (duration > slowCalls || oserror.HasTagFromContext(ctx, "profile_restrict")) {
// 		body, ok := oserror.BodyFromContext(ctx)
// 		if !ok {
// 			body = "unknown body, probably simple request"
// 		}
// 		profile(body, duration, times)
// 	}

// 	return data, nil
// }

// // restrict changes the keys and values in data for the user with the given user
// // id.
// func restrict(ctx context.Context, getter datastore.Getter, uid int, data map[dskey.Key][]byte) (map[string]timeCount, error) {
// 	ds := dsfetch.New(getter)

// 	isSuperAdmin, err := perm.HasOrganizationManagementLevel(ctx, ds, uid, perm.OMLSuperadmin)
// 	if err != nil {
// 		var errDoesNotExist dsfetch.DoesNotExistError
// 		if errors.As(err, &errDoesNotExist) || dskey.Key(errDoesNotExist).Collection == "user" {
// 			// TODO LAST ERROR
// 			return nil, fmt.Errorf("request user %d does not exist", uid)
// 		}
// 		return nil, fmt.Errorf("checking for superadmin: %w", err)
// 	}

// 	if isSuperAdmin {
// 		if err := restrictSuperAdmin(ctx, getter, uid, data); err != nil {
// 			return nil, fmt.Errorf("restrict as superadmin: %w", err)
// 		}
// 		return nil, nil
// 	}

// 	// Get all required collections with there ids.
// 	restrictModeIDs := make(map[collection.CM]set.Set[int])
// 	for key := range data {
// 		if data[key] == nil {
// 			continue
// 		}

// 		if err := groupKeysByCollection(key, data[key], restrictModeIDs); err != nil {
// 			return nil, fmt.Errorf("grouping keys by collection: %w", err)
// 		}
// 	}

// 	if len(restrictModeIDs) == 0 {
// 		return nil, nil
// 	}

// 	// Call restrict Mode function for each collection.
// 	mperms := perm.NewMeetingPermission(ds, uid)
// 	times := make(map[string]timeCount, len(restrictModeIDs))
// 	orderedCMs := sortRestrictModeIDs(restrictModeIDs)
// 	allowedMods := make(map[collection.CM]*set.Set[int])
// 	for _, cm := range orderedCMs {
// 		ids := restrictModeIDs[cm]
// 		idsCount := ids.Len()
// 		start := time.Now()

// 		modeFunc, err := restrictModefunc(cm.Collection, cm.Mode)
// 		if err != nil {
// 			return nil, fmt.Errorf("getting restiction mode for %s/%s: %w", cm.Collection, cm.Mode, err)
// 		}

// 		allowedIDs, err := modeFunc(ctx, ds, mperms, ids.List()...)
// 		if err != nil {
// 			var errDoesNotExist dsfetch.DoesNotExistError
// 			if !errors.As(err, &errDoesNotExist) {
// 				return nil, fmt.Errorf("calling collection %s modefunc %s with ids %v: %w", cm.Collection, cm.Mode, ids.List(), err)
// 			}
// 		}
// 		allowedMods[cm] = set.New(allowedIDs...)

// 		duration := time.Since(start)
// 		times[cm.Collection+"/"+cm.Mode] = timeCount{time: duration, count: idsCount}
// 	}

// 	// Remove restricted keys.
// 	for key := range data {
// 		if data[key] == nil {
// 			continue
// 		}

// 		restrictionMode, err := restrictModeName(key.Collection, key.Field)
// 		if err != nil {
// 			return nil, fmt.Errorf("getting restriction Mode for %s: %w", key, err)
// 		}

// 		cm := collection.CM{Collection: key.Collection, Mode: restrictionMode}
// 		if !allowedMods[cm].Has(key.ID) {
// 			data[key] = nil
// 			continue
// 		}

// 		newValue, ok, err := manipulateRelations(key, data[key], allowedMods)
// 		if err != nil {
// 			return nil, fmt.Errorf("new value for relation key %s: %w", key, err)
// 		}

// 		if ok {
// 			data[key] = newValue
// 		}
// 	}

// 	return times, nil
// }

// func restrictSuperAdmin(ctx context.Context, getter datastore.Getter, uid int, data map[dskey.Key][]byte) error {
// 	ds := dsfetch.New(getter)
// 	mperms := perm.NewMeetingPermission(ds, uid)

// 	for key := range data {
// 		if data[key] == nil {
// 			continue
// 		}

// 		restricter := collection.Collection(key.Collection)
// 		if restricter == nil {
// 			// Superadmins can see unknown collections.
// 			continue
// 		}

// 		type superRestricter interface {
// 			SuperAdmin(mode string) collection.FieldRestricter
// 		}
// 		sr, ok := restricter.(superRestricter)
// 		if !ok {
// 			continue
// 		}

// 		restrictionMode, err := restrictModeName(key.Collection, key.Field)
// 		if err != nil {
// 			return fmt.Errorf("getting restriction Mode for %s: %w", key, err)
// 		}

// 		modefunc := sr.SuperAdmin(restrictionMode)
// 		if modefunc == nil {
// 			// Do not restrict unknown fields that are not implemented.
// 			continue
// 		}

// 		allowed, err := modefunc(ctx, ds, mperms, key.ID)
// 		if err != nil {
// 			return fmt.Errorf("calling mode func: %w", err)
// 		}

// 		if len(allowed) == 0 {
// 			data[key] = nil
// 		}
// 	}
// 	return nil
// }

// // groupKeysByCollection groups all the keys in data by there collection.
// func groupKeysByCollection(key dskey.Key, value []byte, restrictModeIDs map[collection.CM]set.Set[int]) error {
// 	restrictionMode, err := restrictModeName(key.Collection, key.Field)
// 	if err != nil {
// 		return fmt.Errorf("getting restriction Mode for %s: %w", key, err)
// 	}

// 	cm := collection.CM{Collection: key.Collection, Mode: restrictionMode}
// 	if restrictModeIDs[cm].IsNotInitialized() {
// 		restrictModeIDs[cm] = set.New[int]()
// 	}
// 	restrictModeIDs[cm].Add(key.ID)

// 	if err := addRelationToRestrictModeIDs(key, value, restrictModeIDs); err != nil {
// 		return fmt.Errorf("check %s for relation: %w", key, err)
// 	}

// 	return nil
// }

// func addRelationToRestrictModeIDs(key dskey.Key, value []byte, restrictModeIDs map[collection.CM]set.Set[int]) error {
// 	keyPrefix := templateKeyPrefix(key.CollectionField())

// 	cm, id, ok, err := isRelation(keyPrefix, value)
// 	if err != nil {
// 		return fmt.Errorf("checking for relation: %w", err)
// 	}

// 	if ok {
// 		if restrictModeIDs[cm].IsNotInitialized() {
// 			restrictModeIDs[cm] = set.New[int]()
// 		}
// 		restrictModeIDs[cm].Add(id)
// 		return nil
// 	}

// 	cm, ids, ok, err := isRelationList(keyPrefix, value)
// 	if err != nil {
// 		return fmt.Errorf("checking for relation-list: %w", err)
// 	}

// 	if ok {
// 		if restrictModeIDs[cm].IsNotInitialized() {
// 			restrictModeIDs[cm] = set.New[int]()
// 		}
// 		restrictModeIDs[cm].Add(ids...)
// 		return nil
// 	}

// 	cm, id, ok, err = isGenericRelation(keyPrefix, value)
// 	if err != nil {
// 		return fmt.Errorf("checking for generic-relation: %w", err)
// 	}

// 	if ok {
// 		if restrictModeIDs[cm].IsNotInitialized() {
// 			restrictModeIDs[cm] = set.New[int]()
// 		}
// 		restrictModeIDs[cm].Add(id)
// 		return nil
// 	}

// 	mcm, _, ok, err := isGenericRelationList(keyPrefix, value)
// 	if err != nil {
// 		return fmt.Errorf("checking for generic-relation-list: %w", err)
// 	}

// 	if ok {
// 		for _, cmID := range mcm {
// 			if restrictModeIDs[cmID.cm].IsNotInitialized() {
// 				restrictModeIDs[cmID.cm] = set.New[int]()
// 			}
// 			restrictModeIDs[cmID.cm].Add(cmID.id)
// 		}
// 	}

// 	return nil
// }

// // manipulateRelations changes the value of relation fields.
// //
// // The first return value is the new value. The second is, if the value was
// // manipulated.q
// func manipulateRelations(key dskey.Key, value []byte, allowedRestrictions map[collection.CM]*set.Set[int]) ([]byte, bool, error) {
// 	keyPrefix := templateKeyPrefix(key.CollectionField())

// 	cm, id, ok, err := isRelation(keyPrefix, value)
// 	if err != nil {
// 		return nil, false, fmt.Errorf("checking %s for relation: %w", key, err)
// 	}

// 	if ok {
// 		return nil, !allowedRestrictions[cm].Has(id), nil
// 	}

// 	cm, ids, ok, err := isRelationList(keyPrefix, value)
// 	if err != nil {
// 		return nil, false, fmt.Errorf("checking %s for relation-list: %w", key, err)
// 	}

// 	if ok {
// 		allowed := make([]int, 0, len(ids))
// 		for _, id := range ids {
// 			if allowedRestrictions[cm].Has(id) {
// 				allowed = append(allowed, id)
// 			}
// 		}

// 		if len(allowed) != len(ids) {
// 			newValue, err := json.Marshal(allowed)
// 			if err != nil {
// 				return nil, false, fmt.Errorf("marshal new value for key %s: %w", key, err)
// 			}
// 			return newValue, true, nil
// 		}
// 		return nil, false, nil
// 	}

// 	cm, id, ok, err = isGenericRelation(keyPrefix, value)
// 	if err != nil {
// 		return nil, false, fmt.Errorf("checking %s for generic-relation: %w", key, err)
// 	}

// 	if ok {
// 		return nil, !allowedRestrictions[cm].Has(id), nil
// 	}

// 	mcm, genericIDs, ok, err := isGenericRelationList(keyPrefix, value)
// 	if err != nil {
// 		return nil, false, fmt.Errorf("checking %s for generic-relation-list: %w", key, err)
// 	}

// 	if ok {
// 		allowed := make([]string, 0, len(genericIDs))
// 		for genericID, cmID := range mcm {
// 			if allowedRestrictions[cmID.cm].Has(cmID.id) {
// 				allowed = append(allowed, genericID)
// 			}
// 		}

// 		if len(allowed) != len(genericIDs) {
// 			newValue, err := json.Marshal(allowed)
// 			if err != nil {
// 				return nil, false, fmt.Errorf("marshal new value for key %s: %w", key, err)
// 			}
// 			return newValue, true, nil
// 		}

// 		return nil, false, nil
// 	}

// 	return nil, false, nil
// }

// func isRelation(keyPrefix string, value []byte) (collection.CM, int, bool, error) {
// 	toCollectionfield, ok := relationFields[keyPrefix]
// 	if !ok {
// 		return collection.CM{}, 0, false, nil
// 	}

// 	var id int
// 	if err := json.Unmarshal(value, &id); err != nil {
// 		return collection.CM{}, 0, false, fmt.Errorf("decoding %q (`%s`): %w", keyPrefix, value, err)
// 	}

// 	coll, field, _ := strings.Cut(toCollectionfield, "/")
// 	fieldMode, err := restrictModeName(coll, field)
// 	if err != nil {
// 		return collection.CM{}, 0, false, fmt.Errorf("building relation field field mode: %w", err)
// 	}

// 	return collection.CM{Collection: coll, Mode: fieldMode}, id, true, nil
// }

// func isRelationList(keyPrefix string, value []byte) (collection.CM, []int, bool, error) {
// 	toCollectionfield, ok := relationListFields[keyPrefix]
// 	if !ok {
// 		return collection.CM{}, nil, false, nil
// 	}

// 	var ids []int
// 	if err := json.Unmarshal(value, &ids); err != nil {
// 		return collection.CM{}, nil, false, fmt.Errorf("decoding value (size: %d): %w", len(value), err)
// 	}

// 	coll, field, _ := strings.Cut(toCollectionfield, "/")
// 	fieldMode, err := restrictModeName(coll, field)
// 	if err != nil {
// 		return collection.CM{}, nil, false, fmt.Errorf("building relation field field mode: %w", err)
// 	}

// 	return collection.CM{Collection: coll, Mode: fieldMode}, ids, true, nil
// }

// func isGenericRelation(keyPrefix string, value []byte) (collection.CM, int, bool, error) {
// 	toCollectionfield, ok := genericRelationFields[keyPrefix]
// 	if !ok {
// 		return collection.CM{}, 0, false, nil
// 	}

// 	var genericID string
// 	if err := json.Unmarshal(value, &genericID); err != nil {
// 		return collection.CM{}, 0, false, fmt.Errorf("decoding %q: %w", keyPrefix, err)
// 	}

// 	cm, id, err := genericKeyToCollectionMode(genericID, toCollectionfield)
// 	if err != nil {
// 		return collection.CM{}, 0, false, fmt.Errorf("parsing generic key: %w", err)
// 	}

// 	return cm, id, true, nil
// }

// type collectionModeID struct {
// 	cm collection.CM
// 	id int
// }

// func isGenericRelationList(keyPrefix string, value []byte) (map[string]collectionModeID, []string, bool, error) {
// 	toCollectionfield, ok := genericRelationListFields[keyPrefix]
// 	if !ok {
// 		return nil, nil, false, nil
// 	}

// 	var genericIDs []string
// 	if err := json.Unmarshal(value, &genericIDs); err != nil {
// 		return nil, nil, false, fmt.Errorf("decoding %q: %w", keyPrefix, err)
// 	}

// 	mcm := make(map[string]collectionModeID, len(genericIDs))
// 	for _, genericID := range genericIDs {
// 		cm, id, err := genericKeyToCollectionMode(genericID, toCollectionfield)
// 		if err != nil {
// 			return nil, nil, false, fmt.Errorf("parsing generic key: %w", err)
// 		}
// 		mcm[genericID] = collectionModeID{cm, id}
// 	}

// 	return mcm, genericIDs, true, nil
// }

// // genericKeyToCollectionMode calls f for each collection mode.
// func genericKeyToCollectionMode(genericID string, toCollectionFieldMap map[string]string) (collection.CM, int, error) {
// 	coll, rawID, found := strings.Cut(genericID, "/")
// 	if !found {
// 		// TODO LAST ERROR
// 		return collection.CM{}, 0, fmt.Errorf("invalid generic relation: %s", genericID)
// 	}

// 	id, err := strconv.Atoi(rawID)
// 	if err != nil {
// 		// TODO LAST ERROR
// 		return collection.CM{}, 0, fmt.Errorf("invalid generic relation, no id: %s", genericID)
// 	}

// 	toField := toCollectionFieldMap[coll]
// 	if toField == "" {
// 		// TODO LAST ERROR
// 		return collection.CM{}, 0, fmt.Errorf("unknown generic relation: %s", coll)
// 	}

// 	fieldMode, err := restrictModeName(coll, toField)
// 	if err != nil {
// 		return collection.CM{}, 0, fmt.Errorf("building generic relation field mode: %w", err)
// 	}

// 	return collection.CM{Collection: coll, Mode: fieldMode}, id, nil
// }

// FieldsForCollection returns the list of fieldnames for an collection.
func FieldsForCollection(collection string) []string {
	return collectionFields[collection]
}
