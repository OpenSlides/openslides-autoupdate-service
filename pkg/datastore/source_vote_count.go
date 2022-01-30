package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const voteCountPath = "/internal/vote/vote_count"

// VoteCountSource is a datastore source for the poll/vote_count value.
type VoteCountSource struct {
	voteServiceURL string
	client         *http.Client
	id             uint64
}

// NewVoteCountSource initializes the object.
func NewVoteCountSource(url string) *VoteCountSource {
	return &VoteCountSource{
		voteServiceURL: url,
		client:         &http.Client{},
	}
}

// Get is called when a key is not in the cache.
func (s *VoteCountSource) Get(ctx context.Context, keys ...string) (map[string][]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s.url(false), nil)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request to vote service: %w", err)
	}
	defer resp.Body.Close()

	var content struct {
		ID    uint64      `json:"id"`
		Polls map[int]int `json:"polls"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return nil, fmt.Errorf("decoding response body: %w", err)
	}

	out := make(map[string][]byte, len(keys))
	for _, key := range keys {
		out[key] = nil

		var pollID int
		if _, err := fmt.Sscanf(key, "poll/%d/vote_count", &pollID); err != nil {
			continue
		}

		if count, ok := content.Polls[pollID]; ok {
			out[key] = []byte(strconv.Itoa(count))
		}
	}
	return out, nil
}

// Update is called frequently and should block until there is new data.
func (s *VoteCountSource) Update(ctx context.Context) (map[string][]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s.url(true), nil)
	if err != nil {
		return nil, fmt.Errorf("building request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request to vote service: %w", err)
	}
	defer resp.Body.Close()

	var content struct {
		ID    uint64      `json:"id"`
		Polls map[int]int `json:"polls"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return nil, fmt.Errorf("decoding response body: %w", err)
	}

	s.id = content.ID

	out := make(map[string][]byte, len(content.Polls))
	for pollID, count := range content.Polls {
		out[fmt.Sprintf("poll/%d/vote_count", pollID)] = []byte(strconv.Itoa(count))
	}
	return out, nil
}

func (s *VoteCountSource) url(withID bool) string {
	if withID {
		return fmt.Sprintf("%s%s?id=%d", s.voteServiceURL, voteCountPath, s.id)
	}
	return fmt.Sprintf("%s%s?id=0", s.voteServiceURL, voteCountPath)
}
