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

type voteCountContent struct {
	ID    uint64      `json:"id"`
	Polls map[int]int `json:"polls"`
}

func (s *VoteCountSource) voteServiceConnect(ctx context.Context, blocking bool) (voteCountContent, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s.url(blocking), nil)
	if err != nil {
		return voteCountContent{}, fmt.Errorf("building request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		// TODO External Error
		return voteCountContent{}, fmt.Errorf("sending request to vote service: %w", err)
	}
	defer resp.Body.Close()

	var content voteCountContent
	if err := json.NewDecoder(resp.Body).Decode(&content); err != nil {
		// TODO External Error
		return voteCountContent{}, fmt.Errorf("decoding response body: %w", err)
	}

	return content, nil
}

// Get is called when a key is not in the cache.
func (s *VoteCountSource) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	content, err := s.voteServiceConnect(ctx, false)
	if err != nil {
		return nil, fmt.Errorf("connecting to vote service: %w", err)
	}

	out := make(map[Key][]byte, len(keys))
	for _, key := range keys {
		out[key] = nil

		if key.Collection != "poll" || key.Field != "vote_count" {
			continue
		}

		if count, ok := content.Polls[key.ID]; ok {
			out[key] = []byte(strconv.Itoa(count))
		}
	}
	return out, nil
}

// Update is called frequently and should block until there is new data.
func (s *VoteCountSource) Update(ctx context.Context) (map[Key][]byte, error) {
	content, err := s.voteServiceConnect(ctx, true)
	if err != nil {
		return nil, fmt.Errorf("connecting to vote service: %w", err)
	}

	s.id = content.ID

	out := make(map[Key][]byte, len(content.Polls))
	for pollID, count := range content.Polls {
		out[Key{"poll", pollID, "vote_count"}] = []byte(strconv.Itoa(count))
	}
	return out, nil
}

func (s *VoteCountSource) url(blocking bool) string {
	if blocking {
		return fmt.Sprintf("%s%s?id=%d", s.voteServiceURL, voteCountPath, s.id)
	}
	return fmt.Sprintf("%s%s", s.voteServiceURL, voteCountPath)
}
