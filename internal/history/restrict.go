package history

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
)

type restricter struct {
	userID        int
	currentGetter datastore.Getter
	oldGetter     datastore.Getter
}

func newRestricter(current datastore.Getter, old datastore.Getter, userID int) restricter {
	return restricter{userID, current, old}
}

// Get returns the keys the user can see.
//
// In summary, a organization manager can see nearly all keys. A meeting admin
// can see all keys, that belong to there meeting.
func (h restricter) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	if h.userID == 0 {
		return nil, nil
	}

	currentDS := dsfetch.New(h.currentGetter)
	ctx = perm.ContextWithPermissionCache(ctx, h.currentGetter, h.userID)
	oldDS := dsfetch.New(h.oldGetter)

	orgaManager, err := perm.HasOrganizationManagementLevel(ctx, currentDS, h.userID, perm.OMLCanManageOrganization)
	if err != nil {
		return nil, fmt.Errorf("check orga management permission: %w", err)
	}

	requestUserMeetingIDs, err := currentDS.User_MeetingIDs(h.userID).Value(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting relevant meeting ids: %w", err)
	}

	adminInMeeting := make(map[int]struct{}, len(requestUserMeetingIDs))
	for _, meetingID := range requestUserMeetingIDs {
		p, err := perm.FromContext(ctx, meetingID)
		if err != nil {
			return nil, fmt.Errorf("getting permissions for meeting %d: %w", meetingID, err)
		}

		if p.IsAdmin() {
			adminInMeeting[meetingID] = struct{}{}
		}
	}

	if len(adminInMeeting) == 0 && !orgaManager {
		return nil, nil
	}

	allowedKeys := make([]dskey.Key, 0, len(keys))

	for _, key := range keys {
		canSee, err := h.canSeeKey(ctx, oldDS, currentDS, orgaManager, adminInMeeting, key)
		if err != nil {
			return nil, fmt.Errorf("checking key %s: %w", key, err)
		}

		if canSee {
			allowedKeys = append(allowedKeys, key)
		}
	}

	data, err := h.oldGetter.Get(ctx, allowedKeys...)
	if err != nil {
		return nil, fmt.Errorf("get data from history getter: %w", err)
	}
	return data, nil
}

func (h restricter) canSeeKey(
	ctx context.Context,
	oldDS,
	currentDS *dsfetch.Fetch,
	isOrgaManager bool,
	adminInMeeting map[int]struct{},
	key dskey.Key,
) (bool, error) {
	if key.Collection == "user" && key.Field == "password" {
		return false, nil
	}

	if key.Collection == "personal_note" {
		personalNoteUser, err := oldDS.PersonalNote_UserID(key.ID).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting personal note user: %w", err)
		}

		return personalNoteUser == h.userID, nil
	}

	if isOrgaManager {
		return true, nil
	}

	if key.Collection == "theme" || key.Collection == "organization" || key.Collection == "organization_tag" || key.Collection == "mediafile" {
		return true, nil
	}

	if key.Collection == "committee" {
		return false, nil
	}

	meetingID, hasMeeting, err := collection.Collection(ctx, key.Collection).MeetingID(ctx, oldDS, key.ID)
	if err != nil {
		return false, fmt.Errorf("getting meeting id: %w", err)
	}

	if hasMeeting {
		_, isAdmin := adminInMeeting[meetingID]
		return isAdmin, nil
	}

	if key.Collection == "user" {
		for _, r := range (collection.User{}).RequiredObjects(ctx, oldDS) {
			meetingIDs, err := r.TmplFunc(key.ID).Value(ctx)
			if err != nil {
				return false, fmt.Errorf("getting meeting ids for %s: %w", r.Name, err)
			}

			for _, meetingID := range meetingIDs {
				if _, ok := adminInMeeting[meetingID]; ok {
					return true, nil
				}
			}
		}
	}

	return false, nil
}
