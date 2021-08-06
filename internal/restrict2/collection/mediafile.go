package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
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
	meetingID := datastore.Int(ctx, fetch.FetchIfExist, "mediafile/%d/meeting_id", mediafileID)
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

	usedAsLogo := datastore.Strings(ctx, fetch.FetchIfExist, "mediafile/%d/used_as_logo_$_in_meeting_id", mediafileID)
	usedAsFont := datastore.Strings(ctx, fetch.FetchIfExist, "mediafile/%d/used_as_font_$_in_meeting_id", mediafileID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching as logo and as font: %w", err)
	}
	if canSeeMeeting && (len(usedAsFont)+len(usedAsLogo) > 0) {
		return true, nil
	}

	if perms.Has(perm.ProjectorCanSee) {
		p7onIDs := datastore.Ints(ctx, fetch.FetchIfExist, "mediafile/%d/projection_ids", mediafileID)
		for _, p7onID := range p7onIDs {
			current := datastore.Int(ctx, fetch.Fetch, "projection/%d/current_projector_id", p7onID)
			if current != 0 {
				return true, nil
			}
		}

		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("checking projections: %w", err)
		}
	}

	if perms.Has(perm.MediafileCanSee) {
		public := datastore.Bool(ctx, fetch.FetchIfExist, "mediafile/%d/is_public", mediafileID)
		if public {
			return true, nil
		}

		inheritedGroups := datastore.Ints(ctx, fetch.FetchIfExist, "mediafile/%d/inherited_access_group_ids", mediafileID)
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

	losID := datastore.Int(ctx, fetch.FetchIfExist, "mediafile/%d/list_of_speakers_id", mediafileID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting list of speakers id: %w", err)
	}

	if losID == 0 {
		return false, nil
	}

	canSeeLOS, err := ListOfSpeakers{}.see(ctx, fetch, mperms, losID)
	if err != nil {
		return false, fmt.Errorf("can see los: %w", err)
	}
	return canSeeLOS, nil
}
