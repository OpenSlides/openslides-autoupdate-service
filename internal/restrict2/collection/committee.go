package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

// Committee handels permission for committees.
type Committee struct{}

// Modes returns a map from all known modes to there restricter.
func (a Committee) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return a.see
	case "B":
		//return a.modeB
	}
	return nil
}

func (a Committee) see(ctx context.Context, fetch *datastore.Fetcher, mperms *perm.MeetingPermission, committeeID int) (bool, error) {
	userIDs := datastore.Ints(ctx, fetch.FetchIfExist, "committee/%d/user_ids", committeeID)
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting committee users: %w", err)
	}

	for _, uid := range userIDs {
		if uid == mperms.UserID() {
			return true, nil
		}
	}

	oml := datastore.String(ctx, fetch.FetchIfExist, "user/%d/organization_management_level", mperms.UserID())
	if err := fetch.Err(); err != nil {
		return false, fmt.Errorf("getting oml of user %d: %w", mperms.UserID(), err)
	}

	if oml == "can_manage_organization" || oml == "can_manage_users" {
		return true, nil
	}

	return false, nil
}
