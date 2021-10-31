package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Mediafile handels permissions for the collection mediafile.
type Mediafile struct{}

// Modes returns the field modes for the collection mediafile.
func (m Mediafile) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.modeA
	case "B":
		return m.see
	}
	return nil
}

func (m Mediafile) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, mediafileID int) (bool, error) {
	meetingID, err := ds.Mediafile_MeetingID(mediafileID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching meeting_id of mediafile %d: %w", mediafileID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	if perms.IsAdmin() {
		return true, nil
	}

	canSeeMeeting, err := Meeting{}.see(ctx, ds, mperms, meetingID)
	if err != nil {
		return false, fmt.Errorf("can see meeting %d: %w", meetingID, err)
	}

	usedAsLogo := ds.Mediafile_UsedAsLogoInMeetingIDTmpl(mediafileID).ErrorLater(ctx)
	usedAsFont := ds.Mediafile_UsedAsFontInMeetingIDTmpl(mediafileID).ErrorLater(ctx)
	if err := ds.Err(); err != nil {
		return false, fmt.Errorf("fetching as logo and as font: %w", err)
	}
	if canSeeMeeting && (len(usedAsFont)+len(usedAsLogo) > 0) {
		return true, nil
	}

	if perms.Has(perm.ProjectorCanSee) {
		p7onIDs := ds.Mediafile_ProjectionIDs(mediafileID).ErrorLater(ctx)
		for _, p7onID := range p7onIDs {
			if _, exist := ds.Projection_CurrentProjectorID(p7onID).ErrorLater(ctx); exist {
				return true, nil
			}
		}

		if err := ds.Err(); err != nil {
			return false, fmt.Errorf("checking projections: %w", err)
		}
	}

	if perms.Has(perm.MediafileCanSee) {
		public := ds.Mediafile_IsPublic(mediafileID).ErrorLater(ctx)
		if public {
			return true, nil
		}

		inheritedGroups := ds.Mediafile_InheritedAccessGroupIDs(mediafileID).ErrorLater(ctx)
		for _, id := range inheritedGroups {
			if perms.InGroup(id) {
				return true, nil
			}
		}

		if err := ds.Err(); err != nil {
			return false, fmt.Errorf("checking can see conditions: %w", err)
		}
	}
	return false, nil
}

func (m Mediafile) modeA(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, mediafileID int) (bool, error) {
	canSee, err := m.see(ctx, ds, mperms, mediafileID)
	if err != nil {
		return false, fmt.Errorf("see property: %w", err)
	}

	if canSee {
		return true, nil
	}

	losID, exist, err := ds.Mediafile_ListOfSpeakersID(mediafileID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting list of speakers id: %w", err)
	}

	if !exist {
		return false, nil
	}

	canSeeLOS, err := ListOfSpeakers{}.see(ctx, ds, mperms, losID)
	if err != nil {
		return false, fmt.Errorf("can see los: %w", err)
	}
	return canSeeLOS, nil
}
