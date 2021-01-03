package dataprovider

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type externalDataProvider interface {
	// If a field does not exist, it is not returned.
	Get(ctx context.Context, fields ...string) ([]json.RawMessage, error)
}

// DataProvider is a wrapper around permission.DataProvider that provides some
// helper functions.
type DataProvider struct {
	External externalDataProvider
}

func (dp *DataProvider) externalGet(ctx context.Context, fields ...string) ([]json.RawMessage, error) {
	return dp.External.Get(ctx, fields...)
}

// Get returns a value from the datastore and unpacks it in to the argument value.
//
// The argument value has to be an non nil pointer.
func (dp *DataProvider) Get(ctx context.Context, fqfield string, value interface{}) error {
	fields, err := dp.externalGet(ctx, fqfield)
	if err != nil {
		return fmt.Errorf("getting data from datastore: %w", err)
	}

	if fields[0] == nil {
		return DoesNotExistError(fqfield)
	}

	if err := json.Unmarshal(fields[0], value); err != nil {
		return fmt.Errorf("unpacking value: %w", err)
	}
	return nil
}

// GetIfExist behaves like Get() but does not throw an error if the fqfield does
// not exist.
func (dp *DataProvider) GetIfExist(ctx context.Context, fqfield string, value interface{}) error {
	if err := dp.Get(ctx, fqfield, value); err != nil {
		var errDoesNotExist DoesNotExistError
		if !errors.As(err, &errDoesNotExist) {
			return err
		}
	}
	return nil
}

// Exists tells, if a fqfield exist.
//
// If an error happens, it returns false.
func (dp *DataProvider) Exists(ctx context.Context, fqfield string) (bool, error) {
	fields, err := dp.externalGet(ctx, fqfield)
	if err != nil {
		return false, fmt.Errorf("getting fqfield: %w", err)
	}

	return fields[0] != nil, nil
}

// DoesUserExists returns true, if an user exist. Returns allways true for
// userID 0.
func (dp *DataProvider) DoesUserExists(ctx context.Context, userID int) (bool, error) {
	if userID == 0 {
		return true, nil
	}

	exists, err := dp.DoesModelExists(ctx, "user/"+strconv.Itoa(userID))
	if err != nil {
		return false, fmt.Errorf("lockup user: %w", err)
	}
	return exists, nil
}

// DoesModelExists returns true, if an object exists in the datastore.
func (dp *DataProvider) DoesModelExists(ctx context.Context, fqid string) (bool, error) {
	exists, err := dp.Exists(ctx, fqid+"/"+"id")
	if err != nil {
		return false, fmt.Errorf("checking for model existing: %w", err)
	}
	return exists, nil
}

// IsSuperuser returns true, if the user is in the superuser group.
func (dp *DataProvider) IsSuperuser(ctx context.Context, userID int) (bool, error) {
	// The anonymous is never a superadmin.
	if userID == 0 {
		return false, nil
	}

	// Get superadmin role id.
	var superadminRoleID int
	if err := dp.Get(ctx, "organisation/1/superadmin_role_id", &superadminRoleID); err != nil {
		return false, fmt.Errorf("getting superadmin role id: %w", err)
	}

	// Get users role id.
	fqfield := fmt.Sprintf("user/%d/role_id", userID)
	var userRoleID int
	if err := dp.GetIfExist(ctx, fqfield, &userRoleID); err != nil {
		return false, fmt.Errorf("getting role_id: %w", err)
	}

	return superadminRoleID == userRoleID, nil
}

// CommitteeID returns the id of a committee from an meeting id.
func (dp *DataProvider) CommitteeID(ctx context.Context, meetingID int) (int, error) {
	var committeeID int
	if err := dp.Get(ctx, "meeting/"+strconv.Itoa(meetingID)+"/committee_id", &committeeID); err != nil {
		return 0, fmt.Errorf("getting committee id: %w", err)
	}
	return committeeID, nil
}

// IsManager returns true, if the user is a manager in the committee.
func (dp *DataProvider) IsManager(ctx context.Context, userID, committeeID int) (bool, error) {
	// The anonymous is never a manager.
	if userID == 0 {
		return false, nil
	}

	// Get committee manager_ids.
	managerIDs := []int{}
	fqfield := "committee/" + strconv.Itoa(committeeID) + "/manager_ids"
	if err := dp.GetIfExist(ctx, fqfield, &managerIDs); err != nil {
		return false, fmt.Errorf("getting committee ids: %w", err)
	}

	for _, id := range managerIDs {
		if userID == id {
			return true, nil
		}
	}
	return false, nil
}

// MeetingIDFromPayload returns the id of a meeting from the payload.
//
// It expect the payload to have a idfield to an object. This object needs to
// have a field `meeting_id`.
func (dp *DataProvider) MeetingIDFromPayload(ctx context.Context, payload map[string]json.RawMessage, collection, idField string) (int, error) {
	var meetingID int
	if err := dp.Get(ctx, fmt.Sprintf("%s/%s/meeting_id", collection, payload[idField]), &meetingID); err != nil {
		return 0, fmt.Errorf("getting meetingID: %w", err)
	}
	return meetingID, nil
}

// InMeeting returns true, if the user is in the user_ids list or anonymous.
func (dp *DataProvider) InMeeting(ctx context.Context, userID, meetingID int) (bool, error) {
	if userID == 0 {
		var enableAnonymous bool
		fqfield := "meeting/" + strconv.Itoa(meetingID) + "/enable_anonymous"
		if err := dp.GetIfExist(ctx, fqfield, &enableAnonymous); err != nil {
			return false, fmt.Errorf("checking anonymous enabled: %w", err)
		}
		return enableAnonymous, nil
	}

	userIDs := []int{}
	fqfield := "meeting/" + strconv.Itoa(meetingID) + "/user_ids"
	if err := dp.GetIfExist(ctx, fqfield, &userIDs); err != nil {
		return false, fmt.Errorf("getting meeting/user_ids: %w", err)
	}

	for _, id := range userIDs {
		if id == userID {
			return true, nil
		}
	}
	return false, nil
}

// MeetingFromModel returns the meeting id for an model.
func (dp *DataProvider) MeetingFromModel(ctx context.Context, fqid string) (int, error) {
	var id int
	if err := dp.Get(ctx, fqid+"/meeting_id", &id); err != nil {
		return 0, fmt.Errorf("getting meeting id: %w", err)
	}
	return id, nil
}
