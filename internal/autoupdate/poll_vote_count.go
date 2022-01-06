package autoupdate

import (
	"context"
	"fmt"
)

const path = "/internal/vote/vote_count"

func (a *Autoupdate) datastorePollVoteCount(ctx context.Context, fqfield string, changed map[string][]byte) ([]byte, error) {
	if changed != nil {
		return changed[fqfield], nil
	}

	values, err := a.datastore.RequestKeys(ctx, a.voteAddr+path, []string{fqfield})
	if err != nil {
		return nil, fmt.Errorf("loading key %q from vote-service: %w", fqfield, err)
	}
	return values[fqfield], nil
}
