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

// newUser gets the user from datastore and return the user as dbUser struct
// together with keys and error.
// The meeting_id is used only to get the user-level for this meeting.
func newUser(ctx context.Context, fetch *datastore.Fetcher, id, meetingID int) (*dbUser, error) {
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

	data := fetch.Object(ctx, fields, "user/%d", id)
	if fetch.Error() != nil {
		return nil, fmt.Errorf("getting user object: %w", fetch.Error())
	}

	if meetingID != 0 {
		data["structure_level_$"] = data[fmt.Sprintf("structure_level_$%d", meetingID)]
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("encoding user data: %w", err)
	}

	var u dbUser
	if err := json.Unmarshal(bs, &u); err != nil {
		return nil, fmt.Errorf("decoding user data: %w", err)
	}

	if u.FirstName == "" && u.LastName == "" && u.Username == "" {
		return nil, fmt.Errorf("neither firstName, lastName nor username found")
	}
	return &u, nil
}

// UserRepresentation returns the meeting-dependent string for the given user.
func (u *dbUser) UserRepresentation(meetingID int) string {
	name := u.UserShortName()
	level := u.UserStructureLevel(meetingID)
	if level == "" {
		return name
	}
	return fmt.Sprintf("%s (%s)", name, level)
}

// UserStructureLevel returns in first place the meeting specific level,
// otherwise the default level.
// It is assumed that the Level-field in dbUser-struct contains the
// meeting dependent level.
func (u *dbUser) UserStructureLevel(meetingID int) string {
	if u.Level == "" {
		return u.DefaultLevel
	}
	return u.Level
}

// UserShortName returns the short name as "title first_name last_name".
// Without first_name and last_name, uses username instead.
func (u *dbUser) UserShortName() string {
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
	store.RegisterSliderFunc("user", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (responseValue []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
			if err == nil {
				keys = fetch.Keys()
			} else {
				responseValue = nil
			}
		}()

		id, err := strconv.Atoi(strings.Split(p7on.ContentObjectID, "/")[1])
		if err != nil {
			return nil, nil, fmt.Errorf("getting user id: %w", err)
		}

		user, err := newUser(ctx, fetch, id, p7on.MeetingID)
		if err != nil {
			return nil, nil, fmt.Errorf("getting new user id: %w", err)
		}
		out := struct {
			User string `json:"user"`
		}{
			user.UserRepresentation(p7on.MeetingID),
		}
		responseValue, err = json.Marshal(out)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response slide user: %w", err)
		}
		return responseValue, keys, err
	})

	store.RegisterGetTitleInformationFunc("user", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		id, err := strconv.Atoi(strings.Split(fqid, "/")[1])
		if err != nil {
			return nil, fmt.Errorf("getting user id: %w", err)
		}

		user, err := newUser(ctx, fetch, id, meetingID)
		if err != nil {
			return nil, fmt.Errorf("loading user: %w", err)
		}
		out := struct {
			Username string `json:"username"`
		}{
			user.UserRepresentation(meetingID),
		}
		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding title: %w", err)
		}
		return responseValue, err
	})
}
