package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// ReadPerm checks that the user has one permission in a meeting.
//
// If the user has the permission, all fields of the given collection can be
// seen.
func ReadPerm(dp dataprovider.DataProvider, permission perm.TPermission, collections ...string) perm.ConnecterFunc {
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
			s.RegisterRestricter(c, perm.CollectionFunc(isPublic))
		}
	}
}

func isPublic(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	for _, field := range fqfields {
		result[field.String()] = true
	}
	return nil
}

// LoggedIn can be seen by logged in users.
func LoggedIn(dp dataprovider.DataProvider, collections ...string) perm.ConnecterFunc {
	return func(s perm.HandlerStore) {
		for _, c := range collections {
			s.RegisterRestricter(c, perm.CollectionFunc(isLoggedIn))
		}
	}
}

func isLoggedIn(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	if userID == 0 {
		return nil
	}

	for _, field := range fqfields {
		result[field.String()] = true
	}
	return nil
}

func isInMeeting(dp dataprovider.DataProvider, collection string) perm.CollectionFunc {
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

func hasPerm(dp dataprovider.DataProvider, permission perm.TPermission, collection string) perm.CollectionFunc {
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
