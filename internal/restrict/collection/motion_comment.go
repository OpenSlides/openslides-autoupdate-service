package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
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

func (m MotionComment) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, motionCommentID int) (bool, error) {
	motionID := ds.MotionComment_MotionID(motionCommentID).ErrorLater(ctx)
	commentSectionID := ds.MotionComment_SectionID(motionCommentID).ErrorLater(ctx)
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("getting motion id and comment section id: %w", err)
	}

	seeMotion, err := Motion{}.see(ctx, ds, mperms, motionID)
	if err != nil {
		return false, fmt.Errorf("checking motion %d can see: %w", motionID, err)
	}

	seeSection, err := MotionCommentSection{}.see(ctx, ds, mperms, commentSectionID)
	if err != nil {
		return false, fmt.Errorf("checking motion comment section %d can see: %w", commentSectionID, err)
	}

	return seeMotion && seeSection, nil
}
