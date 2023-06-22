package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
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
	UsersPdfWlanSsid                          string `json:"users_pdf_wlan_ssid"`
	UsersPdfWlanPassword                      string `json:"users_pdf_wlan_password"`
	UsersPdfWlanEncryption                    string `json:"users_pdf_wlan_encryption"`
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
	data := fetch.Object(ctx, fmt.Sprintf("meeting/%d", meetingID), fetchFields...)
	if err := fetch.Err(); err != nil {
		return nil, err
	}

	meeting, err = meetingFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("get meeting: %w", err)
	}
	return meeting, nil
}

// CurrentSpeakerChyron renders the current_speaker_chyron slide.
func WiFiAccessData(store *projector.SlideStore) {
	store.RegisterSliderFunc("wifi_access_data", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		meetingID, err := strconv.Atoi(strings.Split(p7on.ContentObjectID, "/")[1])
		if err != nil {
			return nil, fmt.Errorf("error in Atoi with ContentObjectID: %w", err)
		}

		meeting, err := getMeeting(ctx, fetch, meetingID, []string{
			"users_pdf_wlan_ssid",
			"users_pdf_wlan_password",
			"users_pdf_wlan_encryption",
		})

		out := struct {
			UsersPdfWlanSsid       string `json:"users_pdf_wlan_ssid,omitempty"`
			UsersPdfWlanPassword   string `json:"users_pdf_wlan_password,omitempty"`
			UsersPdfWlanEncryption string `json:"users_pdf_wlan_encryption,omitempty"`
		}{
			meeting.UsersPdfWlanSsid,
			meeting.UsersPdfWlanPassword,
			meeting.UsersPdfWlanEncryption,
		}

		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding response slide current_speaker_chyron: %w", err)
		}
		return responseValue, nil
	})
}
