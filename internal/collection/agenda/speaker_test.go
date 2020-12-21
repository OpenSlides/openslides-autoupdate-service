package agenda_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/collection/agenda"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestSpeakerDelete(t *testing.T) {
	tdp := tests.NewTestDataProvider()
	tdp.AddUserToMeeting(1, 1) // Speaker user
	tdp.AddUserToMeeting(2, 1) // Manager user
	tdp.AddUserToGroup(2, 1, 3)
	tdp.AddUserToMeeting(3, 1) // Unprivileg user

	dp := dataprovider.DataProvider{External: tdp}
	s := agenda.NewSpeaker(dp)
	hs := new(tests.HandlerStoreMock)
	s.Connect(hs)
	delete := hs.WriteHandler["speaker.delete"]

	t.Run("delete self", func(t *testing.T) {
		tdp.AddBasicModel("speaker", 1)
		tdp.Set("speaker/1/user_id", "1")

		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}

		if _, err := delete.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error %v", err)
		}
	})

	t.Run("delete other", func(t *testing.T) {
		tdp.AddBasicModel("speaker", 1)
		tdp.Set("speaker/1/user_id", "1")

		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}

		if _, err := delete.IsAllowed(context.Background(), 2, payload); err != nil {
			t.Errorf("Got unexpected error %v", err)
		}
	})

	t.Run("delete other without perm", func(t *testing.T) {
		tdp.AddBasicModel("speaker", 1)
		tdp.Set("speaker/1/user_id", "1")

		payload := map[string]json.RawMessage{
			"id": []byte("1"),
		}

		if _, err := delete.IsAllowed(context.Background(), 3, payload); err == nil {
			t.Errorf("Expected error, got non")
		}
	})
}
