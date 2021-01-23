package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Meeting handels permissions of meeting objects.
func Meeting(dp dataprovider.DataProvider) perm.ConnecterFunc {
	m := &meeting{dp: dp}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("meeting", perm.RestricterCheckerFunc(m.read))
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
		case "enable_anonymous", "id":
		case "welcome_title", "welcome_text":
			if !perms.Has("meeting.can_see_frontpage") {
				continue
			}
		case "conference_stream_url", "conference_stream_poster_url":
			if !perms.Has("meeting.can_see_livestream") {
				continue
			}
		case "present_user_ids":
			if !perms.Has("user.can_see") {
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
