package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OpenSlides/openslides-permission-service/internal/collection"
	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
)

// PersonalNote handels permissions for personal notes.
type PersonalNote struct {
	dp dataprovider.DataProvider
}

// NewPersonalNote initializes a personal note.
func NewPersonalNote(dp dataprovider.DataProvider) *PersonalNote {
	return &PersonalNote{
		dp: dp,
	}
}

// Connect creates the routes.
func (p *PersonalNote) Connect(s collection.HandlerStore) {
	s.RegisterWriteHandler("personal_note.update", collection.WriteCheckerFunc(p.modify))
	s.RegisterWriteHandler("personal_note.delete", collection.WriteCheckerFunc(p.modify))
}

func (p PersonalNote) modify(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	fqfield := fmt.Sprintf("personal_note/%s/user_id", payload["id"])
	var noteUserID int
	if err := p.dp.Get(ctx, fqfield, &noteUserID); err != nil {
		return nil, fmt.Errorf("getting %s from datastore: %w", fqfield, err)
	}

	if noteUserID != userID {
		return nil, collection.NotAllowedf("Not your note")
	}
	return nil, nil
}
