package slide

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/projector/datastore"
)

type pollVoteRepr struct {
	UserId *int          `json:"user_id"`
	User   *pollUserRepr `json:"user"`
	Value  string        `json:"value"`
}

func pollVoteFromMap(in map[string]json.RawMessage) (*pollVoteRepr, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding vote data: %w", err)
	}

	var vr pollVoteRepr
	if err := json.Unmarshal(bs, &vr); err != nil {
		return nil, fmt.Errorf("decoding vote data: %w", err)
	}

	return &vr, nil
}

type pollUserRepr struct {
	Id        int     `json:"id"`
	Title     *string `json:"title,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

func pollUserFromMap(in map[string]json.RawMessage) (*pollUserRepr, error) {
	bs, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("encoding user data: %w", err)
	}

	var ur pollUserRepr
	if err := json.Unmarshal(bs, &ur); err != nil {
		return nil, fmt.Errorf("decoding user data: %w", err)
	}

	return &ur, nil
}

type optionRepr struct {
	Text            string          `json:"text,omitempty"`
	ContentObjectID string          `json:"content_object_id,omitempty"`
	ContentObject   json.RawMessage `json:"content_object,omitempty"`
	Yes             *string         `json:"yes,omitempty"`     // Python-DecimalField
	No              *string         `json:"no,omitempty"`      // Python-DecimalField
	Abstain         *string         `json:"abstain,omitempty"` // Python-DecimalField
	Votes           []*pollVoteRepr `json:"votes,omitempty"`
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

	if in["text"] != nil {
		if err := json.Unmarshal(in["text"], &or.Text); err != nil {
			return nil, fmt.Errorf("decoding option text: %w", err)
		}
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

func pollSlideDataFunction(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection, store *projector.SlideStore) (*dbPoll, error) {
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
	return poll, err
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

		if option.ContentObjectID != "" {
			option.ContentObject, err = getTitleInfoFromContentObject(ctx, fetch, store, option.ContentObjectID, "", meetingID)
			if err != nil {
				return nil, fmt.Errorf("getTitleInfoFromContentObject: %w", err)
			}
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

func getPollUser(ctx context.Context, fetch *datastore.Fetcher, userID int) (user *pollUserRepr, err error) {
	data := fetch.Object(ctx, fmt.Sprintf("user/%d", userID), "id", "title", "first_name", "last_name")
	user, err = pollUserFromMap(data)
	if err != nil {
		return nil, fmt.Errorf("get user data: %w", err)
	}

	return user, nil
}

func getVotes(ctx context.Context, fetch *datastore.Fetcher, optionID int) (votes []*pollVoteRepr, err error) {
	voteIDs := datastore.Ints(ctx, fetch.FetchIfExist, "option/%d/vote_ids", optionID)

	fetchFields := []string{
		"id",
		"user_id",
		"value",
	}

	for _, voteID := range voteIDs {
		data := fetch.Object(ctx, fmt.Sprintf("vote/%d", voteID), fetchFields...)
		vote, err := pollVoteFromMap(data)
		if err != nil {
			return nil, fmt.Errorf("get option data: %w", err)
		}

		if vote.UserId != nil {
			vote.User, err = getPollUser(ctx, fetch, *vote.UserId)
			if err != nil {
				return nil, fmt.Errorf("get user data: %w", err)
			}
		}

		votes = append(votes, vote)
	}
	if err := fetch.Err(); err != nil {
		return nil, err
	}

	return votes, nil
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

// Poll renders the poll slide.
func Poll(store *projector.SlideStore) {
	store.RegisterSliderFunc("poll", func(ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection) (encoded []byte, err error) {
		poll, err := pollSlideDataFunction(ctx, fetch, p7on, store)

		if p7on.Type == "poll_single_votes" {
			if err := PollSingleVotes(store, ctx, fetch, p7on, poll); err != nil {
				return nil, fmt.Errorf("adding single votes additional info : %w", err)
			}
		}

		responseValue, err := json.Marshal(poll)
		if err != nil {
			return nil, fmt.Errorf("encoding response slide poll: %w", err)
		}
		return responseValue, err
	})
}

// PollSingleVotes renders the poll_single_votes slide.
func PollSingleVotes(store *projector.SlideStore, ctx context.Context, fetch *datastore.Fetcher, p7on *projector.Projection, poll *dbPoll) error {
	for i, option := range poll.Options {
		votes, err := getVotes(ctx, fetch, *option.id)
		if err != nil {
			return fmt.Errorf("reading option votes: %w", err)
		}

		poll.Options[i].Votes = votes
	}

	if poll.EntitledUsersAtStop != nil {
		var pollUserData []map[string]json.RawMessage
		if err := json.Unmarshal(*poll.EntitledUsersAtStop, &pollUserData); err != nil {
			return fmt.Errorf("reading entitled users")
		}

		var newUserData []map[string]interface{}
		for _, userDate := range pollUserData {
			entry := make(map[string]interface{}, len(userDate))
			for key, val := range userDate {
				if i, err := strconv.ParseInt(string(val), 10, 64); err == nil {
					entry[key] = int(i)
				} else {
					entry[key] = val
				}
			}

			var userID int
			if _, ok := entry["user_merged_into_id"]; ok {
				userID = entry["user_merged_into_id"].(int)
			} else if _, ok := entry["user_id"]; ok {
				userID = entry["user_id"].(int)
			} else {
				continue
			}

			user, err := getPollUser(ctx, fetch, userID)
			if err != nil {
				return fmt.Errorf("encoding entitled users interpretation")
			}

			entry["user"] = user
			newUserData = append(newUserData, entry)
		}

		var pollUserDataJSON, err = json.Marshal(newUserData)
		if err != nil {
			return fmt.Errorf("encoding entitled users interpretation")
		}

		var pollUserDataJSONRaw = json.RawMessage(pollUserDataJSON)
		poll.EntitledUsersAtStop = &pollUserDataJSONRaw
	}

	return nil
}
