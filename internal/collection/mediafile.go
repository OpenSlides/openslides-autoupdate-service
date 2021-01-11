package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Mediafile implements the permission for the mediafile collection.
func Mediafile(dp dataprovider.DataProvider) perm.ConnecterFunc {
	read := func(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
		return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
			fqid := fmt.Sprintf("mediafile/%d", fqfield.ID)
			meetingID, err := dp.MeetingFromModel(ctx, fqid)
			if err != nil {
				return false, fmt.Errorf("getting meetingID from model %s: %w", fqid, err)
			}

			perms, err := perm.New(ctx, dp, userID, meetingID)
			if err != nil {
				return false, fmt.Errorf("getting user permissions: %w", err)
			}

			hasPerms := perms.Has("mediafile.can_manage")
			if hasPerms {
				return true, nil
			}

			var isPublic bool
			field := fqid + "/is_public"
			if err := dp.GetIfExist(ctx, field, &isPublic); err != nil {
				return false, fmt.Errorf("get %s: %w", field, err)
			}

			if !isPublic {
				return false, nil
			}

			return perms.Has("mediafile.can_see"), nil
		})
	}

	return func(s perm.HandlerStore) {
		s.RegisterReadHandler("mediafile", perm.ReadCheckerFunc(read))
	}
}
