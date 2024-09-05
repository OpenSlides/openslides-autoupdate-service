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
// Mediafiles can be seen if:
//
//	The meeting mediafile belongs to the organization and the user has organization management permissions
//	The user is admin in a meeting and published_to_organization_id is set to the current organization id
//	The user can see any of the meeting mediafiles of the mediafile
//	The field token is set to any non empty value
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
					return nil, fmt.Errorf("getting organization management level: %w", err)
				}

				if managementLevel {
					return ids, nil
				}
			}
		}

		return eachCondition(ids, func(mediafileID int) (bool, error) {
			token, err := ds.Mediafile_Token(mediafileID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting mediafile token: %w", err)
			}

			if token != "" {
				return true, nil
			}

			published, err := ds.Mediafile_PublishedToMeetingsInOrganizationID(mediafileID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting published to meetings in organization: %w", err)
			}

			if val, _ := published.Value(); val == 1 {
				isAdmin, err := isMeetingAdmin(ctx, ds)
				if err != nil {
					return false, fmt.Errorf("checking if user is meeting admin: %w", err)
				}

				if isAdmin {
					return true, nil
				}
			} else if collection == "organization" {
				return false, nil
			}

			meetingMediafileIDs, err := ds.Mediafile_MeetingMediafileIDs(mediafileID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting meeting mediafile ids: %w", err)
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
