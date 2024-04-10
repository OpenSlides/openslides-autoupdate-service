package slide

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dskey"
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

type dbSpeakerWork struct {
	MeetingUserID                  int `json:"meeting_user_id"`
	Weight                         int `json:"weight"`
	EndTime                        int `json:"end_time"`
	TotalPause                     int `json:"total_pause"`
	StructureLevelListOfSpeakersID int `json:"structure_level_list_of_speakers_id"`
}
type dbSpeaker struct {
	User         string         `json:"user"`
	SpeechState  string         `json:"speech_state"`
	Note         string         `json:"note"`
	BeginTime    int            `json:"begin_time,omitempty"`
	PauseTime    int            `json:"pause_time,omitempty"`
	PointOfOrder bool           `json:"point_of_order"`
	SpeakerWork  *dbSpeakerWork `json:",omitempty"`
}

func speakerFromMap(in map[string]json.RawMessage) (*dbSpeaker, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding speaker data: %w", err)
	}

	var speaker dbSpeaker
	var work dbSpeakerWork
	speaker.SpeakerWork = &work
	if err := json.Unmarshal(bs, &speaker); err != nil {
		return nil, fmt.Errorf("decoding speaker data: %w", err)
	}
	if err := json.Unmarshal(bs, &work); err != nil {
		return nil, fmt.Errorf("decoding speaker work data: %w", err)
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
	store.RegisterSliderFunc("list_of_speakers", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		return renderListOfSpeakers(ctx, fetch, p7on.ContentObjectID, p7on.MeetingID, store)
	})
}

// CurrentListOfSpeakers renders the current_list_of_speakers slide.
func CurrentListOfSpeakers(store *projector.SlideStore) {
	store.RegisterSliderFunc("current_list_of_speakers", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		losID, _, err := getLosID(ctx, p7on.ContentObjectID, fetch)
		if err != nil {
			return nil, fmt.Errorf("error in getLosID: %w", err)
		}
		if losID == 0 {
			return []byte("{}"), nil
		}

		if err := fetch.Err(); err != nil {
			return nil, err
		}

		content, err := renderListOfSpeakers(ctx, fetch, fmt.Sprintf("list_of_speakers/%d", losID), p7on.MeetingID, store)
		if err != nil {
			return nil, fmt.Errorf("render list of speakers %d: %w", losID, err)
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}
		return content, nil
	})
}

// CurrentSpeakerChyron renders the current_speaker_chyron slide.
func CurrentSpeakerChyron(store *projector.SlideStore) {
	store.RegisterSliderFunc("current_speaker_chyron", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		losID, referenceProjectorID, err := getLosID(ctx, p7on.ContentObjectID, fetch)
		if err != nil {
			return nil, fmt.Errorf("error in getLosID: %w", err)
		}

		projectorID := p7on.CurrentProjectorID
		if projectorID <= 0 {
			projectorID = referenceProjectorID
		}

		meetingID, err := strconv.Atoi(strings.Split(p7on.ContentObjectID, "/")[1])
		if err != nil {
			return nil, fmt.Errorf("error in Atoi with ContentObjectID: %w", err)
		}

		projector := &dbChyronProjector{}
		if projectorID > 0 {
			data := fetch.Object(ctx, fmt.Sprintf("projector/%d", projectorID), "chyron_background_color", "chyron_font_color")
			projector, err = chyronProjectorFromMap(data)
			if err != nil {
				return nil, fmt.Errorf("error in get chyron projector: %w", err)
			}
		}

		var shortName, structureLevel string
		if losID > 0 {
			shortName, structureLevel, err = getCurrentSpeakerData(ctx, fetch, losID, meetingID)
			if err != nil {
				return nil, fmt.Errorf("get CurrentSpeakerData: %w", err)
			}
			if err := fetch.Err(); err != nil {
				return nil, err
			}
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
			return nil, fmt.Errorf("encoding response slide current_speaker_chyron: %w", err)
		}
		return responseValue, nil
	})
}

type dbStructureLevel struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func structureLevelFromMap(in map[string]json.RawMessage) (*dbStructureLevel, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding motion data: %w", err)
	}

	var m dbStructureLevel
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, fmt.Errorf("decoding motion: %w", err)
	}
	return &m, nil
}

