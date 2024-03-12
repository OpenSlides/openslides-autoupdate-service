package datastore

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/environment"
)

const votePubKeyPath = "/internal/vote/public_main_key"

// VoteDecryptPubKey fetches the public main key from vote decrypt via the
// vote-service.
type VoteDecryptPubKey struct {
	voteServiceURL string
	client         *http.Client
}

// NewVoteDecryptPubKeySource initializes the object.
func NewVoteDecryptPubKeySource(lookup environment.Environmenter) *VoteDecryptPubKey {
	url := fmt.Sprintf(
		"%s://%s:%s",
		envVoteProtocol.Value(lookup),
		envVoteHost.Value(lookup),
		envVotePort.Value(lookup),
	)

	return &VoteDecryptPubKey{
		voteServiceURL: url,
		client:         &http.Client{},
	}
}

// Get is called when a key is not in the cache.
func (s *VoteDecryptPubKey) Get(ctx context.Context, keys ...dskey.Key) (map[dskey.Key][]byte, error) {
	out := make(map[dskey.Key][]byte, len(keys))
	for _, key := range keys {
		out[key] = nil

		if key.Collection() != "organization" || key.Field() != "vote_decrypt_public_main_key" {
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
func (s *VoteDecryptPubKey) Update(ctx context.Context, updateFn func(map[dskey.Key][]byte, error)) {
	<-ctx.Done()
}
