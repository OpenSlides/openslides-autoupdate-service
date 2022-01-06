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

func (m Meeting) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, meetingID int) (bool, error) {
	enableAnonymous, err := ds.Meeting_EnableAnonymous(meetingID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("checking enabled anonymous: %w", err)
	}
	if enableAnonymous {
		return true, nil
	}

	if mperms.UserID() == 0 {
		return false, nil
	}

	oml, err := perm.HasOrganizationManagementLevel(ctx, ds, mperms.UserID(), perm.OMLCanManageOrganization)
	if err != nil {
		return false, fmt.Errorf("checking organization management level: %w", err)
	}

	if oml {
		return true, nil
	}

	userIDs := ds.Meeting_UserIDs(meetingID).ErrorLater(ctx)
	for _, id := range userIDs {
		if mperms.UserID() == id {
			return true, nil
		}
	}

	committeeID := ds.Meeting_CommitteeID(meetingID).ErrorLater(ctx)
	userManagementLvl := ds.User_CommitteeManagementLevel(mperms.UserID(), committeeID).ErrorLater(ctx)
	if userManagementLvl == "can_manage" {
		return true, nil
	}

	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("fetching meeting/%d: %w", meetingID, err)
	}

	return false, nil
}

func (m Meeting) modeC(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, meetingID int) (bool, error) {
	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.MeetingCanSeeFrontpage), nil
}

func (m Meeting) modeD(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, meetingID int) (bool, error) {
	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	return perms.Has(perm.MeetingCanSeeLivestream), nil
}
