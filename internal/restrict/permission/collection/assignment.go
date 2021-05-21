package collection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// Assignment creates the handler for assignment related objects.
func Assignment(dp dataprovider.DataProvider) perm.ConnecterFunc {
	a := &assignment{dp}
	return func(s perm.HandlerStore) {
		s.RegisterAction("assignment_candidate.create", perm.ActionFunc(a.candidateCreate))
		s.RegisterAction("assignment_candidate.delete", perm.ActionFunc(a.candidateDelete))
	}
}

type assignment struct {
	dp dataprovider.DataProvider
}

func (a *assignment) candidateCreate(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	meetingID, err := a.dp.MeetingIDFromPayload(ctx, payload, "assignment", "assignment_id")
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	permissions, err := perm.New(ctx, a.dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("collecting permissions for user %d in meeting %d: %w", userID, meetingID, err)
	}

	if permissions.Has(perm.AssignmentCanManage) {
		return true, nil
	}

	var phase string
	if err := a.dp.Get(ctx, fmt.Sprintf("assignment/%s/phase", payload["assignment_id"]), &phase); err != nil {
		return false, fmt.Errorf("getting phase of assignment: %w", err)
	}

	if phase != "search" {
		perm.LogNotAllowedf("Assignment is already in phase %s. No new candidates allowed.", phase)
		return false, nil
	}

	var cid int
	if err := json.Unmarshal(payload["user_id"], &cid); err != nil {
		return false, fmt.Errorf("getting user_id from payload: %w", err)
	}

	requiredPerm := perm.AssignmentCanNominateOther
	if userID == cid {
		requiredPerm = perm.AssignmentCanNominateSelf
	}

	if permissions.Has(requiredPerm) {
		return true, nil
	}

	perm.LogNotAllowedf("User %d does not have the permission %s", userID, requiredPerm)
	return false, nil
}

func (a *assignment) candidateDelete(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	meetingID, err := a.dp.MeetingIDFromPayload(ctx, payload, "assignment", "assignment_id")
	if err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	permissions, err := perm.New(ctx, a.dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("collecting permissions for user %d in meeting %d: %w", userID, meetingID, err)
	}

	if permissions.Has(perm.AssignmentCanManage) {
		return true, nil
	}

	var phase string
	if err := a.dp.Get(ctx, fmt.Sprintf("assignment/%s/phase", payload["assignment_id"]), &phase); err != nil {
		return false, fmt.Errorf("getting phase of assignment: %w", err)
	}

	if phase != "search" {
		perm.LogNotAllowedf("Assignment is already in phase %S. You can not remove yourself anymore.", phase)
		return false, nil
	}

	var cid int
	if err := json.Unmarshal(payload["user_id"], &cid); err != nil {
		return false, fmt.Errorf("getting user_id from payload: %w", err)
	}

	if userID == cid && permissions.Has(perm.AssignmentCanNominateSelf) {
		return true, nil
	}

	return false, nil
}

func canSeeAssignmentCandidate(p *perm.Permission) bool {
	return p.Has(perm.AssignmentCanSee)
}
