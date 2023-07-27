package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Group handels restrictions of the collection group.
//
// The user can see a group, if the user can see the group's meeting.
//
// Mode A: The user can see the group.
type Group struct{}

// Name returns the collection name.
func (g Group) Name() string {
	return "group"
}

// MeetingID returns the meetingID for the object.
func (g Group) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Group_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching meeting id of group %d: %w", id, err)
	}

	return meetingID, true, nil
}

// Modes returns the field restricters for the collection group.
func (g Group) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return g.see
	}
	return nil
}

func (g Group) see(ctx context.Context, fetcher *dsfetch.Fetch, groupIDs []int) ([]Tuple, error) {
	groupToMeeting := make(map[int]int, len(groupIDs))
	meetingIDs := set.New[int]()
	for _, groupID := range groupIDs {
		meetingID := fetcher.Group_MeetingID(groupID).ErrorLater(ctx)
		groupToMeeting[groupID] = meetingID
		meetingIDs.Add(meetingID)
	}
	if err := fetcher.Err(); err != nil {
		return nil, fmt.Errorf("getting meetingIDs: %w", err)
	}

	motionAttrList, err := Collection(ctx, "meeting").Modes("B")(ctx, fetcher, meetingIDs.List())
	if err != nil {
		return nil, fmt.Errorf("checking motion.see: %w", err)
	}

	indexMeetingToAttr := make(map[int]attribute.Func, len(motionAttrList))
	for _, motionAttr := range motionAttrList {
		indexMeetingToAttr[motionAttr.Key.ID] = motionAttr.Value
	}

	results := make([]Tuple, len(groupIDs))
	for i, id := range groupIDs {
		results[i].Key = modeKey(g, id, "A")
		motionID := groupToMeeting[id]
		results[i].Value = indexMeetingToAttr[motionID]
	}
	return results, nil
}
