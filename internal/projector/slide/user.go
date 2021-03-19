package slide

import (
	"context"
	"fmt"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/projector"
)

type user struct {
	Username  string `json:"username"`
	Title     string `json:"title"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//Level     string `json:"structure_level"`
}

func (u user) String() string {
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

	// if u.Level != "" {
	// 	parts = append(parts, fmt.Sprintf("(%s)", u.Level))
	// }

	return strings.Join(parts, " ")
}

// User renders the user slide.
func User(store *projector.SlideStore) {
	store.AddFunc("user", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		var u user
		if err := datastore.GetObject(ctx, ds, p7on.ContentObjectID, &u); err != nil {
			return nil, nil, fmt.Errorf("getting user object: %w", err)
		}

		return []byte(fmt.Sprintf(`{"user":"%s"}`, u.String())), datastore.ObjectKeys(p7on.ContentObjectID, &u), nil
	})
}
