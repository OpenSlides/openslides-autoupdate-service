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
	Mimetype string `json:"mimetype"`
}

func mediafileItemFromMap(in map[string]json.RawMessage) (*dbMediafile, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding mediafile data")
	}

	var mf dbMediafile
	if err := json.Unmarshal(bs, &mf); err != nil {
		return nil, fmt.Errorf("decoding mediafile item: %w", err)
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
			return nil, nil, err
		}
		responseValue, err := json.Marshal(map[string]interface{}{"id": mediafile.ID, "mimetype": mediafile.Mimetype})
		if err != nil {
			return nil, nil, err
		}
		return responseValue, fetch.Keys(), err
	})

	store.RegisterGetTitleInformationFunc("mediafile", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string) (json.RawMessage, error) {
		title := "title of mediafile (TODO)"

		agendatitle := struct {
			Title string `json:"title"`
		}{
			title,
		}

		bs, err := json.Marshal(agendatitle)
		if err != nil {
			return nil, fmt.Errorf("decoding title: %w", err)
		}
		return bs, err
	})
}
