package restrict

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
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

	log.Printf("Profile: Restrict: Slow request:\nRequest: %s\nDuration: %d ms\n%s\n", request, duration.Milliseconds(), strings.Join(timeStrings, "\n"))
}
