package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Motion handels restrictions of the collection motion.
type Motion struct{}

// Modes returns the restrictions modes for the meeting collection.
func (m Motion) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return m.modeA
	case "B":
		return m.modeB
	case "C":
		return m.see
	case "D":
		return m.modeD
	}
	return nil
}

func (m Motion) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionID int) (bool, error) {
	perms, err := mperms.Meeting(ctx, fetch.Field().Motion_MeetingID(ctx, motionID))
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if !perms.Has(perm.MotionCanSee) {
		return false, nil
	}

	for _, restriction := range fetch.Field().MotionState_Restrictions(ctx, fetch.Field().Motion_StateID(ctx, motionID)) {
		if restriction == "is_submitter" {
			isSubmitter := false
			for _, submitterID := range fetch.Field().Motion_SubmitterIDs(ctx, motionID) {
				if fetch.Field().MotionSubmitter_UserID(ctx, submitterID) == mperms.UserID() {
					isSubmitter = true
					break
				}
			}
			if err := fetch.Err(); err != nil {
				return false, fmt.Errorf("getting submitter: %w", err)
			}

			if !isSubmitter {
				return false, nil
			}
			continue
		}

		if !perms.Has(perm.TPermission(restriction)) {
			return false, nil
		}
	}
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("checking restrictions: %w", err)
	}
	return true, nil
}

func (m Motion) modeA(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionID int) (bool, error) {
	see, err := m.modeB(ctx, fetch, mperms, motionID)
	if err != nil {
		return false, fmt.Errorf("see motion: %w", err)
	}

	if see {
		return true, nil
	}

	agendaID, exist := fetch.Field().Motion_AgendaItemID(ctx, motionID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting agenda item: %w", err)
	}

	if exist {
		seeAgenda, err := AgendaItem{}.see(ctx, fetch, mperms, agendaID)
		if err != nil {
			return false, fmt.Errorf("checking see for agenda item %d: %w", agendaID, err)
		}

		if seeAgenda {
			return true, nil
		}
	}

	losID := fetch.Field().Motion_ListOfSpeakersID(ctx, motionID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting agenda item: %w", err)
	}

	if losID != 0 {
		seeLOS, err := ListOfSpeakers{}.see(ctx, fetch, mperms, losID)
		if err != nil {
			return false, fmt.Errorf("checking see for agenda item %d: %w", agendaID, err)
		}

		if seeLOS {
			return true, nil
		}
	}

	return false, nil
}

func (m Motion) modeB(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, motionID int) (bool, error) {
	motionIDs := append([]int{motionID}, fetch.Field().Motion_AllOriginIDs(ctx, motionID)...)
	motionIDs = append(motionIDs, fetch.Field().Motion_AllDerivedMotionIDs(ctx, motionID)...)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("fetching origin and derived motions: %w", err)
	}

	for _, referenceID := range motionIDs {
		see, err := m.see(ctx, fetch, mperms, referenceID)
		if err != nil {
			return false, fmt.Errorf("see motion %d: %w", referenceID, err)
		}

		if see {
			return true, nil
		}
	}
	return false, nil
}

func (m Motion) modeD(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, UserID int) (bool, error) {
	return false, nil
}
