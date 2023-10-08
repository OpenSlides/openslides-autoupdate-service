package collection

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Poll handels restrictions of the collection poll.
// If the user can see a poll depends on the content object:
//
//	motion: The user can see the linked motion.
//	assignment: The user can see the linked assignment.
//	topic: The user can see the topic.
//
// If the user can manage the poll depends on the content object:
//
//	motion: The user needs motion.can_manage_polls.
//	assignment: The user needs assignment.can_manage.
//	topic: The user needs poll.can_manage.
//
// Mode A: The user can see the poll.
//
// Mode B: Depends on poll/state:
//
//	published: Accessible if the user can see the poll.
//	finished: Accessible if the user can manage the poll.
//	others: Not accessible for anyone.
//
// Mode C: The poll is in the started state and
//
//	the user can manage the poll or
//	the user has the permissions `user.can_see` and `list_of_speakers.can_manage`.
//
// Mode D: Same as Mode B, but for `finished`: Accessible if the user can manage the poll or the user has list_of_speakers.can_manage.
type Poll struct{}

// Name returns the collection name.
func (p Poll) Name() string {
	return "poll"
}

// MeetingID returns the meetingID for the object.
func (p Poll) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Poll_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("getting meetingID: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (p Poll) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return p.see
	case "MANAGE":
		return p.manage // This is a helper to cache the result of p.manage
	case "B":
		return p.modeB
	case "C":
		return p.modeC
	case "D":
		return p.modeD
	}
	return nil
}

