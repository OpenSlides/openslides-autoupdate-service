package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

type dbMeetingMediafile struct {
	ID          int `json:"id"`
	MediafileID int `json:"mediafile_id"`
}

type dbMediafile struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Mimetype string `json:"mimetype"`
}

func meetingMediafileItemFromMap(in map[string]json.RawMessage) (*dbMeetingMediafile, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding mediafile item data: %w", err)
	}

	var mf dbMeetingMediafile
	if err := json.Unmarshal(bs, &mf); err != nil {
		return nil, fmt.Errorf("decoding mediafile item data: %w", err)
	}
	return &mf, nil
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

// MeetingMediafile renders the mediafile slide.
func MeetingMediafile(store *projector.SlideStore) {
	store.RegisterSliderFunc("meeting_mediafile", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		meetingMediafileData := fetch.Object(ctx, p7on.ContentObjectID, "id", "mediafile_id")
		if err := fetch.Err(); err != nil {
			return nil, err
		}
		meetingMediafile, err := meetingMediafileItemFromMap(meetingMediafileData)
		data := fetch.Object(ctx, "mediafile/"+strconv.Itoa(meetingMediafile.MediafileID), "id", "mimetype")
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

	store.RegisterGetTitleInformationFunc("meeting_mediafile", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		meetingMediafileData := fetch.Object(ctx, fqid, "id", "mediafile_id")
		if err := fetch.Err(); err != nil {
			return nil, err
		}
		meetingMediafile, err := meetingMediafileItemFromMap(meetingMediafileData)
		mediafileFqid := "mediafile/" + strconv.Itoa(meetingMediafile.MediafileID)
		data := fetch.Object(ctx, mediafileFqid, "id", "mediafile_id")
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
			mediafileFqid,
			mediafile.Title,
		}

		bs, err := json.Marshal(mediafiletitle)
		if err != nil {
			return nil, fmt.Errorf("decoding title: %w", err)
		}
		return bs, err
	})
}
