package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

func addIntToJSONArray(arrayJSON json.RawMessage, value int) json.RawMessage {
	var array []int
	_ = json.Unmarshal(arrayJSON, &array)
	array = append(array, value)
	newArrayJSON, _ := json.Marshal(array)
	return newArrayJSON
}

func addStringToJSONArray(arrayJSON json.RawMessage, value string) json.RawMessage {
	var array []string
	_ = json.Unmarshal(arrayJSON, &array)
	array = append(array, value)
	newArrayJSON, _ := json.Marshal(array)
	return newArrayJSON
}

// TestDataProvider does ...
type TestDataProvider struct {
	Data map[string]json.RawMessage
}

// NewTestDataProvider does ...
func NewTestDataProvider() *TestDataProvider {
	var testDataProvider = &TestDataProvider{nil}
	testDataProvider.SetDefault()
	return testDataProvider
}

// GetDataprovider does ...
func (t *TestDataProvider) GetDataprovider() dataprovider.DataProvider {
	return dataprovider.DataProvider{External: t}
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
            "agenda.can_see_internal_items",
            "assignment.can_see",
            "mediafile.can_see",
            "motion.can_see",
            "user.can_see_name"
        ]`,
		"group/1/meeting_id": "1",

		// Group 2: Delegate
		"group/2/id": "2",
		"group/2/permissions": `[
            "agenda.can_see_internal_items",
            "agenda.can_be_speaker",
            "assignment.can_nominate_other",
			"assignment.can_nominate_self",
            "mediafile.can_see",
            "motion.can_manage",
            "user.can_see_name"
        ]`,
		"group/2/meeting_id": "1",

		// Group 3: Superadmin
		"group/3/id": "3",
		"group/3/superadmin_group_for_meeting_id": "1",
		"group/3/meeting_id":                      "1",
	}

	t.Data = make(map[string]json.RawMessage)
	for k, v := range data {
		t.Data[k] = []byte(v)
	}
}

// AddBasicModel does ...
func (t *TestDataProvider) AddBasicModel(collection string, id int) {
	t.Set(collection+"/"+strconv.Itoa(id)+"/id", strconv.Itoa(id))
	t.Set(collection+"/"+strconv.Itoa(id)+"/meeting_id", "1")
	t.Set("meeting/1/"+collection+"_ids", "["+strconv.Itoa(id)+"]")
}

// EnableAnonymous does ...
func (t *TestDataProvider) EnableAnonymous() {
	t.Data["meeting/1/enable_anonymous"] = []byte("true")
}

// AddUser does ...
func (t *TestDataProvider) AddUser(id int) {
	t.Data["user/"+strconv.Itoa(id)+"/id"] = []byte(strconv.Itoa(id))
}

// AddUserToMeeting does ...
func (t *TestDataProvider) AddUserToMeeting(userID, meetingID int) {
	t.AddUser(userID)
	meetingField := "meeting/" + strconv.Itoa(meetingID) + "/user_ids"
	t.Data[meetingField] = addIntToJSONArray(t.getFieldWithDefault(meetingField, "[]"), userID)
}

// AddUserToGroup adds the user to the group in the given meeting.
func (t *TestDataProvider) AddUserToGroup(userID, meetingID, groupID int) {
	groupFQField := fmt.Sprintf("group/%d/user_ids", groupID)
	t.Data[groupFQField] = addIntToJSONArray(t.getFieldWithDefault(groupFQField, "[]"), userID)

	userFQField := fmt.Sprintf("user/%d/group_$%d_ids", userID, meetingID)
	t.Data[userFQField] = addIntToJSONArray(t.getFieldWithDefault(userFQField, "[]"), groupID)
}

// AddUserToCommitteeAsManager does ...
func (t *TestDataProvider) AddUserToCommitteeAsManager(userID, committeeID int) {
	t.AddUser(userID)
	userField := "user/" + strconv.Itoa(userID) + "/committee_as_manager_ids"
	t.Data[userField] = addIntToJSONArray(t.getFieldWithDefault(userField, "[]"), committeeID)
	committeeField := "committee/" + strconv.Itoa(committeeID) + "/manager_ids"
	t.Data[committeeField] = addIntToJSONArray(t.getFieldWithDefault(committeeField, "[]"), userID)
}

// AddUserWithSuperadminRole does ...
func (t *TestDataProvider) AddUserWithSuperadminRole(id int) {
	t.Data["role/1/user_ids"] = addIntToJSONArray(t.getFieldWithDefault("role/1/user_ids", "[]"), id)
	t.Data["user/"+strconv.Itoa(id)+"/id"] = []byte(strconv.Itoa(id))
	t.Data["user/"+strconv.Itoa(id)+"/role_id"] = []byte("1")
}

// AddUserWithAdminGroupToMeeting does ...
func (t *TestDataProvider) AddUserWithAdminGroupToMeeting(userID, meetingID int) {
	t.AddUserToMeeting(userID, meetingID)
	t.Data["group/3/user_ids"] = addIntToJSONArray(t.getFieldWithDefault("group/3/user_ids", "[]"), userID)
	t.Data["user/"+strconv.Itoa(userID)+"/group_$"+strconv.Itoa(meetingID)+"_ids"] = []byte("[3]")
}

// AddPermissionToGroup does ...
func (t *TestDataProvider) AddPermissionToGroup(groupID int, permission string) {
	fqfield := "group/" + strconv.Itoa(groupID) + "/permissions"
	t.Data[fqfield] = addStringToJSONArray(t.getFieldWithDefault(fqfield, "[]"), permission)
}

func (t *TestDataProvider) getFieldWithDefault(fqfield string, defaultValue string) json.RawMessage {
	if value, ok := t.Data[fqfield]; ok {
		return value
	}

	return []byte(defaultValue)
}

// Set does ...
func (t *TestDataProvider) Set(fqfield string, value string) {
	t.Data[fqfield] = []byte(value)
}

// Get does ...
func (t TestDataProvider) Get(ctx context.Context, fqfields ...string) ([]json.RawMessage, error) {
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
