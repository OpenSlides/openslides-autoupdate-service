package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/openslides/openslides-autoupdate-service/internal/datastore"
	"github.com/openslides/openslides-autoupdate-service/internal/projector"
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
		var los dbListOfSpeakers
		keys, err := datastore.GetObject(ctx, ds, p7on.ContentObjectID, &los)
		if err != nil {
			return nil, nil, fmt.Errorf("fetch list of speakers: %w", err)
		}
		hotkeys = append(hotkeys, keys...)

		var title string
		if err := datastore.Get(ctx, ds, los.ContentObjectID+"/title", &title); err != nil {
			return nil, nil, fmt.Errorf("fetch title from content object id: %w", err)
		}
		hotkeys = append(hotkeys, los.ContentObjectID+"/title")

		var speakersWaiting []outputSpeaker
		var speakersFinished []outputSpeaker
		var currentSpeaker *outputSpeaker
		for _, id := range los.SpeakerIDs {
			var speaker dbSpeaker
			keys, err := datastore.GetObject(ctx, ds, fmt.Sprintf("speaker/%d", id), &speaker)
			if err != nil {
				return nil, nil, fmt.Errorf("fetch speaker object: %w", err)
			}
			hotkeys = append(hotkeys, keys...)

			var user dbUser
			keys, err = datastore.GetObject(ctx, ds, fmt.Sprintf("user/%d", speaker.UserID), &user)
			if err != nil {
				return nil, nil, fmt.Errorf("fetch user for speaker %d: %w", id, err)
			}
			hotkeys = append(hotkeys, keys...)

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
		return b, hotkeys, nil
	})
}

// CurrentListOfSpeakers renders the current_list_of_speakers slide.
func CurrentListOfSpeakers(store *projector.SlideStore) {
	store.AddFunc("current_list_of_speakers", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}

// CurrentSpeakerChyron renders the current_speaker_chyron slide.
func CurrentSpeakerChyron(store *projector.SlideStore) {
	store.AddFunc("current_speaker_chyron", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		return []byte(`"TODO"`), nil, nil
	})
}
