package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbMotionStatuteParagraph struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type dbMotionChangeRecommendation struct {
	Rejected         bool   `json:"rejected"`
	Type             string `json:"type"`
	OtherDescription string `json:"other_description"`
	LineFrom         int    `json:"line_from"`
	LineTo           int    `json:"line_to"`
	Text             string `json:"text"`
	CreationTime     int    `json:"creation_time"`
}

type amendmentsType struct {
	ID                      int               `json:"id"`
	Title                   string            `json:"title"`
	Number                  string            `json:"number"`
	AmendmentParagraphs     map[string]string `json:"amendment_paragraphs"`
	MergeAmendmentIntoFinal string            `json:"merge_amendment_into_final"`
	MergeAmendmentIntoDiff  string            `json:"merge_amendment_into_diff"`
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
	AmendmentParagraph                           []string `json:"amendment_paragraph_$"`
	ChangeRecommendationIDS                      []int    `json:"change_recommendation_ids"`
	AmendmentIDS                                 []int    `json:"amendment_ids"`
	SubmitterIDS                                 []int    `json:"submitter_ids"`
	ReferencedInMotionRecommendationExtensionIDS []int    `json:"referenced_in_motion_recommendation_extension_ids"`
	RecommendationID                             int      `json:"recommendation_id"`
	RecommendationExtensionReferenceIDS          []string `json:"recommendation_extension_reference_ids"`
	RecommendationExtension                      string   `json:"recommendation_extension"`
	StateID                                      int      `json:"state_id"`
}
type dbMotion struct {
	ID                               int                                     `json:"id"`
	Title                            string                                  `json:"title"`
	Number                           string                                  `json:"number"`
	Submitters                       []map[string]string                     `json:"submitters"`
	ShowSidebox                      bool                                    `json:"show_sidebox"`
	LineLength                       int                                     `json:"line_length"`
	Preamble                         string                                  `json:"preamble"`
	LineNumbering                    string                                  `json:"line_numbering"`
	AmendmentParagraphs              map[string]string                       `json:"amendment_paragraphs,omitempty"`
	LeadMotion                       *leadMotionType                         `json:"lead_motion,omitempty"`
	BaseStatute                      *dbMotionStatuteParagraph               `json:"base_statute,omitempty"`
	ChangeRecommendations            map[string]dbMotionChangeRecommendation `json:"change_recommendations,omitempty"`
	Amendments                       map[string]amendmentsType               `json:"amendments,omitempty"`
	RecommendationReferencingMotions map[string]json.RawMessage              `json:"recommendation_referencing_motions,omitempty"`
	RecommendationLabel              string                                  `json:"recommendation_label,omitempty"`
	RecommendationExtension          string                                  `json:"recommendation_extension,omitempty"`
	RecommendationReferencedMotions  map[string]json.RawMessage              `json:"recommendation_referenced_motions,omitempty"`
	Recommender                      string                                  `json:"recommender,omitempty"`
	Text                             string                                  `json:"text,omitempty"`
	Reason                           string                                  `json:"reason,omitempty"`
	ModifiedFinalVersion             string                                  `json:"modified_final_version,omitempty"`
	MotionWork                       *dbMotionWork                           `json:",omitempty"`
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

type dbMotionBlock struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func motionBlockFromMap(in map[string]json.RawMessage) (*dbMotionBlock, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding motion data: %w", err)
	}

	var m dbMotionBlock
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, fmt.Errorf("decoding motion: %w", err)
	}
	return &m, nil
}

