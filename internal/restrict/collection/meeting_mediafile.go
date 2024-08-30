package collection

import (
	"context"
	"fmt"

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
type MeetingMediafile struct{}

// Name returns the collection name.
func (m MeetingMediafile) Name() string {
	return "meeting_mediafile"
}

// MeetingID returns the meetingID for the object.
func (m MeetingMediafile) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.MeetingMediafile_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("fetching meeting_id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the field modes for the collection mediafile.
func (m MeetingMediafile) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.see
	}
	return nil
}

func (m MeetingMediafile) see(ctx context.Context, ds *dsfetch.Fetch, meetingMediafileIDs ...int) ([]int, error) {
	return eachMeeting(ctx, ds, m, meetingMediafileIDs, func(meetingID int, ids []int) ([]int, error) {
		perms, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting perms for meeting %d: %w", meetingID, err)
		}

		if perms.IsAdmin() {
			return ids, nil
		}

		canSeeMeeting, err := Collection(ctx, Meeting{}.Name()).Modes("B")(ctx, ds, meetingID)
		if err != nil {
			return nil, fmt.Errorf("can see meeting %d: %w", meetingID, err)
		}

		return eachCondition(ids, func(meetingMediafileID int) (bool, error) {
			logoOrFont, err := usedAsLogoOrFont(ctx, ds, meetingMediafileID)
			if err != nil {
				return false, err
			}

			if len(canSeeMeeting) == 1 && logoOrFont {
				return true, nil
			}

			if perms.Has(perm.ProjectorCanSee) {
				p7onIDs, err := ds.MeetingMediafile_ProjectionIDs(meetingMediafileID).Value(ctx)
				if err != nil {
					return false, fmt.Errorf("getting projection ids: %w", err)
				}

				for _, p7onID := range p7onIDs {
					value, err := ds.Projection_CurrentProjectorID(p7onID).Value(ctx)
					if err != nil {
						return false, fmt.Errorf("getting current projector id: %w", err)
					}

					if !value.Null() {
						return true, nil
					}
				}
			}

			if perms.Has(perm.MediafileCanManage) {
				return true, nil
			}

			if perms.Has(perm.MediafileCanSee) {
				public, err := ds.MeetingMediafile_IsPublic(meetingMediafileID).Value(ctx)
				if err != nil {
					return false, fmt.Errorf("getting is public: %w", err)
				}

				if public {
					return true, nil
				}

				inheritedGroups, err := ds.MeetingMediafile_InheritedAccessGroupIDs(meetingMediafileID).Value(ctx)
				if err != nil {
					return false, fmt.Errorf("getting inheritedGroups: %w", err)
				}

				for _, id := range inheritedGroups {
					if perms.InGroup(id) {
						return true, nil
					}
				}
			}

			return false, nil
		})
	})
}

func usedAsLogoOrFont(ctx context.Context, ds *dsfetch.Fetch, meetingMediafileID int) (bool, error) {
	var usedAs struct {
		UsedAsLogoProjectorMainInMeetingID     dsfetch.Maybe[int]
		UsedAsLogoProjectorHeaderInMeetingID   dsfetch.Maybe[int]
		UsedAsLogoWebHeaderInMeetingID         dsfetch.Maybe[int]
		UsedAsLogoPdfHeaderLInMeetingID        dsfetch.Maybe[int]
		UsedAsLogoPdfHeaderRInMeetingID        dsfetch.Maybe[int]
		UsedAsLogoPdfFooterLInMeetingID        dsfetch.Maybe[int]
		UsedAsLogoPdfFooterRInMeetingID        dsfetch.Maybe[int]
		UsedAsLogoPdfBallotPaperInMeetingID    dsfetch.Maybe[int]
		UsedAsFontRegularInMeetingID           dsfetch.Maybe[int]
		UsedAsFontItalicInMeetingID            dsfetch.Maybe[int]
		UsedAsFontBoldInMeetingID              dsfetch.Maybe[int]
		UsedAsFontBoldItalicInMeetingID        dsfetch.Maybe[int]
		UsedAsFontMonospaceInMeetingID         dsfetch.Maybe[int]
		UsedAsFontChyronSpeakerNameInMeetingID dsfetch.Maybe[int]
		UsedAsFontProjectorH1InMeetingID       dsfetch.Maybe[int]
		UsedAsFontProjectorH2InMeetingID       dsfetch.Maybe[int]
	}

	ds.MeetingMediafile_UsedAsLogoProjectorMainInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoProjectorMainInMeetingID)
	ds.MeetingMediafile_UsedAsLogoProjectorHeaderInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoProjectorHeaderInMeetingID)
	ds.MeetingMediafile_UsedAsLogoWebHeaderInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoWebHeaderInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderLInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoPdfHeaderLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfHeaderRInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoPdfHeaderRInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterLInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoPdfFooterLInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfFooterRInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoPdfFooterRInMeetingID)
	ds.MeetingMediafile_UsedAsLogoPdfBallotPaperInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsLogoPdfBallotPaperInMeetingID)
	ds.MeetingMediafile_UsedAsFontRegularInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontRegularInMeetingID)
	ds.MeetingMediafile_UsedAsFontItalicInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontBoldInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontBoldInMeetingID)
	ds.MeetingMediafile_UsedAsFontBoldItalicInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontBoldItalicInMeetingID)
	ds.MeetingMediafile_UsedAsFontMonospaceInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontMonospaceInMeetingID)
	ds.MeetingMediafile_UsedAsFontChyronSpeakerNameInMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontChyronSpeakerNameInMeetingID)
	ds.MeetingMediafile_UsedAsFontProjectorH1InMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontProjectorH1InMeetingID)
	ds.MeetingMediafile_UsedAsFontProjectorH2InMeetingID(meetingMediafileID).Lazy(&usedAs.UsedAsFontProjectorH2InMeetingID)
	if err := ds.Execute(ctx); err != nil {
		return false, fmt.Errorf("fetching as logo and as font: %w", err)
	}

	return !usedAs.UsedAsLogoProjectorMainInMeetingID.Null() ||
		!usedAs.UsedAsLogoProjectorHeaderInMeetingID.Null() ||
		!usedAs.UsedAsLogoWebHeaderInMeetingID.Null() ||
		!usedAs.UsedAsLogoPdfHeaderLInMeetingID.Null() ||
		!usedAs.UsedAsLogoPdfHeaderRInMeetingID.Null() ||
		!usedAs.UsedAsLogoPdfFooterLInMeetingID.Null() ||
		!usedAs.UsedAsLogoPdfFooterRInMeetingID.Null() ||
		!usedAs.UsedAsLogoPdfBallotPaperInMeetingID.Null() ||
		!usedAs.UsedAsFontRegularInMeetingID.Null() ||
		!usedAs.UsedAsFontItalicInMeetingID.Null() ||
		!usedAs.UsedAsFontBoldInMeetingID.Null() ||
		!usedAs.UsedAsFontBoldItalicInMeetingID.Null() ||
		!usedAs.UsedAsFontMonospaceInMeetingID.Null() ||
		!usedAs.UsedAsFontChyronSpeakerNameInMeetingID.Null() ||
		!usedAs.UsedAsFontProjectorH1InMeetingID.Null() ||
		!usedAs.UsedAsFontProjectorH2InMeetingID.Null(), nil
}
