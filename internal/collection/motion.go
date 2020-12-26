package collection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Motion handels permissions of motions objects.
type Motion struct {
	dp dataprovider.DataProvider
}

// NewMotion initializes a motion.
func NewMotion(dp dataprovider.DataProvider) *Motion {
	return &Motion{
		dp: dp,
	}
}

// Connect registers the Motion handlers.
func (m *Motion) Connect(s perm.HandlerStore) {
	s.RegisterWriteHandler("motion.set_state", perm.WriteCheckerFunc(m.setState))
}

func (m *Motion) setState(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	motionFQID := fmt.Sprintf("motion/%s", payload["id"])
	meetingID, err := m.dp.MeetingFromModel(ctx, motionFQID)
	if err != nil {
		return nil, fmt.Errorf("getting meeting for %s: %w", motionFQID, err)
	}

	isManager, err := perm.IsAllowed(perm.EnsurePerm(ctx, m.dp, userID, meetingID, "motion.can_manage_metadata"))
	if err != nil {
		return nil, fmt.Errorf("checking meta manager permission: %w", err)
	}

	if isManager {
		return nil, nil
	}

	var submitterIDs []int
	if err := m.dp.Get(ctx, motionFQID+"/submitter_ids", &submitterIDs); err != nil {
		return nil, fmt.Errorf("getting submitter ids: %w", err)
	}

	var isSubmitter bool
	for _, sid := range submitterIDs {
		var sUserID int
		if err := m.dp.Get(ctx, fmt.Sprintf("motion_submitter/%d/user_id", sid), &sUserID); err != nil {
			return nil, fmt.Errorf("getting userid of sumitter %d: %w", sid, err)
		}
		if sUserID == userID {
			isSubmitter = true
			break
		}
	}

	if !isSubmitter {
		return nil, perm.NotAllowedf("User %d is not a manager and not a submitter of %s", userID, motionFQID)
	}

	var stateID int
	if err := m.dp.Get(ctx, motionFQID+"/state_id", &stateID); err != nil {
		return nil, fmt.Errorf("getting stateID: %w", err)
	}

	var allowSubmitterEdit bool
	if err := m.dp.Get(ctx, fmt.Sprintf("motion_state/%d/allow_submitter_edit", stateID), &allowSubmitterEdit); err != nil {
		return nil, fmt.Errorf("getting allow_submitter_edit: %w", err)
	}

	if !allowSubmitterEdit {
		return nil, perm.NotAllowedf("Motion state does not allow submitter edites")
	}

	return nil, nil
}
