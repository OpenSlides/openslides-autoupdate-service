package collection

import (
	"context"
	"fmt"

	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict/perm"
	"github.com/OpenSlides/openslides-autoupdate-service/internal/restrict2/attribute"
	"github.com/OpenSlides/openslides-autoupdate-service/pkg/datastore/dsfetch"
)

// Speaker handels restrictions of the collection speaker.
//
// The user can see a speaker if he has list_of_speakers.can_see or if user_id is the request_user.
//
// Mode A: The user can see the speaker.
type Speaker struct{}

// Name returns the collection name.
func (s Speaker) Name() string {
	return "speaker"
}

// MeetingID returns the meetingID for the object.
func (s Speaker) MeetingID(ctx context.Context, ds *dsfetch.Fetch, id int) (int, bool, error) {
	meetingID, err := ds.Speaker_MeetingID(id).Value(ctx)
	if err != nil {
		return 0, false, fmt.Errorf("get meeting id: %w", err)
	}

	return meetingID, true, nil
}

// Modes returns the restrictions modes for the meeting collection.
func (s Speaker) Modes(mode string) FieldRestricter {
	switch mode {
	case "A":
		return s.see
	}
	return nil
}

func (s Speaker) see(ctx context.Context, fetcher *dsfetch.Fetch, speakerIDs []int) ([]attribute.Func, error) {
	meetingUserID := make([]int, len(speakerIDs))
	meetingID := make([]int, len(speakerIDs))
	for i, id := range speakerIDs {
		if id == 0 {
			continue
		}
		fetcher.Speaker_MeetingUserID(id).Lazy(&meetingUserID[i])
		fetcher.Speaker_MeetingID(id).Lazy(&meetingID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching speaker data: %w", err)
	}

	userID := make([]int, len(speakerIDs))
	for i, id := range meetingUserID {
		if id == 0 {
			continue
		}
		fetcher.MeetingUser_UserID(id).Lazy(&userID[i])
	}

	if err := fetcher.Execute(ctx); err != nil {
		return nil, fmt.Errorf("fetching user: %w", err)
	}

	attr := make([]attribute.Func, len(speakerIDs))
	for i, id := range speakerIDs {
		if id == 0 {
			continue
		}

		groupMap, err := perm.GroupMapFromContext(ctx, fetcher, meetingID[i])
		if err != nil {
			return nil, fmt.Errorf("getting group map: %w", err)
		}

		attr[i] = attribute.FuncOr(
			attribute.FuncUserIDs([]int{userID[i]}),
			attribute.FuncInGroup(groupMap[perm.ListOfSpeakersCanSee]),
		)
	}
	return attr, nil
}
