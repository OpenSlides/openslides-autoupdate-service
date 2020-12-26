package collection_test

import (
	"context"
	"encoding/json"
	"testing"

	user "github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/tests"
)

func TestMotionSetState(t *testing.T) {
	tdp := tests.NewTestDataProvider()
	tdp.AddUserToMeeting(1, 1)

	tdp.AddBasicModel("motion", 1)
	tdp.Set("motion/1/submitter_ids", "[1]")
	tdp.Set("motion/1/state_id", "1")

	tdp.AddBasicModel("motion_submitter", 1)
	tdp.Set("motion_submitter/1/user_id", "1")
	tdp.Set("motion_submitter/1/motion_id", "1")

	tdp.AddBasicModel("motion_state", 1)
	tdp.Set("motion_state/1/motion_ids", "[1]")
	tdp.Set("motion_state/1/allow_submitter_edit", "true")

	dp := dataprovider.DataProvider{External: tdp}
	m := user.NewMotion(dp)
	hs := new(tests.HandlerStoreMock)
	m.Connect(hs)
	setState, ok := hs.WriteHandler["motion.set_state"]
	if !ok {
		t.Fatalf("Unknown handler `motion.set_state`")
	}

	payload := map[string]json.RawMessage{
		"id": []byte("1"),
	}

	t.Run("correct state", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "true")

		if _, err := setState.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %w", err)
		}
	})

	t.Run("wrong state", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "false")
		_, err := setState.IsAllowed(context.Background(), 1, payload)
		assertNotAllowed(t, err)
	})

	t.Run("manager", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "false")
		tdp.AddPermissionToGroup(1, "motion.can_manage_metadata")

		if _, err := setState.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %w", err)
		}
	})
}
