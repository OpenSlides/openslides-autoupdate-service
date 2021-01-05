package collection

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Speaker handels the permissions of the speaker collection.
type Speaker struct {
	dp dataprovider.DataProvider
}

// NewSpeaker initializes a Speaker.
func NewSpeaker(dp dataprovider.DataProvider) *Speaker {
	return &Speaker{
		dp: dp,
	}
}

// Connect connects the list_of_speakers routes.
func (sp *Speaker) Connect(s perm.HandlerStore) {
	s.RegisterWriteHandler("speaker.delete", perm.WriteCheckerFunc(sp.delete))
	s.RegisterReadHandler("speaker", perm.ReadCheckerFunc(sp.read))

}

func (sp *Speaker) delete(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	fqid := "speaker/" + string(payload["id"])
	var sUserID int
	if err := sp.dp.Get(ctx, fqid+"/user_id", &sUserID); err != nil {
		return nil, fmt.Errorf("getting `%s/user_id` from DB: %w", fqid, err)
	}

	// Speaker is deleting himself.
	if sUserID == userID {
		return nil, nil
	}

	// Check if request-user is list-of-speaker-manager
	meetingID, err := sp.dp.MeetingFromModel(ctx, fqid)
	if err != nil {
		return nil, fmt.Errorf("getting meeting_id from speaker model: %w", err)
	}

	if err := perm.EnsurePerm(ctx, sp.dp, userID, meetingID, "agenda.can_manage_list_of_speakers"); err != nil {
		return nil, fmt.Errorf("ensuring list-of-speaker-manager perms: %w", err)
	}

	return nil, nil
}

func (sp *Speaker) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		fqid := fmt.Sprintf("speaker/%d", fqfield.ID)
		meetingID, err := sp.dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			var errDoesNotExist dataprovider.DoesNotExistError
			if errors.As(err, &errDoesNotExist) {
				return false, nil
			}
			return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
		}

		allowed, err := perm.IsAllowed(perm.EnsurePerm(ctx, sp.dp, userID, meetingID, "agenda.can_see_list_of_speakers"))
		if err != nil {
			return false, fmt.Errorf("ensuring perm %w", err)
		}
		return allowed, nil
	})
}
