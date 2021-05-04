package slide

import (
	"context"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbUser struct {
	Username     string         `json:"username"`
	Title        string         `json:"title"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	Level        map[int]string `json:"structure_level_$"`
	DefaultLevel string         `json:"default_structure_level"`
}

// Get instance representation of the user
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
	}(u.FirstName, u.LastName)

	if len(parts) == 0 {
		parts = append(parts, u.Username)
	} else if u.Title != "" {
		parts = append([]string{u.Title}, parts...)
	}

	level := u.Level[meetingID]
	if level == "" {
		level = u.DefaultLevel
	}
	if level != "" {
		parts = append(parts, fmt.Sprintf("(%s)", level))
	}

	return strings.Join(parts, " ")
}

// Get Representation of the ContentObjectId and MeetingID from the Projection, assuming it's a User
func getUserRepresentation(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
	var u dbUser
	keys, err = datastore.Object(ctx, ds, p7on.ContentObjectID, &u)
	if err != nil {
		return nil, nil, fmt.Errorf("getting user object: %w", err)
	}

	repr := u.String(p7on.MeetingID)
	if repr == "" {
		return nil, nil, slidesError{"Neither firstName, lastName nor username found", "user", p7on.ID, p7on.Type, p7on.ContentObjectID, p7on.MeetingID}
	}
	return []byte(fmt.Sprintf(`{"user":"%s"}`, repr)), keys, nil
}

// User renders the user slide.
func User(store *projector.SlideStore) {
	store.AddFunc("user", getUserRepresentation)
}
