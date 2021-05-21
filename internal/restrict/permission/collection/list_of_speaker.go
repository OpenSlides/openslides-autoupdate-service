package collection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// ListOfSpeaker handels the permissions of list of speakers and speakers.
func ListOfSpeaker(dp dataprovider.DataProvider) perm.ConnecterFunc {
	l := &listOfSpeaker{
		dp: dp,
	}
	return func(s perm.HandlerStore) {
		s.RegisterAction("speaker.create", perm.ActionFunc(l.speakerCreate))
		s.RegisterAction("speaker.delete", perm.ActionFunc(l.speakerDelete))
		s.RegisterRestricter("speaker", perm.CollectionFunc(l.speakerRead))

		s.RegisterAction("list_of_speakers.delete", perm.ActionFunc(l.listDelete))
		s.RegisterRestricter("list_of_speakers", perm.CollectionFunc(l.listRead))
	}
}

type listOfSpeaker struct {
	dp dataprovider.DataProvider
}

func (l *listOfSpeaker) speakerCreate(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var meetingID int
	if err := l.dp.Get(ctx, fmt.Sprintf("list_of_speakers/%s/meeting_id", payload["list_of_speakers_id"]), &meetingID); err != nil {
		return false, fmt.Errorf("getting meeting id: %w", err)
	}

	perms, err := perm.New(ctx, l.dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	puid, err := strconv.Atoi(string(payload["user_id"]))
	if err != nil {
		return false, fmt.Errorf("invalid value in payload['user_id']: %s", payload["user_id"])
	}

	requiredPerm := perm.ListOfSpeakersCanManage
	if puid == userID {
		requiredPerm = perm.ListOfSpeakersCanBeSpeaker
	}

	if perms.Has(requiredPerm) {
		return true, nil
	}
	perm.LogNotAllowedf("User %d can not set user %d on the list of speaker.", userID, puid)
	return false, nil
}

func (l *listOfSpeaker) speakerDelete(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	fqid := "speaker/" + string(payload["id"])
	var sUserID int
	if err := l.dp.Get(ctx, fqid+"/user_id", &sUserID); err != nil {
		return false, fmt.Errorf("getting `%s/user_id` from DB: %w", fqid, err)
	}

	// Speaker is deleting himself.
	if sUserID == userID {
		return true, nil
	}

	// Check if request-user is list-of-speaker-manager
	meetingID, err := l.dp.MeetingFromModel(ctx, fqid)
	if err != nil {
		return false, fmt.Errorf("getting meeting_id from speaker model: %w", err)
	}

	ok, err := perm.HasPerm(ctx, l.dp, userID, meetingID, perm.ListOfSpeakersCanManage)
	if err != nil {
		return false, fmt.Errorf("ensuring list-of-speaker-manager perms: %w", err)
	}

	return ok, nil
}

func (l *listOfSpeaker) speakerRead(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		fqid := fmt.Sprintf("speaker/%d", fqfield.ID)

		var suid int
		if err := l.dp.Get(ctx, fqid+"/user_id", &suid); err != nil {
			return false, fmt.Errorf("getting speaker user id: %w", err)
		}

		if suid == userID {
			return true, nil
		}

		meetingID, err := l.dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
		}

		allowed, err := perm.HasPerm(ctx, l.dp, userID, meetingID, perm.ListOfSpeakersCanSee)
		if err != nil {
			return false, fmt.Errorf("ensuring perm %w", err)
		}
		return allowed, nil
	})
}

func (l *listOfSpeaker) listDelete(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	perm.LogNotAllowedf("list_of_speaker.delete is an internal action.")
	return false, nil
}

func (l *listOfSpeaker) listRead(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		fqid := fmt.Sprintf("list_of_speakers/%d", fqfield.ID)

		// If the request user is a speaker in the list of speakers, he can see the list.
		var sids []int
		if err := l.dp.GetIfExist(ctx, fqid+"/speaker_ids", &sids); err != nil {
			return false, fmt.Errorf("getting speaker object ids: %w", err)
		}

		for _, sid := range sids {
			var suid int
			if err := l.dp.Get(ctx, fmt.Sprintf("speaker/%d/user_id", sid), &suid); err != nil {
				return false, fmt.Errorf("getting speaker user id: %w", err)
			}

			if suid == userID {
				return true, nil
			}
		}

		meetingID, err := l.dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			var errDoesNotExist dataprovider.DoesNotExistError
			if errors.As(err, &errDoesNotExist) {
				return false, nil
			}
			return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
		}

		allowed, err := perm.HasPerm(ctx, l.dp, userID, meetingID, perm.ListOfSpeakersCanSee)
		if err != nil {
			return false, fmt.Errorf("ensuring perm %w", err)
		}
		return allowed, nil
	})
}

func canSeeSpeaker(p *perm.Permission) bool {
	return p.Has(perm.ListOfSpeakersCanSee)
}
