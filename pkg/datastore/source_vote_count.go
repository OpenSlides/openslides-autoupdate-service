package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const voteCountPath = "/internal/vote/vote_count"

// VoteCountSource is a datastore source for the poll/vote_count value.
type VoteCountSource struct {
	voteServiceURL string
	client         *http.Client
	id             uint64

	mu        sync.Mutex
	voteCount map[int]int
	update    chan map[int]int
}

// NewVoteCountSource initializes the object.
func NewVoteCountSource(url string) *VoteCountSource {
	return &VoteCountSource{
		voteServiceURL: url,
		client:         &http.Client{},
		update:         make(chan map[int]int, 1),
	}
}

// Connect creates a connection to the vote service and makes sure, it stays
// open.
func (s *VoteCountSource) Connect(ctx context.Context, eventProvider func() (<-chan time.Time, func() bool), errHandler func(error)) {
	for ctx.Err() == nil {
		if err := s.connect(ctx); err != nil {
			errHandler(fmt.Errorf("connecting to vote service: %w", err))
		}

		s.wait(ctx, eventProvider)
	}
}

// wait waits for an event in s.eventProvider.
func (s *VoteCountSource) wait(ctx context.Context, eventProvider func() (<-chan time.Time, func() bool)) {
	event, close := eventProvider()
	defer close()

	select {
	case <-ctx.Done():
	case <-event:
	}
}

func (s *VoteCountSource) connect(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", s.voteServiceURL+voteCountPath, nil)
	if err != nil {
		return fmt.Errorf("building request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		// TODO External Error
		return fmt.Errorf("sending request to vote service: %w", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	s.voteCount = make(map[int]int)
	for {
		var counts map[int]int
		if err := decoder.Decode(&counts); err != nil {
			if err == io.EOF {
				return nil
			}
			return fmt.Errorf("decoding poll data: %w", err)
		}

		s.mu.Lock()

		for k, v := range counts {
			if v == 0 {
				delete(s.voteCount, k)
				continue
			}
			s.voteCount[k] = v
		}

		s.mu.Unlock()

		select {
		case s.update <- counts:
		default:
		}
	}
}

// Get is called when a key is not in the cache.
func (s *VoteCountSource) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make(map[Key][]byte, len(keys))
	for _, key := range keys {
		out[key] = nil

		if key.Collection != "poll" || key.Field != "vote_count" {
			continue
		}

		if count, ok := s.voteCount[key.ID]; ok {
			out[key] = []byte(strconv.Itoa(count))
		}
	}
	return out, nil
}

// Update is called frequently and should block until there is new data.
func (s *VoteCountSource) Update(ctx context.Context) (map[Key][]byte, error) {
	var data map[int]int
	select {
	case <-ctx.Done():
		return nil, ctx.Err()

	case data = <-s.update:
	}

	out := make(map[Key][]byte, len(data))
	for pollID, count := range data {
		bs := []byte(strconv.Itoa(count))
		if count == 0 {
			bs = nil
		}
		out[Key{"poll", pollID, "vote_count"}] = bs
	}
	return out, nil
}
