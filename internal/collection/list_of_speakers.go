package collection

import (
	"context"
	"encoding/json"

	"github.com/OpenSlides/openslides-permission-service/internal/dataprovider"
	"github.com/OpenSlides/openslides-permission-service/internal/perm"
)

// ListOfSpeaker connects the handers for list of speakers.
func ListOfSpeaker(dp dataprovider.DataProvider) perm.ConnecterFunc {
	l := &listOfSpeaker{dp}
	return func(s perm.HandlerStore) {
		s.RegisterWriteHandler("list_of_speakers.delete", perm.WriteCheckerFunc(l.delete))
	}
}

type listOfSpeaker struct {
	dp dataprovider.DataProvider
}

func (l *listOfSpeaker) delete(ctx context.Context, userID int, payload map[string]json.RawMessage) (map[string]interface{}, error) {
	return nil, perm.NotAllowedf("list_of_speaker.delete is an internal action.")
}
