package collection_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestSpeaker(t *testing.T) {
	tdp := tests.NewTestDataProvider()
	tdp.AddUserToMeeting(1, 1) // Speaker user
	tdp.AddUserToMeeting(2, 1) // Manager user
	tdp.AddUserToGroup(2, 1, 3)
	tdp.AddUserToMeeting(3, 1) // Unprivileg user

	dp := dataprovider.DataProvider{External: tdp}
	s := collection.ListOfSpeaker(dp)
	hs := new(tests.HandlerStoreMock)
	s.Connect(hs)
	delete := hs.WriteHandler["speaker.delete"]
	read := hs.ReadHandler["speaker"]

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

	t.Run("read", func(t *testing.T) {
		tdp.AddBasicModel("speaker", 1)
		tdp.Set("speaker/1/user_id", "1")
		tdp.Set("speaker/1/meeting_id", "1")
		tdp.Set("speaker/2/user_id", "999")
		tdp.AddPermissionToGroup(1, "agenda.can_see_list_of_speakers")
		fqfields := mustFQfields(
			"speaker/1/id",
			"speaker/1/user_id",
			"speaker/2/id",
			"speaker/2/user_id",
		)

		result := make(map[string]bool)
		if err := read.RestrictFQFields(context.Background(), 1, fqfields, result); err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		checkRead(t, result, "speaker/1/id", "speaker/1/user_id")
	})
}
