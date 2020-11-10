package group_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/group"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestCreate(t *testing.T) {
	t.Run("ValidPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id":  []byte("1"),
			"permissions": []byte(`["motions.can_manage", "assignments.can_nominate_other"]`),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, group.Create, params)
	})
	t.Run("EmptyPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, group.Create, params)
	})
	t.Run("InvalidPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id":  []byte("1"),
			"permissions": []byte(`["motions.can_do_it"]`),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, group.Create, params)
	})
	t.Run("InvalidJson", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"meeting_id":  []byte("1"),
			"permissions": []byte(`{"key": 123}`),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, group.Create, params)
	})
}

func TestSetPermission(t *testing.T) {
	t.Run("ValidPermission", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel("group", dp)
		data := definitions.FqfieldData{
			"id":         []byte("1"),
			"permission": []byte(`"motions.can_manage"`),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, group.SetPermission, params)
	})
	t.Run("EmptyPermission", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel("group", dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, group.SetPermission, params)
	})
	t.Run("InvalidPermission", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel("group", dp)
		data := definitions.FqfieldData{
			"id":         []byte("1"),
			"permission": []byte("agenda.not_valid"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, group.SetPermission, params)
	})
	t.Run("InvalidJson", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		allowed.AddBasicModel("group", dp)
		data := definitions.FqfieldData{
			"id":          []byte("1"),
			"permissions": []byte(`{"key": 123}`),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "users.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, group.SetPermission, params)
	})
}
