package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbProjectorCountdown struct {
	ID            int     `json:"id"`
	Description   string  `json:"description"`
	Running       bool    `json:"running"`
	CountdownTime float32 `json:"countdown_time"`
	MeetingID     int     `json:"meeting_id"`
}

func projectorCountdownFromMap(in map[string]json.RawMessage) (*dbProjectorCountdown, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding projector countdown data: %w", err)
	}

	var pc dbProjectorCountdown
	if err := json.Unmarshal(bs, &pc); err != nil {
		return nil, fmt.Errorf("decoding projector countdown data: %w", err)
	}
	return &pc, nil
}

type dbProjectorMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func projectorMessageFromMap(in map[string]json.RawMessage) (*dbProjectorMessage, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding projector message data: %w", err)
	}

	var pm dbProjectorMessage
	if err := json.Unmarshal(bs, &pm); err != nil {
		return nil, fmt.Errorf("decoding projector message data: %w", err)
	}
	return &pm, nil
}

// ProjectorCountdown renders the projector_countdown slide.
func ProjectorCountdown(store *projector.SlideStore) {
	store.RegisterSliderFunc("projector_countdown", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		data := fetch.Object(ctx, p7on.ContentObjectID, "id", "description", "running", "countdown_time", "meeting_id")
		pc, err := projectorCountdownFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get projector countdown from map: %w", err)
		}
		pcwarningTime := datastore.Int(ctx, fetch.FetchIfExist, "meeting/%d/projector_countdown_warning_time", pc.MeetingID)
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		responseValue, err := json.Marshal(map[string]interface{}{"description": pc.Description, "running": pc.Running, "countdown_time": pc.CountdownTime, "warning_time": pcwarningTime})
		if err != nil {
			return nil, fmt.Errorf("encoding response for projector countdown slide: %w", err)
		}
		return responseValue, err
	})
}

// ProjectorMessage renders the projector_message slide.
func ProjectorMessage(store *projector.SlideStore) {
	store.RegisterSliderFunc("projector_message", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		data := fetch.Object(ctx, p7on.ContentObjectID, "id", "message")
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		projectorMessage, err := projectorMessageFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get projector message from map: %w", err)
		}
		responseValue, err := json.Marshal(map[string]interface{}{"message": projectorMessage.Message})
		if err != nil {
			return nil, fmt.Errorf("encoding response for projector message slide: %w", err)
		}
		return responseValue, err
	})
}
