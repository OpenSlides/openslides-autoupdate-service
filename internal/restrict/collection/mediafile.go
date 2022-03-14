package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Mediafile handels permissions for the collection mediafile.
//
// Everyone can see a medafile that belongs to the organization.
//
// The user can see a mediafile of a meeting if any of:
//     The user is an admin of the meeting.
//     The user can see the meeting and used_as_logo_$_in_meeting_id or used_as_font_$_in_meeting_id is not empty.
//     The user has projector.can_see and there exists a mediafile/projection_ids with projection/current_projector_id set.
//     The user has mediafile.can_manage.
//     The user has mediafile.can_see and either:
//         mediafile/is_public is true, or
//         The user has groups in common with meeting/inherited_access_group_ids.
//
// Mode A: The user can see the mediafile.
type Mediafile struct{}

// Modes returns the field modes for the collection mediafile.
func (m Mediafile) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m Mediafile) see(ctx context.Context, ds *datastore.Request, mperms *perm.MeetingPermission, mediafileID int) (bool, error) {
	genericOwnerID, err := ds.Mediafile_OwnerID(mediafileID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("fetching owner_id of mediafile %d: %w", mediafileID, err)
	}

	collection, rawID, found := cutgo118(genericOwnerID, "/")
	if !found {
		return false, fmt.Errorf("invalid ownerID: %s", genericOwnerID)
	}

	ownerID, err := strconv.Atoi(rawID)
	if err != nil {
		return false, fmt.Errorf("invalid id part of ownerID: %s", genericOwnerID)
	}

	if collection == "organization" {
		return true, nil
	}

	perms, err := mperms.Meeting(ctx, ownerID)
	if err != nil {
		return false, fmt.Errorf("getting perms for meeting %d: %w", ownerID, err)
	}

	if perms.IsAdmin() {
		return true, nil
	}

	canSeeMeeting, err := Meeting{}.see(ctx, ds, mperms, ownerID)
	if err != nil {
		return false, fmt.Errorf("can see meeting %d: %w", ownerID, err)
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

	if perms.Has(perm.MediafileCanManage) {
		return true, nil
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

// cutgo118 from go 1.18. Replace me after the 1.18 release.
func cutgo118(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}
