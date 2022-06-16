package restrict

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/oserror"
)

const slowCalls = 100 * time.Millisecond

type timeCount struct {
	time  time.Duration
	count int
}

func (tc timeCount) MarshalJSON() ([]byte, error) {
	ms := tc.time.Milliseconds()
	decodable := struct {
		MS    int64 `json:"duraction_ms"`
		Count int   `json:"count"`
	}{
		ms,
		tc.count,
	}

	return json.Marshal(decodable)
}

func profile(request string, duration time.Duration, times map[string]timeCount) {
	timeStrings := make([]string, 0, len(times))
	for collection, tc := range times {
		timeStrings = append(timeStrings, fmt.Sprintf("%s: %d keys in %d ms", collection, tc.count, tc.time.Milliseconds()))
	}
	sort.Slice(timeStrings, func(i, j int) bool {
		return timeStrings[i] < timeStrings[j]
	})

	log.Printf("Profile: Restrict: Slow request:\nProfile: Request: %s\nProfile: Duration: %d ms\nProfile: %s\n", request, duration.Milliseconds(), strings.Join(timeStrings, "\nProfile: "))
}

func logTimes(ctx context.Context) func(map[string]timeCount) {
	start := time.Now()

	return func(times map[string]timeCount) {
		duration := time.Since(start)

		if times == nil || duration < slowCalls && !oserror.HasTagFromContext(ctx, "profile_restrict") {
			return
		}

		body, ok := oserror.BodyFromContext(ctx)
		if !ok {
			body = "unknown body, probably simple request"
		}
		profile(body, duration, times)
	}
}
