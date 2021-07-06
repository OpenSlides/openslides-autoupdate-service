package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type dbListOfSpeakers struct {
	SpeakerIDs      []int  `json:"speaker_ids"`
	ContentObjectID string `json:"content_object_id"`
	Closed          bool   `json:"closed"`
}

func losFromMap(in map[string]json.RawMessage) (*dbListOfSpeakers, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding list of speakers data: %w", err)
	}

	var los dbListOfSpeakers
	if err := json.Unmarshal(bs, &los); err != nil {
		return nil, fmt.Errorf("decoding list of speakers data: %w", err)
	}
	return &los, nil
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

func speakerFromMap(in map[string]json.RawMessage) (*dbSpeaker, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding speaker data: %w", err)
	}

	var speaker dbSpeaker
	if err := json.Unmarshal(bs, &speaker); err != nil {
		return nil, fmt.Errorf("decoding speaker data: %w", err)
	}
	return &speaker, nil
}

type dbChyronProjector struct {
	ChyronBackgroundColor string `json:"chyron_background_color"`
	ChyronFontColor       string `json:"chyron_font_color"`
}

func chyronProjectorFromMap(in map[string]json.RawMessage) (*dbChyronProjector, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding chyron projector data: %w", err)
	}

	var projector dbChyronProjector
	if err := json.Unmarshal(bs, &projector); err != nil {
		return nil, fmt.Errorf("decoding chyron projector data: %w", err)
	}
	return &projector, nil
}

// ListOfSpeaker renders current list of speaker slide.
func ListOfSpeaker(store *projector.SlideStore) {
	store.RegisterSliderFunc("list_of_speakers", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, hotkeys []string, err error) {
		return renderListOfSpeakers(ctx, ds, p7on.ContentObjectID, p7on.MeetingID)
	})
}

