package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionStatuteParagraph handels restrictions of the collection motion_statute_paragraph.
type MotionStatuteParagraph struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionStatuteParagraph) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionStatuteParagraph) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, MotionStatuteParagraphID int) (bool, error) {
	meetingID := fetch.Field().MotionStatuteParagraph_MeetingID(ctx, MotionStatuteParagraphID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting meetingID: %w", err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permission: %w", err)
	}

	return perms.Has(perm.MotionCanSee), nil
}
