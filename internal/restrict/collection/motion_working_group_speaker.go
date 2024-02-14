package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// MotionWorkingGroupSpeaker handels restrictions of the collection motion_working_group_speaker.
//
// The user can see a motion_working_group_speaker level if he has `motion.can_see`
//
// Mode A: The user can see the motion working group speaker.
type MotionWorkingGroupSpeaker struct{}

// Name returns the collection name.
func (m MotionWorkingGroupSpeaker) Name() string {
	return "motion_working_group_speaker"
}

// MeetingID returns the meetingID for the object.
func (m MotionWorkingGroupSpeaker) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MotionWorkingGroupSpeaker_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionWorkingGroupSpeaker) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	case "B":
		return never // TODO: Remove me after the fix in the backend
	}
	return nil
}

func (m MotionWorkingGroupSpeaker) see(ctx context.Context, ds *dsfetch.Fetch, motionWorkingGroupSpeakerIDs ...int) ([]int, error) {
	return meetingPerm(ctx, ds, m, motionWorkingGroupSpeakerIDs, perm.MotionCanSee)
}
