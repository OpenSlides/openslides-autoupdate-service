package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// Create checks for the mermission to create a new object.
func Create(dp dataprovider.DataProvider, perm, collection string) WriteChecker {
	return WriteCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
		var meetingID int
		if err := json.Unmarshal(payload["meeting_id"], &meetingID); err != nil {
			return nil, fmt.Errorf("no valid meeting id: %w", err)
		}

		return check(ctx, dp, perm, meetingID, userID, payload)
	})
}

// Modify checks for the permissions to alter an existing object.
func Modify(dp dataprovider.DataProvider, perm, collection string) WriteChecker {
	return WriteCheckerFunc(func(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
		id, err := modelID(payload)
		if err != nil {
			return nil, fmt.Errorf("getting model id from payload: %w", err)
		}

		fqid := fmt.Sprintf("%s/%d", collection, id)
		meetingID, err := dp.MeetingFromModel(ctx, fqid)
		if err != nil {
			return nil, fmt.Errorf("getting meeting id for model %s: %w", fqid, err)
		}

		return check(ctx, dp, perm, meetingID, userID, payload)
	})
}

func check(ctx context.Context, dp dataprovider.DataProvider, managePerm string, meetingID int, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	superUser, err := dp.IsSuperuser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if superUser {
		return nil, nil
	}

	if err := EnsurePerms(ctx, dp, userID, meetingID, managePerm); err != nil {
		return nil, fmt.Errorf("ensure manage permission: %w", err)
	}
	return nil, nil
}

func modelID(data map[string]json.RawMessage) (int, error) {
	var id int
	if err := json.Unmarshal(data["id"], &id); err != nil {
		return 0, fmt.Errorf("no valid meeting id: %w", err)
	}
	return id, nil
}

// Restrict tells, if the user has the permission to see the requested
// fields.
func Restrict(dp dataprovider.DataProvider, perm, collection string) ReadeChecker {
	return ReadeCheckerFunc(func(ctx context.Context, userID int, fqfields []string, result map[string]bool) error {
		if len(fqfields) == 0 {
			return nil
		}

		parts := strings.Split(fqfields[0], "/")
		meetingID, err := dp.MeetingFromModel(ctx, collection+"/"+parts[1])
		if err != nil {
			return fmt.Errorf("getting meeting from model: %w", err)
		}

		if err := EnsurePerms(ctx, dp, userID, meetingID, perm); err != nil {
			return nil
		}

		for _, fqfield := range fqfields {
			result[fqfield] = true
		}
		return nil
	})
}

// EnsurePerms makes sure the user has at least one of the given permissions.
func EnsurePerms(ctx context.Context, dp dataprovider.DataProvider, userID int, meetingID int, permissions ...string) error {
	committeeID, err := dp.CommitteeID(ctx, meetingID)
	if err != nil {
		return fmt.Errorf("getting committee id for meeting: %w", err)
	}

	committeeManager, err := dp.IsManager(ctx, userID, committeeID)
	if err != nil {
		return fmt.Errorf("check for manager: %w", err)
	}
	if committeeManager {
		return nil
	}

	canSeeMeeting, err := dp.InMeeting(ctx, userID, meetingID)
	if err != nil {
		return err
	}
	if !canSeeMeeting {
		return NotAllowedf("User %d is not in meeting %d", userID, meetingID)
	}

	perms, err := Perms(ctx, userID, meetingID, dp)
	if err != nil {
		return fmt.Errorf("getting user permissions: %w", err)
	}

	hasPerms := perms.HasOne(permissions...)
	if !hasPerms {
		return NotAllowedf("User %d has not the required permission in meeting %d", userID, meetingID)
	}

	return nil
}
