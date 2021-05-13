package slide

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/models"
)

// UserMeetingDependent gets the instance representation of the user, which is
// meeting dependent with the structur_level
func UserMeetingDependent(u *models.User, meetingID int) string {
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

	level := u.StructureLevel[strconv.Itoa(meetingID)]
	if level == "" {
		level = u.DefaultStructureLevel
	}
	if level != "" {
		parts = append(parts, fmt.Sprintf("(%s)", level))
	}

	return strings.Join(parts, " ")
}

// getUserRepresentation returns the meeting-dependent string for the given user, including database access
func getUserRepresentation(ctx context.Context, ds projector.Datastore, p7on *models.Projection) (encoded []byte, keys []string, err error) {
	user, err := models.LoadUser(ctx, ds, idFromFQID(p7on.ContentObjectID))
	if err != nil {
		return nil, nil, fmt.Errorf("getting user object: %w", err)
	}
	keys = []string{
		p7on.ContentObjectID + "/username",
		p7on.ContentObjectID + "/title",
		p7on.ContentObjectID + "/first_name",
		p7on.ContentObjectID + "/last_name",
		p7on.ContentObjectID + "/default_structure_level",
	}

	if p7on.MeetingID != 0 {
		keys = append(keys, p7on.ContentObjectID+"/structure_level_$"+strconv.Itoa(p7on.MeetingID))
	}

	repr := UserMeetingDependent(user, p7on.MeetingID)
	if repr == "" {
		return nil, nil, slidesError{"Neither firstName, lastName nor username found", "user", p7on.ID, p7on.Type, p7on.ContentObjectID, p7on.MeetingID}
	}
	return []byte(fmt.Sprintf(`{"user":"%s"}`, repr)), keys, nil
}

// User renders the user slide.
func User(store *projector.SlideStore) {
	store.AddFunc("user", getUserRepresentation)
}

func idFromFQID(fqid string) int {
	parts := strings.Split(fqid, "/")
	i, _ := strconv.Atoi(parts[1])
	return i
}
