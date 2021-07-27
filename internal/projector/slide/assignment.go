package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbAssignment struct {
	ID                   int    `json:"id"`
	Title                string `json:"title"`
	Description          string `json:"description"`
	NumberPollCandidates bool   `json:"number_poll_candidates"`
	CandidateIDs         []int  `json:"candidate_ids"`
	AgendaItemID         int    `json:"agenda_item_id"`
}

type dbAssignmentCandidate struct {
	UserID int `json:"user_id"`
	Weight int `json:"weight"`
}

func assignmentFromMap(in map[string]json.RawMessage) (*dbAssignment, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding assignment data: %w", err)
	}

	var a dbAssignment
	if err := json.Unmarshal(bs, &a); err != nil {
		return nil, fmt.Errorf("decoding assignment data: %w", err)
	}
	return &a, nil
}

func assignmentCandidateFromMap(in map[string]json.RawMessage) (*dbAssignmentCandidate, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding assignment candidate data: %w", err)
	}

	var ac dbAssignmentCandidate
	if err := json.Unmarshal(bs, &ac); err != nil {
		return nil, fmt.Errorf("decoding assignment candidate data: %w", err)
	}
	return &ac, nil
}

// Assignment renders the assignment slide.
func Assignment(store *projector.SlideStore) {
	store.RegisterSliderFunc("assignment", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

		data := fetch.Object(
			ctx,
			[]string{
				"id",
				"title",
				"description",
				"number_poll_candidates",
				"candidate_ids",
			},
			p7on.ContentObjectID,
		)

		assignment, err := assignmentFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("get assignment: %w", err)
		}

		var allUsers []*dbAssignmentCandidate
		for _, ac := range assignment.CandidateIDs {
			data = fetch.Object(ctx, []string{"user_id", "weight"}, "assignment_candidate/%d", ac)
			userWeight, err := assignmentCandidateFromMap(data)
			if err != nil {
				return nil, nil, fmt.Errorf("get assignment candidate: %w", err)
			}
			allUsers = append(allUsers, userWeight)
		}
		sort.SliceStable(allUsers, func(i, j int) bool { return allUsers[i].Weight < allUsers[j].Weight })
		titler := store.GetTitleInformationFunc("user")
		if titler == nil {
			return nil, nil, fmt.Errorf("no titler function registered for user")
		}

		var users []string
		for _, candidate := range allUsers {
			user, err := NewUser(ctx, fetch, candidate.UserID, p7on.MeetingID)
			if err != nil {
				return nil, nil, fmt.Errorf("getting new user id: %w", err)
			}
			users = append(users, user.UserRepresentation(p7on.MeetingID))
		}

		out := struct {
			Title                string   `json:"title"`
			Description          string   `json:"description"`
			NumberPollCandidates bool     `json:"number_poll_candidates"`
			Candidates           []string `json:"candidates"`
		}{
			Title:                assignment.Title,
			Description:          assignment.Description,
			NumberPollCandidates: assignment.NumberPollCandidates,
			Candidates:           users,
		}

		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response slide assignment: %w", err)
		}
		return responseValue, fetch.Keys(), err
	})

	store.RegisterGetTitleInformationFunc("assignment", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, []string{"id", "title", "agenda_item_id"}, fqid)
		assignment, err := assignmentFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get assignment: %w", err)
		}

		if itemNumber == "" && assignment.AgendaItemID > 0 {
			itemNumber = fetch.String(ctx, "agenda_item/%d/item_number", assignment.AgendaItemID)
		}

		title := struct {
			Collection       string `json:"collection"`
			ContentObjectID  string `json:"content_object_id"`
			Title            string `json:"title"`
			AgendaItemNumber string `json:"agenda_item_number"`
		}{
			"assignment",
			fqid,
			assignment.Title,
			itemNumber,
		}

		bs, err := json.Marshal(title)
		if err != nil {
			return nil, fmt.Errorf("encoding title: %w", err)
		}
		return bs, err
	})
}
