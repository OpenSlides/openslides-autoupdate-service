package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

type optionRepr struct {
	Text            string          `json:"text"`
	ContentObjectID string          `json:"content_object_id"`
	ContentObject   json.RawMessage `json:"content_object"`
	Yes             *string         `json:"yes,omitempty"`     // Python-DecimalField
	No              *string         `json:"no,omitempty"`      // Python-DecimalField
	Abstain         *string         `json:"abstain,omitempty"` // Python-DecimalField
	id              *int
	weight          *int
}

func optionFromMap(in map[string]json.RawMessage) (*optionRepr, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding option data: %w", err)
	}

	var or optionRepr
	if err := json.Unmarshal(bs, &or); err != nil {
		return nil, fmt.Errorf("decoding option data: %w", err)
	}
	if err := json.Unmarshal(in["weight"], &or.weight); err != nil {
		return nil, fmt.Errorf("decoding option weight: %w", err)
	}
	if err := json.Unmarshal(in["id"], &or.id); err != nil {
		return nil, fmt.Errorf("decoding option id: %w", err)
	}
	return &or, nil
}

type optionGlobRepr struct {
	Yes     string `json:"yes"`     // Python-DecimalField
	No      string `json:"no"`      // Python-DecimalField
	Abstain string `json:"abstain"` // Python-DecimalField
}

func optionGlobFromMap(in map[string]json.RawMessage) (*optionGlobRepr, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding option global data: %w", err)
	}

	var og optionGlobRepr
	if err := json.Unmarshal(bs, &og); err != nil {
		return nil, fmt.Errorf("decoding option glob data: %w", err)
	}
	return &og, nil
}

// Contains fields to be read, but never exported
type dbPollWork struct {
	OptionIDS      []int `json:"option_ids"`
	MeetingID      int   `json:"meeting_id"`
	GlobalOptionID int   `json:"global_option_id"`
}

type dbPoll struct {
	ID                    int              `json:"id"`
	ContentObjectID       string           `json:"content_object_id"`
	TitleInformation      json.RawMessage  `json:"title_information"`
	Title                 string           `json:"title"`
	Description           string           `json:"description"`
	Type                  string           `json:"type"`
	State                 string           `json:"state"`
	GlobalYes             bool             `json:"global_yes"`
	GlobalNo              bool             `json:"global_no"`
	GlobalAbstain         bool             `json:"global_abstain"`
	Options               []*optionRepr    `json:"options"`
	EntitledUsersAtStop   *json.RawMessage `json:"entitled_users_at_stop,omitempty"`
	IsPseudoanonymized    *bool            `json:"is_pseudoanonymized,omitempty"`
	Pollmethod            *string          `json:"pollmethod,omitempty"`
	OnehundredPercentBase *string          `json:"onehundred_percent_base,omitempty"`
	Votesvalid            *string          `json:"votesvalid,omitempty"`   // Python-DecimalField
	Votesinvalid          *string          `json:"votesinvalid,omitempty"` // Python-DecimalField
	Votescast             *string          `json:"votescast,omitempty"`    // Python-DecimalField
	GlobalOption          *optionGlobRepr  `json:"global_option,omitempty"`
	PollWork              *dbPollWork      `json:",omitempty"`
}

func pollFromMap(in map[string]json.RawMessage, state string) (*dbPoll, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding poll data: %w", err)
	}

	var po dbPoll
	var poWork dbPollWork
	po.PollWork = &poWork
	if err := json.Unmarshal(bs, &po); err != nil {
		return nil, fmt.Errorf("decoding poll data: %w", err)
	}
	if err := json.Unmarshal(bs, &poWork); err != nil {
		return nil, fmt.Errorf("decoding poll work data: %w", err)
	}
	return &po, nil
}

