package collection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Assignment creates the handler for assignment related objects.
func Assignment(dp dataprovider.DataProvider) perm.ConnecterFunc {
	a := &assignment{dp}
	return func(s perm.HandlerStore) {
		s.RegisterWriteHandler("assignment_candidate.create", perm.WriteCheckerFunc(a.candidateCreate))
		s.RegisterWriteHandler("assignment_candidate.delete", perm.WriteCheckerFunc(a.candidateDelete))
	}
}

type assignment struct {
	dp dataprovider.DataProvider
}

func (a *assignment) candidateCreate(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	meetingID, err := a.dp.MeetingIDFromPayload(ctx, payload, "assignment", "assignment_id")
	if err != nil {
		return nil, fmt.Errorf("getting meetingID: %w", err)
	}

	permissions, err := perm.New(ctx, a.dp, userID, meetingID)
	if err != nil {
		return nil, fmt.Errorf("collecting permissions for user %d in meeting %d: %w", userID, meetingID, err)
	}

	if permissions.Has("assignment.can_manage") {
		return nil, nil
	}

	var phase int
	if err := a.dp.Get(ctx, fmt.Sprintf("assignment/%s/phase", payload["assignment_id"]), &phase); err != nil {
		return nil, fmt.Errorf("getting phase of assignment: %w", err)
	}

	if phase != 0 {
		return nil, perm.NotAllowedf("Assignment is already in phase %d. No new candidates allowed.", phase)
	}

	var cid int
	if err := json.Unmarshal(payload["user_id"], &cid); err != nil {
		return nil, fmt.Errorf("getting user_id from payload: %w", err)
	}

	requiredPerm := "assignment.can_nominate_other"
	if userID == cid {
		requiredPerm = "assignment.can_nominate_self"
	}

	if permissions.Has(requiredPerm) {
		return nil, nil
	}

	return nil, perm.NotAllowedf("User %d does not have the permission %s", userID, requiredPerm)
}

func (a *assignment) candidateDelete(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	meetingID, err := a.dp.MeetingIDFromPayload(ctx, payload, "assignment", "assignment_id")
	if err != nil {
		return nil, fmt.Errorf("getting meetingID: %w", err)
	}

	permissions, err := perm.New(ctx, a.dp, userID, meetingID)
	if err != nil {
		return nil, fmt.Errorf("collecting permissions for user %d in meeting %d: %w", userID, meetingID, err)
	}

	if permissions.Has("assignment.can_manage") {
		return nil, nil
	}

	var phase int
	if err := a.dp.Get(ctx, fmt.Sprintf("assignment/%s/phase", payload["assignment_id"]), &phase); err != nil {
		return nil, fmt.Errorf("getting phase of assignment: %w", err)
	}

	if phase != 0 {
		return nil, perm.NotAllowedf("Assignment is already in phase %d. You can not remove yourself anymore.", phase)
	}

	var cid int
	if err := json.Unmarshal(payload["user_id"], &cid); err != nil {
		return nil, fmt.Errorf("getting user_id from payload: %w", err)
	}

	if userID == cid && permissions.Has("assignment.can_nominate_self") {
		return nil, nil
	}

	return nil, perm.NotAllowedf("Bad boy")
}
