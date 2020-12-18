package assignment_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/collection/assignment"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestCreate(t *testing.T) {
	t.Run("ValidPermissionSelf", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment", 1)
		payload := map[string]json.RawMessage{
			"assignment_id": []byte("1"),
			"user_id":       []byte("1"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_self")
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.create"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}
	})

	t.Run("ValidPermissionOther", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment", 1)
		payload := map[string]json.RawMessage{
			"assignment_id": []byte("1"),
			"user_id":       []byte("2"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_other")
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.create"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}
	})

	t.Run("InvalidPermissionSelf", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment", 1)
		payload := map[string]json.RawMessage{
			"assignment_id": []byte("1"),
			"user_id":       []byte("1"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_other") // the wrong permission
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.create"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err == nil {
			t.Errorf("Got no error, expected one")
		}
	})

	t.Run("InvalidPermissionOther", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment", 1)
		payload := map[string]json.RawMessage{
			"assignment_id": []byte("1"),
			"user_id":       []byte("2"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_self") // the wrong permission
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.create"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err == nil {
			t.Errorf("Got no error, expected one")
		}
	})

	t.Run("NoUserId", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment", 1)
		payload := map[string]json.RawMessage{
			"assignment_id": []byte("1"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_self")
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_other")
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.create"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err == nil {
			t.Errorf("Got no error, expected one")
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("ValidPermissionSelf", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment_candidate", 1)
		tdp.Set("assignment_candidate/1/user_id", "1")
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_self")
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.delete"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}
	})

	t.Run("ValidPermissionOther", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment_candidate", 1)
		tdp.Set("assignment_candidate/1/user_id", "2")
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_other")
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.delete"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}
	})

	t.Run("InvalidPermissionSelf", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment_candidate", 1)
		tdp.Set("assignment_candidate/1/user_id", "1")
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_other") // wrong permission
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.delete"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err == nil {
			t.Errorf("Got no error, expected one")
		}
	})

	t.Run("InvalidPermissionOther", func(t *testing.T) {
		tdp := tests.NewTestDataProvider()
		tdp.AddBasicModel("assignment_candidate", 1)
		tdp.Set("assignment_candidate/1/user_id", "2")
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}
		tdp.AddUserToMeeting(1, 1)
		tdp.AddPermissionToGroup(1, "assignments.can_nominate_self") // wrong permission
		dp := dataprovider.DataProvider{External: tdp}
		permMock := new(tests.HandlerStoreMock)
		assignment.NewCandidate(dp).Connect(permMock)
		handler := permMock.WriteHandler["assignment_candidate.delete"]

		if _, err := handler.IsAllowed(context.Background(), 1, payload); err == nil {
			t.Errorf("Got no error, expected one")
		}
	})
}
