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

	var m dbAssignment
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, fmt.Errorf("decoding assignment data: %w", err)
	}
	return &m, nil
}

// Assignment renders the assignment slide.
func Assignment(store *projector.SlideStore) {
	store.RegisterSlideFunc("assignment", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
	store.RegisterTitleFunc("assignment", func(ctx context.Context, fetch *datastore.Fetcher, fqid string, meeting_id int, value map[string]interface{}) (*projector.TitlerFuncResult, error) {

		data := fetch.Object(ctx, []string{"id", "title"}, fqid)
		assignment, err := assignmentFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get assignment from map: %w", err)
		}
		agenda_item_number := value["agenda_item_number"].(string)
		titleData := projector.TitlerFuncResult{
			Title:            &assignment.Title,
			AgendaItemNumber: &agenda_item_number,
		}
		return &titleData, err
	})
}