func (p Poll) see(ctx context.Context, fetcher *dsfetch.Fetch, pollIDs []int) ([]attribute.Func, error) {
	contentObect := make([]string, len(pollIDs))
	for i, id := range pollIDs {
		fetcher.Poll_ContentObjectID(id).Lazy(&contentObect[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll data: %w", err)
	}

	out := make([]attribute.Func, len(pollIDs))
	for i := range pollIDs {
		collection, rawContentID, found := strings.Cut(contentObect[i], "/")
		if !found {
			return nil, fmt.Errorf("invalid content object id: %s", contentObect[i])
		}

		contentID, err := strconv.Atoi(rawContentID)
		if err != nil {
			return nil, fmt.Errorf("invalid id part of content object id: %w", err)
		}

		var mode FieldRestricter
		switch collection {
		case "motion":
			mode = Collection(ctx, "motion").Modes("C")
		case "assignment":
			mode = Collection(ctx, "assignment").Modes("A")
		case "topic":
			mode = Collection(ctx, "topic").Modes("A")
		default:
			return nil, fmt.Errorf("invalid collection %s", collection)
		}

		contentAttr, err := mode(ctx, fetcher, []int{contentID})
		if err != nil {
			return nil, fmt.Errorf("checking content object %s: %w", contentObect[i], err)
		}

		out[i] = contentAttr[0]
	}

	return out, nil
}

func (p Poll) manage(ctx context.Context, fetcher *dsfetch.Fetch, pollIDs []int) ([]attribute.Func, error) {
	contentObect := make([]string, len(pollIDs))
	meetingID := make([]int, len(pollIDs))
	for i, id := range pollIDs {
		fetcher.Poll_ContentObjectID(id).Lazy(&contentObect[i])
		fetcher.Poll_MeetingID(id).Lazy(&meetingID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll data: %w", err)
	}

	out := make([]attribute.Func, len(pollIDs))
	for i := range pollIDs {
		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		collection, _, found := strings.Cut(contentObect[i], "/")
		if !found {
			return nil, fmt.Errorf("invalid content object id: %s", contentObect[i])
		}

		var managePerm perm.TPermission
		switch collection {
		case "motion":
			managePerm = perm.MotionCanManagePolls
		case "assignment":
			managePerm = perm.AssignmentCanManage
		case "topic":
			managePerm = perm.PollCanManage
		default:
			return nil, fmt.Errorf("invalid collection %s", collection)
		}

		out[i] = attribute.FuncInGroup(groupMap[managePerm])
	}

	return out, nil
}

func (p Poll) modeB(ctx context.Context, fetcher *dsfetch.Fetch, pollIDs []int) ([]attribute.Func, error) {
	state := make([]string, len(pollIDs))
	for i, id := range pollIDs {
		fetcher.Poll_State(id).Lazy(&state[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll data: %w", err)
	}

	out := make([]attribute.Func, len(pollIDs))
	// TODO: try to group this
	for i, pollID := range pollIDs {
		switch state[i] {
		case "published":
			af, err := Collection(ctx, p.Name()).Modes("A")(ctx, fetcher, []int{pollID})
			if err != nil {
				return nil, fmt.Errorf("checking poll.see: %w", err)
			}
			out[i] = af[0]

		case "finished":
			af, err := Collection(ctx, p.Name()).Modes("MANAGE")(ctx, fetcher, []int{pollID})
			if err != nil {
				return nil, fmt.Errorf("checking poll.see: %w", err)
			}
			out[i] = af[0]

		default:
			out[i] = attribute.FuncNotAllowed
		}
	}
	return out, nil
}

func (p Poll) modeC(ctx context.Context, fetcher *dsfetch.Fetch, pollIDs []int) ([]attribute.Func, error) {
	state := make([]string, len(pollIDs))
	meetingID := make([]int, len(pollIDs))
	for i, id := range pollIDs {
		fetcher.Poll_State(id).Lazy(&state[i])
		fetcher.Poll_MeetingID(id).Lazy(&meetingID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll data: %w", err)
	}

	out := make([]attribute.Func, len(pollIDs))
	for i, pollID := range pollIDs {
		if state[i] != "started" {
			out[i] = attribute.FuncNotAllowed
			continue
		}

		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		canManage, err := Collection(ctx, p.Name()).Modes("MANAGE")(ctx, fetcher, []int{pollID})
		if err != nil {
			return nil, fmt.Errorf("checking poll.see: %w", err)
		}

		out[i] = attribute.FuncOr(
			canManage[0],
			attribute.FuncAnd(
				attribute.FuncInGroup(groupMap[perm.UserCanSee]),
				attribute.FuncInGroup(groupMap[perm.ListOfSpeakersCanManage]),
			),
		)
	}

	return out, nil
}

func (p Poll) modeD(ctx context.Context, fetcher *dsfetch.Fetch, pollIDs []int) ([]attribute.Func, error) {
	// Mode D: Same as Mode B, but for `finished`: Accessible if the user can manage the poll or the user has list_of_speakers.can_manage.
	state := make([]string, len(pollIDs))
	meetingID := make([]int, len(pollIDs))
	for i, id := range pollIDs {
		fetcher.Poll_State(id).Lazy(&state[i])
		fetcher.Poll_MeetingID(id).Lazy(&meetingID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching poll data: %w", err)
	}

	out := make([]attribute.Func, len(pollIDs))
	// TODO: try to group this
	for i, pollID := range pollIDs {
		switch state[i] {
		case "published":
			af, err := Collection(ctx, p.Name()).Modes("A")(ctx, fetcher, []int{pollID})
			if err != nil {
				return nil, fmt.Errorf("checking poll.see: %w", err)
			}
			out[i] = af[0]

		case "finished":
			groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
			if err != nil {
				return nil, fmt.Errorf("getting group map: %w", err)
			}

			canManage, err := Collection(ctx, p.Name()).Modes("MANAGE")(ctx, fetcher, []int{pollID})
			if err != nil {
				return nil, fmt.Errorf("checking poll.see: %w", err)
			}
			out[i] = attribute.FuncOr(
				canManage[0],
				attribute.FuncInGroup(groupMap[perm.ListOfSpeakersCanManage]),
			)

		default:
			out[i] = attribute.FuncNotAllowed
		}
	}
	return out, nil
}
