package restrict

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// History filters the keys for the history.
//
// It checks if the request User is organization manager or is admin in a meeting.
type History struct {
	userID        int
	currentGetter datastore.Getter
	oldGetter     datastore.Getter
}

// NewHistory initializes a History object.
func NewHistory(userID int, current datastore.Getter, old datastore.Getter) History {
	return History{userID, current, old}
}

// Get returns the keys the user can see.
//
// In summary, a organization manager can see nearly all keys. A meeting admin
// can see all keys, that belong to there meeting.
func (h History) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	if h.userID == 0 {
		return nil, nil
	}

	currentDS := datastore.NewRequest(h.currentGetter)
	mperms := perm.NewMeetingPermission(currentDS, h.userID)
	oldDS := datastore.NewRequest(h.oldGetter)

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
		p, err := mperms.Meeting(ctx, meetingID)
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

	allowedKeys := make([]string, 0, len(keys))

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

func (h History) canSeeKey(
	ctx context.Context,
	oldDS,
	currentDS *datastore.Request,
	isOrgaManager bool,
	adminInMeeting map[int]struct{},
	key string,
) (bool, error) {
	coll, id, field, err := collectionIDField(key)
	if err != nil {
		return false, fmt.Errorf("splitting key: %w", err)
	}

	if coll == "user" && field == "password" {
		return false, nil
	}

	if coll == "personal_note" {
		personalNoteUser, err := oldDS.PersonalNote_UserID(id).Value(ctx)
		if err != nil {
			return false, fmt.Errorf("getting personal note user: %w", err)
		}

		return personalNoteUser == h.userID, nil
	}

	if isOrgaManager {
		return true, nil
	}

	if coll == "theme" || coll == "organization" || coll == "organization_tag" || coll == "mediafile" {
		return true, nil
	}

	if coll == "committee" {
		return false, nil
	}

	meetingID, hasMeeting, err := collection.Collection(coll).MeetingID(ctx, oldDS, id)
	if err != nil {
		return false, fmt.Errorf("getting meeting id: %w", err)
	}

	if hasMeeting {
		_, isAdmin := adminInMeeting[meetingID]
		return isAdmin, nil
	}

	if coll == "user" {
		for _, r := range (collection.User{}).RequiredObjects(oldDS) {
			meetingIDs, err := r.TmplFunc(id).Value(ctx)
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

func collectionIDField(key string) (string, int, string, error) {
	parts := strings.Split(key, "/")
	if len(parts) != 3 {
		return "", 0, "", fmt.Errorf("invalid key: %s", key)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, "", fmt.Errorf("second part not an id: %s", parts[1])
	}

	return parts[0], id, parts[2], nil
}