type dbStructureLevelListOfSpeakers struct {
	SpeakerIDs       []int `json:"speaker_ids"`
	StructureLevelID int   `json:"structure_level_id"`
	InitialTime      int   `json:"initial_time"`
	RemainingTime    int   `json:"remaining_time"`
	AdditionalTime   int   `json:"additional_time"`
	CurrentStartTime int   `json:"current_start_time"`
}

func structureLevelListOfSpeakersFromMap(in map[string]json.RawMessage) (*dbStructureLevelListOfSpeakers, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding motion data: %w", err)
	}

	var m dbStructureLevelListOfSpeakers
	if err := json.Unmarshal(bs, &m); err != nil {
		return nil, fmt.Errorf("decoding motion: %w", err)
	}
	return &m, nil
}

type structureLevelRepr struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	Color            string `json:"color"`
	SpeechState      string `json:"speech_state"`
	PointOfOrder     bool   `json:"point_of_order"`
	RemainingTime    *int   `json:"remaining_time,omitempty"`
	CurrentStartTime int    `json:"current_start_time"`
}

// CurrentStructureLevelList renders the current_structure_level_list slide.
func CurrentStructureLevelList(store *projector.SlideStore) {
	store.RegisterSliderFunc("current_structure_level_list", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		losID, _, err := getLosID(ctx, p7on.ContentObjectID, fetch)
		if err != nil {
			return nil, fmt.Errorf("error in getLosID: %w", err)
		}

		var losContentObject string
		fetch.Fetch(ctx, &losContentObject, "list_of_speakers/%d/content_object_id", losID)
		if err := fetch.Err(); err != nil {
			return nil, fmt.Errorf("getting content object for list of speakers %d: %w", losID, err)
		}

		var title string
		fetch.Fetch(ctx, &title, "%s/%s", losContentObject, "title")
		if err := fetch.Err(); err != nil {
			return nil, fmt.Errorf("getting title for list of speakers content object %s: %w", losContentObject, err)
		}

		var structureLevelListOfSpeakersIds []int
		fetch.Fetch(ctx, &structureLevelListOfSpeakersIds, "list_of_speakers/%d/structure_level_list_of_speakers_ids", losID)
		if err := fetch.Err(); err != nil {
			return nil, fmt.Errorf("getting structure_level_list_of_speakers_ids for list of speakers %d: %w", losID, err)
		}

		structureLevels := []structureLevelRepr{}
		for _, slsID := range structureLevelListOfSpeakersIds {
			hasSpeaker, err := structureLevelHasSpeaker(ctx, fetch, slsID)
			if err != nil {
				return nil, fmt.Errorf("checking speakers structure level los %d for list of speakers %d: %w", slsID, losID, err)
			}

			if !hasSpeaker {
				continue
			}

			slsData := fetch.Object(ctx, fmt.Sprintf("structure_level_list_of_speakers/%d", slsID), "structure_level_id", "remaining_time", "current_start_time")
			sls, err := structureLevelListOfSpeakersFromMap(slsData)
			if err != nil {
				return nil, fmt.Errorf("parsing structure level los %d for list of speakers %d: %w", slsID, losID, err)
			}

			slData := fetch.Object(ctx, fmt.Sprintf("structure_level/%d", sls.StructureLevelID), "name", "color")
			sl, err := structureLevelFromMap(slData)
			if err != nil {
				return nil, fmt.Errorf("parsing structure level %d for list of speakers %d: %w", sls.StructureLevelID, losID, err)
			}

			structureLevel := structureLevelRepr{
				ID:               sls.StructureLevelID,
				Name:             sl.Name,
				Color:            sl.Color,
				RemainingTime:    &sls.RemainingTime,
				CurrentStartTime: sls.CurrentStartTime,
			}
			structureLevels = append(structureLevels, structureLevel)
		}

		out := struct {
			Title           string               `json:"title"`
			StructureLevels []structureLevelRepr `json:"structure_levels"`
		}{
			title,
			structureLevels,
		}

		responseValue, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("encoding response slide current_speaker_chyron: %w", err)
		}
		return responseValue, nil
	})
}

