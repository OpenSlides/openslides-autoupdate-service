package collection

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/dataprovider"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/permission/perm"
)

// Poll handels the permissions for poll objects.
func Poll(dp dataprovider.DataProvider) perm.ConnecterFunc {
	p := &poll{dp}
	return func(s perm.HandlerStore) {
		s.RegisterRestricter("poll", perm.CollectionFunc(p.readPoll))
		s.RegisterRestricter("option", perm.CollectionFunc(p.readOption))
		s.RegisterRestricter("vote", perm.CollectionFunc(p.readVote))

		s.RegisterAction("poll.delete", perm.ActionFunc(p.pollDelete))
		s.RegisterAction("option.delete", perm.ActionFunc(p.optionDelete))
		s.RegisterAction("vote.delete", perm.ActionFunc(p.voteDelete))
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
		pollID, err := pollIDFromOption(ctx, p.dp, fqfield.ID)
		if err != nil {
			return 0, fmt.Errorf("fetch poll id: %w", err)
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

		pollID, err := pollIDFromOption(ctx, p.dp, optionID)
		if err != nil {
			return false, fmt.Errorf("fetch poll id: %w", err)
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
				return fmt.Errorf("get permissions for %s: %w", fqfield.FQID(), err)
			}
		}
		if hasPerm == 1 || hasPerm == 2 && !restricted[fqfield.Field] {
			result[fqfield.String()] = true
		}
	}
	return nil
}

func (p *poll) pollDeleteWithID(ctx context.Context, userID int, pollID int) (bool, error) {
	fqid := "poll/" + strconv.Itoa(pollID)
	meetingID, err := p.dp.MeetingFromModel(ctx, fqid)
	if err != nil {
		return false, fmt.Errorf("getting meeting id from %s: %w", fqid, err)
	}

	var contentObjectID string
	if err := p.dp.GetIfExist(ctx, fqid+"/content_object_id", &contentObjectID); err != nil {
		return false, fmt.Errorf("getting content object id: %w", err)
	}
	collection := strings.Split(contentObjectID, "/")[0]
	requiredPerm := p.canManage(collection)
	return perm.HasPerm(ctx, p.dp, userID, meetingID, requiredPerm)
}

func (p *poll) pollDelete(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var pollID int
	if err := json.Unmarshal(payload["id"], &pollID); err != nil {
		return false, fmt.Errorf("no id in payload: %v", err)
	}

	return p.pollDeleteWithID(ctx, userID, pollID)
}

func (p *poll) optionDelete(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var optionID int
	if err := json.Unmarshal(payload["id"], &optionID); err != nil {
		return false, fmt.Errorf("no id in payload: %v", err)
	}

	var pollID int
	if err := p.dp.Get(ctx, fmt.Sprintf("option/%d/poll_id", optionID), &pollID); err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}

	return p.pollDeleteWithID(ctx, userID, pollID)
}

func (p *poll) voteDelete(ctx context.Context, userID int, payload map[string]json.RawMessage) (bool, error) {
	var voteID int
	if err := json.Unmarshal(payload["id"], &voteID); err != nil {
		return false, fmt.Errorf("no id in payload: %v", err)
	}

	var optionID int
	if err := p.dp.Get(ctx, fmt.Sprintf("vote/%d/option_id", voteID), &optionID); err != nil {
		return false, fmt.Errorf("getting option id: %w", err)
	}

	var pollID int
	if err := p.dp.Get(ctx, fmt.Sprintf("option/%d/poll_id", optionID), &pollID); err != nil {
		return false, fmt.Errorf("getting poll id: %w", err)
	}

	return p.pollDeleteWithID(ctx, userID, pollID)
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

	canSeePerm := perm.AgendaItemCanSee
	if collection == "assignment" {
		canSeePerm = perm.AssignmentCanSee
	}
	return perms.Has(canSeePerm), nil
}

func (p *poll) canManage(collection string) perm.TPermission {
	if collection == "motion" {
		return perm.MotionCanManagePolls
	}
	if collection == "assignment" {
		return perm.AssignmentCanManage
	}
	return perm.AgendaItemCanManage
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
		pollID, err := pollIDFromOption(ctx, dp, id)
		if err != nil {
			return false, fmt.Errorf("fetch poll id: %w", err)
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

// pollIDFromOption returns the poll id from an option.
//
// An option is linked to the poll via poll_id xor
// used_as_global_option_in_poll_id.
func pollIDFromOption(ctx context.Context, dp dataprovider.DataProvider, optionID int) (int, error) {
	var pollID int
	if err := dp.GetIfExist(ctx, fmt.Sprintf("option/%d/poll_id", optionID), &pollID); err != nil {
		return 0, fmt.Errorf("poll id from field `poll_id`: %w", err)
	}

	if pollID != 0 {
		// Option has the field `poll_id`
		return pollID, nil
	}

	if err := dp.Get(ctx, fmt.Sprintf("option/%d/used_as_global_option_in_poll_id", optionID), &pollID); err != nil {
		return 0, fmt.Errorf("poll id from field `used_as_global_option_in_poll_id`: %w", err)
	}
	return pollID, nil
}
