package user_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/collection/user"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestPersonalNote(t *testing.T) {
	tdp := tests.NewTestDataProvider()
	tdp.AddUserToMeeting(1, 1) // Speaker user
	tdp.AddUserToMeeting(2, 1) // Unprivileg user

	tdp.AddBasicModel("personal_note", 1)
	tdp.Set("personal_note/1/user_id", "1")

	dp := dataprovider.DataProvider{External: tdp}
	n := user.NewPersonalNote(dp)
	hs := new(tests.HandlerStoreMock)
	n.Connect(hs)
	update := hs.WriteHandler["personal_note.update"]

	t.Run("update own note", func(t *testing.T) {
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}

		if _, err := update.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error %v", err)
		}
	})

	t.Run("update other note", func(t *testing.T) {
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}

		if _, err := update.IsAllowed(context.Background(), 2, payload); err == nil {
			t.Errorf("Expected an error")
		}
	})

	delete := hs.WriteHandler["personal_note.delete"]

	t.Run("delete own note", func(t *testing.T) {
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}

		if _, err := delete.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error %v", err)
		}
	})

	t.Run("delete other note", func(t *testing.T) {
		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}

		if _, err := delete.IsAllowed(context.Background(), 2, payload); err == nil {
			t.Errorf("Expected an error")
		}
	})
}
