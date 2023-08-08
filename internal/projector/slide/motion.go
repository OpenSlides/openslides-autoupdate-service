package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

type dbMotionState struct {
	RecommendationLabel string             `json:"recommendation_label"`
	CSSClass            string             `json:"css_class"`
	MotionStateWork     *dbMotionStateWork `json:",omitempty"`
}
type dbMotionStateWork struct {
	ShowRecommendationExtensionField bool `json:"show_recommendation_extension_field"`
}

func motionStateFromMap(in map[string]json.RawMessage) (*dbMotionState, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding motion state data: %w", err)
	}

	var ms dbMotionState
	var msWork dbMotionStateWork
	ms.MotionStateWork = &msWork
	if err := json.Unmarshal(bs, &ms); err != nil {
		return nil, fmt.Errorf("decoding motion state data: %w", err)
	}
	if err := json.Unmarshal(bs, &msWork); err != nil {
		return nil, fmt.Errorf("decoding motion state work data: %w", err)
	}
	return &ms, nil
}

type dbMotionStatuteParagraph struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type dbMotionChangeRecommendation struct {
	ID               int    `json:"id"`
	Rejected         bool   `json:"rejected"`
	Type             string `json:"type"`
	OtherDescription string `json:"other_description"`
	LineFrom         int    `json:"line_from"`
	LineTo           int    `json:"line_to"`
	Text             string `json:"text"`
	CreationTime     int    `json:"creation_time"`
}

type amendmentsType struct {
	ID                      int                            `json:"id"`
	Title                   string                         `json:"title"`
	Number                  string                         `json:"number"`
	AmendmentParagraph      json.RawMessage                `json:"amendment_paragraphs"`
	ChangeRecommendations   []dbMotionChangeRecommendation `json:"change_recommendations"`
	MergeAmendmentIntoFinal string                         `json:"merge_amendment_into_final"`
	MergeAmendmentIntoDiff  string                         `json:"merge_amendment_into_diff"`
}
type leadMotionType struct {
	Title  string `json:"title"`
	Number string `json:"number"`
	Text   string `json:"text"`
}
type dbMotionWork struct {
	MeetingID                                    int      `json:"meeting_id"`
	LeadMotionID                                 int      `json:"lead_motion_id"`
	StatuteParagraphID                           int      `json:"statute_paragraph_id"`
	ChangeRecommendationIDS                      []int    `json:"change_recommendation_ids"`
	AmendmentIDS                                 []int    `json:"amendment_ids"`
	SubmitterIDS                                 []int    `json:"submitter_ids"`
	ReferencedInMotionRecommendationExtensionIDS []int    `json:"referenced_in_motion_recommendation_extension_ids"`
	RecommendationID                             int      `json:"recommendation_id"`
	RecommendationExtensionReferenceIDS          []string `json:"recommendation_extension_reference_ids"`
	RecommendationExtension                      string   `json:"recommendation_extension"`
	StateID                                      int      `json:"state_id"`
	AgendaItemID                                 int      `json:"agenda_item_id"`
}
type dbMotion struct {
	ID                               int                            `json:"id"`
	Title                            string                         `json:"title"`
	Number                           string                         `json:"number"`
	Submitters                       []string                       `json:"submitters"`
	ShowSidebox                      bool                           `json:"show_sidebox"`
	LineLength                       int                            `json:"line_length"`
	Preamble                         string                         `json:"preamble"`
	LineNumbering                    string                         `json:"line_numbering"`
	AmendmentParagraph               json.RawMessage                `json:"amendment_paragraphs,omitempty"`
	LeadMotion                       *leadMotionType                `json:"lead_motion,omitempty"`
	BaseStatute                      *dbMotionStatuteParagraph      `json:"base_statute,omitempty"`
	ChangeRecommendations            []dbMotionChangeRecommendation `json:"change_recommendations"`
	Amendments                       []amendmentsType               `json:"amendments,omitempty"`
	RecommendationReferencingMotions []json.RawMessage              `json:"recommendation_referencing_motions,omitempty"`
	RecommendationLabel              string                         `json:"recommendation_label,omitempty"`
	RecommendationExtension          string                         `json:"recommendation_extension,omitempty"`
	RecommendationReferencedMotions  map[string]json.RawMessage     `json:"recommendation_referenced_motions,omitempty"`
	Recommender                      string                         `json:"recommender,omitempty"`
	Text                             string                         `json:"text,omitempty"`
	Reason                           string                         `json:"reason,omitempty"`
	ModifiedFinalVersion             string                         `json:"modified_final_version,omitempty"`
	MotionWork                       *dbMotionWork                  `json:",omitempty"`
}

