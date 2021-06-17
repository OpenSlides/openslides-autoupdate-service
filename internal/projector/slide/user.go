package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbUser struct {
	Username     string `json:"username"`
	Title        string `json:"title"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Level        string `json:"structure_level_$"`
	DefaultLevel string `json:"default_structure_level"`
}

func newUser(ctx context.Context, ds datastore.Getter, id, meetingID int) (*dbUser, []string, error) {
	fields := []string{
		"username",
		"title",
		"first_name",
		"last_name",
		"default_structure_level",
	}
	if meetingID != 0 {
		fields = append(fields, fmt.Sprintf("structure_level_$%d", meetingID))
	}

	data, keys, err := datastore.Object(ctx, ds, fmt.Sprintf("user/%d", id), fields)
	if err != nil {
		return nil, nil, fmt.Errorf("getting user object: %w", err)
	}

	if meetingID != 0 {
		data["structure_level_$"] = data[fmt.Sprintf("structure_level_$%d", meetingID)]
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return nil, nil, fmt.Errorf("encoding user data")
	}

	var u dbUser
	if err := json.Unmarshal(bs, &u); err != nil {
		return nil, nil, fmt.Errorf("decoding user: %w", err)
	}

	return &u, keys, nil
}

// StringMeetingDependent gets the instance representation of the user, which is meeting dependent with the structur_level
func (u *dbUser) StringMeetingDependent(meetingID int) string {
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

	level := u.Level
	if level == "" {
		level = u.DefaultLevel
	}
	if level != "" {
		parts = append(parts, fmt.Sprintf("(%s)", level))
	}

	return strings.Join(parts, " ")
}

// getUserRepresentation returns the meeting-dependent string for the given user, including database access
func getUserRepresentation(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
	id, err := strconv.Atoi(strings.Split(p7on.ContentObjectID, "/")[1])
	if err != nil {
		return nil, nil, fmt.Errorf("getting user id: %w", err)
	}

	u, keys, err := newUser(ctx, ds, id, p7on.MeetingID)
	if err != nil {
		return nil, nil, fmt.Errorf("loading user: %w", err)
	}

	repr := u.StringMeetingDependent(p7on.MeetingID)
	if repr == "" {
		return nil, nil, slidesError{"Neither firstName, lastName nor username found", "user", p7on.ID, p7on.Type, p7on.ContentObjectID, p7on.MeetingID}
	}
	return []byte(fmt.Sprintf(`{"user":"%s"}`, repr)), keys, nil
}

// User renders the user slide.
func User(store *projector.SlideStore) {
	store.RegisterSlideFunc("user", getUserRepresentation)
	store.RegisterTitleFunc("user", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (*projector.TitleFuncResult, error) {
		title := "title of user"
		titleData := projector.TitleFuncResult{
			Title: &title,
		}
		return &titleData, nil
	})
}
