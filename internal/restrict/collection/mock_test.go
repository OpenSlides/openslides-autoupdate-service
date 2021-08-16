package collection_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

type testData struct {
	name          string
	data          map[string]string
	expect        bool
	requestUserID int
	elementID     int
}

func testCase(name string, t *testing.T, f collection.FieldRestricter, expect bool, yaml string, op ...testCaseOption) {
	td := testData{
		name:          name,
		expect:        expect,
		data:          dsmock.YAMLData(yaml),
		requestUserID: 1,
		elementID:     1,
	}

	for _, o := range op {
		o(&td)
	}

	td.data[fmt.Sprintf("user/%d/id", td.requestUserID)] = strconv.Itoa(td.requestUserID)

	td.test(t, f)
}

func (tt testData) test(t *testing.T, f collection.FieldRestricter) {
	t.Helper()

	t.Run(tt.name, func(t *testing.T) {
		fetch := datastore.NewFetcher(dsmock.Stub(tt.data))
		perms := perm.NewMeetingPermission(fetch, tt.requestUserID)

		got, err := f(context.Background(), fetch, perms, tt.elementID)

		if err != nil {
			t.Fatalf("See returned unexpected error: %v", err)
		}

		if got != tt.expect {
			t.Errorf("See() returned %t, expected %t", got, tt.expect)
		}
	})
}

type testCaseOption func(*testData)

// withPerms uses the group X337 to add permissions to the request user in the given
// meeting. X is the meetingID.
//
// Make sure to call withRequestUser before withPerms.
func withPerms(meetingID int, perms ...perm.TPermission) testCaseOption {
	return func(td *testData) {
		permString := "["
		for _, p := range perms {
			permString += fmt.Sprintf("%q,", p)
		}
		permString = permString[:len(permString)-1] + "]"
		groupID := 1000 + 337

		groupsKey := fmt.Sprintf("user/1/group_$%d_ids", meetingID)
		td.data[groupsKey] = jsonAppend(td.data[groupsKey], groupID)
		td.data[fmt.Sprintf("group/%d/id", groupID)] = strconv.Itoa(groupID)
		td.data[fmt.Sprintf("group/%d/permissions", groupID)] = permString
		td.data[fmt.Sprintf("meeting/%d/id", meetingID)] = strconv.Itoa(meetingID)
	}
}

func withRequestUser(userID int) testCaseOption {
	return func(td *testData) {
		td.requestUserID = userID
	}
}

func withElementID(id int) testCaseOption {
	return func(td *testData) {
		td.elementID = id
	}
}

func jsonAppend(value string, element ...int) string {
	var list []int
	if value != "" {
		if err := json.Unmarshal([]byte(value), &list); err != nil {
			panic(err)
		}
	}
	list = append(list, element...)
	newValue, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	return string(newValue)

}
