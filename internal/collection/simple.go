package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// ReadPerm checks that the user has one permission in a meeting.
//
// If the user has the permission, all fields of the given collection can be
// seen.
func ReadPerm(dp dataprovider.DataProvider, permission string, collections ...string) perm.ConnecterFunc {
	return func(s perm.HandlerStore) {
		for _, coll := range collections {
			s.RegisterRestricter(coll, hasPerm(dp, permission, coll))
		}

	}
}

// ReadInMeeting lets the user see the collection, if he is in the meeting.
func ReadInMeeting(dp dataprovider.DataProvider, collections ...string) perm.ConnecterFunc {
	return func(s perm.HandlerStore) {
		for _, coll := range collections {
			s.RegisterRestricter(coll, isInMeeting(dp, coll))
		}
	}
}

// Public can be seen by everyone.
func Public(dp dataprovider.DataProvider, collections ...string) perm.ConnecterFunc {
	return func(s perm.HandlerStore) {
		for _, c := range collections {
			s.RegisterRestricter(c, perm.RestricterCheckerFunc(isPublic))
		}
	}
}

func isPublic(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	for _, field := range fqfields {
		result[field.String()] = true
	}
	return nil
}

func isInMeeting(dp dataprovider.DataProvider, collection string) perm.RestricterCheckerFunc {
	return func(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
		return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
			fqid := fmt.Sprintf("%s/%d", collection, fqfield.ID)
			meetingID, err := dp.MeetingFromModel(ctx, fqid)
			if err != nil {
				return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
			}

			inMeeting, err := dp.InMeeting(ctx, userID, meetingID)
			if err != nil {
				return false, fmt.Errorf("see if user %d is in meeting %d: %w", userID, meetingID, err)
			}
			return inMeeting, nil
		})
	}
}

func hasPerm(dp dataprovider.DataProvider, permission, collection string) perm.RestricterCheckerFunc {
	return func(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
		return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
			fqid := fmt.Sprintf("%s/%d", collection, fqfield.ID)
			meetingID, err := dp.MeetingFromModel(ctx, fqid)
			if err != nil {
				return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
			}

			allowed, err := perm.HasPerm(ctx, dp, userID, meetingID, permission)
			if err != nil {
				return false, fmt.Errorf("ensuring perm %w", err)
			}
			return allowed, nil
		})
	}
}

// WritePerm initializes actions, that only need one permission
func WritePerm(dp dataprovider.DataProvider, def map[string]string) perm.ConnecterFunc {
	return func(s perm.HandlerStore) {
		for route, perm := range def {
			parts := strings.Split(route, ".")
			if len(parts) != 2 {
				panic("Invalid WritePerm action: " + route)
			}
			s.RegisterAction(route, (writeChecker(dp, parts[1], perm)))
		}
	}
}

func writeChecker(dp dataprovider.DataProvider, collName, permission string) perm.ActionChecker {
	return perm.ActionCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
		var meetingID int
		if err := json.Unmarshal(payload["meeting_id"], &meetingID); err != nil {
			var id int
			if err := json.Unmarshal(payload["id"], &id); err != nil {
				return false, fmt.Errorf("no valid meeting_id or id in payload")
			}

			fqid := collName + "/" + strconv.Itoa(id)
			meetingID, err = dp.MeetingFromModel(ctx, fqid)
			if err != nil {
				return false, fmt.Errorf("getting meeting id for %s: %w", fqid, err)
			}
		}

		ok, err := perm.HasPerm(ctx, dp, userID, meetingID, permission)
		if err != nil {
			return false, fmt.Errorf("checking permission: %w", err)
		}

		return ok, nil
	})
}
