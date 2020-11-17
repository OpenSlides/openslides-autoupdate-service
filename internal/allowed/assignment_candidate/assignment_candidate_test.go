package assignment_candidate_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/allowed"
	"github.com/OpenSlides/openslides-permission-service/internal/allowed/assignment_candidate"
	"github.com/OpenSlides/openslides-permission-service/internal/definitions"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestCreate(t *testing.T) {
	t.Run("ValidPermissionSelf", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment", 1)
		data := definitions.FqfieldData{
			"assignment_id": []byte("1"),
			"user_id":       []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_self")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, assignment_candidate.Create, params)
	})
	t.Run("ValidPermissionOther", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment", 1)
		data := definitions.FqfieldData{
			"assignment_id": []byte("1"),
			"user_id":       []byte("2"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_other")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, assignment_candidate.Create, params)
	})
	t.Run("InvalidPermissionSelf", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment", 1)
		data := definitions.FqfieldData{
			"assignment_id": []byte("1"),
			"user_id":       []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_other") // the wrong permission
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, assignment_candidate.Create, params)
	})
	t.Run("InvalidPermissionOther", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment", 1)
		data := definitions.FqfieldData{
			"assignment_id": []byte("1"),
			"user_id":       []byte("2"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_self") // the wrong permission
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, assignment_candidate.Create, params)
	})
	t.Run("NoUserId", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment", 1)
		data := definitions.FqfieldData{
			"assignment_id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_self")
		dp.AddPermissionToGroup(1, "assignments.can_nominate_other")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, assignment_candidate.Create, params)
	})
}

func TestDelete(t *testing.T) {
	t.Run("ValidPermissionSelf", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment_candidate", 1)
		dp.Set("assignment_candidate/1/user_id", "1")
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_self")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, assignment_candidate.Delete, params)
	})
	t.Run("ValidPermissionOther", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment_candidate", 1)
		dp.Set("assignment_candidate/1/user_id", "2")
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_other")
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsAllowed(t, assignment_candidate.Delete, params)
	})
	t.Run("InvalidPermissionSelf", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment_candidate", 1)
		dp.Set("assignment_candidate/1/user_id", "1")
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_other") // wrong permission
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, assignment_candidate.Delete, params)
	})
	t.Run("InvalidPermissionOther", func(t *testing.T) {
		dp := tests.NewTestDataProvider(context.Background())
		dp.AddBasicModel("assignment_candidate", 1)
		dp.Set("assignment_candidate/1/user_id", "2")
		data := definitions.FqfieldData{
			"id": []byte("1"),
		}
		dp.AddUserToMeeting(1, 1)
		dp.AddPermissionToGroup(1, "assignments.can_nominate_self") // wrong permission
		params := &allowed.IsAllowedParams{UserID: 1, Data: data, DataProvider: dp.GetDataprovider()}

		allowed.AssertIsNotAllowed(t, assignment_candidate.Delete, params)
	})
}
