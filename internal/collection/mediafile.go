package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Mediafile implements the permission for the mediafile collection.
func Mediafile(dp dataprovider.DataProvider) perm.ConnecterFunc {
	m := &mediafile{dp: dp}

	return func(s perm.HandlerStore) {
		s.RegisterRestricter("mediafile", perm.CollectionFunc(m.read))

		s.RegisterAction("mediafile.can_see_mediafile", perm.ActionFunc(m.canSeeAction))
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

		hasPerms := perms.Has("mediafile.can_manage")
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

		return perms.Has("mediafile.can_see"), nil
	})
}

func (m *mediafile) canSeeAction(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var mediafileID int
	if err := json.Unmarshal(payload["id"], &mediafileID); err != nil {
		return false, fmt.Errorf("no valid id in payload")
	}

	fqid := "mediafile/" + strconv.Itoa(mediafileID)
	meetingID, err := m.dp.MeetingFromModel(ctx, fqid)
	if err != nil {
		return false, fmt.Errorf("getting meeting id for %s: %w", fqid, err)
	}

	perms, err := perm.New(ctx, m.dp, userID, meetingID)
	if err != nil {
		return false, fmt.Errorf("getting permissions: %w", err)
	}

	if perms == nil {
		return false, nil
	}

	if perms.Has("mediafile.can_manage") {
		return true, nil
	}

	if perms.Has("mediafile.can_see") {
		var isPublic bool
		if err := m.dp.GetIfExist(ctx, fqid+"/is_public", &isPublic); err != nil {
			return false, fmt.Errorf("getting is public: %w", err)
		}

		if isPublic {
			return true, nil
		}

		var accessGroups []int
		if err := m.dp.GetIfExist(ctx, fqid+"/inherited_access_group_ids", &accessGroups); err != nil {
			return false, fmt.Errorf("getting inherited_access_group_ids: %w", err)
		}

		for _, gid := range accessGroups {
			if perms.InGroup(gid) {
				return true, nil
			}
		}
	}

	var vars []string
	if err := m.dp.GetIfExist(ctx, fmt.Sprintf("%s/used_as_logo_$_in_meeting_id", fqid), &vars); err != nil {
		return false, fmt.Errorf("getting is as logo: %w", err)
	}
	if len(vars) > 0 {
		return true, nil
	}

	if err := m.dp.GetIfExist(ctx, fmt.Sprintf("%s/used_as_font_$_in_meeting_id", fqid), &vars); err != nil {
		return false, fmt.Errorf("getting is as font: %w", err)
	}
	if len(vars) > 0 {
		return true, nil
	}

	if !perms.Has("projector.can_see") {
		return false, nil
	}

	var currentProjector []int
	if err := m.dp.GetIfExist(ctx, fqid+"/current_projector_ids", &currentProjector); err != nil {
		return false, fmt.Errorf("getting current projector: %w", err)
	}

	return len(currentProjector) > 0, nil

}
