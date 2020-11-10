package topic_test

import (
	"context"
	"errors"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/definitions"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/topic"

	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func assertUpdateIsNotAllowed(t *testing.T, params *allowed.IsAllowedParams) {
	addition, err := topic.Update(params)
	if nil != addition {
		t.Errorf("Expected to fail without an addition: %s", addition)
	}

	if nil == err {
		t.Errorf("Expected to fail (reason must be set).")
	} else {
		var clientError interface {
			Type() string
		}
		if !errors.As(err, &clientError) || clientError.Type() != "ClientError" {
			t.Errorf("Expected to fail with a client error, not %v", err)
		}
	}
}

func assertUpdateIsAllowed(t *testing.T, params *allowed.IsAllowedParams) {
	addition, err := topic.Update(params)
	if nil != addition {
		t.Errorf("Expected to fail without an addition: %s", addition)
	}
	if nil != err {
		t.Errorf("Expected to fail without an error (error: %s)", err)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("UnknownUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsNotAllowed(t, params)
	})

	t.Run("SuperadminRole", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{} // No meeting id needed, it is always possible.
		dp.AddUserWithSuperadminRole(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsAllowed(t, params)
	})

	t.Run("NoId", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		data := definitions.FqfieldData{}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsNotAllowed(t, params)
	})

	t.Run("UserNotInMeeting", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		addBasicTopic(dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUser(1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsNotAllowed(t, params)
	})

	t.Run("AdminUser", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		addBasicTopic(dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserWithAdminGroupToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsAllowed(t, params)
	})

	t.Run("User", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		addBasicTopic(dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "agenda.can_manage")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsAllowed(t, params)
	})

	t.Run("UserNoPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		addBasicTopic(dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsNotAllowed(t, params)
	})

	t.Run("InvaldFields", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddUserWithSuperadminRole(1)
		data := definitions.FqfieldData{
			"not_allowed": []byte("some value"),
		}
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsNotAllowed(t, params)
	})

	t.Run("DisabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		addBasicTopic(dp)
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsNotAllowed(t, params)
	})

	t.Run("EnabledAnonymous", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		addBasicTopic(dp)
		dp.EnableAnonymous()
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsNotAllowed(t, params)
	})

	t.Run("EnabledAnonymousWithPermissions", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		addBasicTopic(dp)
		dp.EnableAnonymous()
		dp.AddPermissionToGroup(1, "agenda.can_manage")
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		params := &allowed.IsAllowedParams{UserID: 0, Data: data, DataProvider: dp.GetDataprovider()}

		assertUpdateIsAllowed(t, params)
	})
}