func motionFromMap(in map[string]json.RawMessage) (*dbMotion, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding motion data: %w", err)
	}

	var mo dbMotion
	var moWork dbMotionWork
	mo.MotionWork = &moWork
	if err := json.Unmarshal(bs, &mo); err != nil {
		return nil, fmt.Errorf("decoding motion data: %w", err)
	}
	if err := json.Unmarshal(bs, &moWork); err != nil {
		return nil, fmt.Errorf("decoding motion work data: %w", err)
	}
	return &mo, nil
}

// Motion renders the motion slide.
func Motion(store *projector.SlideStore) {
	store.RegisterSliderFunc("motion", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		meeting, err := getMeeting(ctx, fetch, p7on.MeetingID, []string{
			"motions_enable_text_on_projector",
			"motions_enable_reason_on_projector",
			"motions_show_referring_motions",
			"motions_enable_recommendation_on_projector",
			"motions_statute_recommendations_by",
			"motions_recommendations_by",
			"motions_enable_sidebox_on_projector",
			"motions_line_length",
			"motions_preamble",
			"motions_default_line_numbering",
		})
		if err != nil {
			return nil, fmt.Errorf("getMeeting: %w", err)
		}

		var options struct {
			Mode string `json:"mode"`
		}
		if p7on.Options != nil {
			if err := json.Unmarshal(p7on.Options, &options); err != nil {
				return nil, fmt.Errorf("decoding projection options: %w", err)
			}
		}

		fetchFields := []string{
			"id",
			"title",
			"number",
			"meeting_id",
			"lead_motion_id",
			"statute_paragraph_id",
			"amendment_paragraphs",
			"change_recommendation_ids",
			"amendment_ids",
			"submitter_ids",
			"referenced_in_motion_recommendation_extension_ids",
			"recommendation_id",
			"recommendation_extension",
			"recommendation_extension_reference_ids",
		}
		if meeting.MotionsEnableTextOnProjector {
			fetchFields = append(fetchFields, "text")
		}
		if meeting.MotionsEnableReasonOnProjector {
			fetchFields = append(fetchFields, "reason")
		}
		if p7on.Options != nil && options.Mode == "final" {
			fetchFields = append(fetchFields, "modified_final_version")
		}

		data := fetch.Object(ctx, p7on.ContentObjectID, fetchFields...)

		motion, err := motionFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motion: %w", err)
		}
		motion.RecommendationExtension = "" // will be (re-)filled conditionally

		fillMotionFromMeeting(motion, meeting)

		if err := fillSubmitters(ctx, fetch, motion); err != nil {
			return nil, fmt.Errorf("fillSubmitters: %w", err)
		}

		if err := fillLeadMotion(ctx, fetch, motion); err != nil {
			return nil, fmt.Errorf("fillLeadMotion: %w", err)
		}

		if err := fillBaseStatute(ctx, fetch, motion); err != nil {
			return nil, fmt.Errorf("fillBaseStatute: %w", err)
		}

		if err := fillChangeRecommendations(ctx, fetch, motion); err != nil {
			return nil, fmt.Errorf("fillChangeRecommendations: %w", err)
		}

		if err := fillAmendments(ctx, fetch, motion); err != nil {
			return nil, fmt.Errorf("fillAmendments: %w", err)
		}

		titlerMotion := store.GetTitleInformationFunc("motion")
		if titlerMotion == nil {
			return nil, fmt.Errorf("no titler function registered for motion")
		}

		if meeting.MotionsShowReferringMotions && len(motion.MotionWork.ReferencedInMotionRecommendationExtensionIDS) > 0 {
			err = fillRecommendationReferencingMotions(ctx, fetch, titlerMotion, motion)
			if err != nil {
				return nil, fmt.Errorf("FillRecommendationReferencingMotions: %w", err)
			}
		}

		if meeting.MotionsEnableRecommendationOnProjector && motion.MotionWork.RecommendationID > 0 {
			err = fillRecommendationLabelEtc(ctx, fetch, titlerMotion, motion, meeting)
			if err != nil {
				return nil, fmt.Errorf("RecommendationLabelEtc: %w", err)
			}
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		motion.MotionWork = nil // do not export worker fields
		responseValue, err := json.Marshal(motion)
		if err != nil {
			return nil, fmt.Errorf("encoding response for slide motion: %w", err)
		}
		return responseValue, err
	})

	store.RegisterGetTitleInformationFunc("motion", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, fqid, "id", "number", "title", "agenda_item_id")
		motion, err := motionFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motion: %w", err)
		}

		if itemNumber == "" && motion.MotionWork.AgendaItemID > 0 {
			itemNumber = datastore.String(ctx, fetch.FetchIfExist, "agenda_item/%d/item_number", motion.MotionWork.AgendaItemID)
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		title := struct {
			Collection       string `json:"collection"`
			ContentObjectID  string `json:"content_object_id"`
			Title            string `json:"title"`
			Number           string `json:"number"`
			AgendaItemNumber string `json:"agenda_item_number"`
		}{
			"motion",
			fqid,
			motion.Title,
			motion.Number,
			itemNumber,
		}

		bs, err := json.Marshal(title)
		if err != nil {
			return nil, fmt.Errorf("encoding title: %w", err)
		}
		return bs, err
	})
}

