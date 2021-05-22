package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// Mediafile implements the permission for the mediafile collection.
func Mediafile(dp dataprovider.DataProvider) perm.ConnecterFunc {
	m := &mediafile{dp: dp}

	return func(s perm.HandlerStore) {
		s.RegisterRestricter("mediafile", perm.CollectionFunc(m.read))
	}
}

type mediafile struct {
	dp dataprovider.DataProvider
}

func (m *mediafile) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		fqid := fmt.Sprintf("mediafile/%d", fqfield.ID)
		meetingID, err := m.dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
		}

		perms, err := perm.New(ctx, m.dp, userID, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting user permissions: %w", err)
		}

		hasPerms := perms.Has(perm.MediafileCanManage)
		if hasPerms {
			return true, nil
		}

		var isPublic bool
		field := fqid + "/is_public"
		if err := m.dp.GetIfExist(ctx, field, &isPublic); err != nil {
			return false, fmt.Errorf("get %s: %w", field, err)
		}

		if !isPublic {
			return false, nil
		}

		return perms.Has(perm.MediafileCanSee), nil
	})
}
