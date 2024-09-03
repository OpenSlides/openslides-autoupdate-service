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
				managementLevel, err := perm.HasOrganizationManagementLevel(ctx, ds, requestUser, perm.OMLCanManageOrganization)
				if err != nil {
					return nil, err
				}

				if managementLevel {
					return ids, nil
				}
			}
		}

		return eachCondition(ids, func(mediafileID int) (bool, error) {
			token, err := ds.Mediafile_Token(mediafileID).Value(ctx)
			if err != nil {
				return false, err
			}

			if token == "web_header" {
				return true, nil
			}

			published, err := ds.Mediafile_PublishedToMeetingsInOrganizationID(mediafileID).Value(ctx)
			if err != nil {
				return false, err
			}

			if val, isSet := published.Value(); isSet && val == 1 {
				isAdmin, err := isMeetingAdmin(ctx, ds)
				if err != nil {
					return false, err
				}

				if isAdmin {
					return true, nil
				}
			} else if collection == "organization" {
				return requestUser != 0, nil
			}

			meetingMediafileIDs, err := ds.Mediafile_MeetingMediafileIDs(mediafileID).Value(ctx)
			if err != nil {
				return false, err
			}

			if len(meetingMediafileIDs) > 0 {
				canSeeMeetingMediafile, err := Collection(ctx, MeetingMediafile{}.Name()).Modes("A")(ctx, ds, meetingMediafileIDs...)
				if err != nil {
					return false, fmt.Errorf("can see meeting mediafile of mediafile %d: %w", mediafileID, err)
				}

				if len(canSeeMeetingMediafile) >= 1 {
					return true, nil
				}
			}

			return false, nil
		})
	})
}

func isMeetingAdmin(ctx context.Context, ds *dsfetch.Fetch) (bool, error) {
	userID, err := perm.RequestUserFromContext(ctx)
	if err != nil {
		return false, fmt.Errorf("getting request user: %w", err)
	}

	if userID == 0 {
		return false, nil
	}

	meetingUserIDs, err := ds.User_MeetingUserIDs(userID).Value(ctx)
	if err != nil {
		return false, fmt.Errorf("getting meeting_user objects: %w", err)
	}

	for _, muid := range meetingUserIDs {
		groupIDs, err := ds.MeetingUser_GroupIDs(muid).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting groupIDs of user: %w", err)
		}

		adminGroups := make([]dsfetch.Maybe[int], len(groupIDs))
		for i := 0; i < len(groupIDs); i++ {
			ds.Group_AdminGroupForMeetingID(groupIDs[i]).Lazy(&adminGroups[i])
		}

		if err := ds.Execute(ctx); err != nil {
			return false, fmt.Errorf("checking for admin groups: %w", err)
		}

		for _, isAdmin := range adminGroups {
			if _, ok := isAdmin.Value(); ok {
				return true, nil
			}
		}

	}

	return false, nil
}
