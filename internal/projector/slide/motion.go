package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbMotion struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Number string `json:"number"`
}

func motionFromMap(in map[string]json.RawMessage) (*dbMotion, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding motion data: %w", err)
	}

	var m dbMotion
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, fmt.Errorf("decoding motion data: %w", err)
	}
	return &m, nil
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
	store.RegisterSlideFunc("motion", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
	store.RegisterTitleFunc("motion", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (*projector.TitlerFuncResult, error) {
		data := fetch.Object(ctx, []string{"id", "number", "title"}, fqid)
		motion, err := motionFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motion from map: %w", err)
		}
		titleData := projector.TitlerFuncResult{
			Title:  &motion.Title,
			Number: &motion.Number,
		}
		return &titleData, err
	})
}

// MotionBlock renders the motion_block slide.
func MotionBlock(store *projector.SlideStore) {
	store.RegisterSlideFunc("motion_block", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
	store.RegisterTitleFunc("motion_block", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (*projector.TitlerFuncResult, error) {
		data := fetch.Object(ctx, []string{"id", "title"}, fqid)
		motionBlock, err := motionBlockFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get motion block from map: %w", err)
		}
		agenda_item_number := value["agenda_item_number"].(string)
		titleData := projector.TitlerFuncResult{
			Title:            &motionBlock.Title,
			AgendaItemNumber: &agenda_item_number,
		}
		return &titleData, err
	})
}
