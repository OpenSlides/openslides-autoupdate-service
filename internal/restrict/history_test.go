package restrict_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func TestHistoryGetter(t *testing.T) {
	ctx := context.Background()

	for _, tt := range []struct {
		name    string
		current string
		old     string
		keys    []string
		expect  []string
	}{
		{
			"anonymous",
			`---
			user/1/id: 1`,
			``,
			[]string{"user/1/name"},
			nil,
		},
		{
			"orga admin",
			`---
			user/1/organization_management_level: can_manage_organization`,
			``,
			[]string{"user/1/name"},
			[]string{"user/1/name"},
		},
		{
			"orga user management",
			`---
			user/1/organization_management_level: can_manage_users`,
			``,
			[]string{"user/1/name"},
			nil,
		},
		{
			"password field",
			`---
			user/1/organization_management_level: can_manage_organization`,
			``,
			[]string{"user/1/password"},
			nil,
		},
		{
			"personal note same user",
			`---
			user/1/organization_management_level: can_manage_organization`,
			`---
			personal_note/5/user_id: 1
			`,
			[]string{"personal_note/5/note"},
			[]string{"personal_note/5/note"},
		},
		{
			"personal note other user",
			`---
			user/1/organization_management_level: can_manage_organization`,
			`---
			personal_note/5/user_id: 2
			`,
			[]string{"personal_note/5/note"},
			nil,
		},
		{
			"meeting object orga manager",
			`---
			user/1/organization_management_level: can_manage_organization
			`,
			``,
			[]string{"topic/5/title"},
			[]string{"topic/5/title"},
		},
		{
			"meeting object not orga manager",
			`---
			user/1/id: 1
			`,
			``,
			[]string{"topic/5/title"},
			nil,
		},
		{
			"meeting object history permission",
			`---
			user/1:
				group_$2_ids: [3]
				meeting_ids: [2]
			group/3/permissions: ["meeting.can_see_history"]
			meeting/2/admin_group_id: 3

			`,
			`topic/5/meeting_id: 2`,
			[]string{"topic/5/title"},
			[]string{"topic/5/title"},
		},
		{
			"meeting object wrong meeting",
			`---
			user/1:
				group_$2_ids: [3]
				meeting_ids: [2]
			group/3/permissions: ["meeting.can_see_history"]
			meeting/2/admin_group_id: 3
			`,
			`topic/5/meeting_id: 404`,
			[]string{"topic/5/title"},
			nil,
		},
		{
			"theme",
			`---
			user/1:
				group_$2_ids: [3]
				meeting_ids: [2]
			group/3/permissions: ["meeting.can_see_history"]
			meeting/2/admin_group_id: 3
			`,
			``,
			[]string{"theme/5/name"},
			[]string{"theme/5/name"},
		},
		{
			"committee",
			`---
			user/1:
				group_$2_ids: [3]
				meeting_ids: [2]
			group/3/permissions: ["meeting.can_see_history"]
			meeting/2/admin_group_id: 3
			`,
			``,
			[]string{"committee/5/name"},
			nil,
		},
		{
			"user is motion submitter",
			`---
			user/1:
				group_$2_ids: [3]
				meeting_ids: [2]
			group/3/permissions: ["meeting.can_see_history"]
			meeting/2/admin_group_id: 3
			`,
			`---
			user/50:
				submitted_motion_$_ids: ["2"]
			`,
			[]string{"user/50/username"},
			[]string{"user/50/username"},
		},
		{
			"unknown collection",
			`---
			user/1:
				group_$2_ids: [3]
				meeting_ids: [2]
			group/3/permissions: ["meeting.can_see_history"]
			meeting/2/admin_group_id: 3
			`,
			``,
			[]string{"unknown/50/name"},
			nil,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			currentDS := dsmock.Stub(dsmock.YAMLData(tt.current))
			oldDS := dsmock.Stub(dsmock.YAMLData(tt.old))
			history := restrict.NewHistory(currentDS, oldDS, 1)

			keys := make([]datastore.Key, len(tt.keys))
			for i, k := range tt.keys {
				keys[i] = MustKey(k)
			}

			got, err := history.Get(ctx, keys...)
			if err != nil {
				t.Fatalf("Get returned: %v", err)
			}

			gotKeys := make([]string, 0, len(got))
			for k := range got {
				gotKeys = append(gotKeys, k.String())
			}

			if !reflect.DeepEqual(gotKeys, tt.expect) {
				t.Errorf("Got %v, expected %v", gotKeys, tt.expect)
			}
		})
	}
}
