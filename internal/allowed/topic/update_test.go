package topic_test

import (
	"testing"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func assertUpdateFailWithError(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Update(ctx)
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

func assertUpdateIsNotAllowed(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Update(ctx)
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

func assertUpdateIsAllowed(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Update(ctx)
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

func TestUpdateUnknownUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateFailWithError(t, ctx)
}

func TestUpdateSuperadminRole(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
	dp.AddUserWithSuperadminRole(1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsAllowed(t, ctx)
}

func TestUpdateNoId(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateFailWithError(t, ctx)
}

func TestUpdateUserNotInMeeting(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUser(1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsNotAllowed(t, ctx)
}

func TestUpdateAdminUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsAllowed(t, ctx)
}

func TestUpdateUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsAllowed(t, ctx)
}

func TestUpdateUserNoPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsNotAllowed(t, ctx)
}

func TestUpdateInvaldFields(t *testing.T) {
	dp := tests.NewTestDataProvider()
	dp.AddUserWithSuperadminRole(1)
	data := definitions.FqfieldData{
		"not_allowed": "some value",
	}
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateFailWithError(t, ctx)
}

func TestUpdateDisabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	data := definitions.FqfieldData{
		"id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsNotAllowed(t, ctx)
}

func TestUpdateEnabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	dp.EnableAnonymous()
	data := definitions.FqfieldData{
		"id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsNotAllowed(t, ctx)
}

func TestUpdateEnabledAnonymousWithPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	addBasicTopic(dp)
	dp.EnableAnonymous()
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	data := definitions.FqfieldData{
		"id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertUpdateIsAllowed(t, ctx)
}
