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

// OrgaLevel returns the organization level of a user. Returns an empty string
// if the user has no orga level.
func (dp *DataProvider) OrgaLevel(ctx context.Context, userID int) (string, error) {
	// The anonymous has no orga level.
	if userID == 0 {
		return "", nil
	}

	var orgaLevel string
	if err := dp.GetIfExist(ctx, fmt.Sprintf("user/%d/organization_management_level", userID), &orgaLevel); err != nil {
		return "", fmt.Errorf("getting organization level: %w", err)
	}
	return orgaLevel, nil
}

// InMeeting returns true, if the user is part of a meeting.
//
// Anonymous is part of a meeting, if anonymous is enabled.
func (dp *DataProvider) InMeeting(ctx context.Context, userID, meetingID int) (bool, error) {
	if userID == 0 {
		var enableAnonymous bool
		fqfield := "meeting/" + strconv.Itoa(meetingID) + "/enable_anonymous"
		if err := dp.GetIfExist(ctx, fqfield, &enableAnonymous); err != nil {
			return false, fmt.Errorf("checking anonymous enabled: %w", err)
		}
		return enableAnonymous, nil
	}

	fqfield := fmt.Sprintf("user/%d/group_$%d_ids", userID, meetingID)
	var groupIDs []int
	if err := dp.GetIfExist(ctx, fqfield, &groupIDs); err != nil {
		return false, fmt.Errorf("getting group ids of the meeting: %w", err)
	}
	return len(groupIDs) != 0, nil
}

// MeetingFromModel returns the meeting id for an model.
func (dp *DataProvider) MeetingFromModel(ctx context.Context, fqid string) (int, error) {
	var id int
	if err := dp.Get(ctx, fqid+"/meeting_id", &id); err != nil {
		return 0, fmt.Errorf("getting meeting id: %w", err)
	}
	return id, nil
}
