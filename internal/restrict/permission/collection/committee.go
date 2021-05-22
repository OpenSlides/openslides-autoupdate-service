package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// Committee handels the permissions of committe objects.
func Committee(dp dataprovider.DataProvider) perm.ConnecterFunc {
	c := &committee{dp: dp}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("committee", perm.CollectionFunc(c.read))
	}
}

type committee struct {
	dp dataprovider.DataProvider
}

func (c *committee) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		var orgaLevel string
		if userID == 0 {
			return false, nil
		}

		if err := c.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/organisation_management_level", userID), &orgaLevel); err != nil {
			return false, fmt.Errorf("getting organisation level: %w", err)
		}

		if orgaLevel == "can_manage_organisation" {
			return true, nil
		}

		for _, field := range []string{"committee_as_member_ids", "committee_as_manager_ids"} {
			var ids []int
			if err := c.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/%s", userID, field), &ids); err != nil {
				return false, fmt.Errorf("getting user field %s: %w", field, err)
			}
			for _, id := range ids {
				if id == fqfield.ID {
					return true, nil
				}
			}
		}

		return false, nil
	})
}
