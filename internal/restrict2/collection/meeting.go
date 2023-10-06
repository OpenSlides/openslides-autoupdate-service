package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Meeting handels restrictions of the collection meeting.
//
// The user can see a meeting if one of the following is True:
//
//	`meeting/enable_anonymous`.
//	The user is in the meeting.
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
type Meeting struct{}

// Name returns the collection name.
func (m Meeting) Name() string {
	return "meeting"
}

// MeetingID returns the meetingID for the object.
func (m Meeting) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return id, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (m Meeting) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return Allways
	case "B":
		return m.see
	case "C":
		return m.modeC
	case "D":
		return m.modeD
	}
	return nil
}

func (m Meeting) see(ctx context.Context, fetcher *dsfetch.Fetch, meetingIDs []int) ([]attribute.Func, error) {
	orgaManger := attribute.FuncGlobalLevel(perm.OMLCanManageOrganization)

	result := make([]attribute.Func, len(meetingIDs))
	for i, meetingID := range meetingIDs {
		enableAnonymous, err := fetcher.Meeting_EnableAnonymous(meetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("checking enabled anonymous: %w", err)
		}

		if enableAnonymous {
			result[i] = attribute.FuncAllow
			continue
		}

		groupIDs, err := fetcher.Meeting_GroupIDs(meetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting group_ids: %w", err)
		}

		inMeeting := attribute.FuncInGroup(groupIDs)

		committeeID, err := fetcher.Meeting_CommitteeID(meetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting committee id of meeting: %w", err)
		}

		committeeManagerIDs, err := fetcher.Committee_ManagerIDs(committeeID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting user ids of committee managers: %w", err)
		}

		committeeManager := attribute.FuncUserIDs(committeeManagerIDs)

		_, isTemplateMeeting, err := fetcher.Meeting_TemplateForOrganizationID(meetingID).Value(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting template meeting: %w", err)
		}

		templateMeeting := attribute.FuncNotAllowed
		if isTemplateMeeting {
			templateMeeting = attribute.FuncIsCommitteeManager
		}

		result[i] = attribute.FuncOr(
			orgaManger,
			inMeeting,
			committeeManager,
			templateMeeting,
		)
	}

	return result, nil
}

func (m Meeting) modeC(ctx context.Context, fetcher *dsfetch.Fetch, meetingIDs []int) ([]attribute.Func, error) {
	result := make([]attribute.Func, len(meetingIDs))
	for i, id := range meetingIDs {
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, id)
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		result[i] = attribute.FuncInGroup(groupMap[perm.MeetingCanSeeFrontpage])
	}
	return result, nil
}

func (m Meeting) modeD(ctx context.Context, fetcher *dsfetch.Fetch, meetingIDs []int) ([]attribute.Func, error) {
	result := make([]attribute.Func, len(meetingIDs))
	for i, id := range meetingIDs {
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, id)
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		result[i] = attribute.FuncInGroup(groupMap[perm.MeetingCanSeeLivestream])
	}
	return result, nil
}