// CurrentSpeakingStructureLevel renders the current_speaking_structure_level slide.
func CurrentSpeakingStructureLevel(store *projector.SlideStore) {
	store.RegisterSliderFunc("current_speaking_structure_level", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		losID, _, err := getLosID(ctx, p7on.ContentObjectID, fetch)
		if err != nil {
			return nil, fmt.Errorf("error in getLosID: %w", err)
		}

		speaker, err := getStructureLevelData(ctx, fetch, losID)
		if err != nil {
			return nil, fmt.Errorf("error in getStructureLevelData: %w", err)
		}

		if speaker != nil {
			slsID := speaker.SpeakerWork.StructureLevelListOfSpeakersID
			out := structureLevelRepr{
				SpeechState:  speaker.SpeechState,
				PointOfOrder: speaker.PointOfOrder,
			}

			if slsID != 0 {
				slsData := fetch.Object(ctx, fmt.Sprintf("structure_level_list_of_speakers/%d", slsID), "structure_level_id", "remaining_time", "current_start_time")
				sls, err := structureLevelListOfSpeakersFromMap(slsData)
				if err != nil {
					return nil, fmt.Errorf("parsing structure level los %d for list of speakers %d: %w", slsID, losID, err)
				}

				slData := fetch.Object(ctx, fmt.Sprintf("structure_level/%d", sls.StructureLevelID), "name", "color")
				sl, err := structureLevelFromMap(slData)
				if err != nil {
					return nil, fmt.Errorf("parsing structure level %d for list of speakers %d: %w", sls.StructureLevelID, losID, err)
				}

				out.ID = sls.StructureLevelID
				out.Name = sl.Name
				out.Color = sl.Color
				out.RemainingTime = &sls.RemainingTime
				out.CurrentStartTime = sls.CurrentStartTime
			}

			if speaker.SpeechState == "interposed_question" || speaker.SpeechState == "intervention" || speaker.PointOfOrder || slsID == 0 {
				out.RemainingTime = nil
				if speaker.SpeechState == "intervention" {
					meetingID := datastore.Int(ctx, fetch.FetchIfExist, "list_of_speakers/%d/meeting_id", losID)
					if err := fetch.Err(); err != nil {
						return nil, fmt.Errorf("Error loading meeting id from los %d %w", losID, err)
					}
					interventionTime := datastore.Int(ctx, fetch.FetchIfExist, "meeting/%d/list_of_speakers_intervention_time", meetingID)
					if err := fetch.Err(); err != nil {
						return nil, fmt.Errorf("Error loading intervention time from meeting %d %w", meetingID, err)
					}
					out.RemainingTime = &interventionTime
				}
				if speaker.PauseTime != 0 {
					if out.RemainingTime == nil {
						out.CurrentStartTime = 0
					} else {
						out.CurrentStartTime = 0
						*out.RemainingTime -= speaker.PauseTime - (speaker.BeginTime + speaker.SpeakerWork.TotalPause)
					}
				} else {
					out.CurrentStartTime = speaker.BeginTime + speaker.SpeakerWork.TotalPause
				}
			}

			responseValue, err := json.Marshal(out)
			if err != nil {
				return nil, fmt.Errorf("encoding response slide current_speaking_structure_level: %w", err)
			}
			return responseValue, nil
		}

		return []byte("{}"), nil
	})
}

func structureLevelHasSpeaker(ctx context.Context, fetch *datastore.Fetcher, structureLevelLosID int) (spoken bool, err error) {
	data := fetch.Object(ctx, fmt.Sprintf("structure_level_list_of_speakers/%d", structureLevelLosID), "speaker_ids", "initial_time", "additional_time", "remaining_time", "current_start_time")
	sllos, err := structureLevelListOfSpeakersFromMap(data)
	if err != nil {
		return false, fmt.Errorf("loading structure level list of speakers: %w", err)
	}

	if sllos.InitialTime+sllos.AdditionalTime != sllos.RemainingTime || sllos.CurrentStartTime != 0 {
		return true, nil
	}

	for _, id := range sllos.SpeakerIDs {
		speechState := datastore.String(ctx, fetch.FetchIfExist, "speaker/%d/speech_state", id)
		if err := fetch.Err(); err != nil {
			return false, fmt.Errorf("Error loading speach state %d %w", id, err)
		}

		if speechState == "interposed_question" || speechState == "intervention" {
			continue
		}

		return true, nil
	}

	return false, nil
}

