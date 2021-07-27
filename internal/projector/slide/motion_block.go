package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbMotionBlock struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	AgendaItemID int    `json:"agenda_item_id"`
	MotionIDS    []int  `json:"motion_ids"`
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

type motionRepr struct {
	Title                   string         `json:"title"`
	Number                  string         `json:"number"`
	AgendaItemNumber        string         `json:"agenda_item_number"`
	Recommendation          *dbMotionState `json:"recommendation,omitempty"`
	RecommendationExtension *string        `json:"recommendation_extension,omitempty"`
}

// MotionBlock renders the motion_block slide.
func MotionBlock(store *projector.SlideStore) {
	store.RegisterSliderFunc("motion_block", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		titlerMotion := store.GetTitleInformationFunc("motion")
		if titlerMotion == nil {
			return nil, fmt.Errorf("no titler function registered for motion")
		}

		data := fetch.Object(ctx, p7on.ContentObjectID, "title", "motion_ids")
		motionBlock, err := motionBlockFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motionBlock: %w", err)
		}
		var motions []motionRepr
		referenced := map[string]json.RawMessage{}
		for _, motionID := range motionBlock.MotionIDS {
			data := fetch.Object(
				ctx,
				fmt.Sprintf("motion/%d", motionID),
				"title",
				"number",
				"agenda_item_id",
				"recommendation_id",
				"recommendation_extension",
				"recommendation_extension_reference_ids",
				"meeting_id",
			)
			motion, err := motionFromMap(data)
			if err != nil {
				return nil, fmt.Errorf("get motion: %w", err)
			}

			var recommendation *dbMotionState
			var recommendationExtension *string
			if motion.MotionWork.RecommendationID > 0 {
				data := fetch.Object(
					ctx,
					fmt.Sprintf("motion_state/%d", motion.MotionWork.RecommendationID),
					"recommendation_label",
					"css_class",
					"show_recommendation_extension_field",
				)
				recommendation, err = motionStateFromMap(data)
				if err != nil {
					return nil, fmt.Errorf("get motion: %w", err)
				}
				if recommendation.MotionStateWork.ShowRecommendationExtensionField {
					recommendationExtension = &motion.RecommendationExtension
				}
				for _, referenceObjectID := range motion.MotionWork.RecommendationExtensionReferenceIDS {
					title, err := titlerMotion.GetTitleInformation(ctx, fetch, referenceObjectID, "", motion.MotionWork.MeetingID)
					if err != nil {
						return nil, fmt.Errorf("encoding GetTitleInformation data: %w", err)
					}
					referenced[referenceObjectID] = title
				}
				recommendation.MotionStateWork = nil // don't export
			}
			itemNumber := datastore.String(ctx, fetch.Fetch, "agenda_item/%d/item_number", motion.MotionWork.AgendaItemID)
			if err := fetch.Err(); err != nil {
				return nil, err
			}

			motions = append(motions, motionRepr{
				Title:                   motion.Title,
				Number:                  motion.Number,
				AgendaItemNumber:        itemNumber,
				Recommendation:          recommendation,
				RecommendationExtension: recommendationExtension,
			})
		}

		out := struct {
			Title      string                     `json:"title"`
			Motions    []motionRepr               `json:"motions"`
			Referenced map[string]json.RawMessage `json:"referenced"`
		}{
			motionBlock.Title,
			motions,
			referenced,
		}
		bs, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding motion_block: %w", err)
		}
		return bs, nil
	})

	store.RegisterGetTitleInformationFunc("motion_block", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, fqid, "id", "title", "agenda_item_id")
		motionBlock, err := motionBlockFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motion block: %w", err)
		}

		if itemNumber == "" && motionBlock.AgendaItemID > 0 {
			itemNumber = datastore.String(ctx, fetch.Fetch, "agenda_item/%d/item_number", motionBlock.AgendaItemID)
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		title := struct {
			Collection       string `json:"collection"`
			ContentObjectID  string `json:"content_object_id"`
			Title            string `json:"title"`
			AgendaItemNumber string `json:"agenda_item_number"`
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
