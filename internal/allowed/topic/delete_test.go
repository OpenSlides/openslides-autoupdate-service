package topic_test

import (
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func assertDeleteFailWithError(t *testing.T, params *allowed.IsAllowedParams) {
	allowed, addition, err := topic.Delete(params)
	if nil != addition {
		t.Errorf("Expected to fail without an addition: %s", addition)
	}
	if nil == err {
		t.Errorf("Expected to fail with an error")
	}

	if allowed {
		t.Errorf("Expected to fail with allowed=false")
	}
}

func assertDeleteIsNotAllowed(t *testing.T, params *allowed.IsAllowedParams) {
	allowed, addition, err := topic.Delete(params)
	if nil != addition {
		t.Errorf("Expected to fail without an addition: %s", addition)
	}
	if nil != err {
		t.Errorf("Expected to fail without an error (error: %s)", err)
	}

	if allowed {
		t.Errorf("Expected to fail with allowed=false")
	}
}

func assertDeleteIsAllowed(t *testing.T, params *allowed.IsAllowedParams) {
	allowed, addition, err := topic.Delete(params)
	if nil != addition {
		t.Errorf("Expected to fail without an addition: %s", addition)
	}
	if nil != err {
		t.Errorf("Expected to fail without an error (error: %s)", err)
	}

	if !allowed {
		t.Errorf("Expected to be allowed")
	}
}

func TestDeleteUnknownUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"id": "1",
	}
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteFailWithError(t, params)
}

func TestDeleteSuperadminRole(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
	dp.AddUserWithSuperadminRole(1)
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, params)
}

func TestDeleteNoId(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteFailWithError(t, params)
}

func addBasicTopic(dp *tests.TestDataProvider) {
	dp.Set("topic/1/id", "1")
	dp.Set("topic/1/meeting_id", "1")
	dp.Set("meeting/1/topic_ids", "[1]")
}

func TestDeleteUserNotInMeeting(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUser(1)
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, params)
}

func TestDeleteAdminUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, params)
}

func TestDeleteUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, params)
}

func TestDeleteUserNoPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, params)
}

func TestDeleteInvaldFields(t *testing.T) {
	dp := tests.NewTestDataProvider()
	dp.AddUserWithSuperadminRole(1)
	data := definitions.FqfieldData{
		"not_allowed": "some value",
	}
	params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteFailWithError(t, params)
}

func TestDeleteDisabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, params)
}

func TestDeleteEnabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	dp.EnableAnonymous()
	data := definitions.FqfieldData{
		"id": "1",
	}
	params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, params)
}

func TestDeleteEnabledAnonymousWithPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	dp.EnableAnonymous()
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	data := definitions.FqfieldData{
		"id": "1",
	}
	params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, params)
}