func renderListOfSpeakers(ctx context.Context, ds projector.Datastore, losFQID string, meetingID int) (encoded []byte, keys []string, err error) {
	fetch := datastore.NewFetcher(ds)
	defer func() {
		if err == nil {
			err = fetch.Error()
		}
	}()

	data := fetch.Object(ctx, []string{"speaker_ids", "content_object_id", "closed"}, losFQID)
	los, err := losFromMap(data)
	if err != nil {
		return nil, nil, fmt.Errorf("loading list of speakers: %w", err)
	}

	title := fetch.String(ctx, los.ContentObjectID+"/title")

	var speakersWaiting []outputSpeaker
	var speakersFinished []outputSpeaker
	var currentSpeaker *outputSpeaker
	for _, id := range los.SpeakerIDs {
		fields := []string{
			"user_id",
			"marked",
			"point_of_order",
			"weight",
			"begin_time",
			"end_time",
		}
		speaker, err := speakerFromMap(fetch.Object(ctx, fields, "speaker/%d", id))
		if err != nil {
			return nil, nil, fmt.Errorf("loading speaker: %w", err)
		}

		user, err := newUser(ctx, fetch, speaker.UserID, meetingID)
		if err != nil {
			return nil, nil, fmt.Errorf("loading user: %w", err)
		}

		s := outputSpeaker{
			User:         user.UserRepresentation(meetingID),
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

// get_losID determines the losID and first current_projection of the reference_projector
func getLosID(ctx context.Context, ContentObjectID string, fetch *datastore.Fetcher) (losID int, referenceProjectorID int, err error) {
	parts := strings.Split(ContentObjectID, "/")
	if len(parts) != 2 || parts[0] != "meeting" {
		return losID, referenceProjectorID, fmt.Errorf("invalid Content_objectID %s. Expected a meeting-objectID", ContentObjectID)
	}
	meetingID, err := strconv.Atoi(parts[1])
	if err != nil {
		return losID, referenceProjectorID, fmt.Errorf("invalid Content_objectID %s. Expected a numeric meeting_id", ContentObjectID)
	}
	referenceProjectorID = fetch.Int(ctx, "meeting/%d/reference_projector_id", meetingID)
	referenceP7onIDs := fetch.Ints(ctx, "projector/%d/current_projection_ids", referenceProjectorID)

	for _, pID := range referenceP7onIDs {
		contentObjectID := fetch.String(ctx, "projection/%d/content_object_id", pID)
		losID = fetch.Int(ctx, "%s/list_of_speakers_id", contentObjectID)

		if losID != 0 {
			break
		}
	}
	return losID, referenceProjectorID, nil
}

// CurrentListOfSpeakers renders the current_list_of_speakers slide.
func CurrentListOfSpeakers(store *projector.SlideStore) {
	store.RegisterSliderFunc("current_list_of_speakers", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

		losID, _, err := getLosID(ctx, p7on.ContentObjectID, fetch)
		if err != nil {
			return nil, nil, fmt.Errorf("error in getLosID: %w", err)
		}
		if losID == 0 {
			return []byte("{}"), fetch.Keys(), nil
		}

		if err := fetch.Error(); err != nil {
			return nil, nil, err
		}

		content, keys, err := renderListOfSpeakers(ctx, ds, fmt.Sprintf("list_of_speakers/%d", losID), p7on.MeetingID)
		if err != nil {
			return nil, nil, fmt.Errorf("render list of speakers %d: %w", losID, err)
		}
		keys = append(keys, fetch.Keys()...)
		return content, keys, nil
	})
}

func getCurrentSpeakerData(ctx context.Context, fetch *datastore.Fetcher, losID int, meetingID int) (shortName string, structureLevel string, err error) {
	data := fetch.Object(ctx, []string{"speaker_ids", "content_object_id", "closed"}, "list_of_speakers/%d", losID)
	los, err := losFromMap(data)
	if err != nil {
		return shortName, structureLevel, fmt.Errorf("loading list of speakers: %w", err)
	}

	fields := []string{
		"user_id",
		"begin_time",
		"end_time",
	}

	for _, id := range los.SpeakerIDs {
		speaker, err := speakerFromMap(fetch.Object(ctx, fields, "speaker/%d", id))
		if err != nil {
			return shortName, structureLevel, fmt.Errorf("loading speaker: %w", err)
		}

		if speaker.BeginTime == 0 || (speaker.BeginTime != 0 && speaker.EndTime != 0) {
			continue
		}

		user, err := newUser(ctx, fetch, speaker.UserID, meetingID)
		if err != nil {
			return shortName, structureLevel, fmt.Errorf("getting newUser: %w", err)
		}
		return user.UserShortName(), user.UserStructureLevel(meetingID), nil
	}
	return shortName, structureLevel, nil
}

// CurrentSpeakerChyron renders the current_speaker_chyron slide.
func CurrentSpeakerChyron(store *projector.SlideStore) {
	store.RegisterSliderFunc("current_speaker_chyron", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

		losID, projectorID, err := getLosID(ctx, p7on.ContentObjectID, fetch)
		if err != nil {
			return nil, nil, fmt.Errorf("error in getLosID: %w", err)
		}
		meetingID, _ := strconv.Atoi(strings.Split(p7on.ContentObjectID, "/")[1])

		data := fetch.Object(ctx, []string{"chyron_background_color", "chyron_font_color"}, "projector/%d", projectorID)
		projector, err := chyronProjectorFromMap(data)
		if err != nil {
			return nil, nil, fmt.Errorf("error in get chyron projector: %w", err)
		}

		shortName, structureLevel, err := getCurrentSpeakerData(ctx, fetch, losID, meetingID)
		if err != nil {
			return nil, nil, fmt.Errorf("get CurrentSpeakerData: %w", err)
		}

		out := struct {
			BackgroundColor     string `json:"background_color"`
			FontColor           string `json:"font_color"`
			CurrentSpeakerName  string `json:"current_speaker_name"`
			CurrentSpeakerLevel string `json:"current_speaker_level"`
		}{
			projector.ChyronBackgroundColor,
			projector.ChyronFontColor,
			shortName,
			structureLevel,
		}

		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response slide current_speaker_chyron: %w", err)
		}
		return responseValue, fetch.Keys(), err
	})
}
