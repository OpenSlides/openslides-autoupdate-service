package assignment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// Candidate is the collection for assignment candidates.
type Candidate struct {
	dp dataprovider.DataProvider
}

// NewCandidate creates a new AssignmentCandidate collection.
func NewCandidate(dp dataprovider.DataProvider) *Candidate {
	return &Candidate{
		dp: dp,
	}
}

// Connect connects the assignment_candidate routes.
func (c *Candidate) Connect(s collection.HandlerStore) {
	s.RegisterWriteHandler("assignment_candidate.create", collection.WriteCheckerFunc(c.create))
	s.RegisterWriteHandler("assignment_candidate.sort", collection.WriteCheckerFunc(c.sort))
	s.RegisterWriteHandler("assignment_candidate.delete", collection.WriteCheckerFunc(c.delete))

	s.RegisterReadHandler("assignment_candidate", c)
}

func (c *Candidate) create(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	superUser, err := c.dp.IsSuperuser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if superUser {
		return nil, nil
	}

	meetingID, err := c.dp.MeetingFromModel(ctx, "assignment/"+string(payload["assignment_id"]))
	if err != nil {
		return nil, fmt.Errorf("getting meeting id: %w", err)
	}

	var candidateUserID int
	if err := json.Unmarshal(payload["user_id"], &candidateUserID); err != nil {
		return nil, fmt.Errorf("no valid user id: %w", err)
	}

	perm := "assignments.can_nominate_other"
	if candidateUserID == userID {
		perm = "assignments.can_nominate_self"
	}

	if err := collection.EnsurePerms(ctx, c.dp, userID, meetingID, perm); err != nil {
		return nil, fmt.Errorf("ensure create permission: %w", err)
	}

	return nil, nil
}

func (c *Candidate) delete(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	superUser, err := c.dp.IsSuperuser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if superUser {
		return nil, nil
	}

	fqid := "assignment_candidate/" + string(payload["id"])

	var candidateUserID int
	if err := c.dp.Get(ctx, fqid+"/user_id", &candidateUserID); err != nil {
		return nil, fmt.Errorf("getting user id of candidate: %w", err)
	}

	perm := "assignments.can_nominate_other"
	if candidateUserID == userID {
		perm = "assignments.can_nominate_self"
	}

	meetingID, err := c.dp.MeetingFromModel(ctx, fqid)
	if err != nil {
		return nil, fmt.Errorf("getting meeting id: %w", err)
	}

	if err := collection.EnsurePerms(ctx, c.dp, userID, meetingID, perm); err != nil {
		return nil, fmt.Errorf("ensure delete permission: %w", err)
	}

	return nil, nil
}

func (c *Candidate) sort(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	superUser, err := c.dp.IsSuperuser(ctx, userID)
	if err != nil {
		return nil, err
	}
	if superUser {
		return nil, nil
	}

	meetingID, err := c.dp.MeetingFromModel(ctx, "assignment/"+string(payload["assignment_id"]))
	if err != nil {
		return nil, fmt.Errorf("getting meeting id: %w", err)
	}

	if err := collection.EnsurePerms(ctx, c.dp, userID, meetingID, "assignments.can_manage"); err != nil {
		return nil, fmt.Errorf("ensure delete permission: %w", err)
	}

	return nil, nil
}

// RestrictFQFields restricts all fields for assignment_candidates.
func (c *Candidate) RestrictFQFields(ctx context.Context, userID int, fqfields []collection.FQField, result map[string]bool) error {
	if len(fqfields) == 0 {
		return nil
	}

	// TODO
	// meetingID, err := c.dp.MeetingFromModel(ctx, "assignment_candidate/"+parts[1])
	// if err != nil {
	// 	return fmt.Errorf("getting meeting from assignment_candidate %s: %w", parts[1], err)
	// }

	// if err := collection.EnsurePerms(ctx, c.dp, userID, meetingID, "assignment.can_see"); err != nil {
	// 	return nil
	// }

	// for _, fqfield := range fqfields {
	// 	result[fqfield] = true
	// }
	return nil
}
