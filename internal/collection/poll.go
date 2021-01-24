package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// Poll handels the permissions for poll objects.
func Poll(dp dataprovider.DataProvider) perm.ConnecterFunc {
	p := &poll{dp}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("poll", perm.CollectionFunc(p.readPoll))
		s.RegisterRestricter("option", perm.CollectionFunc(p.readOption))
		s.RegisterRestricter("vote", perm.CollectionFunc(p.readVote))

	}
}

type poll struct {
	dp dataprovider.DataProvider
}

func (p *poll) readPoll(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	restricted := map[string]bool{
		"votesvalid":   true,
		"votesinvalid": true,
		"votescast":    true,
		"voted_ids":    true,
	}

	return p.fields(fqfields, result, restricted, func(fqfield perm.FQField) (int, error) {
		return p.pollPerm(ctx, userID, fqfield.ID)
	})
}

func (p *poll) readOption(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	restricted := map[string]bool{
		"yes":      true,
		"no":       true,
		"abstain":  true,
		"vote_ids": true,
	}

	return p.fields(fqfields, result, restricted, func(fqfield perm.FQField) (int, error) {
		var pollID int
		if err := p.dp.Get(ctx, fmt.Sprintf("option/%d/poll_id", fqfield.ID), &pollID); err != nil {
			return 0, fmt.Errorf("getting poll id: %w", err)
		}
		return p.pollPerm(ctx, userID, pollID)
	})
}

func (p *poll) readVote(ctx context.Context, userID int, fqfields []perm.FQField, result map[string]bool) error {
	return perm.AllFields(fqfields, result, func(fqfield perm.FQField) (bool, error) {
		var optionID int
		if err := p.dp.Get(ctx, fmt.Sprintf("vote/%d/option_id", fqfield.ID), &optionID); err != nil {
			return false, fmt.Errorf("getting option id: %w", err)
		}

		var pollID int
		if err := p.dp.Get(ctx, fmt.Sprintf("option/%d/poll_id", optionID), &pollID); err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		perms, err := p.pollPerm(ctx, userID, pollID)
		if err != nil {
			return false, fmt.Errorf("getting poll permissions: %w", err)
		}

		if perms == 0 {
			return false, nil
		}

		if perms == 1 {
			return true, nil
		}

		var voteUserID int
		if err := p.dp.Get(ctx, fmt.Sprintf("vote/%d/user_id", fqfield.ID), &voteUserID); err != nil {
			return false, fmt.Errorf("getting vote user id: %w", err)
		}
		if voteUserID == userID {
			return true, nil
		}

		if err := p.dp.GetIfExist(ctx, fmt.Sprintf("vote/%d/delegated_user_id", fqfield.ID), &voteUserID); err != nil {
			return false, fmt.Errorf("getting vote delegated user id: %w", err)
		}
		if voteUserID == userID {
			return true, nil
		}

		return false, nil
	})
}

func (p *poll) fields(fqfields []perm.FQField, result map[string]bool, restricted map[string]bool, f func(perm.FQField) (int, error)) error {
	var hasPerm int
	var lastID int
	for _, fqfield := range fqfields {
		if lastID != fqfield.ID {
			lastID = fqfield.ID
			var err error
			hasPerm, err = f(fqfield)
			if err != nil {
				return fmt.Errorf("get permissions for poll %d: %w", fqfield.ID, err)
			}
		}
		if hasPerm == 1 || hasPerm == 2 && !restricted[fqfield.Field] {
			result[fqfield.String()] = true
		}
	}
	return nil
}

