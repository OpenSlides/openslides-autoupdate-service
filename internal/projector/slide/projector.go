package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbProjectorMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func projectorMessageFromMap(in map[string]json.RawMessage) (*dbProjectorMessage, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding projector message data")
	}

	var pm dbProjectorMessage
	if err := json.Unmarshal(bs, &pm); err != nil {
		return nil, fmt.Errorf("decoding projector message: %w", err)
	}
	return &pm, nil
}

// ProjectorCountdown renders the projector_countdown slide.
func ProjectorCountdown(store *projector.SlideStore) {
	store.RegisterSliderFunc("projector_countdown", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}

// ProjectorMessage renders the projector_message slide.
func ProjectorMessage(store *projector.SlideStore) {
	store.RegisterSliderFunc("projector_message", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()
		data := fetch.Object(ctx, []string{"id", "message"}, p7on.ContentObjectID)
		projectorMessage, err := projectorMessageFromMap(data)
		if err != nil {
			return nil, nil, err
		}
		responseValue, err := json.Marshal(map[string]interface{}{"message": projectorMessage.Message})
		if err != nil {
			return nil, nil, err
		}
		return responseValue, fetch.Keys(), err
	})
}
