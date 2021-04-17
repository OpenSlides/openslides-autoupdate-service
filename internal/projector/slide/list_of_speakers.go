package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbListOfSpeakers struct {
	SpeakerIDs      []int  `json:"speaker_ids"`
	ContentObjectID string `json:"content_object_id"`
	Closed          bool   `json:"closed"`
}

type outputSpeaker struct {
	User         string `json:"user"`
	Marked       bool   `json:"marked"`
	PointOfOrder bool   `json:"point_of_order"`
	Weight       int    `json:"weight"`
	EndTime      int    `json:"end_time,omitempty"`
}

type dbSpeaker struct {
	UserID       int  `json:"user_id"`
	Marked       bool `json:"marked"`
	PointOfOrder bool `json:"point_of_order"`
	Weight       int  `json:"weight"`
	BeginTime    int  `json:"begin_time"`
	EndTime      int  `json:"end_time"`
}

// ListOfSpeaker renders current list of speaker slide.
func ListOfSpeaker(store *projector.SlideStore) {
	store.AddFunc("list_of_speakers", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, hotkeys []string, err error) {
		return renderListOfSpeakers(ctx, ds, p7on.ContentObjectID)
	})
}

func renderListOfSpeakers(ctx context.Context, ds projector.Datastore, losFQID string) (encoded []byte, keys []string, err error) {
	fetch := datastore.NewFetcher(ds)
	defer func() {
		if err == nil {
			err = fetch.Error()
		}
	}()

	var los dbListOfSpeakers
	fetch.Object(ctx, &los, losFQID)
	title := fetch.String(ctx, los.ContentObjectID+"/title")

	var speakersWaiting []outputSpeaker
	var speakersFinished []outputSpeaker
	var currentSpeaker *outputSpeaker
	for _, id := range los.SpeakerIDs {
		var speaker dbSpeaker
		fetch.Object(ctx, &speaker, "speaker/%d", id)

		var user dbUser
		fetch.Object(ctx, &user, "user/%d", speaker.UserID)

		s := outputSpeaker{
			User:         user.String(),
			Marked:       speaker.Marked,
			PointOfOrder: speaker.PointOfOrder,
			Weight:       speaker.Weight,
			EndTime:      speaker.EndTime,
		}

		if speaker.BeginTime == 0 && speaker.EndTime == 0 {
			speakersWaiting = append(speakersWaiting, s)
			continue
		}

		if speaker.EndTime == 0 {
			currentSpeaker = &s
			continue
		}

		speakersFinished = append(speakersFinished, s)
	}

	if err := fetch.Error(); err != nil {
		return nil, nil, err
	}

	idx := strings.Index(los.ContentObjectID, "/")
	collection := los.ContentObjectID[:idx]

	slideData := struct {
		Title                   string          `json:"title"`
		Waiting                 []outputSpeaker `json:"waiting"`
		Current                 *outputSpeaker  `json:"current,"`
		Finished                []outputSpeaker `json:"finished"`
		ContentObjectCollection string          `json:"content_object_collection"`
		TitleInformation        string          `json:"title_information"`
		Closed                  bool            `json:"closed"`
	}{
		title,
		speakersWaiting,
		currentSpeaker,
		speakersFinished,
		collection,
		fmt.Sprintf("title_information for %s", los.ContentObjectID),
		los.Closed,
	}
	b, err := json.Marshal(slideData)
	if err != nil {
		return nil, nil, fmt.Errorf("encoding outgoing data: %w", err)
	}
	return b, fetch.Keys(), nil
}

// CurrentListOfSpeakers renders the current_list_of_speakers slide.
func CurrentListOfSpeakers(store *projector.SlideStore) {
	store.AddFunc("current_list_of_speakers", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

		projectorID := fetch.Int(ctx, "projection/%d/current_projector_id", p7on.ID)
		meetingID := fetch.Int(ctx, "projector/%d/meeting_id", projectorID)
		referenceProjectorID := fetch.Int(ctx, "meeting/%d/reference_projector_id", meetingID)
		referenceP7onIDs := fetch.Ints(ctx, "projector/%d/current_projection_ids", referenceProjectorID)

		var losID int
		for _, pID := range referenceP7onIDs {
			contentObjectID := fetch.String(ctx, "projection/%d/content_object_id", pID)
			losID = fetch.Int(ctx, "%s/list_of_speakers_id", contentObjectID)

			if losID != 0 {
				break
			}
		}
		if losID == 0 {
			return []byte("{}"), fetch.Keys(), nil
		}

		if err := fetch.Error(); err != nil {
			return nil, nil, err
		}

		content, keys, err := renderListOfSpeakers(ctx, ds, fmt.Sprintf("list_of_speakers/%d", losID))
		if err != nil {
			return nil, nil, fmt.Errorf("render list of speakers %d: %w", losID, err)
		}
		keys = append(keys, fetch.Keys()...)
		return content, keys, nil
	})
}

// CurrentSpeakerChyron renders the current_speaker_chyron slide.
func CurrentSpeakerChyron(store *projector.SlideStore) {
	store.AddFunc("current_speaker_chyron", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
