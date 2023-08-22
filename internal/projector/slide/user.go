package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

// DbUser is the class with methods to get needed User Informations
type DbUser struct {
	Username     string `json:"username"`
	Title        string `json:"title"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Level        string `json:"structure_level"`
	DefaultLevel string `json:"default_structure_level"`
}

// NewUser gets the user from datastore and return the user as DbUser struct
// together with keys and error.
// The meeting_id is used only to get the user-level for this meeting.
func NewUser(ctx context.Context, fetch *datastore.Fetcher, id, meetingID int) (*DbUser, error) {
	fields := []string{
		"username",
		"title",
		"first_name",
		"last_name",
		"meeting_user_ids",
		"default_structure_level",
	}

	data := fetch.Object(ctx, fmt.Sprintf("user/%d", id), fields...)
	if err := fetch.Err(); err != nil {
		return nil, fmt.Errorf("getting user object for id %d: %w", id, err)
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("encoding user data: %w", err)
	}

	var u DbUser
	if err := json.Unmarshal(bs, &u); err != nil {
		return nil, fmt.Errorf("decoding user data: %w", err)
	}

	if u.FirstName == "" && u.LastName == "" && u.Username == "" {
		return nil, fmt.Errorf("neither firstName, lastName nor username found")
	}

	if meetingID == 0 || data["meeting_user_ids"] == nil {
		return &u, nil
	}

	var meetingUserIDs []int
	if err := json.Unmarshal(data["meeting_user_ids"], &meetingUserIDs); err != nil {
		return nil, fmt.Errorf("decoding meeting_user_ids: %w", err)
	}

	for _, id := range meetingUserIDs {
		var mid int
		fetch.Fetch(ctx, &mid, "meeting_user/%d/meeting_id", id)
		if err := fetch.Err(); err != nil {
			return nil, fmt.Errorf("get meeting of meeting_user %d: %w", id, err)
		}

		if mid != meetingID {
			continue
		}

		fetch.Fetch(ctx, &u.Level, "meeting_user/%d/structure_level", id)
		if err := fetch.Err(); err != nil {
			return nil, fmt.Errorf("get structure level of meeting_user %d: %w", id, err)
		}
		break
	}

	return &u, nil
}

// UserRepresentation returns the meeting-dependent string for the given user.
func (u *DbUser) UserRepresentation(meetingID int) string {
	name := u.UserShortName()
	level := u.UserStructureLevel(meetingID)
	if level == "" {
		return name
	}
	return fmt.Sprintf("%s (%s)", name, level)
}

// UserStructureLevel returns in first place the meeting specific level,
// otherwise the default level.
// It is assumed that the Level-field in DbUser-struct contains the
// meeting dependent level.
func (u *DbUser) UserStructureLevel(meetingID int) string {
	if u.Level == "" {
		return u.DefaultLevel
	}
	return u.Level
}

// UserShortName returns the short name as "title first_name last_name".
// Without first_name and last_name, uses username instead.
func (u *DbUser) UserShortName() string {
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
	return strings.Join(parts, " ")
}

// User renders the user slide.
func User(store *projector.SlideStore) {
	store.RegisterSliderFunc("user", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (responseValue []byte, err error) {
		id, err := strconv.Atoi(strings.Split(p7on.ContentObjectID, "/")[1])
		if err != nil {
			return nil, fmt.Errorf("getting user id: %w", err)
		}

		user, err := NewUser(ctx, fetch, id, p7on.MeetingID)
		if err != nil {
			return nil, fmt.Errorf("getting new user id: %w", err)
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		out := struct {
			User string `json:"user"`
		}{
			user.UserRepresentation(p7on.MeetingID),
		}
		responseValue, err = json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding response slide user: %w", err)
		}
		return responseValue, err
	})

	store.RegisterGetTitleInformationFunc("user", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		id, err := strconv.Atoi(strings.Split(fqid, "/")[1])
		if err != nil {
			return nil, fmt.Errorf("getting user id: %w", err)
		}

		user, err := NewUser(ctx, fetch, id, meetingID)
		if err != nil {
			return nil, fmt.Errorf("loading user: %w", err)
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		out := struct {
			Collection      string `json:"collection"`
			ContentObjectID string `json:"content_object_id"`
			Username        string `json:"username"`
		}{
			"user",
			fqid,
			user.UserRepresentation(meetingID),
		}
		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding title: %w", err)
		}
		return responseValue, err
	})
}
