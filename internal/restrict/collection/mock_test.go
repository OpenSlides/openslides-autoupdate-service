package collection_test

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/collection"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
)

type testData struct {
	name          string
	data          map[dskey.Key][]byte
	expect        []int
	expectOne     bool
	requestUserID int
	elementIDs    []int
	testElement   dskey.Key
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
		userIDKey := dskey.MustKey(fmt.Sprintf("user/%d/id", td.requestUserID))

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

	userIDKey := dskey.MustKey(fmt.Sprintf("user/%d/id", td.requestUserID))

	td.data[userIDKey] = []byte(strconv.Itoa(td.requestUserID))

	td.test(t, f)
}

func (tt testData) test(t *testing.T, f collection.FieldRestricter) {
	t.Helper()

	t.Run(tt.name, func(t *testing.T) {
		t.Helper()
		ctx := context.Background()

		fetcher := dsfetch.New(dsmock.Stub(tt.data))
		mperms := perm.NewMeetingPermission()

		attrMap := restrict.NewAttributeMap()

		err := f(context.Background(), fetcher, mperms, attrMap, tt.elementIDs...)
		if err != nil {
			t.Fatalf("restriction mode returned unexpected error: %v", err)
		}

		globalPerm, groupIDs, err := restrict.UserPermissions(ctx, fetcher, tt.requestUserID)
		if err != nil {
			t.Fatalf("getting user permissions: %v", err)
		}

		var zeroKey dskey.Key
		if tt.testElement != zeroKey {
			attr, err := attrMap.Get(ctx, fetcher, mperms, tt.testElement)
			if err != nil {
				t.Fatalf("attrMap: %v", err)
			}

			expect := tt.expectOne

			if got := restrict.AllowedByAttr(attr, tt.requestUserID, globalPerm, groupIDs); got != expect {
				t.Errorf("restriction mode returned %t, expedted %t", got, expect)
			}
			return
		}

		for cm, ids := range attrMap.RestrictModeIDs() {
			for _, id := range ids.List() {
				attr, err := attrMap.Get(ctx, fetcher, mperms, dskey.Key{Collection: cm.Collection, ID: id, Field: cm.Mode})
				if err != nil {
					t.Fatalf("attrMap: %v", err)
				}

				expect := tt.expectOne
				if tt.expect != nil {
					// This happens in a multi test
					expect = false
					for _, e := range tt.expect {
						if e == id {
							expect = true
							break
						}
					}
				}

				if got := restrict.AllowedByAttr(attr, tt.requestUserID, globalPerm, groupIDs); got != expect {
					t.Errorf("restriction mode returned %t, expedted %t", got, expect)
				}
			}
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
		permList, _ := json.Marshal(perms)
		groupID := 1000 + 337

		groupsTmplKey := dskey.MustKey(fmt.Sprintf("user/%d/group_$_ids", td.requestUserID))
		groupsKey := dskey.MustKey(fmt.Sprintf("user/%d/group_$%d_ids", td.requestUserID, meetingID))
		groupIDKey := dskey.MustKey(fmt.Sprintf("group/%d/id", groupID))
		groupPermissionKey := dskey.MustKey(fmt.Sprintf("group/%d/permissions", groupID))
		meetingIDKey := dskey.MustKey(fmt.Sprintf("meeting/%d/id", meetingID))
		meetingGroupIDs := dskey.MustKey(fmt.Sprintf("meeting/%d/group_ids", meetingID))

		td.data[groupsTmplKey] = jsonAppend(td.data[groupsTmplKey], strconv.Itoa(meetingID))
		td.data[groupsKey] = jsonAppend(td.data[groupsKey], groupID)
		td.data[groupIDKey] = []byte(strconv.Itoa(groupID))
		td.data[groupPermissionKey] = permList
		td.data[meetingIDKey] = []byte(strconv.Itoa(meetingID))
		td.data[meetingGroupIDs] = jsonAppend(td.data[meetingGroupIDs], groupID)
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

func withTestElement(collection string, id int, mode string) testCaseOption {
	return func(td *testData) {
		td.testElement = dskey.Key{Collection: collection, ID: id, Field: mode}
	}
}

func jsonAppend[T any](value []byte, element ...T) []byte {
	var list []T
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