// pollPerm tells, if the user can see a poll.
// 0 = no perm
// 1 = all fields
// 2 = only some fields
func (p *poll) pollPerm(ctx context.Context, userID, pollID int) (int, error) {
	meetingID, err := p.dp.MeetingFromModel(ctx, fmt.Sprintf("poll/%d", pollID))
	if err != nil {
		return 0, fmt.Errorf("getting meeting id: %w", err)
	}

	var contentObjectID string
	if err := p.dp.GetIfExist(ctx, fmt.Sprintf("poll/%d/content_object_id", pollID), &contentObjectID); err != nil {
		return 0, fmt.Errorf("getting content object id: %w", err)
	}
	collection := strings.Split(contentObjectID, "/")[0]

	perms, err := perm.New(ctx, p.dp, userID, meetingID)
	if err != nil {
		return 0, fmt.Errorf("getting perms: %w", err)
	}

	if perms.Has(p.canManage(collection)) {
		return 1, nil
	}

	canSee, err := canSeePoll(ctx, p.dp, perms, userID, contentObjectID)
	if err != nil {
		return 0, fmt.Errorf("getting can see perm: %w", err)
	}
	if !canSee {
		return 0, nil
	}

	var state string
	if err := p.dp.Get(ctx, fmt.Sprintf("poll/%d/state", pollID), &state); err != nil {
		return 0, fmt.Errorf("getting poll state: %w", err)
	}

	if state == "published" {
		return 1, nil
	}
	return 2, nil
}

func canSeePoll(ctx context.Context, dp dataprovider.DataProvider, perms *perm.Permission, userID int, objectID string) (bool, error) {
	var collection string
	var id int
	if objectID != "" {
		parts := strings.Split(objectID, "/")
		collection = parts[0]
		var err error
		id, err = strconv.Atoi(parts[1])
		if err != nil {
			return false, fmt.Errorf("invalid object id: %w", err)
		}
	}

	if collection == "motion" {
		meetingID, err := dp.MeetingFromModel(ctx, fmt.Sprintf("motion/%d", id))
		if err != nil {
			return false, fmt.Errorf("getting meetingID from motion: %w", err)
		}

		perms, err := perm.New(ctx, dp, userID, meetingID)
		if err != nil {
			return false, fmt.Errorf("getting user permissions: %w", err)
		}

		return canSeeMotion(ctx, dp, userID, id, perms)
	}

	perm := "agenda_item.can_see"
	if collection == "assignment" {
		perm = "assignment.can_see"
	}
	return perms.Has(perm), nil
}

func (p *poll) canManage(collection string) string {
	if collection == "motion" {
		return "motion.can_manage_polls"
	}
	if collection == "assignment" {
		return "assignment.can_manage"
	}
	return "agenda_item.can_manage"
}

func canSeePolls(ctx context.Context, dp dataprovider.DataProvider, perms *perm.Permission, userID int, ids []int) (bool, error) {
	for _, id := range ids {
		var contentObject string
		if err := dp.GetIfExist(ctx, fmt.Sprintf("poll/%d/content_object_id", id), &contentObject); err != nil {
			return false, fmt.Errorf("getting motion id: %w", err)
		}

		b, err := canSeePoll(ctx, dp, perms, userID, contentObject)
		if err != nil {
			return false, fmt.Errorf("can see poll %d: %w", id, err)
		}
		if b {
			return true, nil
		}
	}
	return false, nil
}

func canSeePollOptions(ctx context.Context, dp dataprovider.DataProvider, perms *perm.Permission, userID int, ids []int) (bool, error) {
	for _, id := range ids {
		var pollID int
		if err := dp.Get(ctx, fmt.Sprintf("option/%d/poll_id", id), &pollID); err != nil {
			return false, fmt.Errorf("getting poll id: %w", err)
		}

		var contentObject string
		if err := dp.GetIfExist(ctx, fmt.Sprintf("poll/%d/content_object_id", pollID), &contentObject); err != nil {
			return false, fmt.Errorf("getting motion id: %w", err)
		}

		b, err := canSeePoll(ctx, dp, perms, userID, contentObject)
		if err != nil {
			return false, fmt.Errorf("can see poll %d: %w", id, err)
		}
		if b {
			return true, nil
		}
	}
	return false, nil
}