// getLosID determines the losID and first current_projection of the reference_projector.
func getLosID(ctx context.Context, ContentObjectID string, fetch *datastore.Fetcher) (losID int, referenceProjectorID int, err error) {
	parts := strings.Split(ContentObjectID, "/")
	if len(parts) != 2 || parts[0] != "meeting" {
		return losID, referenceProjectorID, fmt.Errorf("invalid ContentObjectID %s. Expected a meeting-objectID", ContentObjectID)
	}
	meetingID, err := strconv.Atoi(parts[1])
	if err != nil {
		return losID, referenceProjectorID, fmt.Errorf("invalid ContentObjectID %s. Expected a numeric meeting_id", ContentObjectID)
	}
	referenceProjectorID = datastore.Int(ctx, fetch.FetchIfExist, "meeting/%d/reference_projector_id", meetingID)
	referenceP7onIDs := datastore.Ints(ctx, fetch.FetchIfExist, "projector/%d/current_projection_ids", referenceProjectorID)
	if err := fetch.Err(); err != nil {
		return losID, referenceProjectorID, err
	}

	for _, pID := range referenceP7onIDs {
		contentObjectID := datastore.String(ctx, fetch.FetchIfExist, "projection/%d/content_object_id", pID)
		if err := fetch.Err(); err != nil {
			return 0, 0, fmt.Errorf("fetching projection/%d/content_object_id: %w", pID, err)
		}

		if contentObjectID == "" {
			continue
		}
		losID = datastore.Int(ctx, fetch.FetchIfExist, "%s/list_of_speakers_id", contentObjectID)
		if err := fetch.Err(); err != nil {
			var errInvalidKey dskey.InvalidKeyError
			if !errors.As(err, &errInvalidKey) {
				return 0, 0, fmt.Errorf("%s/content_object_id: %w", contentObjectID, err)
			}
		}

		if losID != 0 {
			break
		}
	}

	return losID, referenceProjectorID, nil
}

func getStructureLevelData(ctx context.Context, fetch *datastore.Fetcher, losID int) (speaker *dbSpeaker, err error) {
	data := fetch.Object(ctx, fmt.Sprintf("list_of_speakers/%d", losID), "speaker_ids", "content_object_id", "closed")
	los, err := losFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("loading list of speakers: %w", err)
	}

	fields := []string{
		"begin_time",
		"pause_time",
		"end_time",
		"total_pause",
		"point_of_order",
		"weight",
		"speech_state",
		"structure_level_list_of_speakers_id",
	}

	var fallbackSpeaker *dbSpeaker
	for _, id := range los.SpeakerIDs {
		speaker, err := speakerFromMap(fetch.Object(ctx, fmt.Sprintf("speaker/%d", id), fields...))
		if err != nil {
			return nil, fmt.Errorf("loading speaker %d: %w", id, err)
		}

		if (fallbackSpeaker == nil || fallbackSpeaker.SpeakerWork.Weight > speaker.SpeakerWork.Weight) && speaker.BeginTime != 0 && speaker.SpeakerWork.EndTime == 0 {
			fallbackSpeaker = speaker
		}

		if speaker.BeginTime == 0 || speaker.PauseTime != 0 || (speaker.BeginTime != 0 && speaker.SpeakerWork.EndTime != 0) {
			continue
		}

		return speaker, nil
	}

	if fallbackSpeaker != nil {
		return fallbackSpeaker, nil
	}

	return nil, nil
}

