package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore"
)

type optionRepr struct {
	Text            string          `json:"text"`
	ContentObjectID string          `json:"content_object_id"`
	ContentObject   json.RawMessage `json:"content_object"`
	Yes             *float64        `json:"yes,omitempty"`
	No              *float64        `json:"no,omitempty"`
	Abstain         *float64        `json:"abstain,omitempty"`
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
	Yes     float64 `json:"yes"`
	No      float64 `json:"no"`
	Abstain float64 `json:"abstain"`
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
	ID                    int             `json:"id"`
	ContentObjectID       string          `json:"content_object_id"`
	TitleInformation      json.RawMessage `json:"title_information"`
	Title                 string          `json:"title"`
	Description           string          `json:"description"`
	Type                  string          `json:"type"`
	State                 string          `json:"state"`
	GlobalYes             bool            `json:"global_yes"`
	GlobalNo              bool            `json:"global_no"`
	GlobalAbstain         bool            `json:"global_abstain"`
	Options               []*optionRepr   `json:"options"`
	IsPseudoanonymized    *bool           `json:"is_pseudoanonymized,omitempty"`
	Pollmethod            *string         `json:"pollmethod,omitempty"`
	OnehundredPercentBase *string         `json:"onehundred_percent_base,omitempty"`
	MajorityMethod        *string         `json:"majority_method,omitempty"`
	Votesvalid            *float32        `json:"votesvalid,omitempty"`
	Votesinvalid          *float32        `json:"votesinvalid,omitempty"`
	Votescast             *float32        `json:"votescast,omitempty"`
	GlobalOption          *optionGlobRepr `json:"global_option,omitempty"`
	PollWork              *dbPollWork     `json:",omitempty"`
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
	store.RegisterSliderFunc("poll", func(ctx context.Context, ds projector.Datastore, p7on *projector.Projection) (encoded []byte, keys []string, err error) {
		fetch := datastore.NewFetcher(ds)
		defer func() {
			if err == nil {
				err = fetch.Error()
			}
		}()

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
		state := fetch.String(ctx, "%s/%s", p7on.ContentObjectID, "state")
		if state == "published" {
			fetchFields = append(fetchFields, []string{
				"is_pseudoanonymized",
				"pollmethod",
				"onehundred_percent_base",
				"majority_method",
				"votesvalid",
				"votesinvalid",
				"votescast",
				"global_option_id",
			}...)
		}
		data := fetch.Object(ctx, fetchFields, p7on.ContentObjectID)

		poll, err := pollFromMap(data, state)
		if err != nil {
			return nil, nil, fmt.Errorf("get poll: %w", err)
		}

		poll.TitleInformation, err = getTitleInfoFromContentObject(ctx, fetch, store, poll.ContentObjectID, "", p7on.MeetingID)
		if err != nil {
			return nil, nil, fmt.Errorf("getTitleInfoFromContentObject: %w", err)
		}

		poll.Options, err = getOptions(ctx, fetch, store, poll.PollWork.OptionIDS, state, p7on.MeetingID)
		if err != nil {
			return nil, nil, fmt.Errorf("get Options func: %w", err)
		}
		if state == "published" {
			poll.GlobalOption, err = getGlobalOption(ctx, fetch, store, poll.PollWork.GlobalOptionID)
			if err != nil {
				return nil, nil, fmt.Errorf("get GlobalOption func: %w", err)
			}
		}
		poll.PollWork = nil // don't export
		responseValue, err := json.Marshal(poll)
		if err != nil {
			return nil, nil, fmt.Errorf("encoding response slide poll: %w", err)
		}
		return responseValue, fetch.Keys(), err
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
		data := fetch.Object(ctx, fetchFields, "option/%d", optionID)
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
	sort.Slice(options, func(i, j int) bool {
		if *options[i].weight == *options[j].weight {
			return *options[i].id < *options[j].id
		}
		return *options[i].weight < *options[j].weight
	})

	return options, nil
}

func getGlobalOption(ctx context.Context, fetch *datastore.Fetcher, store *projector.SlideStore, globalOptionID int) (*optionGlobRepr, error) {
	data := fetch.Object(ctx, []string{"yes", "no", "abstain"}, "option/%d", globalOptionID)
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
