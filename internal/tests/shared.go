package tests

import (
	"encoding/json"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"
)

func addIntToJsonArray(arrayJson string, value int) string {
	var array []int
	_ = json.Unmarshal([]byte(arrayJson), &array)
	array = append(array, value)
	newArrayJson, _ := json.Marshal(array)
	return string(newArrayJson)
}

func addStringToJsonArray(arrayJson string, value string) string {
	var array []string
	_ = json.Unmarshal([]byte(arrayJson), &array)
	array = append(array, value)
	newArrayJson, _ := json.Marshal(array)
	return string(newArrayJson)
}

type TestDataProvider struct {
	Data definitions.FqfieldData
}

func NewTestDataProvider() *TestDataProvider {
	var testDataProvider = &TestDataProvider{Data: nil}
	testDataProvider.SetDefault()
	return testDataProvider
}

func (t *TestDataProvider) GetDataprovider() dataprovider.DataProvider {
	return dataprovider.NewDataProvider(t)
}

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

func (t *TestDataProvider) EnableAnonymous() {
	t.Data["meeting/1/enable_anonymous"] = "true"
}

func (t *TestDataProvider) AddUser(id int) {
	t.Data["user/"+strconv.Itoa(id)+"/id"] = strconv.Itoa(id)
}

func (t *TestDataProvider) AddUserToMeeting(userId, meetingId int) {
	t.Data["user/"+strconv.Itoa(userId)+"/id"] = strconv.Itoa(userId)
	meetingField := "meeting/" + strconv.Itoa(meetingId) + "/user_ids"
	t.Data[meetingField] = addIntToJsonArray(t.getFieldWithDefault(meetingField, "[]"), userId)
}

func (t *TestDataProvider) AddUserWithSuperadminRole(id int) {
	t.Data["role/1/user_ids"] = addIntToJsonArray(t.getFieldWithDefault("role/1/user_ids", "[]"), id)
	t.Data["user/"+strconv.Itoa(id)+"/id"] = strconv.Itoa(id)
	t.Data["user/"+strconv.Itoa(id)+"/role_id"] = "1"
}

func (t *TestDataProvider) AddUserWithAdminGroupToMeeting(userId, meetingId int) {
	t.AddUserToMeeting(userId, meetingId)
	t.Data["group/3/user_ids"] = addIntToJsonArray(t.getFieldWithDefault("group/3/user_ids", "[]"), userId)
	t.Data["user/"+strconv.Itoa(userId)+"/group_"+strconv.Itoa(meetingId)+"_ids"] = "[3]"
}

func (t *TestDataProvider) AddPermissionToGroup(groupId int, permission string) {
	fqfield := "group/" + strconv.Itoa(groupId) + "/permissions"
	t.Data[fqfield] = addStringToJsonArray(t.getFieldWithDefault(fqfield, "[]"), permission)
}

func (t *TestDataProvider) getFieldWithDefault(fqfield definitions.Fqfield, defaultValue definitions.Value) string {
	if value, ok := t.Data[fqfield]; ok {
		return value
	} else {
		return defaultValue
	}
}

func (t *TestDataProvider) Set(fqfield definitions.Fqfield, value definitions.Value) {
	t.Data[fqfield] = value
}

func (t TestDataProvider) Get(fqfields []definitions.Fqfield) definitions.FqfieldData {
	data := make(definitions.FqfieldData)
	for _, field := range fqfields {
		value, ok := t.Data[field]
		if ok {
			data[field] = value
		}
	}
	return data
}