// fillMotionFrom Meeting transfers the needed values from meeting object to motion object.
func fillMotionFromMeeting(motion *dbMotion, meeting *dbMeeting) {
	motion.ShowSidebox = meeting.MotionsEnableSideboxOnProjector
	motion.LineLength = meeting.MotionsLineLength
	motion.Preamble = meeting.MotionsPreamble
	motion.LineNumbering = meeting.MotionsDefaultLineNumbering
}

func fillLeadMotion(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) error {
	if motion.MotionWork.LeadMotionID == 0 {
		return nil
	}
	data := fetch.Object(ctx, fmt.Sprintf("motion/%d", motion.MotionWork.LeadMotionID), "title", "number", "text")
	bs, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encoding LeadMotion data: %w", err)
	}
	var mo leadMotionType
	if err := json.Unmarshal(bs, &mo); err != nil {
		return fmt.Errorf("decoding LeadMotion data: %w", err)
	}
	motion.LeadMotion = &mo
	return nil
}

func fillBaseStatute(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) error {
	if motion.MotionWork.StatuteParagraphID == 0 {
		return nil
	}
	data := fetch.Object(ctx, fmt.Sprintf("motion_statute_paragraph/%d", motion.MotionWork.StatuteParagraphID), "title", "text")
	bs, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encoding BaseStatute data: %w", err)
	}
	var mo dbMotionStatuteParagraph
	if err := json.Unmarshal(bs, &mo); err != nil {
		return fmt.Errorf("decoding BaseStatute data: %w", err)
	}
	motion.BaseStatute = &mo
	return nil
}

func fillChangeRecommendations(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) error {
	if len(motion.MotionWork.ChangeRecommendationIDS) == 0 {
		return nil
	}
	for _, id := range motion.MotionWork.ChangeRecommendationIDS {
		data := fetch.Object(
			ctx,
			fmt.Sprintf("motion_change_recommendation/%d", id),
			"id",
			"rejected",
			"type",
			"other_description",
			"line_from",
			"line_to",
			"text",
			"creation_time",
			"internal",
		)
		var internal bool
		if err := json.Unmarshal(data["internal"], &internal); err != nil {
			return fmt.Errorf("decoding internal from ChangeRecommendations: %w", err)
		}
		if internal {
			continue
		}
		bs, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("encoding ChangeRecommendations data: %w", err)
		}
		var mo dbMotionChangeRecommendation
		if err := json.Unmarshal(bs, &mo); err != nil {
			return fmt.Errorf("decoding ChangeRecommendations data: %w", err)
		}
		motion.ChangeRecommendations = append(motion.ChangeRecommendations, mo)
	}
	return nil
}

func fillRecommendationReferencingMotions(ctx context.Context, fetch *datastore.Fetcher, titler projector.Titler, motion *dbMotion) error {
	for _, id := range motion.MotionWork.ReferencedInMotionRecommendationExtensionIDS {
		fqid := fmt.Sprintf("motion/%d", id)
		title, err := titler.GetTitleInformation(ctx, fetch, fqid, "", motion.MotionWork.MeetingID)
		if err != nil {
			return fmt.Errorf("encoding GetTitleInformation data: %w", err)
		}
		motion.RecommendationReferencingMotions = append(motion.RecommendationReferencingMotions, title)
	}
	return nil
}

