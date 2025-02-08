package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-go/datastore/dsfetch"
)

// MotionEditor handels restrictions of the collection motion_editor.
//
// The user can see a motion_editor if he has `motion.can_manage_metadata`
//
// Mode A: The user can see the motion editor.
type MotionEditor struct{}

// Name returns the collection name.
func (m MotionEditor) Name() string {
	return "motion_editor"
}

// MeetingID returns the meetingID for the object.
func (m MotionEditor) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionEditor_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionEditor) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionEditor) see(ctx context.Context, ds *dsfetch.Fetch, motionEditorIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, m, motionEditorIDs, perm.MotionCanManageMetadata)
}
