package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

type dbMediafile struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Mimetype string `json:"mimetype"`
}

func mediafileItemFromMap(in map[string]json.RawMessage) (*dbMediafile, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding mediafile item data: %w", err)
	}

	var mf dbMediafile
	if err := json.Unmarshal(bs, &mf); err != nil {
		return nil, fmt.Errorf("decoding mediafile item data: %w", err)
	}
	return &mf, nil
}

// Mediafile renders the mediafile slide.
func Mediafile(store *projector.SlideStore) {
	store.RegisterSliderFunc("mediafile", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {

		data := fetch.Object(ctx, p7on.ContentObjectID, "id", "mimetype")
		if err := fetch.Err(); err != nil {
			return nil, err
		}
		mediafile, err := mediafileItemFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get mediafile item from map: %w", err)
		}
		responseValue, err := json.Marshal(map[string]interface{}{"id": mediafile.ID, "mimetype": mediafile.Mimetype})
		if err != nil {
			return nil, fmt.Errorf("encoding response slide mediafile item: %w", err)
		}
		return responseValue, err
	})

	store.RegisterGetTitleInformationFunc("mediafile", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, fqid, "id", "title")
		mediafile, err := mediafileItemFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get mediafile: %w", err)
		}

		mediafiletitle := struct {
			Collection      string `json:"collection"`
			ContentObjectID string `json:"content_object_id"`
			Title           string `json:"title"`
		}{
			"mediafile",
			fqid,
			mediafile.Title,
		}

		bs, err := json.Marshal(mediafiletitle)
		if err != nil {
			return nil, fmt.Errorf("decoding title: %w", err)
		}
		return bs, err
	})
}
