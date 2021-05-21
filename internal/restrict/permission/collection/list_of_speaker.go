package collection

import (
	"context"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// ListOfSpeaker handels the permissions of list of speakers and speakers.
func ListOfSpeaker(dp dataprovider.DataProvider) perm.ConnecterFunc {
	l := &listOfSpeaker{
		dp: dp,
	}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("speaker", perm.CollectionFunc(l.speakerRead))
		s.RegisterRestricter("list_of_speakers", perm.CollectionFunc(l.listRead))
	}
}

type listOfSpeaker struct {
	dp dataprovider.DataProvider
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
