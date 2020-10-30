package tests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
)

func addIntToJSONArray(arrayJSON string, value int) string {
	var array []int
	_ = json.Unmarshal([]byte(arrayJSON), &array)
	array = append(array, value)
	newArrayJSON, _ := json.Marshal(array)
	return string(newArrayJSON)
}

func addStringToJSONArray(arrayJSON string, value string) string {
	var array []string
	_ = json.Unmarshal([]byte(arrayJSON), &array)
	array = append(array, value)
	newArrayJSON, _ := json.Marshal(array)
	return string(newArrayJSON)
}

// TestDataProvider does ...
type TestDataProvider struct {
	Data definitions.FqfieldData
}

// NewTestDataProvider does ...
func NewTestDataProvider() *TestDataProvider {
	var testDataProvider = &TestDataProvider{Data: nil}
	testDataProvider.SetDefault()
	return testDataProvider
}

// GetDataprovider does ...
func (t *TestDataProvider) GetDataprovider() dataprovider.DataProvider {
	return dataprovider.NewDataProvider(t)
}

// SetDefault does ....
// Create organisation (1), committe(1) and superadminrole(1)
// Creates meeting 1 with groups:
// - 1: Default
// - 2: Delegates
// - 3: Admin
func (t *TestDataProvider) SetDefault() {
	t.Data = definitions.FqfieldData{
		"organisation/1/committee_ids":      "[1]",
		"organisation/1/role_ids":           "[1]",
		"organisation/1/superadmin_role_id": "1",

		// Role
		"role/1/name":                                "Superadmin role",
		"role/1/permissions":                         "[]",
		"role/1/organisation_id":                     "1",
		"role/1/superadmin_role_for_organisation_id": "1",

		// Committee
		"committee/1/meeting_ids":        "[1]",
		"committee/1/default_meeting_id": "1",
		"committee/1/member_ids":         "[]",
		"committee/1/manager_ids":        "[]",
		"committee/1/organisation_id":    "1",

		// Meeting
		"meeting/1/enable_anonymous":    "false",
		"meeting/1/group_ids":           "[1, 2, 3]",
		"meeting/1/committee_id":        "1",
		"meeting/1/default_group_id":    "1",
		"meeting/1/superadmin_group_id": "3",

		// Group 1: Default
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
		"group/3/superadmin_group_for_meeting_id": "1",
		"group/3/meeting_id":                      "1",
	}
}

// EnableAnonymous does ...
func (t *TestDataProvider) EnableAnonymous() {
	t.Data["meeting/1/enable_anonymous"] = "true"
}

// AddUser does ...
func (t *TestDataProvider) AddUser(id int) {
	t.Data["user/"+strconv.Itoa(id)+"/id"] = strconv.Itoa(id)
}

// AddUserToMeeting does ...
func (t *TestDataProvider) AddUserToMeeting(userID, meetingID int) {
	t.Data["user/"+strconv.Itoa(userID)+"/id"] = strconv.Itoa(userID)
	meetingField := "meeting/" + strconv.Itoa(meetingID) + "/user_ids"
	t.Data[meetingField] = addIntToJSONArray(t.getFieldWithDefault(meetingField, "[]"), userID)
}

// AddUserWithSuperadminRole does ...
func (t *TestDataProvider) AddUserWithSuperadminRole(id int) {
	t.Data["role/1/user_ids"] = addIntToJSONArray(t.getFieldWithDefault("role/1/user_ids", "[]"), id)
	t.Data["user/"+strconv.Itoa(id)+"/id"] = strconv.Itoa(id)
	t.Data["user/"+strconv.Itoa(id)+"/role_id"] = "1"
}

// AddUserWithAdminGroupToMeeting does ...
func (t *TestDataProvider) AddUserWithAdminGroupToMeeting(userID, meetingID int) {
	t.AddUserToMeeting(userID, meetingID)
	t.Data["group/3/user_ids"] = addIntToJSONArray(t.getFieldWithDefault("group/3/user_ids", "[]"), userID)
	t.Data["user/"+strconv.Itoa(userID)+"/group_"+strconv.Itoa(meetingID)+"_ids"] = "[3]"
}

// AddPermissionToGroup does ...
func (t *TestDataProvider) AddPermissionToGroup(groupID int, permission string) {
	fqfield := "group/" + strconv.Itoa(groupID) + "/permissions"
	t.Data[fqfield] = addStringToJSONArray(t.getFieldWithDefault(fqfield, "[]"), permission)
}

func (t *TestDataProvider) getFieldWithDefault(fqfield definitions.Fqfield, defaultValue definitions.Value) string {
	if value, ok := t.Data[fqfield]; ok {
		return value
	}

	return defaultValue
}

// Set does ...
func (t *TestDataProvider) Set(fqfield definitions.Fqfield, value definitions.Value) {
	t.Data[fqfield] = value
}

// Get does ...
func (t TestDataProvider) Get(ctx context.Context, fqfields ...definitions.Fqfield) ([]json.RawMessage, error) {
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
