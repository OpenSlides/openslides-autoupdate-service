package topic_test

import (
	"testing"

	"github.com/OpenSlides/openslides-permission-service/pkg/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func assertCreateFailWithError(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Create(ctx)
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

func assertCreateIsNotAllowed(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Create(ctx)
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

func assertCreateIsAllowed(t *testing.T, ctx *allowed.IsAllowedContext) {
	allowed, addition, err := topic.Create(ctx)
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

func TestCreateUnknownUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateFailWithError(t, ctx)
}

func TestCreateSuperadminRole(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
	dp.AddUserWithSuperadminRole(1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsAllowed(t, ctx)
}

func TestCreateNoMeetingId(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateFailWithError(t, ctx)
}

func TestCreateUserNotInMeeting(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	dp.AddUser(1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsNotAllowed(t, ctx)
}

func TestCreateAdminUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	dp.AddUserWithAdminGroupToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsAllowed(t, ctx)
}

func TestCreateUser(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsAllowed(t, ctx)
}

func TestCreateUserNoPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	dp.AddUserToMeeting(1, 1)
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsNotAllowed(t, ctx)
}

func TestCreateInvaldFields(t *testing.T) {
	dp := tests.NewTestDataProvider()
	dp.AddUserWithSuperadminRole(1)
	data := definitions.FqfieldData{
		"not_allowed": "some value",
	}
	ctx := &allowed.IsAllowedContext{UserId: 1, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateFailWithError(t, ctx)
}

func TestCreateDisabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsNotAllowed(t, ctx)
}

func TestCreateEnabledAnonymous(t *testing.T) {
	dp := tests.NewTestDataProvider()
	dp.EnableAnonymous()
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsNotAllowed(t, ctx)
}

func TestCreateEnabledAnonymousWithPermissions(t *testing.T) {
	dp := tests.NewTestDataProvider()
	dp.EnableAnonymous()
	dp.AddPermissionToGroup(1, "agenda.can_manage")
	data := definitions.FqfieldData{
		"meeting_id": "1",
	}
	ctx := &allowed.IsAllowedContext{UserId: 0, Data: data, DataProvider: dp.GetDataprovider()}

	assertCreateIsAllowed(t, ctx)
}
