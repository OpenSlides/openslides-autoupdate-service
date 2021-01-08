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

		allowed, err := setState.IsAllowed(context.Background(), 1, payload)
		if err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		if !allowed {
			t.Errorf("Got false, expected true")
		}
	})

	t.Run("wrong state", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "false")
		allowed, err := setState.IsAllowed(context.Background(), 1, payload)
		if err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		if allowed {
			t.Errorf("Got true, expected false")
		}
	})

	t.Run("manager", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "false")
		tdp.AddPermissionToGroup(1, "motion.can_manage_metadata")

		allowed, err := setState.IsAllowed(context.Background(), 1, payload)
		if err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		if allowed {
			t.Errorf("Got true, expected false")
		}
	})
}

func TestMotionDelete(t *testing.T) {
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
	delete, ok := hs.WriteHandler["motion.delete"]
	if !ok {
		t.Fatalf("Unknown handler `motion.delete`")
	}

	payload := map[string]json.RawMessage{
		"id": []byte("1"),
	}

	t.Run("correct state", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "true")

		if _, err := delete.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}
	})

	t.Run("wrong state", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "false")
		allowed, err := delete.IsAllowed(context.Background(), 1, payload)
		if err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		if allowed {
			t.Errorf("Got true, expected false")
		}
	})

	t.Run("manager", func(t *testing.T) {
		tdp.Set("motion_state/1/allow_submitter_edit", "false")
		tdp.AddPermissionToGroup(1, "motion.can_manage")

		if _, err := delete.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("Got unexpected error: %v", err)
		}
	})
}

func TestMotionCreate(t *testing.T) {
	tdp := tests.NewTestDataProvider()
	tdp.AddUserToMeeting(1, 1)

	dp := dataprovider.DataProvider{External: tdp}
	m := user.NewMotion(dp)
	hs := new(tests.HandlerStoreMock)
	m.Connect(hs)
	create, ok := hs.WriteHandler["motion.create"]
	if !ok {
		t.Fatalf("Unknown handler `motion.create`")
	}

	t.Run("create simple fields without perm", func(t *testing.T) {
		tdp.Set("group/1/permissions", "[]")
		payload := map[string]json.RawMessage{
			"meeting_id": []byte("1"),
		}

		allowed, err := create.IsAllowed(context.Background(), 1, payload)
		if err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		if allowed {
			t.Errorf("Got true, expected false")
		}

	})

	t.Run("create simple fields with perm", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["motion.can_create"]`)
		payload := map[string]json.RawMessage{
			"meeting_id": []byte("1"),
		}

		if _, err := create.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("got unexpected error: %v", err)
		}
	})

	t.Run("create amentment fields with wrong perm", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["motion.can_create"]`)
		payload := map[string]json.RawMessage{
			"meeting_id": []byte("1"),
			"parent_id":  []byte("1"),
		}

		allowed, err := create.IsAllowed(context.Background(), 1, payload)
		if err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		if allowed {
			t.Errorf("Got true, expected false")
		}

	})

	t.Run("create amentment fields with perm", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["motion.can_create_amendment"]`)
		payload := map[string]json.RawMessage{
			"meeting_id": []byte("1"),
			"parent_id":  []byte("1"),
		}

		if _, err := create.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("got unexpected error: %v", err)
		}
	})

	t.Run("create privileg fields with can_create", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["motion.can_create"]`)
		payload := map[string]json.RawMessage{
			"meeting_id":    []byte("1"),
			"agenda_create": []byte("true"),
		}

		allowed, err := create.IsAllowed(context.Background(), 1, payload)
		if err != nil {
			t.Fatalf("Got unexpected error: %v", err)
		}
		if allowed {
			t.Errorf("Got true, expected false")
		}

	})

	t.Run("create privileg fields with can_manage", func(t *testing.T) {
		tdp.Set("group/1/permissions", `["motion.can_manage"]`)
		payload := map[string]json.RawMessage{
			"meeting_id":    []byte("1"),
			"agenda_create": []byte("true"),
		}

		if _, err := create.IsAllowed(context.Background(), 1, payload); err != nil {
			t.Errorf("got unexpected error: %v", err)
		}
	})
}
