package topic_test

import (
	"testing"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func assertDeleteFailWithError(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Delete(ctx)
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

func assertDeleteIsNotAllowed(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Delete(ctx)
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

func assertDeleteIsAllowed(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Delete(ctx)
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
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteFailWithError(t, ctx)
}

func TestDeleteSuperadminRole(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
	dp.AddUserWithSuperadminRole(1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, ctx)
}

func TestDeleteNoId(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteFailWithError(t, ctx)
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
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, ctx)
}

func TestDeleteAdminUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, ctx)
}

func TestDeleteUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, ctx)
}

func TestDeleteUserNoPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, ctx)
}

func TestDeleteInvaldFields(t *testing.T) {
	dp := tests.NewTestDataProvider()
	dp.AddUserWithSuperadminRole(1)
	data := definitions.FqfieldData{
		"not_allowed": "some value",
	}
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteFailWithError(t, ctx)
}

func TestDeleteDisabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, ctx)
}

func TestDeleteEnabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	dp.EnableAnonymous()
	data := definitions.FqfieldData{
		"id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsNotAllowed(t, ctx)
}

func TestDeleteEnabledAnonymousWithPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	dp.EnableAnonymous()
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	data := definitions.FqfieldData{
		"id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertDeleteIsAllowed(t, ctx)
}
