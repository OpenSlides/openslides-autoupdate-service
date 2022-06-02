package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Mediafile handels permissions for the collection mediafile.
//
// Every logged in user can see a medafile that belongs to the organization.
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

// MeetingID returns the meetingID for the object.
func (m Mediafile) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	genericOwnerID, err := ds.Mediafile_OwnerID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching owner_id of mediafile %d: %w", id, err)
	}

	collection, rawID, found := strings.Cut(genericOwnerID, "/")
	if !found {
		// TODO LAST ERROR
		return 0, false, fmt.Errorf("invalid ownerID: %s", genericOwnerID)
	}

	if collection != "meeting" {
		return 0, false, nil
	}

	ownerID, err := strconv.Atoi(rawID)
	if err != nil {
		// TODO LAST ERROR
		return 0, false, fmt.Errorf("invalid id part of ownerID: %s", genericOwnerID)
	}

	return ownerID, true, nil
}

// Modes returns the field modes for the collection mediafile.
func (m Mediafile) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m Mediafile) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, mediafileIDs ...int) ([]int, error) {
	return eachContentObjectCollection(ctx, ds.Mediafile_OwnerID, mediafileIDs, func(collection string, ownerID int, ids []int) ([]int, error) {
		if collection == "organization" && mperms.UserID() != 0 {
			return ids, nil
		}

		var allowed []int
		for _, mediafileID := range ids {
			perms, err := mperms.Meeting(ctx, ownerID)
			if err != nil {
				return nil, fmt.Errorf("getting perms for meeting %d: %w", ownerID, err)
			}

			if perms.IsAdmin() {
				allowed = append(allowed, mediafileID)
				continue
			}

			canSeeMeeting, err := Meeting{}.see(ctx, ds, mperms, ownerID)
			if err != nil {
				return nil, fmt.Errorf("can see meeting %d: %w", ownerID, err)
			}

			usedAsLogo := ds.Mediafile_UsedAsLogoInMeetingIDTmpl(mediafileID).ErrorLater(ctx)
			usedAsFont := ds.Mediafile_UsedAsFontInMeetingIDTmpl(mediafileID).ErrorLater(ctx)
			if err := ds.Err(); err != nil {
				return nil, fmt.Errorf("fetching as logo and as font: %w", err)
			}
			if canSeeMeeting && (len(usedAsFont)+len(usedAsLogo) > 0) {
				allowed = append(allowed, mediafileID)
				continue
			}

			if perms.Has(perm.ProjectorCanSee) {
				p7onIDs := ds.Mediafile_ProjectionIDs(mediafileID).ErrorLater(ctx)
				for _, p7onID := range p7onIDs {
					if _, exist := ds.Projection_CurrentProjectorID(p7onID).ErrorLater(ctx); exist {
						allowed = append(allowed, mediafileID)
						continue
					}
				}

				if err := ds.Err(); err != nil {
					return nil, fmt.Errorf("checking projections: %w", err)
				}
			}

			if perms.Has(perm.MediafileCanManage) {
				allowed = append(allowed, mediafileID)
				continue
			}

			if perms.Has(perm.MediafileCanSee) {
				public := ds.Mediafile_IsPublic(mediafileID).ErrorLater(ctx)
				if public {
					allowed = append(allowed, mediafileID)
					continue
				}

				inheritedGroups := ds.Mediafile_InheritedAccessGroupIDs(mediafileID).ErrorLater(ctx)
				for _, id := range inheritedGroups {
					if perms.InGroup(id) {
						allowed = append(allowed, mediafileID)
						continue
					}
				}

				if err := ds.Err(); err != nil {
					return nil, fmt.Errorf("checking can see conditions: %w", err)
				}
			}
		}
		return allowed, nil
	})
}
