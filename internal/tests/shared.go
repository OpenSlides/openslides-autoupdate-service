package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

func addIntToJSONArray(arrayJSON definitions.Value, value int) definitions.Value {
	var array []int
	_ = json.Unmarshal(arrayJSON, &array)
	array = append(array, value)
	newArrayJSON, _ := json.Marshal(array)
	return newArrayJSON
}

func addStringToJSONArray(arrayJSON definitions.Value, value string) definitions.Value {
	var array []string
	_ = json.Unmarshal(arrayJSON, &array)
	array = append(array, value)
	newArrayJSON, _ := json.Marshal(array)
	return newArrayJSON
}

// TestDataProvider does ...
type TestDataProvider struct {
	ctx  context.Context
	Data definitions.FqfieldData
}

// NewTestDataProvider does ...
func NewTestDataProvider(ctx context.Context) *TestDataProvider {
	var testDataProvider = &TestDataProvider{ctx, nil}
	testDataProvider.SetDefault()
	return testDataProvider
}

// GetDataprovider does ...
func (t *TestDataProvider) GetDataprovider() dataprovider.DataProvider {
	return dataprovider.NewDataProvider(t.ctx, t)
}

// SetDefault does ....
// Create organisation (1), committe(1) and superadminrole(1)
// Creates meeting 1 with groups:
// - 1: Default
// - 2: Delegates
// - 3: Admin
func (t *TestDataProvider) SetDefault() {
	data := map[string]string{
		"organisation/1/id":                 "1",
		"organisation/1/committee_ids":      "[1]",
		"organisation/1/role_ids":           "[1]",
		"organisation/1/superadmin_role_id": "1",

		// Role
		"role/1/id":              "1",
		"role/1/name":            `"Superadmin role"`,
		"role/1/permissions":     "[]",
		"role/1/organisation_id": "1",
		"role/1/superadmin_role_for_organisation_id": "1",

		// Committee
		"committee/1/id":                 "1",
		"committee/1/meeting_ids":        "[1]",
		"committee/1/default_meeting_id": "1",
		"committee/1/member_ids":         "[]",
		"committee/1/manager_ids":        "[]",
		"committee/1/organisation_id":    "1",

		// Meeting
		"meeting/1/id":                  "1",
		"meeting/1/enable_anonymous":    "false",
		"meeting/1/group_ids":           "[1, 2, 3]",
		"meeting/1/committee_id":        "1",
		"meeting/1/default_group_id":    "1",
		"meeting/1/superadmin_group_id": "3",

		// Group 1: Default
		"group/1/id": "1",
		"group/1/superadmin_group_for_meeting_id": "null",
		"group/1/default_group_for_meeting_id":    "1",
		"group/1/permissions": `[
            "agenda.can_see",
            "agenda.can_see_internal_items",
            "assignments.can_see",
            "core.can_see_frontpage",
            "core.can_see_projector",
            "mediafiles.can_see",
            "motions.can_see",
            "users.can_see_name"
        ]`,
		"group/1/meeting_id": "1",

		// Group 2: Delegate
		"group/2/id": "2",
		"group/2/permissions": `[
			"agenda.can_see",
            "agenda.can_see_internal_items",
            "agenda.can_be_speaker",
            "assignments.can_nominate_other",
            "assignments.can_nominate_self",
            "assignments.can_see",
            "core.can_see_frontpage",
            "core.can_see_projector",
            "mediafiles.can_see",
            "motions.can_create",
            "motions.can_manage",
            "motions.can_see",
            "motions.can_support",
            "users.can_see_name"
        ]`,
		"group/2/meeting_id": "1",

		// Group 3: Superadmin
		"group/3/id": "3",
		"group/3/superadmin_group_for_meeting_id": "1",
		"group/3/meeting_id":                      "1",
	}

	t.Data = make(map[string]definitions.Value)
	for k, v := range data {
		t.Data[k] = []byte(v)
	}
}

