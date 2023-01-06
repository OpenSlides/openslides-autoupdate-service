package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/set"
)

// Meeting handels restrictions of the collection meeting.
//
// The user can see a meeting if one of the following is True:
//
//	`meeting/enable_anonymous`.
//	The user is in meeting/user_ids.
//	The user has the CML can_manage of the meeting's committee.
//	The user has the CML can_manage of any meeting and the meeting is a template meeting.
//	The user has the OML can_manage_organization.
//
// Mode A: Always visible to everyone.
//
// Mode B: The user can see the meeting.
//
// Mode C: The user has meeting.can_see_frontpage.
//
// Mode D: The user has meeting.can_see_livestream.
type Meeting struct {
	name string
}

// Name returns the collection name.
func (m Meeting) Name() string {
	return m.name
}

// MeetingID returns the meetingID for the object.
func (m Meeting) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return id, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m Meeting) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways(m.name, "A")
	case "B":
		return m.see
	case "C":
		return m.modeC
	case "D":
		return m.modeD
	}
	return nil
}

func (m Meeting) see(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, meetingIDs ...int) error {
	for _, meetingID := range meetingIDs {
		var meeting struct {
			enableAnonymous bool
			committeeID     int
			isTemplate      bool
			userIDs         []int
		}
		ds.Meeting_EnableAnonymous(meetingID).Lazy(&meeting.enableAnonymous)
		ds.Meeting_CommitteeID(meetingID).Lazy(&meeting.committeeID)
		ds.Meeting_TemplateForOrganizationID(meetingID).LazyExists(&meeting.isTemplate)
		ds.Meeting_UserIDs(meetingID).Lazy(&meeting.userIDs)

		if err := ds.Execute(ctx); err != nil {
			return fmt.Errorf("getting meeting %d: %w", meetingID, err)
		}

		if meeting.enableAnonymous {
			attrMap.Add(m.name, meetingID, "B", &allwaysAttr)
			continue
		}

		var committeeUsers []int
		if meeting.isTemplate {
			allCommitteeIDs, err := ds.Organization_CommitteeIDs(1).Value(ctx)
			if err != nil {
				return fmt.Errorf("getting all committee ids: %w", err)
			}

			usersIDsList := make([][]int, len(allCommitteeIDs))
			for i, committeeID := range allCommitteeIDs {
				ds.Committee_UserManagementLevel(committeeID, "can_manage").Lazy(&usersIDsList[i])

			}

			if err := ds.Execute(ctx); err != nil {
				return fmt.Errorf("getting all committee managers: %w", err)
			}

			for i := 0; i < len(usersIDsList); i++ {
				committeeUsers = append(committeeUsers, usersIDsList[i]...)
			}

		} else {
			users, err := ds.Committee_UserManagementLevel(meeting.committeeID, "can_manage").Value(ctx)
			if err != nil {
				return fmt.Errorf("getting committee users: %w", err)
			}
			committeeUsers = users
		}

		allUserIDs := append(committeeUsers, meeting.userIDs...)

		attrMap.Add(m.name, meetingID, "B", &Attributes{
			GlobalPermission: byte(perm.OMLCanManageOrganization),
			UserIDs:          set.New(allUserIDs...),
		})
	}
	return nil
}

func (m Meeting) modeC(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, meetingIDs ...int) error {
	for _, meetingID := range meetingIDs {
		permMap, err := mperms.Meeting(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("getting perm map for meeting %d: %w", meetingID, err)
		}

		attrMap.Add(m.name, meetingID, "C", &Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         permMap[perm.MeetingCanSeeFrontpage],
		})
	}

	return nil
}

func (m Meeting) modeD(ctx context.Context, ds *dsfetch.Fetch, mperms *perm.MeetingPermission, attrMap AttributeMap, meetingIDs ...int) error {
	for _, meetingID := range meetingIDs {
		permMap, err := mperms.Meeting(ctx, ds, meetingID)
		if err != nil {
			return fmt.Errorf("getting perm map for meeting %d: %w", meetingID, err)
		}

		attrMap.Add(m.name, meetingID, "C", &Attributes{
			GlobalPermission: byte(perm.OMLSuperadmin),
			GroupIDs:         permMap[perm.MeetingCanSeeLivestream],
		})
	}

	return nil
}
