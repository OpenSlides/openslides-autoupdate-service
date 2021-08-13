package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// MotionComment handels restrictions of the collection motion_comment.
type MotionComment struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m MotionComment) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MotionComment) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionCommentID int) (bool, error) {
	motionID := fetch.Field().MotionComment_MotionID(ctx, motionCommentID)
	commentSectionID := fetch.Field().MotionComment_SectionID(ctx, motionCommentID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting motion id and comment section id: %w", err)
	}

	seeMotion, err := Motion{}.see(ctx, fetch, mperms, motionID)
	if err != nil {
		return false, fmt.Errorf("checking motion %d can see: %w", motionID, err)
	}

	seeSection, err := MotionCommentSection{}.see(ctx, fetch, mperms, commentSectionID)
	if err != nil {
		return false, fmt.Errorf("checking motion comment section %d can see: %w", commentSectionID, err)
	}

	return seeMotion && seeSection, nil
}
