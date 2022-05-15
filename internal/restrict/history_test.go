package restrict_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/dsmock"
)

func TestHistoryGetter(t *testing.T) {
	ctx := context.Background()

	for _, tt := range []struct {
		name         string
		current      string
		old          string
		testKey      string
		expectCanSee bool
	}{
		{
			"anonymous",
			`---
			user/1/id: 1`,
			``,
			"user/1/name",
			false,
		},
		{
			"orga admin",
			`---
			user/1/organization_management_level: can_manage_organization`,
			``,
			"user/1/name",
			true,
		},
		{
			"orga user management",
			`---
			user/1/organization_management_level: can_manage_users`,
			``,
			"user/1/name",
			false,
		},
		{
			"password field",
			`---
			user/1/organization_management_level: can_manage_organization`,
			``,
			"user/1/password",
			false,
		},
		{
			"personal note same user",
			`---
			user/1/organization_management_level: can_manage_organization`,
			`---
			personal_note/5/user_id: 1
			`,
			"personal_note/5/note",
			true,
		},
		{
			"personal note other user",
			`---
			user/1/organization_management_level: can_manage_organization`,
			`---
			personal_note/5/user_id: 2
			`,
			"personal_note/5/note",
			false,
		},
		{
			"meeting object orga manager",
			`---
			user/1/organization_management_level: can_manage_organization
			`,
			``,
			"topic/5/title",
			true,
		},
		{
			"meeting object not orga manager",
			`---
			user/1/id: 1
			`,
			``,
			"topic/5/title",
			false,
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
			"topic/5/title",
			true,
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
			"topic/5/title",
			false,
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
			"theme/5/name",
			true,
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
			"committee/5/name",
			false,
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
			"user/50/username",
			true,
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
			"unknown/50/name",
			false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			currentDS := dsmock.Stub(dsmock.YAMLData(tt.current))
			oldDS := dsmock.Stub(dsmock.YAMLData(tt.old))
			history := restrict.NewHistory(currentDS, oldDS, 1)

			key := MustKey(tt.testKey)

			got, err := history.Get(ctx, key)
			if err != nil {
				t.Fatalf("Get returned: %v", err)
			}

			if tt.expectCanSee && len(got) == 0 {
				t.Errorf("history.Get() did not return %v", tt.testKey)
			}

			if !tt.expectCanSee && len(got) == 1 {
				t.Errorf("histroy.Get() did return %v", tt.testKey)
			}
		})
	}
}
