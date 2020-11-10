package allowed_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

var collection = "test_collection"
var permission = "test_collection.permission"
var create = allowed.BuildCreate([]string{"meeting_id"}, permission)
var modify = allowed.BuildModify([]string{"id"}, collection, permission)

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
		allowed.AddBasicModel(collection, dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUser(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("UserIsCommitteemanager", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel(collection, dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToCommitteeAsManager(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modify, params)
	})

	t.Run("AdminUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel(collection, dp)
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
		allowed.AddBasicModel(collection, dp)
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
		allowed.AddBasicModel(collection, dp)
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
		allowed.AddBasicModel(collection, dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("EnabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel(collection, dp)
		dp.EnableAnonymous()
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, modify, params)
	})

	t.Run("EnabledAnonymousWithPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel(collection, dp)
		dp.EnableAnonymous()
		dp.AddPermissionToGroup(1, permission)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, modify, params)
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
