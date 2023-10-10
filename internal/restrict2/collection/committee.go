package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Committee handels permission for committees.
//
// See user can see a committee, if he is in committee/user_ids or have the OML
// can_manage_users or higher.
//
// Mode A: The user can see the committee.
//
// Mode B: The user must have the OML `can_manage_organization` or higher or the
// CML `can_manage` in the committee.
type Committee struct{}

// Name returns the collection name.
func (c Committee) Name() string {
	return "committee"
}

// MeetingID returns the meetingID for the object.
func (c Committee) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	return 0, false, nil
}

// Modes returns a map from all known modes to there restricter.
func (c Committee) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return c.see
	case "B":
		return c.modeB
	}
	return nil
}

func (c Committee) see(ctx context.Context, fetcher *dsfetch.Fetch, committeeIDs []int) ([]attribute.Func, error) {
	userIDs := make([][]int, len(committeeIDs))
	for i, id := range committeeIDs {
		if id == 0 {
			continue
		}
		fetcher.Committee_UserIDs(id).Lazy(&userIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching committee data: %w", err)
	}

	oml := attribute.FuncOrgaLevel(perm.OMLCanManageUsers)

	attr := make([]attribute.Func, len(committeeIDs))
	for i, id := range committeeIDs {
		if id == 0 {
			continue
		}

		attr[i] = attribute.FuncOr(
			oml,
			// TODO: This are a lot of users. Is this ok?
			attribute.FuncUserIDs(userIDs[i]),
		)
	}
	return attr, nil
}

func (c Committee) modeB(ctx context.Context, fetcher *dsfetch.Fetch, committeeIDs []int) ([]attribute.Func, error) {
	managerIDs := make([][]int, len(committeeIDs))
	for i, id := range committeeIDs {
		if id == 0 {
			continue
		}
		fetcher.Committee_ManagerIDs(id).Lazy(&managerIDs[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching committee data: %w", err)
	}

	oml := attribute.FuncOrgaLevel(perm.OMLCanManageOrganization)

	attr := make([]attribute.Func, len(committeeIDs))
	for i, id := range committeeIDs {
		if id == 0 {
			continue
		}

		attr[i] = attribute.FuncOr(
			oml,
			attribute.FuncUserIDs(managerIDs[i]),
		)
	}
	return attr, nil
}
