package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// Meeting handels permissions of meeting objects.
func Meeting(dp dataprovider.DataProvider) perm.ConnecterFunc {
	m := &meeting{dp: dp}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("meeting", perm.CollectionFunc(m.read))
	}
}

type meeting struct {
	dp dataprovider.DataProvider
}

func (m *meeting) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	var perms *perm.Permission
	var lastID int
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
			lastID = fqfield.ID
			var err error
			perms, err = perm.New(ctx, m.dp, userID, fqfield.ID)
			if err != nil {
				return fmt.Errorf("getting perms: %w", err)
			}
		}

		switch fqfield.Field {
		case "enable_anonymous", "id", "name":
		case "welcome_title", "welcome_text":
			if !perms.Has(perm.MeetingCanSeeFrontpage) && !perms.Has("meeting.can_manage_settings") {
				continue
			}
		case "conference_stream_url", "conference_stream_poster_url":
			if !perms.Has(perm.MeetingCanSeeLivestream) && !perms.Has("meeting.can_manage_settings") {
				continue
			}
		case "present_user_ids", "temporary_user_ids", "guest_ids", "user_ids":
			if !perms.Has(perm.UserCanSee) {
				continue
			}
		default:
			if perms == nil {
				continue
			}
		}
		result[fqfield.String()] = true
	}
	return nil
}
