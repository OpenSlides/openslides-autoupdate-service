package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

// DbPollCandidateList is the class with methods to get needed PollCandidateList Informations
type DbPollCandidateList struct {
	PollCandidateIDs []int `json:"poll_candidate_ids"`
}

// NewPollCandidateList gets the poll_candidate_list from datastore and return it as DbPollCandidateList struct
// together with keys and error.
func NewPollCandidateList(ctx context.Context, fetch *datastore.Fetcher, id int) (*DbPollCandidateList, error) {
	fields := []string{
		"poll_candidate_ids",
	}

	data := fetch.Object(ctx, fmt.Sprintf("poll_candidate_list/%d", id), fields...)
	if err := fetch.Err(); err != nil {
		return nil, fmt.Errorf("getting poll_candidate_list object: %w", err)
	}

	bs, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("encoding poll_candidate_list data: %w", err)
	}

	var u DbPollCandidateList
	if err := json.Unmarshal(bs, &u); err != nil {
		return nil, fmt.Errorf("decoding poll_candidate_list data: %w", err)
	}
	return &u, nil
}

// PollCandidateList renders the poll_candidate_list slide.
func PollCandidateList(store *projector.SlideStore) {
	store.RegisterGetTitleInformationFunc("poll_candidate_list", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		id, err := strconv.Atoi(strings.Split(fqid, "/")[1])
		if err != nil {
			return nil, fmt.Errorf("getting poll_candidate_list id: %w", err)
		}

		poll_candidate_list, err := NewPollCandidateList(ctx, fetch, id)
		if err != nil {
			return nil, fmt.Errorf("loading poll_candidate_list: %w", err)
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		out := struct {
			Collection      string `json:"collection"`
			ContentObjectID string `json:"content_object_id"`
			EntriesAmount   int    `json:"entries_amount"`
		}{
			"poll_candidate_list",
			fqid,
			len(poll_candidate_list.PollCandidateIDs),
		}
		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding title: %w", err)
		}
		return responseValue, err
	})
}
