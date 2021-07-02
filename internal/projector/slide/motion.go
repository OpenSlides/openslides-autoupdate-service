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
	store.RegisterSliderFunc("motion", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
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