// Poll renders the poll slide.
func Poll(store *projector.SlideStore) {
	store.RegisterSliderFunc("poll", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		fetchFields := []string{
			"id",
			"content_object_id",
			"title",
			"description",
			"type",
			"state",
			"global_yes",
			"global_no",
			"global_abstain",
			"option_ids",
			"meeting_id",
		}
		state := datastore.String(ctx, fetch.FetchIfExist, "%s/%s", p7on.ContentObjectID, "state")
		if state == "published" {
			fetchFields = append(fetchFields, []string{
				"entitled_users_at_stop",
				"is_pseudoanonymized",
				"pollmethod",
				"onehundred_percent_base",
				"votesvalid",
				"votesinvalid",
				"votescast",
				"global_option_id",
			}...)
		}
		data := fetch.Object(ctx, p7on.ContentObjectID, fetchFields...)

		poll, err := pollFromMap(data, state)
		if err != nil {
			return nil, fmt.Errorf("get poll: %w", err)
		}

		poll.TitleInformation, err = getTitleInfoFromContentObject(ctx, fetch, store, poll.ContentObjectID, "", p7on.MeetingID)
		if err != nil {
			return nil, fmt.Errorf("getTitleInfoFromContentObject: %w", err)
		}

		poll.Options, err = getOptions(ctx, fetch, store, poll.PollWork.OptionIDS, state, p7on.MeetingID)
		if err != nil {
			return nil, fmt.Errorf("get Options func: %w", err)
		}
		if state == "published" {
			poll.GlobalOption, err = getGlobalOption(ctx, fetch, store, poll.PollWork.GlobalOptionID)
			if err != nil {
				return nil, fmt.Errorf("get GlobalOption func: %w", err)
			}
		}
		if err := fetch.Err(); err != nil {
			return nil, err
		}

		poll.PollWork = nil // don't export
		responseValue, err := json.Marshal(poll)
		if err != nil {
			return nil, fmt.Errorf("encoding response slide poll: %w", err)
		}
		return responseValue, err
	})
}

func getOptions(ctx context.Context, fetch *datastore.Fetcher, store *projector.SlideStore, optionIDS []int, state string, meetingID int) (options []*optionRepr, err error) {
	fetchFields := []string{
		"text",
		"content_object_id",
		"weight",
		"id",
	}
	if state == "published" {
		fetchFields = append(fetchFields, []string{
			"yes",
			"no",
			"abstain",
		}...)
	}

	for _, optionID := range optionIDS {
		data := fetch.Object(ctx, fmt.Sprintf("option/%d", optionID), fetchFields...)
		option, err := optionFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get option data: %w", err)
		}
		option.ContentObject, err = getTitleInfoFromContentObject(ctx, fetch, store, option.ContentObjectID, "", meetingID)
		if err != nil {
			return nil, fmt.Errorf("getTitleInfoFromContentObject: %w", err)
		}
		options = append(options, option)
	}
	if err := fetch.Err(); err != nil {
		return nil, err
	}

	sort.Slice(options, func(i, j int) bool {
		if *options[i].weight == *options[j].weight {
			return *options[i].id < *options[j].id
		}
		return *options[i].weight < *options[j].weight
	})

	return options, nil
}

func getGlobalOption(ctx context.Context, fetch *datastore.Fetcher, store *projector.SlideStore, globalOptionID int) (*optionGlobRepr, error) {
	data := fetch.Object(ctx, fmt.Sprintf("option/%d", globalOptionID), "yes", "no", "abstain")
	globalOption, err := optionGlobFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("get option data: %w", err)
	}
	return globalOption, nil
}

// getTitleInfoFromContentObject gets GetTitleInformation from ContentObject
func getTitleInfoFromContentObject(ctx context.Context, fetch *datastore.Fetcher, store *projector.SlideStore, contentObjectID string, itemNumber string, meetingID int) (json.RawMessage, error) {
	collection := strings.Split(contentObjectID, "/")[0]
	titler := store.GetTitleInformationFunc(collection)
	if titler == nil {
		return nil, fmt.Errorf("no titler function registered for %s", collection)
	}
	titleInfo, err := titler.GetTitleInformation(ctx, fetch, contentObjectID, "", meetingID)
	if err != nil {
		return nil, fmt.Errorf("get title func: %w", err)
	}
	return titleInfo, nil
}
