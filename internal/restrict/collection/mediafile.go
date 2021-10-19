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

func (m Mediafile) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, mediafileID int) (bool, error) {
	meetingID := fetch.Field().Mediafile_MeetingID(ctx, mediafileID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching meeting_id of mediafile %d: %w", mediafileID, err)
	}

	perms, err := mperms.Meeting(ctx, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
	}

	if perms.IsAdmin() {
		return true, nil
	}

	canSeeMeeting, err := Meeting{}.see(ctx, fetch, mperms, meetingID)
	if err != nil {
		return false, fmt.Errorf("can see meeting %d: %w", meetingID, err)
	}

	usedAsLogo := fetch.Field().Mediafile_UsedAsLogoInMeetingIDTmpl(ctx, mediafileID)
	usedAsFont := fetch.Field().Mediafile_UsedAsFontInMeetingIDTmpl(ctx, mediafileID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching as logo and as font: %w", err)
	}
	if canSeeMeeting && (len(usedAsFont)+len(usedAsLogo) > 0) {
		return true, nil
	}

	if perms.Has(perm.ProjectorCanSee) {
		p7onIDs := fetch.Field().Mediafile_ProjectionIDs(ctx, mediafileID)
		for _, p7onID := range p7onIDs {
			if _, exist := fetch.Field().Projection_CurrentProjectorID(ctx, p7onID); exist {
				return true, nil
			}
		}

		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("checking projections: %w", err)
		}
	}

	if perms.Has(perm.MediafileCanSee) {
		public := fetch.Field().Mediafile_IsPublic(ctx, mediafileID)
		if public {
			return true, nil
		}

		inheritedGroups := fetch.Field().Mediafile_InheritedAccessGroupIDs(ctx, mediafileID)
		for _, id := range inheritedGroups {
			if perms.InGroup(id) {
				return true, nil
			}
		}

		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("checking can see conditions: %w", err)
		}
	}
	return false, nil
}

func (m Mediafile) modeA(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, mediafileID int) (bool, error) {
	canSee, err := m.see(ctx, fetch, mperms, mediafileID)
	if err != nil {
		return false, fmt.Errorf("see property: %w", err)
	}

	if canSee {
		return true, nil
	}

	losID, exist := fetch.Field().Mediafile_ListOfSpeakersID(ctx, mediafileID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting list of speakers id: %w", err)
	}

	if !exist {
		return false, nil
	}

	canSeeLOS, err := ListOfSpeakers{}.see(ctx, fetch, mperms, losID)
	if err != nil {
		return false, fmt.Errorf("can see los: %w", err)
	}
	return canSeeLOS, nil
}
