package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
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
	store.RegisterSliderFunc("mediafile", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()
		data := fetch.Object(ctx, []string{"id", "mimetype"}, p7on.ContentObjectID)
		mediafile, err := mediafileItemFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("get mediafile item from map: %w", err)
		}
		responseValue, err := json.Marshal(map[string]interface{}{"id": mediafile.ID, "mimetype": mediafile.Mimetype})
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response slide mediafile item: %w", err)
		}
		return responseValue, fetch.Keys(), err
	})

	store.RegisterGetTitleInformationFunc("mediafile", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, []string{"id", "title"}, fqid)
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
