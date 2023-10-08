package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// PointOfOrderCategory handels restriction for the point_of_order_category collection.
//
// The user can see a point_of_order_category, if he can see the linked meeting.
//
// Mode A: The user can see the point_of_order_category.
type PointOfOrderCategory struct{}

// Name returns the collection name.
func (p PointOfOrderCategory) Name() string {
	return "point_of_order_category"
}

// MeetingID returns the meetingID for the object.
func (p PointOfOrderCategory) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.PointOfOrderCategory_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the field restriction for each mode.
func (p PointOfOrderCategory) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	}
	return nil
}

func (p PointOfOrderCategory) see(ctx context.Context, fetcher *dsfetch.Fetch, pointOfOrderCategoryIDs []int) ([]attribute.Func, error) {
	// TODO: It would be faster to call Collection().Mode(B) with all meeting IDs at once.
	return byMeeting(ctx, fetcher, p, pointOfOrderCategoryIDs, func(meetingID int, ids []int) ([]attribute.Func, error) {
		meetingAttr, err := Collection(ctx, "meeting").Modes("B")(ctx, fetcher, []int{meetingID})
		if err != nil {
			return nil, fmt.Errorf("checking motion.see: %w", err)
		}

		return attributeFuncList(len(ids), meetingAttr[0]), nil
	})
}
