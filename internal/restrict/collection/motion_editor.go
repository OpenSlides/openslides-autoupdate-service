package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionEditor handels restrictions of the collection motion_editor.
//
// The user can see a motion_editor level if he has `motion.can_see`
//
// Mode A: The user can see the motion editor.
type MotionEditor struct{}

// Name returns the collection name.
func (e MotionEditor) Name() string {
	return "motion_editor"
}

// MeetingID returns the meetingID for the object.
func (s MotionEditor) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionEditor_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (s MotionEditor) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return s.see
	case "B":
		return never // TODO: Remove me after the fix in the backend
	}
	return nil
}

func (s MotionEditor) see(ctx context.Context, ds *dsfetch.Fetch, motionEditorIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, s, motionEditorIDs, perm.MotionCanSee)
}