func getCurrentSpeakerData(ctx context.Context, fetch *datastore.Fetcher, losID int, meetingID int) (shortName string, structureLevel string, err error) {
	data := fetch.Object(ctx, fmt.Sprintf("list_of_speakers/%d", losID), "speaker_ids", "content_object_id", "closed")
	los, err := losFromMap(data)
	if err != nil {
		return "", "", fmt.Errorf("loading list of speakers: %w", err)
	}

	fields := []string{
		"meeting_user_id",
		"begin_time",
		"pause_time",
		"end_time",
	}

	for _, id := range los.SpeakerIDs {
		speaker, err := speakerFromMap(fetch.Object(ctx, fmt.Sprintf("speaker/%d", id), fields...))
		if err != nil {
			return "", "", fmt.Errorf("loading speaker %d: %w", id, err)
		}

		if speaker.BeginTime == 0 || speaker.PauseTime != 0 || (speaker.BeginTime != 0 && speaker.SpeakerWork.EndTime != 0) {
			continue
		}

		if speaker.SpeakerWork.MeetingUserID != 0 {
			var userID int
			fetch.FetchIfExist(ctx, &userID, "meeting_user/%d/user_id", speaker.SpeakerWork.MeetingUserID)
			if err := fetch.Err(); err != nil {
				return "", "", fmt.Errorf("getting user for meeting user %d: %w", speaker.SpeakerWork.MeetingUserID, err)
			}

			user, err := NewUser(ctx, fetch, userID, meetingID)
			if err != nil {
				return "", "", fmt.Errorf("getting newUser: %w", err)
			}

			structureLevelTime := datastore.Int(ctx, fetch.FetchIfExist, "meeting/%d/list_of_speakers_default_structure_level_time", meetingID)
			structureLevelName := ""
			if structureLevelTime > 0 {
				var structureLevelListOfSpeakersID int
				fetch.FetchIfExist(ctx, &structureLevelListOfSpeakersID, "speaker/%d/structure_level_list_of_speakers_id", id)
				if err := fetch.Err(); err != nil {
					return "", "", fmt.Errorf("getting structure level for speaker %d: %w", id, err)
				}

				if structureLevelListOfSpeakersID != 0 {
					var structureLevelID int
					fetch.FetchIfExist(ctx, &structureLevelID, "structure_level_list_of_speakers/%d/structure_level_id", structureLevelListOfSpeakersID)
					if err := fetch.Err(); err != nil {
						return "", "", fmt.Errorf("getting structure level for structure_level_list_of_speakers %d: %w", structureLevelListOfSpeakersID, err)
					}

					fetch.Fetch(ctx, &structureLevelName, "structure_level/%d/name", structureLevelID)
					if err := fetch.Err(); err != nil {
						return "", "", fmt.Errorf("getting name for structure level name %d: %w", structureLevelID, err)
					}
				}
			} else {
				var structureLevelIds []int
				fetch.Fetch(ctx, &structureLevelIds, "meeting_user/%d/structure_level_ids", speaker.SpeakerWork.MeetingUserID)
				if len(structureLevelIds) > 0 {
					structureLevelNames := make([]string, len(structureLevelIds))
					for i, id := range structureLevelIds {
						structureLevelNames[i] = datastore.String(ctx, fetch.FetchIfExist, "structure_level/%d/name", id)
					}
					sort.Strings(structureLevelNames)
					structureLevelName = strings.Join(structureLevelNames, ", ")
				}
			}

			return user.UserShortName(), structureLevelName, nil
		}
		return "", "", nil
	}

	return shortName, structureLevel, nil
}

func renderListOfSpeakers(ctx context.Context, fetch *datastore.Fetcher, losFQID string, meetingID int, store *projector.SlideStore) (encoded []byte, err error) {
	data := fetch.Object(ctx, losFQID, "speaker_ids", "content_object_id", "closed")
	los, err := losFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("loading list of speakers: %w", err)
	}

	var speakersWaiting []dbSpeaker
	var speakersFinished []dbSpeaker
	currentSpeaker, numberOfWaitingSpeakers, err := getSpeakerLists(ctx, los, meetingID, fetch, &speakersWaiting, &speakersFinished)
	if err != nil {
		return nil, fmt.Errorf("getSpeakersList: %w", err)
	}

	if err := fetch.Err(); err != nil {
		return nil, err
	}

	parts := strings.Split(los.ContentObjectID, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("splitting ComtentObjectID: %w", err)
	}
	collection := parts[0]
	if err != nil {
		return nil, fmt.Errorf("get ID from ContentObjectID: %w", err)
	}

	titler := store.GetTitleInformationFunc(collection)
	if titler == nil {
		return nil, fmt.Errorf("no titler function registered for %s", collection)
	}

	titleInfo, err := titler.GetTitleInformation(ctx, fetch, los.ContentObjectID, "", meetingID)
	if err != nil {
		return nil, fmt.Errorf("get title func: %w", err)
	}

	slideData := struct {
		Waiting                 []dbSpeaker     `json:"waiting"`
		Current                 *dbSpeaker      `json:"current,"`
		Finished                []dbSpeaker     `json:"finished"`
		TitleInformation        json.RawMessage `json:"title_information"`
		Closed                  bool            `json:"closed"`
		NumberOfWaitingSpeakers *int            `json:"number_of_waiting_speakers,omitempty"`
	}{
		speakersWaiting,
		currentSpeaker,
		speakersFinished,
		titleInfo,
		los.Closed,
		numberOfWaitingSpeakers,
	}
	b, err := json.Marshal(slideData)
	if err != nil {
		return nil, fmt.Errorf("encoding outgoing data: %w", err)
	}
	if err := fetch.Err(); err != nil {
		return nil, err
	}

	return b, nil
}

