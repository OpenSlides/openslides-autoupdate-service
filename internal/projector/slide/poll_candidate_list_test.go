package slide_test

import (
	"context"
	"testing"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/slide"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsmock"
	"github.com/stretchr/testify/assert"
)

func TestPollCandidateList(t *testing.T) {
	s := new(projector.SlideStore)
	slide.PollCandidateList(s)

	pollCandidateListTitler := s.GetTitleInformationFunc("poll_candidate_list")

	for _, tt := range []struct {
		name   string
		data   map[string]string
		expect string
	}{
		{
			"No Candidates",
			map[string]string{
				"poll_candidate_list/1/id":                 "1",
				"poll_candidate_list/1/poll_candidate_ids": `[]`,
			},
			`{"collection":"poll_candidate_list","entries_amount":0,"content_object_id":"poll_candidate_list/1"}`,
		},
		{
			"One Candidate",
			map[string]string{
				"poll_candidate_list/1/id":                 "1",
				"poll_candidate_list/1/poll_candidate_ids": `[1]`,
			},
			`{"collection":"poll_candidate_list","entries_amount":1,"content_object_id":"poll_candidate_list/1"}`,
		},
		{
			"Two Candidates",
			map[string]string{
				"poll_candidate_list/1/id":                 "1",
				"poll_candidate_list/1/poll_candidate_ids": `[1, 3]`,
			},
			`{"collection":"poll_candidate_list","entries_amount":2,"content_object_id":"poll_candidate_list/1"}`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ds := dsmock.Stub(convertData(tt.data))
			fetch := datastore.NewFetcher(ds)

			bs, err := pollCandidateListTitler.GetTitleInformation(context.Background(), fetch, "poll_candidate_list/1", "1", 222)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expect, string(bs))
		})
	}
}
