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
//
//	The user is an admin of the meeting.
//	The user can see the meeting and used_as_logo_*_in_meeting_id or used_as_font_*_in_meeting_id is not empty.
//	The user has projector.can_see and there exists a mediafile/projection_ids with projection/current_projector_id set.
//	The user has mediafile.can_manage.
//	The user has mediafile.can_see and either:
//	    mediafile/is_public is true, or
//	    The user has groups in common with meeting/inherited_access_group_ids.
//
// Mode A: The user can see the mediafile.
type Mediafile struct{}

// Name returns the collection name.
func (m Mediafile) Name() string {
	return "mediafile"
}

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

func (m Mediafile) see(ctx context.Context, ds *dsfetch.Fetch, mediafileIDs ...int) ([]int, error) {
	requestUser, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting request user: %w", err)
	}

	return eachContentObjectCollection(ctx, ds.Mediafile_OwnerID, mediafileIDs, func(collection string, ownerID int, ids []int) ([]int, error) {
		// ownerID can be a meetingID or the organizationID
		if collection == "organization" {
			if requestUser != 0 {
				return ids, nil
			}
			return nil, nil
		}

		perms, err := perm.FromContext(ctx, ownerID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meeting %d: %w", ownerID, err)
		}

		if perms.IsAdmin() {
			return ids, nil
		}

		canSeeMeeting, err := Collection(ctx, Meeting{}.Name()).Modes("B")(ctx, ds, ownerID)
		if err != nil {
			return nil, fmt.Errorf("can see meeting %d: %w", ownerID, err)
		}

		return eachCondition(ids, func(mediafileID int) (bool, error) {
			logoOrFont, err := usedAsLogoOrFont(ctx, ds, mediafileID)
			if err != nil {
				return false, err
			}

			if len(canSeeMeeting) == 1 && logoOrFont {
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
		})
	})
}

func usedAsLogoOrFont(ctx context.Context, ds *dsfetch.Fetch, mediafileID int) (bool, error) {
	var usedAs struct {
		UsedAsLogoProjectorMainInMeetingID     int
		UsedAsLogoProjectorHeaderInMeetingID   int
		UsedAsLogoWebHeaderInMeetingID         int
		UsedAsLogoPdfHeaderLInMeetingID        int
		UsedAsLogoPdfHeaderRInMeetingID        int
		UsedAsLogoPdfFooterLInMeetingID        int
		UsedAsLogoPdfFooterRInMeetingID        int
		UsedAsLogoPdfBallotPaperInMeetingID    int
		UsedAsFontRegularInMeetingID           int
		UsedAsFontItalicInMeetingID            int
		UsedAsFontBoldInMeetingID              int
		UsedAsFontBoldItalicInMeetingID        int
		UsedAsFontMonospaceInMeetingID         int
		UsedAsFontChyronSpeakerNameInMeetingID int
		UsedAsFontProjectorH1InMeetingID       int
		UsedAsFontProjectorH2InMeetingID       int
	}

	ds.Mediafile_UsedAsLogoProjectorMainInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoProjectorMainInMeetingID)
	ds.Mediafile_UsedAsLogoProjectorHeaderInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoProjectorHeaderInMeetingID)
	ds.Mediafile_UsedAsLogoWebHeaderInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoWebHeaderInMeetingID)
	ds.Mediafile_UsedAsLogoPdfHeaderLInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoPdfHeaderLInMeetingID)
	ds.Mediafile_UsedAsLogoPdfHeaderRInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoPdfHeaderRInMeetingID)
	ds.Mediafile_UsedAsLogoPdfFooterLInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoPdfFooterLInMeetingID)
	ds.Mediafile_UsedAsLogoPdfFooterRInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoPdfFooterRInMeetingID)
	ds.Mediafile_UsedAsLogoPdfBallotPaperInMeetingID(mediafileID).Lazy(&usedAs.UsedAsLogoPdfBallotPaperInMeetingID)
	ds.Mediafile_UsedAsFontRegularInMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontRegularInMeetingID)
	ds.Mediafile_UsedAsFontItalicInMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontItalicInMeetingID)
	ds.Mediafile_UsedAsFontBoldInMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontBoldInMeetingID)
	ds.Mediafile_UsedAsFontBoldItalicInMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontBoldItalicInMeetingID)
	ds.Mediafile_UsedAsFontMonospaceInMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontMonospaceInMeetingID)
	ds.Mediafile_UsedAsFontChyronSpeakerNameInMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontChyronSpeakerNameInMeetingID)
	ds.Mediafile_UsedAsFontProjectorH1InMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontProjectorH1InMeetingID)
	ds.Mediafile_UsedAsFontProjectorH2InMeetingID(mediafileID).Lazy(&usedAs.UsedAsFontProjectorH2InMeetingID)
	if err := ds.Execute(ctx); err != nil {
		return false, fmt.Errorf("fetching as logo and as font: %w", err)
	}

	return usedAs.UsedAsLogoProjectorMainInMeetingID > 0 ||
		usedAs.UsedAsLogoProjectorHeaderInMeetingID > 0 ||
		usedAs.UsedAsLogoWebHeaderInMeetingID > 0 ||
		usedAs.UsedAsLogoPdfHeaderLInMeetingID > 0 ||
		usedAs.UsedAsLogoPdfHeaderRInMeetingID > 0 ||
		usedAs.UsedAsLogoPdfFooterLInMeetingID > 0 ||
		usedAs.UsedAsLogoPdfFooterRInMeetingID > 0 ||
		usedAs.UsedAsLogoPdfBallotPaperInMeetingID > 0 ||
		usedAs.UsedAsFontRegularInMeetingID > 0 ||
		usedAs.UsedAsFontItalicInMeetingID > 0 ||
		usedAs.UsedAsFontBoldInMeetingID > 0 ||
		usedAs.UsedAsFontBoldItalicInMeetingID > 0 ||
		usedAs.UsedAsFontMonospaceInMeetingID > 0 ||
		usedAs.UsedAsFontChyronSpeakerNameInMeetingID > 0 ||
		usedAs.UsedAsFontProjectorH1InMeetingID > 0 ||
		usedAs.UsedAsFontProjectorH2InMeetingID > 0, nil
}