func getSpeakerLists(ctx context.Context, los *dbListOfSpeakers, meetingID int, fetch *datastore.Fetcher, speakersWaiting *[]dbSpeaker, speakersFinished *[]dbSpeaker) (*dbSpeaker, *int, error) {
	fields := []string{
		"meeting_user_id",
		"speech_state",
		"note",
		"point_of_order",
		"weight",
		"begin_time",
		"pause_time",
		"end_time",
	}

	var currentSpeaker *dbSpeaker
	var numberOfWaitingSpeakers *int
	for _, id := range los.SpeakerIDs {
		speaker, err := speakerFromMap(fetch.Object(ctx, fmt.Sprintf("speaker/%d", id), fields...))
		if err != nil {
			return nil, nil, fmt.Errorf("loading speaker: %w", err)
		}

		if speaker.SpeakerWork.MeetingUserID != 0 {
			var userID int
			fetch.FetchIfExist(ctx, &userID, "meeting_user/%d/user_id", speaker.SpeakerWork.MeetingUserID)
			if err := fetch.Err(); err != nil {
				return nil, nil, fmt.Errorf("getting user for meeting user %d: %w", speaker.SpeakerWork.MeetingUserID, err)
			}

			user, err := NewUser(ctx, fetch, userID, meetingID)
			if err != nil {
				return nil, nil, fmt.Errorf("loading user: %w", err)
			}

			speaker.User = user.UserRepresentation(meetingID)
		}

		if (speaker.BeginTime == 0 || speaker.SpeechState == "interposed_question") && speaker.SpeakerWork.EndTime == 0 {
			*speakersWaiting = append(*speakersWaiting, *speaker)
			continue
		}

		if speaker.SpeakerWork.EndTime == 0 {
			currentSpeaker = speaker
			continue
		}

		*speakersFinished = append(*speakersFinished, *speaker)
	}

	// Sort ascending by weight
	sort.Slice(*speakersWaiting, func(i, j int) bool {
		if (*speakersWaiting)[i].SpeakerWork.Weight == (*speakersWaiting)[j].SpeakerWork.Weight {
			return (*speakersWaiting)[i].SpeakerWork.MeetingUserID < (*speakersWaiting)[j].SpeakerWork.MeetingUserID
		}
		return (*speakersWaiting)[i].SpeakerWork.Weight < (*speakersWaiting)[j].SpeakerWork.Weight
	})

	// Sort descending by endtime to get lates at top position
	sort.Slice(*speakersFinished, func(i, j int) bool {
		return (*speakersFinished)[i].SpeakerWork.EndTime > (*speakersFinished)[j].SpeakerWork.EndTime
	})

	meeting, err := getMeeting(ctx, fetch, meetingID, []string{"list_of_speakers_amount_next_on_projector", "list_of_speakers_amount_last_on_projector", "list_of_speakers_show_amount_of_speakers_on_slide"})
	if err != nil {
		return nil, nil, fmt.Errorf("reading meeting: %w", err)
	}
	if err := fetch.Err(); err != nil {
		return nil, nil, err
	}

	if meeting.ListOfSpeakersShowAmountOfSpeakersOnSlide {
		number := len(*speakersWaiting)
		numberOfWaitingSpeakers = &number
	}

	if len(*speakersWaiting) >= 1 || len(*speakersFinished) >= 1 {
		if len(*speakersWaiting) >= 1 && meeting.ListOfSpeakersAmountNextOnProjector >= 0 && meeting.ListOfSpeakersAmountNextOnProjector < len(*speakersWaiting) && meeting.ListOfSpeakersShowAmountOfSpeakersOnSlide {
			*speakersWaiting = (*speakersWaiting)[:meeting.ListOfSpeakersAmountNextOnProjector]
		}
		if len(*speakersFinished) >= 1 && meeting.ListOfSpeakersAmountLastOnProjector >= 0 && meeting.ListOfSpeakersAmountLastOnProjector < len(*speakersFinished) {
			*speakersFinished = (*speakersFinished)[:meeting.ListOfSpeakersAmountLastOnProjector]
		}
	}

	// Remove SpeakerWork's
	for i := range *speakersWaiting {
		(*speakersWaiting)[i].SpeakerWork = nil
	}
	for i := range *speakersFinished {
		(*speakersFinished)[i].SpeakerWork = nil
	}
	if currentSpeaker != nil {
		currentSpeaker.SpeakerWork = nil
	}
	return currentSpeaker, numberOfWaitingSpeakers, nil
}
