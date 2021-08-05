package collection_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

type testData struct {
	name   string
	data   map[string]string
	expect bool
}

func testCase(name string, expect bool, yaml string, op ...testCaseOption) testData {
	td := testData{
		name:   name,
		expect: expect,
		data:   dsmock.YAMLData(yaml),
	}

	td.data["user/1/id"] = "1"

	for _, o := range op {
		o(td)
	}

	return td
}

func (tt testData) test(t *testing.T, f collection.FieldRestricter) {
	t.Helper()

	t.Run(tt.name, func(t *testing.T) {
		fetch := datastore.NewFetcher(dsmock.Stub(tt.data))
		perms := perm.NewMeetingPermission(fetch, 1)

		got, err := f(context.Background(), fetch, perms, 1)

		if err != nil {
			t.Fatalf("See returned unexpected error: %v", err)
		}

		if got != tt.expect {
			t.Errorf("See() returned %t, expected %t", got, tt.expect)
		}
	})
}

type testCaseOption func(testData)

// withPerms uses the group X337 to add permissions to user 1 in the given
// meeting. X is the meetingID.
func withPerms(meetingID int, perms ...perm.TPermission) testCaseOption {
	return func(td testData) {
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
