package allowed_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

var collection = "test_collection"
var permission = "test_collection.permission"
var throughCollection = "through"
var throughField = "through_id"
var throughId = 3
var create = allowed.BuildCreate([]string{"meeting_id"}, permission)
var createThroughId = allowed.BuildCreateThroughID([]string{throughField}, throughCollection, throughField, permission)
var modify = allowed.BuildModify([]string{"id"}, collection, permission)
var modifyThroughId = allowed.BuildModifyThroughID([]string{throughField}, collection, throughCollection, throughField, permission)

func TestModify(t *testing.T) {
	t.Run("UnknownUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("SuperadminRole", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
		dp.AddUserWithSuperadminRole(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modify, params)
	})

	t.Run("NoId", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("UserNotInMeeting", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUser(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("UserIsCommitteemanager", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToCommitteeAsManager(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modify, params)
	})

	t.Run("AdminUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modify, params)
	})

	t.Run("DoesNotExist", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("User", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, permission)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modify, params)
	})

	t.Run("UserNoPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("InvaldFields", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddUserWithSuperadminRole(1)
		data := definitions.FqfieldData{
			"not_allowed": []byte("some value"),
		}
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("DisabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("EnabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		dp.EnableAnonymous()
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("EnabledAnonymousWithPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(collection, 1)
		dp.EnableAnonymous()
		dp.AddPermissionToGroup(1, permission)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modify, params)
	})
}

func TestModifyThroughId(t *testing.T) {
	t.Run("NoThroughModel", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modifyThroughId, params)
	})

	t.Run("AdminUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modifyThroughId, params)
	})

	t.Run("User", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, permission)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modifyThroughId, params)
	})

	t.Run("UserNoPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modifyThroughId, params)
	})

	t.Run("DisabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modifyThroughId, params)
	})

	t.Run("EnabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		dp.EnableAnonymous()
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modifyThroughId, params)
	})

	t.Run("EnabledAnonymousWithPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		dp.EnableAnonymous()
		dp.AddPermissionToGroup(1, permission)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modifyThroughId, params)
	})
}

func TestCreate(t *testing.T) {
	t.Run("UnknownUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("SuperadminRole", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
		dp.AddUserWithSuperadminRole(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, create, params)
	})

	t.Run("NoMeetingId", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("MeetingDoesNotExist", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1337"),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("InvalidMeetingId", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("-9"),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("UserNotInMeeting", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		dp.AddUser(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("UserIsCommitteemanager", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		dp.AddUserToCommitteeAsManager(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, create, params)
	})

	t.Run("AdminUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, create, params)
	})

	t.Run("User", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, permission)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, create, params)
	})

	t.Run("UserNoPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("InvaldFields", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddUserWithSuperadminRole(1)
		data := definitions.FqfieldData{
			"not_allowed": []byte("some value"),
		}
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("DisabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("EnabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.EnableAnonymous()
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, create, params)
	})

	t.Run("EnabledAnonymousWithPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.EnableAnonymous()
		dp.AddPermissionToGroup(1, permission)
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, create, params)
	})
}

func TestCreateThroughId(t *testing.T) {
	t.Run("UnknownUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, createThroughId, params)
	})

	t.Run("SuperadminRole", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
		dp.AddUserWithSuperadminRole(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, createThroughId, params)
	})

	t.Run("NoThroughModel", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, createThroughId, params)
	})

	t.Run("NoThroughId", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, createThroughId, params)
	})

	t.Run("AdminUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, createThroughId, params)
	})

	t.Run("User", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, permission)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, createThroughId, params)
	})

	t.Run("UserNoPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		dp.AddUserToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, createThroughId, params)
	})

	t.Run("DisabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, createThroughId, params)
	})

	t.Run("EnabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		dp.EnableAnonymous()
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, createThroughId, params)
	})

	t.Run("EnabledAnonymousWithPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel(throughCollection, throughId)
		dp.EnableAnonymous()
		dp.AddPermissionToGroup(1, permission)
		data := definitions.FqfieldData{
			throughField: []byte(strconv.Itoa(throughId)),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, createThroughId, params)
	})
}
