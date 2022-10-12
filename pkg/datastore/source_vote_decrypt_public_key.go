package datastore

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

const votePubKeyPath = "/internal/vote/public_main_key"

// VoteDecryptPubKey fetches the public main key from vote decrypt via the
// vote-service.
type VoteDecryptPubKey struct {
	voteServiceURL string
	client         *http.Client
}

// NewVoteDecryptPubKeySource initializes the object.
func NewVoteDecryptPubKeySource(url string) *VoteDecryptPubKey {
	return &VoteDecryptPubKey{
		voteServiceURL: url,
		client:         &http.Client{},
	}
}

// Get is called when a key is not in the cache.
func (s *VoteDecryptPubKey) Get(ctx context.Context, keys ...Key) (map[Key][]byte, error) {
	out := make(map[Key][]byte, len(keys))
	for _, key := range keys {
		out[key] = nil

		if key.Collection != "organization" || key.Field != "vote_decrypt_public_main_key" {
			continue
		}

		req, err := http.NewRequestWithContext(ctx, "GET", s.voteServiceURL+votePubKeyPath, nil)
		if err != nil {
			return nil, fmt.Errorf("create request: %w", err)
		}

		resp, err := s.client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("sending request: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("vote service returned %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("reading vote service body: %w", err)
		}

		out[key] = body
	}
	return out, nil
}

// Update does nothing for this source.
func (s *VoteDecryptPubKey) Update(ctx context.Context) (map[Key][]byte, error) {
	<-ctx.Done()
	return nil, ctx.Err()
}