func fillRecommendationLabelEtc(ctx context.Context, fetch *datastore.Fetcher, titler projector.Titler, motion *dbMotion, meeting *dbMeeting) error {
	data := fetch.Object(
		ctx,
		fmt.Sprintf("motion_state/%d", motion.MotionWork.RecommendationID),
		"recommendation_label",
		"show_recommendation_extension_field",
	)
	st, err := motionStateFromMap(data)
	if err != nil {
		return fmt.Errorf("get motion state: %w", err)
	}
	motion.RecommendationLabel = st.RecommendationLabel
	if st.MotionStateWork.ShowRecommendationExtensionField {
		motion.RecommendationExtension = motion.MotionWork.RecommendationExtension
		motion.RecommendationReferencedMotions = make(map[string]json.RawMessage, len(motion.MotionWork.RecommendationExtensionReferenceIDS))
		for _, fqid := range motion.MotionWork.RecommendationExtensionReferenceIDS {
			parts := strings.Split(fqid, "/")
			collection := parts[0]
			if collection != "motion" {
				return fmt.Errorf("implementation of RecommendationReferencedMotions includes only motion-collection, but not %s", collection)
			}
			title, err := titler.GetTitleInformation(ctx, fetch, fqid, "", motion.MotionWork.MeetingID)
			if err != nil {
				return fmt.Errorf("encoding GetTitleInformation data: %w", err)
			}
			motion.RecommendationReferencedMotions[fqid] = title
		}

	}
	if motion.MotionWork.StatuteParagraphID > 0 {
		motion.Recommender = meeting.MotionsStatuteRecommendationsBy
	} else {
		motion.Recommender = meeting.MotionsRecommendationsBy
	}
	return nil
}

func fillSubmitters(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) error {
	type submitterSort struct {
		MeetingUserID int `json:"meeting_user_id"`
		Weight        int `json:"weight"`
	}
	var submitterToSort []*submitterSort

	for _, id := range motion.MotionWork.SubmitterIDS {
		data := fetch.Object(ctx, fmt.Sprintf("motion_submitter/%d", id), "meeting_user_id", "weight")
		bs, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("encoding MotionSubmitter data: %w", err)
		}

		var su submitterSort
		if err := json.Unmarshal(bs, &su); err != nil {
			return fmt.Errorf("decoding MotionSubmitter data: %w", err)
		}
		submitterToSort = append(submitterToSort, &su)
	}

	sort.Slice(submitterToSort, func(i, j int) bool {
		if submitterToSort[i].Weight == submitterToSort[j].Weight {
			return submitterToSort[i].MeetingUserID < submitterToSort[j].MeetingUserID
		}
		return submitterToSort[i].Weight < submitterToSort[j].Weight
	})

	for _, sortedSub := range submitterToSort {
		var userID int
		fetch.Fetch(ctx, &userID, "meeting_user/%d/user_id", sortedSub.MeetingUserID)
		if err := fetch.Err(); err != nil {
			return fmt.Errorf("getting user for meeting user %d: %w", sortedSub.MeetingUserID, err)
		}

		user, err := NewUser(ctx, fetch, userID, motion.MotionWork.MeetingID)
		if err != nil {
			return fmt.Errorf("getting new user id: %w", err)
		}
		motion.Submitters = append(motion.Submitters, user.UserRepresentation(motion.MotionWork.MeetingID))
	}
	return nil
}

func fillAmendments(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) error {
	fetchFields := []string{
		"id",
		"title",
		"number",
		"meeting_id",
		"amendment_paragraphs",
		"state_id",
		"recommendation_id",
		"change_recommendation_ids",
	}
	for _, id := range motion.MotionWork.AmendmentIDS {
		data := fetch.Object(ctx, fmt.Sprintf("motion/%d", id), fetchFields...)
		motionAmend, err := motionFromMap(data)
		if err != nil {
			return fmt.Errorf("motionFromMap: %w", err)
		}

		if err := fillChangeRecommendations(ctx, fetch, motionAmend); err != nil {
			return fmt.Errorf("fill change recommendations: %w", err)
		}

		var amendment amendmentsType
		amendment.ID = id
		amendment.Title = motionAmend.Title
		amendment.Number = motionAmend.Number
		amendment.AmendmentParagraph = motionAmend.AmendmentParagraph
		amendment.ChangeRecommendations = motionAmend.ChangeRecommendations

		maif := datastore.String(ctx, fetch.FetchIfExist, "motion_state/%d/merge_amendment_into_final", motionAmend.MotionWork.StateID)
		if maif == "do_merge" {
			amendment.MergeAmendmentIntoFinal = maif
			amendment.MergeAmendmentIntoDiff = maif
		} else {
			amendment.MergeAmendmentIntoFinal = "undefined"
			if maif == "do_not_merge" || motionAmend.MotionWork.RecommendationID == 0 {
				amendment.MergeAmendmentIntoDiff = "undefined"
			} else {
				maifReco := datastore.String(ctx, fetch.FetchIfExist, "motion_state/%d/merge_amendment_into_final", motionAmend.MotionWork.RecommendationID)
				if maifReco == "do_merge" {
					amendment.MergeAmendmentIntoDiff = maifReco
				} else {
					amendment.MergeAmendmentIntoDiff = "undefined"
				}
			}
		}

		motion.Amendments = append(motion.Amendments, amendment)
	}
	return nil
}