// Motion renders the motion slide.
func Motion(store *projector.SlideStore) {
	store.RegisterSliderFunc("motion", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

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
			return nil, nil, fmt.Errorf("getMeeting: %w", err)
		}

		var options struct {
			Mode string `json:"mode"`
		}
		if p7on.Options != nil {
			if err := json.Unmarshal(p7on.Options, &options); err != nil {
				return nil, nil, fmt.Errorf("decoding projection options: %w", err)
			}
		}

		fetchFields := []string{
			"id",
			"title",
			"number",
			"meeting_id",
			"lead_motion_id",
			"statute_paragraph_id",
			"amendment_paragraph_$",
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

		data := fetch.Object(ctx, fetchFields, p7on.ContentObjectID)

		motion, err := motionFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("get motion: %w", err)
		}
		motion.RecommendationExtension = "" // will be (re-)filled conditionally

		fillMotionFromMeeting(motion, meeting)
		fillAmendmentParagraphs(ctx, fetch, motion)
		err = fillSubmitters(ctx, fetch, motion)
		if err != nil {
			return nil, nil, fmt.Errorf("fillSubmitters: %w", err)
		}
		err = fillLeadMotion(ctx, fetch, motion)
		if err != nil {
			return nil, nil, fmt.Errorf("fillLeadMotion: %w", err)
		}
		err = fillBaseStatute(ctx, fetch, motion)
		if err != nil {
			return nil, nil, fmt.Errorf("fillBaseStatute: %w", err)
		}
		err = fillChangeRecommendations(ctx, fetch, motion)
		if err != nil {
			return nil, nil, fmt.Errorf("fillChangeRecommendations: %w", err)
		}
		err = fillAmendments(ctx, fetch, motion)
		if err != nil {
			return nil, nil, fmt.Errorf("fillAmendments: %w", err)
		}

		titlerMotion := store.GetTitleInformationFunc("motion")
		if titlerMotion == nil {
			return nil, nil, fmt.Errorf("no titler function registered for motion")
		}

		if meeting.MotionsShowReferringMotions && len(motion.MotionWork.ReferencedInMotionRecommendationExtensionIDS) > 0 {
			err = fillRecommendationReferencingMotions(ctx, fetch, titlerMotion, motion)
			if err != nil {
				return nil, nil, fmt.Errorf("FillRecommendationReferencingMotions: %w", err)
			}
		}

		if meeting.MotionsEnableRecommendationOnProjector && motion.MotionWork.RecommendationID > 0 {
			err = fillRecommendationLabelEtc(ctx, fetch, titlerMotion, motion, meeting)
			if err != nil {
				return nil, nil, fmt.Errorf("RecommendationLabelEtc: %w", err)
			}
		}

		motion.MotionWork = nil // do not export worker fields
		responseValue, err := json.Marshal(motion)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response for slide motion: %w", err)
		}
		return responseValue, fetch.Keys(), err
	})

	store.RegisterGetTitleInformationFunc("motion", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, []string{"id", "number", "title"}, fqid)
		motion, err := motionFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motion: %w", err)
		}

		title := struct {
			Collection      string `json:"collection"`
			ContentObjectID string `json:"content_object_id"`
			Title           string `json:"title"`
			Number          string `json:"number"`
			AgendaNumber    string `json:"agenda_item_number"`
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

// MotionBlock renders the motion_block slide.
func MotionBlock(store *projector.SlideStore) {
	store.RegisterSliderFunc("motion_block", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})

	store.RegisterGetTitleInformationFunc("motion_block", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, []string{"id", "title"}, fqid)
		motionBlock, err := motionBlockFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motion block: %w", err)
		}

		title := struct {
			Collection      string `json:"collection"`
			ContentObjectID string `json:"content_object_id"`
			Title           string `json:"title"`
			AgendaNumber    string `json:"agenda_item_number"`
		}{
			"motion_block",
			fqid,
			motionBlock.Title,
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

func fillAmendmentParagraphs(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) {
	if len(motion.MotionWork.AmendmentParagraph) > 0 {
		motion.AmendmentParagraphs = make(map[string]string, len(motion.MotionWork.AmendmentParagraph))
		for _, nr := range motion.MotionWork.AmendmentParagraph {
			text := fetch.String(ctx, "motion/%d/amendment_paragraph_$%s", motion.ID, nr)
			motion.AmendmentParagraphs[nr] = text
		}
	}
}

func fillLeadMotion(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) error {
	if motion.MotionWork.LeadMotionID == 0 {
		return nil
	}
	data := fetch.Object(ctx, []string{"title", "number", "text"}, "motion/%d", motion.MotionWork.LeadMotionID)
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
	data := fetch.Object(ctx, []string{"title", "text"}, "motion_statute_paragraph/%d", motion.MotionWork.StatuteParagraphID)
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
	motion.ChangeRecommendations = make(map[string]dbMotionChangeRecommendation)
	for _, id := range motion.MotionWork.ChangeRecommendationIDS {
		data := fetch.Object(ctx, []string{"rejected", "type", "other_description", "line_from", "line_to", "text", "creation_time", "internal"}, "motion_change_recommendation/%d", id)
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
		motion.ChangeRecommendations[strconv.Itoa(id)] = mo
	}
	return nil
}

func fillRecommendationReferencingMotions(ctx context.Context, fetch *datastore.Fetcher, titler projector.Titler, motion *dbMotion) error {
	motion.RecommendationReferencingMotions = make(map[string]json.RawMessage, len(motion.MotionWork.ReferencedInMotionRecommendationExtensionIDS))
	for _, id := range motion.MotionWork.ReferencedInMotionRecommendationExtensionIDS {
		fqid := fmt.Sprintf("motion/%d", id)
		title, err := titler.GetTitleInformation(ctx, fetch, fqid, "", motion.MotionWork.MeetingID)
		if err != nil {
			return fmt.Errorf("encoding GetTitleInformation data: %w", err)
		}
		motion.RecommendationReferencingMotions[strconv.Itoa(id)] = title
	}
	return nil
}

func fillRecommendationLabelEtc(ctx context.Context, fetch *datastore.Fetcher, titler projector.Titler, motion *dbMotion, meeting *dbMeeting) error {
	data := fetch.Object(ctx, []string{"recommendation_label", "show_recommendation_extension_field"}, "motion_state/%d", motion.MotionWork.RecommendationID)
	bs, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encoding MotionState data: %w", err)
	}
	var st struct {
		RecommendationLabel              string `json:"recommendation_label"`
		ShowRecommendationExtensionField bool   `json:"show_recommendation_extension_field"`
	}
	if err := json.Unmarshal(bs, &st); err != nil {
		return fmt.Errorf("decoding MotionState data: %w", err)
	}
	motion.RecommendationLabel = st.RecommendationLabel
	if st.ShowRecommendationExtensionField {
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
		UserID int `json:"user_id"`
		Weight int `json:"weight"`
	}
	var submitterToSort []*submitterSort

	for _, id := range motion.MotionWork.SubmitterIDS {
		data := fetch.Object(ctx, []string{"user_id", "weight"}, "motion_submitter/%d", id)
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
			return submitterToSort[i].UserID < submitterToSort[j].UserID
		}
		return submitterToSort[i].Weight < submitterToSort[j].Weight
	})

	for _, sortedSub := range submitterToSort {
		user, err := NewUser(ctx, fetch, sortedSub.UserID, motion.MotionWork.MeetingID)
		if err != nil {
			return fmt.Errorf("getting new user id: %w", err)
		}
		fqid := fmt.Sprintf("user/%d", sortedSub.UserID)
		submitter := map[string]string{
			fqid: user.UserRepresentation(motion.MotionWork.MeetingID),
		}
		motion.Submitters = append(motion.Submitters, submitter)
	}
	return nil
}