// AddBasicModel does ...
func (t *TestDataProvider) AddBasicModel(collection definitions.Collection, id definitions.Id) {
	t.Set(collection+"/"+strconv.Itoa(id)+"/id", strconv.Itoa(id))
	t.Set(collection+"/"+strconv.Itoa(id)+"/meeting_id", "1")
	t.Set("meeting/1/"+collection+"_ids", "["+strconv.Itoa(id)+"]")
}

// EnableAnonymous does ...
func (t *TestDataProvider) EnableAnonymous() {
	t.Data["meeting/1/enable_anonymous"] = []byte("true")
}

// AddUser does ...
func (t *TestDataProvider) AddUser(id definitions.Id) {
	t.Data["user/"+strconv.Itoa(id)+"/id"] = []byte(strconv.Itoa(id))
}

// AddUserToMeeting does ...
func (t *TestDataProvider) AddUserToMeeting(userID, meetingID definitions.Id) {
	t.AddUser(userID)
	meetingField := "meeting/" + strconv.Itoa(meetingID) + "/user_ids"
	t.Data[meetingField] = addIntToJSONArray(t.getFieldWithDefault(meetingField, "[]"), userID)
}

// AddUserToCommitteeAsManager does ...
func (t *TestDataProvider) AddUserToCommitteeAsManager(userID, committeeID definitions.Id) {
	t.AddUser(userID)
	userField := "user/" + strconv.Itoa(userID) + "/committee_as_manager_ids"
	t.Data[userField] = addIntToJSONArray(t.getFieldWithDefault(userField, "[]"), committeeID)
	committeeField := "committee/" + strconv.Itoa(committeeID) + "/manager_ids"
	t.Data[committeeField] = addIntToJSONArray(t.getFieldWithDefault(committeeField, "[]"), userID)
}

// AddUserWithSuperadminRole does ...
func (t *TestDataProvider) AddUserWithSuperadminRole(id definitions.Id) {
	t.Data["role/1/user_ids"] = addIntToJSONArray(t.getFieldWithDefault("role/1/user_ids", "[]"), id)
	t.Data["user/"+strconv.Itoa(id)+"/id"] = []byte(strconv.Itoa(id))
	t.Data["user/"+strconv.Itoa(id)+"/role_id"] = []byte("1")
}

// AddUserWithAdminGroupToMeeting does ...
func (t *TestDataProvider) AddUserWithAdminGroupToMeeting(userID, meetingID definitions.Id) {
	t.AddUserToMeeting(userID, meetingID)
	t.Data["group/3/user_ids"] = addIntToJSONArray(t.getFieldWithDefault("group/3/user_ids", "[]"), userID)
	t.Data["user/"+strconv.Itoa(userID)+"/group_"+strconv.Itoa(meetingID)+"_ids"] = []byte("[3]")
}

// AddPermissionToGroup does ...
func (t *TestDataProvider) AddPermissionToGroup(groupID definitions.Id, permission string) {
	fqfield := "group/" + strconv.Itoa(groupID) + "/permissions"
	t.Data[fqfield] = addStringToJSONArray(t.getFieldWithDefault(fqfield, "[]"), permission)
}

func (t *TestDataProvider) getFieldWithDefault(fqfield definitions.Fqfield, defaultValue string) definitions.Value {
	if value, ok := t.Data[fqfield]; ok {
		return value
	}

	return []byte(defaultValue)
}

// Set does ...
func (t *TestDataProvider) Set(fqfield definitions.Fqfield, value string) {
	t.Data[fqfield] = []byte(value)
}

// Get does ...
func (t TestDataProvider) Get(ctx context.Context, fqfields ...definitions.Fqfield) ([]json.RawMessage, error) {
	if ctx != t.ctx {
		return nil, fmt.Errorf("the context was not propagated")
	}

	data := make([]json.RawMessage, len(fqfields))
	for i, field := range fqfields {
		value, ok := t.Data[field]
		if !ok {
			data[i] = nil
			continue
		}

		data[i] = json.RawMessage(value)
	}
	return data, nil
}
