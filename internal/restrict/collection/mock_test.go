package collection_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/sql"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

type testData struct {
	name          string
	data          map[dskey.Key][]byte
	expect        []int
	expectOne     bool
	requestUserID int
	elementIDs    []int
}

func testCase(name string, t *testing.T, f collection.FieldRestricter, expect bool, yaml string, op ...testCaseOption) {
	t.Helper()
	td := testData{
		name:          name,
		expect:        nil,
		expectOne:     expect,
		data:          dsmock.YAMLData(yaml),
		requestUserID: 1,
		elementIDs:    []int{1},
	}

	for _, o := range op {
		o(&td)
	}

	if td.requestUserID != 0 {
		userIDKey, err := dskey.FromString(fmt.Sprintf("user/%d/id", td.requestUserID))
		if err != nil {
			t.Fatalf("invalid key %v", fmt.Sprintf("user/%d/id", td.requestUserID))
		}

		td.data[userIDKey] = []byte(strconv.Itoa(td.requestUserID))
	}

	td.test(t, f)
}

func testCaseMulti(name string, t *testing.T, f collection.FieldRestricter, ids, expect []int, yaml string, op ...testCaseOption) {
	t.Helper()
	td := testData{
		name:          name,
		expect:        expect,
		data:          dsmock.YAMLData(yaml),
		requestUserID: 1,
		elementIDs:    ids,
	}

	for _, o := range op {
		o(&td)
	}

	userIDKey := dskey.MustKey("user/%d/id", td.requestUserID)

	td.data[userIDKey] = []byte(strconv.Itoa(td.requestUserID))
	td.data[dskey.MustKey("user/%d/username", td.requestUserID)] = []byte(`"Request User"`)
	td.data[dskey.MustKey("user/%d/organization_id", td.requestUserID)] = []byte(`1`)

	td.test(t, f)
}

func (tt testData) test(t *testing.T, f collection.FieldRestricter) {
	t.Helper()

	t.Run(tt.name, func(t *testing.T) {
		t.Helper()
		getter := dsmock.Stub(tt.data)

		fmt.Println(sql.Insert(tt.data))

		ds := dsfetch.New(getter)
		ctx := perm.ContextWithPermissionCache(context.Background(), getter, tt.requestUserID)

		allowedIDs, err := f(ctx, ds, tt.elementIDs...)
		if err != nil {
			t.Fatalf("restriction mode returned unexpected error: %v", err)
		}

		if tt.expect == nil {
			// test for one value
			got := len(allowedIDs) == 1

			if got != tt.expectOne {
				t.Errorf("restriction mode returned %t, expected %t", got, tt.expectOne)
			}
			return
		}

		if !set.Equal(set.New(allowedIDs...), set.New(tt.expect...)) {
			t.Errorf("restriction mode returned %v, expected %v", allowedIDs, tt.expect)
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
		jsonPerms, err := json.Marshal(perms)
		if err != nil {
			panic(err)
		}

		meetingUserID := td.requestUserID * 10
		groupID := 1337

		userMeetingUserIDsKey := dskey.MustKey("user/%d/meeting_user_ids", td.requestUserID)
		meetingUserGroupIDsKey := dskey.MustKey("meeting_user/%d/group_ids", meetingUserID)
		meeetingUserMeetingIDKey := dskey.MustKey("meeting_user/%d/meeting_id", meetingUserID)
		meetingUserIDKey := dskey.MustKey("meeting_user/%d/id", meetingUserID)

		groupIDKey := dskey.MustKey("group/%d/id", groupID)
		groupPermissionKey := dskey.MustKey("group/%d/permissions", groupID)

		td.data[userMeetingUserIDsKey] = jsonAppend(td.data[userMeetingUserIDsKey], meetingUserID)
		td.data[dskey.MustKey("user/%d/username", td.requestUserID)] = []byte(`"request user"`)
		td.data[dskey.MustKey("user/%d/organization_id", td.requestUserID)] = []byte(`1`)
		td.data[meetingUserGroupIDsKey] = jsonAppend(td.data[meetingUserGroupIDsKey], groupID)
		td.data[meeetingUserMeetingIDKey] = []byte(strconv.Itoa(meetingID))
		td.data[meetingUserIDKey] = []byte(strconv.Itoa(meetingUserID))
		td.data[dskey.MustKey("meeting_user/%d/user_id", meetingUserID)] = []byte(strconv.Itoa(td.requestUserID))

		td.data[groupIDKey] = []byte(strconv.Itoa(groupID))
		td.data[dskey.MustKey("group/%d/meeting_id", groupID)] = []byte(strconv.Itoa(meetingID))
		td.data[dskey.MustKey("group/%d/name", groupID)] = []byte("test group")
		td.data[groupPermissionKey] = jsonPerms
		//td.data[dskey.MustKey("meeting/%d/id", meetingID)] = []byte(strconv.Itoa(meetingID))
	}
}

func withRequestUser(userID int) testCaseOption {
	return func(td *testData) {
		td.requestUserID = userID
	}
}

func withElementID(id int) testCaseOption {
	return func(td *testData) {
		td.elementIDs = []int{id}
	}
}

func jsonAppend(value []byte, element ...int) []byte {
	var list []int
	if value != nil {
		if err := json.Unmarshal([]byte(value), &list); err != nil {
			panic(err)
		}
	}
	list = append(list, element...)
	newValue, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	return newValue
}