func fillAmendments(ctx context.Context, fetch *datastore.Fetcher, motion *dbMotion) error {
	motion.Amendments = make(map[string]amendmentsType, len(motion.MotionWork.AmendmentIDS))
	fetchFields := []string{
		"id",
		"title",
		"number",
		"meeting_id",
		"amendment_paragraph_$",
		"state_id",
		"recommendation_id",
	}
	for _, id := range motion.MotionWork.AmendmentIDS {
		data := fetch.Object(ctx, fetchFields, "motion/%d", id)
		motionAmend, err := motionFromMap(data)
		if err != nil {
			return fmt.Errorf("motionFromMap: %w", err)
		}
		fillAmendmentParagraphs(ctx, fetch, motionAmend)
		var amendment amendmentsType
		amendment.ID = id
		amendment.Title = motionAmend.Title
		amendment.Number = motionAmend.Number
		amendment.AmendmentParagraphs = motionAmend.AmendmentParagraphs

		maif := fetch.String(ctx, "motion_state/%d/merge_amendment_into_final", motionAmend.MotionWork.StateID)
		if maif == "do_merge" {
			amendment.MergeAmendmentIntoFinal = maif
			amendment.MergeAmendmentIntoDiff = maif
		} else {
			amendment.MergeAmendmentIntoFinal = "undefined"
			if maif == "do_not_merge" || motionAmend.MotionWork.RecommendationID == 0 {
				amendment.MergeAmendmentIntoDiff = "undefined"
			} else {
				maifReco := fetch.String(ctx, "motion_state/%d/merge_amendment_into_final", motionAmend.MotionWork.RecommendationID)
				if maifReco == "do_merge" {
					amendment.MergeAmendmentIntoDiff = maifReco
				} else {
					amendment.MergeAmendmentIntoDiff = "undefined"
				}
			}
		}

		motion.Amendments[strconv.Itoa(id)] = amendment
	}
	return nil
}
