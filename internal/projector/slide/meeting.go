package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbMeeting struct {
	ID                                        int    `json:"id"`
	MotionsEnableTextOnProjector              bool   `json:"motions_enable_text_on_projector"`
	MotionsEnableReasonOnProjector            bool   `json:"motions_enable_reason_on_projector"`
	MotionsShowReferringMotions               bool   `json:"motions_show_referring_motions"`
	MotionsEnableRecommendationOnProjector    bool   `json:"motions_enable_recommendation_on_projector"`
	MotionsStatuteRecommendationsBy           string `json:"motions_statute_recommendations_by"`
	MotionsRecommendationsBy                  string `json:"motions_recommendations_by"`
	MotionsEnableSideboxOnProjector           bool   `json:"motions_enable_sidebox_on_projector"`
	MotionsLineLength                         int    `json:"motions_line_length"`
	MotionsPreamble                           string `json:"motions_preamble"`
	MotionsDefaultLineNumbering               string `json:"motions_default_line_numbering"`
	ListOfSpeakersAmountNextOnProjector       int    `json:"list_of_speakers_amount_next_on_projector"`
	ListOfSpeakersAmountLastOnProjector       int    `json:"list_of_speakers_amount_last_on_projector"`
	ListOfSpeakersShowAmountOfSpeakersOnSlide bool   `json:"list_of_speakers_show_amount_of_speakers_on_slide"`
}

func meetingFromMap(in map[string]json.RawMessage) (*dbMeeting, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding meeting data: %w", err)
	}

	var me dbMeeting
	if err := json.Unmarshal(bs, &me); err != nil {
		return nil, fmt.Errorf("decoding meeting data: %w", err)
	}
	return &me, nil
}

func getMeeting(ctx context.Context, fetch *datastore.Fetcher, meetingID int, fetchFields []string) (meeting *dbMeeting, err error) {
	defer func() {
		if err == nil {
			err = fetch.Error()
		}
	}()

	data := fetch.Object(ctx, fetchFields, "meeting/%d", meetingID)

	meeting, err = meetingFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("get meeting: %w", err)
	}
	return meeting, nil
}
