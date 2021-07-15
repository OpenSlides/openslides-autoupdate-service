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

// inCommittee returns a set of all committee ids the user is in.
func inCommittee(ctx context.Context, dp dataprovider.DataProvider, userID int) (map[int]bool, error) {
	var inCommitteeIDs []int
	if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/committee_ids", userID), &inCommitteeIDs); err != nil {
		return nil, fmt.Errorf("checking committee ids of user: %w", err)
	}

	inCommittee := make(map[int]bool)
	for _, id := range inCommitteeIDs {
		inCommittee[id] = true
	}

	return inCommittee, nil
}

func (c *committee) read(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	if userID == 0 {
		return nil
	}

	orgaLevel, err := c.dp.OrgaLevel(ctx, userID)
	if err != nil {
		return fmt.Errorf("getting organization level: %w", err)
	}

	if orgaLevel == "can_manage_organization" {
		// The user can see all fields of all committees.
		for _, fqfield := range fqfields {
			result[fqfield.String()] = true
		}
		return nil
	}

	committees, err := inCommittee(ctx, c.dp, userID)
	if err != nil {
		return fmt.Errorf("getting all committee ids the user is in: %w", err)
	}

	groupA := map[string]bool{
		"id":                   true,
		"name":                 true,
		"description":          true,
		"meeting_ids":          true,
		"template_meeting_id":  true,
		"default_meeting_id":   true,
		"user_ids":             true,
		"organization_tag_ids": true,
		"organization_id":      true,
	}
	for _, fqfield := range fqfields {
		if !committees[fqfield.ID] && orgaLevel != "can_manage_users" {
			// User is not in commitee.
			continue
		}

		if groupA[fqfield.Field] {
			result[fqfield.String()] = true
		}
	}

	return nil
}
