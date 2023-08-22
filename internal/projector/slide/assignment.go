package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
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
	MeetingUserID int `json:"meeting_user_id"`
	Weight        int `json:"weight"`
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
	store.RegisterSliderFunc("assignment", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		data := fetch.Object(
			ctx,
			p7on.ContentObjectID,
			"id",
			"title",
			"description",
			"number_poll_candidates",
			"candidate_ids",
		)

		assignment, err := assignmentFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get assignment: %w", err)
		}

		var allUsers []*dbAssignmentCandidate
		for _, ac := range assignment.CandidateIDs {
			data = fetch.Object(ctx, fmt.Sprintf("assignment_candidate/%d", ac), "meeting_user_id", "weight")
			userWeight, err := assignmentCandidateFromMap(data)
			if err != nil {
				return nil, fmt.Errorf("get assignment candidate: %w", err)
			}
			allUsers = append(allUsers, userWeight)
		}

		sort.SliceStable(allUsers, func(i, j int) bool { return allUsers[i].Weight < allUsers[j].Weight })

		titler := store.GetTitleInformationFunc("user")
		if titler == nil {
			return nil, fmt.Errorf("no titler function registered for user")
		}

		var users []string
		for _, candidate := range allUsers {
			var userID int
			fetch.FetchIfExist(ctx, &userID, "meeting_user/%d/user_id", candidate.MeetingUserID)
			if err := fetch.Err(); err != nil {
				return nil, fmt.Errorf("getting user for meeting user %d: %w", candidate.MeetingUserID, err)
			}

			user, err := NewUser(ctx, fetch, userID, p7on.MeetingID)
			if err != nil {
				return nil, fmt.Errorf("getting new user id: %w", err)
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
			return nil, fmt.Errorf("encoding response slide assignment: %w", err)
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}
		return responseValue, nil
	})

	store.RegisterGetTitleInformationFunc("assignment", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, fqid, "id", "title", "agenda_item_id")
		assignment, err := assignmentFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get assignment: %w", err)
		}

		if itemNumber == "" && assignment.AgendaItemID > 0 {
			itemNumber = datastore.String(ctx, fetch.FetchIfExist, "agenda_item/%d/item_number", assignment.AgendaItemID)
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
		if err := fetch.Err(); err != nil {
			return nil, err
		}
		return bs, nil
	})
}
