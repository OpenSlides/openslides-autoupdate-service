package slide

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbAssignment struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func assignmentFromMap(in map[string]json.RawMessage) (*dbAssignment, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding assignment data: %w", err)
	}

	var a dbAssignment
	if err := json.Unmarshal(bs, &a); err != nil {
		return nil, fmt.Errorf("decoding assignment data: %w", err)
	}
	return &a, nil
}

// Assignment renders the assignment slide.
func Assignment(store *projector.SlideStore) {
	store.RegisterSliderFunc("assignment", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})

	store.RegisterGetTitleInformationFunc("assignment", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, itemNumber string, meetingID int) (json.RawMessage, error) {
		data := fetch.Object(ctx, []string{"id", "title"}, fqid)
		assignment, err := assignmentFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get assignment: %w", err)
		}

		title := struct {
			Collection       string `json:"collection"`
			ContentObjectID  string `json:"content_object_id"`
			Title            string `json:"title"`
			AgendaItemNumber string `json:"agenda_item_number"`
		}{
			"assignment",
			fqid,
			assignment.Title,
			itemNumber,
		}

		bs, err := json.Marshal(title)
		if err != nil {
			return nil, fmt.Errorf("encoding title: %w", err)
		}
		return bs, err
	})
}
