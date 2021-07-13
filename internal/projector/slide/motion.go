package slide

import (
	"context"
	"encoding/json"
	"fmt"

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
	ID                      int    `json:"id"`
	Title                   string `json:"title"`
	Number                  string `json:"number"`
	amendmentParagraph      string `json:"amendment_paragraph_$"`
	AmendmentParagraphs     map[string]string
	MergeAmendmentIntoFinal string
	MergeAmendmentIntoDiff  string
}
type leadMotionType struct {
	Title  string  `json:"title"`
	Number string  `json:"number"`
	Text   *string `json:"text"`
}
type dbMotionWork struct {
	MeetingID                                    int      `json:"meeting_id"`
	LeadMotionID                                 int      `json:"lead_motion_id"`
	Statute_ParagraphID                          int      `json:"statute_paragraph_id"`
	AmendmentParagraph                           []string `json:"amendment_paragraph_$"`
	ChangeRecommendationIDS                      []int    `json:"change_recommendation_ids"`
	AmendmentIDS                                 []int    `json:"amendment_ids"`
	SubmitterIDS                                 []int    `json:"submitter_ids"`
	ReferencedInMotionRecommendationExtensionIDS []int    `json:"referenced_in_motion_recommendation_extension_ids"`
	RecommendationID                             int      `json:"recommendation_id"`
	RecommendationExtensionReferenceIDS          []int    `json:"recommendation_extension_reference_ids"`
}
type dbMotion struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Number string `json:"number"`

	// zu berechnen
	AmendmentParagraphs              *map[string]string            `json:"amendment_paragraphs"`
	LeadMotion                       *leadMotionType               `json:"lead_motion"`
	BaseStatute                      *dbMotionStatuteParagraph     `json:"base_statute"`
	ChangeRecommendations            *dbMotionChangeRecommendation `json:"change_recommendations"`
	Amendments                       []*amendmentsType             `json:"amendments"`
	Submitters                       map[string]*string            `json:"submitters"`                         // map submitter to UserRepresentation
	RecommendationReferencingMotions map[string]string             `json:"recommendation_referencing_motions"` // map motions from ReferencedInMotionRecommendationExtensionIDS
	RecommendationLabel              *string                       `json:"recommendation_label"`
	RecommendationExtension          *string                       `json:"recommendation_extension"` // LESEN: NICHT hier importieren, aber exportieren
	RecommendationReferencedMotions  map[string]string             `json:"recommendation_referenced_motions"`
	Recommender                      *string                       `json:"recommender"`

	// Abhaengige Felder
	Text                 *string `json:"text"`
	Reason               *string `json:"reason"`
	ModifiedFinalVersion *string `json:"modified_final_version"`

	MotionWork *dbMotionWork
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
		}
		// state := fetch.String(ctx, "%s/%s", p7on.ContentObjectID, "state")
		// if state == "published" {
		// 	fetchFields = append(fetchFields, []string{
		// 		"is_pseudoanonymized",
		// 		"pollmethod",
		// 		"onehundred_percent_base",
		// 		"majority_method",
		// 		"votesvalid",
		// 		"votesinvalid",
		// 		"votescast",
		// 		"global_option_id",
		// 	}...)
		// }
		data := fetch.Object(ctx, fetchFields, p7on.ContentObjectID)

		motion, err := motionFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("get motion: %w", err)
		}
		_ = motion
		return nil, nil, nil
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
