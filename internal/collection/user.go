package collection

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// User handels the permissions of user-actions and the user collection.
func User(dp dataprovider.DataProvider) perm.ConnecterFunc {
	u := &user{dp: dp}
	return func(s perm.HandlerStore) {
		s.RegisterWriteHandler("user.create", perm.WriteCheckerFunc(u.create))
	}
}

type user struct {
	dp dataprovider.DataProvider
}

func (u *user) create(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var orgaLevel string
	if err := u.dp.GetIfExist(ctx, fmt.Sprintf("user/%d/organisation_management_level", userID), &orgaLevel); err != nil {
		return false, fmt.Errorf("getting organisation level: %w", err)
	}
	switch orgaLevel {
	case "is_superuser", "can_manage_organisation", "can_manage_users":
		return true, nil
	default:
		return false, nil
	}
}
