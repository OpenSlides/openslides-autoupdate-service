package slide

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbUser struct {
	Username  string         `json:"username"`
	Title     string         `json:"title"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Level     map[int]string `json:"structure_level_$"`
}

func (u dbUser) String(meetingID int) string {
	parts := func(sp ...string) []string {
		var full []string
		for _, s := range sp {
			if s == "" {
				continue
			}
			full = append(full, s)
		}
		return full
	}(u.Title, u.FirstName, u.LastName)

	if len(parts) == 0 {
		return u.Username
	}

	if level := u.Level[meetingID]; level != "" {
		parts = append(parts, fmt.Sprintf("(%s)", level))
	}

	return strings.Join(parts, " ")
}

// User renders the user slide.
func User(store *projector.SlideStore) {
	store.AddFunc("user", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		var u dbUser
		keys, err = datastore.Object(ctx, ds, p7on.ContentObjectID, &u)
		if err != nil {
			return nil, nil, fmt.Errorf("getting user object: %w", err)
		}

		return []byte(fmt.Sprintf(`{"user":"%s"}`, u.String(1))), keys, nil
	})
}
