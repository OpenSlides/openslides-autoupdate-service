package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Meeting handels restrictions of the collection meeting.
type Meeting struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m Meeting) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	case "B":
		return m.see
	case "C":
		return m.modeC
	case "D":
		return m.modeD
	}
	return nil
}

func (m Meeting) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, meetingID int) (bool, error) {
	enableAnonymous := fetch.Field().Meeting_EnableAnonymous(ctx, meetingID)
	if enableAnonymous {
		return true, nil
	}

	userIDs := fetch.Field().Meeting_UserIDs(ctx, meetingID)
	for _, id := range userIDs {
		if mperms.UserID() == id {
			return true, nil
		}
	}

	committeeID := fetch.Field().Meeting_CommitteeID(ctx, meetingID)
	userManagementLvl := fetch.Field().User_CommitteeManagementLevel(ctx, mperms.UserID(), committeeID)
	if userManagementLvl == "can_manage" {
		return true, nil
	}

	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching meeting/%d: %w", meetingID, err)
	}

	return false, nil
}

func (m Meeting) modeC(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, meetingID int) (bool, error) {
	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.MeetingCanSeeFrontpage), nil
}

func (m Meeting) modeD(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, meetingID int) (bool, error) {
	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.MeetingCanSeeLivestream), nil
}
