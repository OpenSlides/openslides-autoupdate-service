package collection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// ListOfSpeaker handels the permissions of list of speakers and speakers.
func ListOfSpeaker(dp dataprovider.DataProvider) perm.ConnecterFunc {
	l := &listOfSpeaker{
		dp: dp,
	}
	return func(s perm.HandlerStore) {
		s.RegisterAction("speaker.create", perm.ActionCheckerFunc(l.createSpeaker))
		s.RegisterAction("speaker.delete", perm.ActionCheckerFunc(l.deleteSpeaker))
		s.RegisterRestricter("speaker", perm.RestricterCheckerFunc(l.readSpeaker))

		s.RegisterAction("list_of_speakers.delete", perm.ActionCheckerFunc(l.deleteList))
		s.RegisterRestricter("list_of_speakers", perm.RestricterCheckerFunc(l.readList))
	}
}

type listOfSpeaker struct {
	dp dataprovider.DataProvider
}

func (l *listOfSpeaker) createSpeaker(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var meetingID int
	if err := l.dp.Get(ctx, fmt.Sprintf("list_of_speakers/%s/meeting_id", payload["list_of_speakers_id"]), &meetingID); err != nil {
		return false, fmt.Errorf("getting meeting id: %w", err)
	}

	perms, err := perm.New(ctx, l.dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	var puid int
	if err := json.Unmarshal(payload["user_id"], &puid); err != nil {
		return false, fmt.Errorf("invalid value in payload['user_id']: %s", payload["user_id"])
	}

	requiredPerm := "list_of_speakers.can_manage"
	if puid == userID {
		requiredPerm = "list_of_speakers.can_be_speaker"
	}

	if perms.Has(requiredPerm) {
		return true, nil
	}
	perm.LogNotAllowedf("User %d can not set user %d on the list of speaker.", userID, puid)
	return false, nil
}

func (l *listOfSpeaker) deleteSpeaker(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
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

	ok, err := perm.HasPerm(ctx, l.dp, userID, meetingID, "list_of_speakers.can_manage")
	if err != nil {
		return false, fmt.Errorf("ensuring list-of-speaker-manager perms: %w", err)
	}

	return ok, nil
}

func (l *listOfSpeaker) readSpeaker(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
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

		allowed, err := perm.HasPerm(ctx, l.dp, userID, meetingID, "list_of_speakers.can_see")
		if err != nil {
			return false, fmt.Errorf("ensuring perm %w", err)
		}
		return allowed, nil
	})
}

func (l *listOfSpeaker) deleteList(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	perm.LogNotAllowedf("list_of_speaker.delete is an internal action.")
	return false, nil
}

func (l *listOfSpeaker) readList(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
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

		allowed, err := perm.HasPerm(ctx, l.dp, userID, meetingID, "list_of_speakers.can_see")
		if err != nil {
			return false, fmt.Errorf("ensuring perm %w", err)
		}
		return allowed, nil
	})
}

func canSeeSpeaker(p *perm.Permission) bool {
	return p.Has("list_of_speakers.can_see")
}
